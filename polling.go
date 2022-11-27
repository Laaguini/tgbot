package tgbot

import (
	"fmt"
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
					if update.Message.Text != "" {
						i.HandleMessage(*update.Message)
					}
					if update.Message.Sticker.Id != "" {
						i.HandleSticker(*update.Message)
					}
				}
				offset = lastUpdate.Id
			}
		}
	}
}
