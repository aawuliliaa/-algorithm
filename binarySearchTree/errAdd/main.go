package main

import "fmt"

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
	tree.add(tree.root, e)

}
func (tree *binaryTree) add(node *Node, e int) {
	if node == nil {
		node = &Node{
			e:     0,
			left:  nil,
			right: nil,
		}
	} else if node.e > e {
		tree.add(node.left, e)
	} else {
		tree.add(node.right, e)
	}

}
func main() {
	tree := NewBinaryTree()
	for _, i := range []int{3, 5, 1, 2, 4} {
		tree.addNode(i)
	}
	fmt.Println(tree)

}
