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

package hdinsight

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/hdinsight/mgmt/2018-06-01-preview/hdinsight"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AsyncOperationState = original.AsyncOperationState

const (
	Failed     AsyncOperationState = original.Failed
	InProgress AsyncOperationState = original.InProgress
	Succeeded  AsyncOperationState = original.Succeeded
)

type ClusterProvisioningState = original.ClusterProvisioningState

const (
	ClusterProvisioningStateCanceled   ClusterProvisioningState = original.ClusterProvisioningStateCanceled
	ClusterProvisioningStateDeleting   ClusterProvisioningState = original.ClusterProvisioningStateDeleting
	ClusterProvisioningStateFailed     ClusterProvisioningState = original.ClusterProvisioningStateFailed
	ClusterProvisioningStateInProgress ClusterProvisioningState = original.ClusterProvisioningStateInProgress
	ClusterProvisioningStateSucceeded  ClusterProvisioningState = original.ClusterProvisioningStateSucceeded
)

type DaysOfWeek = original.DaysOfWeek

const (
	Friday    DaysOfWeek = original.Friday
	Monday    DaysOfWeek = original.Monday
	Saturday  DaysOfWeek = original.Saturday
	Sunday    DaysOfWeek = original.Sunday
	Thursday  DaysOfWeek = original.Thursday
	Tuesday   DaysOfWeek = original.Tuesday
	Wednesday DaysOfWeek = original.Wednesday
)

type DirectoryType = original.DirectoryType

const (
	ActiveDirectory DirectoryType = original.ActiveDirectory
)

type FilterMode = original.FilterMode

const (
	Exclude FilterMode = original.Exclude
	Include FilterMode = original.Include
)

type JSONWebKeyEncryptionAlgorithm = original.JSONWebKeyEncryptionAlgorithm

const (
	RSA15      JSONWebKeyEncryptionAlgorithm = original.RSA15
	RSAOAEP    JSONWebKeyEncryptionAlgorithm = original.RSAOAEP
	RSAOAEP256 JSONWebKeyEncryptionAlgorithm = original.RSAOAEP256
)

type OSType = original.OSType

const (
	Linux   OSType = original.Linux
	Windows OSType = original.Windows
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	None                       ResourceIdentityType = original.None
	SystemAssigned             ResourceIdentityType = original.SystemAssigned
	SystemAssignedUserAssigned ResourceIdentityType = original.SystemAssignedUserAssigned
	UserAssigned               ResourceIdentityType = original.UserAssigned
)

type Tier = original.Tier

const (
	Premium  Tier = original.Premium
	Standard Tier = original.Standard
)

