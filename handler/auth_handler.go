package handler

import (
	"fmt"
	"net/http"
	"thunder_hoster/config"
	"thunder_hoster/public"
	"time"

	"github.com/gin-gonic/gin"
)

func PasswordAuthenticator(ctx *gin.Context) {
	ip := ctx.ClientIP()
	passwd := ctx.PostForm("password")

	if passwd == config.Cfg.Security.Password {
		public.ValidTime = time.Now().Add(time.Duration(config.Cfg.Service.ValidMin) * time.Minute)
		public.FailedCounter.Delete(ip)

		hostaddress := ctx.Request.Host
		if config.Cfg.Customize.HostAddress != "" {
			hostaddress = config.Cfg.Customize.HostAddress
		}
		ctx.HTML(http.StatusOK, "message.tmpl", gin.H{
			"title":       "Verification Successed",
			"message":     "Verification Successed",
			"description": fmt.Sprintf("Now you can visit %s/map , before %s", hostaddress, public.ValidTime.Format("2006-01-02 15:04:05")),
			"color":       "green",
		})
	} else {
		public.FailedCounter.Add(ip)
		ctx.HTML(http.StatusOK, "message.tmpl", gin.H{
			"title":       "Verification Failed",
			"message":     "Verification Failed",
			"description": "Please try again",
			"color":       "red",
		})
	}

}
