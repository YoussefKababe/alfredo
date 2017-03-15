package main

import (
	"alfredo/config"
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
	config.LoadEnvVars()
	go messenger.SetGetStartedButton()
	go messenger.SetGreetingText()
	e.GET("/webhook", verify)
	e.POST("/webhook", receive)
	e.GET("/mdropbox", messenger.LinkDropbox)
	e.Logger.Fatal(e.Start(":8080"))
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
				messenger.HandleMessage(event)
			} else if event.Postback != nil {
				messenger.HandlePostback(event)
			} else {
				fmt.Println("Webhook received unknown event:", event)
			}
		}
	}

	return c.NoContent(http.StatusOK)
}
