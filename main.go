package main

import (
	"fmt"
	"thunder_hoster/config"
	"thunder_hoster/public"
	"thunder_hoster/storage"
)

func main() {
	config.InitConfig()
	public.InitFailedCounter()
	storage.InitStorage()

	router := RouterSetup()
	router.Run(fmt.Sprintf(":%d", config.Cfg.Port))
}
