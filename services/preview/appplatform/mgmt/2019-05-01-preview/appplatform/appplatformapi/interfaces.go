package appplatformapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/appplatform/mgmt/2019-05-01-preview/appplatform"
	"github.com/Azure/go-autorest/autorest"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	CheckNameAvailability(ctx context.Context, location string, availabilityParameters appplatform.NameAvailabilityParameters) (result appplatform.NameAvailability, err error)
}

var _ BaseClientAPI = (*appplatform.BaseClient)(nil)

// ServicesClientAPI contains the set of methods on the ServicesClient type.
type ServicesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, resource *appplatform.ServiceResource) (result appplatform.ServicesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.ServicesDeleteFuture, err error)
	DisableTestEndpoint(ctx context.Context, resourceGroupName string, serviceName string) (result autorest.Response, err error)
	EnableTestEndpoint(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.TestKeys, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.ServiceResource, err error)
	List(ctx context.Context, resourceGroupName string) (result appplatform.ServiceResourceListPage, err error)
	ListBySubscription(ctx context.Context) (result appplatform.ServiceResourceListPage, err error)
	ListTestKeys(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.TestKeys, err error)
	RegenerateTestKey(ctx context.Context, resourceGroupName string, serviceName string, regenerateTestKeyRequest *appplatform.RegenerateTestKeyRequestPayload) (result appplatform.TestKeys, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, resource *appplatform.ServiceResource) (result appplatform.ServicesUpdateFuture, err error)
}

var _ ServicesClientAPI = (*appplatform.ServicesClient)(nil)

// AppsClientAPI contains the set of methods on the AppsClient type.
type AppsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, appName string, appResource *appplatform.AppResource) (result appplatform.AppResource, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, appName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, appName string, syncStatus string) (result appplatform.AppResource, err error)
	GetResourceUploadURL(ctx context.Context, resourceGroupName string, serviceName string, appName string) (result appplatform.ResourceUploadDefinition, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.AppResourceCollectionPage, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, appName string, appResource *appplatform.AppResource) (result appplatform.AppResource, err error)
}

var _ AppsClientAPI = (*appplatform.AppsClient)(nil)

// BindingsClientAPI contains the set of methods on the BindingsClient type.
type BindingsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, appName string, bindingName string, bindingResource *appplatform.BindingResource) (result appplatform.BindingResource, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, appName string, bindingName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, appName string, bindingName string) (result appplatform.BindingResource, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string, appName string) (result appplatform.BindingResourceCollectionPage, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, appName string, bindingName string, bindingResource *appplatform.BindingResource) (result appplatform.BindingResource, err error)
}

var _ BindingsClientAPI = (*appplatform.BindingsClient)(nil)

// DeploymentsClientAPI contains the set of methods on the DeploymentsClient type.
type DeploymentsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string, deploymentResource *appplatform.DeploymentResource) (result appplatform.DeploymentsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result appplatform.DeploymentResource, err error)
	GetLogFileURL(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result appplatform.LogFileURLResponse, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string, appName string, version []string) (result appplatform.DeploymentResourceCollectionPage, err error)
	ListClusterAllDeployments(ctx context.Context, resourceGroupName string, serviceName string, version []string) (result appplatform.DeploymentResourceCollectionPage, err error)
	Restart(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result appplatform.DeploymentsRestartFuture, err error)
	Start(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result appplatform.DeploymentsStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result appplatform.DeploymentsStopFuture, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string, deploymentResource *appplatform.DeploymentResource) (result appplatform.DeploymentsUpdateFuture, err error)
}

var _ DeploymentsClientAPI = (*appplatform.DeploymentsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result appplatform.AvailableOperationsPage, err error)
}

var _ OperationsClientAPI = (*appplatform.OperationsClient)(nil)
