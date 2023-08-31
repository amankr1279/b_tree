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

func insert(val int, t *bTree) {
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
}

func (t *bTree) Insert(val int) error {
	if t == nil {
		err := fmt.Errorf("empty tree")
		return err
	}
	insert(val, t)
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

func search(n *node, val int) bool {
	if n == nil {
		return false
	}
	numKeys := len(n.keys)
	numChildren := len(n.children)
	if numKeys == 0 {
		return false
	}
	if n.keys[0] < val && val <= n.keys[numKeys-1] {
		for i := 0; i < numKeys-1; i++ {
			if n.keys[i] == val {
				return true
			}
			if n.keys[i] < val && n.keys[i+1] > val && numChildren >= (i+1) {
				return search(n.children[i+1], val)
			}
		}
		if n.keys[numKeys-1] == val {
			return true
		}
	} else if n.keys[numKeys-1] < val {
		if numChildren == 1+numKeys {
			return search(n.children[numKeys], val)
		}
	} else {
		if !n.isLeaf {
			return search(n.children[0], val)
		}
	}
	return false
}

func (t *bTree) Search(val int) (bool, error) {
	if t == nil {
		err := fmt.Errorf("empty tree")
		return false, err
	}
	exists := search(t.root, val)
	if exists {
		fmt.Println("Found", val)
	} else {
		fmt.Println("Does not exist", val)
	}
	return exists, nil
}

func main() {
	fmt.Println("B-Tree Example")
	t := NewBTree()
	t.Insert(5)
	t.Insert(8)
	t.Insert(10)
	t.Insert(12)
	t.Insert(14)
	t.Insert(4)
	t.Insert(6)
	t.Search(14)
	t.PrintTree(t.root)
}
