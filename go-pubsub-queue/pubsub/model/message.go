package model

type Message struct {
	msg string
}

// NewMessage creates a new message
func NewMessage(msg string) *Message {
	return &Message{
		msg: msg,
	}
}

// GetMsg returns the message content
func (m *Message) GetMsg() string {
	return m.msg
}
