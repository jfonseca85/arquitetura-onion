package cloudcontrolservice

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
)

// Atualiza os valores de propriedade especificados no recurso. Você especifica o seu recurso
// atualizações de propriedades como uma lista de operações de patch contidas em um patch JSON
// documento que segue o patch RFC 6902 - JavaScript Object Notation (JSON)
// (https://datatracker.ietf.org/doc/html/rfc6902) padrão. Para obter detalhes sobre como
// Cloud Control API executa operações de atualização de recursos, consulte Atualizar um recurso
// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-update.html)
// no Guia do usuário da API Amazon Web Services Cloud Control. Depois que você tiver
// iniciou uma solicitação de atualização de recurso, você pode monitorar o progresso de seu
// solicitação chamando GetResourceRequestStatus
// (https://docs.aws.amazon.com/cloudcontrolapi/latest/APIReference/API_GetResourceRequestStatus.html)
// usando o RequestToken do ProgressEvent retornado por UpdateResource. Para mais
// informações sobre as propriedades de um recurso específico, consulte o relacionado
// tópico para o recurso na referência de tipos de recurso e propriedade
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
// no Guia do usuário do CloudFormation da Amazon Web Services.
func (srv *Service) Update(params *cloudcontrol.UpdateResourceInput) (*cloudcontrol.UpdateResourceOutput, error) {
	//Metodo que destinado para Deletar os recursos usando o cloudcontrol
	log.Printf("Chamando o cloudControlCreateResourcesrv.Update>>> ")
	output, err := srv.controlsdk.UpdateResource(context.TODO(), params)
	if err != nil {
		return nil, err
	}
	return output, nil
}
