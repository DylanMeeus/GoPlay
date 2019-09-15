package main

import "fmt"

func main() {
    s1 := []int{1,2,3,5,6,7,8,9}
    s2 := []int{4,5,6}
    ch := make(chan int)
    go sum(s1, ch)
    go sum(s2, ch)
    x, y := <-ch, <-ch
    fmt.Println(x, y)
}

func sum(is []int, c chan int) {
    res := 0
    for _, i := range is {
	res += i
    }
    c <-res 
}


