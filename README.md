![](https://github.com/Laaguini/tgbot/blob/main/readme-cover.png?raw=true)

# tgbot

tgbot is a simple lightweight no-dependencies library for [Telegram Bot API](https://core.telegram.org/bots/api)

## Overview

- [Initialization](#initialization)
- [Getting bot info](#getting-bot-info)
- [Handling text messages](#handling-text-messages)
- [Handling stickers](#handling-stickers)
<<<<<<< HEAD
=======
- [Handling files](#handling-files)
>>>>>>> parent of 8fae427 (let's do it again)

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

### Getting bot info

```go
// Info as JSON string
fmt.Println(bot.Info())
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
<<<<<<< HEAD
=======

### Handling files

```go
// Donwnloading voice message
bot.OnVoice(func(message tgbot.Message) {
    voice := message.Voice
    fileReader := bot.GetFile(voice.Id)

    out, err := os.Create("./files/" + voice.Id + voice.Extensions[0])
    if err != nil {
        fmt.Println(err)
    }
    defer out.Close()
    defer fileReader.Close()

    _, err = io.Copy(out, fileReader)
    if err ! nil {
        fmt.Println(err)
    }
})

// Logging audio messsage info 
bot.OnAudio(func(message tgbot.Message) {
    audio := message.Audio	
    fmt.Println(audio.Name + " " + strconv.Itoa(audio.Duration) + " " + strconv.Itoa(audio.Size))
})
```
>>>>>>> parent of 8fae427 (let's do it again)
