package domain

import (
	"time"
)

/**
	Domain API Cloud Control para operações de API
	@author: Jorge Luis
	@version: 0.0.1
	@Documentation: https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/what-is-cloudcontrolapi.html
**/
type CloudControl struct {
	//Um identificador exclusivo para garantir a idempotência da solicitação de recurso.
	ClientToken string `json:"clientToken"`

	//Formato de dados estruturados que representa o estado desejado do recurso
	DesiredState string `json:"desiredState"`

	//O Amazon Resource Name (ARN) da função AWS Identity and Access Management (IAM) para a API Cloud Control usar ao executar esta operação de recurso.
	RoleArn string `json:"roleArn"`

	//O nome do tipo de recurso.
	TypeName string `json:"typeName"`
	//Para tipos de recursos privados, a versão do tipo a ser usada nesta operação de recurso.
	TypeVersionId string `json:"typeVersionId"`

	//O identificador do recurso.
	Identifier string `json:"identifier"`
}

/**
	Representa o status atual da solicitação de operação do recurso.
	Documentation: https://docs.aws.amazon.com/cloudcontrolapi/latest/APIReference/API_GetResourceRequestStatus.html
**/
type ProgressEvent struct {
	// For requests with a status of FAILED, the associated error code. For error code
	// definitions, see Handler error codes
	// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-test-contract-errors.html)
	// in the CloudFormation Command Line Interface User Guide for Extension
	// Development.
	ErrorCode HandlerErrorCode
	// When the resource operation request was initiated.
	// When the resource operation request was initiated.
	EventTime *time.Time
	// The primary identifier for the resource. In some cases, the resource identifier
	// may be available before the resource operation has reached a status of SUCCESS.
	Identifier *string

	// The resource operation type.
	Operation Operation
	// The current status of the resource operation request.
	//
	// * PENDING: The resource
	// operation has not yet started.
	//
	// * IN_PROGRESS: The resource operation is
	// currently in progress.
	//
	// * SUCCESS: The resource operation has successfully
	// completed.
	//
	// * FAILED: The resource operation has failed. Refer to the error code
	// and status message for more information.
	//
	// * CANCEL_IN_PROGRESS: The resource
	// operation is in the process of being canceled.
	//
	// * CANCEL_COMPLETE: The resource
	// operation has been canceled.
	OperationStatus OperationStatus

	// The unique token representing this resource operation request. Use the
	// RequestToken with GetResourceRequestStatus
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/APIReference/API_GetResourceRequestStatus.html)
	// to return the current status of a resource operation request.
	RequestToken *string

	// A JSON string containing the resource model, consisting of each resource
	// property and its current value.
	ResourceModel *string

	// When to next request the status of this resource operation request.
	RetryAfter *time.Time

	// Any message explaining the current status.
	StatusMessage *string

	// The name of the resource type used in the operation.
	TypeName *string
}

/**
	Domain que envia informações para retornar os recursos existentes na conta e região AWS
**/
type CloudControlRequestList struct {

	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next set
	// of results. The default is 20.
	MaxResults *int32
	// If the previous paginated request didn't return all of the remaining results,
	// the response object's NextToken parameter value is set to a token. To retrieve
	// the next set of results, call this action again and assign that token to the
	// request object's NextToken parameter. If there are no remaining results, the
	// previous response object's NextToken parameter is set to null.
	NextToken *string
	// The resource model to use to select the resources to return.
	ResourceModel *string

	// For private resource types, the type version to use in this resource operation.
	// If you do not specify a resource version, CloudFormation uses the default
	// version
	RoleArn *string
	// The name of the resource type.
	//
	// This member is required.
	TypeName *string

	// For private resource types, the type version to use in this resource operation.
	// If you do not specify a resource version, CloudFormation uses the default
	// version
	TypeVersionId *string
}

type ListResourcesOutput struct {

	// If the request doesn't return all of the remaining results, NextToken is set to
	// a token. To retrieve the next set of results, call ListResources again and
	// assign that token to the request object's NextToken parameter. If the request
	// returns all results, NextToken is set to null.
	NextToken *string

	// Information about the specified resources, including primary identifier and
	// resource model.
	ResourceDescriptions []types.ResourceDescription

	// The name of the resource type.
	TypeName *string
}
