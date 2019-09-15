package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
    f, err := ioutil.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    lines := strings.Split(string(f), "\n")
    ptr := 0
    doctors := -1
    for ptr < len(lines) {
        if doctors == -1 {
            doctors = 
        }
        ptr += 1
    }
}
