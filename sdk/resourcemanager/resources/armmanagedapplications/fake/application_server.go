//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armmanagedapplications"
	"net/http"
)

// ApplicationServer is a fake server for instances of the armmanagedapplications.ApplicationClient type.
type ApplicationServer struct {
	// NewListOperationsPager is the fake for method ApplicationClient.NewListOperationsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListOperationsPager func(options *armmanagedapplications.ApplicationClientListOperationsOptions) (resp azfake.PagerResponder[armmanagedapplications.ApplicationClientListOperationsResponse])
}

// NewApplicationServerTransport creates a new instance of ApplicationServerTransport with the provided implementation.
// The returned ApplicationServerTransport instance is connected to an instance of armmanagedapplications.ApplicationClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewApplicationServerTransport(srv *ApplicationServer) *ApplicationServerTransport {
	return &ApplicationServerTransport{
		srv:                    srv,
		newListOperationsPager: newTracker[azfake.PagerResponder[armmanagedapplications.ApplicationClientListOperationsResponse]](),
	}
}

// ApplicationServerTransport connects instances of armmanagedapplications.ApplicationClient to instances of ApplicationServer.
// Don't use this type directly, use NewApplicationServerTransport instead.
type ApplicationServerTransport struct {
	srv                    *ApplicationServer
	newListOperationsPager *tracker[azfake.PagerResponder[armmanagedapplications.ApplicationClientListOperationsResponse]]
}

// Do implements the policy.Transporter interface for ApplicationServerTransport.
func (a *ApplicationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ApplicationClient.NewListOperationsPager":
		resp, err = a.dispatchNewListOperationsPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *ApplicationServerTransport) dispatchNewListOperationsPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListOperationsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListOperationsPager not implemented")}
	}
	newListOperationsPager := a.newListOperationsPager.get(req)
	if newListOperationsPager == nil {
		resp := a.srv.NewListOperationsPager(nil)
		newListOperationsPager = &resp
		a.newListOperationsPager.add(req, newListOperationsPager)
		server.PagerResponderInjectNextLinks(newListOperationsPager, req, func(page *armmanagedapplications.ApplicationClientListOperationsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListOperationsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListOperationsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListOperationsPager) {
		a.newListOperationsPager.remove(req)
	}
	return resp, nil
}
