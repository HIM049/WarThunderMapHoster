package handler

import (
	"errors"
	"net/http"
	"path/filepath"
	"thunder_hoster/config"
	"thunder_hoster/storage"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "upload.tmpl", gin.H{
		"title": "File Uploader",
	})
}

func UploadHandler(ctx *gin.Context) {
	mapName := ctx.PostForm("name")
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

	mapPath := filepath.Join(config.Cfg.Service.MapDir, mapName+".blk")
	err = ctx.SaveUploadedFile(file, mapPath)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "message.tmpl", gin.H{
			"title":       "Upload Failed",
			"message":     "Failed to save file",
			"description": "Error: " + err.Error(),
			"color":       "red",
		})
		return
	}

	newMap := storage.MapInformation{
		MapName:    mapName,
		FilePath:   mapPath,
		UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	err = storage.Storage.Add(&newMap)
	if err != nil {
		if errors.Is(err, storage.ErrDuplicatedName) {
			ctx.HTML(http.StatusInternalServerError, "message.tmpl", gin.H{
				"title":       "Upload Failed",
				"message":     "Map name already exists",
				"description": "Error: " + err.Error(),
				"color":       "red",
			})
		} else {
			ctx.HTML(http.StatusInternalServerError, "message.tmpl", gin.H{
				"title":       "Upload Failed",
				"message":     "Failed to save storage file",
				"description": "Error: " + err.Error(),
				"color":       "red",
			})
		}
		return
	}

	storage.Storage.GenerateIndex()

	ctx.HTML(http.StatusOK, "message.tmpl", gin.H{
		"title":       "Upload Successfully",
		"message":     "Upload Successfully",
		"description": "",
		"color":       "green",
	})
}
