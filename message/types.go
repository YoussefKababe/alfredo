package message

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
