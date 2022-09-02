package main

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
func (t *tree) getSize() int{
	return t.size
}
func (t *tree) isEmpty() bool{
	return t.size==0
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

func (t *tree) min() *Node {
	return t.minR(t.root)
}
func (t *tree) minR(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.left == nil {
		return node
	}
	return t.minR(node.left)
}
func (t *tree) removeMin() *Node {
	t.root = t.removeMinR(t.root)
	return t.min()
}
func (t *tree) removeMinR(node *Node) *Node {
	if node.left == nil {
		t.size--
		rightNode := node.right
		node.right = nil
		return rightNode
	}
	node.left = t.removeMinR(node.left)
	return node
}
func (t *tree) max() *Node {
	return t.minR(t.root)
}
func (t *tree) maxR(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.right == nil {
		return node
	}
	return t.maxR(node.right)
}
func (t *tree) removeMax() *Node {
	t.root = t.removeMaxR(t.root)
	return t.max()
}
func (t *tree) removeMaxR(node *Node) *Node {
	if node.right == nil {
		t.size--
		leftNode := node.left
		node.left = nil
		return leftNode
	}
	node.right = t.removeMaxR(node.right)
	return node
}
func (t *tree) remove(e int) {
	t.root = t.removeR(t.root, e)
}
func (t *tree) removeR(node *Node, e int) *Node {
	if node == nil {
		return nil
	}
	if node.e > e {
		node.left = t.removeR(node.left, e)
		return node
	} else if node.e < e {
		node.right = t.removeR(node.right, e)
		return node
	} else { //e == node.e
		//没有左子树
		if node.left == nil {
			t.size--
			rightNode := node.right
			node.right = nil
			return rightNode
		}
		//没有右子树
		if node.right == nil {
			t.size--
			leftNode := node.left
			node.left = nil
			return leftNode
		}
		//待删除节点左右子树都不为空
		//找到比待删除节点大的最小节点，即待删除节点右子树的最小节点
		//用这个节点顶替待删除的节点

		//由于在removeMinR中进行过一次size--，实际那个节点没有被删除，所以这里不需要再多进行一次size--了
		successor := t.minR(node.right)
		//这个要写在前面，先把successor从右子树中移除，然后在进行赋值操作
		successor.right = t.removeMinR(node.right)
		successor.left = node.left

		node.left = nil
		node.right = nil

		return successor
	}
}
type BSTSet struct {
	BST *tree
}

func New() *BSTSet {
	return &BSTSet{
		BST: NewTree(),
	}
}

func (bs *BSTSet) Add(e int) {
	bs.BST.add(e)
}

func (bs *BSTSet) Remove(e int) {
	bs.BST.remove(e)
}

func (bs *BSTSet) Contains(e int) bool {
	return bs.BST.contains(e)
}

func (bs *BSTSet) GetSize() int {
	return bs.BST.getSize()
}

func (bs *BSTSet) IsEmpty() bool {
	return bs.BST.isEmpty()
}