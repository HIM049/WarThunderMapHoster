package pages

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PageMain(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")
	ctx.HTML(http.StatusOK, "index", nil)

}
