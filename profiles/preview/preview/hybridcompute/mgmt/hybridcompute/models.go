// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package hybridcompute

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/hybridcompute/mgmt/2019-03-18-preview/hybridcompute"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type InstanceViewTypes = original.InstanceViewTypes

const (
	InstanceView InstanceViewTypes = original.InstanceView
)

type StatusTypes = original.StatusTypes

const (
	Connected    StatusTypes = original.Connected
	Disconnected StatusTypes = original.Disconnected
	Error        StatusTypes = original.Error
)

type BaseClient = original.BaseClient
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type Identity = original.Identity
type Machine = original.Machine
type MachineListResult = original.MachineListResult
type MachineListResultIterator = original.MachineListResultIterator
type MachineListResultPage = original.MachineListResultPage
type MachineProperties = original.MachineProperties
type MachineReconnect = original.MachineReconnect
type MachineReconnectProperties = original.MachineReconnectProperties
type MachineUpdate = original.MachineUpdate
type MachineUpdateProperties = original.MachineUpdateProperties
type MachinesClient = original.MachinesClient
type OSProfile = original.OSProfile
type OperationListResult = original.OperationListResult
type OperationValue = original.OperationValue
type OperationValueDisplay = original.OperationValueDisplay
type OperationsClient = original.OperationsClient
type Resource = original.Resource
type UpdateResource = original.UpdateResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewMachineListResultIterator(page MachineListResultPage) MachineListResultIterator {
	return original.NewMachineListResultIterator(page)
}
func NewMachineListResultPage(getNextPage func(context.Context, MachineListResult) (MachineListResult, error)) MachineListResultPage {
	return original.NewMachineListResultPage(getNextPage)
}
func NewMachinesClient(subscriptionID string) MachinesClient {
	return original.NewMachinesClient(subscriptionID)
}
func NewMachinesClientWithBaseURI(baseURI string, subscriptionID string) MachinesClient {
	return original.NewMachinesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleInstanceViewTypesValues() []InstanceViewTypes {
	return original.PossibleInstanceViewTypesValues()
}
func PossibleStatusTypesValues() []StatusTypes {
	return original.PossibleStatusTypesValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
