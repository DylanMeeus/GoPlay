// Infix parser using the Shunting-yard algorithm

package main

import (
    "fmt"
    "strings"
)

func main(){
    eval("13+250 - 2")
}


// remove some of the spaces - that is the only cleanup at the moment
func formatInputString(input string) string{
    spacerune := 32
    return strings.Map(func(r rune) rune{
        if r == spacerune{
            return -1
        }
        return r
    },input)
}

func tokenize(input string) []string{
    tokens := []string{}
    index := 0
    for index < len(input) {
        token := ""
        charvalue := input[index]
        for charIsDigit(charvalue) && index+1 < len(input){ // also make sure it is in the bounds
            token += string(charvalue)
            index++
            charvalue = input[index]
        }

        if token == ""{
            token = string(charvalue)
            index++
        }
        tokens = append(tokens, token)
    }
    return tokens
}

// evaluate a string
func eval(input string) int{
    input = formatInputString(input)
    tokens := tokenize(input)
    fmt.Println(tokens)
    return 0
}


func charIsDigit(char byte) bool{
    return char >= 48 && char <= 57
}
