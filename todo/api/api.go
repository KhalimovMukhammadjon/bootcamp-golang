package api

import (
	"Muhammadjon/bootcamp/todo/api/docs"
	// _ "Muhammadjon/bootcamp/todo/api/docs"
	"Muhammadjon/bootcamp/todo/api/handlers"
	"Muhammadjon/bootcamp/todo/config"

	"github.com/gin-gonic/gin"
	// ginSwagger "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/gin-swagger/swaggerFiles"
)

// @description This is a sample article demo.
// @termsOfService https://udevs.io
func SetUpAPI(r *gin.Engine, h handlers.Handler, cfg config.Config) {
	docs.SwaggerInfo.Title = cfg.App
	docs.SwaggerInfo.Version = cfg.Version
	docs.SwaggerInfo.Host = cfg.ServiceHost + cfg.HTTPPort
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.POST("/todo", h.CreateTodo)
	//r.GET("/todo", h.GetArticleList)
	// r.GET("/articles/:id", GetByIDHandler)
	// r.PUT("/articles/:id", UpdateHandler)
	// r.DELETE("/articles/:id", DeleteHandler)

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
