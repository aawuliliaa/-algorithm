package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func product(ch chan int)  {

	for i:=0;i<5;i++{
		ch<-i
	}
	close(ch)

}
func consumer(ch chan int)  {

	for {
		val ,ok :=<-ch
		if !ok{
			break
		}else{
			fmt.Println("consume---",val)
		}
	}


}
func main()  {
	ch := make(chan int,3)

	wg.Add(2)
	go func() {
		defer wg.Done()
		product(ch)
	}()
	go func() {
		defer wg.Done()
		consumer(ch)
	}()
	wg.Wait()
	fmt.Println("finish")
}