package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thunder_hoster/storage"
)

func MapList(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "maplist.tmpl", gin.H{
		"title":   "Map list",
		"maplist": storage.Storage.Maps,
	})
}
