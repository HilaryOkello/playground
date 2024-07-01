package main

import "fmt"

type Node struct {
	Data string
	Next *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func ListPushBack(l *List, data string) {
	newNode := &Node{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}

func main() {
	link := &List{}
	ListPushBack(link, "Hello")
	ListPushBack(link, "man")
	ListPushBack(link, "how are you")
	current := link.Head
	for current != nil {
		fmt.Print(current.Data, "->")
		current = current.Next
	}
	fmt.Println()
}
