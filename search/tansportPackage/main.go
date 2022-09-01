package main

import "fmt"

func shipWithinDays(weights []int, days int) int {
	l := getMaxWight(weights)
	r := sumWeight(weights)
	for {
		if l >= r {
			return l
		}
		mid := l + (r-l)/2
		daysNeeded := daysNeed(weights, mid)
		if daysNeeded > days {
			l = mid + 1
		} else {
			r = mid
		}

	}
}
func daysNeed(weights []int, shipWeight int) int {
	sum := 0
	days := 0
	for i := 0; i < len(weights); i++ {
		sum = sum + weights[i]
		if sum > shipWeight {
			days =days+1
			sum = weights[i]
		}

	}
	days =days+1
	return days

}
func sumWeight(weights []int) int {
	sum := 0
	for i := 0; i < len(weights); i++ {
		sum += weights[i]
	}
	return sum
}
func getMaxWight(weights []int) int {
	pivot := weights[0]
	for i := 1; i < len(weights); i++ {
		if weights[i] > pivot {
			pivot = weights[i]
		}
	}
	return pivot
}
func main() {
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(shipWithinDays(arr, 5))
	//fmt.Println(daysNeed(arr,13))
}
