package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thunder_hoster/services"
)

func LoginCheckMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie(services.JWT_COOKIE_NAME)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/login")
			return
		}

		claims, err := services.VerifyJWT(token)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/login")
			return
		}

		group, find := claims["group"]
		if !find {
			ctx.Redirect(http.StatusFound, "/login")
			return
		}

		ctx.Set("group", group)
		ctx.Next()
	}
}
