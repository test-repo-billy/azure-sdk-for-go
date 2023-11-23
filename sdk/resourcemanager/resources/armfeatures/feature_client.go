//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfeatures

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// FeatureClient contains the methods for the FeatureClient group.
// Don't use this type directly, use NewFeatureClient() instead.
type FeatureClient struct {
	internal *arm.Client
}

// NewFeatureClient creates a new instance of FeatureClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFeatureClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*FeatureClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FeatureClient{
		internal: cl,
	}
	return client, nil
}

// NewListOperationsPager - Lists all of the available Microsoft.Features REST API operations.
//
// Generated from API version 2021-07-01
//   - options - FeatureClientListOperationsOptions contains the optional parameters for the FeatureClient.NewListOperationsPager
//     method.
func (client *FeatureClient) NewListOperationsPager(options *FeatureClientListOperationsOptions) *runtime.Pager[FeatureClientListOperationsResponse] {
	return runtime.NewPager(runtime.PagingHandler[FeatureClientListOperationsResponse]{
		More: func(page FeatureClientListOperationsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FeatureClientListOperationsResponse) (FeatureClientListOperationsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "FeatureClient.NewListOperationsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listOperationsCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return FeatureClientListOperationsResponse{}, err
			}
			return client.listOperationsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listOperationsCreateRequest creates the ListOperations request.
func (client *FeatureClient) listOperationsCreateRequest(ctx context.Context, options *FeatureClientListOperationsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Features/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json, text/json"}
	return req, nil
}

// listOperationsHandleResponse handles the ListOperations response.
func (client *FeatureClient) listOperationsHandleResponse(resp *http.Response) (FeatureClientListOperationsResponse, error) {
	result := FeatureClientListOperationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationListResult); err != nil {
		return FeatureClientListOperationsResponse{}, err
	}
	return result, nil
}
