package databoxedge

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

// UsersClient is the client for the Users methods of the Databoxedge service.
type UsersClient struct {
	BaseClient
}

// NewUsersClient creates an instance of the UsersClient client.
func NewUsersClient(subscriptionID string) UsersClient {
	return NewUsersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUsersClientWithBaseURI creates an instance of the UsersClient client.
func NewUsersClientWithBaseURI(baseURI string, subscriptionID string) UsersClient {
	return UsersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a new user or updates an existing user's information on a data box edge/gateway device.
// Parameters:
// deviceName - the device name.
// name - the user name.
// userParameter - the user details.
// resourceGroupName - the resource group name.
func (client UsersClient) CreateOrUpdate(ctx context.Context, deviceName string, name string, userParameter User, resourceGroupName string) (result UsersCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: userParameter,
			Constraints: []validation.Constraint{{Target: "userParameter.UserProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "userParameter.UserProperties.EncryptedPassword", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "userParameter.UserProperties.EncryptedPassword.Value", Name: validation.Null, Rule: true, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewError("databoxedge.UsersClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, deviceName, name, userParameter, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.UsersClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.UsersClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client UsersClient) CreateOrUpdatePreparer(ctx context.Context, deviceName string, name string, userParameter User, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/users/{name}", pathParameters),
		autorest.WithJSON(userParameter),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) CreateOrUpdateSender(req *http.Request) (future UsersCreateOrUpdateFuture, err error) {
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
func (client UsersClient) CreateOrUpdateResponder(resp *http.Response) (result User, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the user on a databox edge/gateway device.
// Parameters:
// deviceName - the device name.
// name - the user name.
// resourceGroupName - the resource group name.
func (client UsersClient) Delete(ctx context.Context, deviceName string, name string, resourceGroupName string) (result UsersDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, deviceName, name, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.UsersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.UsersClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client UsersClient) DeletePreparer(ctx context.Context, deviceName string, name string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/users/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) DeleteSender(req *http.Request) (future UsersDeleteFuture, err error) {
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
func (client UsersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the properties of the specified user.
// Parameters:
// deviceName - the device name.
// name - the user name.
// resourceGroupName - the resource group name.
func (client UsersClient) Get(ctx context.Context, deviceName string, name string, resourceGroupName string) (result User, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, deviceName, name, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.UsersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databoxedge.UsersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.UsersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client UsersClient) GetPreparer(ctx context.Context, deviceName string, name string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/users/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client UsersClient) GetResponder(resp *http.Response) (result User, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDataBoxEdgeDevice gets all the users registered on a data box edge/gateway device.
// Parameters:
// deviceName - the device name.
// resourceGroupName - the resource group name.
func (client UsersClient) ListByDataBoxEdgeDevice(ctx context.Context, deviceName string, resourceGroupName string) (result UserListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.ListByDataBoxEdgeDevice")
		defer func() {
			sc := -1
			if result.ul.Response.Response != nil {
				sc = result.ul.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByDataBoxEdgeDeviceNextResults
	req, err := client.ListByDataBoxEdgeDevicePreparer(ctx, deviceName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.UsersClient", "ListByDataBoxEdgeDevice", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDataBoxEdgeDeviceSender(req)
	if err != nil {
		result.ul.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "databoxedge.UsersClient", "ListByDataBoxEdgeDevice", resp, "Failure sending request")
		return
	}

	result.ul, err = client.ListByDataBoxEdgeDeviceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.UsersClient", "ListByDataBoxEdgeDevice", resp, "Failure responding to request")
	}

	return
}

// ListByDataBoxEdgeDevicePreparer prepares the ListByDataBoxEdgeDevice request.
func (client UsersClient) ListByDataBoxEdgeDevicePreparer(ctx context.Context, deviceName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName":        autorest.Encode("path", deviceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/users", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDataBoxEdgeDeviceSender sends the ListByDataBoxEdgeDevice request. The method will close the
// http.Response Body if it receives an error.
func (client UsersClient) ListByDataBoxEdgeDeviceSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByDataBoxEdgeDeviceResponder handles the response to the ListByDataBoxEdgeDevice request. The method always
// closes the http.Response Body.
func (client UsersClient) ListByDataBoxEdgeDeviceResponder(resp *http.Response) (result UserList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByDataBoxEdgeDeviceNextResults retrieves the next set of results, if any.
func (client UsersClient) listByDataBoxEdgeDeviceNextResults(ctx context.Context, lastResults UserList) (result UserList, err error) {
	req, err := lastResults.userListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "databoxedge.UsersClient", "listByDataBoxEdgeDeviceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByDataBoxEdgeDeviceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "databoxedge.UsersClient", "listByDataBoxEdgeDeviceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByDataBoxEdgeDeviceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databoxedge.UsersClient", "listByDataBoxEdgeDeviceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByDataBoxEdgeDeviceComplete enumerates all values, automatically crossing page boundaries as required.
func (client UsersClient) ListByDataBoxEdgeDeviceComplete(ctx context.Context, deviceName string, resourceGroupName string) (result UserListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsersClient.ListByDataBoxEdgeDevice")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByDataBoxEdgeDevice(ctx, deviceName, resourceGroupName)
	return
}
