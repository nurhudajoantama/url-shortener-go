package main

import (
	"url-shortener/config"
	"url-shortener/handler"
	"url-shortener/helper"
	"url-shortener/repository"
	"url-shortener/service"

	"github.com/gin-gonic/gin"
)

var (
	database        = config.SetupDatabaseConnection()
	urlRepository   = repository.NewUrlRepository(database)
	urlService      = service.NewUrlService(urlRepository)
	redirectHandler = handler.NewRedirectHandler(urlService)
	urlHandler      = handler.NewUrlHandler(urlService)
)

func main() {

	defer config.CloseDatabaseConnection(database)

	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/shorten", urlHandler.Create)
		api.PUT("/shorten/:id", urlHandler.Update)
		api.DELETE("/shorten/:id", urlHandler.Delete)
	}
	router.GET("/:url", redirectHandler.RedirectUrl)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, helper.NotFoundResponse)
	})

	router.Run(":8080")
}
