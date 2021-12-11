//	Centralizamos a criação das interfaces ports do cloudcontrol
//	@author: Jorge Luis
//	@version: 0.0.1
//	@Documentation: https://github.com/aws/aws-sdk-go-v2/tree/main/service/cloudcontrol
package ports

import (
	"github.com/jfonseca85/controlplaneagent/internal/core/domain"

	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
)

type CloudControlService interface {
	//Cria o recurso especificado.
	Create(domain.CloudControlModel) (*domain.ProgressEvent, error)
	//Exclui o recurso especificado.
	Delete(domain.CloudControlModel) (*domain.ProgressEvent, error)
	//Retorna informações sobre o estado atual do recurso especificado.
	Get(domain.CloudControlModel) (*cloudcontrol.GetResourceOutput, error)
	// Returns information about the specified resources. For more information, see
	// Discovering resources in the Amazon Web Services Cloud Control API User Guide.
	// You can use this action to return information about existing resources in your
	// account and Amazon Web Services Region, whether or not those resources were
	// provisioned using Cloud Control API.
	List(params *cloudcontrol.ListResourcesInput) (*cloudcontrol.ListResourcesOutput, error)

	//Retorna o status atual de uma solicitação de operação de recurso.
	GetStatus(params *cloudcontrol.GetResourceRequestStatusInput) (*cloudcontrol.GetResourceRequestStatusOutput, error)
	// Returns existing resource operation requests. This includes requests of all
	// status types. For more information, see Listing active resource operation
	// requests
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-manage-requests.html#resource-operations-manage-requests-list)
	// in the Amazon Web Services Cloud Control API User Guide. Resource operation
	// requests expire after seven days.
	ListRequests(params *cloudcontrol.ListResourceRequestsInput) (*cloudcontrol.ListResourceRequestsOutput, error)
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
	Update(params *cloudcontrol.UpdateResourceInput) (*cloudcontrol.UpdateResourceOutput, error)
	// Cancels the specified resource operation request. For more information, see
	// Canceling resource operation requests
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-manage-requests.html#resource-operations-manage-requests-cancel)
	// in the Amazon Web Services Cloud Control API User Guide. Only resource
	// operations requests with a status of PENDING or IN_PROGRESS can be cancelled.
	Cancel(params *cloudcontrol.CancelResourceRequestInput) (*cloudcontrol.CancelResourceRequestOutput, error)
}
