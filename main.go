package main

import (
	"fmt"
	"thunder_hoster/config"
	"thunder_hoster/public"
)

func main() {
	config.InitConfig()
	public.InitFailedCounter()

	router := RouterSetup()
	router.Run(fmt.Sprintf(":%d", config.Cfg.Port))
}
