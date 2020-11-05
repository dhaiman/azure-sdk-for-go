package operationalinsights

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

// BillingType enumerates the values for billing type.
type BillingType string

const (
	// BillingTypeCluster ...
	BillingTypeCluster BillingType = "Cluster"
	// BillingTypeWorkspaces ...
	BillingTypeWorkspaces BillingType = "Workspaces"
)

// PossibleBillingTypeValues returns an array of possible values for the BillingType const type.
func PossibleBillingTypeValues() []BillingType {
	return []BillingType{BillingTypeCluster, BillingTypeWorkspaces}
}

// ClusterEntityStatus enumerates the values for cluster entity status.
type ClusterEntityStatus string

const (
	// Canceled ...
	Canceled ClusterEntityStatus = "Canceled"
	// Creating ...
	Creating ClusterEntityStatus = "Creating"
	// Deleting ...
	Deleting ClusterEntityStatus = "Deleting"
	// Failed ...
	Failed ClusterEntityStatus = "Failed"
	// ProvisioningAccount ...
	ProvisioningAccount ClusterEntityStatus = "ProvisioningAccount"
	// Succeeded ...
	Succeeded ClusterEntityStatus = "Succeeded"
	// Updating ...
	Updating ClusterEntityStatus = "Updating"
)

// PossibleClusterEntityStatusValues returns an array of possible values for the ClusterEntityStatus const type.
func PossibleClusterEntityStatusValues() []ClusterEntityStatus {
	return []ClusterEntityStatus{Canceled, Creating, Deleting, Failed, ProvisioningAccount, Succeeded, Updating}
}

// ClusterSkuNameEnum enumerates the values for cluster sku name enum.
type ClusterSkuNameEnum string

const (
	// CapacityReservation ...
	CapacityReservation ClusterSkuNameEnum = "CapacityReservation"
)

// PossibleClusterSkuNameEnumValues returns an array of possible values for the ClusterSkuNameEnum const type.
func PossibleClusterSkuNameEnumValues() []ClusterSkuNameEnum {
	return []ClusterSkuNameEnum{CapacityReservation}
}

// IdentityType enumerates the values for identity type.
type IdentityType string

const (
	// None ...
	None IdentityType = "None"
	// SystemAssigned ...
	SystemAssigned IdentityType = "SystemAssigned"
	// UserAssigned ...
	UserAssigned IdentityType = "UserAssigned"
)

// PossibleIdentityTypeValues returns an array of possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{None, SystemAssigned, UserAssigned}
}
