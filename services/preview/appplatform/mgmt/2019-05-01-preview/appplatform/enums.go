package appplatform

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

// AppResourceProvisioningState enumerates the values for app resource provisioning state.
type AppResourceProvisioningState string

const (
	// Creating ...
	Creating AppResourceProvisioningState = "Creating"
	// Failed ...
	Failed AppResourceProvisioningState = "Failed"
	// Succeeded ...
	Succeeded AppResourceProvisioningState = "Succeeded"
	// Updating ...
	Updating AppResourceProvisioningState = "Updating"
)

// PossibleAppResourceProvisioningStateValues returns an array of possible values for the AppResourceProvisioningState const type.
func PossibleAppResourceProvisioningStateValues() []AppResourceProvisioningState {
	return []AppResourceProvisioningState{Creating, Failed, Succeeded, Updating}
}

// ConfigServerState enumerates the values for config server state.
type ConfigServerState string

const (
	// ConfigServerStateDeleted ...
	ConfigServerStateDeleted ConfigServerState = "Deleted"
	// ConfigServerStateFailed ...
	ConfigServerStateFailed ConfigServerState = "Failed"
	// ConfigServerStateNotAvailable ...
	ConfigServerStateNotAvailable ConfigServerState = "NotAvailable"
	// ConfigServerStateSucceeded ...
	ConfigServerStateSucceeded ConfigServerState = "Succeeded"
	// ConfigServerStateUpdating ...
	ConfigServerStateUpdating ConfigServerState = "Updating"
)

// PossibleConfigServerStateValues returns an array of possible values for the ConfigServerState const type.
func PossibleConfigServerStateValues() []ConfigServerState {
	return []ConfigServerState{ConfigServerStateDeleted, ConfigServerStateFailed, ConfigServerStateNotAvailable, ConfigServerStateSucceeded, ConfigServerStateUpdating}
}

// DeploymentResourceProvisioningState enumerates the values for deployment resource provisioning state.
type DeploymentResourceProvisioningState string

const (
	// DeploymentResourceProvisioningStateCreating ...
	DeploymentResourceProvisioningStateCreating DeploymentResourceProvisioningState = "Creating"
	// DeploymentResourceProvisioningStateFailed ...
	DeploymentResourceProvisioningStateFailed DeploymentResourceProvisioningState = "Failed"
	// DeploymentResourceProvisioningStateSucceeded ...
	DeploymentResourceProvisioningStateSucceeded DeploymentResourceProvisioningState = "Succeeded"
	// DeploymentResourceProvisioningStateUpdating ...
	DeploymentResourceProvisioningStateUpdating DeploymentResourceProvisioningState = "Updating"
)

// PossibleDeploymentResourceProvisioningStateValues returns an array of possible values for the DeploymentResourceProvisioningState const type.
func PossibleDeploymentResourceProvisioningStateValues() []DeploymentResourceProvisioningState {
	return []DeploymentResourceProvisioningState{DeploymentResourceProvisioningStateCreating, DeploymentResourceProvisioningStateFailed, DeploymentResourceProvisioningStateSucceeded, DeploymentResourceProvisioningStateUpdating}
}

// DeploymentResourceStatus enumerates the values for deployment resource status.
type DeploymentResourceStatus string

const (
	// DeploymentResourceStatusAllocating ...
	DeploymentResourceStatusAllocating DeploymentResourceStatus = "Allocating"
	// DeploymentResourceStatusCompiling ...
	DeploymentResourceStatusCompiling DeploymentResourceStatus = "Compiling"
	// DeploymentResourceStatusFailed ...
	DeploymentResourceStatusFailed DeploymentResourceStatus = "Failed"
	// DeploymentResourceStatusRunning ...
	DeploymentResourceStatusRunning DeploymentResourceStatus = "Running"
	// DeploymentResourceStatusStopped ...
	DeploymentResourceStatusStopped DeploymentResourceStatus = "Stopped"
	// DeploymentResourceStatusUnknown ...
	DeploymentResourceStatusUnknown DeploymentResourceStatus = "Unknown"
	// DeploymentResourceStatusUpgrading ...
	DeploymentResourceStatusUpgrading DeploymentResourceStatus = "Upgrading"
)

