package computeapi

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
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-04-01/compute"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result compute.OperationListResult, err error)
}

var _ OperationsClientAPI = (*compute.OperationsClient)(nil)

// AvailabilitySetsClientAPI contains the set of methods on the AvailabilitySetsClient type.
type AvailabilitySetsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters compute.AvailabilitySet) (result compute.AvailabilitySet, err error)
	Delete(ctx context.Context, resourceGroupName string, availabilitySetName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, availabilitySetName string) (result compute.AvailabilitySet, err error)
	List(ctx context.Context, resourceGroupName string) (result compute.AvailabilitySetListResultPage, err error)
	ListAvailableSizes(ctx context.Context, resourceGroupName string, availabilitySetName string) (result compute.VirtualMachineSizeListResult, err error)
	ListBySubscription(ctx context.Context, expand string) (result compute.AvailabilitySetListResultPage, err error)
	Update(ctx context.Context, resourceGroupName string, availabilitySetName string, parameters compute.AvailabilitySetUpdate) (result compute.AvailabilitySet, err error)
}

var _ AvailabilitySetsClientAPI = (*compute.AvailabilitySetsClient)(nil)

// ProximityPlacementGroupsClientAPI contains the set of methods on the ProximityPlacementGroupsClient type.
type ProximityPlacementGroupsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters compute.ProximityPlacementGroup) (result compute.ProximityPlacementGroup, err error)
	Delete(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string) (result compute.ProximityPlacementGroup, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result compute.ProximityPlacementGroupListResultPage, err error)
	ListBySubscription(ctx context.Context) (result compute.ProximityPlacementGroupListResultPage, err error)
	Update(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters compute.ProximityPlacementGroupUpdate) (result compute.ProximityPlacementGroup, err error)
}

var _ ProximityPlacementGroupsClientAPI = (*compute.ProximityPlacementGroupsClient)(nil)

// VirtualMachineExtensionImagesClientAPI contains the set of methods on the VirtualMachineExtensionImagesClient type.
type VirtualMachineExtensionImagesClientAPI interface {
	Get(ctx context.Context, location string, publisherName string, typeParameter string, version string) (result compute.VirtualMachineExtensionImage, err error)
	ListTypes(ctx context.Context, location string, publisherName string) (result compute.ListVirtualMachineExtensionImage, err error)
	ListVersions(ctx context.Context, location string, publisherName string, typeParameter string, filter string, top *int32, orderby string) (result compute.ListVirtualMachineExtensionImage, err error)
}

var _ VirtualMachineExtensionImagesClientAPI = (*compute.VirtualMachineExtensionImagesClient)(nil)

// VirtualMachineExtensionsClientAPI contains the set of methods on the VirtualMachineExtensionsClient type.
type VirtualMachineExtensionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, VMName string, VMExtensionName string, extensionParameters compute.VirtualMachineExtension) (result compute.VirtualMachineExtensionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, VMName string, VMExtensionName string) (result compute.VirtualMachineExtensionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, VMName string, VMExtensionName string, expand string) (result compute.VirtualMachineExtension, err error)
	List(ctx context.Context, resourceGroupName string, VMName string, expand string) (result compute.VirtualMachineExtensionsListResult, err error)
	Update(ctx context.Context, resourceGroupName string, VMName string, VMExtensionName string, extensionParameters compute.VirtualMachineExtensionUpdate) (result compute.VirtualMachineExtensionsUpdateFuture, err error)
}

var _ VirtualMachineExtensionsClientAPI = (*compute.VirtualMachineExtensionsClient)(nil)

// VirtualMachineImagesClientAPI contains the set of methods on the VirtualMachineImagesClient type.
type VirtualMachineImagesClientAPI interface {
	Get(ctx context.Context, location string, publisherName string, offer string, skus string, version string) (result compute.VirtualMachineImage, err error)
	List(ctx context.Context, location string, publisherName string, offer string, skus string, filter string, top *int32, orderby string) (result compute.ListVirtualMachineImageResource, err error)
	ListOffers(ctx context.Context, location string, publisherName string) (result compute.ListVirtualMachineImageResource, err error)
	ListPublishers(ctx context.Context, location string) (result compute.ListVirtualMachineImageResource, err error)
	ListSkus(ctx context.Context, location string, publisherName string, offer string) (result compute.ListVirtualMachineImageResource, err error)
}

