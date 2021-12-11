package cloudcontrolservice

import (
	"context"
	"log"

	"github.com/jfonseca85/arquitetura-onion/internal/core/domain"

	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
)

// Retorna informações sobre o estado atual do recurso especificado. Para mais
// detalhes, consulte a documentação da Amazon
// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-read.html).
// Você pode usar esta ação para retornar informações sobre um recurso existente em sua
// conta e região de serviços da Web da Amazon, sejam esses recursos ou não
// provisionado usando a API Cloud Control.
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
