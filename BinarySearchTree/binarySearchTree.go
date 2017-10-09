package main

import (
	"container/list"
	"fmt"
)

func newNode(key, value int) *node {
	var n = node{
		key:   key,
		value: value,
	}
	return &n
}

type node struct {
	key         int
	value       int
	left, right *node
}

func newBST() *BST {
	var bst = BST{
		root:  nil,
		count: 0,
	}

	return &bst
}

type BST struct {
	root  *node
	count int
}

func (t *BST) InsertRecursive(key, value int) {
	t.root = t.insertRecursive(t.root, key, value)
}

func (t *BST) insertRecursive(root *node, key, value int) *node {
	if root == nil {
		t.count++
		return newNode(key, value)
	}

	if root.key == key {
		root.value = value
	}
	if root.key < key {
		root.left = t.insertRecursive(root.left, key, value)
	} else {
		root.right = t.insertRecursive(root.right, key, value)
	}
	return root
}

func (t *BST) Insert(key, value int) {
	if t.root == nil {
		t.root = newNode(key, value)
		return
	}

	current := t.root
	for {
		if current.key == key {
			current.value = value
			return
		}
		if current.key > key {
			if current.left == nil {
				current.left = newNode(key, value)
				t.count++
				return
			}
			current = current.left
		}
		if current.key < key {
			if current.right == nil {
				current.right = newNode(key, value)
				t.count++
				return
			}
			current = current.right
		}
	}
}

func (t *BST) Contain(key int) bool {
	return t.contain(t.root, key)
}

func (t *BST) contain(n *node, key int) bool {
	if n == nil {
		return false
	}

	if n.key == key {
		return true
	}
	if n.key > key {
		return t.contain(n.left, key)
	}
	if n.key < key {
		return t.contain(n.right, key)
	}

	return false
}

func (t *BST) Search(key int) *int {
	return t.search(t.root, key)
}

func (t *BST) search(n *node, key int) *int {
	if n == nil {
		return nil
	}
	if n.key == key {
		return &n.value
	}
	if n.key > key {
		return t.search(n.left, key)
	}
	if n.key < key {
		return t.search(n.right, key)
	}

	return nil
}

func (t *BST) PreOrder() {
	t.preOrder(t.root)
}

func (t *BST) preOrder(n *node) {
	if n != nil {
		fmt.Println(n.key)
		t.preOrder(n.left)
		t.preOrder(n.right)
	}
}

func (t *BST) InOrder() {
	t.inOrder(t.root)
}

func (t *BST) inOrder(n *node) {
	if n != nil {
		t.inOrder(n.left)
		fmt.Println(n.key)
		t.inOrder(n.right)
	}
}

func (t *BST) PostOrder() {
	t.postOrder(t.root)
}

func (t *BST) postOrder(n *node) {
	if n != nil {
		t.postOrder(n.left)
		fmt.Println(n.key)
		t.postOrder(n.right)
	}
}

func (t *BST) LevelOrder() {
	if t.root == nil {
		return
	}

	l := list.New()
	l.PushBack(t.root)
	for {
		e := l.Front()
		if e == nil {
			return
		}
		l.Remove(e)
		v := e.Value.(*node)
		fmt.Println(v.value)

		if v.left != nil {
			l.PushBack(v.left)
		}
		if v.right != nil {
			l.PushBack(v.right)
		}
	}
}

func (t *BST) Minimum() int {
	if t.root == nil {
		panic("no data")
	}

	node := t.root
	for {
		if node.left == nil {
			return node.key
		}
		node = node.left
	}
}

func (t *BST) Maximum() int {
	if t.root == nil {
		panic("no data")
	}

	node := t.root
	for {
		if node.right == nil {
			return node.key
		}
		node = node.right
	}
}

func (t *BST) RemoveMin() {
	t.root = t.removeMin(t.root)
}

func (t *BST) removeMin(n *node) *node {
	if n.left == nil {
		t.count--
		return n.right
	}
	n.left = t.removeMin(n.left)
	return n
}

func (t *BST) RemoveMax() {
	t.root = t.removeMax(t.root)
}

func (t *BST) removeMax(n *node) *node {
	if n.right == nil {
		t.count--
		return n.left
	}
	n.right = t.removeMax(n.right)
	return n
}

func (t *BST) Remove(n *node, key int) *node {
	if n == nil {
		panic("nil node")
	}

	if key < n.key {
		n.left = t.Remove(n.left, key)
		return n
	}
	if key > n.key {
		n.right = t.Remove(n.right, key)
		return n
	}
	if n.left == nil {
		t.count--
		return n.right
	}
	if n.right == nil {
		t.count--
		return n.left
	}
	min := n.right
	for {
		if min.left == nil {
			break
		}
		min = min.left
	}
	min.left = n.left
	min.right = t.removeMin(n.right)
	t.count--
	return min
}
