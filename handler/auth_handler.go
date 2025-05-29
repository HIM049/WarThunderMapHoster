package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thunder_hoster/config"
	"thunder_hoster/public"
	"thunder_hoster/services"
	"time"
)

func AuthHandler(ctx *gin.Context) {
	ip := ctx.ClientIP()
	passwd := ctx.PostForm("password")

	var userGroup string
	if passwd == config.Cfg.Security.Password {
		// 用户密码验证成功
		userGroup = services.GROUP_USER
	} else if passwd == config.Cfg.Security.AdminPasswd {
		// 管理员密码验证成功
		userGroup = services.GROUP_ADMIN
	} else {
		// 验证失败
		public.FailedCounter.Add(ip)
		ctx.Redirect(http.StatusFound, "/login?error=1")
		return
	}

	public.FailedCounter.Delete(ip)
	public.ValidTime = time.Now().Add(time.Duration(config.Cfg.Service.ValidMin) * time.Minute)

	jwtToken, err := services.GenerateJWT(userGroup)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/login?error=1") // TODO: Detailed error msg
		return
	}

	ctx.SetCookie(
		services.JWT_COOKIE_NAME,
		jwtToken,
		config.Cfg.ValidMin*60,
		"/",
		"",
		config.Cfg.NetWork.Https,
		true,
	)

	if userGroup == services.GROUP_USER {
		ctx.Redirect(http.StatusFound, "/pages/list")
	} else {
		ctx.Redirect(http.StatusFound, "/pages/list?admin=1")
	}
}
