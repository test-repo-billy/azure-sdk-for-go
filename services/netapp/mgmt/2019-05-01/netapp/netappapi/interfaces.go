package netappapi

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
	"github.com/Azure/azure-sdk-for-go/services/netapp/mgmt/2019-05-01/netapp"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	CheckFilePathAvailability(ctx context.Context, location string) (result netapp.ResourceNameAvailability, err error)
	CheckNameAvailability(ctx context.Context, location string) (result netapp.ResourceNameAvailability, err error)
}

var _ BaseClientAPI = (*netapp.BaseClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result netapp.OperationListResult, err error)
}

var _ OperationsClientAPI = (*netapp.OperationsClient)(nil)

// AccountsClientAPI contains the set of methods on the AccountsClient type.
type AccountsClientAPI interface {
	CreateOrUpdate(ctx context.Context, body netapp.Account, resourceGroupName string, accountName string) (result netapp.AccountsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result netapp.AccountsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result netapp.Account, err error)
	List(ctx context.Context, resourceGroupName string) (result netapp.AccountList, err error)
	Update(ctx context.Context, body netapp.AccountPatch, resourceGroupName string, accountName string) (result netapp.Account, err error)
}

var _ AccountsClientAPI = (*netapp.AccountsClient)(nil)

// PoolsClientAPI contains the set of methods on the PoolsClient type.
type PoolsClientAPI interface {
	CreateOrUpdate(ctx context.Context, body netapp.CapacityPool, resourceGroupName string, accountName string, poolName string) (result netapp.PoolsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, poolName string) (result netapp.PoolsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, poolName string) (result netapp.CapacityPool, err error)
	List(ctx context.Context, resourceGroupName string, accountName string) (result netapp.CapacityPoolList, err error)
	Update(ctx context.Context, body netapp.CapacityPoolPatch, resourceGroupName string, accountName string, poolName string) (result netapp.CapacityPool, err error)
}

var _ PoolsClientAPI = (*netapp.PoolsClient)(nil)

// VolumesClientAPI contains the set of methods on the VolumesClient type.
type VolumesClientAPI interface {
	CreateOrUpdate(ctx context.Context, body netapp.Volume, resourceGroupName string, accountName string, poolName string, volumeName string) (result netapp.VolumesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string) (result netapp.VolumesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string) (result netapp.Volume, err error)
	List(ctx context.Context, resourceGroupName string, accountName string, poolName string) (result netapp.VolumeList, err error)
	Update(ctx context.Context, body netapp.VolumePatch, resourceGroupName string, accountName string, poolName string, volumeName string) (result netapp.Volume, err error)
}

var _ VolumesClientAPI = (*netapp.VolumesClient)(nil)

// MountTargetsClientAPI contains the set of methods on the MountTargetsClient type.
type MountTargetsClientAPI interface {
	List(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string) (result netapp.MountTargetList, err error)
}

var _ MountTargetsClientAPI = (*netapp.MountTargetsClient)(nil)

// SnapshotsClientAPI contains the set of methods on the SnapshotsClient type.
type SnapshotsClientAPI interface {
	Create(ctx context.Context, body netapp.Snapshot, resourceGroupName string, accountName string, poolName string, volumeName string, snapshotName string) (result netapp.SnapshotsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, snapshotName string) (result netapp.SnapshotsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, snapshotName string) (result netapp.Snapshot, err error)
	List(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string) (result netapp.SnapshotsList, err error)
	Update(ctx context.Context, body netapp.SnapshotPatch, resourceGroupName string, accountName string, poolName string, volumeName string, snapshotName string) (result netapp.Snapshot, err error)
}

var _ SnapshotsClientAPI = (*netapp.SnapshotsClient)(nil)