// PossibleDeploymentResourceStatusValues returns an array of possible values for the DeploymentResourceStatus const type.
func PossibleDeploymentResourceStatusValues() []DeploymentResourceStatus {
	return []DeploymentResourceStatus{DeploymentResourceStatusAllocating, DeploymentResourceStatusCompiling, DeploymentResourceStatusFailed, DeploymentResourceStatusRunning, DeploymentResourceStatusStopped, DeploymentResourceStatusUnknown, DeploymentResourceStatusUpgrading}
}

// ManagedIdentityType enumerates the values for managed identity type.
type ManagedIdentityType string

const (
	// None ...
	None ManagedIdentityType = "None"
	// SystemAssigned ...
	SystemAssigned ManagedIdentityType = "SystemAssigned"
	// SystemAssignedUserAssigned ...
	SystemAssignedUserAssigned ManagedIdentityType = "SystemAssigned,UserAssigned"
	// UserAssigned ...
	UserAssigned ManagedIdentityType = "UserAssigned"
)

// PossibleManagedIdentityTypeValues returns an array of possible values for the ManagedIdentityType const type.
func PossibleManagedIdentityTypeValues() []ManagedIdentityType {
	return []ManagedIdentityType{None, SystemAssigned, SystemAssignedUserAssigned, UserAssigned}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateCreating ...
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleted ...
	ProvisioningStateDeleted ProvisioningState = "Deleted"
	// ProvisioningStateDeleting ...
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed ...
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateMoved ...
	ProvisioningStateMoved ProvisioningState = "Moved"
	// ProvisioningStateMoveFailed ...
	ProvisioningStateMoveFailed ProvisioningState = "MoveFailed"
	// ProvisioningStateMoving ...
	ProvisioningStateMoving ProvisioningState = "Moving"
	// ProvisioningStateSucceeded ...
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating ...
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{ProvisioningStateCreating, ProvisioningStateDeleted, ProvisioningStateDeleting, ProvisioningStateFailed, ProvisioningStateMoved, ProvisioningStateMoveFailed, ProvisioningStateMoving, ProvisioningStateSucceeded, ProvisioningStateUpdating}
}

// ResourceSkuRestrictionsReasonCode enumerates the values for resource sku restrictions reason code.
type ResourceSkuRestrictionsReasonCode string

const (
	// NotAvailableForSubscription ...
	NotAvailableForSubscription ResourceSkuRestrictionsReasonCode = "NotAvailableForSubscription"
	// QuotaID ...
	QuotaID ResourceSkuRestrictionsReasonCode = "QuotaId"
)

// PossibleResourceSkuRestrictionsReasonCodeValues returns an array of possible values for the ResourceSkuRestrictionsReasonCode const type.
func PossibleResourceSkuRestrictionsReasonCodeValues() []ResourceSkuRestrictionsReasonCode {
	return []ResourceSkuRestrictionsReasonCode{NotAvailableForSubscription, QuotaID}
}

// ResourceSkuRestrictionsType enumerates the values for resource sku restrictions type.
type ResourceSkuRestrictionsType string

const (
	// Location ...
	Location ResourceSkuRestrictionsType = "Location"
	// Zone ...
	Zone ResourceSkuRestrictionsType = "Zone"
)

// PossibleResourceSkuRestrictionsTypeValues returns an array of possible values for the ResourceSkuRestrictionsType const type.
func PossibleResourceSkuRestrictionsTypeValues() []ResourceSkuRestrictionsType {
	return []ResourceSkuRestrictionsType{Location, Zone}
}

// RuntimeVersion enumerates the values for runtime version.
type RuntimeVersion string

const (
	// Java11 ...
	Java11 RuntimeVersion = "Java_11"
	// Java8 ...
	Java8 RuntimeVersion = "Java_8"
	// NetCore31 ...
	NetCore31 RuntimeVersion = "NetCore_31"
)

// PossibleRuntimeVersionValues returns an array of possible values for the RuntimeVersion const type.
func PossibleRuntimeVersionValues() []RuntimeVersion {
	return []RuntimeVersion{Java11, Java8, NetCore31}
}

