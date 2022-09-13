package main

//小顶堆
type heap struct {
	data []int
	size int
}

func NewHeap(capacity int) *heap {
	data := make([]int, 0, capacity)
	return &heap{data: data}
}

func (h *heap) leftIndex(index int) int {
	return index*2 + 1
}
func (h *heap) rightIndex(index int) int {
	return index*2 + 2
}
func (h *heap) parentIndex(index int) int {
	return (index - 1) / 2
}
func (h *heap) siftUp(index int) {
	for index > 0 {
		parentIndex := h.parentIndex(index)
		if h.data[index] < h.data[parentIndex] {
			h.data[index], h.data[parentIndex] = h.data[parentIndex], h.data[index]
			index = parentIndex
		} else {
			break
		}
	}
}
func (h *heap) siftDown(index int) {
	for h.leftIndex(index) <= h.size-1 {
		indexToCompare := h.leftIndex(index)
		rightIndex := h.rightIndex(index)
		if rightIndex <= h.size-1 && h.data[indexToCompare] > h.data[rightIndex] {
			indexToCompare = rightIndex
		}
		if h.data[index] > h.data[indexToCompare] {
			h.data[index], h.data[indexToCompare] = h.data[indexToCompare], h.data[index]
			index = indexToCompare

		} else {
			break
		}

	}
}
func (h *heap) add(val int) {
	h.data = append(h.data, val)
	h.size++
	h.siftUp(h.size - 1)
}

func findKthLargest(nums []int, k int) int {
	heap := NewHeap(k)
	for index, num := range nums {
		if index < k {
			heap.add(num)
		} else {
			if num > heap.data[0] {
				heap.data[0] = num
				heap.siftDown(0)
			}

		}
	}
	return heap.data[0]

}
