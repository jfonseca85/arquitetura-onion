/*
	Pacote que centraliza os model do cloudcontrol
	@author: Jorge Luis
	@version: 0.0.1
	@Documentation: https://github.com/aws/aws-sdk-go-v2/tree/main/service/cloudcontrol
*/
package domain

import (
	"time"

	"github.com/jfonseca85/controlplaneagent/internal/types"

	"github.com/aws/aws-sdk-go-v2/services/cloudcontrol"
)

type CloudControlModel struct {
	// Structured data format representing the desired state of the resource,
	// consisting of that resource's properties and their desired values. Cloud Control
	// API currently supports JSON as a structured data format. Specify the desired
	// state as one of the following:
	//
	// * A JSON blob
	//
	// * A local path containing the
	// desired state in JSON data format
	//
	// For more information, see Composing the
	// desired state of the resource
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-create.html#resource-operations-create-desiredstate)
	// in the Amazon Web Services Cloud Control API User Guide. For more information
	// about the properties of a specific resource, refer to the related topic for the
	// resource in the Resource and property types reference
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
	// in the Amazon Web Services CloudFormation Users Guide.
	//
	// This member is required.
	DesiredState string

	// The name of the resource type.
	//
	// This member is required.
	TypeName string

	// A unique identifier to ensure the idempotency of the resource request. As a best
	// practice, specify this token to ensure idempotency, so that Amazon Web Services
	// Cloud Control API can accurately distinguish between request retries and new
	// resource requests. You might retry a resource request to ensure that it was
	// successfully received. A client token is valid for 36 hours once used. After
	// that, a resource request with the same client token is treated as a new request.
	// If you do not specify a client token, one is generated for inclusion in the
	// request. For more information, see Ensuring resource operation requests are
	// unique
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations.html#resource-operations-idempotency)
	// in the Amazon Web Services Cloud Control API User Guide.
	ClientToken string

	// The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) for
	// Cloud Control API to use when performing this resource operation. The role
	// specified must have the permissions required for this operation. The necessary
	// permissions for each event handler are defined in the handlers
	// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html#schema-properties-handlers)
	// section of the resource type definition schema
	// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html).
	// If you do not specify a role, Cloud Control API uses a temporary session created
	// using your Amazon Web Services user credentials. For more information, see
	// Specifying credentials
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations.html#resource-operations-permissions)
	// in the Amazon Web Services Cloud Control API User Guide.
	RoleArn string

	// For private resource types, the type version to use in this resource operation.
	// If you do not specify a resource version, CloudFormation uses the default
	// version.
	TypeVersionId string
}

func ToResourceInput(model CloudControlModel) *cloudcontrol.CreateResourceInput {
	result := cloudcontrol.CreateResourceInput{
		TypeName:      &(model.TypeName),
		DesiredState:  &(model.DesiredState),
		ClientToken:   &(model.ClientToken),
		RoleArn:       &(model.RoleArn),
		TypeVersionId: &(model.TypeVersionId),
	}
	return &result
}

func ToProgressEvent(model *cloudcontrol.CreateResourceOutput) *ProgressEvent {
	result := ProgressEvent{
		ErrorCode:       types.HandlerErrorCode(model.ProgressEvent.ErrorCode),
		EventTime:       model.ProgressEvent.EventTime,
		Identifier:      model.ProgressEvent.Identifier,
		Operation:       types.Operation(model.ProgressEvent.Operation),
		OperationStatus: types.OperationStatus(model.ProgressEvent.Operation),
		RequestToken:    model.ProgressEvent.RequestToken,
		ResourceModel:   model.ProgressEvent.ResourceModel,
		RetryAfter:      model.ProgressEvent.RetryAfter,
		StatusMessage:   model.ProgressEvent.StatusMessage,
		TypeName:        model.ProgressEvent.TypeName,
	}
	return &result
}