// SkuScaleType enumerates the values for sku scale type.
type SkuScaleType string

const (
	// SkuScaleTypeAutomatic ...
	SkuScaleTypeAutomatic SkuScaleType = "Automatic"
	// SkuScaleTypeManual ...
	SkuScaleTypeManual SkuScaleType = "Manual"
	// SkuScaleTypeNone ...
	SkuScaleTypeNone SkuScaleType = "None"
)

// PossibleSkuScaleTypeValues returns an array of possible values for the SkuScaleType const type.
func PossibleSkuScaleTypeValues() []SkuScaleType {
	return []SkuScaleType{SkuScaleTypeAutomatic, SkuScaleTypeManual, SkuScaleTypeNone}
}

// SupportedRuntimePlatform enumerates the values for supported runtime platform.
type SupportedRuntimePlatform string

const (
	// Java ...
	Java SupportedRuntimePlatform = "Java"
	// NETCore ...
	NETCore SupportedRuntimePlatform = ".NET Core"
)

// PossibleSupportedRuntimePlatformValues returns an array of possible values for the SupportedRuntimePlatform const type.
func PossibleSupportedRuntimePlatformValues() []SupportedRuntimePlatform {
	return []SupportedRuntimePlatform{Java, NETCore}
}

// SupportedRuntimeVersion enumerates the values for supported runtime version.
type SupportedRuntimeVersion string

const (
	// SupportedRuntimeVersionJava11 ...
	SupportedRuntimeVersionJava11 SupportedRuntimeVersion = "Java_11"
	// SupportedRuntimeVersionJava8 ...
	SupportedRuntimeVersionJava8 SupportedRuntimeVersion = "Java_8"
	// SupportedRuntimeVersionNetCore31 ...
	SupportedRuntimeVersionNetCore31 SupportedRuntimeVersion = "NetCore_31"
)

// PossibleSupportedRuntimeVersionValues returns an array of possible values for the SupportedRuntimeVersion const type.
func PossibleSupportedRuntimeVersionValues() []SupportedRuntimeVersion {
	return []SupportedRuntimeVersion{SupportedRuntimeVersionJava11, SupportedRuntimeVersionJava8, SupportedRuntimeVersionNetCore31}
}

// TestKeyType enumerates the values for test key type.
type TestKeyType string

const (
	// Primary ...
	Primary TestKeyType = "Primary"
	// Secondary ...
	Secondary TestKeyType = "Secondary"
)

// PossibleTestKeyTypeValues returns an array of possible values for the TestKeyType const type.
func PossibleTestKeyTypeValues() []TestKeyType {
	return []TestKeyType{Primary, Secondary}
}

// TraceProxyState enumerates the values for trace proxy state.
type TraceProxyState string

const (
	// TraceProxyStateFailed ...
	TraceProxyStateFailed TraceProxyState = "Failed"
	// TraceProxyStateNotAvailable ...
	TraceProxyStateNotAvailable TraceProxyState = "NotAvailable"
	// TraceProxyStateSucceeded ...
	TraceProxyStateSucceeded TraceProxyState = "Succeeded"
	// TraceProxyStateUpdating ...
	TraceProxyStateUpdating TraceProxyState = "Updating"
)

// PossibleTraceProxyStateValues returns an array of possible values for the TraceProxyState const type.
func PossibleTraceProxyStateValues() []TraceProxyState {
	return []TraceProxyState{TraceProxyStateFailed, TraceProxyStateNotAvailable, TraceProxyStateSucceeded, TraceProxyStateUpdating}
}

// UserSourceType enumerates the values for user source type.
type UserSourceType string

const (
	// Jar ...
	Jar UserSourceType = "Jar"
	// NetCoreZip ...
	NetCoreZip UserSourceType = "NetCoreZip"
	// Source ...
	Source UserSourceType = "Source"
)

// PossibleUserSourceTypeValues returns an array of possible values for the UserSourceType const type.
func PossibleUserSourceTypeValues() []UserSourceType {
	return []UserSourceType{Jar, NetCoreZip, Source}
}
