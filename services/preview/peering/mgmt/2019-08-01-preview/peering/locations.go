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

// LocationsClient is the peering Client
type LocationsClient struct {
	BaseClient
}

// NewLocationsClient creates an instance of the LocationsClient client.
func NewLocationsClient(subscriptionID string) LocationsClient {
	return NewLocationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLocationsClientWithBaseURI creates an instance of the LocationsClient client.
func NewLocationsClientWithBaseURI(baseURI string, subscriptionID string) LocationsClient {
	return LocationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists all of the available peering locations for the specified kind of peering.
// Parameters:
// kind - the kind of the peering.
// directPeeringType - the type of direct peering.
func (client LocationsClient) List(ctx context.Context, kind string, directPeeringType string) (result LocationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationsClient.List")
		defer func() {
			sc := -1
			if result.llr.Response.Response != nil {
				sc = result.llr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, kind, directPeeringType)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.LocationsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.llr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "peering.LocationsClient", "List", resp, "Failure sending request")
		return
	}

	result.llr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.LocationsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client LocationsClient) ListPreparer(ctx context.Context, kind string, directPeeringType string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"kind":        autorest.Encode("query", kind),
	}
	if len(string(directPeeringType)) > 0 {
		queryParameters["directPeeringType"] = autorest.Encode("query", directPeeringType)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Peering/peeringLocations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client LocationsClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client LocationsClient) ListResponder(resp *http.Response) (result LocationListResult, err error) {
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
func (client LocationsClient) listNextResults(ctx context.Context, lastResults LocationListResult) (result LocationListResult, err error) {
	req, err := lastResults.locationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "peering.LocationsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "peering.LocationsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.LocationsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client LocationsClient) ListComplete(ctx context.Context, kind string, directPeeringType string) (result LocationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, kind, directPeeringType)
	return
}
