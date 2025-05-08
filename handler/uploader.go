package handler

import (
	"net/http"
	"thunder_hoster/config"

	"github.com/gin-gonic/gin"
)

func UploadPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "upload.tmpl", gin.H{
		"title": "File Uploader",
	})
}

func UploadHandler(ctx *gin.Context) {
	passwd := ctx.PostForm("password")

	if passwd != config.Cfg.Security.AdminPasswd {

		ctx.HTML(http.StatusForbidden, "message.tmpl", gin.H{
			"title":       "Wrong Password",
			"message":     "Wrong Password",
			"description": "",
			"color":       "red",
		})
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "message.tmpl", gin.H{
			"title":       "Upload Failed",
			"message":     "Upload Failed",
			"description": "Error: " + err.Error(),
			"color":       "red",
		})
		return
	}

	err = ctx.SaveUploadedFile(file, config.Cfg.Service.FilePath)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "message.tmpl", gin.H{
			"title":       "Upload Failed",
			"message":     "Failed when save file",
			"description": "Error: " + err.Error(),
			"color":       "red",
		})
		return
	}

	ctx.HTML(http.StatusOK, "message.tmpl", gin.H{
		"title":       "Upload Successed",
		"message":     "Upload Successed",
		"description": "",
		"color":       "green",
	})
}
