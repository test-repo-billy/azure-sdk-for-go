//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armchanges

// ChangeAttributes - Details about the change resource
type ChangeAttributes struct {
	// READ-ONLY; The number of changes this resource captures
	ChangesCount *int64

	// READ-ONLY; The ARM correlation ID of the change resource
	CorrelationID *string

	// READ-ONLY; The GUID of the new snapshot
	NewResourceSnapshotID *string

	// READ-ONLY; The GUID of the previous snapshot
	PreviousResourceSnapshotID *string

	// READ-ONLY; The time the change(s) on the target resource ocurred
	Timestamp *string
}

// ChangeBase - An individual change on the target resource
type ChangeBase struct {
	// READ-ONLY; The entity that made the change
	ChangeCategory *ChangeCategory

	// READ-ONLY; The target resource property value after the change
	NewValue *string

	// READ-ONLY; The target resource property value before the change
	PreviousValue *string

	// READ-ONLY; The type of change that occurred
	PropertyChangeType *PropertyChangeType
}

// ChangeProperties - The properties of a change
type ChangeProperties struct {
	// Details about the change resource
	ChangeAttributes *ChangeAttributes

	// A dictionary with changed property name as a key and the change details as the value
	Changes map[string]*ChangeBase

	// READ-ONLY; The type of change that was captured in the resource
	ChangeType *ChangeType

	// READ-ONLY; The fully qualified ID of the target resource that was changed
	TargetResourceID *string

	// READ-ONLY; The namespace and type of the resource
	TargetResourceType *string
}

// ChangeResourceListResult - The list of resources
type ChangeResourceListResult struct {
	// The link used to get the next page of Change Resources
	NextLink *string

	// The list of resources
	Value []*ChangeResourceResult
}

// ChangeResourceResult - Change Resource
type ChangeResourceResult struct {
	// The properties of a change
	Properties *ChangeProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any

	// READ-ONLY; The additional info type.
	Type *string
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*ErrorDetail

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}
