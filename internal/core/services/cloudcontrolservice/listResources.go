package cloudcontrolservice

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
)

// Retorna informações sobre os recursos especificados. Para mais informações, veja
// Descobrindo recursos no Guia do usuário da API Amazon Web Services Cloud Control.
// Você pode usar esta ação para retornar informações sobre os recursos existentes na sua
// conta e região de serviços da Web da Amazon, sejam esses recursos ou não
// provisionado usando a API Cloud Control.
func (srv *Service) List(params *cloudcontrol.ListResourcesInput) (*cloudcontrol.ListResourcesOutput, error) {
	//Metodo que destinado para Deletar os recursos usando o cloudcontrol
	log.Printf("Chamando o cloudControlCreateResourcesrv.List>>> ")
	output, err := srv.controlsdk.ListResources(context.TODO(), params)
	if err != nil {
		return nil, err
	}
	return output, nil
}
