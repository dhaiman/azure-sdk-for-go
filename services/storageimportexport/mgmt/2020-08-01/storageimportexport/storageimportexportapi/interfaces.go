package storageimportexportapi

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
	"github.com/Azure/azure-sdk-for-go/services/storageimportexport/mgmt/2020-08-01/storageimportexport"
	"github.com/Azure/go-autorest/autorest"
)

// LocationsClientAPI contains the set of methods on the LocationsClient type.
type LocationsClientAPI interface {
	Get(ctx context.Context, locationName string) (result storageimportexport.Location, err error)
	List(ctx context.Context) (result storageimportexport.LocationsResponse, err error)
}

var _ LocationsClientAPI = (*storageimportexport.LocationsClient)(nil)

// JobsClientAPI contains the set of methods on the JobsClient type.
type JobsClientAPI interface {
	Create(ctx context.Context, jobName string, resourceGroupName string, body storageimportexport.PutJobParameters, clientTenantID string) (result storageimportexport.JobResponse, err error)
	Delete(ctx context.Context, jobName string, resourceGroupName string) (result autorest.Response, err error)
	Get(ctx context.Context, jobName string, resourceGroupName string) (result storageimportexport.JobResponse, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, top *int32, filter string) (result storageimportexport.ListJobsResponsePage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, top *int32, filter string) (result storageimportexport.ListJobsResponseIterator, err error)
	ListBySubscription(ctx context.Context, top *int32, filter string) (result storageimportexport.ListJobsResponsePage, err error)
	ListBySubscriptionComplete(ctx context.Context, top *int32, filter string) (result storageimportexport.ListJobsResponseIterator, err error)
	Update(ctx context.Context, jobName string, resourceGroupName string, body storageimportexport.UpdateJobParameters) (result storageimportexport.JobResponse, err error)
}

var _ JobsClientAPI = (*storageimportexport.JobsClient)(nil)

// BitLockerKeysClientAPI contains the set of methods on the BitLockerKeysClient type.
type BitLockerKeysClientAPI interface {
	List(ctx context.Context, jobName string, resourceGroupName string) (result storageimportexport.GetBitLockerKeysResponse, err error)
}

var _ BitLockerKeysClientAPI = (*storageimportexport.BitLockerKeysClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result storageimportexport.ListOperationsResponse, err error)
}

var _ OperationsClientAPI = (*storageimportexport.OperationsClient)(nil)
