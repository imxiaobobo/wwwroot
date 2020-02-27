/*
   @Time : 2020/2/1 19:22
   @Author : wangbo
   @File : page
*/
package model

type Page struct {
	Book     []*Book
	PageNo   int64 //当前页
	PageSize int64 //每页显示的条数
	Count    int64 //总页数
	Total    int64 //总记录数
}

func (p *Page) IsHasPrev() bool {
	return p.PageNo > 1
}

func (p *Page) IsHasNext() bool {
	return p.PageNo < p.Total
}

func (p *Page) GetPrev() int64 {
	if p.IsHasPrev() {
		return p.PageNo - 1
	} else {
		return 1
	}
}

func (p *Page) GetNext() int64 {
	if p.IsHasNext() {
		return p.PageNo + 1
	} else {
		return p.Total
	}
}
