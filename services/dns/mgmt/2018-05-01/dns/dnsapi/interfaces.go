package dnsapi

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
    "github.com/Azure/azure-sdk-for-go/services/dns/mgmt/2018-05-01/dns"
    "github.com/Azure/go-autorest/autorest"
)

        // RecordSetsClientAPI contains the set of methods on the RecordSetsClient type.
        type RecordSetsClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, zoneName string, relativeRecordSetName string, recordType dns.RecordType, parameters dns.RecordSet, ifMatch string, ifNoneMatch string) (result dns.RecordSet, err error)
            Delete(ctx context.Context, resourceGroupName string, zoneName string, relativeRecordSetName string, recordType dns.RecordType, ifMatch string) (result autorest.Response, err error)
            Get(ctx context.Context, resourceGroupName string, zoneName string, relativeRecordSetName string, recordType dns.RecordType) (result dns.RecordSet, err error)
            ListAllByDNSZone(ctx context.Context, resourceGroupName string, zoneName string, top *int32, recordSetNameSuffix string) (result dns.RecordSetListResultPage, err error)
                ListAllByDNSZoneComplete(ctx context.Context, resourceGroupName string, zoneName string, top *int32, recordSetNameSuffix string) (result dns.RecordSetListResultIterator, err error)
            ListByDNSZone(ctx context.Context, resourceGroupName string, zoneName string, top *int32, recordsetnamesuffix string) (result dns.RecordSetListResultPage, err error)
                ListByDNSZoneComplete(ctx context.Context, resourceGroupName string, zoneName string, top *int32, recordsetnamesuffix string) (result dns.RecordSetListResultIterator, err error)
            ListByType(ctx context.Context, resourceGroupName string, zoneName string, recordType dns.RecordType, top *int32, recordsetnamesuffix string) (result dns.RecordSetListResultPage, err error)
                ListByTypeComplete(ctx context.Context, resourceGroupName string, zoneName string, recordType dns.RecordType, top *int32, recordsetnamesuffix string) (result dns.RecordSetListResultIterator, err error)
            Update(ctx context.Context, resourceGroupName string, zoneName string, relativeRecordSetName string, recordType dns.RecordType, parameters dns.RecordSet, ifMatch string) (result dns.RecordSet, err error)
        }

        var _ RecordSetsClientAPI = (*dns.RecordSetsClient)(nil)
        // ZonesClientAPI contains the set of methods on the ZonesClient type.
        type ZonesClientAPI interface {
            CreateOrUpdate(ctx context.Context, resourceGroupName string, zoneName string, parameters dns.Zone, ifMatch string, ifNoneMatch string) (result dns.Zone, err error)
            Delete(ctx context.Context, resourceGroupName string, zoneName string, ifMatch string) (result dns.ZonesDeleteFuture, err error)
            Get(ctx context.Context, resourceGroupName string, zoneName string) (result dns.Zone, err error)
            List(ctx context.Context, top *int32) (result dns.ZoneListResultPage, err error)
                ListComplete(ctx context.Context, top *int32) (result dns.ZoneListResultIterator, err error)
            ListByResourceGroup(ctx context.Context, resourceGroupName string, top *int32) (result dns.ZoneListResultPage, err error)
                ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, top *int32) (result dns.ZoneListResultIterator, err error)
            Update(ctx context.Context, resourceGroupName string, zoneName string, parameters dns.ZoneUpdate, ifMatch string) (result dns.Zone, err error)
        }

        var _ ZonesClientAPI = (*dns.ZonesClient)(nil)
        // ResourceReferenceClientAPI contains the set of methods on the ResourceReferenceClient type.
        type ResourceReferenceClientAPI interface {
            GetByTargetResources(ctx context.Context, parameters dns.ResourceReferenceRequest) (result dns.ResourceReferenceResult, err error)
        }

        var _ ResourceReferenceClientAPI = (*dns.ResourceReferenceClient)(nil)
