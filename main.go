package main

import (
	"fmt"
	"sync"
)

var T int // minimum order
type node struct {
	keys     []int
	children []*node
	isLeaf   bool
}

type BTree struct {
	root *node
	mu   sync.Mutex
}

func NewNode() *node {
	return &node{
		keys:     make([]int, 0),
		isLeaf:   true,
		children: make([]*node, 0),
	}
}

func NewBTree(val ...int) *BTree {
	if len(val) == 0 {
		val = append(val, 10)
	}
	T = val[0]
	return &BTree{
		root: NewNode(),
	}
}

func (t *BTree) Insert(val int) error {
	if t == nil {
		err := fmt.Errorf("empty tree")
		return err
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	insert(val, t)
	return nil
}

func (t *BTree) PrintTree() error {
	if t == nil {
		err := fmt.Errorf("empty tree")
		return err
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	printTree(t.root)
	fmt.Println("")
	return nil
}

func (t *BTree) Search(val int) (bool, error) {
	if t == nil {
		err := fmt.Errorf("empty tree")
		return false, err
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	foundNode := search(t.root, val)
	exists := false
	if foundNode != nil {
		fmt.Println("Found", val)
		exists = true
	} else {
		fmt.Println("Does not exist", val)
	}
	return exists, nil
}

func (t *BTree) Delete(key int) error {
	if t.root == nil {
		err := fmt.Errorf("empty tree")
		return err
	}
	t.mu.Lock()
	defer t.mu.Unlock()

	if !t.root.isLeaf && len(t.root.keys) == 0 {
		// If the root has no keys and is not a leaf, set the root to its only child.
		t.root = t.root.children[0]
	}

	deleteNode(&t.root, key)
	return nil
}

func (t *BTree) Update(oldVal int, newVal int) error {
	var err error
	if t == nil {
		err = fmt.Errorf("empty tree")
		return err
	}
	n := search(t.root, oldVal)
	if n == nil {
		err = fmt.Errorf("old key not in tree")
		return err
	}
	// delete oldVal and insert newVal
	err = t.Delete(oldVal)
	if err != nil {
		return err
	}
	err = t.Insert(newVal)
	if err != nil {
		return err
	}
	return nil
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
	t.Search(10)
	t.Update(10, 13)
	t.PrintTree()
	// benchmark(t)
	// concurrency(t)
}
