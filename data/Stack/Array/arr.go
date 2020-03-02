/*
   @Time : 2020/2/23 13:55
   @Author : wangbo
   @File : arr
*/
package Array

import (
	"bytes"
	"fmt"
)

type Array struct {
	data []interface{} //切片
	size int           //元素个数,下标
}

//新建一个数组
func NewArr(cap int) *Array {
	return &Array{
		data: make([]interface{}, cap),
	}
}

//获取数据的容量
func (a *Array) GetArrCap() int {
	return len(a.data)
}

//获取数据的个数
func (a *Array) GetArrSize() int {
	return a.size
}

//判断是否为空
func (a *Array) IsEmpty() bool {
	return a.size == 0
}

//向数组尾部添加元素
func (a *Array) AddLast(e interface{}) {
	a.Add(a.size, e) //因为index大于size就会报错,所以只有size等于index的时候才会直接添加
}

//向数组头部添加元素
func (a *Array) AddFirst(e int) {
	a.Add(0, e)
}

//向数组指定位置添加元素
func (a *Array) Add(index int, e interface{}) {
	if a.size == len(a.data) { //如果数组的元素个数和长度相等,说明这个数组已经满了
		a.reSize(len(a.data) * 2)
	}
	if index < 0 || index > a.size { //如果传入的位置小于0或者大于目前下标所在位置
		panic("index < 0 || index >= a.size")
	}
	//从当前下标前一个开始,到index结束,数组的前一个值后移 arr[1,0，0，0，0]
	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[index] = e
	a.size++
}

//取出数组元素
func (a *Array) Get(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("index < 0 || index >= a.size")
	}
	return a.data[index]
}

//更新数组元素
func (a *Array) Set(index, e int) {
	if index < 0 || index >= a.size {
		panic("index < 0 || index >= a.size")
	}
	a.data[index] = e
}

//数组中是否包含该元素
func (a *Array) Contains(e int) bool {
	for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			return true
		}
	}
	return false
}

//返回索引
func (a *Array) Find(e int) int {
	for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			return i
		}
	}
	return -1
}

//删除所有[5,4,2,2] 5 4 2    i = 2 s = 3
func (a *Array) DelAll(e int) {
	for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			a.Del(i)
			i-- //需要i-- 是因为 每次删除一个元素,索引就会少一个,如果不减去一个i直接加就会少算一个
		}
	}
}

//删除指定位置的元素
func (a *Array) Del(index int) {
	if index < 0 || index >= a.size { //如果传入的位置小于0或者大于目前下标所在位置
		panic("index < 0 || index >= a.size")
	}
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}
	a.size--
	a.data[a.size] = nil

	//前面做了删除操作之后,如果该数组的元素个数只有整个数组长度的一办,那么就把数组的长度砍掉一半
	if a.size == len(a.data)/2 {
		a.reSize(len(a.data) / 2)
	}
}

//给定某个数,有就删除
func (a *Array) DelElement(e int) {
	index := a.Find(e)
	if index != -1 {
		a.Del(index)
	}
}

//删除首元素
func (a *Array) DelFirst() {
	a.Del(0)
}

//删除尾部元素
func (a *Array) DelLast() {
	a.Del(a.size - 1)
}

//数组扩容(让数组拥有动态数组的能力)
func (a *Array) reSize(newCapacity int) {
	newData := make([]interface{}, newCapacity) //make一个新的数组
	for i := 0; i < a.size; i++ {               //将原本数组里的值赋值到新数组里
		newData[i] = a.data[i]
	}
	a.data = newData //再将新数组赋值给老数组
}

//重写array的string方法
func (a *Array) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", a.size, len(a.data)))
	buf.WriteString("[")
	for i := 0; i < a.size; i++ {
		buf.WriteString(fmt.Sprintf("%v", a.data[i]))
		if i != (a.size - 1) {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("]")
	return buf.String()
}
