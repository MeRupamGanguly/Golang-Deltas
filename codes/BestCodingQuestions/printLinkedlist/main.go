package main

import (
	"fmt"
)

type Node struct {
	data  int
	right *Node
}

type list struct {
	Nodes *Node
}

func NewList() *list {
	return &list{}
}
func (l *list) Add(d int) {
	node := Node{
		data: d,
	}
	if l.Nodes == nil {
		l.Nodes = &node
		return
	}
	c := l.Nodes
	for c.right != nil {
		c = c.right
	}
	c.right = &node
}
func (l *list) Remove(n int) {
	if l.Nodes == nil {
		return
	}
	c := l.Nodes
	for i := 0; i < n-1; i++ {
		if c.right != nil {
			c = c.right
		}
	}
	if c.right.right != nil {
		c.right = c.right.right
	}
}
func (l *list) print() {
	c := l.Nodes
	for c.right != nil {
		fmt.Println(c.data)
		c = c.right
	}
	fmt.Println(c.data)
	fmt.Println("------------------------")
}
func (l *list) NthAdd(d int, n int) {
	node := Node{
		data: d,
	}
	if l.Nodes == nil {
		l.Nodes = &node
	}
	if n == 0 {
		node.right = l.Nodes
		l.Nodes = &node
	}

	c := l.Nodes
	for i := 0; i < n-1; i++ {
		if c.right != nil {
			c = c.right
		}
	}
	node.right = c.right
	c.right = &node
}
func main() {
	l := NewList()
	l.Add(2)
	l.Add(4)
	l.Add(6)
	l.Add(9)
	l.Add(12)
	l.Add(16)
	l.print()
	l.NthAdd(17, 8)
	l.print()
	l.NthAdd(14, 3)
	l.print()
	l.Remove(3)
	l.print()
	l.Remove(2)
	l.print()
}