var _ VirtualMachineImagesClientAPI = (*compute.VirtualMachineImagesClient)(nil)

// UsageClientAPI contains the set of methods on the UsageClient type.
type UsageClientAPI interface {
	List(ctx context.Context, location string) (result compute.ListUsagesResultPage, err error)
}

var _ UsageClientAPI = (*compute.UsageClient)(nil)

// VirtualMachinesClientAPI contains the set of methods on the VirtualMachinesClient type.
type VirtualMachinesClientAPI interface {
	Capture(ctx context.Context, resourceGroupName string, VMName string, parameters compute.VirtualMachineCaptureParameters) (result compute.VirtualMachinesCaptureFuture, err error)
	ConvertToManagedDisks(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachinesConvertToManagedDisksFuture, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, VMName string, parameters compute.VirtualMachine) (result compute.VirtualMachinesCreateOrUpdateFuture, err error)
	Deallocate(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachinesDeallocateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachinesDeleteFuture, err error)
	Generalize(ctx context.Context, resourceGroupName string, VMName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, VMName string, expand compute.InstanceViewTypes) (result compute.VirtualMachine, err error)
	InstanceView(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachineInstanceView, err error)
	List(ctx context.Context, resourceGroupName string) (result compute.VirtualMachineListResultPage, err error)
	ListAll(ctx context.Context) (result compute.VirtualMachineListResultPage, err error)
	ListAvailableSizes(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachineSizeListResult, err error)
	ListByLocation(ctx context.Context, location string) (result compute.VirtualMachineListResultPage, err error)
	PerformMaintenance(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachinesPerformMaintenanceFuture, err error)
	PowerOff(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachinesPowerOffFuture, err error)
	Redeploy(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachinesRedeployFuture, err error)
	Restart(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachinesRestartFuture, err error)
	RunCommand(ctx context.Context, resourceGroupName string, VMName string, parameters compute.RunCommandInput) (result compute.VirtualMachinesRunCommandFuture, err error)
	Start(ctx context.Context, resourceGroupName string, VMName string) (result compute.VirtualMachinesStartFuture, err error)
	Update(ctx context.Context, resourceGroupName string, VMName string, parameters compute.VirtualMachineUpdate) (result compute.VirtualMachinesUpdateFuture, err error)
}

var _ VirtualMachinesClientAPI = (*compute.VirtualMachinesClient)(nil)

// VirtualMachineSizesClientAPI contains the set of methods on the VirtualMachineSizesClient type.
type VirtualMachineSizesClientAPI interface {
	List(ctx context.Context, location string) (result compute.VirtualMachineSizeListResult, err error)
}

var _ VirtualMachineSizesClientAPI = (*compute.VirtualMachineSizesClient)(nil)

// ImagesClientAPI contains the set of methods on the ImagesClient type.
type ImagesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, imageName string, parameters compute.Image) (result compute.ImagesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, imageName string) (result compute.ImagesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, imageName string, expand string) (result compute.Image, err error)
	List(ctx context.Context) (result compute.ImageListResultPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result compute.ImageListResultPage, err error)
	Update(ctx context.Context, resourceGroupName string, imageName string, parameters compute.ImageUpdate) (result compute.ImagesUpdateFuture, err error)
}

var _ ImagesClientAPI = (*compute.ImagesClient)(nil)

