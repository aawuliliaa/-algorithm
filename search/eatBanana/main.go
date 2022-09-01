package main

import "fmt"

func minEatingSpeed(piles []int, h int) int {
	l :=1
	r :=getMaxPile(piles)
	for {
		if l>=r{
			return l
		}
		mid := l+(r-l)/2
		if eatTime(piles,mid)>h{
			l = mid+1
		}else{
			r= mid
		}
	}

}
func eatTime(piles []int, speed int) int {
	time := 0
	for _, pile := range piles {
		if pile%speed == 0 {
			time += pile / speed
		} else {
			time += pile/speed + 1
		}
	}
	return time
}
func getMaxPile(piles []int) int {

	pivot := piles[0]
	for i := 1; i < len(piles); i++ {
		if piles[i] > pivot {
			pivot = piles[i]
		}
	}
	return pivot
}
func main() {
	arr := []int{30,11,23,4,20}
	fmt.Println(minEatingSpeed(arr, 5))
}