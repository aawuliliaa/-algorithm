package main
//303--预处理数组
type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = nums[i] + sums[i-1]
	}
	return NumArray{
		sum: sums,
	}

}

func (this *NumArray) SumRange(left int, right int) int {
	if left==0{
		return this.sum[right]
	}
	return this.sum[right]-this.sum[left-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
