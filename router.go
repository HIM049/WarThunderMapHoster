package main

import (
	"thunder_hoster/config"
	"thunder_hoster/handler"
	"thunder_hoster/middleware"
	"thunder_hoster/pages"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.Use(middleware.FailedCountLimiter())

	router.GET("/", pages.PageMain)
	router.POST("/", handler.PasswordAuthenticator)

	mapGroup := router.Group(config.Cfg.DownloadRouter)
	mapGroup.Use(middleware.AccessControlMiddleware())
	{
		router.GET("/list", pages.PageMapList)
		mapGroup.GET("/:map", handler.SendFile)
		mapGroup.POST("/:map", handler.SendFile)
	}

	adminGroup := router.Group("/admin")
	{
		adminGroup.GET("/", pages.UploadPage)
		adminGroup.POST("/upload", handler.UploadHandler)
		adminGroup.POST("remove", handler.RemoveMap)
	}

	return router
}
