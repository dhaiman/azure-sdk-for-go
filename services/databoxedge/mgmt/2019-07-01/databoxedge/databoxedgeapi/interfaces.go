package databoxedgeapi

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
    "github.com/Azure/azure-sdk-for-go/services/databoxedge/mgmt/2019-07-01/databoxedge"
)

        // OperationsClientAPI contains the set of methods on the OperationsClient type.
        type OperationsClientAPI interface {
            List(ctx context.Context) (result databoxedge.OperationsListPage, err error)
                ListComplete(ctx context.Context) (result databoxedge.OperationsListIterator, err error)
        }

        var _ OperationsClientAPI = (*databoxedge.OperationsClient)(nil)
        // DevicesClientAPI contains the set of methods on the DevicesClient type.
        type DevicesClientAPI interface {
            CreateOrUpdate(ctx context.Context, deviceName string, dataBoxEdgeDevice databoxedge.Device, resourceGroupName string) (result databoxedge.DevicesCreateOrUpdateFuture, err error)
            CreateOrUpdateSecuritySettings(ctx context.Context, deviceName string, securitySettings databoxedge.SecuritySettings, resourceGroupName string) (result databoxedge.DevicesCreateOrUpdateSecuritySettingsFuture, err error)
            Delete(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.DevicesDeleteFuture, err error)
            DownloadUpdates(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.DevicesDownloadUpdatesFuture, err error)
            Get(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.Device, err error)
            GetExtendedInformation(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.DeviceExtendedInfo, err error)
            GetNetworkSettings(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.NetworkSettings, err error)
            GetUpdateSummary(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.UpdateSummary, err error)
            InstallUpdates(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.DevicesInstallUpdatesFuture, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string, expand string) (result databoxedge.DeviceListPage, err error)
                ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, expand string) (result databoxedge.DeviceListIterator, err error)
            ListBySubscription(ctx context.Context, expand string) (result databoxedge.DeviceListPage, err error)
                ListBySubscriptionComplete(ctx context.Context, expand string) (result databoxedge.DeviceListIterator, err error)
            ScanForUpdates(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.DevicesScanForUpdatesFuture, err error)
            Update(ctx context.Context, deviceName string, parameters databoxedge.DevicePatch, resourceGroupName string) (result databoxedge.Device, err error)
            UploadCertificate(ctx context.Context, deviceName string, parameters databoxedge.UploadCertificateRequest, resourceGroupName string) (result databoxedge.UploadCertificateResponse, err error)
        }

        var _ DevicesClientAPI = (*databoxedge.DevicesClient)(nil)
        // AlertsClientAPI contains the set of methods on the AlertsClient type.
        type AlertsClientAPI interface {
            Get(ctx context.Context, deviceName string, name string, resourceGroupName string) (result databoxedge.Alert, err error)
            ListByDataBoxEdgeDevice(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.AlertListPage, err error)
                ListByDataBoxEdgeDeviceComplete(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.AlertListIterator, err error)
        }

        var _ AlertsClientAPI = (*databoxedge.AlertsClient)(nil)
        // BandwidthSchedulesClientAPI contains the set of methods on the BandwidthSchedulesClient type.
        type BandwidthSchedulesClientAPI interface {
            CreateOrUpdate(ctx context.Context, deviceName string, name string, parameters databoxedge.BandwidthSchedule, resourceGroupName string) (result databoxedge.BandwidthSchedulesCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, deviceName string, name string, resourceGroupName string) (result databoxedge.BandwidthSchedulesDeleteFuture, err error)
            Get(ctx context.Context, deviceName string, name string, resourceGroupName string) (result databoxedge.BandwidthSchedule, err error)
            ListByDataBoxEdgeDevice(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.BandwidthSchedulesListPage, err error)
                ListByDataBoxEdgeDeviceComplete(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.BandwidthSchedulesListIterator, err error)
        }

        var _ BandwidthSchedulesClientAPI = (*databoxedge.BandwidthSchedulesClient)(nil)
        // JobsClientAPI contains the set of methods on the JobsClient type.
        type JobsClientAPI interface {
            Get(ctx context.Context, deviceName string, name string, resourceGroupName string) (result databoxedge.Job, err error)
        }

        var _ JobsClientAPI = (*databoxedge.JobsClient)(nil)
        // NodesClientAPI contains the set of methods on the NodesClient type.
        type NodesClientAPI interface {
            ListByDataBoxEdgeDevice(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.NodeList, err error)
        }

        var _ NodesClientAPI = (*databoxedge.NodesClient)(nil)
        // OperationsStatusClientAPI contains the set of methods on the OperationsStatusClient type.
        type OperationsStatusClientAPI interface {
            Get(ctx context.Context, deviceName string, name string, resourceGroupName string) (result databoxedge.Job, err error)
        }

        var _ OperationsStatusClientAPI = (*databoxedge.OperationsStatusClient)(nil)
        // OrdersClientAPI contains the set of methods on the OrdersClient type.
        type OrdersClientAPI interface {
            CreateOrUpdate(ctx context.Context, deviceName string, order databoxedge.Order, resourceGroupName string) (result databoxedge.OrdersCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.OrdersDeleteFuture, err error)
            Get(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.Order, err error)
            ListByDataBoxEdgeDevice(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.OrderListPage, err error)
                ListByDataBoxEdgeDeviceComplete(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.OrderListIterator, err error)
        }

        var _ OrdersClientAPI = (*databoxedge.OrdersClient)(nil)
        // RolesClientAPI contains the set of methods on the RolesClient type.
        type RolesClientAPI interface {
            CreateOrUpdate(ctx context.Context, deviceName string, name string, role databoxedge.BasicRole, resourceGroupName string) (result databoxedge.RolesCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, deviceName string, name string, resourceGroupName string) (result databoxedge.RolesDeleteFuture, err error)
            Get(ctx context.Context, deviceName string, name string, resourceGroupName string) (result databoxedge.RoleModel, err error)
            ListByDataBoxEdgeDevice(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.RoleListPage, err error)
                ListByDataBoxEdgeDeviceComplete(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.RoleListIterator, err error)
        }

        var _ RolesClientAPI = (*databoxedge.RolesClient)(nil)
        // SharesClientAPI contains the set of methods on the SharesClient type.
        type SharesClientAPI interface {
            CreateOrUpdate(ctx context.Context, deviceName string, name string, share databoxedge.Share, resourceGroupName string) (result databoxedge.SharesCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, deviceName string, name string, resourceGroupName string) (result databoxedge.SharesDeleteFuture, err error)
            Get(ctx context.Context, deviceName string, name string, resourceGroupName string) (result databoxedge.Share, err error)
            ListByDataBoxEdgeDevice(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.ShareListPage, err error)
                ListByDataBoxEdgeDeviceComplete(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.ShareListIterator, err error)
            Refresh(ctx context.Context, deviceName string, name string, resourceGroupName string) (result databoxedge.SharesRefreshFuture, err error)
        }

        var _ SharesClientAPI = (*databoxedge.SharesClient)(nil)
        // StorageAccountCredentialsClientAPI contains the set of methods on the StorageAccountCredentialsClient type.
        type StorageAccountCredentialsClientAPI interface {
            CreateOrUpdate(ctx context.Context, deviceName string, name string, storageAccountCredential databoxedge.StorageAccountCredential, resourceGroupName string) (result databoxedge.StorageAccountCredentialsCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, deviceName string, name string, resourceGroupName string) (result databoxedge.StorageAccountCredentialsDeleteFuture, err error)
            Get(ctx context.Context, deviceName string, name string, resourceGroupName string) (result databoxedge.StorageAccountCredential, err error)
            ListByDataBoxEdgeDevice(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.StorageAccountCredentialListPage, err error)
                ListByDataBoxEdgeDeviceComplete(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.StorageAccountCredentialListIterator, err error)
        }

        var _ StorageAccountCredentialsClientAPI = (*databoxedge.StorageAccountCredentialsClient)(nil)
        // TriggersClientAPI contains the set of methods on the TriggersClient type.
        type TriggersClientAPI interface {
            CreateOrUpdate(ctx context.Context, deviceName string, name string, trigger databoxedge.BasicTrigger, resourceGroupName string) (result databoxedge.TriggersCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, deviceName string, name string, resourceGroupName string) (result databoxedge.TriggersDeleteFuture, err error)
            Get(ctx context.Context, deviceName string, name string, resourceGroupName string) (result databoxedge.TriggerModel, err error)
            ListByDataBoxEdgeDevice(ctx context.Context, deviceName string, resourceGroupName string, expand string) (result databoxedge.TriggerListPage, err error)
                ListByDataBoxEdgeDeviceComplete(ctx context.Context, deviceName string, resourceGroupName string, expand string) (result databoxedge.TriggerListIterator, err error)
        }

        var _ TriggersClientAPI = (*databoxedge.TriggersClient)(nil)
        // UsersClientAPI contains the set of methods on the UsersClient type.
        type UsersClientAPI interface {
            CreateOrUpdate(ctx context.Context, deviceName string, name string, userParameter databoxedge.User, resourceGroupName string) (result databoxedge.UsersCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, deviceName string, name string, resourceGroupName string) (result databoxedge.UsersDeleteFuture, err error)
            Get(ctx context.Context, deviceName string, name string, resourceGroupName string) (result databoxedge.User, err error)
            ListByDataBoxEdgeDevice(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.UserListPage, err error)
                ListByDataBoxEdgeDeviceComplete(ctx context.Context, deviceName string, resourceGroupName string) (result databoxedge.UserListIterator, err error)
        }

        var _ UsersClientAPI = (*databoxedge.UsersClient)(nil)
