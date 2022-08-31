package main

import "fmt"

func binarySearch(arr []int,target int) int {
	l :=0
	r := len(arr)-1
	for{
		if l>r{
			return -1
		}
		mid := l+(r-l)/2
		if arr[mid]==target{
			return mid
		}else if arr[mid]>target{
			r = mid-1
		}else{
			l = mid+1
		}
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res :=binarySearch(arr,11)
	fmt.Println(res)

}
