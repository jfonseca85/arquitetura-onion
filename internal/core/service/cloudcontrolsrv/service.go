//Pacote de criação das regras de negocio do AWS SDK GO V2 service cloudcontrol
package cloudcontrolsrv

import (
	"context"
	"log"

	"github.com/jfonseca85/controlplaneagent/internal/core/domain/cloudcontrolmdl"

	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
)

type Service struct {
	controlsdk *cloudcontrol.Client
	//controlsdk cloudcontrolapi.SDK
}

func New(controlsdk *cloudcontrol.Client) Service {
	return Service{
		controlsdk: controlsdk,
	}
}

func (srv *Service) Create(model cloudcontrolmdl.Model) (*cloudcontrolmdl.ProgressEvent, error) {
	//Metodo que destinado para criar os recursos usando o cloudcontrol
	log.Printf("Chamando o CreateResource>>> ", cloudcontrolmdl.ToResourceInput(&model))
	output, err := srv.controlsdk.CreateResource(context.TODO(), cloudcontrolmdl.ToResourceInput(&model))
	progressEvent := cloudcontrolmdl.ToProgressEvent(output)
	if err != nil {
		return nil, err
	}
	return progressEvent, nil
}
