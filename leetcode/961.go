package main

import "fmt"

func main() { 
    arr := []int{5,1,5,2,5,3,5,4}
    fmt.Printf("%v\n", repeatedNTimes(arr))
}

func repeatedNTimes(arr []int) int {
    n := len(arr) / 2
    m := make(map[int]int, 0)
    for _,el := range arr {
        m[el]++
    }
    for k,v := range m {
        if v == n {
            return k
        }
    }
    return 0
}
