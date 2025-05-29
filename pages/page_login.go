package pages

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PageLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login", gin.H{})
}
