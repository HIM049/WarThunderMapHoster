package handler

import (
	"fmt"
	"os"
	"strconv"
	"thunder_hoster/config"

	"github.com/gin-gonic/gin"
)

func SendFile(ctx *gin.Context) {

	file, err := os.Open(config.Cfg.Service.FilePath)
	if err != nil {
		panic(err)
	}

	fileInfo, _ := file.Stat()

	ctx.Header("Content-Description", " File Transfer")
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileInfo.Name()))
	ctx.Header("Expires", "0")
	ctx.Header("Cache-Control", "must-revalidate")
	ctx.Header("Pragma", "public")
	ctx.Header("Content-Length", strconv.Itoa(int(fileInfo.Size())))
	ctx.File(config.Cfg.Service.FilePath)
}
