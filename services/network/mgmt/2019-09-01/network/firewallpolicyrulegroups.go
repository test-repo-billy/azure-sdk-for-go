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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// FirewallPolicyRuleGroupsClient is the network Client
type FirewallPolicyRuleGroupsClient struct {
	BaseClient
}

// NewFirewallPolicyRuleGroupsClient creates an instance of the FirewallPolicyRuleGroupsClient client.
func NewFirewallPolicyRuleGroupsClient(subscriptionID string) FirewallPolicyRuleGroupsClient {
	return NewFirewallPolicyRuleGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFirewallPolicyRuleGroupsClientWithBaseURI creates an instance of the FirewallPolicyRuleGroupsClient client.
func NewFirewallPolicyRuleGroupsClientWithBaseURI(baseURI string, subscriptionID string) FirewallPolicyRuleGroupsClient {
	return FirewallPolicyRuleGroupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates the specified FirewallPolicyRuleGroup.
// Parameters:
// resourceGroupName - the name of the resource group.
// firewallPolicyName - the name of the Firewall Policy.
// ruleGroupName - the name of the FirewallPolicyRuleGroup.
// parameters - parameters supplied to the create or update FirewallPolicyRuleGroup operation.
func (client FirewallPolicyRuleGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, parameters FirewallPolicyRuleGroup) (result FirewallPolicyRuleGroupsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FirewallPolicyRuleGroupsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.FirewallPolicyRuleGroupProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.FirewallPolicyRuleGroupProperties.Priority", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.FirewallPolicyRuleGroupProperties.Priority", Name: validation.InclusiveMaximum, Rule: int64(65000), Chain: nil},
						{Target: "parameters.FirewallPolicyRuleGroupProperties.Priority", Name: validation.InclusiveMinimum, Rule: 100, Chain: nil},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("network.FirewallPolicyRuleGroupsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, firewallPolicyName, ruleGroupName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyRuleGroupsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyRuleGroupsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client FirewallPolicyRuleGroupsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string, parameters FirewallPolicyRuleGroup) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"firewallPolicyName": autorest.Encode("path", firewallPolicyName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"ruleGroupName":      autorest.Encode("path", ruleGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Etag = nil
	parameters.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups/{ruleGroupName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client FirewallPolicyRuleGroupsClient) CreateOrUpdateSender(req *http.Request) (future FirewallPolicyRuleGroupsCreateOrUpdateFuture, err error) {
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
func (client FirewallPolicyRuleGroupsClient) CreateOrUpdateResponder(resp *http.Response) (result FirewallPolicyRuleGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified FirewallPolicyRuleGroup.
// Parameters:
// resourceGroupName - the name of the resource group.
// firewallPolicyName - the name of the Firewall Policy.
// ruleGroupName - the name of the FirewallPolicyRuleGroup.
func (client FirewallPolicyRuleGroupsClient) Delete(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string) (result FirewallPolicyRuleGroupsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FirewallPolicyRuleGroupsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, firewallPolicyName, ruleGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyRuleGroupsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyRuleGroupsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client FirewallPolicyRuleGroupsClient) DeletePreparer(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"firewallPolicyName": autorest.Encode("path", firewallPolicyName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"ruleGroupName":      autorest.Encode("path", ruleGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups/{ruleGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client FirewallPolicyRuleGroupsClient) DeleteSender(req *http.Request) (future FirewallPolicyRuleGroupsDeleteFuture, err error) {
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
func (client FirewallPolicyRuleGroupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified FirewallPolicyRuleGroup.
// Parameters:
// resourceGroupName - the name of the resource group.
// firewallPolicyName - the name of the Firewall Policy.
// ruleGroupName - the name of the FirewallPolicyRuleGroup.
func (client FirewallPolicyRuleGroupsClient) Get(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string) (result FirewallPolicyRuleGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FirewallPolicyRuleGroupsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, firewallPolicyName, ruleGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyRuleGroupsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyRuleGroupsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyRuleGroupsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client FirewallPolicyRuleGroupsClient) GetPreparer(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"firewallPolicyName": autorest.Encode("path", firewallPolicyName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"ruleGroupName":      autorest.Encode("path", ruleGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups/{ruleGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client FirewallPolicyRuleGroupsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client FirewallPolicyRuleGroupsClient) GetResponder(resp *http.Response) (result FirewallPolicyRuleGroup, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all FirewallPolicyRuleGroups in a FirewallPolicy resource.
// Parameters:
// resourceGroupName - the name of the resource group.
// firewallPolicyName - the name of the Firewall Policy.
func (client FirewallPolicyRuleGroupsClient) List(ctx context.Context, resourceGroupName string, firewallPolicyName string) (result FirewallPolicyRuleGroupListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FirewallPolicyRuleGroupsClient.List")
		defer func() {
			sc := -1
			if result.fprglr.Response.Response != nil {
				sc = result.fprglr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, firewallPolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyRuleGroupsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.fprglr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyRuleGroupsClient", "List", resp, "Failure sending request")
		return
	}

	result.fprglr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyRuleGroupsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client FirewallPolicyRuleGroupsClient) ListPreparer(ctx context.Context, resourceGroupName string, firewallPolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"firewallPolicyName": autorest.Encode("path", firewallPolicyName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleGroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client FirewallPolicyRuleGroupsClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client FirewallPolicyRuleGroupsClient) ListResponder(resp *http.Response) (result FirewallPolicyRuleGroupListResult, err error) {
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
func (client FirewallPolicyRuleGroupsClient) listNextResults(ctx context.Context, lastResults FirewallPolicyRuleGroupListResult) (result FirewallPolicyRuleGroupListResult, err error) {
	req, err := lastResults.firewallPolicyRuleGroupListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.FirewallPolicyRuleGroupsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.FirewallPolicyRuleGroupsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.FirewallPolicyRuleGroupsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client FirewallPolicyRuleGroupsClient) ListComplete(ctx context.Context, resourceGroupName string, firewallPolicyName string) (result FirewallPolicyRuleGroupListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FirewallPolicyRuleGroupsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, firewallPolicyName)
	return
}
