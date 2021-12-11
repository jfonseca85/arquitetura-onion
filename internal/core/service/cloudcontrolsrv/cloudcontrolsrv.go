//Pacote de criação das regras de negocio do AWS SDK GO V2 service cloudcontrol
package cloudcontrolsrv

import (
	"context"
	"log"

	"github.com/jfonseca85/controlplaneagent/internal/core/domain"

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

func (srv *Service) Create(model domain.CloudControlModel) (*domain.ProgressEvent, error) {
	//Metodo que destinado para criar os recursos usando o cloudcontrol
	log.Printf("Chamando o CreateResource>>> ")
	output, err := srv.controlsdk.CreateResource(context.TODO(), domain.ToResourceInput(model))
	progressEvent := domain.ToProgressEvent(output)
	if err != nil {
		return nil, err
	}
	return progressEvent, nil
}
