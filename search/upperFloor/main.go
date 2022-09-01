package main

import "fmt"

func upperFloor(arr []int,target int) int {
	l := 0
	r := len(arr)-1
	if arr[r]<=target{
		return r
	}
	if arr[l]>target{
		return -1
	}
	for {
		if l>=r{
			return l
		}
		mid := l+(r-l+1)/2
		if arr[mid]>target{
			r=mid-1
		}else{
			l = mid
		}
	}
}

func main()  {
	arr := []int{1,2,2,5,5}
	fmt.Println(upperFloor(arr,3))
}
