package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

type args struct {
	val int
}

type testObj struct {
	name    string
	args    args
	want    any
	wantErr bool
}

func TestInsert(t *testing.T) {
	bTree := NewBTree()
	tests := make([]testObj, 0)
	for i := 0; i < 2*T-1; i++ {
		name := fmt.Sprintf("success_leaf_add_%v", i)
		tests = append(tests, unitTestMaker(name, args{rand.Intn(10000)}, nil))
	}
	for i := 0; i < 2*T-1; i++ {
		name := fmt.Sprintf("success_node_add_%v", i)
		tests = append(tests, unitTestMaker(name, args{rand.Intn(10000)}, nil))
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := bTree.Insert(tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("bTree.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	bTree.PrintTree(bTree.root)
}

func unitTestMaker(testName string, testAgrs args, want any) testObj {
	var test testObj
	isSuccess := true
	test.name = testName
	test.wantErr = isSuccess
	if strings.Contains(testName, "success") {
		test.wantErr = !isSuccess
	}
	test.args = testAgrs
	test.want = want

	return test
}

func Test_bTree_Search(t *testing.T) {
	bTree := NewBTree()
	tests := make([]testObj, 0)
	nums := make([]int, 0)
	for i := 0; i < 3*T; i++ {
		x := rand.Intn(100)
		nums = append(nums, x)
		bTree.Insert(x)
	}
	bTree.PrintTree(bTree.root)
	for i := 0; i < len(nums); i++ {
		name := fmt.Sprintf("success_find_%v", i+1)
		tests = append(tests, unitTestMaker(name, args{val: nums[i]}, true))
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := bTree
			got, err := tr.Search(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("bTree.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("bTree.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
