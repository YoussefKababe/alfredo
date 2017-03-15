package message

// Call represents a call from the messenger platform.
type Call struct {
	Object  string   `json:"object"`
	Entries []*Entry `json:"entry"`
}

// Entry represents an entry from the messenger platform call.
type Entry struct {
	ID     string   `json:"id"`
	Time   uint     `json:"time"`
	Events []*Event `json:"messaging"`
}

// Event represents an event from an Entry.
type Event struct {
	Message   *Message  `json:"message"`
	Postback  *Postback `json:"postback"`
	Sender    *Person   `json:"sender"`
	Recipient *Person   `json:"recipient"`
	Timestamp uint      `json:"timestamp"`
}

// Message represents a message from an Event.
type Message struct {
	Text        string        `json:"text"`
	Attachments []*Attachment `json:"attachments"`
}

// Attachment represents an attachment from a Message.
type Attachment struct {
	Type    string            `json:"type"`
	Payload map[string]string `json:"payload"`
}

// Postback represents a postback from a Message.
type Postback struct {
	Payload string `json:"payload"`
}

// Person represents a sender or a recipient from an Event.
type Person struct {
	ID string `json:"id"`
}
