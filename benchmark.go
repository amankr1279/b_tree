package main

import (
	"fmt"
	"time"
)

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
