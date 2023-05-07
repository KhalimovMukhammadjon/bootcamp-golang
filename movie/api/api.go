package api

import (
	"Muhammadjon/bootcamp/movie/api/docs"
	"Muhammadjon/bootcamp/movie/api/handlers"
	"Muhammadjon/bootcamp/movie/config"

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

	r.POST("/actor", h.CreateActor)
	r.GET("/actor", h.GetList)
	r.GET("/actor/:id", h.GetById)
	r.PUT("/actor/update", h.Update)

	// movie
	r.POST("/movie", h.CreateMovie)
	r.GET("/movie", h.GetMovie)
	// r.GET("/min", h.GetMinPublisher)

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
