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

package maintenance

import original "github.com/Azure/azure-sdk-for-go/services/preview/maintenance/mgmt/2018-06-01-preview/maintenance"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ImpactType = original.ImpactType

const (
	Freeze   ImpactType = original.Freeze
	None     ImpactType = original.None
	Redeploy ImpactType = original.Redeploy
	Restart  ImpactType = original.Restart
)

type Scope = original.Scope

const (
	ScopeAll        Scope = original.ScopeAll
	ScopeHost       Scope = original.ScopeHost
	ScopeInResource Scope = original.ScopeInResource
	ScopeResource   Scope = original.ScopeResource
)

type UpdateStatus = original.UpdateStatus

const (
	Completed  UpdateStatus = original.Completed
	InProgress UpdateStatus = original.InProgress
	Pending    UpdateStatus = original.Pending
	RetryLater UpdateStatus = original.RetryLater
	RetryNow   UpdateStatus = original.RetryNow
)

type ApplyUpdate = original.ApplyUpdate
type ApplyUpdateProperties = original.ApplyUpdateProperties
type ApplyUpdatesClient = original.ApplyUpdatesClient
type BaseClient = original.BaseClient
type Configuration = original.Configuration
type ConfigurationAssignment = original.ConfigurationAssignment
type ConfigurationAssignmentProperties = original.ConfigurationAssignmentProperties
type ConfigurationAssignmentsClient = original.ConfigurationAssignmentsClient
type ConfigurationProperties = original.ConfigurationProperties
type ConfigurationsClient = original.ConfigurationsClient
type Error = original.Error
type ErrorDetails = original.ErrorDetails
type ListConfigurationAssignmentsResult = original.ListConfigurationAssignmentsResult
type ListMaintenanceConfigurationsResult = original.ListMaintenanceConfigurationsResult
type ListUpdatesResult = original.ListUpdatesResult
type Operation = original.Operation
type OperationInfo = original.OperationInfo
type OperationsClient = original.OperationsClient
type OperationsListResult = original.OperationsListResult
type Resource = original.Resource
type Update = original.Update
type UpdateProperties = original.UpdateProperties
type UpdatesClient = original.UpdatesClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewApplyUpdatesClient(subscriptionID string) ApplyUpdatesClient {
	return original.NewApplyUpdatesClient(subscriptionID)
}
func NewApplyUpdatesClientWithBaseURI(baseURI string, subscriptionID string) ApplyUpdatesClient {
	return original.NewApplyUpdatesClientWithBaseURI(baseURI, subscriptionID)
}
func NewConfigurationAssignmentsClient(subscriptionID string) ConfigurationAssignmentsClient {
	return original.NewConfigurationAssignmentsClient(subscriptionID)
}
func NewConfigurationAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationAssignmentsClient {
	return original.NewConfigurationAssignmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewConfigurationsClient(subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClient(subscriptionID)
}
func NewConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationsClient {
	return original.NewConfigurationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewUpdatesClient(subscriptionID string) UpdatesClient {
	return original.NewUpdatesClient(subscriptionID)
}
func NewUpdatesClientWithBaseURI(baseURI string, subscriptionID string) UpdatesClient {
	return original.NewUpdatesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleImpactTypeValues() []ImpactType {
	return original.PossibleImpactTypeValues()
}
func PossibleScopeValues() []Scope {
	return original.PossibleScopeValues()
}
func PossibleUpdateStatusValues() []UpdateStatus {
	return original.PossibleUpdateStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
