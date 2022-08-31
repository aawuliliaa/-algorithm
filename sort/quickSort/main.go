package main

import "fmt"

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(arr, l, r)
	quickSort(arr, l, p-1) //这里是到p-1
	quickSort(arr, p+1, r)//这里是从p+1开始
}
func partition(arr []int, l, r int) int {
	// 以当前数据序列最后一个元素作为初始 pivot
	pivot := arr[l]
	j := l
	for i := l+1; i <=r; i++ {
		if arr[i] < pivot {
			j++
			arr[i], arr[j] = arr[j], arr[i]

		}
	}
	arr[j], arr[l] = arr[l], arr[j]
	return j
}
func main() {
	arr := []int{3, 5, 6, 1, 1, 2, 2, 4, 8}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)

}
