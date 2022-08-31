package main

import "fmt"

func binarySearch(arr []int, target int) int {
	return search(arr,target,0, len(arr)-1)
}
func search(arr []int, target, l, r int) int{
	if l>r{
		return -1
	}
	mid := l+(r-l)/2
	if arr[mid]==target{
		return mid
	}else if arr[mid]>target{
		return search(arr,target,l,mid-1)

	}else{
		return search(arr,target,mid+1,r)
	}

}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res :=binarySearch(arr,7)
	fmt.Println(res)

}
