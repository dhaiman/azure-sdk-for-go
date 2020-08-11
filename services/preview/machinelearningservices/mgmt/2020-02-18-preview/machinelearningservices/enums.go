package machinelearningservices

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

// AllocationState enumerates the values for allocation state.
type AllocationState string

const (
	// Resizing ...
	Resizing AllocationState = "Resizing"
	// Steady ...
	Steady AllocationState = "Steady"
)

// PossibleAllocationStateValues returns an array of possible values for the AllocationState const type.
func PossibleAllocationStateValues() []AllocationState {
	return []AllocationState{Resizing, Steady}
}

// ApplicationSharingPolicy enumerates the values for application sharing policy.
type ApplicationSharingPolicy string

const (
	// Personal ...
	Personal ApplicationSharingPolicy = "Personal"
	// Shared ...
	Shared ApplicationSharingPolicy = "Shared"
)

// PossibleApplicationSharingPolicyValues returns an array of possible values for the ApplicationSharingPolicy const type.
func PossibleApplicationSharingPolicyValues() []ApplicationSharingPolicy {
	return []ApplicationSharingPolicy{Personal, Shared}
}

// ComputeInstanceState enumerates the values for compute instance state.
type ComputeInstanceState string

const (
	// CreateFailed ...
	CreateFailed ComputeInstanceState = "CreateFailed"
	// Creating ...
	Creating ComputeInstanceState = "Creating"
	// Deleting ...
	Deleting ComputeInstanceState = "Deleting"
	// JobRunning ...
	JobRunning ComputeInstanceState = "JobRunning"
	// RestartFailed ...
	RestartFailed ComputeInstanceState = "RestartFailed"
	// Restarting ...
	Restarting ComputeInstanceState = "Restarting"
	// Running ...
	Running ComputeInstanceState = "Running"
	// SettingUp ...
	SettingUp ComputeInstanceState = "SettingUp"
	// StartFailed ...
	StartFailed ComputeInstanceState = "StartFailed"
	// Starting ...
	Starting ComputeInstanceState = "Starting"
	// StopFailed ...
	StopFailed ComputeInstanceState = "StopFailed"
	// Stopped ...
	Stopped ComputeInstanceState = "Stopped"
	// Stopping ...
	Stopping ComputeInstanceState = "Stopping"
	// Unknown ...
	Unknown ComputeInstanceState = "Unknown"
	// Unusable ...
	Unusable ComputeInstanceState = "Unusable"
	// UserSettingUp ...
	UserSettingUp ComputeInstanceState = "UserSettingUp"
)

// PossibleComputeInstanceStateValues returns an array of possible values for the ComputeInstanceState const type.
func PossibleComputeInstanceStateValues() []ComputeInstanceState {
	return []ComputeInstanceState{CreateFailed, Creating, Deleting, JobRunning, RestartFailed, Restarting, Running, SettingUp, StartFailed, Starting, StopFailed, Stopped, Stopping, Unknown, Unusable, UserSettingUp}
}

// ComputeType enumerates the values for compute type.
type ComputeType string

const (
	// ComputeTypeAKS ...
	ComputeTypeAKS ComputeType = "AKS"
	// ComputeTypeAmlCompute ...
	ComputeTypeAmlCompute ComputeType = "AmlCompute"
	// ComputeTypeComputeInstance ...
	ComputeTypeComputeInstance ComputeType = "ComputeInstance"
	// ComputeTypeDatabricks ...
	ComputeTypeDatabricks ComputeType = "Databricks"
	// ComputeTypeDataFactory ...
	ComputeTypeDataFactory ComputeType = "DataFactory"
	// ComputeTypeDataLakeAnalytics ...
	ComputeTypeDataLakeAnalytics ComputeType = "DataLakeAnalytics"
	// ComputeTypeHDInsight ...
	ComputeTypeHDInsight ComputeType = "HDInsight"
	// ComputeTypeVirtualMachine ...
	ComputeTypeVirtualMachine ComputeType = "VirtualMachine"
)

// PossibleComputeTypeValues returns an array of possible values for the ComputeType const type.
func PossibleComputeTypeValues() []ComputeType {
	return []ComputeType{ComputeTypeAKS, ComputeTypeAmlCompute, ComputeTypeComputeInstance, ComputeTypeDatabricks, ComputeTypeDataFactory, ComputeTypeDataLakeAnalytics, ComputeTypeHDInsight, ComputeTypeVirtualMachine}
}

// ComputeTypeBasicCompute enumerates the values for compute type basic compute.
type ComputeTypeBasicCompute string

