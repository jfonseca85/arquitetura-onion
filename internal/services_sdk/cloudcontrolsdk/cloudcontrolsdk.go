package cloudcontrolsdk

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
)

func NewClient() *cloudcontrol.Client {
	//Carregando  a configuração do aws
	log.Println("Carregando o contexto da AWS>>>>>>")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("sa-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	service := cloudcontrol.NewFromConfig(cfg)
	return service
}
