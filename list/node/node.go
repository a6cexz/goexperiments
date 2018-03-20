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

func sumAcc(list *Node, acc int) int {
	if list == nil {
		return acc
	}
	return sumAcc(list.next, list.value+acc)
}

func maxAcc(list *Node, max int) int {
	if list == nil {
		return max
	}
	if list.GetValue() > max {
		return maxAcc(list.next, list.GetValue())
	} else {
		return maxAcc(list.next, max)
	}
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

// Sum calculates sum of all elements in the list
func (l *Node) Sum() int {
	return sumAcc(l, 0)
}

// Max find maximum value in the list
func (l *Node) Max() int {
	return maxAcc(l, l.GetValue())
}
