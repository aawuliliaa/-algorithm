package main

//307--区域检索，数据可修改
type NumArray struct {
	segmentTree []int
	data        []int
	merger      func(i, j int) int
}

func Constructor(nums []int) NumArray {
	numArray := NumArray{}
	numArray.data = nums
	numArray.merger = func(i, j int) int {
		return i + j
	}
	numArray.buildSegmentTree()
	return numArray
}
func (this *NumArray) leftChild(index int) int {
	return index*2 + 1
}
func (this *NumArray) rightChild(index int) int {
	return index*2 + 2
}
func (this *NumArray) buildSegmentTree() {
	this.segmentTree = make([]int, len(this.data)*4)
	this.buildSegmentTreeR(0, 0, len(this.data)-1)
}

func (this *NumArray) buildSegmentTreeR(treeIndex, l, r int) {
	if l == r {
		this.segmentTree[treeIndex] = this.data[l]
		return
	}
	mid := l + (r-l)/2
	leftChildIndex := this.leftChild(treeIndex)
	rightChildIndex := this.rightChild(treeIndex)
	this.buildSegmentTreeR(leftChildIndex, l, mid)
	this.buildSegmentTreeR(rightChildIndex, mid+1, r)
	this.segmentTree[treeIndex] = this.merger(this.segmentTree[leftChildIndex], this.segmentTree[rightChildIndex])

}
func (this *NumArray) Update(index int, val int) {
	this.data[index] = val
	this.updateR(0,0,len(this.data)-1,index,val)

}
func (this *NumArray) updateR(treeIndex, l, r, index, val int) {
	if l == r {
		this.segmentTree[treeIndex] = val
		return
	}
	mid := l + (r-l)/2
	leftChildIndex := this.leftChild(treeIndex)
	rightChildIndex := this.rightChild(treeIndex)
	if index > mid {
		this.updateR(rightChildIndex, mid+1, r, index, val)
	} else {
		this.updateR(leftChildIndex, l, mid, index, val)
	}
	this.segmentTree[treeIndex] = this.merger(this.segmentTree[leftChildIndex], this.segmentTree[rightChildIndex])
}

func (this *NumArray) SumRange(left int, right int) int {

	return this.queryR(0, 0, len(this.data)-1, left, right)
}
func (this *NumArray) queryR(treeIndex, l, r, queryLeft, queryRight int) int {
	if l == queryLeft && r == queryRight {
		return this.segmentTree[treeIndex]
	}
	mid := l + (r-l)/2
	leftChildIndex := this.leftChild(treeIndex)
	rightChildIndex := this.rightChild(treeIndex)
	if queryLeft > mid {
		return this.queryR(rightChildIndex, mid+1, r, queryLeft, queryRight)
	} else if queryRight <= mid {
		return this.queryR(leftChildIndex, l, mid, queryLeft, queryRight)
	}
	leftResult := this.queryR(leftChildIndex, l, mid, queryLeft, mid)
	rightResult := this.queryR(rightChildIndex, mid+1, r, mid+1, queryRight)
	return this.merger(leftResult, rightResult)
}