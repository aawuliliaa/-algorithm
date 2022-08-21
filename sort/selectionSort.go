package main

import "fmt"

//先把最小的拿出来
//剩下的再把最小的拿出来
//每次选择还没处理的元素中的最小元素
func selectionSort(sourceList []int) []int {
	for i := 0; i < len(sourceList); i++ {

		for j := i + 1; j < len(sourceList); j++ {
			if sourceList[j] < sourceList[i] {
				temp := sourceList[i]
				sourceList[i] = sourceList[j]
				sourceList[j] = temp
			}
		}
	}
	return sourceList
}

func anotherWaySort(sourceList []int) []int {
	for i := len(sourceList)-1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if sourceList[j]>sourceList[i]{
				temp :=sourceList[i]
				sourceList[i] = sourceList[j]
				sourceList[j] = temp
			}
		}
	}
	return sourceList
}
func main() {
	//方式一：从前到后的排序
	sourceList := []int{3, 6, 2, 1, 10, 33, 22, 34, 45, 15}
	desList := selectionSort(sourceList)
	fmt.Println(desList)
	//方式二：从后到前的排序

	desListTwo := anotherWaySort(sourceList)
	fmt.Println(desListTwo)

}
