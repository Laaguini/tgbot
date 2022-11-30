package tgbot

import (
	"fmt"
	"mime"
	"time"
)

func (i *New) Poll() {
	ticker := time.NewTicker(i.PollingInterval)
	defer ticker.Stop()

	done := make(chan bool)

	var offset int
	if upd := i.GetUpdates(0); len(upd) > 0 {
		offset = upd[len(upd)-1].Id
	} else {
		offset = 0
	}

	for {
		select {
		case <-done:
			fmt.Println("Done")
			return
		case _ = <-ticker.C:
			updates := i.GetUpdates(offset)

			if len(updates) == 0 {
				continue
			}

			lastUpdate := updates[len(updates)-1]

			if offset != lastUpdate.Id {
				for _, update := range updates[1:] {
					message := update.Message
					if message.Text != "" {
						i.HandleMessage(*message)
					}
					if message.Sticker.Id != "" {
						i.HandleSticker(*message)
					}
					if message.Audio.Id != "" {
						i.HandleAudio(*message)
					}
					if message.Video.Id != "" {
						i.HandleVideo(*message)
					}
					if message.Document.Id != "" {
						i.HandleDocument(*message)
					}
					if message.VideoNote.Id != "" {
						i.HandleVideoNote(*message)
					}
					if message.Voice.Id != "" {
						file_exts, _ := mime.ExtensionsByType(message.Voice.MimeType)
						message.Voice.Extensions = file_exts
						i.HandleVoice(*message)
					}
				}
				offset = lastUpdate.Id
			}
		}
	}
}
