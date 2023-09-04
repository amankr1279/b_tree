package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wgI sync.WaitGroup
var wgD sync.WaitGroup
var wgS sync.WaitGroup

func benchmark(t *BTree) {
	st := time.Now()
	numOps := 1_000_000
	for i := 0; i < numOps; i++ {
		t.Insert(i + 1)
		if i%1000 == 0 {
			fmt.Println("Epoch : #", i/1000)
		}
	}
	x := time.Since(st)
	fmt.Printf("Elapsed time: %v\n", x)
	fmt.Printf("Number of operations per milliisecond : %v\n", int64(numOps)/x.Milliseconds())
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

func concurrency(t *BTree) {
	nums := make([]int, 0)
	for i := 0; i < 10; i++ {
		wgI.Add(1)
		go func() {
			x := rand.Intn(100)
			nums = append(nums, x)
			t.Insert(x)
			wgI.Done()
		}()
	}
	wgI.Wait()
	for i := 0; i < 10; i++ {
		wgS.Add(1)
		go func(idx int) {
			t.Search(nums[idx])
			wgS.Done()
		}(i)
	}
	wgS.Wait()
	for i := 0; i < 10; i++ {
		wgI.Add(1)
		wgS.Add(1)
		wgD.Add(1)
		go func(idx int) {
			t.Update(nums[idx], nums[idx]+100)
			wgI.Done()
			wgS.Done()
			wgD.Done()
		}(i)
	}
	wgI.Wait()
	wgS.Wait()
	wgD.Wait()
	t.PrintTree()
}
