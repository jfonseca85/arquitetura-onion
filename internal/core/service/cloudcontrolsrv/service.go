package cloudcontrolsrv

import (
	"context"
	"log"

	"github.com/jfonseca85/controlplaneagent/internal/core/domain/cloudcontrolmdl"
	"github.com/jfonseca85/controlplaneagent/internal/core/ports/cloudcontrolapi"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
)

type service struct {
	cloudControlSDK cloudcontrolapi.SDK
}

func New() *service {
	return &service{}
}

func (srv *service) Create(model cloudcontrolmdl.Model) (*cloudcontrolmdl.ProgressEvent, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("sa-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	client := cloudcontrol.NewFromConfig(cfg)
	output, err := client.CreateResource(context.TODO(), cloudcontrolmdl.ToResourceInput(&model))
	//progressEvent, err := srv.cloudControlSDK.Save(instanceAWS.ctx, model)
	if err != nil {
		return nil, err
	}

	return output, err
}
