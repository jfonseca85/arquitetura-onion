package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jfonseca85/controlplaneagent/internal/core/service/cloudcontrolsrv"
	"github.com/jfonseca85/controlplaneagent/internal/handlers/cloudcontrolhdl"
	"github.com/jfonseca85/controlplaneagent/internal/sdk"
)

func main() {
	//criação do client
	client := sdk.NewCloudControlClient()
	controlService := cloudcontrolsrv.New(client)
	controlHandler := cloudcontrolhdl.NewHTTPHandler(controlService)
	router := gin.New()
	router.POST("/resources", controlHandler.Create)

	router.Run(":8080")
}
