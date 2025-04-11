package main

import (
	"fmt"

	"github.com/collinglass/bptree"
)

func main() {
	//Example1()
	Example2()
}

func Example1() {
	key := 1
	value := []byte("hello friend")

	t := bptree.NewTree()

	err := t.Insert(key, value)
	if err != nil {
		fmt.Printf("error: %s\n\n", err)
	}

	r, err := t.Find(key, true)
	if err != nil {
		fmt.Printf("error: %s\n\n", err)
	}

	fmt.Printf("%s\n\n", r.Value)

	t.FindAndPrint(key, true)
	t.PrintLeaves()
}

// http://www.cburch.com/cs/340/reading/btree/index.html, 文章里是B+-tree而不是单纯的B+tree
func Example2() {
	t := bptree.NewTree()
	t.Insert(13, nil)
	t.Insert(9, nil)
	t.Insert(11, nil)
	t.Insert(16, nil)
	t.Insert(1, nil)
	t.Insert(4, nil)
	t.Insert(9, nil)
	t.Insert(10, nil)
	t.Insert(11, nil)
	t.Insert(12, nil)
	t.Insert(13, nil)
	t.Insert(15, nil)
	t.Insert(16, nil)
	t.Insert(20, nil)
	t.Insert(25, nil)

	t.PrintLeaves()
	t.PrintLeaves()
	t.PrintTree()
	t.PrintTree()
}
