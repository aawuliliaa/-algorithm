package main

import (
	"bytes"
	"fmt"
)

type Array struct {
	data []interface{}
	size int
}

func New(capacity int) *Array {
	return &Array{
		data: make([]interface{}, capacity),
		size: 0,
	}
}
func (a *Array) GetCapacity() int {
	return len(a.data)
}
func (a *Array) GetSize() int {
	return a.size
}
func (a *Array) IsEmpty() bool {
	return a.size == 0
}
func (a *Array) Resize(newCapacity int) {
	newData := make([]interface{}, newCapacity)
	for i, item := range a.data {
		newData[i] = item
	}
	a.data = newData
}
func (a *Array) AddLast(e interface{}) {
	a.Add(a.size, e)
}
func (a *Array) AddFist(e interface{}) {
	a.Add(0, e)
}
func (a *Array) Get(index int) interface{} {
	return a.data[index]
}
func (a *Array) Set(index int, e interface{}) {
	a.data[index] = e
}
func (a *Array) Add(index int, e interface{}) {
	if index < 0 || index > a.size {
		panic("add error,index out of range")
	}
	//如果数组放满了，就扩容一倍
	if index == a.size {
		a.Resize(2 * len(a.data))
	}
	for i := a.size; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[index] = e
	a.size++
}
func (a *Array) RemoveLast() {
	a.Remove(a.size - 1)
}
func (a *Array) RemoveFist() {
	a.Remove(0)
}
func (a *Array) Remove(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("index out of range")
	}
	e := a.data[index]
	for i := index + 1; i <= a.size; i++ {
		a.data[i-1] = a.data[i]
	}
	a.size--
	a.data[a.size] = nil
	if a.size == len(a.data)/4 && len(a.data)/2 != 0 {
		a.Resize(len(a.data) / 2)
	}

	return e

}
// 重写 Array 的 string 方法
func (a *Array) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("array: size = %d, capacity = %d\n", a.size, len(a.data)))
	buffer.WriteString("[")
	for i := 0; i < a.size; i++ {
		// fmt.Sprint 将 interface{} 类型转换为字符串
		buffer.WriteString(fmt.Sprint(a.data[i]))
		if i != (a.size - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}
func main() {
	a :=New(10)
	a.AddLast(1)
	fmt.Println(a)
}
