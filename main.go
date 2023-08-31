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
		isLeaf:   true,
		children: make([]*node, 0),
	}
}

func NewBTree() *bTree {
	return &bTree{
		root: NewNode(),
	}
}

func splitNode(x *node, index int) {
	rightChild := NewNode()
	leftChild := x.children[index]
	n := len(x.keys)
	rightChild.isLeaf = leftChild.isLeaf
	for i := 0; i < T-1; i++ {
		rightChild.keys = append(rightChild.keys, leftChild.keys[i+T])
	}
	leftChild.keys = leftChild.keys[:T]
	if !leftChild.isLeaf {
		for i := 0; i < T; i++ {
			rightChild.children = append(rightChild.children, leftChild.children[i+T])
		}
		leftChild.children = leftChild.children[:T]
	}
	x.children = append(x.children, nil)
	for i := n; i > index+1; i-- {
		x.children[i] = x.children[i-1]
	}
	x.children[index+1] = rightChild
	x.keys = append(x.keys, 0)
	for i := n - 1; i > index; i-- {
		x.keys[i] = x.keys[i-1]
	}
	x.keys[index] = leftChild.keys[T-1]
	leftChild.keys = leftChild.keys[:T-1]
}

func insertNonFull(x *node, val int) {
	i := len(x.keys) - 1 // last filled index
	if x.isLeaf {
		x.keys = append(x.keys, val) // just for maintaining indices
		for i >= 0 && val < x.keys[i] {
			x.keys[i+1] = x.keys[i]
			i--
		}
		x.keys[i+1] = val
	} else {
		for i >= 0 && val < x.keys[i] {
			i--
		}
		i++
		if len(x.children) == (2*T - 1) {
			splitNode(x, i)
			if val > x.keys[i] {
				i++
			}
		}
		insertNonFull(x.children[i], val)
	}
}

func Insert(val int, t *bTree) error {
	if t == nil {
		err := fmt.Errorf("empty tree")
		return err
	}
	n := len(t.root.keys)
	if n == 2*T-1 {
		newRoot := NewNode()
		newRoot.isLeaf = false
		newRoot.children = append(newRoot.children, t.root)
		t.root = newRoot
		splitNode(t.root, 0)
		insertNonFull(t.root, val)
	} else {
		insertNonFull(t.root, val)
	}
	return nil
}

func (t *bTree) PrintTree(n *node) {
	if n.isLeaf {
		for j := 0; j < len(n.keys); j++ {
			fmt.Printf("%v ", n.keys[j])
		}
	} else {
		for i := 0; i < len(n.keys); i++ {
			t.PrintTree(n.children[i])
			fmt.Printf("%v ", n.keys[i])
		}
		t.PrintTree(n.children[len(n.keys)])
	}
}

func main() {
	fmt.Println("B-Tree Example")
	t := NewBTree()
	Insert(21, t)
	Insert(22, t)
	Insert(20, t)
	Insert(23, t)
	Insert(18, t)
	Insert(25, t)
	t.PrintTree(t.root)
}
