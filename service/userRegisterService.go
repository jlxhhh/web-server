package service

import (
	"web-server/model"
	"web-server/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Username string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// valid 验证表单
func (service *UserRegisterService) valid() *serializer.Response {
	count := 0
	model.DB.Model(&model.User{}).Where("username = ?", service.Username).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Data:nil,
			Msg:  "用户名已经注册",
			Error: "",
		}
	}
	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() serializer.Response {
	user := model.User{
		Username: service.Username,
		Password:service.Password,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}
	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Data:  nil,
			Msg:   "注册失败",
			Error: "",
		}
	}
	return serializer.BuildUserResponse(user)
}
