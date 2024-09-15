package routes

import (
	"21BLC1564/controllers"
	"21BLC1564/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	api := r.Group("/api", middleware.AuthMiddleware())
	{
		api.POST("/upload", controllers.UploadFile)
		api.GET("/files", controllers.ListFiles)
		api.GET("/share/:ID", middleware.AuthMiddleware(), controllers.ShareFile)
		api.GET("/search/files", middleware.AuthMiddleware(), controllers.SearchFiles)

	}
	

	return r
}
