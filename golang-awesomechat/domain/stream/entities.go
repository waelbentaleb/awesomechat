package stream

import "time"

// Stream represent a user message channel uses to stream messages
type Stream struct {
	Username       string
	MessageChannel chan Message
}

// Message represent a user received message
type Message struct {
	Sender    string
	Content   string
	Type      string
	GroupName string
	Date      time.Time
}
