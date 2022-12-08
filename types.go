package tgbot

type Photosize struct {
	Width  int
	Height int
	Size   int `json:"file_size"`
}

type File struct {
	Id       string `json:"file_id"`
	UniqueId string `json:"file_unique_id"`
	Size     int    `json:"file_size"`
	Path     string `json:"file_path"`
}

type Audio struct {
	Id         string `json:"file_id"`
	UniqueId   string `json:"file_unique_id"`
	Duration   int
	Performer  string
	Title      string
	Name       string `json:"file_name"`
	MimeType   string `json:"mime_type"`
	Size       int    `json:"file_size"`
	Thumb      Photosize
	Extensions []string
}

type Document struct {
	Id       string `json:"file_id"`
	UniqueId string `json:"file_unique_id"`
	Name     string `json:"file_name"`
	MimeType string `json:"mime_type"`
	Size     int    `json:"file_size"`
	Thumb    Photosize
}

type Video struct {
	Id         string `json:"file_id"`
	UniqueId   string `json:"file_unique_id"`
	Width      int
	Height     int
	Duration   int
	Name       string `json:"file_name"`
	MimeType   string `json:"mime_type"`
	Size       int    `json:"file_size"`
	Thumb      Photosize
	Extensions []string
}

type VideoNote struct {
	Id       string `json:"file_id"`
	UniqueId string `json:"file_unique_id"`
	Length   int
	Duration int
	Size     int `json:"file_size"`
	Thumb    Photosize
}

type Voice struct {
	Id         string `json:"file_id"`
	UniqueId   string `json:"file_unique_id"`
	Duration   int
	Size       int `json:"file_size"`
	Thumb      Photosize
	MimeType   string `json:"mime_type"`
	Extensions []string
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
	Audio              Audio
	Document           Document
	Video              Video
	VideoNote          VideoNote
	Voice              Voice
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
