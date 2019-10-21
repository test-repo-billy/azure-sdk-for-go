package network

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

// PrivateEndpointsClient is the network Client
type PrivateEndpointsClient struct {
	BaseClient
}

// NewPrivateEndpointsClient creates an instance of the PrivateEndpointsClient client.
func NewPrivateEndpointsClient(subscriptionID string) PrivateEndpointsClient {
	return NewPrivateEndpointsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateEndpointsClientWithBaseURI creates an instance of the PrivateEndpointsClient client.
func NewPrivateEndpointsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointsClient {
	return PrivateEndpointsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates an private endpoint in the specified resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// privateEndpointName - the name of the private endpoint.
// parameters - parameters supplied to the create or update private endpoint operation.
func (client PrivateEndpointsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, privateEndpointName string, parameters PrivateEndpoint) (result PrivateEndpointsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, privateEndpointName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateEndpointsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateEndpointsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client PrivateEndpointsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, privateEndpointName string, parameters PrivateEndpoint) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateEndpointName": autorest.Encode("path", privateEndpointName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointsClient) CreateOrUpdateSender(req *http.Request) (future PrivateEndpointsCreateOrUpdateFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client PrivateEndpointsClient) CreateOrUpdateResponder(resp *http.Response) (result PrivateEndpoint, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified private endpoint.
// Parameters:
// resourceGroupName - the name of the resource group.
// privateEndpointName - the name of the private endpoint.
func (client PrivateEndpointsClient) Delete(ctx context.Context, resourceGroupName string, privateEndpointName string) (result PrivateEndpointsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, privateEndpointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateEndpointsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateEndpointsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client PrivateEndpointsClient) DeletePreparer(ctx context.Context, resourceGroupName string, privateEndpointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateEndpointName": autorest.Encode("path", privateEndpointName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointsClient) DeleteSender(req *http.Request) (future PrivateEndpointsDeleteFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PrivateEndpointsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified private endpoint by resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// privateEndpointName - the name of the private endpoint.
// expand - expands referenced resources.
func (client PrivateEndpointsClient) Get(ctx context.Context, resourceGroupName string, privateEndpointName string, expand string) (result PrivateEndpoint, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, privateEndpointName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateEndpointsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PrivateEndpointsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateEndpointsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client PrivateEndpointsClient) GetPreparer(ctx context.Context, resourceGroupName string, privateEndpointName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"privateEndpointName": autorest.Encode("path", privateEndpointName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PrivateEndpointsClient) GetResponder(resp *http.Response) (result PrivateEndpoint, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all private endpoints in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client PrivateEndpointsClient) List(ctx context.Context, resourceGroupName string) (result PrivateEndpointListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointsClient.List")
		defer func() {
			sc := -1
			if result.pelr.Response.Response != nil {
				sc = result.pelr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateEndpointsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.pelr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PrivateEndpointsClient", "List", resp, "Failure sending request")
		return
	}

	result.pelr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateEndpointsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client PrivateEndpointsClient) ListPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointsClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PrivateEndpointsClient) ListResponder(resp *http.Response) (result PrivateEndpointListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client PrivateEndpointsClient) listNextResults(ctx context.Context, lastResults PrivateEndpointListResult) (result PrivateEndpointListResult, err error) {
	req, err := lastResults.privateEndpointListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.PrivateEndpointsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.PrivateEndpointsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateEndpointsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client PrivateEndpointsClient) ListComplete(ctx context.Context, resourceGroupName string) (result PrivateEndpointListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName)
	return
}

// ListBySubscription gets all private endpoints in a subscription.
func (client PrivateEndpointsClient) ListBySubscription(ctx context.Context) (result PrivateEndpointListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.pelr.Response.Response != nil {
				sc = result.pelr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateEndpointsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.pelr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.PrivateEndpointsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.pelr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateEndpointsClient", "ListBySubscription", resp, "Failure responding to request")
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client PrivateEndpointsClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/privateEndpoints", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client PrivateEndpointsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client PrivateEndpointsClient) ListBySubscriptionResponder(resp *http.Response) (result PrivateEndpointListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client PrivateEndpointsClient) listBySubscriptionNextResults(ctx context.Context, lastResults PrivateEndpointListResult) (result PrivateEndpointListResult, err error) {
	req, err := lastResults.privateEndpointListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.PrivateEndpointsClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.PrivateEndpointsClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.PrivateEndpointsClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client PrivateEndpointsClient) ListBySubscriptionComplete(ctx context.Context) (result PrivateEndpointListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrivateEndpointsClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx)
	return
}
