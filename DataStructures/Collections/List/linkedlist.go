package main


// implementation of a linked list in go

import (
    "fmt"
    "strconv"
)

// CODE FOR THE NODE
type Node struct{
    value interface{}
    next *Node
}

func (node Node) String() string {
    return strconv.Itoa(node.value.(int)) // todo: generalize this for any type.
}

// CODE FOR THE LIST
type LinkedList struct{
    size int
    start *Node
}

func (linkedList *LinkedList) add(value interface{}){
    if linkedList.start == nil {
        linkedList.start = &Node{value: value, next: nil}
    } else {
        // get the last node and add the new node
        newNode := Node{value: value, next : nil}
        startNode := linkedList.start
        for startNode.next != nil{
            startNode = startNode.next
        }
        // now startNode is the last node that does not have a _next_ node.
        startNode.next = &newNode
    }
}

// CODE TO TEST THE LIST

func main(){
    // todo: move this away from main after the collection is done

    list := LinkedList{}
    list.add(1)
    list.add(2)
    list.add(3)
    list.add(5)


    start := list.start

    for start != nil {
        fmt.Println(start)
        start = start.next
    }

    //fmt.Println(list.start)

}
