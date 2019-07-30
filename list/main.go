package main

import "github.com/a6cexz/goexperiments/list/node"

func main() {
	list := node.New(1, 2, 3)
	list.Print()

	rlist := list.Reverse()
	rlist.Print()

	rlist2 := list.Reverse()
	rlist2.Print()
}
