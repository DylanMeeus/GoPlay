// pascal triangle implementation in Go
package main

import (
    "fmt"
)


func main() {
    fmt.Printf("%v\n", generate(5))
}


func generate(levels int) [][]int {
    ct := []int{}
    ts := [][]int{}
    for i := 0; i < levels; i++ {
        ct = nextLevel(ct)
        ts = append(ts, ct)
    }
    return ts
}

func nextLevel(current []int) (next []int) {
    next = []int{1}
    if len(current) == 0 {
        return next 
    }
    for i := 0; i < len(current) - 1; i++ {
            next = append(next, current[i] + current[i+1])
    }
    return append(next, 1)
}
