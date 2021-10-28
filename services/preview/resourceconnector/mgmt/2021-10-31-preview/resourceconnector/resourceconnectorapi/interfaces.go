package resourceconnectorapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/resourceconnector/mgmt/2021-10-31-preview/resourceconnector"
)

// AppliancesClientAPI contains the set of methods on the AppliancesClient type.
type AppliancesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, parameters resourceconnector.Appliance) (result resourceconnector.AppliancesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result resourceconnector.AppliancesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result resourceconnector.Appliance, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result resourceconnector.ApplianceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result resourceconnector.ApplianceListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result resourceconnector.ApplianceListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result resourceconnector.ApplianceListResultIterator, err error)
	ListClusterUserCredential(ctx context.Context, resourceGroupName string, resourceName string) (result resourceconnector.ApplianceListCredentialResults, err error)
	ListOperations(ctx context.Context) (result resourceconnector.ApplianceOperationsListPage, err error)
	ListOperationsComplete(ctx context.Context) (result resourceconnector.ApplianceOperationsListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, parameters resourceconnector.PatchableAppliance) (result resourceconnector.Appliance, err error)
}

var _ AppliancesClientAPI = (*resourceconnector.AppliancesClient)(nil)
