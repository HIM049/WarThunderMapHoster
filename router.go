package main

import (
	"github.com/gin-contrib/multitemplate"
	"thunder_hoster/config"
	"thunder_hoster/handler"
	"thunder_hoster/middleware"
	"thunder_hoster/pages"

	"github.com/gin-gonic/gin"
)

func NewTemplateRender() multitemplate.Renderer {
	r := multitemplate.New()
	r.AddFromFiles("index", "templates/index.tmpl", "templates/base.tmpl")
	r.AddFromFiles("login", "templates/login.tmpl", "templates/base.tmpl")
	r.AddFromFiles("list", "templates/list.tmpl", "templates/base.tmpl")
	return r
}

func RouterSetup() *gin.Engine {

	router := gin.Default()
	router.HTMLRender = NewTemplateRender()

	router.Use(middleware.FailedCountLimiter())

	router.GET("/", pages.PageMain)

	router.GET("/login", pages.PageLogin)

	pageGroup := router.Group("/pages")
	pageGroup.Use(middleware.LoginCheckMiddleware())
	{
		pageGroup.GET("/list", pages.PageMapList)
	}

	// 文件下载路由
	mapGroup := router.Group(config.Cfg.DownloadRouter)
	mapGroup.Use(middleware.DownloadControlMiddleware())
	{
		mapGroup.GET("/:map", handler.SendFile)
		mapGroup.POST("/:map", handler.SendFile)
	}

	apiGroup := router.Group("/api")
	{
		apiGroup.POST("/login", handler.AuthHandler)
		apiGroup.POST("/upload", handler.UploadHandler)
		apiGroup.POST("/delete", handler.DeleteHandler)
	}

	return router
}
