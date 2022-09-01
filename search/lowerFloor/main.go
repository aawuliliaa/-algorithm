package main

import "fmt"

func lower(arr []int, target int) int {
	l := 0
	r := len(arr) - 1

	if arr[l] >= target {
		return -1
	}

	if arr[r] < target {
		return r
	}
	for {
		mid := l + (r-l+1)/2
		if l >= r {
			return l
		}
		if arr[mid] < target {
			l = mid
		} else {
			r = mid - 1
		}
	}
}
func lowerFloor(arr []int, target int) int {
	p := lower(arr, target)
	if  p == len(arr)-1 || arr[p+1] != target {
		return p
	}
	return p + 1

}

func main() {
	arr := []int{1, 1, 2, 2, 5, 5}
	fmt.Println(lowerFloor(arr, 5))
}
