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

func (tree *binaryTree) add(e int) {
	if tree.root == nil {
		tree.size++
		tree.root = &Node{
			e:     e,
			left:  nil,
			right: nil,
		}
		return
	}
	node := tree.root
	for {
		if node.e > e {
			if node.left == nil {
				tree.size++
				node.left = &Node{
					e:     e,
					left:  nil,
					right: nil,
				}
				return
			} else {
				node = node.left
			}
		} else {
			if node.right == nil {
				tree.size++
				node.right = &Node{
					e:     e,
					left:  nil,
					right: nil,
				}
				return
			} else {
				node = node.right
			}
		}
	}

}
func main() {
	tree := NewBinaryTree()
	for _, i := range []int{3, 5, 1, 2, 4} {
		tree.add(i)
	}
	fmt.Println(tree)

}
