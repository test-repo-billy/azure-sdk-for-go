package servicebusapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2018-01-01-preview/servicebus"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result servicebus.OperationListResultPage, err error)
}

var _ OperationsClientAPI = (*servicebus.OperationsClient)(nil)

// NamespacesClientAPI contains the set of methods on the NamespacesClient type.
type NamespacesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, parameters servicebus.SBNamespace) (result servicebus.NamespacesCreateOrUpdateFuture, err error)
	CreateOrUpdateIPFilterRule(ctx context.Context, resourceGroupName string, namespaceName string, IPFilterRuleName string, parameters servicebus.IPFilterRule) (result servicebus.IPFilterRule, err error)
	CreateOrUpdateNetworkRuleSet(ctx context.Context, resourceGroupName string, namespaceName string, parameters servicebus.NetworkRuleSet) (result servicebus.NetworkRuleSet, err error)
	CreateOrUpdateVirtualNetworkRule(ctx context.Context, resourceGroupName string, namespaceName string, virtualNetworkRuleName string, parameters servicebus.VirtualNetworkRule) (result servicebus.VirtualNetworkRule, err error)
	Delete(ctx context.Context, resourceGroupName string, namespaceName string) (result servicebus.NamespacesDeleteFuture, err error)
	DeleteIPFilterRule(ctx context.Context, resourceGroupName string, namespaceName string, IPFilterRuleName string) (result autorest.Response, err error)
	DeleteVirtualNetworkRule(ctx context.Context, resourceGroupName string, namespaceName string, virtualNetworkRuleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, namespaceName string) (result servicebus.SBNamespace, err error)
	GetIPFilterRule(ctx context.Context, resourceGroupName string, namespaceName string, IPFilterRuleName string) (result servicebus.IPFilterRule, err error)
	GetNetworkRuleSet(ctx context.Context, resourceGroupName string, namespaceName string) (result servicebus.NetworkRuleSet, err error)
	GetVirtualNetworkRule(ctx context.Context, resourceGroupName string, namespaceName string, virtualNetworkRuleName string) (result servicebus.VirtualNetworkRule, err error)
	List(ctx context.Context) (result servicebus.SBNamespaceListResultPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result servicebus.SBNamespaceListResultPage, err error)
	ListIPFilterRules(ctx context.Context, resourceGroupName string, namespaceName string) (result servicebus.IPFilterRuleListResultPage, err error)
	ListVirtualNetworkRules(ctx context.Context, resourceGroupName string, namespaceName string) (result servicebus.VirtualNetworkRuleListResultPage, err error)
	Update(ctx context.Context, resourceGroupName string, namespaceName string, parameters servicebus.SBNamespaceUpdateParameters) (result servicebus.SBNamespace, err error)
}

var _ NamespacesClientAPI = (*servicebus.NamespacesClient)(nil)
