package tgbot

type CommandHandler struct {
	Command string
	Handler func(message Message)
}

type Handlers struct {
	message    []func(message Message)
	anyMessage []func(message Message)
	command    []CommandHandler
	sticker    []func(message Message)
	audio      []func(message Message)
	video      []func(message Message)
	document   []func(message Message)
	videoNote  []func(message Message)
	voice      []func(message Message)
}

func (i *New) OnAnyMessage(handler func(message Message)) {
	i.Handlers.anyMessage = append(i.Handlers.anyMessage, handler)
}

func (i *New) OnMessage(handler func(message Message)) {
	i.Handlers.message = append(i.Handlers.message, handler)
}

func (i *New) OnCommand(command string, handler func(message Message)) {
	i.Handlers.command = append(i.Handlers.command, CommandHandler{Command: command, Handler: handler})
}

func (i *New) OnSticker(handler func(message Message)) {
	i.Handlers.sticker = append(i.Handlers.sticker, handler)
}

func (i *New) OnAudio(handler func(message Message)) {
	i.Handlers.audio = append(i.Handlers.audio, handler)
}

func (i *New) OnVideo(handler func(message Message)) {
	i.Handlers.video = append(i.Handlers.video, handler)
}

func (i *New) OnDocument(handler func(message Message)) {
	i.Handlers.document = append(i.Handlers.document, handler)
}

func (i *New) OnVideoNote(handler func(message Message)) {
	i.Handlers.videoNote = append(i.Handlers.videoNote, handler)
}

func (i *New) OnVoice(handler func(message Message)) {
	i.Handlers.voice = append(i.Handlers.voice, handler)
}

func (i *New) HandleMessage(message Message) {
	text := message.Text
	isCommand := text[0] == '/'

	for _, h := range i.Handlers.anyMessage {
		h(message)
	}

	if isCommand {
		for _, ch := range i.Handlers.command {
			if ch.Command == text[1:] {
				ch.Handler(message)
			}
		}
	} else {
		for _, h := range i.Handlers.message {
			h(message)
		}
	}
}

func (i *New) HandleSticker(message Message) {
	for _, h := range i.Handlers.sticker {
		h(message)
	}
}

func (i *New) HandleAudio(message Message) {
	for _, h := range i.Handlers.audio {
		h(message)
	}
}

func (i *New) HandleVideo(message Message) {
	for _, h := range i.Handlers.video {
		h(message)
	}
}

func (i *New) HandleDocument(message Message) {
	for _, h := range i.Handlers.document {
		h(message)
	}
}

func (i *New) HandleVideoNote(message Message) {
	for _, h := range i.Handlers.videoNote {
		h(message)
	}
}

func (i *New) HandleVoice(message Message) {
	for _, h := range i.Handlers.voice {
		h(message)
	}
}
