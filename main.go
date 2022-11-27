package tgbot

import (
	"time"
)

type New struct {
	Token           string
	PollingInterval time.Duration
	Handlers        Handlers
}

func (i *New) Info() string {
	return i.get(GetArgs{Action: "getMe"})
}

func (i *New) Send(chatId int, text string) {
	type Body struct {
		ChatId int    `json:"chat_id"`
		Text   string `json:"text"`
	}

	i.post(&PostArgs{
		Action: "sendMessage",
		Body: Body{
			ChatId: chatId,
			Text:   text,
		},
	})
}
