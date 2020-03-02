/*
   @Time : 2020/2/28 19:52
   @Author : wangbo
   @File : stacker
*/
package Stacker

type Stack interface {
	GetSize() int
	IsEmpty() bool
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
}
