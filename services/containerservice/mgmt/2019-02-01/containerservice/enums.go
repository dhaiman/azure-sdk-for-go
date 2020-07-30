package containerservice

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

// AgentPoolType enumerates the values for agent pool type.
type AgentPoolType string

const (
	// AvailabilitySet ...
	AvailabilitySet AgentPoolType = "AvailabilitySet"
	// VirtualMachineScaleSets ...
	VirtualMachineScaleSets AgentPoolType = "VirtualMachineScaleSets"
)

// PossibleAgentPoolTypeValues returns an array of possible values for the AgentPoolType const type.
func PossibleAgentPoolTypeValues() []AgentPoolType {
	return []AgentPoolType{AvailabilitySet, VirtualMachineScaleSets}
}

// Format enumerates the values for format.
type Format string

const (
	// Azure ...
	Azure Format = "azure"
	// Exec ...
	Exec Format = "exec"
)

// PossibleFormatValues returns an array of possible values for the Format const type.
func PossibleFormatValues() []Format {
	return []Format{Azure, Exec}
}

// Kind enumerates the values for kind.
type Kind string

const (
	// KindAADIdentityProvider ...
	KindAADIdentityProvider Kind = "AADIdentityProvider"
	// KindOpenShiftManagedClusterBaseIdentityProvider ...
	KindOpenShiftManagedClusterBaseIdentityProvider Kind = "OpenShiftManagedClusterBaseIdentityProvider"
)

// PossibleKindValues returns an array of possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{KindAADIdentityProvider, KindOpenShiftManagedClusterBaseIdentityProvider}
}

// Login enumerates the values for login.
type Login string

const (
	// Devicecode ...
	Devicecode Login = "devicecode"
	// Msi ...
	Msi Login = "msi"
	// Spn ...
	Spn Login = "spn"
)

// PossibleLoginValues returns an array of possible values for the Login const type.
func PossibleLoginValues() []Login {
	return []Login{Devicecode, Msi, Spn}
}

// NetworkPlugin enumerates the values for network plugin.
type NetworkPlugin string

const (
	// NetworkPluginAzure ...
	NetworkPluginAzure NetworkPlugin = "azure"
	// NetworkPluginKubenet ...
	NetworkPluginKubenet NetworkPlugin = "kubenet"
)

// PossibleNetworkPluginValues returns an array of possible values for the NetworkPlugin const type.
func PossibleNetworkPluginValues() []NetworkPlugin {
	return []NetworkPlugin{NetworkPluginAzure, NetworkPluginKubenet}
}

// NetworkPolicy enumerates the values for network policy.
type NetworkPolicy string

const (
	// NetworkPolicyAzure ...
	NetworkPolicyAzure NetworkPolicy = "azure"
	// NetworkPolicyCalico ...
	NetworkPolicyCalico NetworkPolicy = "calico"
)

// PossibleNetworkPolicyValues returns an array of possible values for the NetworkPolicy const type.
func PossibleNetworkPolicyValues() []NetworkPolicy {
	return []NetworkPolicy{NetworkPolicyAzure, NetworkPolicyCalico}
}

// OpenShiftAgentPoolProfileRole enumerates the values for open shift agent pool profile role.
type OpenShiftAgentPoolProfileRole string

const (
	// Compute ...
	Compute OpenShiftAgentPoolProfileRole = "compute"
	// Infra ...
	Infra OpenShiftAgentPoolProfileRole = "infra"
)

// PossibleOpenShiftAgentPoolProfileRoleValues returns an array of possible values for the OpenShiftAgentPoolProfileRole const type.
func PossibleOpenShiftAgentPoolProfileRoleValues() []OpenShiftAgentPoolProfileRole {
	return []OpenShiftAgentPoolProfileRole{Compute, Infra}
}

// OpenShiftContainerServiceVMSize enumerates the values for open shift container service vm size.
type OpenShiftContainerServiceVMSize string

