package tgbot

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Photosize struct {
	Width  int
	Height int
	Size   int `json:"file_size"`
}

type Sticker struct {
	Id       string `json:"file_id"`
	Type     string
	Thumb    Photosize
	Emoji    string
	SetName  string `json:"set_name"`
	FileSize int    `json:"file_size"`
}

type User struct {
	Id           int
	IsBot        bool
	FirstName    string `json:"first_name"`
	Username     string
	LanguageCode string
}

type Chat struct {
	Id        int
	FirstName string `json:"first_name"`
	Username  string
	Type      string
}

type Message struct {
	Id                 int `json:"message_id"`
	ThreadId           int `json:"message_thread_id"`
	From               User
	ForwardFrom        User
	ForwardFromChat    User
	ForwardedMessageId int `json:"forward_from_message_id"`
	ForwardDate        int
	SenderChat         Chat
	Chat               Chat
	Date               int
	Text               string
	ReplyToMessage     *Message
	ViaBot             User
	EditDate           int
	Sticker            Sticker
}

type Update struct {
	Id                int `json:"update_id"`
	Message           *Message
	EditedMessage     *Message
	ChannelPost       *Message
	EditedChannelPost *Message
}

type UpdateResult struct {
	Ok      bool
	Updates []Update `json:"result"`
}

func (i *New) GetUpdates(offset int) []Update {
	res := i.get(GetArgs{Action: "getUpdates", Params: []Param{{Name: "offset", Value: strconv.Itoa(offset)}}})
	updateResult := UpdateResult{}
	err := json.Unmarshal([]byte(res), &updateResult)

	if err != nil {
		fmt.Println(err)
	}
	if !updateResult.Ok {
		fmt.Println("An error occured with update result parsing")
	}

	return updateResult.Updates
}
