package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"jwtDemo/model"
)

var Db *gorm.DB

func init() {
	driverName := "mysql"
	host := "localhost"
	port := "8889"
	database := "jwt"
	userName := "root"
	password := "root"
	charset := "utf8"
	str := "%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true"
	args := fmt.Sprintf(str, userName, password, host, port, database, charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.UserModel{})
	Db = db
}

func NewDB() *gorm.DB {
	return Db
}
