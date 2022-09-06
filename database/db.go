package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

// 连接数据库

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gdopo?charset=utf8&parseTime=true")
	if err != nil {
		fmt.Println("connect to database error.")
	}
	DB.SingularTable(true)

}
