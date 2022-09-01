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
	if tree.root == nil {
		tree.root = &Node{
			e:     e,
			left:  nil,
			right: nil,
		}
	} else {
		tree.add(tree.root, e)
	}

}
func (tree *binaryTree) add(node *Node, e int) {
	if node.e == e {
		return
	} else if node.e > e && node.left == nil {
		tree.size++
		node.left = &Node{
			e:     e,
			left:  nil,
			right: nil,
		}
		return
	} else if node.e < e && node.right == nil {
		tree.size++
		node.right = &Node{
			e:     e,
			left:  nil,
			right: nil,
		}
		return
	}
	if node.e > e {
		tree.add(node.left, e)
	} else {
		tree.add(node.right, e)
	}

}
func main() {
	tree :=NewBinaryTree()
	for _,i := range []int{3, 5, 1, 2, 4} {
		tree.addNode(i)
	}
	fmt.Println(tree)

}