const (
	// ComputeTypeAKS1 ...
	ComputeTypeAKS1 ComputeTypeBasicCompute = "AKS"
	// ComputeTypeAmlCompute1 ...
	ComputeTypeAmlCompute1 ComputeTypeBasicCompute = "AmlCompute"
	// ComputeTypeCompute ...
	ComputeTypeCompute ComputeTypeBasicCompute = "Compute"
	// ComputeTypeComputeInstance1 ...
	ComputeTypeComputeInstance1 ComputeTypeBasicCompute = "ComputeInstance"
	// ComputeTypeDatabricks1 ...
	ComputeTypeDatabricks1 ComputeTypeBasicCompute = "Databricks"
	// ComputeTypeDataFactory1 ...
	ComputeTypeDataFactory1 ComputeTypeBasicCompute = "DataFactory"
	// ComputeTypeDataLakeAnalytics1 ...
	ComputeTypeDataLakeAnalytics1 ComputeTypeBasicCompute = "DataLakeAnalytics"
	// ComputeTypeHDInsight1 ...
	ComputeTypeHDInsight1 ComputeTypeBasicCompute = "HDInsight"
	// ComputeTypeVirtualMachine1 ...
	ComputeTypeVirtualMachine1 ComputeTypeBasicCompute = "VirtualMachine"
)

// PossibleComputeTypeBasicComputeValues returns an array of possible values for the ComputeTypeBasicCompute const type.
func PossibleComputeTypeBasicComputeValues() []ComputeTypeBasicCompute {
	return []ComputeTypeBasicCompute{ComputeTypeAKS1, ComputeTypeAmlCompute1, ComputeTypeCompute, ComputeTypeComputeInstance1, ComputeTypeDatabricks1, ComputeTypeDataFactory1, ComputeTypeDataLakeAnalytics1, ComputeTypeHDInsight1, ComputeTypeVirtualMachine1}
}

// ComputeTypeBasicComputeNodesInformation enumerates the values for compute type basic compute nodes
// information.
type ComputeTypeBasicComputeNodesInformation string

const (
	// ComputeTypeBasicComputeNodesInformationComputeTypeAmlCompute ...
	ComputeTypeBasicComputeNodesInformationComputeTypeAmlCompute ComputeTypeBasicComputeNodesInformation = "AmlCompute"
	// ComputeTypeBasicComputeNodesInformationComputeTypeComputeNodesInformation ...
	ComputeTypeBasicComputeNodesInformationComputeTypeComputeNodesInformation ComputeTypeBasicComputeNodesInformation = "ComputeNodesInformation"
)

// PossibleComputeTypeBasicComputeNodesInformationValues returns an array of possible values for the ComputeTypeBasicComputeNodesInformation const type.
func PossibleComputeTypeBasicComputeNodesInformationValues() []ComputeTypeBasicComputeNodesInformation {
	return []ComputeTypeBasicComputeNodesInformation{ComputeTypeBasicComputeNodesInformationComputeTypeAmlCompute, ComputeTypeBasicComputeNodesInformationComputeTypeComputeNodesInformation}
}

// ComputeTypeBasicComputeSecrets enumerates the values for compute type basic compute secrets.
type ComputeTypeBasicComputeSecrets string

const (
	// ComputeTypeBasicComputeSecretsComputeTypeAKS ...
	ComputeTypeBasicComputeSecretsComputeTypeAKS ComputeTypeBasicComputeSecrets = "AKS"
	// ComputeTypeBasicComputeSecretsComputeTypeComputeSecrets ...
	ComputeTypeBasicComputeSecretsComputeTypeComputeSecrets ComputeTypeBasicComputeSecrets = "ComputeSecrets"
	// ComputeTypeBasicComputeSecretsComputeTypeDatabricks ...
	ComputeTypeBasicComputeSecretsComputeTypeDatabricks ComputeTypeBasicComputeSecrets = "Databricks"
	// ComputeTypeBasicComputeSecretsComputeTypeVirtualMachine ...
	ComputeTypeBasicComputeSecretsComputeTypeVirtualMachine ComputeTypeBasicComputeSecrets = "VirtualMachine"
)

// PossibleComputeTypeBasicComputeSecretsValues returns an array of possible values for the ComputeTypeBasicComputeSecrets const type.
func PossibleComputeTypeBasicComputeSecretsValues() []ComputeTypeBasicComputeSecrets {
	return []ComputeTypeBasicComputeSecrets{ComputeTypeBasicComputeSecretsComputeTypeAKS, ComputeTypeBasicComputeSecretsComputeTypeComputeSecrets, ComputeTypeBasicComputeSecretsComputeTypeDatabricks, ComputeTypeBasicComputeSecretsComputeTypeVirtualMachine}
}

