package machinelearningservicesapi

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
	"github.com/Azure/azure-sdk-for-go/services/machinelearningservices/mgmt/2019-06-01/machinelearningservices"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result machinelearningservices.OperationListResult, err error)
}

var _ OperationsClientAPI = (*machinelearningservices.OperationsClient)(nil)

// WorkspacesClientAPI contains the set of methods on the WorkspacesClient type.
type WorkspacesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, parameters machinelearningservices.Workspace) (result machinelearningservices.Workspace, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string) (result machinelearningservices.Workspace, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, skiptoken string) (result machinelearningservices.WorkspaceListResultPage, err error)
	ListBySubscription(ctx context.Context, skiptoken string) (result machinelearningservices.WorkspaceListResultPage, err error)
	ListKeys(ctx context.Context, resourceGroupName string, workspaceName string) (result machinelearningservices.ListWorkspaceKeysResult, err error)
	ResyncKeys(ctx context.Context, resourceGroupName string, workspaceName string) (result autorest.Response, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, parameters machinelearningservices.WorkspaceUpdateParameters) (result machinelearningservices.Workspace, err error)
}

var _ WorkspacesClientAPI = (*machinelearningservices.WorkspacesClient)(nil)

// UsagesClientAPI contains the set of methods on the UsagesClient type.
type UsagesClientAPI interface {
	List(ctx context.Context, location string, expandChildren string) (result machinelearningservices.ListUsagesResultPage, err error)
}

var _ UsagesClientAPI = (*machinelearningservices.UsagesClient)(nil)

// VirtualMachineSizesClientAPI contains the set of methods on the VirtualMachineSizesClient type.
type VirtualMachineSizesClientAPI interface {
	List(ctx context.Context, location string) (result machinelearningservices.VirtualMachineSizeListResult, err error)
}

var _ VirtualMachineSizesClientAPI = (*machinelearningservices.VirtualMachineSizesClient)(nil)

// QuotasClientAPI contains the set of methods on the QuotasClient type.
type QuotasClientAPI interface {
	List(ctx context.Context, location string) (result machinelearningservices.ListWorkspaceQuotasPage, err error)
	Update(ctx context.Context, location string, parameters machinelearningservices.QuotaUpdateParameters) (result machinelearningservices.UpdateWorkspaceQuotasResult, err error)
}

var _ QuotasClientAPI = (*machinelearningservices.QuotasClient)(nil)

// MachineLearningComputeClientAPI contains the set of methods on the MachineLearningComputeClient type.
type MachineLearningComputeClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters machinelearningservices.ComputeResource) (result machinelearningservices.MachineLearningComputeCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, underlyingResourceAction machinelearningservices.UnderlyingResourceAction) (result machinelearningservices.MachineLearningComputeDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (result machinelearningservices.ComputeResource, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, skiptoken string) (result machinelearningservices.PaginatedComputeResourcesListPage, err error)
	ListKeys(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (result machinelearningservices.ComputeSecretsModel, err error)
	ListNodes(ctx context.Context, resourceGroupName string, workspaceName string, computeName string) (result machinelearningservices.AmlComputeNodesInformation, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, computeName string, parameters machinelearningservices.ClusterUpdateParameters) (result machinelearningservices.MachineLearningComputeUpdateFuture, err error)
}

var _ MachineLearningComputeClientAPI = (*machinelearningservices.MachineLearningComputeClient)(nil)
