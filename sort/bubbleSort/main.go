package main

import "fmt"

func sort(arr []int) []int {
	//i表示已经有多少个排好序了
	for i:=0;i< len(arr);i++{
		for j :=0;j<len(arr)-i-1;j++{
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
	return arr

}

func main()  {
	arr :=[]int{4,6,5,7,9,8}
	arr = sort(arr)
	fmt.Println(arr)

}
