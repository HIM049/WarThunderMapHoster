package main

import (
	"fmt"
	"log"
	"os"
	"thunder_hoster/config"
	"thunder_hoster/public"
)

func main() {
	config.InitConfig()
	public.InitFailedCounter()

	err := os.MkdirAll("./maps", os.ModeDir)
	if err != nil {
		log.Fatalln(err)
	}

	router := RouterSetup()
	router.Run(fmt.Sprintf(":%d", config.Cfg.Port))
}
