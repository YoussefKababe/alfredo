package main

import (
	"alfredo"
	"alfredo/config"
	"alfredo/dropbox"
	"alfredo/firebase"
	"alfredo/messenger"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	config.Initialize()
	go messenger.SetGetStartedButton()
	go messenger.SetGreetingText()
	e.GET("/webhook", verify)
	e.POST("/webhook", receive)
	e.GET("/mdropbox", linkDropbox)
	e.Logger.Fatal(e.Start(":" + config.AppPort))
}

func verify(c echo.Context) error {
	mode, token := c.QueryParam("hub.mode"), c.QueryParam("hub.verify_token")
	if mode == "subscribe" && token == config.VerifyToken {
		fmt.Println("Validating webhook...")
		return c.String(http.StatusOK, c.QueryParam("hub.challenge"))
	}

	log.Panic("Failed validation. Make sure the validation tokens match.")
	return c.NoContent(http.StatusForbidden)
}

func receive(c echo.Context) error {
	call := new(messenger.Call)
	c.Bind(call)

	for _, entry := range call.Entries {
		for _, event := range entry.Events {
			if event.Message != nil {
				alfredo.HandleMessage(event)
			} else if event.Postback != nil {
				alfredo.HandlePostback(event)
			} else {
				fmt.Println("Webhook received unknown event:", event)
			}
		}
	}

	return c.NoContent(http.StatusOK)
}

func linkDropbox(c echo.Context) error {
	fb := firebase.New(config.FirebaseProjectID, config.FirebaseSecret)
	db := dropbox.New(config.DropboxKey, config.DropboxSecret)
	code := c.QueryParam("code")
	userID := c.QueryParam("state")

	token := db.GetAuthToken(code)

	fb.SaveUser(userID, token)
	messenger.SendText("Awesome! Your Dropbox account is now linked. Send or forward"+
		" any file to me and I'll instantly save it to your Dropbox!", userID)
	return c.String(200, "You're Dropbox account was successfully linked! You"+
		" can close this tab and go back to messenger.")
}
