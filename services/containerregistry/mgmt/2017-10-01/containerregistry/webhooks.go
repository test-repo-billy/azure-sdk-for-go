package containerregistry

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// WebhooksClient is the client for the Webhooks methods of the Containerregistry service.
type WebhooksClient struct {
	BaseClient
}

// NewWebhooksClient creates an instance of the WebhooksClient client.
func NewWebhooksClient(subscriptionID string) WebhooksClient {
	return NewWebhooksClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWebhooksClientWithBaseURI creates an instance of the WebhooksClient client.
func NewWebhooksClientWithBaseURI(baseURI string, subscriptionID string) WebhooksClient {
	return WebhooksClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a webhook for a container registry with the specified parameters.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// webhookName - the name of the webhook.
// webhookCreateParameters - the parameters for creating a webhook.
func (client WebhooksClient) Create(ctx context.Context, resourceGroupName string, registryName string, webhookName string, webhookCreateParameters WebhookCreateParameters) (result WebhooksCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WebhooksClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: webhookName,
			Constraints: []validation.Constraint{{Target: "webhookName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "webhookName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "webhookName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: webhookCreateParameters,
			Constraints: []validation.Constraint{{Target: "webhookCreateParameters.Location", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "webhookCreateParameters.WebhookPropertiesCreateParameters", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "webhookCreateParameters.WebhookPropertiesCreateParameters.ServiceURI", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "webhookCreateParameters.WebhookPropertiesCreateParameters.Actions", Name: validation.Null, Rule: true, Chain: nil},
					}}}}}); err != nil {
		return result, validation.NewError("containerregistry.WebhooksClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, registryName, webhookName, webhookCreateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client WebhooksClient) CreatePreparer(ctx context.Context, resourceGroupName string, registryName string, webhookName string, webhookCreateParameters WebhookCreateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"webhookName":       autorest.Encode("path", webhookName),
	}

	const APIVersion = "2017-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}", pathParameters),
		autorest.WithJSON(webhookCreateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client WebhooksClient) CreateSender(req *http.Request) (future WebhooksCreateFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client WebhooksClient) CreateResponder(resp *http.Response) (result Webhook, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a webhook from a container registry.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// webhookName - the name of the webhook.
func (client WebhooksClient) Delete(ctx context.Context, resourceGroupName string, registryName string, webhookName string) (result WebhooksDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WebhooksClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: webhookName,
			Constraints: []validation.Constraint{{Target: "webhookName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "webhookName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "webhookName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.WebhooksClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, registryName, webhookName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client WebhooksClient) DeletePreparer(ctx context.Context, resourceGroupName string, registryName string, webhookName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"webhookName":       autorest.Encode("path", webhookName),
	}

	const APIVersion = "2017-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client WebhooksClient) DeleteSender(req *http.Request) (future WebhooksDeleteFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client WebhooksClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the properties of the specified webhook.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// webhookName - the name of the webhook.
func (client WebhooksClient) Get(ctx context.Context, resourceGroupName string, registryName string, webhookName string) (result Webhook, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WebhooksClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: webhookName,
			Constraints: []validation.Constraint{{Target: "webhookName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "webhookName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "webhookName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.WebhooksClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, registryName, webhookName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client WebhooksClient) GetPreparer(ctx context.Context, resourceGroupName string, registryName string, webhookName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"webhookName":       autorest.Encode("path", webhookName),
	}

	const APIVersion = "2017-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WebhooksClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WebhooksClient) GetResponder(resp *http.Response) (result Webhook, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetCallbackConfig gets the configuration of service URI and custom headers for the webhook.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// webhookName - the name of the webhook.
func (client WebhooksClient) GetCallbackConfig(ctx context.Context, resourceGroupName string, registryName string, webhookName string) (result CallbackConfig, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WebhooksClient.GetCallbackConfig")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: webhookName,
			Constraints: []validation.Constraint{{Target: "webhookName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "webhookName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "webhookName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.WebhooksClient", "GetCallbackConfig", err.Error())
	}

	req, err := client.GetCallbackConfigPreparer(ctx, resourceGroupName, registryName, webhookName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "GetCallbackConfig", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetCallbackConfigSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "GetCallbackConfig", resp, "Failure sending request")
		return
	}

	result, err = client.GetCallbackConfigResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "GetCallbackConfig", resp, "Failure responding to request")
	}

	return
}

// GetCallbackConfigPreparer prepares the GetCallbackConfig request.
func (client WebhooksClient) GetCallbackConfigPreparer(ctx context.Context, resourceGroupName string, registryName string, webhookName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"webhookName":       autorest.Encode("path", webhookName),
	}

	const APIVersion = "2017-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}/getCallbackConfig", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetCallbackConfigSender sends the GetCallbackConfig request. The method will close the
// http.Response Body if it receives an error.
func (client WebhooksClient) GetCallbackConfigSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetCallbackConfigResponder handles the response to the GetCallbackConfig request. The method always
// closes the http.Response Body.
func (client WebhooksClient) GetCallbackConfigResponder(resp *http.Response) (result CallbackConfig, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all the webhooks for the specified container registry.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
func (client WebhooksClient) List(ctx context.Context, resourceGroupName string, registryName string) (result WebhookListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WebhooksClient.List")
		defer func() {
			sc := -1
			if result.wlr.Response.Response != nil {
				sc = result.wlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.WebhooksClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, registryName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.wlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "List", resp, "Failure sending request")
		return
	}

	result.wlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client WebhooksClient) ListPreparer(ctx context.Context, resourceGroupName string, registryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client WebhooksClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client WebhooksClient) ListResponder(resp *http.Response) (result WebhookListResult, err error) {
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
func (client WebhooksClient) listNextResults(ctx context.Context, lastResults WebhookListResult) (result WebhookListResult, err error) {
	req, err := lastResults.webhookListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client WebhooksClient) ListComplete(ctx context.Context, resourceGroupName string, registryName string) (result WebhookListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WebhooksClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, registryName)
	return
}

// ListEvents lists recent events for the specified webhook.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// webhookName - the name of the webhook.
func (client WebhooksClient) ListEvents(ctx context.Context, resourceGroupName string, registryName string, webhookName string) (result EventListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WebhooksClient.ListEvents")
		defer func() {
			sc := -1
			if result.elr.Response.Response != nil {
				sc = result.elr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: webhookName,
			Constraints: []validation.Constraint{{Target: "webhookName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "webhookName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "webhookName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.WebhooksClient", "ListEvents", err.Error())
	}

	result.fn = client.listEventsNextResults
	req, err := client.ListEventsPreparer(ctx, resourceGroupName, registryName, webhookName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "ListEvents", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListEventsSender(req)
	if err != nil {
		result.elr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "ListEvents", resp, "Failure sending request")
		return
	}

	result.elr, err = client.ListEventsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "ListEvents", resp, "Failure responding to request")
	}

	return
}

// ListEventsPreparer prepares the ListEvents request.
func (client WebhooksClient) ListEventsPreparer(ctx context.Context, resourceGroupName string, registryName string, webhookName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"webhookName":       autorest.Encode("path", webhookName),
	}

	const APIVersion = "2017-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}/listEvents", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListEventsSender sends the ListEvents request. The method will close the
// http.Response Body if it receives an error.
func (client WebhooksClient) ListEventsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListEventsResponder handles the response to the ListEvents request. The method always
// closes the http.Response Body.
func (client WebhooksClient) ListEventsResponder(resp *http.Response) (result EventListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listEventsNextResults retrieves the next set of results, if any.
func (client WebhooksClient) listEventsNextResults(ctx context.Context, lastResults EventListResult) (result EventListResult, err error) {
	req, err := lastResults.eventListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "listEventsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListEventsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "listEventsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListEventsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "listEventsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListEventsComplete enumerates all values, automatically crossing page boundaries as required.
func (client WebhooksClient) ListEventsComplete(ctx context.Context, resourceGroupName string, registryName string, webhookName string) (result EventListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WebhooksClient.ListEvents")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListEvents(ctx, resourceGroupName, registryName, webhookName)
	return
}

// Ping triggers a ping event to be sent to the webhook.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// webhookName - the name of the webhook.
func (client WebhooksClient) Ping(ctx context.Context, resourceGroupName string, registryName string, webhookName string) (result EventInfo, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WebhooksClient.Ping")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: webhookName,
			Constraints: []validation.Constraint{{Target: "webhookName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "webhookName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "webhookName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.WebhooksClient", "Ping", err.Error())
	}

	req, err := client.PingPreparer(ctx, resourceGroupName, registryName, webhookName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "Ping", nil, "Failure preparing request")
		return
	}

	resp, err := client.PingSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "Ping", resp, "Failure sending request")
		return
	}

	result, err = client.PingResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "Ping", resp, "Failure responding to request")
	}

	return
}

