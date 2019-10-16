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

// ServiceEndpointPolicyDefinitionsClient is the network Client
type ServiceEndpointPolicyDefinitionsClient struct {
	BaseClient
}

// NewServiceEndpointPolicyDefinitionsClient creates an instance of the ServiceEndpointPolicyDefinitionsClient client.
func NewServiceEndpointPolicyDefinitionsClient(subscriptionID string) ServiceEndpointPolicyDefinitionsClient {
	return NewServiceEndpointPolicyDefinitionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServiceEndpointPolicyDefinitionsClientWithBaseURI creates an instance of the
// ServiceEndpointPolicyDefinitionsClient client.
func NewServiceEndpointPolicyDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) ServiceEndpointPolicyDefinitionsClient {
	return ServiceEndpointPolicyDefinitionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a service endpoint policy definition in the specified service endpoint policy.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceEndpointPolicyName - the name of the service endpoint policy.
// serviceEndpointPolicyDefinitionName - the name of the service endpoint policy definition name.
// serviceEndpointPolicyDefinitions - parameters supplied to the create or update service endpoint policy
// operation.
func (client ServiceEndpointPolicyDefinitionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string, serviceEndpointPolicyDefinitions ServiceEndpointPolicyDefinition) (result ServiceEndpointPolicyDefinitionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceEndpointPolicyDefinitionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName, serviceEndpointPolicyDefinitions)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ServiceEndpointPolicyDefinitionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string, serviceEndpointPolicyDefinitions ServiceEndpointPolicyDefinition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"serviceEndpointPolicyDefinitionName": autorest.Encode("path", serviceEndpointPolicyDefinitionName),
		"serviceEndpointPolicyName":           autorest.Encode("path", serviceEndpointPolicyName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}", pathParameters),
		autorest.WithJSON(serviceEndpointPolicyDefinitions),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceEndpointPolicyDefinitionsClient) CreateOrUpdateSender(req *http.Request) (future ServiceEndpointPolicyDefinitionsCreateOrUpdateFuture, err error) {
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
func (client ServiceEndpointPolicyDefinitionsClient) CreateOrUpdateResponder(resp *http.Response) (result ServiceEndpointPolicyDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified ServiceEndpoint policy definitions.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceEndpointPolicyName - the name of the Service Endpoint Policy.
// serviceEndpointPolicyDefinitionName - the name of the service endpoint policy definition.
func (client ServiceEndpointPolicyDefinitionsClient) Delete(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string) (result ServiceEndpointPolicyDefinitionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceEndpointPolicyDefinitionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ServiceEndpointPolicyDefinitionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"serviceEndpointPolicyDefinitionName": autorest.Encode("path", serviceEndpointPolicyDefinitionName),
		"serviceEndpointPolicyName":           autorest.Encode("path", serviceEndpointPolicyName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceEndpointPolicyDefinitionsClient) DeleteSender(req *http.Request) (future ServiceEndpointPolicyDefinitionsDeleteFuture, err error) {
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
func (client ServiceEndpointPolicyDefinitionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the specified service endpoint policy definitions from service endpoint policy.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceEndpointPolicyName - the name of the service endpoint policy name.
// serviceEndpointPolicyDefinitionName - the name of the service endpoint policy definition name.
func (client ServiceEndpointPolicyDefinitionsClient) Get(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string) (result ServiceEndpointPolicyDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceEndpointPolicyDefinitionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ServiceEndpointPolicyDefinitionsClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"serviceEndpointPolicyDefinitionName": autorest.Encode("path", serviceEndpointPolicyDefinitionName),
		"serviceEndpointPolicyName":           autorest.Encode("path", serviceEndpointPolicyName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceEndpointPolicyDefinitionsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServiceEndpointPolicyDefinitionsClient) GetResponder(resp *http.Response) (result ServiceEndpointPolicyDefinition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup gets all service endpoint policy definitions in a service end point policy.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceEndpointPolicyName - the name of the service endpoint policy name.
func (client ServiceEndpointPolicyDefinitionsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string) (result ServiceEndpointPolicyDefinitionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceEndpointPolicyDefinitionsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.sepdlr.Response.Response != nil {
				sc = result.sepdlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, serviceEndpointPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.sepdlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.sepdlr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ServiceEndpointPolicyDefinitionsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"serviceEndpointPolicyName": autorest.Encode("path", serviceEndpointPolicyName),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceEndpointPolicyDefinitionsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ServiceEndpointPolicyDefinitionsClient) ListByResourceGroupResponder(resp *http.Response) (result ServiceEndpointPolicyDefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client ServiceEndpointPolicyDefinitionsClient) listByResourceGroupNextResults(ctx context.Context, lastResults ServiceEndpointPolicyDefinitionListResult) (result ServiceEndpointPolicyDefinitionListResult, err error) {
	req, err := lastResults.serviceEndpointPolicyDefinitionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ServiceEndpointPolicyDefinitionsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ServiceEndpointPolicyDefinitionsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string) (result ServiceEndpointPolicyDefinitionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServiceEndpointPolicyDefinitionsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, serviceEndpointPolicyName)
	return
}
