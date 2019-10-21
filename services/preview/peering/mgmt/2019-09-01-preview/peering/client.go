// Package peering implements the Azure ARM Peering service API version 2019-09-01-preview.
//
// Peering Client
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

const (
	// DefaultBaseURI is the default URI used for the service Peering
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Peering.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client.
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// CheckServiceProviderAvailability checks if the peering service provider is present within 1000 miles of customer's
// location
// Parameters:
// checkServiceProviderAvailabilityInput - the CheckServiceProviderAvailabilityInput indicating customer
// location and service provider.
func (client BaseClient) CheckServiceProviderAvailability(ctx context.Context, checkServiceProviderAvailabilityInput CheckServiceProviderAvailabilityInput) (result String, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.CheckServiceProviderAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CheckServiceProviderAvailabilityPreparer(ctx, checkServiceProviderAvailabilityInput)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.BaseClient", "CheckServiceProviderAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckServiceProviderAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "peering.BaseClient", "CheckServiceProviderAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckServiceProviderAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "peering.BaseClient", "CheckServiceProviderAvailability", resp, "Failure responding to request")
	}

	return
}

// CheckServiceProviderAvailabilityPreparer prepares the CheckServiceProviderAvailability request.
func (client BaseClient) CheckServiceProviderAvailabilityPreparer(ctx context.Context, checkServiceProviderAvailabilityInput CheckServiceProviderAvailabilityInput) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Peering/CheckServiceProviderAvailability", pathParameters),
		autorest.WithJSON(checkServiceProviderAvailabilityInput),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckServiceProviderAvailabilitySender sends the CheckServiceProviderAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) CheckServiceProviderAvailabilitySender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CheckServiceProviderAvailabilityResponder handles the response to the CheckServiceProviderAvailability request. The method always
// closes the http.Response Body.
func (client BaseClient) CheckServiceProviderAvailabilityResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
