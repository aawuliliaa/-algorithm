package main

import "fmt"

func insertSort(sourceList []int) []int {
	for i := 1; i < len(sourceList); i++ {
		for j := i; j > 0; j-- {
			if sourceList[j] < sourceList[j-1] {
				sourceList[j], sourceList[j-1] = sourceList[j-1], sourceList[j]
			}
		}
	}
	return sourceList
}

//从前向后的排序
func insertSortBetter(sourceList []int) []int {
	for i := 1; i < len(sourceList); i++ {
		itemToCompare := sourceList[i]
		var j int
		for j = i; j > 0; j-- {
			if sourceList[j-1] > itemToCompare {
				sourceList[j] = sourceList[j-1]
			} else {
				break

			}
		}
		sourceList[j] = itemToCompare
	}
	return sourceList
}

//从后向前的排序
func insertSortAnother(sourceList []int) []int {
	for i := len(sourceList) - 1; i >= 0; i-- {
		itemToCompare := sourceList[i]
		var j int
		for j = i; j+1 < len(sourceList); j++ {
			if sourceList[j+1] < itemToCompare {
				sourceList[j] = sourceList[j+1]
			} else {
				break
			}
		}
		sourceList[j] = itemToCompare
	}
	return sourceList
}
func main() {
	sourceList := []int{3, 6, 2, 1, 10, 33, 22, 34, 45, 15}
	desList := insertSortAnother(sourceList)
	fmt.Println(desList)
}
