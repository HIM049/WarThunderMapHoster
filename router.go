package main

import (
	"thunder_hoster/config"
	"thunder_hoster/handler"
	"thunder_hoster/middleware"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.Use(middleware.FailedCountLimiter())

	router.GET("/", handler.MainPage)
	router.POST("/", handler.PasswordAuthenticator)

	mapGroup := router.Group(config.Cfg.DownloadRouter)
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
