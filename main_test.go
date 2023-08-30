package main

import (
	"fmt"
	"strings"
	"testing"
)

type args struct {
	val int
	t   *bTree
}

type testObj struct {
	name    string
	args    args
	wantErr bool
}

func TestInsert(t *testing.T) {
	bTree := NewBTree()
	tests := make([]testObj, 0)
	tests = append(tests, unitTestMaker("success_leaf_add_1", args{21, bTree}))
	tests = append(tests, unitTestMaker("success_leaf_add_2", args{19, bTree}))
	tests = append(tests, unitTestMaker("success_leaf_add_3", args{22, bTree}))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Insert(tt.args.val, tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	fmt.Printf("Tree root: %+v", bTree.root.keys)
}

func unitTestMaker(testName string, testAgrs args) testObj {
	var test testObj
	isSuccess := true
	test.name = testName
	test.wantErr = isSuccess
	if strings.Contains(testName, "success") {
		test.wantErr = !isSuccess
	}
	test.args = testAgrs

	return test
}