// EncryptionStatus enumerates the values for encryption status.
type EncryptionStatus string

const (
	// Disabled ...
	Disabled EncryptionStatus = "Disabled"
	// Enabled ...
	Enabled EncryptionStatus = "Enabled"
)

// PossibleEncryptionStatusValues returns an array of possible values for the EncryptionStatus const type.
func PossibleEncryptionStatusValues() []EncryptionStatus {
	return []EncryptionStatus{Disabled, Enabled}
}

// NodeState enumerates the values for node state.
type NodeState string

const (
	// NodeStateIdle ...
	NodeStateIdle NodeState = "idle"
	// NodeStateLeaving ...
	NodeStateLeaving NodeState = "leaving"
	// NodeStatePreempted ...
	NodeStatePreempted NodeState = "preempted"
	// NodeStatePreparing ...
	NodeStatePreparing NodeState = "preparing"
	// NodeStateRunning ...
	NodeStateRunning NodeState = "running"
	// NodeStateUnusable ...
	NodeStateUnusable NodeState = "unusable"
)

// PossibleNodeStateValues returns an array of possible values for the NodeState const type.
func PossibleNodeStateValues() []NodeState {
	return []NodeState{NodeStateIdle, NodeStateLeaving, NodeStatePreempted, NodeStatePreparing, NodeStateRunning, NodeStateUnusable}
}

// OsType enumerates the values for os type.
type OsType string

const (
	// Linux ...
	Linux OsType = "Linux"
	// Windows ...
	Windows OsType = "Windows"
)

// PossibleOsTypeValues returns an array of possible values for the OsType const type.
func PossibleOsTypeValues() []OsType {
	return []OsType{Linux, Windows}
}

// PrivateEndpointConnectionProvisioningState enumerates the values for private endpoint connection
// provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	// PrivateEndpointConnectionProvisioningStateCreating ...
	PrivateEndpointConnectionProvisioningStateCreating PrivateEndpointConnectionProvisioningState = "Creating"
	// PrivateEndpointConnectionProvisioningStateDeleting ...
	PrivateEndpointConnectionProvisioningStateDeleting PrivateEndpointConnectionProvisioningState = "Deleting"
	// PrivateEndpointConnectionProvisioningStateFailed ...
	PrivateEndpointConnectionProvisioningStateFailed PrivateEndpointConnectionProvisioningState = "Failed"
	// PrivateEndpointConnectionProvisioningStateSucceeded ...
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns an array of possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{PrivateEndpointConnectionProvisioningStateCreating, PrivateEndpointConnectionProvisioningStateDeleting, PrivateEndpointConnectionProvisioningStateFailed, PrivateEndpointConnectionProvisioningStateSucceeded}
}

// PrivateEndpointServiceConnectionStatus enumerates the values for private endpoint service connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	// Approved ...
	Approved PrivateEndpointServiceConnectionStatus = "Approved"
	// Pending ...
	Pending PrivateEndpointServiceConnectionStatus = "Pending"
	// Rejected ...
	Rejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns an array of possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{Approved, Pending, Rejected}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateCanceled ...
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateCreating ...
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleting ...
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed ...
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateSucceeded ...
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUnknown ...
	ProvisioningStateUnknown ProvisioningState = "Unknown"
	// ProvisioningStateUpdating ...
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{ProvisioningStateCanceled, ProvisioningStateCreating, ProvisioningStateDeleting, ProvisioningStateFailed, ProvisioningStateSucceeded, ProvisioningStateUnknown, ProvisioningStateUpdating}
}

// QuotaUnit enumerates the values for quota unit.
type QuotaUnit string

const (
	// Count ...
	Count QuotaUnit = "Count"
)

// PossibleQuotaUnitValues returns an array of possible values for the QuotaUnit const type.
func PossibleQuotaUnitValues() []QuotaUnit {
	return []QuotaUnit{Count}
}

// ReasonCode enumerates the values for reason code.
type ReasonCode string

const (
	// NotAvailableForRegion ...
	NotAvailableForRegion ReasonCode = "NotAvailableForRegion"
	// NotAvailableForSubscription ...
	NotAvailableForSubscription ReasonCode = "NotAvailableForSubscription"
	// NotSpecified ...
	NotSpecified ReasonCode = "NotSpecified"
)

// PossibleReasonCodeValues returns an array of possible values for the ReasonCode const type.
func PossibleReasonCodeValues() []ReasonCode {
	return []ReasonCode{NotAvailableForRegion, NotAvailableForSubscription, NotSpecified}
}

// RemoteLoginPortPublicAccess enumerates the values for remote login port public access.
type RemoteLoginPortPublicAccess string