type Application = original.Application
type ApplicationGetEndpoint = original.ApplicationGetEndpoint
type ApplicationGetHTTPSEndpoint = original.ApplicationGetHTTPSEndpoint
type ApplicationListResult = original.ApplicationListResult
type ApplicationListResultIterator = original.ApplicationListResultIterator
type ApplicationListResultPage = original.ApplicationListResultPage
type ApplicationProperties = original.ApplicationProperties
type ApplicationsClient = original.ApplicationsClient
type ApplicationsCreateFuture = original.ApplicationsCreateFuture
type ApplicationsDeleteFuture = original.ApplicationsDeleteFuture
type Autoscale = original.Autoscale
type AutoscaleCapacity = original.AutoscaleCapacity
type AutoscaleRecurrence = original.AutoscaleRecurrence
type AutoscaleSchedule = original.AutoscaleSchedule
type AutoscaleTimeAndCapacity = original.AutoscaleTimeAndCapacity
type BaseClient = original.BaseClient
type BillingMeters = original.BillingMeters
type BillingResources = original.BillingResources
type BillingResponseListResult = original.BillingResponseListResult
type CapabilitiesResult = original.CapabilitiesResult
type Cluster = original.Cluster
type ClusterConfigurations = original.ClusterConfigurations
type ClusterCreateParametersExtended = original.ClusterCreateParametersExtended
type ClusterCreateProperties = original.ClusterCreateProperties
type ClusterDefinition = original.ClusterDefinition
type ClusterDiskEncryptionParameters = original.ClusterDiskEncryptionParameters
type ClusterGetProperties = original.ClusterGetProperties
type ClusterIdentity = original.ClusterIdentity
type ClusterIdentityUserAssignedIdentitiesValue = original.ClusterIdentityUserAssignedIdentitiesValue
type ClusterListPersistedScriptActionsResult = original.ClusterListPersistedScriptActionsResult
type ClusterListResult = original.ClusterListResult
type ClusterListResultIterator = original.ClusterListResultIterator
type ClusterListResultPage = original.ClusterListResultPage
type ClusterListRuntimeScriptActionDetailResult = original.ClusterListRuntimeScriptActionDetailResult
type ClusterMonitoringRequest = original.ClusterMonitoringRequest
type ClusterMonitoringResponse = original.ClusterMonitoringResponse
type ClusterPatchParameters = original.ClusterPatchParameters
type ClusterResizeParameters = original.ClusterResizeParameters
type ClustersClient = original.ClustersClient
type ClustersCreateFuture = original.ClustersCreateFuture
type ClustersDeleteFuture = original.ClustersDeleteFuture
type ClustersExecuteScriptActionsFuture = original.ClustersExecuteScriptActionsFuture
type ClustersResizeFuture = original.ClustersResizeFuture
type ClustersRotateDiskEncryptionKeyFuture = original.ClustersRotateDiskEncryptionKeyFuture
type ClustersUpdateGatewaySettingsFuture = original.ClustersUpdateGatewaySettingsFuture
type ComputeProfile = original.ComputeProfile
type ConfigurationsClient = original.ConfigurationsClient
type ConfigurationsUpdateFuture = original.ConfigurationsUpdateFuture
type ConnectivityEndpoint = original.ConnectivityEndpoint
type DataDisksGroups = original.DataDisksGroups
type DiskBillingMeters = original.DiskBillingMeters
type DiskEncryptionProperties = original.DiskEncryptionProperties
type ErrorResponse = original.ErrorResponse
type Errors = original.Errors
type ExecuteScriptActionParameters = original.ExecuteScriptActionParameters
type Extension = original.Extension
type ExtensionsClient = original.ExtensionsClient
type ExtensionsCreateFuture = original.ExtensionsCreateFuture
type ExtensionsDeleteFuture = original.ExtensionsDeleteFuture
type ExtensionsDisableMonitoringFuture = original.ExtensionsDisableMonitoringFuture
type ExtensionsEnableMonitoringFuture = original.ExtensionsEnableMonitoringFuture
type GatewaySettings = original.GatewaySettings
type HardwareProfile = original.HardwareProfile
type LinuxOperatingSystemProfile = original.LinuxOperatingSystemProfile
type LocalizedName = original.LocalizedName
type LocationsClient = original.LocationsClient
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationResource = original.OperationResource
type OperationsClient = original.OperationsClient
type OsProfile = original.OsProfile
type ProxyResource = original.ProxyResource
type QuotaCapability = original.QuotaCapability
type QuotaInfo = original.QuotaInfo
type RegionalQuotaCapability = original.RegionalQuotaCapability
type RegionsCapability = original.RegionsCapability
type Resource = original.Resource
type Role = original.Role
type RuntimeScriptAction = original.RuntimeScriptAction
type RuntimeScriptActionDetail = original.RuntimeScriptActionDetail
type SSHProfile = original.SSHProfile
type SSHPublicKey = original.SSHPublicKey
type ScriptAction = original.ScriptAction
type ScriptActionExecutionHistoryList = original.ScriptActionExecutionHistoryList
type ScriptActionExecutionHistoryListIterator = original.ScriptActionExecutionHistoryListIterator
type ScriptActionExecutionHistoryListPage = original.ScriptActionExecutionHistoryListPage
type ScriptActionExecutionSummary = original.ScriptActionExecutionSummary
type ScriptActionPersistedGetResponseSpec = original.ScriptActionPersistedGetResponseSpec
type ScriptActionsClient = original.ScriptActionsClient
type ScriptActionsList = original.ScriptActionsList
type ScriptActionsListIterator = original.ScriptActionsListIterator
type ScriptActionsListPage = original.ScriptActionsListPage
type ScriptExecutionHistoryClient = original.ScriptExecutionHistoryClient
type SecurityProfile = original.SecurityProfile
type SetString = original.SetString
type StorageAccount = original.StorageAccount
type StorageProfile = original.StorageProfile
type TrackedResource = original.TrackedResource
type UpdateGatewaySettingsParameters = original.UpdateGatewaySettingsParameters
type Usage = original.Usage
type UsagesListResult = original.UsagesListResult
type VMSizeCompatibilityFilter = original.VMSizeCompatibilityFilter
type VMSizeCompatibilityFilterV2 = original.VMSizeCompatibilityFilterV2
type VMSizesCapability = original.VMSizesCapability
type VersionSpec = original.VersionSpec
type VersionsCapability = original.VersionsCapability
type VirtualNetworkProfile = original.VirtualNetworkProfile

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewApplicationListResultIterator(page ApplicationListResultPage) ApplicationListResultIterator {
	return original.NewApplicationListResultIterator(page)
}
func NewApplicationListResultPage(getNextPage func(context.Context, ApplicationListResult) (ApplicationListResult, error)) ApplicationListResultPage {
	return original.NewApplicationListResultPage(getNextPage)
}
func NewApplicationsClient(subscriptionID string) ApplicationsClient {
	return original.NewApplicationsClient(subscriptionID)
}
func NewApplicationsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationsClient {
	return original.NewApplicationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewClusterListResultIterator(page ClusterListResultPage) ClusterListResultIterator {
	return original.NewClusterListResultIterator(page)
}
func NewClusterListResultPage(getNextPage func(context.Context, ClusterListResult) (ClusterListResult, error)) ClusterListResultPage {
	return original.NewClusterListResultPage(getNextPage)
}
func NewClustersClient(subscriptionID string) ClustersClient {
	return original.NewClustersClient(subscriptionID)
}
func NewClustersClientWithBaseURI(baseURI string, subscriptionID string) ClustersClient {
	return original.NewClustersClientWithBaseURI(baseURI, subscriptionID)
}
func NewConfigurationsClient(subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClient(subscriptionID)
}
func NewConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewExtensionsClient(subscriptionID string) ExtensionsClient {
	return original.NewExtensionsClient(subscriptionID)
}
func NewExtensionsClientWithBaseURI(baseURI string, subscriptionID string) ExtensionsClient {
	return original.NewExtensionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewLocationsClient(subscriptionID string) LocationsClient {
	return original.NewLocationsClient(subscriptionID)
}
func NewLocationsClientWithBaseURI(baseURI string, subscriptionID string) LocationsClient {
	return original.NewLocationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewScriptActionExecutionHistoryListIterator(page ScriptActionExecutionHistoryListPage) ScriptActionExecutionHistoryListIterator {
	return original.NewScriptActionExecutionHistoryListIterator(page)
}
func NewScriptActionExecutionHistoryListPage(getNextPage func(context.Context, ScriptActionExecutionHistoryList) (ScriptActionExecutionHistoryList, error)) ScriptActionExecutionHistoryListPage {
	return original.NewScriptActionExecutionHistoryListPage(getNextPage)
}
func NewScriptActionsClient(subscriptionID string) ScriptActionsClient {
	return original.NewScriptActionsClient(subscriptionID)
}
func NewScriptActionsClientWithBaseURI(baseURI string, subscriptionID string) ScriptActionsClient {
	return original.NewScriptActionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewScriptActionsListIterator(page ScriptActionsListPage) ScriptActionsListIterator {
	return original.NewScriptActionsListIterator(page)
}
func NewScriptActionsListPage(getNextPage func(context.Context, ScriptActionsList) (ScriptActionsList, error)) ScriptActionsListPage {
	return original.NewScriptActionsListPage(getNextPage)
}
func NewScriptExecutionHistoryClient(subscriptionID string) ScriptExecutionHistoryClient {
	return original.NewScriptExecutionHistoryClient(subscriptionID)
}
func NewScriptExecutionHistoryClientWithBaseURI(baseURI string, subscriptionID string) ScriptExecutionHistoryClient {
	return original.NewScriptExecutionHistoryClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAsyncOperationStateValues() []AsyncOperationState {
	return original.PossibleAsyncOperationStateValues()
}
func PossibleClusterProvisioningStateValues() []ClusterProvisioningState {
	return original.PossibleClusterProvisioningStateValues()
}
func PossibleDaysOfWeekValues() []DaysOfWeek {
	return original.PossibleDaysOfWeekValues()
}
func PossibleDirectoryTypeValues() []DirectoryType {
	return original.PossibleDirectoryTypeValues()
}
func PossibleFilterModeValues() []FilterMode {
	return original.PossibleFilterModeValues()
}
func PossibleJSONWebKeyEncryptionAlgorithmValues() []JSONWebKeyEncryptionAlgorithm {
	return original.PossibleJSONWebKeyEncryptionAlgorithmValues()
}
func PossibleOSTypeValues() []OSType {
	return original.PossibleOSTypeValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleTierValues() []Tier {
	return original.PossibleTierValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
