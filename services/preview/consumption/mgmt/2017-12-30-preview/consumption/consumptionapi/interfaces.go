package consumptionapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/consumption/mgmt/2017-12-30-preview/consumption"
	"github.com/Azure/go-autorest/autorest"
)

// BudgetsClientAPI contains the set of methods on the BudgetsClient type.
type BudgetsClientAPI interface {
	CreateOrUpdate(ctx context.Context, parameters consumption.Budget) (result consumption.Budget, err error)
	Delete(ctx context.Context) (result autorest.Response, err error)
	Get(ctx context.Context) (result consumption.Budget, err error)
	List(ctx context.Context) (result consumption.BudgetsListResult, err error)
}

var _ BudgetsClientAPI = (*consumption.BudgetsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result consumption.OperationListResultPage, err error)
}

var _ OperationsClientAPI = (*consumption.OperationsClient)(nil)
