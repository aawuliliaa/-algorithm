package fastSort

func fastSortR(nums []int, index, l, r int) int {
	partition := partition(nums, l, r)
	if partition == index {
		return nums[index]
	}
	if partition < index {
		return fastSortR(nums, index, partition+1, r)

	} else {
		return fastSortR(nums, index, l, partition-1)
	}
}
func partition(nums []int, l, r int) int {
	pivot := nums[l]
	i := l + 1
	j := r
	for {
		switch {
		case i > j:
			nums[l], nums[j] = nums[j], nums[l]
			return j
		case nums[i] < pivot:
			i++
		case nums[j] > pivot:
			j--
		default:
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
}
func findKthLargest(nums []int, k int) int {
	return fastSortR(nums, len(nums)-k, 0, len(nums)-1)
}
