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

// VpnConnectionsClient is the network Client
type VpnConnectionsClient struct {
	BaseClient
}

// NewVpnConnectionsClient creates an instance of the VpnConnectionsClient client.
func NewVpnConnectionsClient(subscriptionID string) VpnConnectionsClient {
	return NewVpnConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVpnConnectionsClientWithBaseURI creates an instance of the VpnConnectionsClient client.
func NewVpnConnectionsClientWithBaseURI(baseURI string, subscriptionID string) VpnConnectionsClient {
	return VpnConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a vpn connection to a scalable vpn gateway if it doesn't exist else updates the existing
// connection.
// Parameters:
// resourceGroupName - the resource group name of the VpnGateway.
// gatewayName - the name of the gateway.
// connectionName - the name of the connection.
// vpnConnectionParameters - parameters supplied to create or Update a VPN Connection.
func (client VpnConnectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters VpnConnection) (result VpnConnectionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnConnectionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, gatewayName, connectionName, vpnConnectionParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VpnConnectionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters VpnConnection) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":    autorest.Encode("path", connectionName),
		"gatewayName":       autorest.Encode("path", gatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	vpnConnectionParameters.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}", pathParameters),
		autorest.WithJSON(vpnConnectionParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VpnConnectionsClient) CreateOrUpdateSender(req *http.Request) (future VpnConnectionsCreateOrUpdateFuture, err error) {
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
func (client VpnConnectionsClient) CreateOrUpdateResponder(resp *http.Response) (result VpnConnection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a vpn connection.
// Parameters:
// resourceGroupName - the resource group name of the VpnGateway.
// gatewayName - the name of the gateway.
// connectionName - the name of the connection.
func (client VpnConnectionsClient) Delete(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string) (result VpnConnectionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnConnectionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, gatewayName, connectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VpnConnectionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":    autorest.Encode("path", connectionName),
		"gatewayName":       autorest.Encode("path", gatewayName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VpnConnectionsClient) DeleteSender(req *http.Request) (future VpnConnectionsDeleteFuture, err error) {
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
func (client VpnConnectionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get retrieves the details of a vpn connection.
// Parameters:
// resourceGroupName - the resource group name of the VpnGateway.
// gatewayName - the name of the gateway.
// connectionName - the name of the vpn connection.
func (client VpnConnectionsClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string) (result VpnConnection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnConnectionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, gatewayName, connectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VpnConnectionsClient) GetPreparer(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionName":    autorest.Encode("path", connectionName),
		"gatewayName":       autorest.Encode("path", gatewayName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VpnConnectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VpnConnectionsClient) GetResponder(resp *http.Response) (result VpnConnection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByVpnGateway retrieves all vpn connections for a particular virtual wan vpn gateway.
// Parameters:
// resourceGroupName - the resource group name of the VpnGateway.
// gatewayName - the name of the gateway.
func (client VpnConnectionsClient) ListByVpnGateway(ctx context.Context, resourceGroupName string, gatewayName string) (result ListVpnConnectionsResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnConnectionsClient.ListByVpnGateway")
		defer func() {
			sc := -1
			if result.lvcr.Response.Response != nil {
				sc = result.lvcr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByVpnGatewayNextResults
	req, err := client.ListByVpnGatewayPreparer(ctx, resourceGroupName, gatewayName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "ListByVpnGateway", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByVpnGatewaySender(req)
	if err != nil {
		result.lvcr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "ListByVpnGateway", resp, "Failure sending request")
		return
	}

	result.lvcr, err = client.ListByVpnGatewayResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "ListByVpnGateway", resp, "Failure responding to request")
	}

	return
}

// ListByVpnGatewayPreparer prepares the ListByVpnGateway request.
func (client VpnConnectionsClient) ListByVpnGatewayPreparer(ctx context.Context, resourceGroupName string, gatewayName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayName":       autorest.Encode("path", gatewayName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByVpnGatewaySender sends the ListByVpnGateway request. The method will close the
// http.Response Body if it receives an error.
func (client VpnConnectionsClient) ListByVpnGatewaySender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByVpnGatewayResponder handles the response to the ListByVpnGateway request. The method always
// closes the http.Response Body.
func (client VpnConnectionsClient) ListByVpnGatewayResponder(resp *http.Response) (result ListVpnConnectionsResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByVpnGatewayNextResults retrieves the next set of results, if any.
func (client VpnConnectionsClient) listByVpnGatewayNextResults(ctx context.Context, lastResults ListVpnConnectionsResult) (result ListVpnConnectionsResult, err error) {
	req, err := lastResults.listVpnConnectionsResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "listByVpnGatewayNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByVpnGatewaySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "listByVpnGatewayNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByVpnGatewayResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VpnConnectionsClient", "listByVpnGatewayNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByVpnGatewayComplete enumerates all values, automatically crossing page boundaries as required.
func (client VpnConnectionsClient) ListByVpnGatewayComplete(ctx context.Context, resourceGroupName string, gatewayName string) (result ListVpnConnectionsResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VpnConnectionsClient.ListByVpnGateway")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByVpnGateway(ctx, resourceGroupName, gatewayName)
	return
}
