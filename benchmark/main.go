package main

import "fmt"

//var empty = struct{}{}

func main() {
    numbers := []int{1,2,3,4,5,5,6,7,8,9,19,1,2,5,76,53}
    f := make(map[int]*struct{},0)
    for _,n := range numbers {
        if _,ok := f[n]; !ok {
           f[n] = &empty 
        }
    }
    for k,_ := range f {
        fmt.Printf("%v ", k)
    }
}

