package main

import "fmt"

//查找大于target的最小值
func search(arr []int,target int)int  {
	l :=0
	r := len(arr)-1
	if arr[r]<=target{
		return -1
	}
	if arr[l]>target{
		return arr[l]
	}
	for{
		if l>=r{
			return arr[r]
		}
		mid := l+(r-l)/2
		if arr[mid]>target{
			r = mid
		}else{
			l = mid +1
		}
	}
}

func main()  {
	arr := []int{1,1,3,3,5,7}
	fmt.Println(search(arr,2))

}
