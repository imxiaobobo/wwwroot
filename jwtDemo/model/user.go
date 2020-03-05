package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type UserModel struct {
	gorm.Model
	Db       *gorm.DB
	Name     string `gorm:"type:varchar(20);not null" form:"name" json:"name"`
	Phone    string `gorm:"type:varchar(11);not null;unique" form:"phone" json:"phone"`
	Password string `gorm:"type:varchar(255);not null" form:"password" json:"password"`
}

//创建用户
func (u *UserModel) CreateUser() {
	u.Db.Create(u)
}

//判断手机号
func (u *UserModel) IsPhoneExist() bool {
	u.Db.Where("phone = ?", u.Phone).First(u)
	fmt.Println(u)
	if u.ID != 0 {
		return true
	}
	return false
}
