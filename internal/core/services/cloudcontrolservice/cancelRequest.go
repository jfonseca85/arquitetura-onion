package cloudcontrolservice

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
)

// Cancela a solicitação de operação do recurso especificado. Para mais informações, veja
// Cancelando solicitações de operação de recursos
// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-manage-requests.html#resource-operations-manage-requests-cancel)
// no Guia do usuário da API Amazon Web Services Cloud Control. Apenas recurso
// solicitações de operações com status PENDING ou IN_PROGRESS podem ser canceladas.
func (srv *Service) Cancel(params *cloudcontrol.CancelResourceRequestInput) (*cloudcontrol.CancelResourceRequestOutput, error) {
	//Metodo que destinado para Deletar os recursos usando o cloudcontrol
	log.Printf("Chamando o cloudControlCreateResourcesrv.Cancel>>> ")
	output, err := srv.controlsdk.CancelResourceRequest(context.TODO(), params)
	if err != nil {
		return nil, err
	}
	return output, nil
}
