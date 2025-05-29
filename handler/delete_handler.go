package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thunder_hoster/services"
	"thunder_hoster/storage"
)

func DeleteHandler(ctx *gin.Context) {
	group, exists := ctx.Get("group")
	if !exists {
		// 没有找到用户权限组标记
		ctx.Redirect(http.StatusFound, "/login")
		return
	}

	if group != services.GROUP_ADMIN {
		// 不是管理员用户
		ctx.Redirect(http.StatusFound, "/pages/list")
		return
	}
	mapName := ctx.PostForm("name")

	err := storage.Storage.Remove(mapName)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "message.tmpl", gin.H{
			"title":       "Delete Failed",
			"message":     "Failed to save file",
			"description": "Error: " + err.Error(),
			"color":       "red",
		})
		return
	}

	storage.Storage.GenerateIndex()

	ctx.HTML(http.StatusOK, "message.tmpl", gin.H{
		"title":       "Delete Successfully",
		"message":     "Delete Successfully",
		"description": "",
		"color":       "green",
	})
}
