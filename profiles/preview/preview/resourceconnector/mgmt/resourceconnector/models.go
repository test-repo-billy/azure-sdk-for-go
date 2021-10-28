//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package resourceconnector

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/resourceconnector/mgmt/2021-10-31-preview/resourceconnector"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccessProfileType = original.AccessProfileType

const (
	AccessProfileTypeClusterUser AccessProfileType = original.AccessProfileTypeClusterUser
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type Distro = original.Distro

const (
	DistroAKSEdge Distro = original.DistroAKSEdge
)

type Provider = original.Provider

const (
	ProviderHCI    Provider = original.ProviderHCI
	ProviderSCVMM  Provider = original.ProviderSCVMM
	ProviderVMWare Provider = original.ProviderVMWare
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	ResourceIdentityTypeNone           ResourceIdentityType = original.ResourceIdentityTypeNone
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = original.ResourceIdentityTypeSystemAssigned
)

type Status = original.Status

const (
	StatusConnected           Status = original.StatusConnected
	StatusRunning             Status = original.StatusRunning
	StatusValidating          Status = original.StatusValidating
	StatusWaitingForHeartbeat Status = original.StatusWaitingForHeartbeat
)

type Appliance = original.Appliance
type ApplianceCredentialKubeconfig = original.ApplianceCredentialKubeconfig
type ApplianceListCredentialResults = original.ApplianceListCredentialResults
type ApplianceListResult = original.ApplianceListResult
type ApplianceListResultIterator = original.ApplianceListResultIterator
type ApplianceListResultPage = original.ApplianceListResultPage
type ApplianceOperation = original.ApplianceOperation
type ApplianceOperationValueDisplay = original.ApplianceOperationValueDisplay
type ApplianceOperationsList = original.ApplianceOperationsList
type ApplianceOperationsListIterator = original.ApplianceOperationsListIterator
type ApplianceOperationsListPage = original.ApplianceOperationsListPage
type ApplianceProperties = original.ApplianceProperties
type AppliancePropertiesInfrastructureConfig = original.AppliancePropertiesInfrastructureConfig
type AppliancesClient = original.AppliancesClient
type AppliancesCreateOrUpdateFuture = original.AppliancesCreateOrUpdateFuture
type AppliancesDeleteFuture = original.AppliancesDeleteFuture
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type HybridConnectionConfig = original.HybridConnectionConfig
type Identity = original.Identity
type PatchableAppliance = original.PatchableAppliance
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type SystemData = original.SystemData
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewApplianceListResultIterator(page ApplianceListResultPage) ApplianceListResultIterator {
	return original.NewApplianceListResultIterator(page)
}
func NewApplianceListResultPage(cur ApplianceListResult, getNextPage func(context.Context, ApplianceListResult) (ApplianceListResult, error)) ApplianceListResultPage {
	return original.NewApplianceListResultPage(cur, getNextPage)
}
func NewApplianceOperationsListIterator(page ApplianceOperationsListPage) ApplianceOperationsListIterator {
	return original.NewApplianceOperationsListIterator(page)
}
func NewApplianceOperationsListPage(cur ApplianceOperationsList, getNextPage func(context.Context, ApplianceOperationsList) (ApplianceOperationsList, error)) ApplianceOperationsListPage {
	return original.NewApplianceOperationsListPage(cur, getNextPage)
}
func NewAppliancesClient(subscriptionID string) AppliancesClient {
	return original.NewAppliancesClient(subscriptionID)
}
func NewAppliancesClientWithBaseURI(baseURI string, subscriptionID string) AppliancesClient {
	return original.NewAppliancesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccessProfileTypeValues() []AccessProfileType {
	return original.PossibleAccessProfileTypeValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleDistroValues() []Distro {
	return original.PossibleDistroValues()
}
func PossibleProviderValues() []Provider {
	return original.PossibleProviderValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
