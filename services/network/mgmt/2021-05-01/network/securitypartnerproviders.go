package network

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
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

// SecurityPartnerProvidersClient is the network Client
type SecurityPartnerProvidersClient struct {
	BaseClient
}

// NewSecurityPartnerProvidersClient creates an instance of the SecurityPartnerProvidersClient client.
func NewSecurityPartnerProvidersClient(subscriptionID string) SecurityPartnerProvidersClient {
	return NewSecurityPartnerProvidersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSecurityPartnerProvidersClientWithBaseURI creates an instance of the SecurityPartnerProvidersClient client using
// a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewSecurityPartnerProvidersClientWithBaseURI(baseURI string, subscriptionID string) SecurityPartnerProvidersClient {
	return SecurityPartnerProvidersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates the specified Security Partner Provider.
// Parameters:
// resourceGroupName - the name of the resource group.
// securityPartnerProviderName - the name of the Security Partner Provider.
// parameters - parameters supplied to the create or update Security Partner Provider operation.
func (client SecurityPartnerProvidersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, parameters SecurityPartnerProvider) (result SecurityPartnerProvidersCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityPartnerProvidersClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, securityPartnerProviderName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityPartnerProvidersClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityPartnerProvidersClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SecurityPartnerProvidersClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, parameters SecurityPartnerProvider) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"securityPartnerProviderName": autorest.Encode("path", securityPartnerProviderName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Etag = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/securityPartnerProviders/{securityPartnerProviderName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityPartnerProvidersClient) CreateOrUpdateSender(req *http.Request) (future SecurityPartnerProvidersCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SecurityPartnerProvidersClient) CreateOrUpdateResponder(resp *http.Response) (result SecurityPartnerProvider, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified Security Partner Provider.
// Parameters:
// resourceGroupName - the name of the resource group.
// securityPartnerProviderName - the name of the Security Partner Provider.
func (client SecurityPartnerProvidersClient) Delete(ctx context.Context, resourceGroupName string, securityPartnerProviderName string) (result SecurityPartnerProvidersDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityPartnerProvidersClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, securityPartnerProviderName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityPartnerProvidersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityPartnerProvidersClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SecurityPartnerProvidersClient) DeletePreparer(ctx context.Context, resourceGroupName string, securityPartnerProviderName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"securityPartnerProviderName": autorest.Encode("path", securityPartnerProviderName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/securityPartnerProviders/{securityPartnerProviderName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityPartnerProvidersClient) DeleteSender(req *http.Request) (future SecurityPartnerProvidersDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SecurityPartnerProvidersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified Security Partner Provider.
// Parameters:
// resourceGroupName - the name of the resource group.
// securityPartnerProviderName - the name of the Security Partner Provider.
func (client SecurityPartnerProvidersClient) Get(ctx context.Context, resourceGroupName string, securityPartnerProviderName string) (result SecurityPartnerProvider, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityPartnerProvidersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, securityPartnerProviderName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityPartnerProvidersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.SecurityPartnerProvidersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityPartnerProvidersClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SecurityPartnerProvidersClient) GetPreparer(ctx context.Context, resourceGroupName string, securityPartnerProviderName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"securityPartnerProviderName": autorest.Encode("path", securityPartnerProviderName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/securityPartnerProviders/{securityPartnerProviderName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityPartnerProvidersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SecurityPartnerProvidersClient) GetResponder(resp *http.Response) (result SecurityPartnerProvider, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all the Security Partner Providers in a subscription.
func (client SecurityPartnerProvidersClient) List(ctx context.Context) (result SecurityPartnerProviderListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityPartnerProvidersClient.List")
		defer func() {
			sc := -1
			if result.spplr.Response.Response != nil {
				sc = result.spplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityPartnerProvidersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.spplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.SecurityPartnerProvidersClient", "List", resp, "Failure sending request")
		return
	}

	result.spplr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityPartnerProvidersClient", "List", resp, "Failure responding to request")
		return
	}
	if result.spplr.hasNextLink() && result.spplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client SecurityPartnerProvidersClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/securityPartnerProviders", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityPartnerProvidersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SecurityPartnerProvidersClient) ListResponder(resp *http.Response) (result SecurityPartnerProviderListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client SecurityPartnerProvidersClient) listNextResults(ctx context.Context, lastResults SecurityPartnerProviderListResult) (result SecurityPartnerProviderListResult, err error) {
	req, err := lastResults.securityPartnerProviderListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.SecurityPartnerProvidersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.SecurityPartnerProvidersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityPartnerProvidersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client SecurityPartnerProvidersClient) ListComplete(ctx context.Context) (result SecurityPartnerProviderListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityPartnerProvidersClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup lists all Security Partner Providers in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client SecurityPartnerProvidersClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result SecurityPartnerProviderListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityPartnerProvidersClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.spplr.Response.Response != nil {
				sc = result.spplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityPartnerProvidersClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.spplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.SecurityPartnerProvidersClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.spplr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityPartnerProvidersClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.spplr.hasNextLink() && result.spplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client SecurityPartnerProvidersClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/securityPartnerProviders", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityPartnerProvidersClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client SecurityPartnerProvidersClient) ListByResourceGroupResponder(resp *http.Response) (result SecurityPartnerProviderListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client SecurityPartnerProvidersClient) listByResourceGroupNextResults(ctx context.Context, lastResults SecurityPartnerProviderListResult) (result SecurityPartnerProviderListResult, err error) {
	req, err := lastResults.securityPartnerProviderListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.SecurityPartnerProvidersClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.SecurityPartnerProvidersClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityPartnerProvidersClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client SecurityPartnerProvidersClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result SecurityPartnerProviderListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityPartnerProvidersClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// UpdateTags updates tags of a Security Partner Provider resource.
// Parameters:
// resourceGroupName - the name of the resource group.
// securityPartnerProviderName - the name of the Security Partner Provider.
// parameters - parameters supplied to update Security Partner Provider tags.
func (client SecurityPartnerProvidersClient) UpdateTags(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, parameters TagsObject) (result SecurityPartnerProvider, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityPartnerProvidersClient.UpdateTags")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateTagsPreparer(ctx, resourceGroupName, securityPartnerProviderName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityPartnerProvidersClient", "UpdateTags", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateTagsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.SecurityPartnerProvidersClient", "UpdateTags", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateTagsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityPartnerProvidersClient", "UpdateTags", resp, "Failure responding to request")
		return
	}

	return
}

// UpdateTagsPreparer prepares the UpdateTags request.
func (client SecurityPartnerProvidersClient) UpdateTagsPreparer(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, parameters TagsObject) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":           autorest.Encode("path", resourceGroupName),
		"securityPartnerProviderName": autorest.Encode("path", securityPartnerProviderName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/securityPartnerProviders/{securityPartnerProviderName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateTagsSender sends the UpdateTags request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityPartnerProvidersClient) UpdateTagsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateTagsResponder handles the response to the UpdateTags request. The method always
// closes the http.Response Body.
func (client SecurityPartnerProvidersClient) UpdateTagsResponder(resp *http.Response) (result SecurityPartnerProvider, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
