package main

import "fmt"

func quickSort3Way(arr []int, l, r int) {
	if l >= r {
		return
	}
	lt, gt := partition(arr, l, r)
	quickSort3Way(arr, l, lt-1)
	quickSort3Way(arr, gt, r)
}
func partition(arr []int, l, r int) (lt, gt int) {
	//arr[l+1,lt]<v arr[lt+1,i-1]=v arr[gt,r]>v
	//下面i和j的初始值 可使上面的数组左边界大于右边界
	i := l + 1
	lt = l
	gt = r + 1
	pivot := arr[l]
	for {
		switch {
		case i >= gt:
			arr[l], arr[lt] = arr[lt], arr[l]
			return lt,gt
		case arr[i] < pivot:
			lt++
			arr[i], arr[lt] = arr[lt], arr[i]
			i++
		case arr[i] > pivot:
			gt--
			arr[i], arr[gt] = arr[gt], arr[i]
		default://arr[i] = pivot
			i++
		}
	}
	return lt, gt
}

func main() {
	arr := []int{3, 5, 6, 1, 1, 2, 2, 4, 8}
	quickSort3Way(arr, 0, len(arr)-1)
	fmt.Println(arr)

}
