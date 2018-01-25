package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) Insert(key interface{}) {
	list := &Node{
		next: l.head,
		key:  key,
	}

	if l.head != nil {
		l.head.prev = list
	}

	l.head = list

	lh := l.head

	for lh.next != nil {
		lh = lh.next
	}

	l.tail = lh
}

func (l *List) Display() {
	list := l.head

	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}

	fmt.Println()
}

func Display(list *Node) {
	for list != nil {
		fmt.Printf("%v ->", list.key)
		list = list.next
	}

	fmt.Println()
}

func main() {
	link := List{}
	var path string
	for {
		path = os.Args[1]
		if path != "" {
			break
		} else {
			fmt.Println("Haven't Found a Dictionary!")
		}
	}

	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	desiredWord := "cat"

	for scanner.Scan() {
		link.Insert(scanner.Text())

		if link.head.key == desiredWord {
			fmt.Println("We Found It!", link.head.key)
		} else {
			break
		}
	}

	fmt.Printf("Head: %v\n", link.head.key)
	fmt.Printf("Tail: %v\n", link.tail.key)
	link.Display()
}
