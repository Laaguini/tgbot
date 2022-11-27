# tgbot

tgbot is a simple lightweight no-dependencies library for [Telegram Bot API](https://core.telegram.org/bots/api)

## Overview

- [Initialization](#initialization)
- [Handling text messages](#handling-text-messages)
- [Handling stickers](#handling-stickers)

## Usage

### Initialization

```go
import "github.com/Laaguini/tgbot"

func main() {
    bot := &tgbot.New{
        Token:           "TOKEN",
        PollingInterval: 500 * time.Millisecond,
    }
    
    bot.Poll()
}
```

### Handling text messages

```go
// Any text message 
bot.OnAnyMessage(func(message tgbot.Message) {
    fmt.Printf("%s From: %s Message: %s \n", time.Now().Format("15:51:15"), message.From.Username, message.Text)
})
    
// Message that starts from '/', e.g '/start'
bot.OnCommand("start", func(message tgbot.Message) {
    bot.Send(message.Chat.Id, "Hi, "+message.From.FirstName+" !")
})
    
// Handling not a command
bot.OnMessage(func(message tgbot.Message) {
    bot.Send(message.Chat.Id, message.Text)
})
```

### Handling stickers

```go
bot.OnSticker(func(message tgbot.Message) {
    bot.Send(message.Chat.Id, message.Sticker.Emoji)
})
```