const (
	// StandardD16sV3 ...
	StandardD16sV3 OpenShiftContainerServiceVMSize = "Standard_D16s_v3"
	// StandardD2sV3 ...
	StandardD2sV3 OpenShiftContainerServiceVMSize = "Standard_D2s_v3"
	// StandardD32sV3 ...
	StandardD32sV3 OpenShiftContainerServiceVMSize = "Standard_D32s_v3"
	// StandardD4sV3 ...
	StandardD4sV3 OpenShiftContainerServiceVMSize = "Standard_D4s_v3"
	// StandardD64sV3 ...
	StandardD64sV3 OpenShiftContainerServiceVMSize = "Standard_D64s_v3"
	// StandardD8sV3 ...
	StandardD8sV3 OpenShiftContainerServiceVMSize = "Standard_D8s_v3"
	// StandardDS12V2 ...
	StandardDS12V2 OpenShiftContainerServiceVMSize = "Standard_DS12_v2"
	// StandardDS13V2 ...
	StandardDS13V2 OpenShiftContainerServiceVMSize = "Standard_DS13_v2"
	// StandardDS14V2 ...
	StandardDS14V2 OpenShiftContainerServiceVMSize = "Standard_DS14_v2"
	// StandardDS15V2 ...
	StandardDS15V2 OpenShiftContainerServiceVMSize = "Standard_DS15_v2"
	// StandardDS4V2 ...
	StandardDS4V2 OpenShiftContainerServiceVMSize = "Standard_DS4_v2"
	// StandardDS5V2 ...
	StandardDS5V2 OpenShiftContainerServiceVMSize = "Standard_DS5_v2"
	// StandardE16sV3 ...
	StandardE16sV3 OpenShiftContainerServiceVMSize = "Standard_E16s_v3"
	// StandardE20sV3 ...
	StandardE20sV3 OpenShiftContainerServiceVMSize = "Standard_E20s_v3"
	// StandardE32sV3 ...
	StandardE32sV3 OpenShiftContainerServiceVMSize = "Standard_E32s_v3"
	// StandardE4sV3 ...
	StandardE4sV3 OpenShiftContainerServiceVMSize = "Standard_E4s_v3"
	// StandardE64sV3 ...
	StandardE64sV3 OpenShiftContainerServiceVMSize = "Standard_E64s_v3"
	// StandardE8sV3 ...
	StandardE8sV3 OpenShiftContainerServiceVMSize = "Standard_E8s_v3"
	// StandardF16s ...
	StandardF16s OpenShiftContainerServiceVMSize = "Standard_F16s"
	// StandardF16sV2 ...
	StandardF16sV2 OpenShiftContainerServiceVMSize = "Standard_F16s_v2"
	// StandardF32sV2 ...
	StandardF32sV2 OpenShiftContainerServiceVMSize = "Standard_F32s_v2"
	// StandardF64sV2 ...
	StandardF64sV2 OpenShiftContainerServiceVMSize = "Standard_F64s_v2"
	// StandardF72sV2 ...
	StandardF72sV2 OpenShiftContainerServiceVMSize = "Standard_F72s_v2"
	// StandardF8s ...
	StandardF8s OpenShiftContainerServiceVMSize = "Standard_F8s"
	// StandardF8sV2 ...
	StandardF8sV2 OpenShiftContainerServiceVMSize = "Standard_F8s_v2"
	// StandardGS2 ...
	StandardGS2 OpenShiftContainerServiceVMSize = "Standard_GS2"
	// StandardGS3 ...
	StandardGS3 OpenShiftContainerServiceVMSize = "Standard_GS3"
	// StandardGS4 ...
	StandardGS4 OpenShiftContainerServiceVMSize = "Standard_GS4"
	// StandardGS5 ...
	StandardGS5 OpenShiftContainerServiceVMSize = "Standard_GS5"
	// StandardL16s ...
	StandardL16s OpenShiftContainerServiceVMSize = "Standard_L16s"
	// StandardL32s ...
	StandardL32s OpenShiftContainerServiceVMSize = "Standard_L32s"
	// StandardL4s ...
	StandardL4s OpenShiftContainerServiceVMSize = "Standard_L4s"
	// StandardL8s ...
	StandardL8s OpenShiftContainerServiceVMSize = "Standard_L8s"
)

// PossibleOpenShiftContainerServiceVMSizeValues returns an array of possible values for the OpenShiftContainerServiceVMSize const type.
func PossibleOpenShiftContainerServiceVMSizeValues() []OpenShiftContainerServiceVMSize {
	return []OpenShiftContainerServiceVMSize{StandardD16sV3, StandardD2sV3, StandardD32sV3, StandardD4sV3, StandardD64sV3, StandardD8sV3, StandardDS12V2, StandardDS13V2, StandardDS14V2, StandardDS15V2, StandardDS4V2, StandardDS5V2, StandardE16sV3, StandardE20sV3, StandardE32sV3, StandardE4sV3, StandardE64sV3, StandardE8sV3, StandardF16s, StandardF16sV2, StandardF32sV2, StandardF64sV2, StandardF72sV2, StandardF8s, StandardF8sV2, StandardGS2, StandardGS3, StandardGS4, StandardGS5, StandardL16s, StandardL32s, StandardL4s, StandardL8s}
}

// OrchestratorTypes enumerates the values for orchestrator types.
type OrchestratorTypes string

const (
	// Custom ...
	Custom OrchestratorTypes = "Custom"
	// DCOS ...
	DCOS OrchestratorTypes = "DCOS"
	// DockerCE ...
	DockerCE OrchestratorTypes = "DockerCE"
	// Kubernetes ...
	Kubernetes OrchestratorTypes = "Kubernetes"
	// Swarm ...
	Swarm OrchestratorTypes = "Swarm"
)

// PossibleOrchestratorTypesValues returns an array of possible values for the OrchestratorTypes const type.
func PossibleOrchestratorTypesValues() []OrchestratorTypes {
	return []OrchestratorTypes{Custom, DCOS, DockerCE, Kubernetes, Swarm}
}

// OSType enumerates the values for os type.
type OSType string

const (
	// Linux ...
	Linux OSType = "Linux"
	// Windows ...
	Windows OSType = "Windows"
)

// PossibleOSTypeValues returns an array of possible values for the OSType const type.
func PossibleOSTypeValues() []OSType {
	return []OSType{Linux, Windows}
}

// StorageProfileTypes enumerates the values for storage profile types.
type StorageProfileTypes string

const (
	// ManagedDisks ...
	ManagedDisks StorageProfileTypes = "ManagedDisks"
	// StorageAccount ...
	StorageAccount StorageProfileTypes = "StorageAccount"
)

// PossibleStorageProfileTypesValues returns an array of possible values for the StorageProfileTypes const type.
func PossibleStorageProfileTypesValues() []StorageProfileTypes {
	return []StorageProfileTypes{ManagedDisks, StorageAccount}
}

// VMSizeTypes enumerates the values for vm size types.
type VMSizeTypes string

