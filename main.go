package main

import (
	"fmt"
	"log"
	"th-release/dcm/api"
	"th-release/dcm/utils"
)

func main() {
	config := utils.GetConfig()

	if config == nil {
		log.Fatalln("Not Found Config")
		return
	}

	app := api.InitServer(config)

	if app == nil {
		log.Fatalln("Init Server Error")
		return
	}

	app.App.Listen(fmt.Sprintf(":%d", config.Port))
}
