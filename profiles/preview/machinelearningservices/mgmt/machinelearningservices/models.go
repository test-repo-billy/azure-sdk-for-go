// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package machinelearningservices

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/machinelearningservices/mgmt/2019-06-01/machinelearningservices"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AllocationState = original.AllocationState

const (
	Resizing AllocationState = original.Resizing
	Steady   AllocationState = original.Steady
)

type ComputeType = original.ComputeType

const (
	ComputeTypeAKS               ComputeType = original.ComputeTypeAKS
	ComputeTypeAmlCompute        ComputeType = original.ComputeTypeAmlCompute
	ComputeTypeDatabricks        ComputeType = original.ComputeTypeDatabricks
	ComputeTypeDataFactory       ComputeType = original.ComputeTypeDataFactory
	ComputeTypeDataLakeAnalytics ComputeType = original.ComputeTypeDataLakeAnalytics
	ComputeTypeHDInsight         ComputeType = original.ComputeTypeHDInsight
	ComputeTypeVirtualMachine    ComputeType = original.ComputeTypeVirtualMachine
)

type ComputeTypeBasicCompute = original.ComputeTypeBasicCompute

const (
	ComputeTypeAKS1               ComputeTypeBasicCompute = original.ComputeTypeAKS1
	ComputeTypeAmlCompute1        ComputeTypeBasicCompute = original.ComputeTypeAmlCompute1
	ComputeTypeCompute            ComputeTypeBasicCompute = original.ComputeTypeCompute
	ComputeTypeDatabricks1        ComputeTypeBasicCompute = original.ComputeTypeDatabricks1
	ComputeTypeDataFactory1       ComputeTypeBasicCompute = original.ComputeTypeDataFactory1
	ComputeTypeDataLakeAnalytics1 ComputeTypeBasicCompute = original.ComputeTypeDataLakeAnalytics1
	ComputeTypeHDInsight1         ComputeTypeBasicCompute = original.ComputeTypeHDInsight1
	ComputeTypeVirtualMachine1    ComputeTypeBasicCompute = original.ComputeTypeVirtualMachine1
)

type ComputeTypeBasicComputeNodesInformation = original.ComputeTypeBasicComputeNodesInformation

const (
	ComputeTypeBasicComputeNodesInformationComputeTypeAmlCompute              ComputeTypeBasicComputeNodesInformation = original.ComputeTypeBasicComputeNodesInformationComputeTypeAmlCompute
	ComputeTypeBasicComputeNodesInformationComputeTypeComputeNodesInformation ComputeTypeBasicComputeNodesInformation = original.ComputeTypeBasicComputeNodesInformationComputeTypeComputeNodesInformation
)

type ComputeTypeBasicComputeSecrets = original.ComputeTypeBasicComputeSecrets

const (
	ComputeTypeBasicComputeSecretsComputeTypeAKS            ComputeTypeBasicComputeSecrets = original.ComputeTypeBasicComputeSecretsComputeTypeAKS
	ComputeTypeBasicComputeSecretsComputeTypeComputeSecrets ComputeTypeBasicComputeSecrets = original.ComputeTypeBasicComputeSecretsComputeTypeComputeSecrets
	ComputeTypeBasicComputeSecretsComputeTypeDatabricks     ComputeTypeBasicComputeSecrets = original.ComputeTypeBasicComputeSecretsComputeTypeDatabricks
	ComputeTypeBasicComputeSecretsComputeTypeVirtualMachine ComputeTypeBasicComputeSecrets = original.ComputeTypeBasicComputeSecretsComputeTypeVirtualMachine
)

type NodeState = original.NodeState

const (
	Idle      NodeState = original.Idle
	Leaving   NodeState = original.Leaving
	Preempted NodeState = original.Preempted
	Preparing NodeState = original.Preparing
	Running   NodeState = original.Running
	Unusable  NodeState = original.Unusable
)

type ProvisioningState = original.ProvisioningState

const (
	Canceled  ProvisioningState = original.Canceled
	Creating  ProvisioningState = original.Creating
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
	Unknown   ProvisioningState = original.Unknown
	Updating  ProvisioningState = original.Updating
)

type QuotaUnit = original.QuotaUnit

const (
	Count QuotaUnit = original.Count
)

type RemoteLoginPortPublicAccess = original.RemoteLoginPortPublicAccess

const (
	Disabled     RemoteLoginPortPublicAccess = original.Disabled
	Enabled      RemoteLoginPortPublicAccess = original.Enabled
	NotSpecified RemoteLoginPortPublicAccess = original.NotSpecified
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	SystemAssigned ResourceIdentityType = original.SystemAssigned
)

type Status = original.Status

const (
	Failure                              Status = original.Failure
	InvalidQuotaBelowClusterMinimum      Status = original.InvalidQuotaBelowClusterMinimum
	InvalidQuotaExceedsSubscriptionLimit Status = original.InvalidQuotaExceedsSubscriptionLimit
	InvalidVMFamilyName                  Status = original.InvalidVMFamilyName
	Success                              Status = original.Success
	Undefined                            Status = original.Undefined
)

