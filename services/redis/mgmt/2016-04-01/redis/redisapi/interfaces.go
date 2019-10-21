package redisapi

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
    "github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2016-04-01/redis"
    "github.com/Azure/go-autorest/autorest"
)

        // OperationsClientAPI contains the set of methods on the OperationsClient type.
        type OperationsClientAPI interface {
            List(ctx context.Context) (result redis.OperationListResultPage, err error)
        }

        var _ OperationsClientAPI = (*redis.OperationsClient)(nil)
        // ClientAPI contains the set of methods on the Client type.
        type ClientAPI interface {
            Create(ctx context.Context, resourceGroupName string, name string, parameters redis.CreateParameters) (result redis.CreateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, name string) (result redis.DeleteFuture, err error)
            ExportData(ctx context.Context, resourceGroupName string, name string, parameters redis.ExportRDBParameters) (result redis.ExportDataFuture, err error)
            ForceReboot(ctx context.Context, resourceGroupName string, name string, parameters redis.RebootParameters) (result redis.ForceRebootResponse, err error)
            Get(ctx context.Context, resourceGroupName string, name string) (result redis.ResourceType, err error)
            ImportData(ctx context.Context, resourceGroupName string, name string, parameters redis.ImportRDBParameters) (result redis.ImportDataFuture, err error)
            List(ctx context.Context) (result redis.ListResultPage, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result redis.ListResultPage, err error)
            ListKeys(ctx context.Context, resourceGroupName string, name string) (result redis.AccessKeys, err error)
            RegenerateKey(ctx context.Context, resourceGroupName string, name string, parameters redis.RegenerateKeyParameters) (result redis.AccessKeys, err error)
            Update(ctx context.Context, resourceGroupName string, name string, parameters redis.UpdateParameters) (result redis.ResourceType, err error)
        }

        var _ ClientAPI = (*redis.Client)(nil)
        // FirewallRulesClientAPI contains the set of methods on the FirewallRulesClient type.
        type FirewallRulesClientAPI interface {
            List(ctx context.Context, resourceGroupName string, cacheName string) (result redis.FirewallRuleListResultPage, err error)
        }

        var _ FirewallRulesClientAPI = (*redis.FirewallRulesClient)(nil)
        // FirewallRuleClientAPI contains the set of methods on the FirewallRuleClient type.
        type FirewallRuleClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, cacheName string, ruleName string, parameters redis.FirewallRule) (result redis.FirewallRule, err error)
            Delete(ctx context.Context, resourceGroupName string, cacheName string, ruleName string) (result autorest.Response, err error)
            Get(ctx context.Context, resourceGroupName string, cacheName string, ruleName string) (result redis.FirewallRule, err error)
        }

        var _ FirewallRuleClientAPI = (*redis.FirewallRuleClient)(nil)
        // PatchSchedulesClientAPI contains the set of methods on the PatchSchedulesClient type.
        type PatchSchedulesClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, name string, parameters redis.PatchSchedule) (result redis.PatchSchedule, err error)
            Delete(ctx context.Context, resourceGroupName string, name string) (result autorest.Response, err error)
            Get(ctx context.Context, resourceGroupName string, name string) (result redis.PatchSchedule, err error)
        }

        var _ PatchSchedulesClientAPI = (*redis.PatchSchedulesClient)(nil)
