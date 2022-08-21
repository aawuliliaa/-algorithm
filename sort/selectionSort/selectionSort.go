package main

import "fmt"

//选择排序----算法复杂度=O(n^2)
//先把最小的拿出来
//剩下的再把最小的拿出来
//每次选择还没处理的元素中的最小元素
func selectionSort(sourceList []int) []int {
	for i := 0; i < len(sourceList); i++ {

		for j := i + 1; j < len(sourceList); j++ {
			if sourceList[j] < sourceList[i] {
				sourceList[j], sourceList[i] = sourceList[i], sourceList[j]
			}
		}
	}
	return sourceList
}

func anotherWaySort(sourceList []int) []int {
	for i := len(sourceList) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if sourceList[j] > sourceList[i] {
				sourceList[j], sourceList[i] = sourceList[i], sourceList[j]
			}
		}
	}
	return sourceList
}
func main() {
	//方式一：从前到后的排序,前面是排序好的，每次从后面未排序部分中找出最小的值放到前面
	sourceList := []int{3, 6, 2, 1, 10, 33, 22, 34, 45, 15}
	desList := selectionSort(sourceList)
	fmt.Println(desList)

	//方式二：从后到前的排序，后面是排序好的，每次从前面未排序部分中找出最小的值放到后面
	desListTwo := anotherWaySort(sourceList)
	fmt.Println(desListTwo)

}
