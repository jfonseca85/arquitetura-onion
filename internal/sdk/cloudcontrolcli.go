package sdk

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
)

func NewCloudControlClient() *cloudcontrol.Client {
	//Carregando  a configuração do aws
	log.Println("Carregando o contexto da AWS>>>>>>")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("sa-east-1"), config.WithClientLogMode(aws.LogRetries|aws.LogRequest))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	service := cloudcontrol.NewFromConfig(cfg)
	return service
}
