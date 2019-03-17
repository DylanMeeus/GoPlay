package main

import "fmt"

func main() {

    fmt.Printf("%v\n", fib(3))
}

// return the n'th fibonacci number
func fib(N int) int {
    a, b := 0, 1
    nums := 1
    for nums < N {
        c := a + b
        a = b
        b = c
        nums++
    }
    return b
}
