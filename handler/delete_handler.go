package handler

import (
	"fmt"
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
		ctx.Redirect(http.StatusFound, fmt.Sprintf("/pages/list?admin=1&error=Failed+to+save+change:+%v", err))
		return
	}

	storage.Storage.GenerateIndex()

	ctx.Redirect(http.StatusFound, "/pages/list?admin=1&success=Map+was+deleted")
}
