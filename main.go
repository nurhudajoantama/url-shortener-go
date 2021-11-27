package main

import (
	"url-shortener/config"
	"url-shortener/handler"
	"url-shortener/helper"
	"url-shortener/middleware"
	"url-shortener/repository"
	"url-shortener/service"

	"github.com/gin-gonic/gin"
)

var (
	database        = config.SetupDatabaseConnection()
	urlRepository   = repository.NewUrlRepository(database)
	userRepository  = repository.NewUserRepository(database)
	urlService      = service.NewUrlService(urlRepository)
	authService     = service.NewAuthService(userRepository)
	jwtService      = service.NewJwtService()
	redirectHandler = handler.NewRedirectHandler(urlService)
	urlHandler      = handler.NewUrlHandler(urlService)
	authHandler     = handler.NewAuthHandler(authService, jwtService)

	jwtMiddleware = middleware.NewAuthMiddleware(authService, jwtService)
)

func main() {

	defer config.CloseDatabaseConnection(database)

	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("login", authHandler.Login)
		api.POST("register", authHandler.Register)

		authorized := api.Group("/")
		authorized.Use(jwtMiddleware.JwtAuthMiddleware)
		{
			authorized.POST("/shorten", urlHandler.Create)
			authorized.PUT("/shorten/:id", urlHandler.Update)
			authorized.DELETE("/shorten/:id", urlHandler.Delete)
		}

	}
	router.GET("/:url", redirectHandler.RedirectUrl)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, helper.NotFoundResponse)
	})

	router.Run(":8080")
}
