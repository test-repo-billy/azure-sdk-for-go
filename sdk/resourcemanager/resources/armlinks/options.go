//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlinks

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// ResourceLinksClientCreateOrUpdateOptions contains the optional parameters for the ResourceLinksClient.CreateOrUpdate method.
type ResourceLinksClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ResourceLinksClientDeleteOptions contains the optional parameters for the ResourceLinksClient.Delete method.
type ResourceLinksClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ResourceLinksClientGetOptions contains the optional parameters for the ResourceLinksClient.Get method.
type ResourceLinksClientGetOptions struct {
	// placeholder for future optional parameters
}

// ResourceLinksClientListAtSourceScopeOptions contains the optional parameters for the ResourceLinksClient.NewListAtSourceScopePager
// method.
type ResourceLinksClientListAtSourceScopeOptions struct {
	// The filter to apply when getting resource links. To get links only at the specified scope (not below the scope), use Filter.atScope()..
	// Specifying any value will set the value to atScope().
	Filter *string
}

// ResourceLinksClientListAtSubscriptionOptions contains the optional parameters for the ResourceLinksClient.NewListAtSubscriptionPager
// method.
type ResourceLinksClientListAtSubscriptionOptions struct {
	// The filter to apply on the list resource links operation. The supported filter for list resource links is targetId. For
	// example, $filter=targetId eq {value}
	Filter *string
}
