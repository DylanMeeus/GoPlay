package main

import (
    "fmt"
)

func calculate(base int) (out int) {
    out = base
    raise := func() {
        out *= out 
    }

    double := func() {
        out *= 2
    }
    defer raise()
    defer double()
    out += 2 
    return out 
}

func main() {
    fmt.Printf("%v\n",calculate(1))
}
