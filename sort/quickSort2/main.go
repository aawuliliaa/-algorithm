package main

import "fmt"

func dubWaySort(arr []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(arr, l, r)
	dubWaySort(arr, l, p-1)
	dubWaySort(arr, p+1, r)
}
func partition(arr []int, l, r int) int {
	pivot := arr[l]
	//arr[l+1...i-1]<=v  arr[j+1...r]>=v
	i := l + 1
	j := r
	for {
		switch {
		case i > j:
			arr[j], arr[l] = arr[l], arr[j]
			return j
		case arr[i] < pivot:
			i++
		case arr[j] > pivot:
			j--

		default:
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
}

func main() {
	arr := []int{3, 5, 6, 1, 1, 2, 2, 4, 8}
	dubWaySort(arr, 0, len(arr)-1)
	fmt.Println(arr)

}
