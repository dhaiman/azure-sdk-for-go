package analysisservicesapi

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
	"github.com/Azure/azure-sdk-for-go/services/analysisservices/mgmt/2017-07-14/analysisservices"
	"github.com/Azure/go-autorest/autorest"
)

// ServersClientAPI contains the set of methods on the ServersClient type.
type ServersClientAPI interface {
	CheckNameAvailability(ctx context.Context, location string, serverParameters analysisservices.CheckServerNameAvailabilityParameters) (result analysisservices.CheckServerNameAvailabilityResult, err error)
	Create(ctx context.Context, resourceGroupName string, serverName string, serverParameters analysisservices.Server) (result analysisservices.ServersCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serverName string) (result analysisservices.ServersDeleteFuture, err error)
	DissociateGateway(ctx context.Context, resourceGroupName string, serverName string) (result autorest.Response, err error)
	GetDetails(ctx context.Context, resourceGroupName string, serverName string) (result analysisservices.Server, err error)
	List(ctx context.Context) (result analysisservices.Servers, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result analysisservices.Servers, err error)
	ListGatewayStatus(ctx context.Context, resourceGroupName string, serverName string) (result analysisservices.GatewayListStatusLive, err error)
	ListOperationResults(ctx context.Context, location string, operationID string) (result autorest.Response, err error)
	ListOperationStatuses(ctx context.Context, location string, operationID string) (result analysisservices.OperationStatus, err error)
	ListSkusForExisting(ctx context.Context, resourceGroupName string, serverName string) (result analysisservices.SkuEnumerationForExistingResourceResult, err error)
	ListSkusForNew(ctx context.Context) (result analysisservices.SkuEnumerationForNewResourceResult, err error)
	Resume(ctx context.Context, resourceGroupName string, serverName string) (result analysisservices.ServersResumeFuture, err error)
	Suspend(ctx context.Context, resourceGroupName string, serverName string) (result analysisservices.ServersSuspendFuture, err error)
	Update(ctx context.Context, resourceGroupName string, serverName string, serverUpdateParameters analysisservices.ServerUpdateParameters) (result analysisservices.ServersUpdateFuture, err error)
}

var _ ServersClientAPI = (*analysisservices.ServersClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result analysisservices.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result analysisservices.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*analysisservices.OperationsClient)(nil)
