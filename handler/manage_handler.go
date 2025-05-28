package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thunder_hoster/config"
	"thunder_hoster/storage"
)

func RemoveMap(ctx *gin.Context) {
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

	err := storage.Storage.Remove(mapName)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "message.tmpl", gin.H{
			"title":       "Remove Error",
			"message":     "Remove Error",
			"description": err.Error(),
			"color":       "red",
		})
	}

	ctx.HTML(http.StatusOK, "message.tmpl", gin.H{
		"title":       "Upload Successfully",
		"message":     "Upload Successfully",
		"description": "",
		"color":       "green",
	})
}
