package main


import (
    "fmt"
    "github.com/bwmarrin/discordgo"
    "io/ioutil"
    "./Handlers"
)

var me *discordgo.User
func main(){
    fmt.Println("Bot is alive!")
    setupDiscord()
    lock := make(chan struct{})
    <- lock
}

func getBotToken() string{
    // get the bot token from the botdata file
    byteContent, err := ioutil.ReadFile("./resources/botdata.txt")
    if err != nil {
        panic(err)
    }
    return string(byteContent)
}

func setupDiscord(){
    token := getBotToken()
    discord, err := discordgo.New("Bot " + token)
    if err != nil{
        panic(err)
    }

    err = discord.Open()
    if err != nil {
        panic(err)
    }

    // add command handler
    discord.AddHandler(Handlers.CommandHandler)

    me, err = discord.User("@me")
    if err != nil{
        panic(err)
    }
}