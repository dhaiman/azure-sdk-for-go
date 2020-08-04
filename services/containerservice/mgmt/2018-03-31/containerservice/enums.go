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

// NetworkPlugin enumerates the values for network plugin.
type NetworkPlugin string

const (
	// Azure ...
	Azure NetworkPlugin = "azure"
	// Kubenet ...
	Kubenet NetworkPlugin = "kubenet"
)

// PossibleNetworkPluginValues returns an array of possible values for the NetworkPlugin const type.
func PossibleNetworkPluginValues() []NetworkPlugin {
	return []NetworkPlugin{Azure, Kubenet}
}

// NetworkPolicy enumerates the values for network policy.
type NetworkPolicy string

const (
	// Calico ...
	Calico NetworkPolicy = "calico"
)

// PossibleNetworkPolicyValues returns an array of possible values for the NetworkPolicy const type.
func PossibleNetworkPolicyValues() []NetworkPolicy {
	return []NetworkPolicy{Calico}
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
	// StandardA1 ...
	StandardA1 VMSizeTypes = "Standard_A1"
	// StandardA10 ...
	StandardA10 VMSizeTypes = "Standard_A10"
	// StandardA11 ...
	StandardA11 VMSizeTypes = "Standard_A11"
	// StandardA1V2 ...
	StandardA1V2 VMSizeTypes = "Standard_A1_v2"
	// StandardA2 ...
	StandardA2 VMSizeTypes = "Standard_A2"
	// StandardA2mV2 ...
	StandardA2mV2 VMSizeTypes = "Standard_A2m_v2"
	// StandardA2V2 ...
	StandardA2V2 VMSizeTypes = "Standard_A2_v2"
	// StandardA3 ...
	StandardA3 VMSizeTypes = "Standard_A3"
	// StandardA4 ...
	StandardA4 VMSizeTypes = "Standard_A4"
	// StandardA4mV2 ...
	StandardA4mV2 VMSizeTypes = "Standard_A4m_v2"
	// StandardA4V2 ...
	StandardA4V2 VMSizeTypes = "Standard_A4_v2"
	// StandardA5 ...
	StandardA5 VMSizeTypes = "Standard_A5"
	// StandardA6 ...
	StandardA6 VMSizeTypes = "Standard_A6"
	// StandardA7 ...
	StandardA7 VMSizeTypes = "Standard_A7"
	// StandardA8 ...
	StandardA8 VMSizeTypes = "Standard_A8"
	// StandardA8mV2 ...
	StandardA8mV2 VMSizeTypes = "Standard_A8m_v2"
	// StandardA8V2 ...
	StandardA8V2 VMSizeTypes = "Standard_A8_v2"
	// StandardA9 ...
	StandardA9 VMSizeTypes = "Standard_A9"
	// StandardB2ms ...
	StandardB2ms VMSizeTypes = "Standard_B2ms"
	// StandardB2s ...
	StandardB2s VMSizeTypes = "Standard_B2s"
	// StandardB4ms ...
	StandardB4ms VMSizeTypes = "Standard_B4ms"
	// StandardB8ms ...
	StandardB8ms VMSizeTypes = "Standard_B8ms"
	// StandardD1 ...
	StandardD1 VMSizeTypes = "Standard_D1"
	// StandardD11 ...
	StandardD11 VMSizeTypes = "Standard_D11"
	// StandardD11V2 ...
	StandardD11V2 VMSizeTypes = "Standard_D11_v2"
	// StandardD11V2Promo ...
	StandardD11V2Promo VMSizeTypes = "Standard_D11_v2_Promo"
	// StandardD12 ...
	StandardD12 VMSizeTypes = "Standard_D12"
	// StandardD12V2 ...
	StandardD12V2 VMSizeTypes = "Standard_D12_v2"
	// StandardD12V2Promo ...
	StandardD12V2Promo VMSizeTypes = "Standard_D12_v2_Promo"
	// StandardD13 ...
	StandardD13 VMSizeTypes = "Standard_D13"
	// StandardD13V2 ...
	StandardD13V2 VMSizeTypes = "Standard_D13_v2"
	// StandardD13V2Promo ...
	StandardD13V2Promo VMSizeTypes = "Standard_D13_v2_Promo"
	// StandardD14 ...
	StandardD14 VMSizeTypes = "Standard_D14"
	// StandardD14V2 ...
	StandardD14V2 VMSizeTypes = "Standard_D14_v2"
	// StandardD14V2Promo ...
	StandardD14V2Promo VMSizeTypes = "Standard_D14_v2_Promo"
	// StandardD15V2 ...
	StandardD15V2 VMSizeTypes = "Standard_D15_v2"
	// StandardD16sV3 ...
	StandardD16sV3 VMSizeTypes = "Standard_D16s_v3"
	// StandardD16V3 ...
	StandardD16V3 VMSizeTypes = "Standard_D16_v3"
	// StandardD1V2 ...
	StandardD1V2 VMSizeTypes = "Standard_D1_v2"
	// StandardD2 ...
	StandardD2 VMSizeTypes = "Standard_D2"
	// StandardD2sV3 ...
	StandardD2sV3 VMSizeTypes = "Standard_D2s_v3"
	// StandardD2V2 ...
	StandardD2V2 VMSizeTypes = "Standard_D2_v2"
	// StandardD2V2Promo ...
	StandardD2V2Promo VMSizeTypes = "Standard_D2_v2_Promo"
	// StandardD2V3 ...
	StandardD2V3 VMSizeTypes = "Standard_D2_v3"
	// StandardD3 ...
	StandardD3 VMSizeTypes = "Standard_D3"
	// StandardD32sV3 ...
	StandardD32sV3 VMSizeTypes = "Standard_D32s_v3"
	// StandardD32V3 ...
	StandardD32V3 VMSizeTypes = "Standard_D32_v3"
	// StandardD3V2 ...
	StandardD3V2 VMSizeTypes = "Standard_D3_v2"
	// StandardD3V2Promo ...
	StandardD3V2Promo VMSizeTypes = "Standard_D3_v2_Promo"
	// StandardD4 ...
	StandardD4 VMSizeTypes = "Standard_D4"
	// StandardD4sV3 ...
	StandardD4sV3 VMSizeTypes = "Standard_D4s_v3"
	// StandardD4V2 ...
	StandardD4V2 VMSizeTypes = "Standard_D4_v2"
	// StandardD4V2Promo ...
	StandardD4V2Promo VMSizeTypes = "Standard_D4_v2_Promo"
	// StandardD4V3 ...
	StandardD4V3 VMSizeTypes = "Standard_D4_v3"
	// StandardD5V2 ...
	StandardD5V2 VMSizeTypes = "Standard_D5_v2"
	// StandardD5V2Promo ...
	StandardD5V2Promo VMSizeTypes = "Standard_D5_v2_Promo"
	// StandardD64sV3 ...
	StandardD64sV3 VMSizeTypes = "Standard_D64s_v3"
	// StandardD64V3 ...
	StandardD64V3 VMSizeTypes = "Standard_D64_v3"
	// StandardD8sV3 ...
	StandardD8sV3 VMSizeTypes = "Standard_D8s_v3"
	// StandardD8V3 ...
	StandardD8V3 VMSizeTypes = "Standard_D8_v3"
	// StandardDS1 ...
	StandardDS1 VMSizeTypes = "Standard_DS1"
	// StandardDS11 ...
	StandardDS11 VMSizeTypes = "Standard_DS11"
	// StandardDS11V2 ...
	StandardDS11V2 VMSizeTypes = "Standard_DS11_v2"
	// StandardDS11V2Promo ...
	StandardDS11V2Promo VMSizeTypes = "Standard_DS11_v2_Promo"
	// StandardDS12 ...
	StandardDS12 VMSizeTypes = "Standard_DS12"
	// StandardDS12V2 ...
	StandardDS12V2 VMSizeTypes = "Standard_DS12_v2"
	// StandardDS12V2Promo ...
	StandardDS12V2Promo VMSizeTypes = "Standard_DS12_v2_Promo"
	// StandardDS13 ...
	StandardDS13 VMSizeTypes = "Standard_DS13"
	// StandardDS132V2 ...
	StandardDS132V2 VMSizeTypes = "Standard_DS13-2_v2"
	// StandardDS134V2 ...
	StandardDS134V2 VMSizeTypes = "Standard_DS13-4_v2"
	// StandardDS13V2 ...
	StandardDS13V2 VMSizeTypes = "Standard_DS13_v2"
	// StandardDS13V2Promo ...
	StandardDS13V2Promo VMSizeTypes = "Standard_DS13_v2_Promo"
	// StandardDS14 ...
	StandardDS14 VMSizeTypes = "Standard_DS14"
	// StandardDS144V2 ...
	StandardDS144V2 VMSizeTypes = "Standard_DS14-4_v2"
	// StandardDS148V2 ...
	StandardDS148V2 VMSizeTypes = "Standard_DS14-8_v2"
	// StandardDS14V2 ...
	StandardDS14V2 VMSizeTypes = "Standard_DS14_v2"
	// StandardDS14V2Promo ...
	StandardDS14V2Promo VMSizeTypes = "Standard_DS14_v2_Promo"
	// StandardDS15V2 ...
	StandardDS15V2 VMSizeTypes = "Standard_DS15_v2"
	// StandardDS1V2 ...
	StandardDS1V2 VMSizeTypes = "Standard_DS1_v2"
	// StandardDS2 ...
	StandardDS2 VMSizeTypes = "Standard_DS2"
	// StandardDS2V2 ...
	StandardDS2V2 VMSizeTypes = "Standard_DS2_v2"
	// StandardDS2V2Promo ...
	StandardDS2V2Promo VMSizeTypes = "Standard_DS2_v2_Promo"
	// StandardDS3 ...
	StandardDS3 VMSizeTypes = "Standard_DS3"
	// StandardDS3V2 ...
	StandardDS3V2 VMSizeTypes = "Standard_DS3_v2"
	// StandardDS3V2Promo ...
	StandardDS3V2Promo VMSizeTypes = "Standard_DS3_v2_Promo"
	// StandardDS4 ...
	StandardDS4 VMSizeTypes = "Standard_DS4"
	// StandardDS4V2 ...
	StandardDS4V2 VMSizeTypes = "Standard_DS4_v2"
	// StandardDS4V2Promo ...
	StandardDS4V2Promo VMSizeTypes = "Standard_DS4_v2_Promo"
	// StandardDS5V2 ...
	StandardDS5V2 VMSizeTypes = "Standard_DS5_v2"
	// StandardDS5V2Promo ...
	StandardDS5V2Promo VMSizeTypes = "Standard_DS5_v2_Promo"
	// StandardE16sV3 ...
	StandardE16sV3 VMSizeTypes = "Standard_E16s_v3"
	// StandardE16V3 ...
	StandardE16V3 VMSizeTypes = "Standard_E16_v3"
	// StandardE2sV3 ...
	StandardE2sV3 VMSizeTypes = "Standard_E2s_v3"
	// StandardE2V3 ...
	StandardE2V3 VMSizeTypes = "Standard_E2_v3"
	// StandardE3216sV3 ...
	StandardE3216sV3 VMSizeTypes = "Standard_E32-16s_v3"
	// StandardE328sV3 ...
	StandardE328sV3 VMSizeTypes = "Standard_E32-8s_v3"
	// StandardE32sV3 ...
	StandardE32sV3 VMSizeTypes = "Standard_E32s_v3"
	// StandardE32V3 ...
	StandardE32V3 VMSizeTypes = "Standard_E32_v3"
	// StandardE4sV3 ...
	StandardE4sV3 VMSizeTypes = "Standard_E4s_v3"
	// StandardE4V3 ...
	StandardE4V3 VMSizeTypes = "Standard_E4_v3"
	// StandardE6416sV3 ...
	StandardE6416sV3 VMSizeTypes = "Standard_E64-16s_v3"
	// StandardE6432sV3 ...
	StandardE6432sV3 VMSizeTypes = "Standard_E64-32s_v3"
	// StandardE64sV3 ...
	StandardE64sV3 VMSizeTypes = "Standard_E64s_v3"
	// StandardE64V3 ...
	StandardE64V3 VMSizeTypes = "Standard_E64_v3"
	// StandardE8sV3 ...
	StandardE8sV3 VMSizeTypes = "Standard_E8s_v3"
	// StandardE8V3 ...
	StandardE8V3 VMSizeTypes = "Standard_E8_v3"
	// StandardF1 ...
	StandardF1 VMSizeTypes = "Standard_F1"
	// StandardF16 ...
	StandardF16 VMSizeTypes = "Standard_F16"
	// StandardF16s ...
	StandardF16s VMSizeTypes = "Standard_F16s"
	// StandardF16sV2 ...
	StandardF16sV2 VMSizeTypes = "Standard_F16s_v2"
	// StandardF1s ...
	StandardF1s VMSizeTypes = "Standard_F1s"
	// StandardF2 ...
	StandardF2 VMSizeTypes = "Standard_F2"
	// StandardF2s ...
	StandardF2s VMSizeTypes = "Standard_F2s"
	// StandardF2sV2 ...
	StandardF2sV2 VMSizeTypes = "Standard_F2s_v2"
	// StandardF32sV2 ...
	StandardF32sV2 VMSizeTypes = "Standard_F32s_v2"
	// StandardF4 ...
	StandardF4 VMSizeTypes = "Standard_F4"
	// StandardF4s ...
	StandardF4s VMSizeTypes = "Standard_F4s"
	// StandardF4sV2 ...
	StandardF4sV2 VMSizeTypes = "Standard_F4s_v2"
	// StandardF64sV2 ...
	StandardF64sV2 VMSizeTypes = "Standard_F64s_v2"
	// StandardF72sV2 ...
	StandardF72sV2 VMSizeTypes = "Standard_F72s_v2"
	// StandardF8 ...
	StandardF8 VMSizeTypes = "Standard_F8"
	// StandardF8s ...
	StandardF8s VMSizeTypes = "Standard_F8s"
	// StandardF8sV2 ...
	StandardF8sV2 VMSizeTypes = "Standard_F8s_v2"
	// StandardG1 ...
	StandardG1 VMSizeTypes = "Standard_G1"
	// StandardG2 ...
	StandardG2 VMSizeTypes = "Standard_G2"
	// StandardG3 ...
	StandardG3 VMSizeTypes = "Standard_G3"
	// StandardG4 ...
	StandardG4 VMSizeTypes = "Standard_G4"
	// StandardG5 ...
	StandardG5 VMSizeTypes = "Standard_G5"
	// StandardGS1 ...
	StandardGS1 VMSizeTypes = "Standard_GS1"
	// StandardGS2 ...
	StandardGS2 VMSizeTypes = "Standard_GS2"
	// StandardGS3 ...
	StandardGS3 VMSizeTypes = "Standard_GS3"
	// StandardGS4 ...
	StandardGS4 VMSizeTypes = "Standard_GS4"
	// StandardGS44 ...
	StandardGS44 VMSizeTypes = "Standard_GS4-4"
	// StandardGS48 ...
	StandardGS48 VMSizeTypes = "Standard_GS4-8"
	// StandardGS5 ...
	StandardGS5 VMSizeTypes = "Standard_GS5"
	// StandardGS516 ...
	StandardGS516 VMSizeTypes = "Standard_GS5-16"
	// StandardGS58 ...
	StandardGS58 VMSizeTypes = "Standard_GS5-8"
	// StandardH16 ...
	StandardH16 VMSizeTypes = "Standard_H16"
	// StandardH16m ...
	StandardH16m VMSizeTypes = "Standard_H16m"
	// StandardH16mr ...
	StandardH16mr VMSizeTypes = "Standard_H16mr"
	// StandardH16r ...
	StandardH16r VMSizeTypes = "Standard_H16r"
	// StandardH8 ...
	StandardH8 VMSizeTypes = "Standard_H8"
	// StandardH8m ...
	StandardH8m VMSizeTypes = "Standard_H8m"
	// StandardL16s ...
	StandardL16s VMSizeTypes = "Standard_L16s"
	// StandardL32s ...
	StandardL32s VMSizeTypes = "Standard_L32s"
	// StandardL4s ...
	StandardL4s VMSizeTypes = "Standard_L4s"
	// StandardL8s ...
	StandardL8s VMSizeTypes = "Standard_L8s"
	// StandardM12832ms ...
	StandardM12832ms VMSizeTypes = "Standard_M128-32ms"
	// StandardM12864ms ...
	StandardM12864ms VMSizeTypes = "Standard_M128-64ms"
	// StandardM128ms ...
	StandardM128ms VMSizeTypes = "Standard_M128ms"
	// StandardM128s ...
	StandardM128s VMSizeTypes = "Standard_M128s"
	// StandardM6416ms ...
	StandardM6416ms VMSizeTypes = "Standard_M64-16ms"
	// StandardM6432ms ...
	StandardM6432ms VMSizeTypes = "Standard_M64-32ms"
	// StandardM64ms ...
	StandardM64ms VMSizeTypes = "Standard_M64ms"
	// StandardM64s ...
	StandardM64s VMSizeTypes = "Standard_M64s"
	// StandardNC12 ...
	StandardNC12 VMSizeTypes = "Standard_NC12"
	// StandardNC12sV2 ...
	StandardNC12sV2 VMSizeTypes = "Standard_NC12s_v2"
	// StandardNC12sV3 ...
	StandardNC12sV3 VMSizeTypes = "Standard_NC12s_v3"
	// StandardNC24 ...
	StandardNC24 VMSizeTypes = "Standard_NC24"
	// StandardNC24r ...
	StandardNC24r VMSizeTypes = "Standard_NC24r"
	// StandardNC24rsV2 ...
	StandardNC24rsV2 VMSizeTypes = "Standard_NC24rs_v2"
	// StandardNC24rsV3 ...
	StandardNC24rsV3 VMSizeTypes = "Standard_NC24rs_v3"
	// StandardNC24sV2 ...
	StandardNC24sV2 VMSizeTypes = "Standard_NC24s_v2"
	// StandardNC24sV3 ...
	StandardNC24sV3 VMSizeTypes = "Standard_NC24s_v3"
	// StandardNC6 ...
	StandardNC6 VMSizeTypes = "Standard_NC6"
	// StandardNC6sV2 ...
	StandardNC6sV2 VMSizeTypes = "Standard_NC6s_v2"
	// StandardNC6sV3 ...
	StandardNC6sV3 VMSizeTypes = "Standard_NC6s_v3"
	// StandardND12s ...
	StandardND12s VMSizeTypes = "Standard_ND12s"
	// StandardND24rs ...
	StandardND24rs VMSizeTypes = "Standard_ND24rs"
	// StandardND24s ...
	StandardND24s VMSizeTypes = "Standard_ND24s"
	// StandardND6s ...
	StandardND6s VMSizeTypes = "Standard_ND6s"
	// StandardNV12 ...
	StandardNV12 VMSizeTypes = "Standard_NV12"
	// StandardNV24 ...
	StandardNV24 VMSizeTypes = "Standard_NV24"
	// StandardNV6 ...
	StandardNV6 VMSizeTypes = "Standard_NV6"
)

