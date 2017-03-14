package main

import (
	"alfredo/config"
	"alfredo/message"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.LoadEnvVars()
	message.SetGetStartedButton()
	e.GET("/webhook", verify)
	e.POST("/webhook", receive)
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
	call := new(message.Call)
	c.Bind(call)

	for _, entry := range call.Entries {
		for _, event := range entry.Events {
			if event.Message != nil {
				message.HandleMessage(event)
			} else {
				fmt.Println("Webhook received unknown event:", event)
			}
		}
	}

	return c.NoContent(http.StatusOK)
}
