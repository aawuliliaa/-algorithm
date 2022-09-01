package main

import "fmt"

//查找大于target的最小索引
func upper(arr []int,target int)int  {
	l :=0
	r := len(arr)-1
	if arr[r]<=target{
		return -1
	}
	if arr[l]>target{
		return l
	}
	for{
		if l>=r{
			return l
		}
		mid := l+(r-l)/2
		if arr[mid]>target{
			r = mid
		}else{
			l = mid +1
		}
	}
}
func ceil(arr []int,target int)int  {
	p :=upper(arr,target)
	if p==-1||p==0||arr[p-1]!=target{
		return p
	}
	return p-1
}

func main()  {
	arr := []int{1,1,3,3,5,7}
	fmt.Println(ceil(arr,0))

}