type Status1 = original.Status1

const (
	Status1Disabled Status1 = original.Status1Disabled
	Status1Enabled  Status1 = original.Status1Enabled
)

type UnderlyingResourceAction = original.UnderlyingResourceAction

const (
	Delete UnderlyingResourceAction = original.Delete
	Detach UnderlyingResourceAction = original.Detach
)

type UsageUnit = original.UsageUnit

const (
	UsageUnitCount UsageUnit = original.UsageUnitCount
)

type VMPriority = original.VMPriority

const (
	Dedicated   VMPriority = original.Dedicated
	LowPriority VMPriority = original.LowPriority
)

type AKS = original.AKS
type AKSProperties = original.AKSProperties
type AksComputeSecrets = original.AksComputeSecrets
type AksNetworkingConfiguration = original.AksNetworkingConfiguration
type AmlCompute = original.AmlCompute
type AmlComputeNodeInformation = original.AmlComputeNodeInformation
type AmlComputeNodesInformation = original.AmlComputeNodesInformation
type AmlComputeProperties = original.AmlComputeProperties
type BaseClient = original.BaseClient
type BasicCompute = original.BasicCompute
type BasicComputeNodesInformation = original.BasicComputeNodesInformation
type BasicComputeSecrets = original.BasicComputeSecrets
type ClusterUpdateParameters = original.ClusterUpdateParameters
type ClusterUpdateProperties = original.ClusterUpdateProperties
type Compute = original.Compute
type ComputeNodesInformation = original.ComputeNodesInformation
type ComputeResource = original.ComputeResource
type ComputeSecrets = original.ComputeSecrets
type ComputeSecretsModel = original.ComputeSecretsModel
type DataFactory = original.DataFactory
type DataLakeAnalytics = original.DataLakeAnalytics
type DataLakeAnalyticsProperties = original.DataLakeAnalyticsProperties
type Databricks = original.Databricks
type DatabricksComputeSecrets = original.DatabricksComputeSecrets
type DatabricksProperties = original.DatabricksProperties
type Error = original.Error
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type HDInsight = original.HDInsight
type HDInsightProperties = original.HDInsightProperties
type Identity = original.Identity
type ListUsagesResult = original.ListUsagesResult
type ListUsagesResultIterator = original.ListUsagesResultIterator
type ListUsagesResultPage = original.ListUsagesResultPage
type ListWorkspaceKeysResult = original.ListWorkspaceKeysResult
type ListWorkspaceQuotas = original.ListWorkspaceQuotas
type ListWorkspaceQuotasIterator = original.ListWorkspaceQuotasIterator
type ListWorkspaceQuotasPage = original.ListWorkspaceQuotasPage
type MachineLearningComputeClient = original.MachineLearningComputeClient
type MachineLearningComputeCreateOrUpdateFuture = original.MachineLearningComputeCreateOrUpdateFuture
type MachineLearningComputeDeleteFuture = original.MachineLearningComputeDeleteFuture
type MachineLearningComputeUpdateFuture = original.MachineLearningComputeUpdateFuture
type NodeStateCounts = original.NodeStateCounts
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationsClient = original.OperationsClient
type PaginatedComputeResourcesList = original.PaginatedComputeResourcesList
type PaginatedComputeResourcesListIterator = original.PaginatedComputeResourcesListIterator
type PaginatedComputeResourcesListPage = original.PaginatedComputeResourcesListPage
type Password = original.Password
type QuotaBaseProperties = original.QuotaBaseProperties
type QuotaUpdateParameters = original.QuotaUpdateParameters
type QuotasClient = original.QuotasClient
type RegistryListCredentialsResult = original.RegistryListCredentialsResult
type Resource = original.Resource
type ResourceID = original.ResourceID
type ResourceName = original.ResourceName
type ResourceQuota = original.ResourceQuota
type ScaleSettings = original.ScaleSettings
type ServicePrincipalCredentials = original.ServicePrincipalCredentials
type SslConfiguration = original.SslConfiguration
type SystemService = original.SystemService
type UpdateWorkspaceQuotas = original.UpdateWorkspaceQuotas
type UpdateWorkspaceQuotasResult = original.UpdateWorkspaceQuotasResult
type Usage = original.Usage
type UsageName = original.UsageName
type UsagesClient = original.UsagesClient
type UserAccountCredentials = original.UserAccountCredentials
type VirtualMachine = original.VirtualMachine
type VirtualMachineProperties = original.VirtualMachineProperties
type VirtualMachineSSHCredentials = original.VirtualMachineSSHCredentials
type VirtualMachineSecrets = original.VirtualMachineSecrets
type VirtualMachineSize = original.VirtualMachineSize
type VirtualMachineSizeListResult = original.VirtualMachineSizeListResult
type VirtualMachineSizesClient = original.VirtualMachineSizesClient
type Workspace = original.Workspace
type WorkspaceListResult = original.WorkspaceListResult
type WorkspaceListResultIterator = original.WorkspaceListResultIterator
type WorkspaceListResultPage = original.WorkspaceListResultPage
type WorkspaceProperties = original.WorkspaceProperties
type WorkspacePropertiesUpdateParameters = original.WorkspacePropertiesUpdateParameters
type WorkspaceUpdateParameters = original.WorkspaceUpdateParameters
type WorkspacesClient = original.WorkspacesClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewListUsagesResultIterator(page ListUsagesResultPage) ListUsagesResultIterator {
	return original.NewListUsagesResultIterator(page)
}
func NewListUsagesResultPage(getNextPage func(context.Context, ListUsagesResult) (ListUsagesResult, error)) ListUsagesResultPage {
	return original.NewListUsagesResultPage(getNextPage)
}
func NewListWorkspaceQuotasIterator(page ListWorkspaceQuotasPage) ListWorkspaceQuotasIterator {
	return original.NewListWorkspaceQuotasIterator(page)
}
func NewListWorkspaceQuotasPage(getNextPage func(context.Context, ListWorkspaceQuotas) (ListWorkspaceQuotas, error)) ListWorkspaceQuotasPage {
	return original.NewListWorkspaceQuotasPage(getNextPage)
}
func NewMachineLearningComputeClient(subscriptionID string) MachineLearningComputeClient {
	return original.NewMachineLearningComputeClient(subscriptionID)
}
func NewMachineLearningComputeClientWithBaseURI(baseURI string, subscriptionID string) MachineLearningComputeClient {
	return original.NewMachineLearningComputeClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPaginatedComputeResourcesListIterator(page PaginatedComputeResourcesListPage) PaginatedComputeResourcesListIterator {
	return original.NewPaginatedComputeResourcesListIterator(page)
}
func NewPaginatedComputeResourcesListPage(getNextPage func(context.Context, PaginatedComputeResourcesList) (PaginatedComputeResourcesList, error)) PaginatedComputeResourcesListPage {
	return original.NewPaginatedComputeResourcesListPage(getNextPage)
}
func NewQuotasClient(subscriptionID string) QuotasClient {
	return original.NewQuotasClient(subscriptionID)
}
func NewQuotasClientWithBaseURI(baseURI string, subscriptionID string) QuotasClient {
	return original.NewQuotasClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsagesClient(subscriptionID string) UsagesClient {
	return original.NewUsagesClient(subscriptionID)
}
func NewUsagesClientWithBaseURI(baseURI string, subscriptionID string) UsagesClient {
	return original.NewUsagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewVirtualMachineSizesClient(subscriptionID string) VirtualMachineSizesClient {
	return original.NewVirtualMachineSizesClient(subscriptionID)
}
func NewVirtualMachineSizesClientWithBaseURI(baseURI string, subscriptionID string) VirtualMachineSizesClient {
	return original.NewVirtualMachineSizesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewWorkspaceListResultIterator(page WorkspaceListResultPage) WorkspaceListResultIterator {
	return original.NewWorkspaceListResultIterator(page)
}
func NewWorkspaceListResultPage(getNextPage func(context.Context, WorkspaceListResult) (WorkspaceListResult, error)) WorkspaceListResultPage {
	return original.NewWorkspaceListResultPage(getNextPage)
}
func NewWorkspacesClient(subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClient(subscriptionID)
}
func NewWorkspacesClientWithBaseURI(baseURI string, subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleAllocationStateValues() []AllocationState {
	return original.PossibleAllocationStateValues()
}
func PossibleComputeTypeBasicComputeNodesInformationValues() []ComputeTypeBasicComputeNodesInformation {
	return original.PossibleComputeTypeBasicComputeNodesInformationValues()
}
func PossibleComputeTypeBasicComputeSecretsValues() []ComputeTypeBasicComputeSecrets {
	return original.PossibleComputeTypeBasicComputeSecretsValues()
}
func PossibleComputeTypeBasicComputeValues() []ComputeTypeBasicCompute {
	return original.PossibleComputeTypeBasicComputeValues()
}
func PossibleComputeTypeValues() []ComputeType {
	return original.PossibleComputeTypeValues()
}
func PossibleNodeStateValues() []NodeState {
	return original.PossibleNodeStateValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleQuotaUnitValues() []QuotaUnit {
	return original.PossibleQuotaUnitValues()
}
func PossibleRemoteLoginPortPublicAccessValues() []RemoteLoginPortPublicAccess {
	return original.PossibleRemoteLoginPortPublicAccessValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleStatus1Values() []Status1 {
	return original.PossibleStatus1Values()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func PossibleUnderlyingResourceActionValues() []UnderlyingResourceAction {
	return original.PossibleUnderlyingResourceActionValues()
}
func PossibleUsageUnitValues() []UsageUnit {
	return original.PossibleUsageUnitValues()
}
func PossibleVMPriorityValues() []VMPriority {
	return original.PossibleVMPriorityValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
