package main

import "fmt"

const T int = 1000 // minimum degree --> Could be made part of the tree
type node struct {
	keys     []int
	children []*node
	isLeaf   bool
}

type BTree struct {
	root *node
}

func NewNode() *node {
	return &node{
		keys:     make([]int, 0),
		isLeaf:   true,
		children: make([]*node, 0),
	}
}

func NewBTree() *BTree {
	return &BTree{
		root: NewNode(),
	}
}

func (t *BTree) Insert(val int) error {
	if t == nil {
		err := fmt.Errorf("empty tree")
		return err
	}
	insert(val, t)
	return nil
}

func printTree(n *node) {
	if n.isLeaf {
		for j := 0; j < len(n.keys); j++ {
			fmt.Printf("%v ", n.keys[j])
		}
	} else {
		for i := 0; i < len(n.keys); i++ {
			printTree(n.children[i])
			fmt.Printf("%v ", n.keys[i])
		}
		printTree(n.children[len(n.keys)])
	}
}

func (t *BTree) PrintTree() error {
	if t == nil {
		err := fmt.Errorf("empty tree")
		return err
	}
	printTree(t.root)
	fmt.Println("")
	return nil
}

func (t *BTree) Search(val int) (bool, error) {
	if t == nil {
		err := fmt.Errorf("empty tree")
		return false, err
	}
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
	benchmark(t)
}
