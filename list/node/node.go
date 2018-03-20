package node

import (
	"bytes"
	"fmt"
	"io"
)

// Node represents simple linked list
type Node struct {
	value int
	next  *Node
}

// New creates new list from the given elements
func New(data ...int) *Node {
	var h *Node
	var r *Node
	for _, v := range data {
		n := &Node{value: v}
		if r == nil {
			r = n
			h = r
		} else {
			r.next = n
			r = n
		}
	}
	return h
}

// GetValue returns node value
func (l *Node) GetValue() int {
	return l.value
}

// GetTail returns list tail
func (l *Node) GetTail() *Node {
	return l.next
}

func reverseListAcc(list *Node, acc *Node) *Node {
	if list == nil {
		return acc
	}
	head := &Node{value: list.value, next: acc}
	return reverseListAcc(list.next, head)
}

func write(writer io.Writer, list *Node) {
	fmt.Fprint(writer, list.value)
	if list.next != nil {
		fmt.Fprint(writer, ", ")
		write(writer, list.next)
	}
}

func (l *Node) Write(writer io.Writer) {
	fmt.Fprint(writer, "[")
	write(writer, l)
	fmt.Fprint(writer, "]")
}

// Print prints list
func (l *Node) Print() {
	fmt.Println(l.String())
}

func (l *Node) String() string {
	buffer := &bytes.Buffer{}
	l.Write(buffer)
	return buffer.String()
}

// Reverse creates reversed list copy
func (l *Node) Reverse() *Node {
	return reverseListAcc(l, nil)
}
