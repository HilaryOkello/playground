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

func ListLast(l *List) string {
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	return current.Data
}

func main() {
	link := &List{}
	ListPushBack(link, "Hello")
	ListPushBack(link, "man")
	ListPushBack(link, "how are you")
	fmt.Println(ListLast(link))
}
