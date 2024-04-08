//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhealthcareapis

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DicomServicesClient contains the methods for the DicomServices group.
// Don't use this type directly, use NewDicomServicesClient() instead.
type DicomServicesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDicomServicesClient creates a new instance of DicomServicesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDicomServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DicomServicesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DicomServicesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a DICOM Service resource with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - workspaceName - The name of workspace resource.
//   - dicomServiceName - The name of DICOM Service resource.
//   - dicomservice - The parameters for creating or updating a Dicom Service resource.
//   - options - DicomServicesClientBeginCreateOrUpdateOptions contains the optional parameters for the DicomServicesClient.BeginCreateOrUpdate
//     method.
func (client *DicomServicesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, dicomServiceName string, dicomservice DicomService, options *DicomServicesClientBeginCreateOrUpdateOptions) (*runtime.Poller[DicomServicesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, dicomServiceName, dicomservice, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DicomServicesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DicomServicesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a DICOM Service resource with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
func (client *DicomServicesClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, dicomServiceName string, dicomservice DicomService, options *DicomServicesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DicomServicesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, dicomServiceName, dicomservice, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DicomServicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dicomServiceName string, dicomservice DicomService, options *DicomServicesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/dicomservices/{dicomServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dicomServiceName == "" {
		return nil, errors.New("parameter dicomServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dicomServiceName}", url.PathEscape(dicomServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, dicomservice); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a DICOM Service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - dicomServiceName - The name of DICOM Service resource.
//   - workspaceName - The name of workspace resource.
//   - options - DicomServicesClientBeginDeleteOptions contains the optional parameters for the DicomServicesClient.BeginDelete
//     method.
func (client *DicomServicesClient) BeginDelete(ctx context.Context, resourceGroupName string, dicomServiceName string, workspaceName string, options *DicomServicesClientBeginDeleteOptions) (*runtime.Poller[DicomServicesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, dicomServiceName, workspaceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DicomServicesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DicomServicesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a DICOM Service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
func (client *DicomServicesClient) deleteOperation(ctx context.Context, resourceGroupName string, dicomServiceName string, workspaceName string, options *DicomServicesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "DicomServicesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, dicomServiceName, workspaceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DicomServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, dicomServiceName string, workspaceName string, options *DicomServicesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/dicomservices/{dicomServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dicomServiceName == "" {
		return nil, errors.New("parameter dicomServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dicomServiceName}", url.PathEscape(dicomServiceName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the properties of the specified DICOM Service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - workspaceName - The name of workspace resource.
//   - dicomServiceName - The name of DICOM Service resource.
//   - options - DicomServicesClientGetOptions contains the optional parameters for the DicomServicesClient.Get method.
func (client *DicomServicesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, dicomServiceName string, options *DicomServicesClientGetOptions) (DicomServicesClientGetResponse, error) {
	var err error
	const operationName = "DicomServicesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, dicomServiceName, options)
	if err != nil {
		return DicomServicesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DicomServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DicomServicesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DicomServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dicomServiceName string, options *DicomServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/dicomservices/{dicomServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dicomServiceName == "" {
		return nil, errors.New("parameter dicomServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dicomServiceName}", url.PathEscape(dicomServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DicomServicesClient) getHandleResponse(resp *http.Response) (DicomServicesClientGetResponse, error) {
	result := DicomServicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DicomService); err != nil {
		return DicomServicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByWorkspacePager - Lists all DICOM Services for the given workspace
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - workspaceName - The name of workspace resource.
//   - options - DicomServicesClientListByWorkspaceOptions contains the optional parameters for the DicomServicesClient.NewListByWorkspacePager
//     method.
func (client *DicomServicesClient) NewListByWorkspacePager(resourceGroupName string, workspaceName string, options *DicomServicesClientListByWorkspaceOptions) *runtime.Pager[DicomServicesClientListByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[DicomServicesClientListByWorkspaceResponse]{
		More: func(page DicomServicesClientListByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DicomServicesClientListByWorkspaceResponse) (DicomServicesClientListByWorkspaceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DicomServicesClient.NewListByWorkspacePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
			}, nil)
			if err != nil {
				return DicomServicesClientListByWorkspaceResponse{}, err
			}
			return client.listByWorkspaceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *DicomServicesClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *DicomServicesClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/dicomservices"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *DicomServicesClient) listByWorkspaceHandleResponse(resp *http.Response) (DicomServicesClientListByWorkspaceResponse, error) {
	result := DicomServicesClientListByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DicomServiceCollection); err != nil {
		return DicomServicesClientListByWorkspaceResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patch DICOM Service details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group that contains the service instance.
//   - dicomServiceName - The name of DICOM Service resource.
//   - workspaceName - The name of workspace resource.
//   - dicomservicePatchResource - The parameters for updating a Dicom Service.
//   - options - DicomServicesClientBeginUpdateOptions contains the optional parameters for the DicomServicesClient.BeginUpdate
//     method.
func (client *DicomServicesClient) BeginUpdate(ctx context.Context, resourceGroupName string, dicomServiceName string, workspaceName string, dicomservicePatchResource DicomServicePatchResource, options *DicomServicesClientBeginUpdateOptions) (*runtime.Poller[DicomServicesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, dicomServiceName, workspaceName, dicomservicePatchResource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DicomServicesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DicomServicesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Patch DICOM Service details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
func (client *DicomServicesClient) update(ctx context.Context, resourceGroupName string, dicomServiceName string, workspaceName string, dicomservicePatchResource DicomServicePatchResource, options *DicomServicesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DicomServicesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, dicomServiceName, workspaceName, dicomservicePatchResource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *DicomServicesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, dicomServiceName string, workspaceName string, dicomservicePatchResource DicomServicePatchResource, options *DicomServicesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}/dicomservices/{dicomServiceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if dicomServiceName == "" {
		return nil, errors.New("parameter dicomServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dicomServiceName}", url.PathEscape(dicomServiceName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, dicomservicePatchResource); err != nil {
		return nil, err
	}
	return req, nil
}
