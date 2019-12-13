package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type ABC struct {
	gorm.Model
	id int
	user string
}

//全局的数据库实例
var DB *gorm.DB

func InitDatabase(connArgs string){
	db, err := gorm.Open("mysql", "root:123456@/web?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	//创建表，会自动转换为小写，并加s作为表名
	db.AutoMigrate(&User{})
//	db.AutoMigrate(&ABC{})
}