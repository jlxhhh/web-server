package api

import (
	"github.com/gin-gonic/gin"
	"web-server/serializer"
	"web-server/service"
)


func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(0, serializer.Response{
			Data:  nil,
			Msg:   "注册错误",
			Error: "",
		})
	}
}
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(0, serializer.Response{
			Data:  nil,
			Msg:   "登录错误",
			Error: "",
		})
	}
}


