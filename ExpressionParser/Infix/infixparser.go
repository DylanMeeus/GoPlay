// Infix parser for the Shunting-yard algorithm

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
    return strings.Map(func(r rune) rune{
        if r == 32{
            return -1
        }
        return r
    },input)
}

// evaluate a string
func eval(input string) int{

    t := "0123456789"

    for i := 0; i < len(t); i++{
        fmt.Println(t[i])
    }

    index := 0
    input = formatInputString(input)
    fmt.Println(input)
    tokens := []string{}
    for index < len(input) {
        token := ""
        charvalue := input[index]
        for charvalue >= 48 && charvalue <= 57 && index+1 < len(input){
            // it is a number
            token += string(charvalue)
            index++
            charvalue = input[index]
        }

        if token == ""{ // we did not encounter a number
            token = string(charvalue)
            index++
        }
        tokens = append(tokens, token)
    }

    fmt.Println(tokens)
    return 0
}

