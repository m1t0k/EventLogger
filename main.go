package main

import (
	"log"

	AppConfig "../eventLogger/v1/config/"
	Controllers "../eventLogger/v1/controllers/"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	err := AppConfig.ReadConfig()
	if err != nil {
		log.Fatalf("Can't start http server:%s.\n", err)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	v1 := router.Group("/v1")
	evCtrl := Controllers.EventController{}
	evCtrl.Register(v1)

	err = router.Run(":8999")
	if err != nil {
		log.Fatalf("Can't start http server:%s.\n", err)
	}
}
