package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"thunder_hoster/storage"
)

func SendFile(ctx *gin.Context) {
	mapName := ctx.Param("map")
	fmt.Println(mapName)
	mapInf, found := storage.Storage.ListMap[mapName]
	if !found {
		ctx.Redirect(http.StatusFound, "/pages/list?error=File+not+found")
		return
	}

	file, err := os.Open(mapInf.FilePath)
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
	ctx.File(mapInf.FilePath)
}
