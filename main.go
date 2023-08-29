package main

import "fmt"

const T int = 3 // minimum degree --> Could be made part of the tree
type node struct {
	keys     []int
	children []*node
	isLeaf   bool
}

type bTree struct {
	root *node
}

func NewNode() *node {
	return &node{
		keys:     make([]int, 0),
		isLeaf:   false,
		children: make([]*node, 0),
	}
}

func NewBTree() *bTree {
	return &bTree{
		root: NewNode(),
	}
}

func splitNode(x *node, index int) {
	z := NewNode()
	y := x.children[index]
	n := len(x.keys)
	z.isLeaf = y.isLeaf
	for i := 0; i < T-1; i++ {
		z.keys = append(z.keys, y.keys[i+T])
	}
	if !y.isLeaf {
		for i := 0; i < T; i++ {
			z.children = append(z.children, y.children[i+T])
		}
	}
	x.children = append(x.children, nil)
	for i := n + 1; i >= index+1; i-- {
		x.children[i] = x.children[i-1]
	}
	x.children[index+1] = z
	x.keys = append(x.keys, 0)
	for i := n; i > index; i-- {
		x.keys[i] = x.keys[i-1]
	}
	x.keys[index] = y.keys[T]
}

func insertNonFull(s *node, val int) {
	
}

func Insert(val int, t *bTree) {
	n := len(t.root.keys)
	if n == 2*T-1 {
		s := NewNode()
		s.children = append(s.children, t.root)
		splitNode(s, 0)
		insertNonFull(s, val)
	} else {
		insertNonFull(t.root, val)
	}
}

func main() {
	fmt.Println("B-Tree Example")
	t := NewBTree()
	fmt.Printf("Root : %+v\n", t.root.keys)
	Insert(21, t)
}
