package pages

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thunder_hoster/config"
)

func PageMain(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": config.Cfg.Customize.SideName,
	})

}
