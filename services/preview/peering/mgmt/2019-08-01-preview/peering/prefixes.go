package peering

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

// PrefixesClient is the peering Client
type PrefixesClient struct {
	BaseClient
}

// NewPrefixesClient creates an instance of the PrefixesClient client.
func NewPrefixesClient(subscriptionID string) PrefixesClient {
	return NewPrefixesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrefixesClientWithBaseURI creates an instance of the PrefixesClient client.
func NewPrefixesClientWithBaseURI(baseURI string, subscriptionID string) PrefixesClient {
	return PrefixesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByPeeringService lists the peerings prefix in the resource group.
// Parameters:
// resourceGroupName - the resource group name.
// peeringServiceName - the peering service name.
func (client PrefixesClient) ListByPeeringService(ctx context.Context, resourceGroupName string, peeringServiceName string) (result ServicePrefixListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrefixesClient.ListByPeeringService")
		defer func() {
			sc := -1
			if result.splr.Response.Response != nil {
				sc = result.splr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByPeeringServiceNextResults
	req, err := client.ListByPeeringServicePreparer(ctx, resourceGroupName, peeringServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.PrefixesClient", "ListByPeeringService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByPeeringServiceSender(req)
	if err != nil {
		result.splr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "peering.PrefixesClient", "ListByPeeringService", resp, "Failure sending request")
		return
	}

	result.splr, err = client.ListByPeeringServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.PrefixesClient", "ListByPeeringService", resp, "Failure responding to request")
	}

	return
}

// ListByPeeringServicePreparer prepares the ListByPeeringService request.
func (client PrefixesClient) ListByPeeringServicePreparer(ctx context.Context, resourceGroupName string, peeringServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"peeringServiceName": autorest.Encode("path", peeringServiceName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peeringServices/{peeringServiceName}/prefixes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByPeeringServiceSender sends the ListByPeeringService request. The method will close the
// http.Response Body if it receives an error.
func (client PrefixesClient) ListByPeeringServiceSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByPeeringServiceResponder handles the response to the ListByPeeringService request. The method always
// closes the http.Response Body.
func (client PrefixesClient) ListByPeeringServiceResponder(resp *http.Response) (result ServicePrefixListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByPeeringServiceNextResults retrieves the next set of results, if any.
func (client PrefixesClient) listByPeeringServiceNextResults(ctx context.Context, lastResults ServicePrefixListResult) (result ServicePrefixListResult, err error) {
	req, err := lastResults.servicePrefixListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "peering.PrefixesClient", "listByPeeringServiceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByPeeringServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "peering.PrefixesClient", "listByPeeringServiceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByPeeringServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.PrefixesClient", "listByPeeringServiceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByPeeringServiceComplete enumerates all values, automatically crossing page boundaries as required.
func (client PrefixesClient) ListByPeeringServiceComplete(ctx context.Context, resourceGroupName string, peeringServiceName string) (result ServicePrefixListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PrefixesClient.ListByPeeringService")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByPeeringService(ctx, resourceGroupName, peeringServiceName)
	return
}
