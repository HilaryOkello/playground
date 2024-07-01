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

func ListAt(l *Node, pos int) *Node {
	count := 0
	for l != nil {
		if count == pos {
			return l
		}
		count++
		l = l.Next
	}
	return nil
}

func main() {
	link := &List{}
	ListPushBack(link, "hello")
	ListPushBack(link, "how are")
	ListPushBack(link, "you")
	ListPushBack(link, "1")

	fmt.Println(ListAt(link.Head, 3).Data)
	fmt.Println(ListAt(link.Head, 1).Data)
	fmt.Println(ListAt(link.Head, 7))
}
