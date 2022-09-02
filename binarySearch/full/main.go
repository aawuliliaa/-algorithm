package main

import "fmt"

type queue struct {
	data []interface{}
	size int
}

func newQueue(cp int) *queue {
	return &queue{
		size: 0,
		data: make([]interface{}, 0, cp),
	}

}
func (q *queue) Push(e interface{}) {
	q.data = append(q.data, e)
	q.size++
}
func (q *queue) Pop() interface{} {
	e := q.data[0]
	q.data = q.data[1:]
	q.size--
	return e
}
func (q *queue) Size() int {
	return q.size
}
func (q *queue) IsEmpty() bool {
	return q.size == 0
}

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
	return t.containsR(t.root, e)

}
func (t *tree) containsR(node *Node, e int) bool {
	if node == nil {
		return false
	}
	if node.e == e {
		return true
	} else if node.e > e {
		return t.containsR(node.left, e)
	} else {
		return t.containsR(node.right, e)
	}
}
func (t *tree) preOrder() {
	t.preOrderR(t.root)

}
func (t *tree) preOrderR(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.e)
	t.preOrderR(node.left)
	t.preOrderR(node.right)
}
func (t *tree) inOrder() {
	t.inOrderR(t.root)
}
func (t *tree) inOrderR(node *Node) {
	if node == nil {
		return
	}

	t.preOrderR(node.left)
	fmt.Println(node.e)
	t.preOrderR(node.right)
}
func (t *tree) postOrder() {
	t.postOrderR(t.root)
}
func (t *tree) postOrderR(node *Node) {
	if node == nil {
		return
	}

	t.preOrderR(node.left)
	t.preOrderR(node.right)
	fmt.Println(node.e)
}
func (t *tree) levelOrder() {
	q := newQueue(3)
	q.Push(t.root)
	for {
		if q.IsEmpty() {
			return
		}
		cur := q.Pop().(*Node)
		fmt.Println(cur.e)
		if cur.left != nil {
			q.Push(cur.left)
		}
		if cur.right != nil {
			q.Push(cur.right)
		}

	}
}
func main() {
	tree := NewTree()
	for _, i := range []int{3, 5, 1, 2, 4} {
		tree.add(i)
	}
	tree.levelOrder()

}
