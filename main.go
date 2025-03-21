package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sakana9826/sakana-resume-backend/config"
	"github.com/sakana9826/sakana-resume-backend/routes"
)

func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// 初始化数据库连接
	config.InitDB()

	// 设置路由
	r := routes.SetupRouter()

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
