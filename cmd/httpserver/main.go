//Client Http Server consumindo o port - cloudcontrol, adapter - http
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jfonseca85/arquitetura-onion/internal/core/services/cloudcontrolservice"
	"github.com/jfonseca85/arquitetura-onion/internal/handlers/cloudcontrolhdl"
	"github.com/jfonseca85/arquitetura-onion/internal/sdk"
)

func main() {
	//criação do client do cloudcontrol
	client := sdk.NewCloudControlClient()
	//Consumindo a implementacao do port - cloudcontrol
	controlService := cloudcontrolservice.New(client)
	//Realizando o acoplamento entre o Adpater http com o port http
	controlHandler := cloudcontrolhdl.NewHTTPHandler(controlService)
	//Subindo o server do http
	router := gin.New()
	//criando a rota do resources com o metodo POST
	router.POST("/resources", controlHandler.Create)
	//Externalizando na port 8080
	router.Run(":8080")
}
