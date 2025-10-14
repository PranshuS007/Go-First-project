package datastructures

import "fmt"

// Node represents a node in the linked list
type Node struct {
    Data int
    Next *Node
}

// LinkedList represents a singly linked list
type LinkedList struct {
    Head *Node
    Size int
}

// NewLinkedList creates and returns a new empty linked list
func NewLinkedList() *LinkedList {
    return &LinkedList{
        Head: nil,
        Size: 0,
    }
}

// Add appends a new node with the given data to the end of the list
func (ll *LinkedList) Add(data int) {
    newNode := &Node{Data: data, Next: nil}
    
    if ll.Head == nil {
        ll.Head = newNode
    } else {
        current := ll.Head
        for current.Next != nil {
            current = current.Next
        }
        current.Next = newNode
    }
    ll.Size++
}

// AddAtBeginning adds a new node with the given data at the beginning of the list
func (ll *LinkedList) AddAtBeginning(data int) {
    newNode := &Node{Data: data, Next: ll.Head}
    ll.Head = newNode
    ll.Size++
}

// Delete removes the first occurrence of a node with the given data
func (ll *LinkedList) Delete(data int) bool {
    if ll.Head == nil {
        return false
    }
    
    // If head needs to be deleted
    if ll.Head.Data == data {
        ll.Head = ll.Head.Next
        ll.Size--
        return true
    }
    
    current := ll.Head
    for current.Next != nil {
        if current.Next.Data == data {
            current.Next = current.Next.Next
            ll.Size--
            return true
        }
        current = current.Next
    }
    return false
}

// Traverse prints all elements in the linked list
func (ll *LinkedList) Traverse() {
    if ll.Head == nil {
        fmt.Println("List is empty")
        return
    }
    
    current := ll.Head
    fmt.Print("LinkedList: ")
    for current != nil {
        fmt.Print(current.Data)
        if current.Next != nil {
            fmt.Print(" -> ")
        }
        current = current.Next
    }
    fmt.Println()
}

// Search checks if a value exists in the linked list
func (ll *LinkedList) Search(data int) bool {
    current := ll.Head
    for current != nil {
        if current.Data == data {
            return true
        }
        current = current.Next
    }
    return false
}

// GetSize returns the number of nodes in the linked list
func (ll *LinkedList) GetSize() int {
    return ll.Size
}

// IsEmpty checks if the linked list is empty
func (ll *LinkedList) IsEmpty() bool {
    return ll.Head == nil
}
