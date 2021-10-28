package desktopvirtualization

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
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

// DesktopsClient is the client for the Desktops methods of the Desktopvirtualization service.
type DesktopsClient struct {
	BaseClient
}

// NewDesktopsClient creates an instance of the DesktopsClient client.
func NewDesktopsClient(subscriptionID string) DesktopsClient {
	return NewDesktopsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDesktopsClientWithBaseURI creates an instance of the DesktopsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDesktopsClientWithBaseURI(baseURI string, subscriptionID string) DesktopsClient {
	return DesktopsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get a desktop.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// applicationGroupName - the name of the application group
// desktopName - the name of the desktop within the specified desktop group
func (client DesktopsClient) Get(ctx context.Context, resourceGroupName string, applicationGroupName string, desktopName string) (result Desktop, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DesktopsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: applicationGroupName,
			Constraints: []validation.Constraint{{Target: "applicationGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applicationGroupName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: desktopName,
			Constraints: []validation.Constraint{{Target: "desktopName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "desktopName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("desktopvirtualization.DesktopsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, applicationGroupName, desktopName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.DesktopsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "desktopvirtualization.DesktopsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.DesktopsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DesktopsClient) GetPreparer(ctx context.Context, resourceGroupName string, applicationGroupName string, desktopName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGroupName": autorest.Encode("path", applicationGroupName),
		"desktopName":          autorest.Encode("path", desktopName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-09-03-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups/{applicationGroupName}/desktops/{desktopName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DesktopsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DesktopsClient) GetResponder(resp *http.Response) (result Desktop, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list desktops.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// applicationGroupName - the name of the application group
func (client DesktopsClient) List(ctx context.Context, resourceGroupName string, applicationGroupName string) (result DesktopListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DesktopsClient.List")
		defer func() {
			sc := -1
			if result.dl.Response.Response != nil {
				sc = result.dl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: applicationGroupName,
			Constraints: []validation.Constraint{{Target: "applicationGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applicationGroupName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("desktopvirtualization.DesktopsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, applicationGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.DesktopsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.dl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "desktopvirtualization.DesktopsClient", "List", resp, "Failure sending request")
		return
	}

	result.dl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.DesktopsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.dl.hasNextLink() && result.dl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client DesktopsClient) ListPreparer(ctx context.Context, resourceGroupName string, applicationGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGroupName": autorest.Encode("path", applicationGroupName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-09-03-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups/{applicationGroupName}/desktops", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DesktopsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DesktopsClient) ListResponder(resp *http.Response) (result DesktopList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client DesktopsClient) listNextResults(ctx context.Context, lastResults DesktopList) (result DesktopList, err error) {
	req, err := lastResults.desktopListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "desktopvirtualization.DesktopsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "desktopvirtualization.DesktopsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.DesktopsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DesktopsClient) ListComplete(ctx context.Context, resourceGroupName string, applicationGroupName string) (result DesktopListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DesktopsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, applicationGroupName)
	return
}

// Update update a desktop.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// applicationGroupName - the name of the application group
// desktopName - the name of the desktop within the specified desktop group
// desktop - object containing Desktop definitions.
func (client DesktopsClient) Update(ctx context.Context, resourceGroupName string, applicationGroupName string, desktopName string, desktop *DesktopPatch) (result Desktop, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DesktopsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: applicationGroupName,
			Constraints: []validation.Constraint{{Target: "applicationGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applicationGroupName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: desktopName,
			Constraints: []validation.Constraint{{Target: "desktopName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "desktopName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("desktopvirtualization.DesktopsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, applicationGroupName, desktopName, desktop)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.DesktopsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "desktopvirtualization.DesktopsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.DesktopsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client DesktopsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, applicationGroupName string, desktopName string, desktop *DesktopPatch) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationGroupName": autorest.Encode("path", applicationGroupName),
		"desktopName":          autorest.Encode("path", desktopName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-09-03-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups/{applicationGroupName}/desktops/{desktopName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if desktop != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(desktop))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client DesktopsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client DesktopsClient) UpdateResponder(resp *http.Response) (result Desktop, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
