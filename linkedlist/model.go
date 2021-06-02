package linkedlist

import "fmt"

type node struct {
	val  int
	next *node
}

func (n *node) printN() {
	header := n
	for header != nil {
		fmt.Println(header.val)
		header = header.next
	}
}
