package main

import "fmt"

func sort(arr []int) []int {
	if len(arr)<2{
		return arr
	}
	mid := len(arr) / 2
	arrL := sort(arr[:mid])
	arrR := sort(arr[mid:])
	return merge(arrL, arrR)

}
func merge(arrL, arrR []int) []int {
	arrNew := make([]int, 0, len(arrL)+len(arrR))
	l := 0
	r := 0
	for l < len(arrL) && r < len(arrR) {
		if arrL[l] > arrR[r] {
			arrNew = append(arrNew, arrR[r])
			r++
		} else {
			arrNew = append(arrNew, arrL[l])
			l++
		}
	}
	if l < len(arrL) {
		arrNew = append(arrNew, arrL[l:]...)
	}
	if r < len(arrR) {
		arrNew = append(arrNew, arrR[r:]...)
	}
	return arrNew

}
func main() {
	arr := []int{3, 5, 6, 1,1,2, 2, 4, 8}

	fmt.Println(sort(arr))
}