const (
	// VMSizeTypesStandardA1 ...
	VMSizeTypesStandardA1 VMSizeTypes = "Standard_A1"
	// VMSizeTypesStandardA10 ...
	VMSizeTypesStandardA10 VMSizeTypes = "Standard_A10"
	// VMSizeTypesStandardA11 ...
	VMSizeTypesStandardA11 VMSizeTypes = "Standard_A11"
	// VMSizeTypesStandardA1V2 ...
	VMSizeTypesStandardA1V2 VMSizeTypes = "Standard_A1_v2"
	// VMSizeTypesStandardA2 ...
	VMSizeTypesStandardA2 VMSizeTypes = "Standard_A2"
	// VMSizeTypesStandardA2mV2 ...
	VMSizeTypesStandardA2mV2 VMSizeTypes = "Standard_A2m_v2"
	// VMSizeTypesStandardA2V2 ...
	VMSizeTypesStandardA2V2 VMSizeTypes = "Standard_A2_v2"
	// VMSizeTypesStandardA3 ...
	VMSizeTypesStandardA3 VMSizeTypes = "Standard_A3"
	// VMSizeTypesStandardA4 ...
	VMSizeTypesStandardA4 VMSizeTypes = "Standard_A4"
	// VMSizeTypesStandardA4mV2 ...
	VMSizeTypesStandardA4mV2 VMSizeTypes = "Standard_A4m_v2"
	// VMSizeTypesStandardA4V2 ...
	VMSizeTypesStandardA4V2 VMSizeTypes = "Standard_A4_v2"
	// VMSizeTypesStandardA5 ...
	VMSizeTypesStandardA5 VMSizeTypes = "Standard_A5"
	// VMSizeTypesStandardA6 ...
	VMSizeTypesStandardA6 VMSizeTypes = "Standard_A6"
	// VMSizeTypesStandardA7 ...
	VMSizeTypesStandardA7 VMSizeTypes = "Standard_A7"
	// VMSizeTypesStandardA8 ...
	VMSizeTypesStandardA8 VMSizeTypes = "Standard_A8"
	// VMSizeTypesStandardA8mV2 ...
	VMSizeTypesStandardA8mV2 VMSizeTypes = "Standard_A8m_v2"
	// VMSizeTypesStandardA8V2 ...
	VMSizeTypesStandardA8V2 VMSizeTypes = "Standard_A8_v2"
	// VMSizeTypesStandardA9 ...
	VMSizeTypesStandardA9 VMSizeTypes = "Standard_A9"
	// VMSizeTypesStandardB2ms ...
	VMSizeTypesStandardB2ms VMSizeTypes = "Standard_B2ms"
	// VMSizeTypesStandardB2s ...
	VMSizeTypesStandardB2s VMSizeTypes = "Standard_B2s"
	// VMSizeTypesStandardB4ms ...
	VMSizeTypesStandardB4ms VMSizeTypes = "Standard_B4ms"
	// VMSizeTypesStandardB8ms ...
	VMSizeTypesStandardB8ms VMSizeTypes = "Standard_B8ms"
	// VMSizeTypesStandardD1 ...
	VMSizeTypesStandardD1 VMSizeTypes = "Standard_D1"
	// VMSizeTypesStandardD11 ...
	VMSizeTypesStandardD11 VMSizeTypes = "Standard_D11"
	// VMSizeTypesStandardD11V2 ...
	VMSizeTypesStandardD11V2 VMSizeTypes = "Standard_D11_v2"
	// VMSizeTypesStandardD11V2Promo ...
	VMSizeTypesStandardD11V2Promo VMSizeTypes = "Standard_D11_v2_Promo"
	// VMSizeTypesStandardD12 ...
	VMSizeTypesStandardD12 VMSizeTypes = "Standard_D12"
	// VMSizeTypesStandardD12V2 ...
	VMSizeTypesStandardD12V2 VMSizeTypes = "Standard_D12_v2"
	// VMSizeTypesStandardD12V2Promo ...
	VMSizeTypesStandardD12V2Promo VMSizeTypes = "Standard_D12_v2_Promo"
	// VMSizeTypesStandardD13 ...
	VMSizeTypesStandardD13 VMSizeTypes = "Standard_D13"
	// VMSizeTypesStandardD13V2 ...
	VMSizeTypesStandardD13V2 VMSizeTypes = "Standard_D13_v2"
	// VMSizeTypesStandardD13V2Promo ...
	VMSizeTypesStandardD13V2Promo VMSizeTypes = "Standard_D13_v2_Promo"
	// VMSizeTypesStandardD14 ...
	VMSizeTypesStandardD14 VMSizeTypes = "Standard_D14"
	// VMSizeTypesStandardD14V2 ...
	VMSizeTypesStandardD14V2 VMSizeTypes = "Standard_D14_v2"
	// VMSizeTypesStandardD14V2Promo ...
	VMSizeTypesStandardD14V2Promo VMSizeTypes = "Standard_D14_v2_Promo"
	// VMSizeTypesStandardD15V2 ...
	VMSizeTypesStandardD15V2 VMSizeTypes = "Standard_D15_v2"
	// VMSizeTypesStandardD16sV3 ...
	VMSizeTypesStandardD16sV3 VMSizeTypes = "Standard_D16s_v3"
	// VMSizeTypesStandardD16V3 ...
	VMSizeTypesStandardD16V3 VMSizeTypes = "Standard_D16_v3"
	// VMSizeTypesStandardD1V2 ...
	VMSizeTypesStandardD1V2 VMSizeTypes = "Standard_D1_v2"
	// VMSizeTypesStandardD2 ...
	VMSizeTypesStandardD2 VMSizeTypes = "Standard_D2"
	// VMSizeTypesStandardD2sV3 ...
	VMSizeTypesStandardD2sV3 VMSizeTypes = "Standard_D2s_v3"
	// VMSizeTypesStandardD2V2 ...
	VMSizeTypesStandardD2V2 VMSizeTypes = "Standard_D2_v2"
	// VMSizeTypesStandardD2V2Promo ...
	VMSizeTypesStandardD2V2Promo VMSizeTypes = "Standard_D2_v2_Promo"
	// VMSizeTypesStandardD2V3 ...
	VMSizeTypesStandardD2V3 VMSizeTypes = "Standard_D2_v3"
	// VMSizeTypesStandardD3 ...
	VMSizeTypesStandardD3 VMSizeTypes = "Standard_D3"
	// VMSizeTypesStandardD32sV3 ...
	VMSizeTypesStandardD32sV3 VMSizeTypes = "Standard_D32s_v3"
	// VMSizeTypesStandardD32V3 ...
	VMSizeTypesStandardD32V3 VMSizeTypes = "Standard_D32_v3"
	// VMSizeTypesStandardD3V2 ...
	VMSizeTypesStandardD3V2 VMSizeTypes = "Standard_D3_v2"
	// VMSizeTypesStandardD3V2Promo ...
	VMSizeTypesStandardD3V2Promo VMSizeTypes = "Standard_D3_v2_Promo"
	// VMSizeTypesStandardD4 ...
	VMSizeTypesStandardD4 VMSizeTypes = "Standard_D4"
	// VMSizeTypesStandardD4sV3 ...
	VMSizeTypesStandardD4sV3 VMSizeTypes = "Standard_D4s_v3"
	// VMSizeTypesStandardD4V2 ...
	VMSizeTypesStandardD4V2 VMSizeTypes = "Standard_D4_v2"
	// VMSizeTypesStandardD4V2Promo ...
	VMSizeTypesStandardD4V2Promo VMSizeTypes = "Standard_D4_v2_Promo"
	// VMSizeTypesStandardD4V3 ...
	VMSizeTypesStandardD4V3 VMSizeTypes = "Standard_D4_v3"
	// VMSizeTypesStandardD5V2 ...
	VMSizeTypesStandardD5V2 VMSizeTypes = "Standard_D5_v2"
	// VMSizeTypesStandardD5V2Promo ...
	VMSizeTypesStandardD5V2Promo VMSizeTypes = "Standard_D5_v2_Promo"
	// VMSizeTypesStandardD64sV3 ...
	VMSizeTypesStandardD64sV3 VMSizeTypes = "Standard_D64s_v3"
	// VMSizeTypesStandardD64V3 ...
	VMSizeTypesStandardD64V3 VMSizeTypes = "Standard_D64_v3"
	// VMSizeTypesStandardD8sV3 ...
	VMSizeTypesStandardD8sV3 VMSizeTypes = "Standard_D8s_v3"
	// VMSizeTypesStandardD8V3 ...
	VMSizeTypesStandardD8V3 VMSizeTypes = "Standard_D8_v3"
	// VMSizeTypesStandardDS1 ...
	VMSizeTypesStandardDS1 VMSizeTypes = "Standard_DS1"
	// VMSizeTypesStandardDS11 ...
	VMSizeTypesStandardDS11 VMSizeTypes = "Standard_DS11"
	// VMSizeTypesStandardDS11V2 ...
	VMSizeTypesStandardDS11V2 VMSizeTypes = "Standard_DS11_v2"
	// VMSizeTypesStandardDS11V2Promo ...
	VMSizeTypesStandardDS11V2Promo VMSizeTypes = "Standard_DS11_v2_Promo"
	// VMSizeTypesStandardDS12 ...
	VMSizeTypesStandardDS12 VMSizeTypes = "Standard_DS12"
	// VMSizeTypesStandardDS12V2 ...
	VMSizeTypesStandardDS12V2 VMSizeTypes = "Standard_DS12_v2"
	// VMSizeTypesStandardDS12V2Promo ...
	VMSizeTypesStandardDS12V2Promo VMSizeTypes = "Standard_DS12_v2_Promo"
	// VMSizeTypesStandardDS13 ...
	VMSizeTypesStandardDS13 VMSizeTypes = "Standard_DS13"
	// VMSizeTypesStandardDS132V2 ...
	VMSizeTypesStandardDS132V2 VMSizeTypes = "Standard_DS13-2_v2"
	// VMSizeTypesStandardDS134V2 ...
	VMSizeTypesStandardDS134V2 VMSizeTypes = "Standard_DS13-4_v2"
	// VMSizeTypesStandardDS13V2 ...
	VMSizeTypesStandardDS13V2 VMSizeTypes = "Standard_DS13_v2"
	// VMSizeTypesStandardDS13V2Promo ...
	VMSizeTypesStandardDS13V2Promo VMSizeTypes = "Standard_DS13_v2_Promo"
	// VMSizeTypesStandardDS14 ...
	VMSizeTypesStandardDS14 VMSizeTypes = "Standard_DS14"
	// VMSizeTypesStandardDS144V2 ...
	VMSizeTypesStandardDS144V2 VMSizeTypes = "Standard_DS14-4_v2"
	// VMSizeTypesStandardDS148V2 ...
	VMSizeTypesStandardDS148V2 VMSizeTypes = "Standard_DS14-8_v2"
	// VMSizeTypesStandardDS14V2 ...
	VMSizeTypesStandardDS14V2 VMSizeTypes = "Standard_DS14_v2"
	// VMSizeTypesStandardDS14V2Promo ...
	VMSizeTypesStandardDS14V2Promo VMSizeTypes = "Standard_DS14_v2_Promo"
	// VMSizeTypesStandardDS15V2 ...
	VMSizeTypesStandardDS15V2 VMSizeTypes = "Standard_DS15_v2"
	// VMSizeTypesStandardDS1V2 ...
	VMSizeTypesStandardDS1V2 VMSizeTypes = "Standard_DS1_v2"
	// VMSizeTypesStandardDS2 ...
	VMSizeTypesStandardDS2 VMSizeTypes = "Standard_DS2"
	// VMSizeTypesStandardDS2V2 ...
	VMSizeTypesStandardDS2V2 VMSizeTypes = "Standard_DS2_v2"
	// VMSizeTypesStandardDS2V2Promo ...
	VMSizeTypesStandardDS2V2Promo VMSizeTypes = "Standard_DS2_v2_Promo"
	// VMSizeTypesStandardDS3 ...
	VMSizeTypesStandardDS3 VMSizeTypes = "Standard_DS3"
	// VMSizeTypesStandardDS3V2 ...
	VMSizeTypesStandardDS3V2 VMSizeTypes = "Standard_DS3_v2"
	// VMSizeTypesStandardDS3V2Promo ...
	VMSizeTypesStandardDS3V2Promo VMSizeTypes = "Standard_DS3_v2_Promo"
	// VMSizeTypesStandardDS4 ...
	VMSizeTypesStandardDS4 VMSizeTypes = "Standard_DS4"
	// VMSizeTypesStandardDS4V2 ...
	VMSizeTypesStandardDS4V2 VMSizeTypes = "Standard_DS4_v2"
	// VMSizeTypesStandardDS4V2Promo ...
	VMSizeTypesStandardDS4V2Promo VMSizeTypes = "Standard_DS4_v2_Promo"
	// VMSizeTypesStandardDS5V2 ...
	VMSizeTypesStandardDS5V2 VMSizeTypes = "Standard_DS5_v2"
	// VMSizeTypesStandardDS5V2Promo ...
	VMSizeTypesStandardDS5V2Promo VMSizeTypes = "Standard_DS5_v2_Promo"
	// VMSizeTypesStandardE16sV3 ...
	VMSizeTypesStandardE16sV3 VMSizeTypes = "Standard_E16s_v3"
	// VMSizeTypesStandardE16V3 ...
	VMSizeTypesStandardE16V3 VMSizeTypes = "Standard_E16_v3"
	// VMSizeTypesStandardE2sV3 ...
	VMSizeTypesStandardE2sV3 VMSizeTypes = "Standard_E2s_v3"
	// VMSizeTypesStandardE2V3 ...
	VMSizeTypesStandardE2V3 VMSizeTypes = "Standard_E2_v3"
	// VMSizeTypesStandardE3216sV3 ...
	VMSizeTypesStandardE3216sV3 VMSizeTypes = "Standard_E32-16s_v3"
	// VMSizeTypesStandardE328sV3 ...
	VMSizeTypesStandardE328sV3 VMSizeTypes = "Standard_E32-8s_v3"
	// VMSizeTypesStandardE32sV3 ...
	VMSizeTypesStandardE32sV3 VMSizeTypes = "Standard_E32s_v3"
	// VMSizeTypesStandardE32V3 ...
	VMSizeTypesStandardE32V3 VMSizeTypes = "Standard_E32_v3"
	// VMSizeTypesStandardE4sV3 ...
	VMSizeTypesStandardE4sV3 VMSizeTypes = "Standard_E4s_v3"
	// VMSizeTypesStandardE4V3 ...
	VMSizeTypesStandardE4V3 VMSizeTypes = "Standard_E4_v3"
	// VMSizeTypesStandardE6416sV3 ...
	VMSizeTypesStandardE6416sV3 VMSizeTypes = "Standard_E64-16s_v3"
	// VMSizeTypesStandardE6432sV3 ...
	VMSizeTypesStandardE6432sV3 VMSizeTypes = "Standard_E64-32s_v3"
	// VMSizeTypesStandardE64sV3 ...
	VMSizeTypesStandardE64sV3 VMSizeTypes = "Standard_E64s_v3"
	// VMSizeTypesStandardE64V3 ...
	VMSizeTypesStandardE64V3 VMSizeTypes = "Standard_E64_v3"
	// VMSizeTypesStandardE8sV3 ...
	VMSizeTypesStandardE8sV3 VMSizeTypes = "Standard_E8s_v3"
	// VMSizeTypesStandardE8V3 ...
	VMSizeTypesStandardE8V3 VMSizeTypes = "Standard_E8_v3"
	// VMSizeTypesStandardF1 ...
	VMSizeTypesStandardF1 VMSizeTypes = "Standard_F1"
	// VMSizeTypesStandardF16 ...
	VMSizeTypesStandardF16 VMSizeTypes = "Standard_F16"
	// VMSizeTypesStandardF16s ...
	VMSizeTypesStandardF16s VMSizeTypes = "Standard_F16s"
	// VMSizeTypesStandardF16sV2 ...
	VMSizeTypesStandardF16sV2 VMSizeTypes = "Standard_F16s_v2"
	// VMSizeTypesStandardF1s ...
	VMSizeTypesStandardF1s VMSizeTypes = "Standard_F1s"
	// VMSizeTypesStandardF2 ...
	VMSizeTypesStandardF2 VMSizeTypes = "Standard_F2"
	// VMSizeTypesStandardF2s ...
	VMSizeTypesStandardF2s VMSizeTypes = "Standard_F2s"
	// VMSizeTypesStandardF2sV2 ...
	VMSizeTypesStandardF2sV2 VMSizeTypes = "Standard_F2s_v2"
	// VMSizeTypesStandardF32sV2 ...
	VMSizeTypesStandardF32sV2 VMSizeTypes = "Standard_F32s_v2"
	// VMSizeTypesStandardF4 ...
	VMSizeTypesStandardF4 VMSizeTypes = "Standard_F4"
	// VMSizeTypesStandardF4s ...
	VMSizeTypesStandardF4s VMSizeTypes = "Standard_F4s"
	// VMSizeTypesStandardF4sV2 ...
	VMSizeTypesStandardF4sV2 VMSizeTypes = "Standard_F4s_v2"
	// VMSizeTypesStandardF64sV2 ...
	VMSizeTypesStandardF64sV2 VMSizeTypes = "Standard_F64s_v2"
	// VMSizeTypesStandardF72sV2 ...
	VMSizeTypesStandardF72sV2 VMSizeTypes = "Standard_F72s_v2"
	// VMSizeTypesStandardF8 ...
	VMSizeTypesStandardF8 VMSizeTypes = "Standard_F8"
	// VMSizeTypesStandardF8s ...
	VMSizeTypesStandardF8s VMSizeTypes = "Standard_F8s"
	// VMSizeTypesStandardF8sV2 ...
	VMSizeTypesStandardF8sV2 VMSizeTypes = "Standard_F8s_v2"
	// VMSizeTypesStandardG1 ...
	VMSizeTypesStandardG1 VMSizeTypes = "Standard_G1"
	// VMSizeTypesStandardG2 ...
	VMSizeTypesStandardG2 VMSizeTypes = "Standard_G2"
	// VMSizeTypesStandardG3 ...
	VMSizeTypesStandardG3 VMSizeTypes = "Standard_G3"
	// VMSizeTypesStandardG4 ...
	VMSizeTypesStandardG4 VMSizeTypes = "Standard_G4"
	// VMSizeTypesStandardG5 ...
	VMSizeTypesStandardG5 VMSizeTypes = "Standard_G5"
	// VMSizeTypesStandardGS1 ...
	VMSizeTypesStandardGS1 VMSizeTypes = "Standard_GS1"
	// VMSizeTypesStandardGS2 ...
	VMSizeTypesStandardGS2 VMSizeTypes = "Standard_GS2"
	// VMSizeTypesStandardGS3 ...
	VMSizeTypesStandardGS3 VMSizeTypes = "Standard_GS3"
	// VMSizeTypesStandardGS4 ...
	VMSizeTypesStandardGS4 VMSizeTypes = "Standard_GS4"
	// VMSizeTypesStandardGS44 ...
	VMSizeTypesStandardGS44 VMSizeTypes = "Standard_GS4-4"
	// VMSizeTypesStandardGS48 ...
	VMSizeTypesStandardGS48 VMSizeTypes = "Standard_GS4-8"
	// VMSizeTypesStandardGS5 ...
	VMSizeTypesStandardGS5 VMSizeTypes = "Standard_GS5"
	// VMSizeTypesStandardGS516 ...
	VMSizeTypesStandardGS516 VMSizeTypes = "Standard_GS5-16"
	// VMSizeTypesStandardGS58 ...
	VMSizeTypesStandardGS58 VMSizeTypes = "Standard_GS5-8"
	// VMSizeTypesStandardH16 ...
	VMSizeTypesStandardH16 VMSizeTypes = "Standard_H16"
	// VMSizeTypesStandardH16m ...
	VMSizeTypesStandardH16m VMSizeTypes = "Standard_H16m"
	// VMSizeTypesStandardH16mr ...
	VMSizeTypesStandardH16mr VMSizeTypes = "Standard_H16mr"
	// VMSizeTypesStandardH16r ...
	VMSizeTypesStandardH16r VMSizeTypes = "Standard_H16r"
	// VMSizeTypesStandardH8 ...
	VMSizeTypesStandardH8 VMSizeTypes = "Standard_H8"
	// VMSizeTypesStandardH8m ...
	VMSizeTypesStandardH8m VMSizeTypes = "Standard_H8m"
	// VMSizeTypesStandardL16s ...
	VMSizeTypesStandardL16s VMSizeTypes = "Standard_L16s"
	// VMSizeTypesStandardL32s ...
	VMSizeTypesStandardL32s VMSizeTypes = "Standard_L32s"
	// VMSizeTypesStandardL4s ...
	VMSizeTypesStandardL4s VMSizeTypes = "Standard_L4s"
	// VMSizeTypesStandardL8s ...
	VMSizeTypesStandardL8s VMSizeTypes = "Standard_L8s"
	// VMSizeTypesStandardM12832ms ...
	VMSizeTypesStandardM12832ms VMSizeTypes = "Standard_M128-32ms"
	// VMSizeTypesStandardM12864ms ...
	VMSizeTypesStandardM12864ms VMSizeTypes = "Standard_M128-64ms"
	// VMSizeTypesStandardM128ms ...
	VMSizeTypesStandardM128ms VMSizeTypes = "Standard_M128ms"
	// VMSizeTypesStandardM128s ...
	VMSizeTypesStandardM128s VMSizeTypes = "Standard_M128s"
	// VMSizeTypesStandardM6416ms ...
	VMSizeTypesStandardM6416ms VMSizeTypes = "Standard_M64-16ms"
	// VMSizeTypesStandardM6432ms ...
	VMSizeTypesStandardM6432ms VMSizeTypes = "Standard_M64-32ms"
	// VMSizeTypesStandardM64ms ...
	VMSizeTypesStandardM64ms VMSizeTypes = "Standard_M64ms"
	// VMSizeTypesStandardM64s ...
	VMSizeTypesStandardM64s VMSizeTypes = "Standard_M64s"
	// VMSizeTypesStandardNC12 ...
	VMSizeTypesStandardNC12 VMSizeTypes = "Standard_NC12"
	// VMSizeTypesStandardNC12sV2 ...
	VMSizeTypesStandardNC12sV2 VMSizeTypes = "Standard_NC12s_v2"
	// VMSizeTypesStandardNC12sV3 ...
	VMSizeTypesStandardNC12sV3 VMSizeTypes = "Standard_NC12s_v3"
	// VMSizeTypesStandardNC24 ...
	VMSizeTypesStandardNC24 VMSizeTypes = "Standard_NC24"
	// VMSizeTypesStandardNC24r ...
	VMSizeTypesStandardNC24r VMSizeTypes = "Standard_NC24r"
	// VMSizeTypesStandardNC24rsV2 ...
	VMSizeTypesStandardNC24rsV2 VMSizeTypes = "Standard_NC24rs_v2"
	// VMSizeTypesStandardNC24rsV3 ...
	VMSizeTypesStandardNC24rsV3 VMSizeTypes = "Standard_NC24rs_v3"
	// VMSizeTypesStandardNC24sV2 ...
	VMSizeTypesStandardNC24sV2 VMSizeTypes = "Standard_NC24s_v2"
	// VMSizeTypesStandardNC24sV3 ...
	VMSizeTypesStandardNC24sV3 VMSizeTypes = "Standard_NC24s_v3"
	// VMSizeTypesStandardNC6 ...
	VMSizeTypesStandardNC6 VMSizeTypes = "Standard_NC6"
	// VMSizeTypesStandardNC6sV2 ...
	VMSizeTypesStandardNC6sV2 VMSizeTypes = "Standard_NC6s_v2"
	// VMSizeTypesStandardNC6sV3 ...
	VMSizeTypesStandardNC6sV3 VMSizeTypes = "Standard_NC6s_v3"
	// VMSizeTypesStandardND12s ...
	VMSizeTypesStandardND12s VMSizeTypes = "Standard_ND12s"
	// VMSizeTypesStandardND24rs ...
	VMSizeTypesStandardND24rs VMSizeTypes = "Standard_ND24rs"
	// VMSizeTypesStandardND24s ...
	VMSizeTypesStandardND24s VMSizeTypes = "Standard_ND24s"
	// VMSizeTypesStandardND6s ...
	VMSizeTypesStandardND6s VMSizeTypes = "Standard_ND6s"
	// VMSizeTypesStandardNV12 ...
	VMSizeTypesStandardNV12 VMSizeTypes = "Standard_NV12"
	// VMSizeTypesStandardNV24 ...
	VMSizeTypesStandardNV24 VMSizeTypes = "Standard_NV24"
	// VMSizeTypesStandardNV6 ...
	VMSizeTypesStandardNV6 VMSizeTypes = "Standard_NV6"
)

