/*
   @Time : 2020/2/28 19:55
   @Author : wangbo
   @File : arrStack
   @Project ：以动态数组实现栈
*/
package ArrayStack

import (
	"data/Array"
)

type Stack struct {
	arr *Array.Array
}

//实例化stack
func NewStack(cap int) *Stack {
	return &Stack{Array.NewArr(cap)}
}

//获取元素个数
func (s *Stack) GetSize() int {
	return s.arr.GetArrSize()
}

//判断是否为空
func (s *Stack) IsEmpty() bool {
	return s.arr.IsEmpty()
}

//获取容量
func (s *Stack) GetCap() int {
	return s.arr.GetArrCap()
}

//向栈中压入元素
func (s *Stack) Push(i interface{}) {
	s.arr.AddLast(i)
}

func (s *Stack) Pop() interface{} {
	return s.arr.DelFirst
}
