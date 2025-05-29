package main

import (
	"fmt"
	"thunder_hoster/config"
	"thunder_hoster/public"
	"thunder_hoster/services"
	"thunder_hoster/storage"
)

func main() {
	config.InitConfig()
	storage.InitStorage()
	services.InitKeys()
	public.InitFailedCounter()

	router := RouterSetup()
	router.Run(fmt.Sprintf(":%d", config.Cfg.Port))
}
