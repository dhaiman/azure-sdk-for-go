package frontdoorapi

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
    "github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2019-10-01/frontdoor"
)

        // BaseClientAPI contains the set of methods on the BaseClient type.
        type BaseClientAPI interface {
            CheckFrontDoorNameAvailability(ctx context.Context, checkFrontDoorNameAvailabilityInput frontdoor.CheckNameAvailabilityInput) (result frontdoor.CheckNameAvailabilityOutput, err error)
            CheckFrontDoorNameAvailabilityWithSubscription(ctx context.Context, checkFrontDoorNameAvailabilityInput frontdoor.CheckNameAvailabilityInput) (result frontdoor.CheckNameAvailabilityOutput, err error)
        }

        var _ BaseClientAPI = (*frontdoor.BaseClient)(nil)
        // FrontDoorsClientAPI contains the set of methods on the FrontDoorsClient type.
        type FrontDoorsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, frontDoorName string, frontDoorParameters frontdoor.FrontDoor) (result frontdoor.FrontDoorsCreateOrUpdateFutureType, err error)
            Delete(ctx context.Context, resourceGroupName string, frontDoorName string) (result frontdoor.FrontDoorsDeleteFutureType, err error)
            Get(ctx context.Context, resourceGroupName string, frontDoorName string) (result frontdoor.FrontDoor, err error)
            List(ctx context.Context) (result frontdoor.ListResultPage, err error)
                ListComplete(ctx context.Context) (result frontdoor.ListResultIterator, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string) (result frontdoor.ListResultPage, err error)
                ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result frontdoor.ListResultIterator, err error)
            ValidateCustomDomain(ctx context.Context, resourceGroupName string, frontDoorName string, customDomainProperties frontdoor.ValidateCustomDomainInput) (result frontdoor.ValidateCustomDomainOutput, err error)
        }

        var _ FrontDoorsClientAPI = (*frontdoor.FrontDoorsClient)(nil)
        // FrontendEndpointsClientAPI contains the set of methods on the FrontendEndpointsClient type.
        type FrontendEndpointsClientAPI interface {
            DisableHTTPS(ctx context.Context, resourceGroupName string, frontDoorName string, frontendEndpointName string) (result frontdoor.FrontendEndpointsDisableHTTPSFuture, err error)
            EnableHTTPS(ctx context.Context, resourceGroupName string, frontDoorName string, frontendEndpointName string, customHTTPSConfiguration frontdoor.CustomHTTPSConfiguration) (result frontdoor.FrontendEndpointsEnableHTTPSFuture, err error)
            Get(ctx context.Context, resourceGroupName string, frontDoorName string, frontendEndpointName string) (result frontdoor.FrontendEndpoint, err error)
            ListByFrontDoor(ctx context.Context, resourceGroupName string, frontDoorName string) (result frontdoor.FrontendEndpointsListResultPage, err error)
                ListByFrontDoorComplete(ctx context.Context, resourceGroupName string, frontDoorName string) (result frontdoor.FrontendEndpointsListResultIterator, err error)
        }

        var _ FrontendEndpointsClientAPI = (*frontdoor.FrontendEndpointsClient)(nil)
        // EndpointsClientAPI contains the set of methods on the EndpointsClient type.
        type EndpointsClientAPI interface {
            PurgeContent(ctx context.Context, resourceGroupName string, frontDoorName string, contentFilePaths frontdoor.PurgeParameters) (result frontdoor.EndpointsPurgeContentFuture, err error)
        }

        var _ EndpointsClientAPI = (*frontdoor.EndpointsClient)(nil)
        // PoliciesClientAPI contains the set of methods on the PoliciesClient type.
        type PoliciesClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, policyName string, parameters frontdoor.WebApplicationFirewallPolicy) (result frontdoor.PoliciesCreateOrUpdateFuture, err error)
            Delete(ctx context.Context, resourceGroupName string, policyName string) (result frontdoor.PoliciesDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, policyName string) (result frontdoor.WebApplicationFirewallPolicy, err error)
            List(ctx context.Context, resourceGroupName string) (result frontdoor.WebApplicationFirewallPolicyListPage, err error)
                ListComplete(ctx context.Context, resourceGroupName string) (result frontdoor.WebApplicationFirewallPolicyListIterator, err error)
        }

        var _ PoliciesClientAPI = (*frontdoor.PoliciesClient)(nil)
        // ManagedRuleSetsClientAPI contains the set of methods on the ManagedRuleSetsClient type.
        type ManagedRuleSetsClientAPI interface {
            List(ctx context.Context) (result frontdoor.ManagedRuleSetDefinitionListPage, err error)
                ListComplete(ctx context.Context) (result frontdoor.ManagedRuleSetDefinitionListIterator, err error)
        }

        var _ ManagedRuleSetsClientAPI = (*frontdoor.ManagedRuleSetsClient)(nil)
