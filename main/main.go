package main

import (
  "os"
  "fmt"
  "log"
  "bytes"
  "net/http"
  "encoding/json"
  "github.com/labstack/echo"
  "github.com/joho/godotenv"
)

var verifyToken string
var pageToken string

func main() {
  e := echo.New()
  if err := godotenv.Load(); err != nil {
    log.Fatal(err)
  }
  verifyToken = os.Getenv("VERIFY_TOKEN")
  pageToken = os.Getenv("PAGE_TOKEN")
  e.GET("/webhook", verify)
  e.POST("/webhook", receive)
  e.Logger.Fatal(e.Start(":8080"))
}

func verify(c echo.Context) error {
  mode, token := c.QueryParam("hub.mode"), c.QueryParam("hub.verify_token")
  if mode == "subscribe" && token == verifyToken {
    fmt.Println("Validating webhook...")
    return c.String(http.StatusOK, c.QueryParam("hub.challenge"))
  } else {
    log.Panic("Failed validation. Make sure the validation tokens match.")
    return c.NoContent(http.StatusForbidden)
  }
}

type Call struct {
  Object string `json:"object"`
  Entries []*Entry `json:"entry"`
}

type Entry struct {
  Id string `json:"id"`
  Time uint `json:"time"`
  Events []*Event `json:"messaging"`
}

type Event struct {
  Message *Message `json:"message"`
  Sender *Person `json:"sender"`
  Recipient *Person `json:"recipient"`
  Timestamp uint `json:"timestamp"`
}

type Message struct {
  Text string `json:"text"`
}

type Person struct {
  Id string `json:"id"`
}

func receive(c echo.Context) error {
  call := new(Call)
  c.Bind(call)

  for _, entry := range call.Entries {
    for _, event := range entry.Events {
      if event.Message != nil {
        receivedMessage(event)
      } else {
        fmt.Println("Webhook received unknown event:", event)
      }
    }
  }

  return c.NoContent(http.StatusOK)
}

func receivedMessage(event *Event) {
  senderId := event.Sender.Id
  message := event.Message

  newMessage := Event{
    Message: &Message{
      Text: message.Text,
    },
    Recipient: &Person{
      Id: senderId,
    },
  }

  mm, _ := json.Marshal(newMessage)

  request, _ := http.NewRequest("POST", "https://graph.facebook.com/v2.6/me/messages?access_token=" + pageToken, bytes.NewBuffer(mm))
  request.Header.Set("Content-Type", "application/json")

  client := http.Client{}
  client.Do(request)
}
