package main


import (
    "fmt"
)

type Stack []int

func main(){
    // for testing
    stack := Stack{}
    stack.push(4)
    stack.push(5)
    stack.pop()
    fmt.Println(stack)
}

func (s *Stack) push(i int){
    *s = append(*s,i)
}

func (s *Stack) pop() int{
    if len(*s) == 0{
        panic("Can not pop from an empty stack!")
    }
    last := (*s)[len(*s)-1]
    (*s) = (*s)[0:len(*s)-1]
    return last
}