const (
	// RemoteLoginPortPublicAccessDisabled ...
	RemoteLoginPortPublicAccessDisabled RemoteLoginPortPublicAccess = "Disabled"
	// RemoteLoginPortPublicAccessEnabled ...
	RemoteLoginPortPublicAccessEnabled RemoteLoginPortPublicAccess = "Enabled"
	// RemoteLoginPortPublicAccessNotSpecified ...
	RemoteLoginPortPublicAccessNotSpecified RemoteLoginPortPublicAccess = "NotSpecified"
)

// PossibleRemoteLoginPortPublicAccessValues returns an array of possible values for the RemoteLoginPortPublicAccess const type.
func PossibleRemoteLoginPortPublicAccessValues() []RemoteLoginPortPublicAccess {
	return []RemoteLoginPortPublicAccess{RemoteLoginPortPublicAccessDisabled, RemoteLoginPortPublicAccessEnabled, RemoteLoginPortPublicAccessNotSpecified}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// SystemAssigned ...
	SystemAssigned ResourceIdentityType = "SystemAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{SystemAssigned}
}

// SSHPublicAccess enumerates the values for ssh public access.
type SSHPublicAccess string

const (
	// SSHPublicAccessDisabled ...
	SSHPublicAccessDisabled SSHPublicAccess = "Disabled"
	// SSHPublicAccessEnabled ...
	SSHPublicAccessEnabled SSHPublicAccess = "Enabled"
)

// PossibleSSHPublicAccessValues returns an array of possible values for the SSHPublicAccess const type.
func PossibleSSHPublicAccessValues() []SSHPublicAccess {
	return []SSHPublicAccess{SSHPublicAccessDisabled, SSHPublicAccessEnabled}
}

// Status enumerates the values for status.
type Status string

const (
	// Failure ...
	Failure Status = "Failure"
	// InvalidQuotaBelowClusterMinimum ...
	InvalidQuotaBelowClusterMinimum Status = "InvalidQuotaBelowClusterMinimum"
	// InvalidQuotaExceedsSubscriptionLimit ...
	InvalidQuotaExceedsSubscriptionLimit Status = "InvalidQuotaExceedsSubscriptionLimit"
	// InvalidVMFamilyName ...
	InvalidVMFamilyName Status = "InvalidVMFamilyName"
	// OperationNotEnabledForRegion ...
	OperationNotEnabledForRegion Status = "OperationNotEnabledForRegion"
	// OperationNotSupportedForSku ...
	OperationNotSupportedForSku Status = "OperationNotSupportedForSku"
	// Success ...
	Success Status = "Success"
	// Undefined ...
	Undefined Status = "Undefined"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{Failure, InvalidQuotaBelowClusterMinimum, InvalidQuotaExceedsSubscriptionLimit, InvalidVMFamilyName, OperationNotEnabledForRegion, OperationNotSupportedForSku, Success, Undefined}
}

// Status1 enumerates the values for status 1.
type Status1 string

const (
	// Status1Disabled ...
	Status1Disabled Status1 = "Disabled"
	// Status1Enabled ...
	Status1Enabled Status1 = "Enabled"
)

// PossibleStatus1Values returns an array of possible values for the Status1 const type.
func PossibleStatus1Values() []Status1 {
	return []Status1{Status1Disabled, Status1Enabled}
}

// UnderlyingResourceAction enumerates the values for underlying resource action.
type UnderlyingResourceAction string

const (
	// Delete ...
	Delete UnderlyingResourceAction = "Delete"
	// Detach ...
	Detach UnderlyingResourceAction = "Detach"
)

// PossibleUnderlyingResourceActionValues returns an array of possible values for the UnderlyingResourceAction const type.
func PossibleUnderlyingResourceActionValues() []UnderlyingResourceAction {
	return []UnderlyingResourceAction{Delete, Detach}
}

// UsageUnit enumerates the values for usage unit.
type UsageUnit string

const (
	// UsageUnitCount ...
	UsageUnitCount UsageUnit = "Count"
)

// PossibleUsageUnitValues returns an array of possible values for the UsageUnit const type.
func PossibleUsageUnitValues() []UsageUnit {
	return []UsageUnit{UsageUnitCount}
}

// VMPriority enumerates the values for vm priority.
type VMPriority string

const (
	// Dedicated ...
	Dedicated VMPriority = "Dedicated"
	// LowPriority ...
	LowPriority VMPriority = "LowPriority"
)

// PossibleVMPriorityValues returns an array of possible values for the VMPriority const type.
func PossibleVMPriorityValues() []VMPriority {
	return []VMPriority{Dedicated, LowPriority}
}
