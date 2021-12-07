/**
	Centralizarmos a criação das interfaces ports do control agent
	@author: Jorge Luis
	@version: 0.0.1
	@Documentation: https://github.com/aws/aws-sdk-go-v2/tree/main/service/cloudcontrol
**/
package sdk

import (
	"github.com/jfonseca85/controlplaneagent/internal/core/domain/cloudcontrol"
	"github.com/jfonseca85/controlplaneagent/internal/core/types"
)

/**
	Interface do port de conexão com o AWS SDK
**/
type CloudControlSDK interface {
	//Cria o recurso especificado.
	Save(cloudcontrol.Model) (*cloudcontrol.ProgressEvent, error)
	//Exclui o recurso especificado.
	Delete(cloudcontrol.Model) (*cloudcontrol.ProgressEvent, error)
	//Retorna informações sobre o estado atual do recurso especificado.
	Get(cloudcontrol.Model) (types.ResourceDescription, error)
	// Returns information about the specified resources. For more information, see
	// Discovering resources in the Amazon Web Services Cloud Control API User Guide.
	// You can use this action to return information about existing resources in your
	// account and Amazon Web Services Region, whether or not those resources were
	// provisioned using Cloud Control API.
	List(cloudcontrol.RequestList) (cloudcontrol.ProgressEvent, error)

	//Retorna o status atual de uma solicitação de operação de recurso.
	GetRequestStatus(requestToken string) (cloudcontrol.ProgressEvent, error)
	// Returns existing resource operation requests. This includes requests of all
	// status types. For more information, see Listing active resource operation
	// requests
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-manage-requests.html#resource-operations-manage-requests-list)
	// in the Amazon Web Services Cloud Control API User Guide. Resource operation
	// requests expire after seven days.
	ListRequests(cloudcontrol.ListRequestsInput) (cloudcontrol.ListRequestsOutput, error)
	// Updates the specified property values in the resource. You specify your resource
	// property updates as a list of patch operations contained in a JSON patch
	// document that adheres to the  RFC 6902 - JavaScript Object Notation (JSON) Patch
	// (https://datatracker.ietf.org/doc/html/rfc6902) standard. For details on how
	// Cloud Control API performs resource update operations, see Updating a resource
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-update.html)
	// in the Amazon Web Services Cloud Control API User Guide. After you have
	// initiated a resource update request, you can monitor the progress of your
	// request by calling GetResourceRequestStatus
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/APIReference/API_GetResourceRequestStatus.html)
	// using the RequestToken of the ProgressEvent returned by UpdateResource. For more
	// information about the properties of a specific resource, refer to the related
	// topic for the resource in the Resource and property types reference
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
	// in the Amazon Web Services CloudFormation Users Guide.
	Update(cloudcontrol.UpdateInput) (cloudcontrol.ListRequestsOutput, error)
}
