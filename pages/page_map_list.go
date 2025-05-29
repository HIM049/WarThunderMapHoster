package pages

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thunder_hoster/config"
	"thunder_hoster/storage"
)

func PageMapList(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "list", gin.H{
		"maplist":       storage.Storage.Maps,
		"downloadRoute": config.Cfg.DownloadRouter,
	})
}
