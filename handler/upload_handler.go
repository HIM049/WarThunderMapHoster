package handler

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"thunder_hoster/config"
	"thunder_hoster/services"
	"thunder_hoster/storage"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadHandler(ctx *gin.Context) {
	group, exists := ctx.Get("group")
	if !exists {
		// 没有找到用户权限组标记
		fmt.Println("upload redirect1")
		ctx.Redirect(http.StatusFound, "/login")
		return
	}

	if group != services.GROUP_ADMIN {
		// 不是管理员用户
		fmt.Println("upload redirect2")
		ctx.Redirect(http.StatusFound, "/pages/list")
		return
	}
	mapName := ctx.PostForm("name")

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.Redirect(http.StatusFound, fmt.Sprintf("/pages/list?admin=1&error=Failed+to+upload:+%v", err))
		return
	}

	mapPath := filepath.Join(config.Cfg.Service.MapDir, mapName+".blk")
	err = ctx.SaveUploadedFile(file, mapPath)
	if err != nil {
		ctx.Redirect(http.StatusFound, fmt.Sprintf("/pages/list?admin=1&error=Failed+to+upload:+%v", err))
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
			ctx.Redirect(http.StatusFound, "/pages/list?admin=1&error=File+name+already+exists")
		} else {
			ctx.Redirect(http.StatusFound, fmt.Sprintf("/pages/list?admin=1&error=Failed+to+upload:+%v", err))
		}
		return
	}

	storage.Storage.GenerateIndex()

	ctx.Redirect(http.StatusFound, "/pages/list?admin=1&success=Upload+successfully")

}
