package api

import (
	"Muhammadjon/bootcamp/stadium/api/handlers"
	"Muhammadjon/bootcamp/stadium/api/handlers/docs"
	"Muhammadjon/bootcamp/stadium/config"

	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"     // swagger embed files
	// ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @description This is a sample article demo.
// @termsOfService https://udevs.io
func SetUpAPI(r *gin.Engine, h handlers.Handler, cfg config.Config) {
	docs.SwaggerInfo.Title = cfg.App
	docs.SwaggerInfo.Version = cfg.Version
	docs.SwaggerInfo.Host = cfg.ServiceHost + cfg.HTTPPort
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.POST("/stadium", h.CreateStadium)
	r.GET("/stadium", h.GetStadiumList)
	// r.GET("/actor/:id", h.GetById)

	// // movie
	// r.POST("/movie", h.CreateMovie)
	// r.GET("/movie", h.GetMovie)
	// r.GET("/min", h.GetMinPublisher)

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