// PossibleVMSizeTypesValues returns an array of possible values for the VMSizeTypes const type.
func PossibleVMSizeTypesValues() []VMSizeTypes {
	return []VMSizeTypes{VMSizeTypesStandardA1, VMSizeTypesStandardA10, VMSizeTypesStandardA11, VMSizeTypesStandardA1V2, VMSizeTypesStandardA2, VMSizeTypesStandardA2mV2, VMSizeTypesStandardA2V2, VMSizeTypesStandardA3, VMSizeTypesStandardA4, VMSizeTypesStandardA4mV2, VMSizeTypesStandardA4V2, VMSizeTypesStandardA5, VMSizeTypesStandardA6, VMSizeTypesStandardA7, VMSizeTypesStandardA8, VMSizeTypesStandardA8mV2, VMSizeTypesStandardA8V2, VMSizeTypesStandardA9, VMSizeTypesStandardB2ms, VMSizeTypesStandardB2s, VMSizeTypesStandardB4ms, VMSizeTypesStandardB8ms, VMSizeTypesStandardD1, VMSizeTypesStandardD11, VMSizeTypesStandardD11V2, VMSizeTypesStandardD11V2Promo, VMSizeTypesStandardD12, VMSizeTypesStandardD12V2, VMSizeTypesStandardD12V2Promo, VMSizeTypesStandardD13, VMSizeTypesStandardD13V2, VMSizeTypesStandardD13V2Promo, VMSizeTypesStandardD14, VMSizeTypesStandardD14V2, VMSizeTypesStandardD14V2Promo, VMSizeTypesStandardD15V2, VMSizeTypesStandardD16sV3, VMSizeTypesStandardD16V3, VMSizeTypesStandardD1V2, VMSizeTypesStandardD2, VMSizeTypesStandardD2sV3, VMSizeTypesStandardD2V2, VMSizeTypesStandardD2V2Promo, VMSizeTypesStandardD2V3, VMSizeTypesStandardD3, VMSizeTypesStandardD32sV3, VMSizeTypesStandardD32V3, VMSizeTypesStandardD3V2, VMSizeTypesStandardD3V2Promo, VMSizeTypesStandardD4, VMSizeTypesStandardD4sV3, VMSizeTypesStandardD4V2, VMSizeTypesStandardD4V2Promo, VMSizeTypesStandardD4V3, VMSizeTypesStandardD5V2, VMSizeTypesStandardD5V2Promo, VMSizeTypesStandardD64sV3, VMSizeTypesStandardD64V3, VMSizeTypesStandardD8sV3, VMSizeTypesStandardD8V3, VMSizeTypesStandardDS1, VMSizeTypesStandardDS11, VMSizeTypesStandardDS11V2, VMSizeTypesStandardDS11V2Promo, VMSizeTypesStandardDS12, VMSizeTypesStandardDS12V2, VMSizeTypesStandardDS12V2Promo, VMSizeTypesStandardDS13, VMSizeTypesStandardDS132V2, VMSizeTypesStandardDS134V2, VMSizeTypesStandardDS13V2, VMSizeTypesStandardDS13V2Promo, VMSizeTypesStandardDS14, VMSizeTypesStandardDS144V2, VMSizeTypesStandardDS148V2, VMSizeTypesStandardDS14V2, VMSizeTypesStandardDS14V2Promo, VMSizeTypesStandardDS15V2, VMSizeTypesStandardDS1V2, VMSizeTypesStandardDS2, VMSizeTypesStandardDS2V2, VMSizeTypesStandardDS2V2Promo, VMSizeTypesStandardDS3, VMSizeTypesStandardDS3V2, VMSizeTypesStandardDS3V2Promo, VMSizeTypesStandardDS4, VMSizeTypesStandardDS4V2, VMSizeTypesStandardDS4V2Promo, VMSizeTypesStandardDS5V2, VMSizeTypesStandardDS5V2Promo, VMSizeTypesStandardE16sV3, VMSizeTypesStandardE16V3, VMSizeTypesStandardE2sV3, VMSizeTypesStandardE2V3, VMSizeTypesStandardE3216sV3, VMSizeTypesStandardE328sV3, VMSizeTypesStandardE32sV3, VMSizeTypesStandardE32V3, VMSizeTypesStandardE4sV3, VMSizeTypesStandardE4V3, VMSizeTypesStandardE6416sV3, VMSizeTypesStandardE6432sV3, VMSizeTypesStandardE64sV3, VMSizeTypesStandardE64V3, VMSizeTypesStandardE8sV3, VMSizeTypesStandardE8V3, VMSizeTypesStandardF1, VMSizeTypesStandardF16, VMSizeTypesStandardF16s, VMSizeTypesStandardF16sV2, VMSizeTypesStandardF1s, VMSizeTypesStandardF2, VMSizeTypesStandardF2s, VMSizeTypesStandardF2sV2, VMSizeTypesStandardF32sV2, VMSizeTypesStandardF4, VMSizeTypesStandardF4s, VMSizeTypesStandardF4sV2, VMSizeTypesStandardF64sV2, VMSizeTypesStandardF72sV2, VMSizeTypesStandardF8, VMSizeTypesStandardF8s, VMSizeTypesStandardF8sV2, VMSizeTypesStandardG1, VMSizeTypesStandardG2, VMSizeTypesStandardG3, VMSizeTypesStandardG4, VMSizeTypesStandardG5, VMSizeTypesStandardGS1, VMSizeTypesStandardGS2, VMSizeTypesStandardGS3, VMSizeTypesStandardGS4, VMSizeTypesStandardGS44, VMSizeTypesStandardGS48, VMSizeTypesStandardGS5, VMSizeTypesStandardGS516, VMSizeTypesStandardGS58, VMSizeTypesStandardH16, VMSizeTypesStandardH16m, VMSizeTypesStandardH16mr, VMSizeTypesStandardH16r, VMSizeTypesStandardH8, VMSizeTypesStandardH8m, VMSizeTypesStandardL16s, VMSizeTypesStandardL32s, VMSizeTypesStandardL4s, VMSizeTypesStandardL8s, VMSizeTypesStandardM12832ms, VMSizeTypesStandardM12864ms, VMSizeTypesStandardM128ms, VMSizeTypesStandardM128s, VMSizeTypesStandardM6416ms, VMSizeTypesStandardM6432ms, VMSizeTypesStandardM64ms, VMSizeTypesStandardM64s, VMSizeTypesStandardNC12, VMSizeTypesStandardNC12sV2, VMSizeTypesStandardNC12sV3, VMSizeTypesStandardNC24, VMSizeTypesStandardNC24r, VMSizeTypesStandardNC24rsV2, VMSizeTypesStandardNC24rsV3, VMSizeTypesStandardNC24sV2, VMSizeTypesStandardNC24sV3, VMSizeTypesStandardNC6, VMSizeTypesStandardNC6sV2, VMSizeTypesStandardNC6sV3, VMSizeTypesStandardND12s, VMSizeTypesStandardND24rs, VMSizeTypesStandardND24s, VMSizeTypesStandardND6s, VMSizeTypesStandardNV12, VMSizeTypesStandardNV24, VMSizeTypesStandardNV6}
}
