package model

type PushMessage struct {
	To       string    `json:"to"`
	Messages []Message `json:"messages"`
}
