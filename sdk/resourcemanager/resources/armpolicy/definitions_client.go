//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpolicy

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// DefinitionsClient contains the methods for the PolicyDefinitions group.
// Don't use this type directly, use NewDefinitionsClient() instead.
type DefinitionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDefinitionsClient creates a new instance of DefinitionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDefinitionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DefinitionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DefinitionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - This operation creates or updates a policy definition in the given subscription with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - policyDefinitionName - The name of the policy definition to create.
//   - parameters - The policy definition properties.
//   - options - DefinitionsClientCreateOrUpdateOptions contains the optional parameters for the DefinitionsClient.CreateOrUpdate
//     method.
func (client *DefinitionsClient) CreateOrUpdate(ctx context.Context, policyDefinitionName string, parameters Definition, options *DefinitionsClientCreateOrUpdateOptions) (DefinitionsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "DefinitionsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, policyDefinitionName, parameters, options)
	if err != nil {
		return DefinitionsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DefinitionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return DefinitionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DefinitionsClient) createOrUpdateCreateRequest(ctx context.Context, policyDefinitionName string, parameters Definition, options *DefinitionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}"
	if policyDefinitionName == "" {
		return nil, errors.New("parameter policyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyDefinitionName}", url.PathEscape(policyDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DefinitionsClient) createOrUpdateHandleResponse(resp *http.Response) (DefinitionsClientCreateOrUpdateResponse, error) {
	result := DefinitionsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Definition); err != nil {
		return DefinitionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// CreateOrUpdateAtManagementGroup - This operation creates or updates a policy definition in the given management group with
// the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - policyDefinitionName - The name of the policy definition to create.
//   - managementGroupID - The ID of the management group.
//   - parameters - The policy definition properties.
//   - options - DefinitionsClientCreateOrUpdateAtManagementGroupOptions contains the optional parameters for the DefinitionsClient.CreateOrUpdateAtManagementGroup
//     method.
func (client *DefinitionsClient) CreateOrUpdateAtManagementGroup(ctx context.Context, policyDefinitionName string, managementGroupID string, parameters Definition, options *DefinitionsClientCreateOrUpdateAtManagementGroupOptions) (DefinitionsClientCreateOrUpdateAtManagementGroupResponse, error) {
	var err error
	const operationName = "DefinitionsClient.CreateOrUpdateAtManagementGroup"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateAtManagementGroupCreateRequest(ctx, policyDefinitionName, managementGroupID, parameters, options)
	if err != nil {
		return DefinitionsClientCreateOrUpdateAtManagementGroupResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DefinitionsClientCreateOrUpdateAtManagementGroupResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return DefinitionsClientCreateOrUpdateAtManagementGroupResponse{}, err
	}
	resp, err := client.createOrUpdateAtManagementGroupHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateAtManagementGroupCreateRequest creates the CreateOrUpdateAtManagementGroup request.
func (client *DefinitionsClient) createOrUpdateAtManagementGroupCreateRequest(ctx context.Context, policyDefinitionName string, managementGroupID string, parameters Definition, options *DefinitionsClientCreateOrUpdateAtManagementGroupOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}"
	if policyDefinitionName == "" {
		return nil, errors.New("parameter policyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyDefinitionName}", url.PathEscape(policyDefinitionName))
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateAtManagementGroupHandleResponse handles the CreateOrUpdateAtManagementGroup response.
func (client *DefinitionsClient) createOrUpdateAtManagementGroupHandleResponse(resp *http.Response) (DefinitionsClientCreateOrUpdateAtManagementGroupResponse, error) {
	result := DefinitionsClientCreateOrUpdateAtManagementGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Definition); err != nil {
		return DefinitionsClientCreateOrUpdateAtManagementGroupResponse{}, err
	}
	return result, nil
}

// Delete - This operation deletes the policy definition in the given subscription with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - policyDefinitionName - The name of the policy definition to delete.
//   - options - DefinitionsClientDeleteOptions contains the optional parameters for the DefinitionsClient.Delete method.
func (client *DefinitionsClient) Delete(ctx context.Context, policyDefinitionName string, options *DefinitionsClientDeleteOptions) (DefinitionsClientDeleteResponse, error) {
	var err error
	const operationName = "DefinitionsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, policyDefinitionName, options)
	if err != nil {
		return DefinitionsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DefinitionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DefinitionsClientDeleteResponse{}, err
	}
	return DefinitionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DefinitionsClient) deleteCreateRequest(ctx context.Context, policyDefinitionName string, options *DefinitionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}"
	if policyDefinitionName == "" {
		return nil, errors.New("parameter policyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyDefinitionName}", url.PathEscape(policyDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// DeleteAtManagementGroup - This operation deletes the policy definition in the given management group with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - policyDefinitionName - The name of the policy definition to delete.
//   - managementGroupID - The ID of the management group.
//   - options - DefinitionsClientDeleteAtManagementGroupOptions contains the optional parameters for the DefinitionsClient.DeleteAtManagementGroup
//     method.
func (client *DefinitionsClient) DeleteAtManagementGroup(ctx context.Context, policyDefinitionName string, managementGroupID string, options *DefinitionsClientDeleteAtManagementGroupOptions) (DefinitionsClientDeleteAtManagementGroupResponse, error) {
	var err error
	const operationName = "DefinitionsClient.DeleteAtManagementGroup"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteAtManagementGroupCreateRequest(ctx, policyDefinitionName, managementGroupID, options)
	if err != nil {
		return DefinitionsClientDeleteAtManagementGroupResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DefinitionsClientDeleteAtManagementGroupResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DefinitionsClientDeleteAtManagementGroupResponse{}, err
	}
	return DefinitionsClientDeleteAtManagementGroupResponse{}, nil
}

// deleteAtManagementGroupCreateRequest creates the DeleteAtManagementGroup request.
func (client *DefinitionsClient) deleteAtManagementGroupCreateRequest(ctx context.Context, policyDefinitionName string, managementGroupID string, options *DefinitionsClientDeleteAtManagementGroupOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}"
	if policyDefinitionName == "" {
		return nil, errors.New("parameter policyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyDefinitionName}", url.PathEscape(policyDefinitionName))
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - This operation retrieves the policy definition in the given subscription with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - policyDefinitionName - The name of the policy definition to get.
//   - options - DefinitionsClientGetOptions contains the optional parameters for the DefinitionsClient.Get method.
func (client *DefinitionsClient) Get(ctx context.Context, policyDefinitionName string, options *DefinitionsClientGetOptions) (DefinitionsClientGetResponse, error) {
	var err error
	const operationName = "DefinitionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, policyDefinitionName, options)
	if err != nil {
		return DefinitionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DefinitionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DefinitionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DefinitionsClient) getCreateRequest(ctx context.Context, policyDefinitionName string, options *DefinitionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}"
	if policyDefinitionName == "" {
		return nil, errors.New("parameter policyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyDefinitionName}", url.PathEscape(policyDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DefinitionsClient) getHandleResponse(resp *http.Response) (DefinitionsClientGetResponse, error) {
	result := DefinitionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Definition); err != nil {
		return DefinitionsClientGetResponse{}, err
	}
	return result, nil
}

// GetAtManagementGroup - This operation retrieves the policy definition in the given management group with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - policyDefinitionName - The name of the policy definition to get.
//   - managementGroupID - The ID of the management group.
//   - options - DefinitionsClientGetAtManagementGroupOptions contains the optional parameters for the DefinitionsClient.GetAtManagementGroup
//     method.
func (client *DefinitionsClient) GetAtManagementGroup(ctx context.Context, policyDefinitionName string, managementGroupID string, options *DefinitionsClientGetAtManagementGroupOptions) (DefinitionsClientGetAtManagementGroupResponse, error) {
	var err error
	const operationName = "DefinitionsClient.GetAtManagementGroup"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getAtManagementGroupCreateRequest(ctx, policyDefinitionName, managementGroupID, options)
	if err != nil {
		return DefinitionsClientGetAtManagementGroupResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DefinitionsClientGetAtManagementGroupResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DefinitionsClientGetAtManagementGroupResponse{}, err
	}
	resp, err := client.getAtManagementGroupHandleResponse(httpResp)
	return resp, err
}

// getAtManagementGroupCreateRequest creates the GetAtManagementGroup request.
func (client *DefinitionsClient) getAtManagementGroupCreateRequest(ctx context.Context, policyDefinitionName string, managementGroupID string, options *DefinitionsClientGetAtManagementGroupOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}"
	if policyDefinitionName == "" {
		return nil, errors.New("parameter policyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyDefinitionName}", url.PathEscape(policyDefinitionName))
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAtManagementGroupHandleResponse handles the GetAtManagementGroup response.
func (client *DefinitionsClient) getAtManagementGroupHandleResponse(resp *http.Response) (DefinitionsClientGetAtManagementGroupResponse, error) {
	result := DefinitionsClientGetAtManagementGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Definition); err != nil {
		return DefinitionsClientGetAtManagementGroupResponse{}, err
	}
	return result, nil
}

// GetBuiltIn - This operation retrieves the built-in policy definition with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - policyDefinitionName - The name of the built-in policy definition to get.
//   - options - DefinitionsClientGetBuiltInOptions contains the optional parameters for the DefinitionsClient.GetBuiltIn method.
func (client *DefinitionsClient) GetBuiltIn(ctx context.Context, policyDefinitionName string, options *DefinitionsClientGetBuiltInOptions) (DefinitionsClientGetBuiltInResponse, error) {
	var err error
	const operationName = "DefinitionsClient.GetBuiltIn"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getBuiltInCreateRequest(ctx, policyDefinitionName, options)
	if err != nil {
		return DefinitionsClientGetBuiltInResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DefinitionsClientGetBuiltInResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DefinitionsClientGetBuiltInResponse{}, err
	}
	resp, err := client.getBuiltInHandleResponse(httpResp)
	return resp, err
}

// getBuiltInCreateRequest creates the GetBuiltIn request.
func (client *DefinitionsClient) getBuiltInCreateRequest(ctx context.Context, policyDefinitionName string, options *DefinitionsClientGetBuiltInOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}"
	if policyDefinitionName == "" {
		return nil, errors.New("parameter policyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyDefinitionName}", url.PathEscape(policyDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getBuiltInHandleResponse handles the GetBuiltIn response.
func (client *DefinitionsClient) getBuiltInHandleResponse(resp *http.Response) (DefinitionsClientGetBuiltInResponse, error) {
	result := DefinitionsClientGetBuiltInResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Definition); err != nil {
		return DefinitionsClientGetBuiltInResponse{}, err
	}
	return result, nil
}

// NewListPager - This operation retrieves a list of all the policy definitions in a given subscription that match the optional
// given $filter. Valid values for $filter are: 'atExactScope()', 'policyType -eq {value}' or
// 'category eq '{value}”. If $filter is not provided, the unfiltered list includes all policy definitions associated with
// the subscription, including those that apply directly or from management groups
// that contain the given subscription. If $filter=atExactScope() is provided, the returned list only includes all policy
// definitions that at the given subscription. If $filter='policyType -eq {value}'
// is provided, the returned list only includes all policy definitions whose type match the {value}. Possible policyType values
// are NotSpecified, BuiltIn, Custom, and Static. If $filter='category -eq
// {value}' is provided, the returned list only includes all policy definitions whose category match the {value}.
//
// Generated from API version 2021-06-01
//   - options - DefinitionsClientListOptions contains the optional parameters for the DefinitionsClient.NewListPager method.
func (client *DefinitionsClient) NewListPager(options *DefinitionsClientListOptions) *runtime.Pager[DefinitionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DefinitionsClientListResponse]{
		More: func(page DefinitionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DefinitionsClientListResponse) (DefinitionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DefinitionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return DefinitionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *DefinitionsClient) listCreateRequest(ctx context.Context, options *DefinitionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DefinitionsClient) listHandleResponse(resp *http.Response) (DefinitionsClientListResponse, error) {
	result := DefinitionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DefinitionListResult); err != nil {
		return DefinitionsClientListResponse{}, err
	}
	return result, nil
}

// NewListBuiltInPager - This operation retrieves a list of all the built-in policy definitions that match the optional given
// $filter. If $filter='policyType -eq {value}' is provided, the returned list only includes all
// built-in policy definitions whose type match the {value}. Possible policyType values are NotSpecified, BuiltIn, Custom,
// and Static. If $filter='category -eq {value}' is provided, the returned list
// only includes all built-in policy definitions whose category match the {value}.
//
// Generated from API version 2021-06-01
//   - options - DefinitionsClientListBuiltInOptions contains the optional parameters for the DefinitionsClient.NewListBuiltInPager
//     method.
func (client *DefinitionsClient) NewListBuiltInPager(options *DefinitionsClientListBuiltInOptions) *runtime.Pager[DefinitionsClientListBuiltInResponse] {
	return runtime.NewPager(runtime.PagingHandler[DefinitionsClientListBuiltInResponse]{
		More: func(page DefinitionsClientListBuiltInResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DefinitionsClientListBuiltInResponse) (DefinitionsClientListBuiltInResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DefinitionsClient.NewListBuiltInPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBuiltInCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return DefinitionsClientListBuiltInResponse{}, err
			}
			return client.listBuiltInHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBuiltInCreateRequest creates the ListBuiltIn request.
func (client *DefinitionsClient) listBuiltInCreateRequest(ctx context.Context, options *DefinitionsClientListBuiltInOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/policyDefinitions"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBuiltInHandleResponse handles the ListBuiltIn response.
func (client *DefinitionsClient) listBuiltInHandleResponse(resp *http.Response) (DefinitionsClientListBuiltInResponse, error) {
	result := DefinitionsClientListBuiltInResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DefinitionListResult); err != nil {
		return DefinitionsClientListBuiltInResponse{}, err
	}
	return result, nil
}

// NewListByManagementGroupPager - This operation retrieves a list of all the policy definitions in a given management group
// that match the optional given $filter. Valid values for $filter are: 'atExactScope()', 'policyType -eq
// {value}' or 'category eq '{value}”. If $filter is not provided, the unfiltered list includes all policy definitions associated
// with the management group, including those that apply directly or from
// management groups that contain the given management group. If $filter=atExactScope() is provided, the returned list only
// includes all policy definitions that at the given management group. If
// $filter='policyType -eq {value}' is provided, the returned list only includes all policy definitions whose type match the
// {value}. Possible policyType values are NotSpecified, BuiltIn, Custom, and
// Static. If $filter='category -eq {value}' is provided, the returned list only includes all policy definitions whose category
// match the {value}.
//
// Generated from API version 2021-06-01
//   - managementGroupID - The ID of the management group.
//   - options - DefinitionsClientListByManagementGroupOptions contains the optional parameters for the DefinitionsClient.NewListByManagementGroupPager
//     method.
func (client *DefinitionsClient) NewListByManagementGroupPager(managementGroupID string, options *DefinitionsClientListByManagementGroupOptions) *runtime.Pager[DefinitionsClientListByManagementGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[DefinitionsClientListByManagementGroupResponse]{
		More: func(page DefinitionsClientListByManagementGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DefinitionsClientListByManagementGroupResponse) (DefinitionsClientListByManagementGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DefinitionsClient.NewListByManagementGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByManagementGroupCreateRequest(ctx, managementGroupID, options)
			}, nil)
			if err != nil {
				return DefinitionsClientListByManagementGroupResponse{}, err
			}
			return client.listByManagementGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByManagementGroupCreateRequest creates the ListByManagementGroup request.
func (client *DefinitionsClient) listByManagementGroupCreateRequest(ctx context.Context, managementGroupID string, options *DefinitionsClientListByManagementGroupOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions"
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByManagementGroupHandleResponse handles the ListByManagementGroup response.
func (client *DefinitionsClient) listByManagementGroupHandleResponse(resp *http.Response) (DefinitionsClientListByManagementGroupResponse, error) {
	result := DefinitionsClientListByManagementGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DefinitionListResult); err != nil {
		return DefinitionsClientListByManagementGroupResponse{}, err
	}
	return result, nil
}
