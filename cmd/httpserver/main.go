package main

import (
	"github.com/jfonseca85/controlplaneagent/internal/core/service/cloudcontrolsrv"
)

func main() {

	cloudcontrolsrv.New()

	//gamesRepository := gamesrepo.NewMemKVS()
	//gamesService := gamesrv.New(gamesRepository, uidgen.New())
	//gamesHandler := gamehdl.NewHTTPHandler(gamesService)

	//router := gin.New()
	//router.POST("/resources", gamesHandler.Create)

	//router.Run(":8080")
}
