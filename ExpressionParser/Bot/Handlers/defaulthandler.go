package Handlers

import (
    "github.com/bwmarrin/discordgo"
    "strings"
    "../../Infix"
    "strconv"
    "../../../RomanNumerals"
)


// prefixes
var calcprefix string = "#calc" // calculate prefix
var convprefix string = "#conv" // convert prefix

// handle the incoming message
func CommandHandler(session *discordgo.Session, message *discordgo.MessageCreate){
    messageContent := strings.TrimSpace(message.Content)

    if strings.HasPrefix(messageContent, calcprefix){
        equation := messageContent[len(calcprefix):]
        result := InfixParser.Eval(equation)
        output := strconv.FormatFloat(result,'f',-1,64)
        _,_ = session.ChannelMessageSend(message.ChannelID, output)
    } else if strings.HasPrefix(messageContent, convprefix) {
        input := messageContent[len(convprefix):]
        decimal := romannumerals.ParseNumeral(input)
        output := strconv.Itoa(decimal)
        _,_ = session.ChannelMessageSend(message.ChannelID, output)
    }

}
