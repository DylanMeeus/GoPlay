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

/*
    Create the first node in the list, or iterate until the last node and insert it there.
 */
func (linkedList *LinkedList) add(value interface{}){
    if linkedList.start == nil {
        linkedList.start = &Node{value: value, next: nil}
        linkedList.size += 1
    } else {
        newNode := Node{value: value, next : nil}
        startNode := linkedList.start
        for startNode.next != nil{
            startNode = startNode.next
        }
        // now startNode is the last node that does not have a _next_ node.
        startNode.next = &newNode
        linkedList.size += 1
    }
}

/*
    Return the n'th node in the LinkedList
 */
func (linkedList *LinkedList) get(index int) *Node{
    node := linkedList.start
    for i := 0; i < index; i++{
        if node.next != nil {
            node = node.next
        } else {
            panic("Node index out of range!")
        }
    }
    return node
}

// CODE TO TEST THE LIST

func main(){
    // todo: move this away from main after the collection is done

    list := LinkedList{}
    list.add(1)
    list.add(2)
    list.add(3)
    list.add(5)


    fmt.Println(list.get(list.size-1))

    //fmt.Println(list.start)

}
