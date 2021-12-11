package cloudcontrolservice

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
)

// Retorna o status atual de uma solicitação de operação de recurso. Para mais
// informações, consulte Rastreando o progresso das solicitações de operação de recursos
// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-manage-requests.html#resource-operations-manage-requests-track)
// no Guia do usuário da API Amazon Web Services Cloud Control.
func (srv *Service) GetStatus(params *cloudcontrol.GetResourceRequestStatusInput) (*cloudcontrol.GetResourceRequestStatusOutput, error) {
	//Metodo que destinado para Deletar os recursos usando o cloudcontrol
	log.Printf("Chamando o cloudControlCreateResourcesrv.GetStatus>>> ")
	output, err := srv.controlsdk.GetResourceRequestStatus(context.TODO(), params)
	if err != nil {
		return nil, err
	}
	return output, nil
}
