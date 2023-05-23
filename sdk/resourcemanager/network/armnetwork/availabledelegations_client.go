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

// AvailableDelegationsClient contains the methods for the AvailableDelegations group.
// Don't use this type directly, use NewAvailableDelegationsClient() instead.
type AvailableDelegationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAvailableDelegationsClient creates a new instance of AvailableDelegationsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAvailableDelegationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AvailableDelegationsClient, error) {
	cl, err := arm.NewClient(moduleName+".AvailableDelegationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AvailableDelegationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Gets all of the available subnet delegations for this subscription in this region.
//
// Generated from API version 2022-11-01
//   - location - The location of the subnet.
//   - options - AvailableDelegationsClientListOptions contains the optional parameters for the AvailableDelegationsClient.NewListPager
//     method.
func (client *AvailableDelegationsClient) NewListPager(location string, options *AvailableDelegationsClientListOptions) *runtime.Pager[AvailableDelegationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AvailableDelegationsClientListResponse]{
		More: func(page AvailableDelegationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AvailableDelegationsClientListResponse) (AvailableDelegationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, location, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AvailableDelegationsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return AvailableDelegationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AvailableDelegationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AvailableDelegationsClient) listCreateRequest(ctx context.Context, location string, options *AvailableDelegationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availableDelegations"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *AvailableDelegationsClient) listHandleResponse(resp *http.Response) (AvailableDelegationsClientListResponse, error) {
	result := AvailableDelegationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailableDelegationsResult); err != nil {
		return AvailableDelegationsClientListResponse{}, err
	}
	return result, nil
}
