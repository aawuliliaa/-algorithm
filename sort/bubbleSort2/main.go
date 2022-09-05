package main

import "fmt"

func sort(arr []int) []int {
	//i表示已经有多少个排好序了
	for i:=0;i< len(arr);i++{
		isSwapped :=false
		for j :=0;j<len(arr)-i-1;j++{
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
				isSwapped =true
			}
		}
		//没有数据交换，说明全部都是有序的了，不需要排序了
		if !isSwapped{
			return  arr
		}
	}
	return arr

}

func main()  {
	arr :=[]int{4,5,6,7,8,9}
	arr = sort(arr)
	fmt.Println(arr)

}
