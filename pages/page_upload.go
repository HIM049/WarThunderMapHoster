package pages

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "upload.tmpl", gin.H{
		"title": "File Uploader",
	})
}
