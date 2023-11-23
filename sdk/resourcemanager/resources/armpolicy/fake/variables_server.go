//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
	"net/http"
	"net/url"
	"regexp"
)

// VariablesServer is a fake server for instances of the armpolicy.VariablesClient type.
type VariablesServer struct {
	// CreateOrUpdate is the fake for method VariablesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, variableName string, parameters armpolicy.Variable, options *armpolicy.VariablesClientCreateOrUpdateOptions) (resp azfake.Responder[armpolicy.VariablesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// CreateOrUpdateAtManagementGroup is the fake for method VariablesClient.CreateOrUpdateAtManagementGroup
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdateAtManagementGroup func(ctx context.Context, managementGroupID string, variableName string, parameters armpolicy.Variable, options *armpolicy.VariablesClientCreateOrUpdateAtManagementGroupOptions) (resp azfake.Responder[armpolicy.VariablesClientCreateOrUpdateAtManagementGroupResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method VariablesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, variableName string, options *armpolicy.VariablesClientDeleteOptions) (resp azfake.Responder[armpolicy.VariablesClientDeleteResponse], errResp azfake.ErrorResponder)

	// DeleteAtManagementGroup is the fake for method VariablesClient.DeleteAtManagementGroup
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	DeleteAtManagementGroup func(ctx context.Context, managementGroupID string, variableName string, options *armpolicy.VariablesClientDeleteAtManagementGroupOptions) (resp azfake.Responder[armpolicy.VariablesClientDeleteAtManagementGroupResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VariablesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, variableName string, options *armpolicy.VariablesClientGetOptions) (resp azfake.Responder[armpolicy.VariablesClientGetResponse], errResp azfake.ErrorResponder)

	// GetAtManagementGroup is the fake for method VariablesClient.GetAtManagementGroup
	// HTTP status codes to indicate success: http.StatusOK
	GetAtManagementGroup func(ctx context.Context, managementGroupID string, variableName string, options *armpolicy.VariablesClientGetAtManagementGroupOptions) (resp azfake.Responder[armpolicy.VariablesClientGetAtManagementGroupResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method VariablesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armpolicy.VariablesClientListOptions) (resp azfake.PagerResponder[armpolicy.VariablesClientListResponse])

	// NewListForManagementGroupPager is the fake for method VariablesClient.NewListForManagementGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListForManagementGroupPager func(managementGroupID string, options *armpolicy.VariablesClientListForManagementGroupOptions) (resp azfake.PagerResponder[armpolicy.VariablesClientListForManagementGroupResponse])
}

// NewVariablesServerTransport creates a new instance of VariablesServerTransport with the provided implementation.
// The returned VariablesServerTransport instance is connected to an instance of armpolicy.VariablesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVariablesServerTransport(srv *VariablesServer) *VariablesServerTransport {
	return &VariablesServerTransport{
		srv:                            srv,
		newListPager:                   newTracker[azfake.PagerResponder[armpolicy.VariablesClientListResponse]](),
		newListForManagementGroupPager: newTracker[azfake.PagerResponder[armpolicy.VariablesClientListForManagementGroupResponse]](),
	}
}

// VariablesServerTransport connects instances of armpolicy.VariablesClient to instances of VariablesServer.
// Don't use this type directly, use NewVariablesServerTransport instead.
type VariablesServerTransport struct {
	srv                            *VariablesServer
	newListPager                   *tracker[azfake.PagerResponder[armpolicy.VariablesClientListResponse]]
	newListForManagementGroupPager *tracker[azfake.PagerResponder[armpolicy.VariablesClientListForManagementGroupResponse]]
}

// Do implements the policy.Transporter interface for VariablesServerTransport.
func (v *VariablesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VariablesClient.CreateOrUpdate":
		resp, err = v.dispatchCreateOrUpdate(req)
	case "VariablesClient.CreateOrUpdateAtManagementGroup":
		resp, err = v.dispatchCreateOrUpdateAtManagementGroup(req)
	case "VariablesClient.Delete":
		resp, err = v.dispatchDelete(req)
	case "VariablesClient.DeleteAtManagementGroup":
		resp, err = v.dispatchDeleteAtManagementGroup(req)
	case "VariablesClient.Get":
		resp, err = v.dispatchGet(req)
	case "VariablesClient.GetAtManagementGroup":
		resp, err = v.dispatchGetAtManagementGroup(req)
	case "VariablesClient.NewListPager":
		resp, err = v.dispatchNewListPager(req)
	case "VariablesClient.NewListForManagementGroupPager":
		resp, err = v.dispatchNewListForManagementGroupPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VariablesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/variables/(?P<variableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armpolicy.Variable](req)
	if err != nil {
		return nil, err
	}
	variableNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("variableName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.CreateOrUpdate(req.Context(), variableNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Variable, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VariablesServerTransport) dispatchCreateOrUpdateAtManagementGroup(req *http.Request) (*http.Response, error) {
	if v.srv.CreateOrUpdateAtManagementGroup == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdateAtManagementGroup not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/variables/(?P<variableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armpolicy.Variable](req)
	if err != nil {
		return nil, err
	}
	managementGroupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
	if err != nil {
		return nil, err
	}
	variableNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("variableName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.CreateOrUpdateAtManagementGroup(req.Context(), managementGroupIDParam, variableNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Variable, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VariablesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if v.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/variables/(?P<variableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	variableNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("variableName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Delete(req.Context(), variableNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VariablesServerTransport) dispatchDeleteAtManagementGroup(req *http.Request) (*http.Response, error) {
	if v.srv.DeleteAtManagementGroup == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteAtManagementGroup not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/variables/(?P<variableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	managementGroupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
	if err != nil {
		return nil, err
	}
	variableNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("variableName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.DeleteAtManagementGroup(req.Context(), managementGroupIDParam, variableNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VariablesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/variables/(?P<variableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	variableNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("variableName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), variableNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Variable, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VariablesServerTransport) dispatchGetAtManagementGroup(req *http.Request) (*http.Response, error) {
	if v.srv.GetAtManagementGroup == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAtManagementGroup not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/variables/(?P<variableName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	managementGroupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
	if err != nil {
		return nil, err
	}
	variableNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("variableName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.GetAtManagementGroup(req.Context(), managementGroupIDParam, variableNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Variable, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VariablesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := v.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/variables`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := v.srv.NewListPager(nil)
		newListPager = &resp
		v.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armpolicy.VariablesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		v.newListPager.remove(req)
	}
	return resp, nil
}

func (v *VariablesServerTransport) dispatchNewListForManagementGroupPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListForManagementGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListForManagementGroupPager not implemented")}
	}
	newListForManagementGroupPager := v.newListForManagementGroupPager.get(req)
	if newListForManagementGroupPager == nil {
		const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/variables`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		managementGroupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListForManagementGroupPager(managementGroupIDParam, nil)
		newListForManagementGroupPager = &resp
		v.newListForManagementGroupPager.add(req, newListForManagementGroupPager)
		server.PagerResponderInjectNextLinks(newListForManagementGroupPager, req, func(page *armpolicy.VariablesClientListForManagementGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListForManagementGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListForManagementGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListForManagementGroupPager) {
		v.newListForManagementGroupPager.remove(req)
	}
	return resp, nil
}
