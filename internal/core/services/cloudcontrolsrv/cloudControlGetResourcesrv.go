//Obtem o resource usando o aws sdk go v2 - service cloudcontrol
package cloudcontrolsrv

import (
	"context"
	"log"

	"github.com/jfonseca85/controlplaneagent/internal/core/domain"

	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
)

func (srv *Service) Get(model domain.CloudControlModel) (*cloudcontrol.GetResourceOutput, error) {
	//Metodo que destinado para Deletar os recursos usando o cloudcontrol
	log.Printf("Chamando o cloudControlCreateResourcesrv.Get>>> ")
	output, err := srv.controlsdk.GetResource(context.TODO(), toGetResourceInput(model))
	//getResourceOutput := toGetResourceOutput(output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func toGetResourceInput(model domain.CloudControlModel) *cloudcontrol.GetResourceInput {
	//Realiza a conversao do domain.CloudControlModel para cloudcontrol.DeleteResourceInput
	log.Printf("Chamando o cloudControlCreateResourcesrv.toGetResourceInput>>> ")
	result := cloudcontrol.GetResourceInput{
		TypeName:   &(model.TypeName),
		RoleArn:    &(model.RoleArn),
		Identifier: &(model.Identifier),
	}
	return &result
}
