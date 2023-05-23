//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

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

// VirtualApplianceConnectionsClient contains the methods for the NetworkVirtualApplianceConnections group.
// Don't use this type directly, use NewVirtualApplianceConnectionsClient() instead.
type VirtualApplianceConnectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualApplianceConnectionsClient creates a new instance of VirtualApplianceConnectionsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualApplianceConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualApplianceConnectionsClient, error) {
	cl, err := arm.NewClient(moduleName+".VirtualApplianceConnectionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualApplianceConnectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a connection to Network Virtual Appliance, if it doesn't exist else updates the existing
// NVA connection'
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group.
//   - networkVirtualApplianceName - The name of the Network Virtual Appliance.
//   - connectionName - The name of the NVA connection.
//   - networkVirtualApplianceConnectionParameters - Parameters supplied in an NetworkVirtualApplianceConnection PUT operation.
//   - options - VirtualApplianceConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualApplianceConnectionsClient.BeginCreateOrUpdate
//     method.
func (client *VirtualApplianceConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, connectionName string, networkVirtualApplianceConnectionParameters VirtualApplianceConnection, options *VirtualApplianceConnectionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualApplianceConnectionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, networkVirtualApplianceName, connectionName, networkVirtualApplianceConnectionParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualApplianceConnectionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VirtualApplianceConnectionsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates a connection to Network Virtual Appliance, if it doesn't exist else updates the existing NVA connection'
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
func (client *VirtualApplianceConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, connectionName string, networkVirtualApplianceConnectionParameters VirtualApplianceConnection, options *VirtualApplianceConnectionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, connectionName, networkVirtualApplianceConnectionParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualApplianceConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, connectionName string, networkVirtualApplianceConnectionParameters VirtualApplianceConnection, options *VirtualApplianceConnectionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/networkVirtualApplianceConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, networkVirtualApplianceConnectionParameters)
}

// BeginDelete - Deletes a NVA connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group.
//   - networkVirtualApplianceName - The name of the Network Virtual Appliance.
//   - connectionName - The name of the NVA connection.
//   - options - VirtualApplianceConnectionsClientBeginDeleteOptions contains the optional parameters for the VirtualApplianceConnectionsClient.BeginDelete
//     method.
func (client *VirtualApplianceConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, connectionName string, options *VirtualApplianceConnectionsClientBeginDeleteOptions) (*runtime.Poller[VirtualApplianceConnectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, networkVirtualApplianceName, connectionName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualApplianceConnectionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VirtualApplianceConnectionsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a NVA connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
func (client *VirtualApplianceConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, connectionName string, options *VirtualApplianceConnectionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, connectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualApplianceConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, connectionName string, options *VirtualApplianceConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/networkVirtualApplianceConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves the details of specified NVA connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group.
//   - networkVirtualApplianceName - The name of the Network Virtual Appliance.
//   - connectionName - The name of the NVA connection.
//   - options - VirtualApplianceConnectionsClientGetOptions contains the optional parameters for the VirtualApplianceConnectionsClient.Get
//     method.
func (client *VirtualApplianceConnectionsClient) Get(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, connectionName string, options *VirtualApplianceConnectionsClientGetOptions) (VirtualApplianceConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, connectionName, options)
	if err != nil {
		return VirtualApplianceConnectionsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualApplianceConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualApplianceConnectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualApplianceConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, connectionName string, options *VirtualApplianceConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/networkVirtualApplianceConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualApplianceConnectionsClient) getHandleResponse(resp *http.Response) (VirtualApplianceConnectionsClientGetResponse, error) {
	result := VirtualApplianceConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualApplianceConnection); err != nil {
		return VirtualApplianceConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists NetworkVirtualApplianceConnections under the NVA.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group.
//   - networkVirtualApplianceName - The name of the Network Virtual Appliance.
//   - options - VirtualApplianceConnectionsClientListOptions contains the optional parameters for the VirtualApplianceConnectionsClient.NewListPager
//     method.
func (client *VirtualApplianceConnectionsClient) NewListPager(resourceGroupName string, networkVirtualApplianceName string, options *VirtualApplianceConnectionsClientListOptions) *runtime.Pager[VirtualApplianceConnectionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualApplianceConnectionsClientListResponse]{
		More: func(page VirtualApplianceConnectionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualApplianceConnectionsClientListResponse) (VirtualApplianceConnectionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, networkVirtualApplianceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualApplianceConnectionsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return VirtualApplianceConnectionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualApplianceConnectionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VirtualApplianceConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *VirtualApplianceConnectionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkVirtualAppliances/{networkVirtualApplianceName}/networkVirtualApplianceConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkVirtualApplianceName == "" {
		return nil, errors.New("parameter networkVirtualApplianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkVirtualApplianceName}", url.PathEscape(networkVirtualApplianceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualApplianceConnectionsClient) listHandleResponse(resp *http.Response) (VirtualApplianceConnectionsClientListResponse, error) {
	result := VirtualApplianceConnectionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualApplianceConnectionList); err != nil {
		return VirtualApplianceConnectionsClientListResponse{}, err
	}
	return result, nil
}
