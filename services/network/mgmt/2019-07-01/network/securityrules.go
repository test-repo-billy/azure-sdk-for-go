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

// SecurityRulesClient is the network Client
type SecurityRulesClient struct {
	BaseClient
}

// NewSecurityRulesClient creates an instance of the SecurityRulesClient client.
func NewSecurityRulesClient(subscriptionID string) SecurityRulesClient {
	return NewSecurityRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSecurityRulesClientWithBaseURI creates an instance of the SecurityRulesClient client.
func NewSecurityRulesClientWithBaseURI(baseURI string, subscriptionID string) SecurityRulesClient {
	return SecurityRulesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a security rule in the specified network security group.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkSecurityGroupName - the name of the network security group.
// securityRuleName - the name of the security rule.
// securityRuleParameters - parameters supplied to the create or update network security rule operation.
func (client SecurityRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, securityRuleParameters SecurityRule) (result SecurityRulesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityRulesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, networkSecurityGroupName, securityRuleName, securityRuleParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityRulesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityRulesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SecurityRulesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, securityRuleParameters SecurityRule) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkSecurityGroupName": autorest.Encode("path", networkSecurityGroupName),
		"resourceGroupName":        autorest.Encode("path", resourceGroupName),
		"securityRuleName":         autorest.Encode("path", securityRuleName),
		"subscriptionId":           autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}", pathParameters),
		autorest.WithJSON(securityRuleParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityRulesClient) CreateOrUpdateSender(req *http.Request) (future SecurityRulesCreateOrUpdateFuture, err error) {
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
func (client SecurityRulesClient) CreateOrUpdateResponder(resp *http.Response) (result SecurityRule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified network security rule.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkSecurityGroupName - the name of the network security group.
// securityRuleName - the name of the security rule.
func (client SecurityRulesClient) Delete(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string) (result SecurityRulesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityRulesClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, networkSecurityGroupName, securityRuleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityRulesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityRulesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SecurityRulesClient) DeletePreparer(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkSecurityGroupName": autorest.Encode("path", networkSecurityGroupName),
		"resourceGroupName":        autorest.Encode("path", resourceGroupName),
		"securityRuleName":         autorest.Encode("path", securityRuleName),
		"subscriptionId":           autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityRulesClient) DeleteSender(req *http.Request) (future SecurityRulesDeleteFuture, err error) {
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
func (client SecurityRulesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the specified network security rule.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkSecurityGroupName - the name of the network security group.
// securityRuleName - the name of the security rule.
func (client SecurityRulesClient) Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string) (result SecurityRule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityRulesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, networkSecurityGroupName, securityRuleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityRulesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.SecurityRulesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityRulesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SecurityRulesClient) GetPreparer(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkSecurityGroupName": autorest.Encode("path", networkSecurityGroupName),
		"resourceGroupName":        autorest.Encode("path", resourceGroupName),
		"securityRuleName":         autorest.Encode("path", securityRuleName),
		"subscriptionId":           autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityRulesClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SecurityRulesClient) GetResponder(resp *http.Response) (result SecurityRule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all security rules in a network security group.
// Parameters:
// resourceGroupName - the name of the resource group.
// networkSecurityGroupName - the name of the network security group.
func (client SecurityRulesClient) List(ctx context.Context, resourceGroupName string, networkSecurityGroupName string) (result SecurityRuleListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityRulesClient.List")
		defer func() {
			sc := -1
			if result.srlr.Response.Response != nil {
				sc = result.srlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, networkSecurityGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityRulesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.srlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.SecurityRulesClient", "List", resp, "Failure sending request")
		return
	}

	result.srlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityRulesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client SecurityRulesClient) ListPreparer(ctx context.Context, resourceGroupName string, networkSecurityGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkSecurityGroupName": autorest.Encode("path", networkSecurityGroupName),
		"resourceGroupName":        autorest.Encode("path", resourceGroupName),
		"subscriptionId":           autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SecurityRulesClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SecurityRulesClient) ListResponder(resp *http.Response) (result SecurityRuleListResult, err error) {
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
func (client SecurityRulesClient) listNextResults(ctx context.Context, lastResults SecurityRuleListResult) (result SecurityRuleListResult, err error) {
	req, err := lastResults.securityRuleListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.SecurityRulesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.SecurityRulesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.SecurityRulesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client SecurityRulesClient) ListComplete(ctx context.Context, resourceGroupName string, networkSecurityGroupName string) (result SecurityRuleListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SecurityRulesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, networkSecurityGroupName)
	return
}
