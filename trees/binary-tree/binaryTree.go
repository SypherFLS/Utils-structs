package binarytree

import (

)

type tree struct {
	root *node
}

type node struct {
	value int
	left *node
	right *node
}

func NewBinTree()  *tree{
	return &tree{root: nil}
}

func (t *tree) Insert(num int) {
	t.root = t.InsertRecursive(t.root, num)
}

func (t *tree) InsertRecursive(n *node, num int) *node {
	if n == nil {
		return &node{
			value: num,
			left: nil,
			right: nil,
		}
	}

	if num < n.value {
		n.left = t.InsertRecursive(n.left, num)
	}

	if num > n.value {
		n.right = t.InsertRecursive(n.right, num)
	}

	return n
}