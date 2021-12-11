package cloudcontrolservice

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
)

// Retorna as solicitações de operação de recursos existentes. Isso inclui solicitações de todos
// tipos de status. Para obter mais informações, consulte Listando operação de recurso ativo
// solicitações de
// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-manage-requests.html#resource-operations-manage-requests-list)
// no Guia do usuário da API Amazon Web Services Cloud Control. Operação de recursos
// os pedidos expiram após sete dias.
func (srv *Service) ListRequests(params *cloudcontrol.ListResourceRequestsInput) (*cloudcontrol.ListResourceRequestsOutput, error) {
	//Metodo que destinado para Deletar os recursos usando o cloudcontrol
	log.Printf("Chamando o cloudControlCreateResourcesrv.List>>> ")
	output, err := srv.controlsdk.ListResourceRequests(context.TODO(), params)
	if err != nil {
		return nil, err
	}
	return output, nil
}
