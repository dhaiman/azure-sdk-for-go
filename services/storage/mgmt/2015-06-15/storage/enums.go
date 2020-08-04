package storage

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

// AccountStatus enumerates the values for account status.
type AccountStatus string

const (
	// Available ...
	Available AccountStatus = "Available"
	// Unavailable ...
	Unavailable AccountStatus = "Unavailable"
)

// PossibleAccountStatusValues returns an array of possible values for the AccountStatus const type.
func PossibleAccountStatusValues() []AccountStatus {
	return []AccountStatus{Available, Unavailable}
}

// AccountType enumerates the values for account type.
type AccountType string

const (
	// PremiumLRS ...
	PremiumLRS AccountType = "Premium_LRS"
	// StandardGRS ...
	StandardGRS AccountType = "Standard_GRS"
	// StandardLRS ...
	StandardLRS AccountType = "Standard_LRS"
	// StandardRAGRS ...
	StandardRAGRS AccountType = "Standard_RAGRS"
	// StandardZRS ...
	StandardZRS AccountType = "Standard_ZRS"
)

// PossibleAccountTypeValues returns an array of possible values for the AccountType const type.
func PossibleAccountTypeValues() []AccountType {
	return []AccountType{PremiumLRS, StandardGRS, StandardLRS, StandardRAGRS, StandardZRS}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Creating ...
	Creating ProvisioningState = "Creating"
	// ResolvingDNS ...
	ResolvingDNS ProvisioningState = "ResolvingDNS"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Creating, ResolvingDNS, Succeeded}
}

// Reason enumerates the values for reason.
type Reason string

const (
	// AccountNameInvalid ...
	AccountNameInvalid Reason = "AccountNameInvalid"
	// AlreadyExists ...
	AlreadyExists Reason = "AlreadyExists"
)

// PossibleReasonValues returns an array of possible values for the Reason const type.
func PossibleReasonValues() []Reason {
	return []Reason{AccountNameInvalid, AlreadyExists}
}

// UsageUnit enumerates the values for usage unit.
type UsageUnit string

const (
	// Bytes ...
	Bytes UsageUnit = "Bytes"
	// BytesPerSecond ...
	BytesPerSecond UsageUnit = "BytesPerSecond"
	// Count ...
	Count UsageUnit = "Count"
	// CountsPerSecond ...
	CountsPerSecond UsageUnit = "CountsPerSecond"
	// Percent ...
	Percent UsageUnit = "Percent"
	// Seconds ...
	Seconds UsageUnit = "Seconds"
)

// PossibleUsageUnitValues returns an array of possible values for the UsageUnit const type.
func PossibleUsageUnitValues() []UsageUnit {
	return []UsageUnit{Bytes, BytesPerSecond, Count, CountsPerSecond, Percent, Seconds}
}
