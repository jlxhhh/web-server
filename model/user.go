package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
}

//用接口来定义任意类型的参数
func GetUser(uid interface{}) (User ,error) {
	var user User
	result := DB.First(&user,uid)
	return user ,result.Error
}