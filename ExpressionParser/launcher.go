package main

import (
    "fmt"
    "./Infix"
)
func main(){
    fmt.Println("Hello World!")
    res := InfixParser.Eval("1 + 200")
    fmt.Println(res)
}