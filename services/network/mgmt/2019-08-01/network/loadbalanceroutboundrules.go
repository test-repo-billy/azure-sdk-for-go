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

// LoadBalancerOutboundRulesClient is the network Client
type LoadBalancerOutboundRulesClient struct {
	BaseClient
}

// NewLoadBalancerOutboundRulesClient creates an instance of the LoadBalancerOutboundRulesClient client.
func NewLoadBalancerOutboundRulesClient(subscriptionID string) LoadBalancerOutboundRulesClient {
	return NewLoadBalancerOutboundRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLoadBalancerOutboundRulesClientWithBaseURI creates an instance of the LoadBalancerOutboundRulesClient client.
func NewLoadBalancerOutboundRulesClientWithBaseURI(baseURI string, subscriptionID string) LoadBalancerOutboundRulesClient {
	return LoadBalancerOutboundRulesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets the specified load balancer outbound rule.
// Parameters:
// resourceGroupName - the name of the resource group.
// loadBalancerName - the name of the load balancer.
// outboundRuleName - the name of the outbound rule.
func (client LoadBalancerOutboundRulesClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, outboundRuleName string) (result OutboundRule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoadBalancerOutboundRulesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, loadBalancerName, outboundRuleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerOutboundRulesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.LoadBalancerOutboundRulesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerOutboundRulesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client LoadBalancerOutboundRulesClient) GetPreparer(ctx context.Context, resourceGroupName string, loadBalancerName string, outboundRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"loadBalancerName":  autorest.Encode("path", loadBalancerName),
		"outboundRuleName":  autorest.Encode("path", outboundRuleName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/outboundRules/{outboundRuleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client LoadBalancerOutboundRulesClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LoadBalancerOutboundRulesClient) GetResponder(resp *http.Response) (result OutboundRule, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all the outbound rules in a load balancer.
// Parameters:
// resourceGroupName - the name of the resource group.
// loadBalancerName - the name of the load balancer.
func (client LoadBalancerOutboundRulesClient) List(ctx context.Context, resourceGroupName string, loadBalancerName string) (result LoadBalancerOutboundRuleListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoadBalancerOutboundRulesClient.List")
		defer func() {
			sc := -1
			if result.lborlr.Response.Response != nil {
				sc = result.lborlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, loadBalancerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerOutboundRulesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lborlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.LoadBalancerOutboundRulesClient", "List", resp, "Failure sending request")
		return
	}

	result.lborlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerOutboundRulesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client LoadBalancerOutboundRulesClient) ListPreparer(ctx context.Context, resourceGroupName string, loadBalancerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"loadBalancerName":  autorest.Encode("path", loadBalancerName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/outboundRules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client LoadBalancerOutboundRulesClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client LoadBalancerOutboundRulesClient) ListResponder(resp *http.Response) (result LoadBalancerOutboundRuleListResult, err error) {
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
func (client LoadBalancerOutboundRulesClient) listNextResults(ctx context.Context, lastResults LoadBalancerOutboundRuleListResult) (result LoadBalancerOutboundRuleListResult, err error) {
	req, err := lastResults.loadBalancerOutboundRuleListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.LoadBalancerOutboundRulesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.LoadBalancerOutboundRulesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.LoadBalancerOutboundRulesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client LoadBalancerOutboundRulesClient) ListComplete(ctx context.Context, resourceGroupName string, loadBalancerName string) (result LoadBalancerOutboundRuleListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LoadBalancerOutboundRulesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, loadBalancerName)
	return
}
