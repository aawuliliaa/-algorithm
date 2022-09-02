package main

import (
	"fmt"
)

type Node struct {
	e     int
	left  *Node
	right *Node
}
type binaryTree struct {
	size int
	root *Node
}

func NewBinaryTree() *binaryTree {
	return &binaryTree{}
}

func (tree *binaryTree) addNode(e int) {
	tree.root = tree.add(tree.root, e)

}
func (tree *binaryTree) add(node *Node, e int) *Node {
	if node == nil {
		tree.size++
		return &Node{
			e:     e,
			left:  nil,
			right: nil,
		}
	} else if node.e > e {
		node.left = tree.add(node.left, e)
	} else {
		node.right = tree.add(node.right, e)
	}
	return node

}
func main() {
	tree := NewBinaryTree()
	for _, i := range []int{3, 5, 1, 2, 4} {
		tree.addNode(i)
	}
	fmt.Println(tree)

}
