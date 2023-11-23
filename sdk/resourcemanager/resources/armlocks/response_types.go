//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlocks

// AuthorizationOperationsClientListResponse contains the response from method AuthorizationOperationsClient.NewListPager.
type AuthorizationOperationsClientListResponse struct {
	// Result of the request to list Microsoft.Authorization operations. It contains a list of operations and a URL link to get
	// the next set of results.
	OperationListResult
}

// ManagementLocksClientCreateOrUpdateAtResourceGroupLevelResponse contains the response from method ManagementLocksClient.CreateOrUpdateAtResourceGroupLevel.
type ManagementLocksClientCreateOrUpdateAtResourceGroupLevelResponse struct {
	// The lock information.
	ManagementLockObject
}

// ManagementLocksClientCreateOrUpdateAtResourceLevelResponse contains the response from method ManagementLocksClient.CreateOrUpdateAtResourceLevel.
type ManagementLocksClientCreateOrUpdateAtResourceLevelResponse struct {
	// The lock information.
	ManagementLockObject
}

// ManagementLocksClientCreateOrUpdateAtSubscriptionLevelResponse contains the response from method ManagementLocksClient.CreateOrUpdateAtSubscriptionLevel.
type ManagementLocksClientCreateOrUpdateAtSubscriptionLevelResponse struct {
	// The lock information.
	ManagementLockObject
}

// ManagementLocksClientCreateOrUpdateByScopeResponse contains the response from method ManagementLocksClient.CreateOrUpdateByScope.
type ManagementLocksClientCreateOrUpdateByScopeResponse struct {
	// The lock information.
	ManagementLockObject
}

// ManagementLocksClientDeleteAtResourceGroupLevelResponse contains the response from method ManagementLocksClient.DeleteAtResourceGroupLevel.
type ManagementLocksClientDeleteAtResourceGroupLevelResponse struct {
	// placeholder for future response values
}

// ManagementLocksClientDeleteAtResourceLevelResponse contains the response from method ManagementLocksClient.DeleteAtResourceLevel.
type ManagementLocksClientDeleteAtResourceLevelResponse struct {
	// placeholder for future response values
}

// ManagementLocksClientDeleteAtSubscriptionLevelResponse contains the response from method ManagementLocksClient.DeleteAtSubscriptionLevel.
type ManagementLocksClientDeleteAtSubscriptionLevelResponse struct {
	// placeholder for future response values
}

// ManagementLocksClientDeleteByScopeResponse contains the response from method ManagementLocksClient.DeleteByScope.
type ManagementLocksClientDeleteByScopeResponse struct {
	// placeholder for future response values
}

// ManagementLocksClientGetAtResourceGroupLevelResponse contains the response from method ManagementLocksClient.GetAtResourceGroupLevel.
type ManagementLocksClientGetAtResourceGroupLevelResponse struct {
	// The lock information.
	ManagementLockObject
}

// ManagementLocksClientGetAtResourceLevelResponse contains the response from method ManagementLocksClient.GetAtResourceLevel.
type ManagementLocksClientGetAtResourceLevelResponse struct {
	// The lock information.
	ManagementLockObject
}

// ManagementLocksClientGetAtSubscriptionLevelResponse contains the response from method ManagementLocksClient.GetAtSubscriptionLevel.
type ManagementLocksClientGetAtSubscriptionLevelResponse struct {
	// The lock information.
	ManagementLockObject
}

// ManagementLocksClientGetByScopeResponse contains the response from method ManagementLocksClient.GetByScope.
type ManagementLocksClientGetByScopeResponse struct {
	// The lock information.
	ManagementLockObject
}

// ManagementLocksClientListAtResourceGroupLevelResponse contains the response from method ManagementLocksClient.NewListAtResourceGroupLevelPager.
type ManagementLocksClientListAtResourceGroupLevelResponse struct {
	// The list of locks.
	ManagementLockListResult
}

// ManagementLocksClientListAtResourceLevelResponse contains the response from method ManagementLocksClient.NewListAtResourceLevelPager.
type ManagementLocksClientListAtResourceLevelResponse struct {
	// The list of locks.
	ManagementLockListResult
}

// ManagementLocksClientListAtSubscriptionLevelResponse contains the response from method ManagementLocksClient.NewListAtSubscriptionLevelPager.
type ManagementLocksClientListAtSubscriptionLevelResponse struct {
	// The list of locks.
	ManagementLockListResult
}

// ManagementLocksClientListByScopeResponse contains the response from method ManagementLocksClient.NewListByScopePager.
type ManagementLocksClientListByScopeResponse struct {
	// The list of locks.
	ManagementLockListResult
}
