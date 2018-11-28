package main

import (
	"time"
)

type Message struct {
	Id      string `gorethink:"id,omitempty"`
	Text    string
	Created time.Time
}

func NewMassage(text string) *Message {
	return &Message{
		Text: text,
	}
}
