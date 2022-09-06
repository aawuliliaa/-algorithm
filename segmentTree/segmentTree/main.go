package main

type SegmentTree struct {
	tree   []int
	data   []int
	merger func(i, j int) int
}

func New(arr []int) *SegmentTree {
	segmentTree := &SegmentTree{
		tree: make([]int, len(arr)*4),
		data: make([]int, len(arr)),
	}
	segmentTree.data = arr
	return segmentTree
}

func (st *SegmentTree) GetSize() int {
	return len(st.data)
}

func (st *SegmentTree) Get(index int) int {
	if index < 0 || index >= len(st.data) {
		panic("index is out of range")
	}

	return st.data[index]
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的左孩子节点的索引
func (st *SegmentTree) leftChild(index int) int {
	return index*2 + 1
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的右孩子节点的索引
func (st *SegmentTree) rightChild(index int) int {
	return index*2 + 2
}
func (st *SegmentTree) buildSegmentTree() {
	st.buildSegmentTreeR(0, 0, len(st.data)-1)

}
func (st *SegmentTree) buildSegmentTreeR(treeIndex, l, r int) {
	if l == r {
		st.tree[treeIndex] = st.data[l]
		return
	}
	mid := l + (r-l)/2
	leftIndex := st.leftChild(treeIndex)
	rightIndex := st.rightChild(treeIndex)
	st.buildSegmentTreeR(leftIndex, l, mid)
	st.buildSegmentTreeR(rightIndex, mid+1, r)
	st.tree[treeIndex] = st.merger(st.tree[leftIndex], st.tree[rightIndex])
}

// 在以treeIndex为根的线段树中[l...r]的范围里，搜索区间[queryL...queryR]的值
func (st *SegmentTree) query(queryL, queryR int) int {
	if queryL < 0 || queryL >= len(st.data) || queryR < 0 || queryR > len(st.data) {
		panic("index is out of range")
	}
	return st.queryR(0, 0, len(st.data)-1, queryL, queryR)
}
func (st *SegmentTree) queryR(treeIndex, l, r, queryL, queryR int) int {
	if l == queryL && r == queryR {
		return st.tree[treeIndex]
	}
	mid := l + (r-l)/2
	leftIndex := st.leftChild(treeIndex)
	rightIndex := st.rightChild(treeIndex)
	if queryL >= mid+1 {
		return st.queryR(rightIndex, mid+1, r, queryL, queryR)
	} else if queryR <= mid { //<=mid
		return st.queryR(leftIndex, l, mid, queryL, queryR)
	}
	leftResult := st.queryR(leftIndex, l, mid, queryL, mid)
	rightResult := st.queryR(rightIndex, mid+1, r, mid+1, queryR)
	return st.merger(leftResult, rightResult)

}
func (st *SegmentTree) set(index, e int) {
	st.data[index] = e
	st.setR(0, 0, len(st.data)-1, index, e)

}
func (st *SegmentTree) setR(treeIndex, l, r, index, e int) {
	if l == r {
		st.tree[treeIndex] = e
	}
	mid := l + (r-l)/2
	leftIndex := st.leftChild(treeIndex)
	rightIndex := st.rightChild(treeIndex)
	if index > mid {
		st.setR(rightIndex, mid+1, r, index, e)
	} else {
		st.setR(leftIndex, l, mid, index, e)
	}
	st.tree[treeIndex] = st.merger(st.tree[leftIndex], st.tree[rightIndex])
}
