package main

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	var t1, t2 = newBST(), newBST()

	t1.Insert(1, 20)
	t1.Insert(2, 30)
	t1.Insert(3, 10)
	fmt.Println("PreOrder")
	t1.PreOrder()
	fmt.Println("InOrder")
	t1.InOrder()
	fmt.Println("PostOrder")
	t1.PostOrder()
	fmt.Println("LevelOrder")
	t1.LevelOrder()

	t2.InsertRecursive(1, 20)
	t2.InsertRecursive(2, 10)
	t2.InsertRecursive(3, 30)
	fmt.Println("PreOrder")
	t2.PreOrder()
	fmt.Println("InOrder")
	t2.InOrder()
	fmt.Println("PostOrder")
	t2.PostOrder()
	fmt.Println("LevelOrder")
	t2.LevelOrder()

	fmt.Println(t1.Contain(2))
	fmt.Println(t1.Contain(4))
	fmt.Println(*t1.Search(2))
	fmt.Println(t1.Search(4))
}
