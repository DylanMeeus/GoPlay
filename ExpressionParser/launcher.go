package main

import (
    "fmt"
    "./Infix"
)
func main(){
    res := InfixParser.Eval("(1 - 200) * 5")
    fmt.Println(res)
}