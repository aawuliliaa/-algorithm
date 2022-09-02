package main

import "fmt"

type Node struct {
	e     int
	left  *Node
	right *Node
}
type tree struct {
	size int
	root *Node
}

func NewTree() *tree {
	return &tree{}
}

func (t *tree) add(e int) {
	t.root = t.addR(t.root, e)

}
func (t *tree) addR(node *Node, e int) *Node {
	if node == nil {
		t.size++
		return &Node{e: e}
	} else if node.e > e {
		node.left = t.addR(node.left, e)
	} else {
		node.right = t.addR(node.right, e)
	}
	return node
}
func (t *tree) contains(e int) bool {
	return t.containsR(t.root,e)

}
func (t *tree) containsR(node *Node, e int) bool {
	if node==nil{
		return false
	}
	if node.e==e{
		return true
	}else if node.e>e{
		return t.containsR(node.left,e)
	}else{
		return t.containsR(node.right,e)
	}
}

func main() {
	tree := NewTree()
	for _, i := range []int{3, 5, 1, 2, 4} {
		tree.add(i)
	}
	fmt.Println(tree.contains(6))

}
