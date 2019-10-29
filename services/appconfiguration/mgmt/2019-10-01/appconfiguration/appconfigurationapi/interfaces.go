package appconfigurationapi

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
    "github.com/Azure/azure-sdk-for-go/services/appconfiguration/mgmt/2019-10-01/appconfiguration"
)

        // ConfigurationStoresClientAPI contains the set of methods on the ConfigurationStoresClient type.
        type ConfigurationStoresClientAPI interface {
            Create(ctx context.Context, resourceGroupName string, configStoreName string, configStoreCreationParameters appconfiguration.ConfigurationStore) (result appconfiguration.ConfigurationStoresCreateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, configStoreName string) (result appconfiguration.ConfigurationStoresDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, configStoreName string) (result appconfiguration.ConfigurationStore, err error)
            List(ctx context.Context, skipToken string) (result appconfiguration.ConfigurationStoreListResultPage, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string, skipToken string) (result appconfiguration.ConfigurationStoreListResultPage, err error)
            ListKeys(ctx context.Context, resourceGroupName string, configStoreName string, skipToken string) (result appconfiguration.APIKeyListResultPage, err error)
            ListKeyValue(ctx context.Context, resourceGroupName string, configStoreName string, listKeyValueParameters appconfiguration.ListKeyValueParameters) (result appconfiguration.KeyValue, err error)
            RegenerateKey(ctx context.Context, resourceGroupName string, configStoreName string, regenerateKeyParameters appconfiguration.RegenerateKeyParameters) (result appconfiguration.APIKey, err error)
            Update(ctx context.Context, resourceGroupName string, configStoreName string, configStoreUpdateParameters appconfiguration.ConfigurationStoreUpdateParameters) (result appconfiguration.ConfigurationStoresUpdateFuture, err error)
        }

        var _ ConfigurationStoresClientAPI = (*appconfiguration.ConfigurationStoresClient)(nil)
        // OperationsClientAPI contains the set of methods on the OperationsClient type.
        type OperationsClientAPI interface {
            CheckNameAvailability(ctx context.Context, checkNameAvailabilityParameters appconfiguration.CheckNameAvailabilityParameters) (result appconfiguration.NameAvailabilityStatus, err error)
            List(ctx context.Context, skipToken string) (result appconfiguration.OperationDefinitionListResultPage, err error)
        }

        var _ OperationsClientAPI = (*appconfiguration.OperationsClient)(nil)
