package netapp

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

// MountTargetsClient is the microsoft NetApp Azure Resource Provider specification
type MountTargetsClient struct {
	BaseClient
}

// NewMountTargetsClient creates an instance of the MountTargetsClient client.
func NewMountTargetsClient(subscriptionID string) MountTargetsClient {
	return NewMountTargetsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMountTargetsClientWithBaseURI creates an instance of the MountTargetsClient client.
func NewMountTargetsClientWithBaseURI(baseURI string, subscriptionID string) MountTargetsClient {
	return MountTargetsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List list all mount targets associated with the volume
// Parameters:
// resourceGroupName - the name of the resource group.
// accountName - the name of the NetApp account
// poolName - the name of the capacity pool
// volumeName - the name of the volume
func (client MountTargetsClient) List(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string) (result MountTargetList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MountTargetsClient.List")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("netapp.MountTargetsClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, resourceGroupName, accountName, poolName, volumeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.MountTargetsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "netapp.MountTargetsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netapp.MountTargetsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client MountTargetsClient) ListPreparer(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"poolName":          autorest.Encode("path", poolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"volumeName":        autorest.Encode("path", volumeName),
	}

	const APIVersion = "2019-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/mountTargets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client MountTargetsClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client MountTargetsClient) ListResponder(resp *http.Response) (result MountTargetList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
