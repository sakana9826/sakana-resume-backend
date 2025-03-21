package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sakana9826/sakana-resume-backend/controllers"
	"github.com/sakana9826/sakana-resume-backend/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 添加 CORS 中间件
	r.Use(middleware.Cors())

	api := r.Group("/api")
	{
		api.POST("/generate-access-code", controllers.GenerateAccessCode)
		api.POST("/verify-access-code", controllers.VerifyAccessCode)
	}

	return r
}
