package storagecacheapi

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
	"github.com/Azure/azure-sdk-for-go/services/storagecache/mgmt/2019-11-01/storagecache"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result storagecache.APIOperationListResultPage, err error)
}

var _ OperationsClientAPI = (*storagecache.OperationsClient)(nil)

// SkusClientAPI contains the set of methods on the SkusClient type.
type SkusClientAPI interface {
	List(ctx context.Context) (result storagecache.ResourceSkusResultPage, err error)
}

var _ SkusClientAPI = (*storagecache.SkusClient)(nil)

// UsageModelsClientAPI contains the set of methods on the UsageModelsClient type.
type UsageModelsClientAPI interface {
	List(ctx context.Context) (result storagecache.UsageModelsResultPage, err error)
}

var _ UsageModelsClientAPI = (*storagecache.UsageModelsClient)(nil)

// CachesClientAPI contains the set of methods on the CachesClient type.
type CachesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, cacheName string, cache *storagecache.Cache) (result storagecache.CachesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, cacheName string) (result storagecache.CachesDeleteFuture, err error)
	Flush(ctx context.Context, resourceGroupName string, cacheName string) (result storagecache.CachesFlushFuture, err error)
	Get(ctx context.Context, resourceGroupName string, cacheName string) (result storagecache.Cache, err error)
	List(ctx context.Context) (result storagecache.CachesListResultPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result storagecache.CachesListResultPage, err error)
	Start(ctx context.Context, resourceGroupName string, cacheName string) (result storagecache.CachesStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, cacheName string) (result storagecache.CachesStopFuture, err error)
	Update(ctx context.Context, resourceGroupName string, cacheName string, cache *storagecache.Cache) (result storagecache.Cache, err error)
	UpgradeFirmware(ctx context.Context, resourceGroupName string, cacheName string) (result storagecache.CachesUpgradeFirmwareFuture, err error)
}

var _ CachesClientAPI = (*storagecache.CachesClient)(nil)

// StorageTargetsClientAPI contains the set of methods on the StorageTargetsClient type.
type StorageTargetsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, storagetarget *storagecache.StorageTarget) (result storagecache.StorageTargetsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string) (result storagecache.StorageTargetsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string) (result storagecache.StorageTarget, err error)
	ListByCache(ctx context.Context, resourceGroupName string, cacheName string) (result storagecache.StorageTargetsResultPage, err error)
}

var _ StorageTargetsClientAPI = (*storagecache.StorageTargetsClient)(nil)
