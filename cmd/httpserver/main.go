package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jfonseca85/controlplaneagent/internal/core/service/cloudcontrolsrv"
	"github.com/jfonseca85/controlplaneagent/internal/handlers/cloudcontrolhdl"
	"github.com/jfonseca85/controlplaneagent/internal/services_sdk/cloudcontrolsdk"
)

func main() {

	//criação do client
	controlsdk := cloudcontrolsdk.NewClient()
	controlService := cloudcontrolsrv.New(controlsdk)
	controlHandler := cloudcontrolhdl.NewHTTPHandler(controlService)
	//controlHandler.Create()
	//dto := cloudcontrolHandler.BodyCreate{}
	//cloudcontrolHandler.Create()

	//gamesRepository := gamesrepo.NewMemKVS()
	//gamesService := gamesrv.New(gamesRepository, uidgen.New())
	//gamesHandler := gamehdl.NewHTTPHandler(gamesService)

	router := gin.New()
	router.POST("/resources", controlHandler.Create)

	router.Run(":8080")
}
