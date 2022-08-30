package main

import (
	"fmt"
)

func sort(arr []int, begin, end int) {
	if begin >= end {
		return
	}
	mid := (begin + end) / 2
	sort(arr, begin, mid)
	sort(arr, mid+1, end)
	//if arr[mid] > arr[mid+1] {
	//	merge(arr, begin, mid, end)
	//}

}
func merge(arr []int, begin, mid, end int) {
	temp := make([]int, len(arr))       //开辟临时数组
	pLeft := begin                      //左边指针
	pRight := mid + 1                   //右边指针
	pTemp := begin                      //记录临时数组的当前指针
	for pLeft <= mid && pRight <= end { //遍历两个数组
		if arr[pLeft] < arr[pRight] { //较小的元素先进入
			temp[pTemp] = arr[pLeft]
			pLeft++ //左边指向下一位
		} else {
			temp[pTemp] = arr[pRight]
			pRight++ //右边指向下一位
		}
		pTemp++
	}
	for pLeft <= mid { //左边继续合并
		temp[pTemp] = arr[pLeft]
		pLeft++
		pTemp++

	}
	for pRight <= end { //右边继续合并
		temp[pTemp] = arr[pRight]
		pRight++
		pTemp++

	}
	//临时数组中数据放到arr中
	for i := begin; i <= end; i++ {
		arr[i] = temp[i]
	}

}
func main() {
	arr := []int{3, 5, 6, 1, 1, 2, 2, 4, 8}
	sort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
