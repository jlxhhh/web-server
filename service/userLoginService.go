package service

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"web-server/model"
	"web-server/serializer"
)



// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	Username string `form:"username" json:"username" binding:"required,min=5,max=10"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=20"`
}

func (service *UserLoginService) setSession(c *gin.Context, user model.User)  {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}


func (service *UserLoginService) Login(c *gin.Context) serializer.Response{
	var user model.User
	usernameErr := model.DB.Where("username = ?",service.Username).Find(&user).Error
	if usernameErr != nil {
		fmt.Print(usernameErr)
		return serializer.Response{
			Code:  0,
			Data:  nil,
			Msg:   "用户名或者密码错误",
			Error: "",
		}
	}
	model.DB.Select("password").Where("username = ?",service.Username).Find(&user)

	if user.Password != service.Password {
		return serializer.Response{
			Code:  0,
			Data:  nil,
			Msg:   "用户名或者密码错误",
			Error: "",
		}
	}
	service.setSession(c, user)
	return serializer.BuildUserResponse(user)
}