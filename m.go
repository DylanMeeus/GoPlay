package main

import "fmt"

func main() {
    fmt.Println(struct{a,_,c int}{1,2,3})
}
