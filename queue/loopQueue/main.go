package main

import "fmt"

type LoopQueue struct {
	data  []interface{}
	front int
	tail  int
	size  int
}

func New(capacity int) (lq *LoopQueue) {
	lq = &LoopQueue{}
	// 队列存在以空位，所以参数为20就需要开辟21的位置
	lq.data = make([]interface{}, capacity+1)
	lq.front = 0
	lq.tail = 0
	lq.size = 0

	return
}

func (lq *LoopQueue) GetCapacity() int {
	// 与上对应，需要删除空位
	return len(lq.data) - 1
}

func (lq *LoopQueue) GetSize() int {
	return lq.size
}

func (lq *LoopQueue) IsEmpty() bool {
	return lq.front == lq.tail
}

func (lq *LoopQueue) Enqueue(e interface{}) {
	cp := lq.GetCapacity()
	//已经满了，扩容
	if (lq.tail+1)%len(lq.data) == lq.front {
		lq.resize(cp * 2)
	}
	//
	lq.data[lq.tail] = e
	lq.tail = (lq.tail + 1) % len(lq.data)
	lq.size += 1

}

func (lq *LoopQueue) Eequeue() interface{} {
	if lq.IsEmpty() {
		panic("lq is empty")
	}

	e := lq.data[lq.front]
	lq.data[lq.front] = nil
	lq.front = (lq.front + 1) % len(lq.data)
	lq.size--
	if lq.size == lq.GetCapacity()/4 && lq.GetCapacity()/2 != 0 {
		lq.resize(lq.GetCapacity() / 2)
	}
	return e

}
func (lq *LoopQueue) resize(capacity int) {
	newList := make([]interface{}, capacity)
	for i := 0; i <= lq.size; i++ {
		newList[i]=lq.data[(lq.front+i)%len(lq.data)]
	}
	lq.data = newList
	lq.front=0
	lq.tail=lq.size
}

func (lq *LoopQueue) GetFront() interface{} {
	return lq.data[lq.front]
}
func main()  {
	lq :=New(3)
	for _,i :=range []int{3,5,7,3,4,5}{
		lq.Enqueue(i)
		fmt.Println(lq)
	}
	for i:=0;i<3;i++{
		lq.Eequeue()
		fmt.Println(lq)
	}

}