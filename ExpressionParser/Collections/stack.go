package collections


type Stack []interface{}


func (s *Stack) Push(i interface{}){
    *s = append(*s,i)
}

func (s *Stack) Peek() interface{} {
    return (*s)[len(*s)-1]
}

func (s *Stack) Pop() interface{}{
    if len(*s) == 0{
        panic("Can not pop from an empty stack!")
    }
    last := (*s).Peek()
    (*s) = (*s)[0:len(*s)-1]
    return last
}

func (s *Stack) Empty() bool {
    return len(*s) == 0
}