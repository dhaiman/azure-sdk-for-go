package search

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

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Failed ...
	Failed ProvisioningState = "failed"
	// Provisioning ...
	Provisioning ProvisioningState = "provisioning"
	// Succeeded ...
	Succeeded ProvisioningState = "succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Failed, Provisioning, Succeeded}
}

// ServiceStatus enumerates the values for service status.
type ServiceStatus string

const (
	// ServiceStatusDegraded ...
	ServiceStatusDegraded ServiceStatus = "degraded"
	// ServiceStatusDeleting ...
	ServiceStatusDeleting ServiceStatus = "deleting"
	// ServiceStatusDisabled ...
	ServiceStatusDisabled ServiceStatus = "disabled"
	// ServiceStatusError ...
	ServiceStatusError ServiceStatus = "error"
	// ServiceStatusProvisioning ...
	ServiceStatusProvisioning ServiceStatus = "provisioning"
	// ServiceStatusRunning ...
	ServiceStatusRunning ServiceStatus = "running"
)

// PossibleServiceStatusValues returns an array of possible values for the ServiceStatus const type.
func PossibleServiceStatusValues() []ServiceStatus {
	return []ServiceStatus{ServiceStatusDegraded, ServiceStatusDeleting, ServiceStatusDisabled, ServiceStatusError, ServiceStatusProvisioning, ServiceStatusRunning}
}

// SkuType enumerates the values for sku type.
type SkuType string

const (
	// Free ...
	Free SkuType = "free"
	// Standard ...
	Standard SkuType = "standard"
	// Standard2 ...
	Standard2 SkuType = "standard2"
)

// PossibleSkuTypeValues returns an array of possible values for the SkuType const type.
func PossibleSkuTypeValues() []SkuType {
	return []SkuType{Free, Standard, Standard2}
}
