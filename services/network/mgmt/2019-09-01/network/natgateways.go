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

// NatGatewaysClient is the network Client
type NatGatewaysClient struct {
	BaseClient
}

// NewNatGatewaysClient creates an instance of the NatGatewaysClient client.
func NewNatGatewaysClient(subscriptionID string) NatGatewaysClient {
	return NewNatGatewaysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewNatGatewaysClientWithBaseURI creates an instance of the NatGatewaysClient client.
func NewNatGatewaysClientWithBaseURI(baseURI string, subscriptionID string) NatGatewaysClient {
	return NatGatewaysClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a nat gateway.
// Parameters:
// resourceGroupName - the name of the resource group.
// natGatewayName - the name of the nat gateway.
// parameters - parameters supplied to the create or update nat gateway operation.
func (client NatGatewaysClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, natGatewayName string, parameters NatGateway) (result NatGatewaysCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NatGatewaysClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, natGatewayName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NatGatewaysClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NatGatewaysClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client NatGatewaysClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, natGatewayName string, parameters NatGateway) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"natGatewayName":    autorest.Encode("path", natGatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client NatGatewaysClient) CreateOrUpdateSender(req *http.Request) (future NatGatewaysCreateOrUpdateFuture, err error) {
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
func (client NatGatewaysClient) CreateOrUpdateResponder(resp *http.Response) (result NatGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified nat gateway.
// Parameters:
// resourceGroupName - the name of the resource group.
// natGatewayName - the name of the nat gateway.
func (client NatGatewaysClient) Delete(ctx context.Context, resourceGroupName string, natGatewayName string) (result NatGatewaysDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NatGatewaysClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, natGatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NatGatewaysClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NatGatewaysClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client NatGatewaysClient) DeletePreparer(ctx context.Context, resourceGroupName string, natGatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"natGatewayName":    autorest.Encode("path", natGatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client NatGatewaysClient) DeleteSender(req *http.Request) (future NatGatewaysDeleteFuture, err error) {
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
func (client NatGatewaysClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified nat gateway in a specified resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// natGatewayName - the name of the nat gateway.
// expand - expands referenced resources.
func (client NatGatewaysClient) Get(ctx context.Context, resourceGroupName string, natGatewayName string, expand string) (result NatGateway, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NatGatewaysClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, natGatewayName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NatGatewaysClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.NatGatewaysClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NatGatewaysClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client NatGatewaysClient) GetPreparer(ctx context.Context, resourceGroupName string, natGatewayName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"natGatewayName":    autorest.Encode("path", natGatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client NatGatewaysClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client NatGatewaysClient) GetResponder(resp *http.Response) (result NatGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all nat gateways in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client NatGatewaysClient) List(ctx context.Context, resourceGroupName string) (result NatGatewayListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NatGatewaysClient.List")
		defer func() {
			sc := -1
			if result.nglr.Response.Response != nil {
				sc = result.nglr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NatGatewaysClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.nglr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.NatGatewaysClient", "List", resp, "Failure sending request")
		return
	}

	result.nglr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NatGatewaysClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client NatGatewaysClient) ListPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client NatGatewaysClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client NatGatewaysClient) ListResponder(resp *http.Response) (result NatGatewayListResult, err error) {
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
func (client NatGatewaysClient) listNextResults(ctx context.Context, lastResults NatGatewayListResult) (result NatGatewayListResult, err error) {
	req, err := lastResults.natGatewayListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.NatGatewaysClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.NatGatewaysClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NatGatewaysClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client NatGatewaysClient) ListComplete(ctx context.Context, resourceGroupName string) (result NatGatewayListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NatGatewaysClient.List")
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

// ListAll gets all the Nat Gateways in a subscription.
func (client NatGatewaysClient) ListAll(ctx context.Context) (result NatGatewayListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NatGatewaysClient.ListAll")
		defer func() {
			sc := -1
			if result.nglr.Response.Response != nil {
				sc = result.nglr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listAllNextResults
	req, err := client.ListAllPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NatGatewaysClient", "ListAll", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.nglr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.NatGatewaysClient", "ListAll", resp, "Failure sending request")
		return
	}

	result.nglr, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NatGatewaysClient", "ListAll", resp, "Failure responding to request")
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client NatGatewaysClient) ListAllPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/natGateways", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client NatGatewaysClient) ListAllSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client NatGatewaysClient) ListAllResponder(resp *http.Response) (result NatGatewayListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAllNextResults retrieves the next set of results, if any.
func (client NatGatewaysClient) listAllNextResults(ctx context.Context, lastResults NatGatewayListResult) (result NatGatewayListResult, err error) {
	req, err := lastResults.natGatewayListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.NatGatewaysClient", "listAllNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.NatGatewaysClient", "listAllNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NatGatewaysClient", "listAllNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAllComplete enumerates all values, automatically crossing page boundaries as required.
func (client NatGatewaysClient) ListAllComplete(ctx context.Context) (result NatGatewayListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NatGatewaysClient.ListAll")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAll(ctx)
	return
}

// UpdateTags updates nat gateway tags.
// Parameters:
// resourceGroupName - the name of the resource group.
// natGatewayName - the name of the nat gateway.
// parameters - parameters supplied to update nat gateway tags.
func (client NatGatewaysClient) UpdateTags(ctx context.Context, resourceGroupName string, natGatewayName string, parameters TagsObject) (result NatGateway, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NatGatewaysClient.UpdateTags")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateTagsPreparer(ctx, resourceGroupName, natGatewayName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NatGatewaysClient", "UpdateTags", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateTagsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.NatGatewaysClient", "UpdateTags", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateTagsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.NatGatewaysClient", "UpdateTags", resp, "Failure responding to request")
	}

	return
}

// UpdateTagsPreparer prepares the UpdateTags request.
func (client NatGatewaysClient) UpdateTagsPreparer(ctx context.Context, resourceGroupName string, natGatewayName string, parameters TagsObject) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"natGatewayName":    autorest.Encode("path", natGatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/natGateways/{natGatewayName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateTagsSender sends the UpdateTags request. The method will close the
// http.Response Body if it receives an error.
func (client NatGatewaysClient) UpdateTagsSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// UpdateTagsResponder handles the response to the UpdateTags request. The method always
// closes the http.Response Body.
func (client NatGatewaysClient) UpdateTagsResponder(resp *http.Response) (result NatGateway, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
