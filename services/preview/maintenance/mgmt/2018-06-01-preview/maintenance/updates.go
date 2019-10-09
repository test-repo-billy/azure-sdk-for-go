package maintenance

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// UpdatesClient is the azure Maintenance Management Client
type UpdatesClient struct {
	BaseClient
}

// NewUpdatesClient creates an instance of the UpdatesClient client.
func NewUpdatesClient(subscriptionID string) UpdatesClient {
	return NewUpdatesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUpdatesClientWithBaseURI creates an instance of the UpdatesClient client.
func NewUpdatesClientWithBaseURI(baseURI string, subscriptionID string) UpdatesClient {
	return UpdatesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List get updates to resources.
// Parameters:
// resourceGroupName - resource group name
// providerName - resource provider name
// resourceType - resource type
// resourceName - resource identifier
func (client UpdatesClient) List(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string) (result ListUpdatesResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UpdatesClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, resourceGroupName, providerName, resourceType, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.UpdatesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "maintenance.UpdatesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.UpdatesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client UpdatesClient) ListPreparer(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"providerName":      autorest.Encode("path", providerName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"resourceType":      autorest.Encode("path", resourceType),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/updates", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client UpdatesClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client UpdatesClient) ListResponder(resp *http.Response) (result ListUpdatesResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListParent get updates to resources.
// Parameters:
// resourceGroupName - resource group name
// providerName - resource provider name
// resourceParentType - resource parent type
// resourceParentName - resource parent identifier
// resourceType - resource type
// resourceName - resource identifier
func (client UpdatesClient) ListParent(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string) (result ListUpdatesResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UpdatesClient.ListParent")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListParentPreparer(ctx, resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.UpdatesClient", "ListParent", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListParentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "maintenance.UpdatesClient", "ListParent", resp, "Failure sending request")
		return
	}

	result, err = client.ListParentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "maintenance.UpdatesClient", "ListParent", resp, "Failure responding to request")
	}

	return
}

// ListParentPreparer prepares the ListParent request.
func (client UpdatesClient) ListParentPreparer(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"providerName":       autorest.Encode("path", providerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"resourceName":       autorest.Encode("path", resourceName),
		"resourceParentName": autorest.Encode("path", resourceParentName),
		"resourceParentType": autorest.Encode("path", resourceParentType),
		"resourceType":       autorest.Encode("path", resourceType),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceParentType}/{resourceParentName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/updates", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListParentSender sends the ListParent request. The method will close the
// http.Response Body if it receives an error.
func (client UpdatesClient) ListParentSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListParentResponder handles the response to the ListParent request. The method always
// closes the http.Response Body.
func (client UpdatesClient) ListParentResponder(resp *http.Response) (result ListUpdatesResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
