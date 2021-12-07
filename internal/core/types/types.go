package types

// Represents information about a provisioned resource.
type ResourceDescription struct {

	// The primary identifier for the resource. For more information, see Identifying
	// resources
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-identifier.html)
	// in the Amazon Web Services Cloud Control API User Guide.
	Identifier *string

	// A list of the resource properties and their current values.
	Properties *string
}

// The filter criteria to use in determining the requests returned.
type ResourceRequestStatusFilter struct {

	// The operation statuses to include in the filter.
	//
	// * PENDING: The operation has
	// been requested, but not yet initiated.
	//
	// * IN_PROGRESS: The operation is
	// currently in progress.
	//
	// * SUCCESS: The operation has successfully completed.
	//
	// *
	// FAILED: The operation has failed.
	//
	// * CANCEL_IN_PROGRESS: The operation is
	// currently in the process of being canceled.
	//
	// * CANCEL_COMPLETE: The operation
	// has been canceled.
	OperationStatuses []OperationStatus

	// The operation types to include in the filter.
	Operations []Operation
}
