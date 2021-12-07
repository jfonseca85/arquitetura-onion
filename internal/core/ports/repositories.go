package ports

import "github.com/jfonseca85/controlplaneagent/internal/core/domain"

/**
	Centralizarmos a criação das interfaces ports do control agent
	@author: Jorge Luis
	@version: 0.0.1
	@Documentation: https://github.com/aws/aws-sdk-go-v2/tree/main/service/cloudcontrol
**/

/**
	Interface do port de conexão com o AWS SDK
**/
type CloudControlSDK interface {
	//Cria o recurso especificado.
	Save(domain.CloudControl) error
	//Exclui o recurso especificado.
	Delete(domain.CloudControl) error
	//Retorna informações sobre o estado atual do recurso especificado.
	Get(domain.CloudControl) (domain.CloudControl, error)
	//Retorna o status atual de uma solicitação de operação de recurso.
	GetRequestStatus(requestToken string) (domain.ProgressEvent, error)
	// Returns information about the specified resources. For more information, see
	// Discovering resources in the Amazon Web Services Cloud Control API User Guide.
	// You can use this action to return information about existing resources in your
	// account and Amazon Web Services Region, whether or not those resources were
	// provisioned using Cloud Control API.
	List(domain.CloudControlRequestList) (domain.ProgressEvent, error)
}
