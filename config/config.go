package config

import (

	"github.com/joho/godotenv"
	"os"
	"web-server/model"
)

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		panic("加载配置文件错误")
	}
	model.InitDatabase(os.Getenv("MYSQL_DSN"))
}
