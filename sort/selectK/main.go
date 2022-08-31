package main

import "fmt"

//查找数组中第K个最大的元素
func findKthLargest(nums []int, k int) int {

	return selectK(nums,0,len(nums)-1,len(nums)-k+1)
}
func selectK(arr []int, l, r ,k int) int{

	p := partition(arr, l, r)
	if p==k-1{
		return arr[p]
	}else if p>k-1{
		return  selectK(arr,l,p-1,k)
	}else{
		return selectK(arr,p+1,r,k)
	}


}
func partition(arr []int, l, r int) int{
	pivot := arr[l]
	i:=l+1
	j:=r
	for  {
		switch  {
		case i>j:
			arr[j],arr[l]=arr[l],arr[j]
			return j
		case arr[i]<pivot:
			i++
		case arr[j]>pivot:
			j--
		default:
			arr[i],arr[j]=arr[j],arr[i]
			i++
			j--

		}

	}

}

func main() {

	arr := []int{1,3,5,7,2,4,6,8}

	fmt.Println(findKthLargest(arr, 2))
}
