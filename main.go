package main

import (
	AppConfig "../eventservice/v1/config/"
	Controllers "../eventservice/v1/controllers/"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
)

func main() {
	err := AppConfig.ReadConfig()
	if err != nil {
		log.Fatalf("Can't start http server:%s.\n", err)
	}

	router := gin.Default()
	evCtrl := Controllers.EventController{}
	evCtrl.Run(router)

	router.Run(":8999")
}