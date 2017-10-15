package main

import (
    "fmt"
    "strings"
)


var numerals = map[string] int{
    "I" : 1,
    "V" : 5,
    "X" : 10,
    "L" : 50,
    "C" : 100,
    "D" : 500,
    "M" : 1000,
}

func main(){
    fmt.Println(parseNumeral("VI"))
    fmt.Println(parseNumeral("LXX"))
    fmt.Println(parseNumeral("MCC"))
    fmt.Println(parseNumeral("IV"))
    fmt.Println(parseNumeral("XC"))
    fmt.Println(parseNumeral("CM"))
    fmt.Println(parseNumeral("XCV"))
    fmt.Println(parseNumeral("XCIX"))
    fmt.Println(parseNumeral("XXX"))
}


func parseNumeral(numeral string) int{
    tokens := tokenize(numeral)

    values := make([]int, 0)
    for i := 0; i < len(tokens); i++{
        literals := strings.Split(tokens[i],"") // string of identical chars
        value := numerals[literals[0]]
        // should the value be a negative?, e.g is the next token group of higher importance?
        if i+1 < len(tokens) {
            nextliteral := strings.Split(tokens[i+1],"")
            nextvalue := numerals[nextliteral[0]]
            if nextvalue > value{
                value *= -1
            }
        }
        values = append(values,value * len(literals))
    }

    sum := 0
    for i := 0; i < len(values); i++{
        sum += values[i]
    }
    //fmt.Println(values)
    return sum
}



// group the tokens by character
func tokenize(numeral string) []string{
    tokens := strings.Split(numeral, "")
    groups := make([]string,0)

    for i := 0; i < len(tokens); i++{
        token := tokens[i]
        for i+1 < len(tokens) && tokens[i] == tokens[i+1]{
            token += tokens[i]
            i++
        }
        groups = append(groups, token)
    }
    return groups
}