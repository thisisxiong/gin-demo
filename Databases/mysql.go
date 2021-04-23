package Databases

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB
func init() {
	var err error
	Db, err := gorm.Open("mysql","root:root@localhost:3306/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("mysql connect error %v",err)
	}

	if Db.Error != nil {
		fmt.Printf("database error %v",Db.Error)
	}

	Db.SingularTable(true) // 全局禁用表名复数

}
