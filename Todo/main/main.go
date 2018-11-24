// todo application to test a stack
package main

import (
    "fmt"
    "../todo"
)

func main() {
    t := todo.Todo{
        Id: 0,
        Description: "Hello World",
    }
    fmt.Printf("%v\n", t)
}
