package quotaapi

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
	"github.com/Azure/azure-sdk-for-go/services/quota/mgmt/2020-10-25/quota"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	CreateOrUpdate(ctx context.Context, subscriptionID string, providerID string, location string, resourceName string, createQuotaRequest quota.CurrentQuotaLimitBase) (result quota.CreateOrUpdateFuture, err error)
	Get(ctx context.Context, subscriptionID string, providerID string, location string, resourceName string) (result quota.CurrentQuotaLimitBase, err error)
	List(ctx context.Context, subscriptionID string, providerID string, location string) (result quota.LimitsPage, err error)
	ListComplete(ctx context.Context, subscriptionID string, providerID string, location string) (result quota.LimitsIterator, err error)
	Update(ctx context.Context, subscriptionID string, providerID string, location string, resourceName string, createQuotaRequest quota.CurrentQuotaLimitBase) (result quota.UpdateFuture, err error)
}

var _ ClientAPI = (*quota.Client)(nil)

// RequestStatusClientAPI contains the set of methods on the RequestStatusClient type.
type RequestStatusClientAPI interface {
	Get(ctx context.Context, subscriptionID string, providerID string, location string, ID string) (result quota.RequestDetails, err error)
	List(ctx context.Context, subscriptionID string, providerID string, location string, filter string, top *int32, skiptoken string) (result quota.RequestDetailsListPage, err error)
	ListComplete(ctx context.Context, subscriptionID string, providerID string, location string, filter string, top *int32, skiptoken string) (result quota.RequestDetailsListIterator, err error)
}

var _ RequestStatusClientAPI = (*quota.RequestStatusClient)(nil)