// VirtualMachineScaleSetsClientAPI contains the set of methods on the VirtualMachineScaleSetsClient type.
type VirtualMachineScaleSetsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, VMScaleSetName string, parameters compute.VirtualMachineScaleSet) (result compute.VirtualMachineScaleSetsCreateOrUpdateFuture, err error)
	Deallocate(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs *compute.VirtualMachineScaleSetVMInstanceIDs) (result compute.VirtualMachineScaleSetsDeallocateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSetsDeleteFuture, err error)
	DeleteInstances(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs compute.VirtualMachineScaleSetVMInstanceRequiredIDs) (result compute.VirtualMachineScaleSetsDeleteInstancesFuture, err error)
	ForceRecoveryServiceFabricPlatformUpdateDomainWalk(ctx context.Context, resourceGroupName string, VMScaleSetName string, platformUpdateDomain int32) (result compute.RecoveryWalkResponse, err error)
	Get(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSet, err error)
	GetInstanceView(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSetInstanceView, err error)
	GetOSUpgradeHistory(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSetListOSUpgradeHistoryPage, err error)
	List(ctx context.Context, resourceGroupName string) (result compute.VirtualMachineScaleSetListResultPage, err error)
	ListAll(ctx context.Context) (result compute.VirtualMachineScaleSetListWithLinkResultPage, err error)
	ListSkus(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSetListSkusResultPage, err error)
	PerformMaintenance(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs *compute.VirtualMachineScaleSetVMInstanceIDs) (result compute.VirtualMachineScaleSetsPerformMaintenanceFuture, err error)
	PowerOff(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs *compute.VirtualMachineScaleSetVMInstanceIDs) (result compute.VirtualMachineScaleSetsPowerOffFuture, err error)
	Redeploy(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs *compute.VirtualMachineScaleSetVMInstanceIDs) (result compute.VirtualMachineScaleSetsRedeployFuture, err error)
	Reimage(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs *compute.VirtualMachineScaleSetVMInstanceIDs) (result compute.VirtualMachineScaleSetsReimageFuture, err error)
	ReimageAll(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs *compute.VirtualMachineScaleSetVMInstanceIDs) (result compute.VirtualMachineScaleSetsReimageAllFuture, err error)
	Restart(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs *compute.VirtualMachineScaleSetVMInstanceIDs) (result compute.VirtualMachineScaleSetsRestartFuture, err error)
	Start(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs *compute.VirtualMachineScaleSetVMInstanceIDs) (result compute.VirtualMachineScaleSetsStartFuture, err error)
	Update(ctx context.Context, resourceGroupName string, VMScaleSetName string, parameters compute.VirtualMachineScaleSetUpdate) (result compute.VirtualMachineScaleSetsUpdateFuture, err error)
	UpdateInstances(ctx context.Context, resourceGroupName string, VMScaleSetName string, VMInstanceIDs compute.VirtualMachineScaleSetVMInstanceRequiredIDs) (result compute.VirtualMachineScaleSetsUpdateInstancesFuture, err error)
}

var _ VirtualMachineScaleSetsClientAPI = (*compute.VirtualMachineScaleSetsClient)(nil)

// VirtualMachineScaleSetExtensionsClientAPI contains the set of methods on the VirtualMachineScaleSetExtensionsClient type.
type VirtualMachineScaleSetExtensionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, VMScaleSetName string, vmssExtensionName string, extensionParameters compute.VirtualMachineScaleSetExtension) (result compute.VirtualMachineScaleSetExtensionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, VMScaleSetName string, vmssExtensionName string) (result compute.VirtualMachineScaleSetExtensionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, VMScaleSetName string, vmssExtensionName string, expand string) (result compute.VirtualMachineScaleSetExtension, err error)
	List(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSetExtensionListResultPage, err error)
}

var _ VirtualMachineScaleSetExtensionsClientAPI = (*compute.VirtualMachineScaleSetExtensionsClient)(nil)

// VirtualMachineScaleSetRollingUpgradesClientAPI contains the set of methods on the VirtualMachineScaleSetRollingUpgradesClient type.
type VirtualMachineScaleSetRollingUpgradesClientAPI interface {
	Cancel(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSetRollingUpgradesCancelFuture, err error)
	GetLatest(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.RollingUpgradeStatusInfo, err error)
	StartOSUpgrade(ctx context.Context, resourceGroupName string, VMScaleSetName string) (result compute.VirtualMachineScaleSetRollingUpgradesStartOSUpgradeFuture, err error)
}

var _ VirtualMachineScaleSetRollingUpgradesClientAPI = (*compute.VirtualMachineScaleSetRollingUpgradesClient)(nil)

