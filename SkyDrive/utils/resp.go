/*
   @Time : 2020/2/20 20:31
   @Author : wangbo
   @File : resp
*/
package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

//ResMsg 消息结构体
type ResMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//NewResMsg 创建消息结构体
func NewResMsg(code int, msg string, data interface{}) *ResMsg {
	return &ResMsg{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

//JsonBytes 对象
func (res *ResMsg) JsonBytes() ([]byte, error) {
	r, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return r, nil
}

// JSONString : 对象转json格式的string
func (res *ResMsg) JSONString() (string, error) {
	r, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(r), nil
}
