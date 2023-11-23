//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtemplatespecs

// ClientCreateOrUpdateResponse contains the response from method Client.CreateOrUpdate.
type ClientCreateOrUpdateResponse struct {
	// Template Spec object.
	TemplateSpec
}

// ClientDeleteResponse contains the response from method Client.Delete.
type ClientDeleteResponse struct {
	// placeholder for future response values
}

// ClientGetBuiltInResponse contains the response from method Client.GetBuiltIn.
type ClientGetBuiltInResponse struct {
	// Template Spec object.
	TemplateSpec
}

// ClientGetResponse contains the response from method Client.Get.
type ClientGetResponse struct {
	// Template Spec object.
	TemplateSpec
}

// ClientListBuiltInsResponse contains the response from method Client.NewListBuiltInsPager.
type ClientListBuiltInsResponse struct {
	// List of Template Specs.
	ListResult
}

// ClientListByResourceGroupResponse contains the response from method Client.NewListByResourceGroupPager.
type ClientListByResourceGroupResponse struct {
	// List of Template Specs.
	ListResult
}

// ClientListBySubscriptionResponse contains the response from method Client.NewListBySubscriptionPager.
type ClientListBySubscriptionResponse struct {
	// List of Template Specs.
	ListResult
}

// ClientUpdateResponse contains the response from method Client.Update.
type ClientUpdateResponse struct {
	// Template Spec object.
	TemplateSpec
}

// TemplateSpecVersionsClientCreateOrUpdateResponse contains the response from method TemplateSpecVersionsClient.CreateOrUpdate.
type TemplateSpecVersionsClientCreateOrUpdateResponse struct {
	// Template Spec Version object.
	TemplateSpecVersion
}

// TemplateSpecVersionsClientDeleteResponse contains the response from method TemplateSpecVersionsClient.Delete.
type TemplateSpecVersionsClientDeleteResponse struct {
	// placeholder for future response values
}

// TemplateSpecVersionsClientGetBuiltInResponse contains the response from method TemplateSpecVersionsClient.GetBuiltIn.
type TemplateSpecVersionsClientGetBuiltInResponse struct {
	// Template Spec Version object.
	TemplateSpecVersion
}

// TemplateSpecVersionsClientGetResponse contains the response from method TemplateSpecVersionsClient.Get.
type TemplateSpecVersionsClientGetResponse struct {
	// Template Spec Version object.
	TemplateSpecVersion
}

// TemplateSpecVersionsClientListBuiltInsResponse contains the response from method TemplateSpecVersionsClient.NewListBuiltInsPager.
type TemplateSpecVersionsClientListBuiltInsResponse struct {
	// List of Template Specs versions
	TemplateSpecVersionsListResult
}

// TemplateSpecVersionsClientListResponse contains the response from method TemplateSpecVersionsClient.NewListPager.
type TemplateSpecVersionsClientListResponse struct {
	// List of Template Specs versions
	TemplateSpecVersionsListResult
}

// TemplateSpecVersionsClientUpdateResponse contains the response from method TemplateSpecVersionsClient.Update.
type TemplateSpecVersionsClientUpdateResponse struct {
	// Template Spec Version object.
	TemplateSpecVersion
}
