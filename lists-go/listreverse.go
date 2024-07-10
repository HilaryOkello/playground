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

func ListReverse(l *List) {
	current := l.Head
	var prev *Node
	var next *Node
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current =next
	}
	l.Head = l.Tail
	l.Tail = prev

}

func main() {
	link := &List{}
	ListPushBack(link, "1")
	ListPushBack(link, "2")
	ListPushBack(link, "3")
	ListPushBack(link, "4")

	ListReverse(link)

	it := link.Head

	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}

	fmt.Println("Tail", link.Tail)
	fmt.Println("Head", link.Head)
}
