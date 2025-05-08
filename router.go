package main

import (
	"thunder_hoster/handler"
	"thunder_hoster/middleware"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/resources", "./resources")

	router.Use(middleware.FailedCountLimiter())

	router.GET("/", handler.MainPage)
	router.POST("/auth", handler.PasswordAuthenticator)

	mapGroup := router.Group("/map")
	mapGroup.Use(middleware.AccessControlMiddleware())
	{
		mapGroup.GET("/*path", handler.SendFile)
		mapGroup.POST("/*path", handler.SendFile)
	}

	adminGroup := router.Group("/admin")
	{
		adminGroup.GET("/upload", handler.UploadPage)
		adminGroup.POST("/upload", handler.UploadHandler)
	}

	return router
}
