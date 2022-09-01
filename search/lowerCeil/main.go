package main

import "fmt"


func lowerCeil(arr []int,target int) int {
	l :=0
	r := len(arr)-1
	if arr[l]>=target{
		return 0
	}
	if arr[r]<target{
		return -1
	}
	for{

		if l>=r{
			return l
		}
		mid := l+(r-l)/2
		//只是这里和upper不同
		if arr[mid]>=target{
			r = mid
		}else{
			l=mid+1
		}
	}

}
func main()  {
	arr := []int{1,2,2,5,5}
	fmt.Println(lowerCeil(arr,2))

}