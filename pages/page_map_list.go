package pages

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thunder_hoster/storage"
)

func PageMapList(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "maplist.tmpl", gin.H{
		"title":   "Map list",
		"maplist": storage.Storage.Maps,
	})
}