type ProgressEvent struct {
	// For requests with a status of FAILED, the associated error code. For error code
	// definitions, see Handler error codes
	// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-test-contract-errors.html)
	// in the CloudFormation Command Line Interface User Guide for Extension
	// Development.
	ErrorCode types.HandlerErrorCode
	// When the resource operation request was initiated.
	// When the resource operation request was initiated.
	EventTime *time.Time
	// The primary identifier for the resource. In some cases, the resource identifier
	// may be available before the resource operation has reached a status of SUCCESS.
	Identifier *string

	// The resource operation type.
	Operation types.Operation
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
	OperationStatus types.OperationStatus

	// The unique token representing this resource operation request. Use the
	// RequestToken with GetResourceRequestStatus
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/APIReference/API_GetResourceRequestStatus.html)
	// to return the current status of a resource operation request.
	RequestToken *string

	// A JSON string containing the resource Model, consisting of each resource
	// property and its current value.
	ResourceModel *string

	// When to next request the status of this resource operation request.
	RetryAfter *time.Time

	// Any message explaining the current status.
	StatusMessage *string

	// The name of the resource type used in the operation.
	TypeName *string
}

type RequestList struct {

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
	// The resource Model to use to select the resources to return.
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

type ListOutput struct {

	// If the request doesn't return all of the remaining results, NextToken is set to
	// a token. To retrieve the next set of results, call ListResources again and
	// assign that token to the request object's NextToken parameter. If the request
	// returns all results, NextToken is set to null.
	NextToken *string

	// Information about the specified resources, including primary identifier and
	// resource Model.
	ResourceDescriptions []types.ResourceDescription

	// The name of the resource type.
	TypeName *string
}

type ListRequestsInput struct {

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

	// The filter criteria to apply to the requests returned.
	ResourceRequestStatusFilter *types.ResourceRequestStatusFilter
}

type ListRequestsOutput struct {

	// If the request doesn't return all of the remaining results, NextToken is set to
	// a token. To retrieve the next set of results, call ListResources again and
	// assign that token to the request object's NextToken parameter. If the request
	// returns all results, NextToken is set to null.
	NextToken *string

	// The requests that match the specified filter criteria.
	ResourceRequestStatusSummaries []ProgressEvent
}

type UpdateInput struct {

	// The identifier for the resource. You can specify the primary identifier, or any
	// secondary identifier defined for the resource type in its resource schema. You
	// can only specify one identifier. Primary identifiers can be specified as a
	// string or JSON; secondary identifiers must be specified as JSON. For compound
	// primary identifiers (that is, one that consists of multiple resource properties
	// strung together), to specify the primary identifier as a string, list the
	// property values in the order they are specified in the primary identifier
	// definition, separated by |. For more information, see Identifying resources
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-identifier.html)
	// in the Amazon Web Services Cloud Control API User Guide.
	//
	// This member is required.
	Identifier *string

	// A JavaScript Object Notation (JSON) document listing the patch operations that
	// represent the updates to apply to the current resource properties. For details,
	// see Composing the patch document
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-update.html#resource-operations-update-patch)
	// in the Amazon Web Services Cloud Control API User Guide.
	//
	// This member is required.
	PatchDocument *string

	// The name of the resource type.
	//
	// This member is required.
	TypeName *string

	// A unique identifier to ensure the idempotency of the resource request. As a best
	// practice, specify this token to ensure idempotency, so that Amazon Web Services
	// Cloud Control API can accurately distinguish between request retries and new
	// resource requests. You might retry a resource request to ensure that it was
	// successfully received. A client token is valid for 36 hours once used. After
	// that, a resource request with the same client token is treated as a new request.
	// If you do not specify a client token, one is generated for inclusion in the
	// request. For more information, see Ensuring resource operation requests are
	// unique
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations.html#resource-operations-idempotency)
	// in the Amazon Web Services Cloud Control API User Guide.
	ClientToken *string

	// The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) for
	// Cloud Control API to use when performing this resource operation. The role
	// specified must have the permissions required for this operation. The necessary
	// permissions for each event handler are defined in the handlers
	// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html#schema-properties-handlers)
	// section of the resource type definition schema
	// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html).
	// If you do not specify a role, Cloud Control API uses a temporary session created
	// using your Amazon Web Services user credentials. For more information, see
	// Specifying credentials
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations.html#resource-operations-permissions)
	// in the Amazon Web Services Cloud Control API User Guide.
	RoleArn *string

	// For private resource types, the type version to use in this resource operation.
	// If you do not specify a resource version, CloudFormation uses the default
	// version.
	TypeVersionId *string
}

type UpdateOutput struct {

	// Represents the current status of the resource update request. Use the
	// RequestToken of the ProgressEvent with GetResourceRequestStatus
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/APIReference/API_GetResourceRequestStatus.html)
	// to return the current status of a resource operation request.
	ProgressEvent *ProgressEvent
}
