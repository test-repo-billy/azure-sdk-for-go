package sql

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

// DatabaseConnectionPoliciesClient is the the Azure SQL Database management API provides a RESTful set of web services
// that interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve,
// update, and delete databases.
type DatabaseConnectionPoliciesClient struct {
	BaseClient
}

// NewDatabaseConnectionPoliciesClient creates an instance of the DatabaseConnectionPoliciesClient client.
func NewDatabaseConnectionPoliciesClient(subscriptionID string) DatabaseConnectionPoliciesClient {
	return NewDatabaseConnectionPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDatabaseConnectionPoliciesClientWithBaseURI creates an instance of the DatabaseConnectionPoliciesClient client.
func NewDatabaseConnectionPoliciesClientWithBaseURI(baseURI string, subscriptionID string) DatabaseConnectionPoliciesClient {
	return DatabaseConnectionPoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a database's connection policy, which is used with table auditing. Table auditing
// is deprecated, use blob auditing instead.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database for which the connection policy will be defined.
// parameters - the database connection policy.
func (client DatabaseConnectionPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters DatabaseConnectionPolicy) (result DatabaseConnectionPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabaseConnectionPoliciesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, databaseName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseConnectionPoliciesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.DatabaseConnectionPoliciesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseConnectionPoliciesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DatabaseConnectionPoliciesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters DatabaseConnectionPolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionPolicyName": autorest.Encode("path", "default"),
		"databaseName":         autorest.Encode("path", databaseName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"serverName":           autorest.Encode("path", serverName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Kind = nil
	parameters.Location = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/connectionPolicies/{connectionPolicyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DatabaseConnectionPoliciesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DatabaseConnectionPoliciesClient) CreateOrUpdateResponder(resp *http.Response) (result DatabaseConnectionPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a database's connection policy, which is used with table auditing. Table auditing is deprecated, use blob
// auditing instead.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database for which the connection policy is defined.
func (client DatabaseConnectionPoliciesClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result DatabaseConnectionPolicy, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DatabaseConnectionPoliciesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseConnectionPoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.DatabaseConnectionPoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseConnectionPoliciesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DatabaseConnectionPoliciesClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionPolicyName": autorest.Encode("path", "default"),
		"databaseName":         autorest.Encode("path", databaseName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"serverName":           autorest.Encode("path", serverName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/connectionPolicies/{connectionPolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DatabaseConnectionPoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DatabaseConnectionPoliciesClient) GetResponder(resp *http.Response) (result DatabaseConnectionPolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
