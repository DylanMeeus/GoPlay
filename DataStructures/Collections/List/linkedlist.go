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
    head *Node
    end *Node       // last node in the list
}

/*
    Create the first node in the list, or iterate until the last node and insert it there.
 */
func (linkedList *LinkedList) add(value interface{}){
    if linkedList.head == nil {
        linkedList.head = &Node{value: value, next: nil}
        linkedList.end = linkedList.head
        linkedList.size += 1
    } else {
        newNode := Node{value: value, next : nil}
        lastNode := linkedList.end
        lastNode.next = &newNode
        linkedList.end = &newNode
        linkedList.size += 1
    }
}

/*
    Return the n'th node in the LinkedList
 */
func (linkedList *LinkedList) get(index int) interface{}{
    node := linkedList.head
    for i := 0; i < index; i++{
        if node.next != nil {
            node = node.next
        } else {
            panic("Node index out of range!")
        }
    }
    return node.value
}

/*
    Remove a node given an index
    Panic's if the index is not inside the bounds of the list!
 */
func (linkedList *LinkedList) remove(index uint){
    if index == 0 {
        linkedList.head = linkedList.head.next
        return
    }
    node := linkedList.head
    for i := uint(0); i < index - 1; i++{
        if node.next != nil {
            node = node.next
        } else {
            panic("Node index out of range!")
        }
    }
    // now we have the node at the index, we want to chain the n-1'th element, to the n+1'th element.
    node.next = node.next.next
}

// CODE TO TEST THE LIST

func main(){
    // todo: move this away from main after the collection is done

    list := LinkedList{}


    // perf testing

    for i := 0; i < 10000000; i++{
        list.add(i)
    }

    //list.printList()

}



// Convenience function to test the list
func (LinkedList *LinkedList) printList(){
    node := LinkedList.head
    for node != nil{
        fmt.Println(node)
        node = node.next
    }
}