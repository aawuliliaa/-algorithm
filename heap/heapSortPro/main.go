package main

import "fmt"

func sort(arr []int) []int {
	//首先把任意数组变成大顶堆
	for i := (len(arr) - 1 - 1) / 2; i >= 0; i-- {
		siftDown(arr, i, len(arr))
	}
	//取出堆顶元素
	for i := len(arr) - 1; i >= 0; i-- {
		//把堆顶元素放在数组末尾
		arr[i], arr[0] = arr[0], arr[i]
		//数组中剩下的内容形成大顶堆
		siftDown(arr, 0, i)
	}
	return arr

}

//对数组arr[0,end）所形成的最大堆中，索引index的元素执行siftDown
func siftDown(arr []int, index, indexEnd int) {
	for index*2+1 < indexEnd {
		j := index*2 + 1
		if j+1 < indexEnd && arr[j] < arr[j+1] {
			j++
		}
		if arr[index] < arr[j] {
			arr[index], arr[j] = arr[j], arr[index]
			index = j
		} else {
			break
		}
	}
}
func main() {
	arr := []int{3, 5, 7, 8, 4, 6, 2, 1}
	fmt.Println(sort(arr))

}
