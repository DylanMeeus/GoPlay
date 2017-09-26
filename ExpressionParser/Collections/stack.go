package collections


type Stack []interface{}


func (s *Stack) push(i interface{}){
    *s = append(*s,i)
}

func (s *Stack) peek() interface{} {
    return (*s)[len(*s)-1]
}

func (s *Stack) pop() interface{}{
    if len(*s) == 0{
        panic("Can not pop from an empty stack!")
    }
    last := (*s).peek()
    (*s) = (*s)[0:len(*s)-1]
    return last
}