// PossibleVMSizeTypesValues returns an array of possible values for the VMSizeTypes const type.
func PossibleVMSizeTypesValues() []VMSizeTypes {
	return []VMSizeTypes{StandardA1, StandardA10, StandardA11, StandardA1V2, StandardA2, StandardA2mV2, StandardA2V2, StandardA3, StandardA4, StandardA4mV2, StandardA4V2, StandardA5, StandardA6, StandardA7, StandardA8, StandardA8mV2, StandardA8V2, StandardA9, StandardB2ms, StandardB2s, StandardB4ms, StandardB8ms, StandardD1, StandardD11, StandardD11V2, StandardD11V2Promo, StandardD12, StandardD12V2, StandardD12V2Promo, StandardD13, StandardD13V2, StandardD13V2Promo, StandardD14, StandardD14V2, StandardD14V2Promo, StandardD15V2, StandardD16sV3, StandardD16V3, StandardD1V2, StandardD2, StandardD2sV3, StandardD2V2, StandardD2V2Promo, StandardD2V3, StandardD3, StandardD32sV3, StandardD32V3, StandardD3V2, StandardD3V2Promo, StandardD4, StandardD4sV3, StandardD4V2, StandardD4V2Promo, StandardD4V3, StandardD5V2, StandardD5V2Promo, StandardD64sV3, StandardD64V3, StandardD8sV3, StandardD8V3, StandardDS1, StandardDS11, StandardDS11V2, StandardDS11V2Promo, StandardDS12, StandardDS12V2, StandardDS12V2Promo, StandardDS13, StandardDS132V2, StandardDS134V2, StandardDS13V2, StandardDS13V2Promo, StandardDS14, StandardDS144V2, StandardDS148V2, StandardDS14V2, StandardDS14V2Promo, StandardDS15V2, StandardDS1V2, StandardDS2, StandardDS2V2, StandardDS2V2Promo, StandardDS3, StandardDS3V2, StandardDS3V2Promo, StandardDS4, StandardDS4V2, StandardDS4V2Promo, StandardDS5V2, StandardDS5V2Promo, StandardE16sV3, StandardE16V3, StandardE2sV3, StandardE2V3, StandardE3216sV3, StandardE328sV3, StandardE32sV3, StandardE32V3, StandardE4sV3, StandardE4V3, StandardE6416sV3, StandardE6432sV3, StandardE64sV3, StandardE64V3, StandardE8sV3, StandardE8V3, StandardF1, StandardF16, StandardF16s, StandardF16sV2, StandardF1s, StandardF2, StandardF2s, StandardF2sV2, StandardF32sV2, StandardF4, StandardF4s, StandardF4sV2, StandardF64sV2, StandardF72sV2, StandardF8, StandardF8s, StandardF8sV2, StandardG1, StandardG2, StandardG3, StandardG4, StandardG5, StandardGS1, StandardGS2, StandardGS3, StandardGS4, StandardGS44, StandardGS48, StandardGS5, StandardGS516, StandardGS58, StandardH16, StandardH16m, StandardH16mr, StandardH16r, StandardH8, StandardH8m, StandardL16s, StandardL32s, StandardL4s, StandardL8s, StandardM12832ms, StandardM12864ms, StandardM128ms, StandardM128s, StandardM6416ms, StandardM6432ms, StandardM64ms, StandardM64s, StandardNC12, StandardNC12sV2, StandardNC12sV3, StandardNC24, StandardNC24r, StandardNC24rsV2, StandardNC24rsV3, StandardNC24sV2, StandardNC24sV3, StandardNC6, StandardNC6sV2, StandardNC6sV3, StandardND12s, StandardND24rs, StandardND24s, StandardND6s, StandardNV12, StandardNV24, StandardNV6}
}
