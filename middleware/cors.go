package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

)

func Cors() gin.HandlerFunc {

	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	config.AllowOrigins = []string{"http://localhost:8080",}

	return cors.New(config)
}