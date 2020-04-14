package managedapplicationsapi

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
    "github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2016-09-01-preview/managedapplications"
)

        // BaseClientAPI contains the set of methods on the BaseClient type.
        type BaseClientAPI interface {
            ListOperations(ctx context.Context) (result managedapplications.OperationListResultPage, err error)
                ListOperationsComplete(ctx context.Context) (result managedapplications.OperationListResultIterator, err error)
        }

        var _ BaseClientAPI = (*managedapplications.BaseClient)(nil)
        // AppliancesClientAPI contains the set of methods on the AppliancesClient type.
        type AppliancesClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, applianceName string, parameters managedapplications.Appliance) (result managedapplications.AppliancesCreateOrUpdateFuture, err error)
            CreateOrUpdateByID(ctx context.Context, applianceID string, parameters managedapplications.Appliance) (result managedapplications.AppliancesCreateOrUpdateByIDFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, applianceName string) (result managedapplications.AppliancesDeleteFuture, err error)
            DeleteByID(ctx context.Context, applianceID string) (result managedapplications.AppliancesDeleteByIDFuture, err error)
            Get(ctx context.Context, resourceGroupName string, applianceName string) (result managedapplications.Appliance, err error)
            GetByID(ctx context.Context, applianceID string) (result managedapplications.Appliance, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result managedapplications.ApplianceListResultPage, err error)
                ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result managedapplications.ApplianceListResultIterator, err error)
            ListBySubscription(ctx context.Context) (result managedapplications.ApplianceListResultPage, err error)
                ListBySubscriptionComplete(ctx context.Context) (result managedapplications.ApplianceListResultIterator, err error)
            Update(ctx context.Context, resourceGroupName string, applianceName string, parameters *managedapplications.Appliance) (result managedapplications.Appliance, err error)
            UpdateByID(ctx context.Context, applianceID string, parameters *managedapplications.Appliance) (result managedapplications.Appliance, err error)
        }

        var _ AppliancesClientAPI = (*managedapplications.AppliancesClient)(nil)
        // ApplianceDefinitionsClientAPI contains the set of methods on the ApplianceDefinitionsClient type.
        type ApplianceDefinitionsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, applianceDefinitionName string, parameters managedapplications.ApplianceDefinition) (result managedapplications.ApplianceDefinitionsCreateOrUpdateFuture, err error)
            CreateOrUpdateByID(ctx context.Context, applianceDefinitionID string, parameters managedapplications.ApplianceDefinition) (result managedapplications.ApplianceDefinitionsCreateOrUpdateByIDFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, applianceDefinitionName string) (result managedapplications.ApplianceDefinitionsDeleteFuture, err error)
            DeleteByID(ctx context.Context, applianceDefinitionID string) (result managedapplications.ApplianceDefinitionsDeleteByIDFuture, err error)
            Get(ctx context.Context, resourceGroupName string, applianceDefinitionName string) (result managedapplications.ApplianceDefinition, err error)
            GetByID(ctx context.Context, applianceDefinitionID string) (result managedapplications.ApplianceDefinition, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result managedapplications.ApplianceDefinitionListResultPage, err error)
                ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result managedapplications.ApplianceDefinitionListResultIterator, err error)
        }

        var _ ApplianceDefinitionsClientAPI = (*managedapplications.ApplianceDefinitionsClient)(nil)
