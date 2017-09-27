package Handlers

import (
    "fmt"
    "github.com/bwmarrin/discordgo"
    "strings"
    "../../Infix"
    "strconv"
)


var prefix string = "#calc"

func CommandHandler(session *discordgo.Session, message *discordgo.MessageCreate){
    fmt.Println("Parsing command")

    messageContent := strings.TrimSpace(message.Content)

    if strings.HasPrefix(messageContent,prefix){
        equation := messageContent[len(prefix):]
        result := InfixParser.Eval(equation)
        output := strconv.FormatFloat(result,'f',-1,64)
        _,_ = session.ChannelMessageSend(message.ChannelID, output)
    }

}
