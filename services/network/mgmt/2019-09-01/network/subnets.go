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

// SubnetsClient is the network Client
type SubnetsClient struct {
	BaseClient
}

// NewSubnetsClient creates an instance of the SubnetsClient client.
func NewSubnetsClient(subscriptionID string) SubnetsClient {
	return NewSubnetsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSubnetsClientWithBaseURI creates an instance of the SubnetsClient client.
func NewSubnetsClientWithBaseURI(baseURI string, subscriptionID string) SubnetsClient {
	return SubnetsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a subnet in the specified virtual network.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualNetworkName - the name of the virtual network.
// subnetName - the name of the subnet.
// subnetParameters - parameters supplied to the create or update subnet operation.
func (client SubnetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters Subnet) (result SubnetsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubnetsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, virtualNetworkName, subnetName, subnetParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SubnetsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SubnetsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SubnetsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters Subnet) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subnetName":         autorest.Encode("path", subnetName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName": autorest.Encode("path", virtualNetworkName),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}", pathParameters),
		autorest.WithJSON(subnetParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SubnetsClient) CreateOrUpdateSender(req *http.Request) (future SubnetsCreateOrUpdateFuture, err error) {
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
func (client SubnetsClient) CreateOrUpdateResponder(resp *http.Response) (result Subnet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified subnet.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualNetworkName - the name of the virtual network.
// subnetName - the name of the subnet.
func (client SubnetsClient) Delete(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string) (result SubnetsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubnetsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, virtualNetworkName, subnetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SubnetsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SubnetsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SubnetsClient) DeletePreparer(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subnetName":         autorest.Encode("path", subnetName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName": autorest.Encode("path", virtualNetworkName),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SubnetsClient) DeleteSender(req *http.Request) (future SubnetsDeleteFuture, err error) {
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
func (client SubnetsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified subnet by virtual network and resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualNetworkName - the name of the virtual network.
// subnetName - the name of the subnet.
// expand - expands referenced resources.
func (client SubnetsClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, expand string) (result Subnet, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubnetsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, virtualNetworkName, subnetName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SubnetsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.SubnetsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SubnetsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SubnetsClient) GetPreparer(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subnetName":         autorest.Encode("path", subnetName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName": autorest.Encode("path", virtualNetworkName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SubnetsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SubnetsClient) GetResponder(resp *http.Response) (result Subnet, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all subnets in a virtual network.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualNetworkName - the name of the virtual network.
func (client SubnetsClient) List(ctx context.Context, resourceGroupName string, virtualNetworkName string) (result SubnetListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubnetsClient.List")
		defer func() {
			sc := -1
			if result.slr.Response.Response != nil {
				sc = result.slr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, virtualNetworkName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SubnetsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.slr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.SubnetsClient", "List", resp, "Failure sending request")
		return
	}

	result.slr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SubnetsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client SubnetsClient) ListPreparer(ctx context.Context, resourceGroupName string, virtualNetworkName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName": autorest.Encode("path", virtualNetworkName),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SubnetsClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SubnetsClient) ListResponder(resp *http.Response) (result SubnetListResult, err error) {
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
func (client SubnetsClient) listNextResults(ctx context.Context, lastResults SubnetListResult) (result SubnetListResult, err error) {
	req, err := lastResults.subnetListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.SubnetsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.SubnetsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SubnetsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client SubnetsClient) ListComplete(ctx context.Context, resourceGroupName string, virtualNetworkName string) (result SubnetListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubnetsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, virtualNetworkName)
	return
}

// PrepareNetworkPolicies prepares a subnet by applying network intent policies.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualNetworkName - the name of the virtual network.
// subnetName - the name of the subnet.
// prepareNetworkPoliciesRequestParameters - parameters supplied to prepare subnet by applying network intent
// policies.
func (client SubnetsClient) PrepareNetworkPolicies(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, prepareNetworkPoliciesRequestParameters PrepareNetworkPoliciesRequest) (result SubnetsPrepareNetworkPoliciesFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubnetsClient.PrepareNetworkPolicies")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PrepareNetworkPoliciesPreparer(ctx, resourceGroupName, virtualNetworkName, subnetName, prepareNetworkPoliciesRequestParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SubnetsClient", "PrepareNetworkPolicies", nil, "Failure preparing request")
		return
	}

	result, err = client.PrepareNetworkPoliciesSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SubnetsClient", "PrepareNetworkPolicies", result.Response(), "Failure sending request")
		return
	}

	return
}

// PrepareNetworkPoliciesPreparer prepares the PrepareNetworkPolicies request.
func (client SubnetsClient) PrepareNetworkPoliciesPreparer(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, prepareNetworkPoliciesRequestParameters PrepareNetworkPoliciesRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subnetName":         autorest.Encode("path", subnetName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName": autorest.Encode("path", virtualNetworkName),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/PrepareNetworkPolicies", pathParameters),
		autorest.WithJSON(prepareNetworkPoliciesRequestParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PrepareNetworkPoliciesSender sends the PrepareNetworkPolicies request. The method will close the
// http.Response Body if it receives an error.
func (client SubnetsClient) PrepareNetworkPoliciesSender(req *http.Request) (future SubnetsPrepareNetworkPoliciesFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// PrepareNetworkPoliciesResponder handles the response to the PrepareNetworkPolicies request. The method always
// closes the http.Response Body.
func (client SubnetsClient) PrepareNetworkPoliciesResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// UnprepareNetworkPolicies unprepares a subnet by removing network intent policies.
// Parameters:
// resourceGroupName - the name of the resource group.
// virtualNetworkName - the name of the virtual network.
// subnetName - the name of the subnet.
// unprepareNetworkPoliciesRequestParameters - parameters supplied to unprepare subnet to remove network intent
// policies.
func (client SubnetsClient) UnprepareNetworkPolicies(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, unprepareNetworkPoliciesRequestParameters UnprepareNetworkPoliciesRequest) (result SubnetsUnprepareNetworkPoliciesFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubnetsClient.UnprepareNetworkPolicies")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UnprepareNetworkPoliciesPreparer(ctx, resourceGroupName, virtualNetworkName, subnetName, unprepareNetworkPoliciesRequestParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SubnetsClient", "UnprepareNetworkPolicies", nil, "Failure preparing request")
		return
	}

	result, err = client.UnprepareNetworkPoliciesSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SubnetsClient", "UnprepareNetworkPolicies", result.Response(), "Failure sending request")
		return
	}

	return
}

// UnprepareNetworkPoliciesPreparer prepares the UnprepareNetworkPolicies request.
func (client SubnetsClient) UnprepareNetworkPoliciesPreparer(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, unprepareNetworkPoliciesRequestParameters UnprepareNetworkPoliciesRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subnetName":         autorest.Encode("path", subnetName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkName": autorest.Encode("path", virtualNetworkName),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/UnprepareNetworkPolicies", pathParameters),
		autorest.WithJSON(unprepareNetworkPoliciesRequestParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UnprepareNetworkPoliciesSender sends the UnprepareNetworkPolicies request. The method will close the
// http.Response Body if it receives an error.
func (client SubnetsClient) UnprepareNetworkPoliciesSender(req *http.Request) (future SubnetsUnprepareNetworkPoliciesFuture, err error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req, sd...)
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UnprepareNetworkPoliciesResponder handles the response to the UnprepareNetworkPolicies request. The method always
// closes the http.Response Body.
func (client SubnetsClient) UnprepareNetworkPoliciesResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
