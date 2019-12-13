package router

import (
	"github.com/gin-gonic/gin"
	"web-server/api"
	"web-server/middleware"

)

func NewRouter() *gin.Engine{
	r := gin.New()
	//gin.Default()包含下面两个中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(middleware.Session())
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

		//r.Use(middleware.CurrentUser())
	v1 := r.Group("/api/v1")
	{
		v1.POST("user/register",api.UserRegister)
		v1.POST("user/login",api.UserLogin)
		v1.GET("ping",api.Ping)
	}

	return r
}