// VirtualMachineScaleSetVMsClientAPI contains the set of methods on the VirtualMachineScaleSetVMsClient type.
type VirtualMachineScaleSetVMsClientAPI interface {
	Deallocate(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMsDeallocateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVM, err error)
	GetInstanceView(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMInstanceView, err error)
	List(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, filter string, selectParameter string, expand string) (result compute.VirtualMachineScaleSetVMListResultPage, err error)
	PerformMaintenance(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMsPerformMaintenanceFuture, err error)
	PowerOff(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMsPowerOffFuture, err error)
	Redeploy(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMsRedeployFuture, err error)
	Reimage(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMsReimageFuture, err error)
	ReimageAll(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMsReimageAllFuture, err error)
	Restart(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMsRestartFuture, err error)
	RunCommand(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, parameters compute.RunCommandInput) (result compute.VirtualMachineScaleSetVMsRunCommandFuture, err error)
	Start(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string) (result compute.VirtualMachineScaleSetVMsStartFuture, err error)
	Update(ctx context.Context, resourceGroupName string, VMScaleSetName string, instanceID string, parameters compute.VirtualMachineScaleSetVM) (result compute.VirtualMachineScaleSetVMsUpdateFuture, err error)
}

var _ VirtualMachineScaleSetVMsClientAPI = (*compute.VirtualMachineScaleSetVMsClient)(nil)

// LogAnalyticsClientAPI contains the set of methods on the LogAnalyticsClient type.
type LogAnalyticsClientAPI interface {
	ExportRequestRateByInterval(ctx context.Context, parameters compute.RequestRateByIntervalInput, location string) (result compute.LogAnalyticsExportRequestRateByIntervalFuture, err error)
	ExportThrottledRequests(ctx context.Context, parameters compute.ThrottledRequestsInput, location string) (result compute.LogAnalyticsExportThrottledRequestsFuture, err error)
}

var _ LogAnalyticsClientAPI = (*compute.LogAnalyticsClient)(nil)

// VirtualMachineRunCommandsClientAPI contains the set of methods on the VirtualMachineRunCommandsClient type.
type VirtualMachineRunCommandsClientAPI interface {
	Get(ctx context.Context, location string, commandID string) (result compute.RunCommandDocument, err error)
	List(ctx context.Context, location string) (result compute.RunCommandListResultPage, err error)
}

var _ VirtualMachineRunCommandsClientAPI = (*compute.VirtualMachineRunCommandsClient)(nil)

// DisksClientAPI contains the set of methods on the DisksClient type.
type DisksClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, diskName string, disk compute.Disk) (result compute.DisksCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, diskName string) (result compute.DisksDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, diskName string) (result compute.Disk, err error)
	GrantAccess(ctx context.Context, resourceGroupName string, diskName string, grantAccessData compute.GrantAccessData) (result compute.DisksGrantAccessFuture, err error)
	List(ctx context.Context) (result compute.DiskListPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result compute.DiskListPage, err error)
	RevokeAccess(ctx context.Context, resourceGroupName string, diskName string) (result compute.DisksRevokeAccessFuture, err error)
	Update(ctx context.Context, resourceGroupName string, diskName string, disk compute.DiskUpdate) (result compute.DisksUpdateFuture, err error)
}

var _ DisksClientAPI = (*compute.DisksClient)(nil)

// SnapshotsClientAPI contains the set of methods on the SnapshotsClient type.
type SnapshotsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, snapshotName string, snapshot compute.Snapshot) (result compute.SnapshotsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, snapshotName string) (result compute.SnapshotsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, snapshotName string) (result compute.Snapshot, err error)
	GrantAccess(ctx context.Context, resourceGroupName string, snapshotName string, grantAccessData compute.GrantAccessData) (result compute.SnapshotsGrantAccessFuture, err error)
	List(ctx context.Context) (result compute.SnapshotListPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result compute.SnapshotListPage, err error)
	RevokeAccess(ctx context.Context, resourceGroupName string, snapshotName string) (result compute.SnapshotsRevokeAccessFuture, err error)
	Update(ctx context.Context, resourceGroupName string, snapshotName string, snapshot compute.SnapshotUpdate) (result compute.SnapshotsUpdateFuture, err error)
}

var _ SnapshotsClientAPI = (*compute.SnapshotsClient)(nil)
