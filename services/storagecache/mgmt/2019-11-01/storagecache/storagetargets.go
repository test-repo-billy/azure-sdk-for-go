package storagecache

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

// StorageTargetsClient is the a Storage Cache provides scalable caching service for NAS clients, serving data from
// either NFSv3 or Blob at-rest storage (referred to as "Storage Targets"). These operations allow you to manage
// Caches.
type StorageTargetsClient struct {
	BaseClient
}

// NewStorageTargetsClient creates an instance of the StorageTargetsClient client.
func NewStorageTargetsClient(subscriptionID string) StorageTargetsClient {
	return NewStorageTargetsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewStorageTargetsClientWithBaseURI creates an instance of the StorageTargetsClient client.
func NewStorageTargetsClientWithBaseURI(baseURI string, subscriptionID string) StorageTargetsClient {
	return StorageTargetsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a Storage Target. This operation is allowed at any time, but if the Cache is down or
// unhealthy, the actual creation/modification of the Storage Target may be delayed until the Cache is healthy again.
// Parameters:
// resourceGroupName - target resource group.
// cacheName - name of Cache.
// storageTargetName - name of the Storage Target.
// storagetarget - object containing the definition of a Storage Target.
func (client StorageTargetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, storagetarget *StorageTarget) (result StorageTargetsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageTargetsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: cacheName,
			Constraints: []validation.Constraint{{Target: "cacheName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,31}$`, Chain: nil}}},
		{TargetValue: storageTargetName,
			Constraints: []validation.Constraint{{Target: "storageTargetName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,31}$`, Chain: nil}}},
		{TargetValue: storagetarget,
			Constraints: []validation.Constraint{{Target: "storagetarget", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "storagetarget.StorageTargetProperties", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "storagetarget.StorageTargetProperties.Nfs3", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "storagetarget.StorageTargetProperties.Nfs3.Target", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "storagetarget.StorageTargetProperties.Nfs3.Target", Name: validation.Pattern, Rule: `^[-.0-9a-zA-Z]+$`, Chain: nil}}},
						}},
					}},
				}}}}}); err != nil {
		return result, validation.NewError("storagecache.StorageTargetsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, cacheName, storageTargetName, storagetarget)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client StorageTargetsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, storagetarget *StorageTarget) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"storageTargetName": autorest.Encode("path", storageTargetName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	storagetarget.Name = nil
	storagetarget.ID = nil
	storagetarget.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if storagetarget != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(storagetarget))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client StorageTargetsClient) CreateOrUpdateSender(req *http.Request) (future StorageTargetsCreateOrUpdateFuture, err error) {
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
func (client StorageTargetsClient) CreateOrUpdateResponder(resp *http.Response) (result StorageTarget, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete removes a Storage Target from a Cache. This operation is allowed at any time, but if the Cache is down or
// unhealthy, the actual removal of the Storage Target may be delayed until the Cache is healthy again. Note that if
// the Cache has data to flush to the Storage Target, the data will be flushed before the Storage Target will be
// deleted.
// Parameters:
// resourceGroupName - target resource group.
// cacheName - name of Cache.
// storageTargetName - name of Storage Target.
func (client StorageTargetsClient) Delete(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string) (result StorageTargetsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageTargetsClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: cacheName,
			Constraints: []validation.Constraint{{Target: "cacheName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,31}$`, Chain: nil}}},
		{TargetValue: storageTargetName,
			Constraints: []validation.Constraint{{Target: "storageTargetName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,31}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagecache.StorageTargetsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, cacheName, storageTargetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client StorageTargetsClient) DeletePreparer(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"storageTargetName": autorest.Encode("path", storageTargetName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client StorageTargetsClient) DeleteSender(req *http.Request) (future StorageTargetsDeleteFuture, err error) {
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
func (client StorageTargetsClient) DeleteResponder(resp *http.Response) (result SetObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get returns a Storage Target from a Cache.
// Parameters:
// resourceGroupName - target resource group.
// cacheName - name of Cache.
// storageTargetName - name of the Storage Target.
func (client StorageTargetsClient) Get(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string) (result StorageTarget, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageTargetsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: cacheName,
			Constraints: []validation.Constraint{{Target: "cacheName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,31}$`, Chain: nil}}},
		{TargetValue: storageTargetName,
			Constraints: []validation.Constraint{{Target: "storageTargetName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,31}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagecache.StorageTargetsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, cacheName, storageTargetName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client StorageTargetsClient) GetPreparer(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"storageTargetName": autorest.Encode("path", storageTargetName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client StorageTargetsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client StorageTargetsClient) GetResponder(resp *http.Response) (result StorageTarget, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByCache returns a list of Storage Targets for the specified Cache.
// Parameters:
// resourceGroupName - target resource group.
// cacheName - name of Cache.
func (client StorageTargetsClient) ListByCache(ctx context.Context, resourceGroupName string, cacheName string) (result StorageTargetsResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageTargetsClient.ListByCache")
		defer func() {
			sc := -1
			if result.str.Response.Response != nil {
				sc = result.str.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: cacheName,
			Constraints: []validation.Constraint{{Target: "cacheName", Name: validation.Pattern, Rule: `^[-0-9a-zA-Z_]{1,31}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("storagecache.StorageTargetsClient", "ListByCache", err.Error())
	}

	result.fn = client.listByCacheNextResults
	req, err := client.ListByCachePreparer(ctx, resourceGroupName, cacheName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "ListByCache", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByCacheSender(req)
	if err != nil {
		result.str.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "ListByCache", resp, "Failure sending request")
		return
	}

	result.str, err = client.ListByCacheResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "ListByCache", resp, "Failure responding to request")
	}

	return
}

// ListByCachePreparer prepares the ListByCache request.
func (client StorageTargetsClient) ListByCachePreparer(ctx context.Context, resourceGroupName string, cacheName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"cacheName":         autorest.Encode("path", cacheName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByCacheSender sends the ListByCache request. The method will close the
// http.Response Body if it receives an error.
func (client StorageTargetsClient) ListByCacheSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByCacheResponder handles the response to the ListByCache request. The method always
// closes the http.Response Body.
func (client StorageTargetsClient) ListByCacheResponder(resp *http.Response) (result StorageTargetsResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByCacheNextResults retrieves the next set of results, if any.
func (client StorageTargetsClient) listByCacheNextResults(ctx context.Context, lastResults StorageTargetsResult) (result StorageTargetsResult, err error) {
	req, err := lastResults.storageTargetsResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "listByCacheNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByCacheSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "listByCacheNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByCacheResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "storagecache.StorageTargetsClient", "listByCacheNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByCacheComplete enumerates all values, automatically crossing page boundaries as required.
func (client StorageTargetsClient) ListByCacheComplete(ctx context.Context, resourceGroupName string, cacheName string) (result StorageTargetsResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/StorageTargetsClient.ListByCache")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByCache(ctx, resourceGroupName, cacheName)
	return
}