// PingPreparer prepares the Ping request.
func (client WebhooksClient) PingPreparer(ctx context.Context, resourceGroupName string, registryName string, webhookName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"webhookName":       autorest.Encode("path", webhookName),
	}

	const APIVersion = "2017-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}/ping", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PingSender sends the Ping request. The method will close the
// http.Response Body if it receives an error.
func (client WebhooksClient) PingSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// PingResponder handles the response to the Ping request. The method always
// closes the http.Response Body.
func (client WebhooksClient) PingResponder(resp *http.Response) (result EventInfo, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates a webhook with the specified parameters.
// Parameters:
// resourceGroupName - the name of the resource group to which the container registry belongs.
// registryName - the name of the container registry.
// webhookName - the name of the webhook.
// webhookUpdateParameters - the parameters for updating a webhook.
func (client WebhooksClient) Update(ctx context.Context, resourceGroupName string, registryName string, webhookName string, webhookUpdateParameters WebhookUpdateParameters) (result WebhooksUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WebhooksClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: registryName,
			Constraints: []validation.Constraint{{Target: "registryName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "registryName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "registryName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}},
		{TargetValue: webhookName,
			Constraints: []validation.Constraint{{Target: "webhookName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "webhookName", Name: validation.MinLength, Rule: 5, Chain: nil},
				{Target: "webhookName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("containerregistry.WebhooksClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, registryName, webhookName, webhookUpdateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.WebhooksClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client WebhooksClient) UpdatePreparer(ctx context.Context, resourceGroupName string, registryName string, webhookName string, webhookUpdateParameters WebhookUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"webhookName":       autorest.Encode("path", webhookName),
	}

	const APIVersion = "2017-10-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/webhooks/{webhookName}", pathParameters),
		autorest.WithJSON(webhookUpdateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client WebhooksClient) UpdateSender(req *http.Request) (future WebhooksUpdateFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client WebhooksClient) UpdateResponder(resp *http.Response) (result Webhook, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
