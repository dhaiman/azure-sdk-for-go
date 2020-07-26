package network

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

// ApplicationGatewayCookieBasedAffinity enumerates the values for application gateway cookie based affinity.
type ApplicationGatewayCookieBasedAffinity string

const (
	// Disabled ...
	Disabled ApplicationGatewayCookieBasedAffinity = "Disabled"
	// Enabled ...
	Enabled ApplicationGatewayCookieBasedAffinity = "Enabled"
)

// PossibleApplicationGatewayCookieBasedAffinityValues returns an array of possible values for the ApplicationGatewayCookieBasedAffinity const type.
func PossibleApplicationGatewayCookieBasedAffinityValues() []ApplicationGatewayCookieBasedAffinity {
	return []ApplicationGatewayCookieBasedAffinity{Disabled, Enabled}
}

// ApplicationGatewayOperationalState enumerates the values for application gateway operational state.
type ApplicationGatewayOperationalState string

const (
	// Running ...
	Running ApplicationGatewayOperationalState = "Running"
	// Starting ...
	Starting ApplicationGatewayOperationalState = "Starting"
	// Stopped ...
	Stopped ApplicationGatewayOperationalState = "Stopped"
	// Stopping ...
	Stopping ApplicationGatewayOperationalState = "Stopping"
)

// PossibleApplicationGatewayOperationalStateValues returns an array of possible values for the ApplicationGatewayOperationalState const type.
func PossibleApplicationGatewayOperationalStateValues() []ApplicationGatewayOperationalState {
	return []ApplicationGatewayOperationalState{Running, Starting, Stopped, Stopping}
}

// ApplicationGatewayProtocol enumerates the values for application gateway protocol.
type ApplicationGatewayProtocol string

const (
	// HTTP ...
	HTTP ApplicationGatewayProtocol = "Http"
	// HTTPS ...
	HTTPS ApplicationGatewayProtocol = "Https"
)

// PossibleApplicationGatewayProtocolValues returns an array of possible values for the ApplicationGatewayProtocol const type.
func PossibleApplicationGatewayProtocolValues() []ApplicationGatewayProtocol {
	return []ApplicationGatewayProtocol{HTTP, HTTPS}
}

// ApplicationGatewayRequestRoutingRuleType enumerates the values for application gateway request routing rule
// type.
type ApplicationGatewayRequestRoutingRuleType string

const (
	// Basic ...
	Basic ApplicationGatewayRequestRoutingRuleType = "Basic"
)

// PossibleApplicationGatewayRequestRoutingRuleTypeValues returns an array of possible values for the ApplicationGatewayRequestRoutingRuleType const type.
func PossibleApplicationGatewayRequestRoutingRuleTypeValues() []ApplicationGatewayRequestRoutingRuleType {
	return []ApplicationGatewayRequestRoutingRuleType{Basic}
}

// ApplicationGatewaySkuName enumerates the values for application gateway sku name.
type ApplicationGatewaySkuName string

const (
	// StandardLarge ...
	StandardLarge ApplicationGatewaySkuName = "Standard_Large"
	// StandardMedium ...
	StandardMedium ApplicationGatewaySkuName = "Standard_Medium"
	// StandardSmall ...
	StandardSmall ApplicationGatewaySkuName = "Standard_Small"
)

// PossibleApplicationGatewaySkuNameValues returns an array of possible values for the ApplicationGatewaySkuName const type.
func PossibleApplicationGatewaySkuNameValues() []ApplicationGatewaySkuName {
	return []ApplicationGatewaySkuName{StandardLarge, StandardMedium, StandardSmall}
}

// ApplicationGatewayTier enumerates the values for application gateway tier.
type ApplicationGatewayTier string

const (
	// Standard ...
	Standard ApplicationGatewayTier = "Standard"
)

// PossibleApplicationGatewayTierValues returns an array of possible values for the ApplicationGatewayTier const type.
func PossibleApplicationGatewayTierValues() []ApplicationGatewayTier {
	return []ApplicationGatewayTier{Standard}
}

// AuthorizationUseStatus enumerates the values for authorization use status.
type AuthorizationUseStatus string

const (
	// Available ...
	Available AuthorizationUseStatus = "Available"
	// InUse ...
	InUse AuthorizationUseStatus = "InUse"
)

// PossibleAuthorizationUseStatusValues returns an array of possible values for the AuthorizationUseStatus const type.
func PossibleAuthorizationUseStatusValues() []AuthorizationUseStatus {
	return []AuthorizationUseStatus{Available, InUse}
}

// ExpressRouteCircuitPeeringAdvertisedPublicPrefixState enumerates the values for express route circuit
// peering advertised public prefix state.
type ExpressRouteCircuitPeeringAdvertisedPublicPrefixState string

const (
	// Configured ...
	Configured ExpressRouteCircuitPeeringAdvertisedPublicPrefixState = "Configured"
	// Configuring ...
	Configuring ExpressRouteCircuitPeeringAdvertisedPublicPrefixState = "Configuring"
	// NotConfigured ...
	NotConfigured ExpressRouteCircuitPeeringAdvertisedPublicPrefixState = "NotConfigured"
	// ValidationNeeded ...
	ValidationNeeded ExpressRouteCircuitPeeringAdvertisedPublicPrefixState = "ValidationNeeded"
)

// PossibleExpressRouteCircuitPeeringAdvertisedPublicPrefixStateValues returns an array of possible values for the ExpressRouteCircuitPeeringAdvertisedPublicPrefixState const type.
func PossibleExpressRouteCircuitPeeringAdvertisedPublicPrefixStateValues() []ExpressRouteCircuitPeeringAdvertisedPublicPrefixState {
	return []ExpressRouteCircuitPeeringAdvertisedPublicPrefixState{Configured, Configuring, NotConfigured, ValidationNeeded}
}

// ExpressRouteCircuitPeeringState enumerates the values for express route circuit peering state.
type ExpressRouteCircuitPeeringState string

const (
	// ExpressRouteCircuitPeeringStateDisabled ...
	ExpressRouteCircuitPeeringStateDisabled ExpressRouteCircuitPeeringState = "Disabled"
	// ExpressRouteCircuitPeeringStateEnabled ...
	ExpressRouteCircuitPeeringStateEnabled ExpressRouteCircuitPeeringState = "Enabled"
)

// PossibleExpressRouteCircuitPeeringStateValues returns an array of possible values for the ExpressRouteCircuitPeeringState const type.
func PossibleExpressRouteCircuitPeeringStateValues() []ExpressRouteCircuitPeeringState {
	return []ExpressRouteCircuitPeeringState{ExpressRouteCircuitPeeringStateDisabled, ExpressRouteCircuitPeeringStateEnabled}
}

// ExpressRouteCircuitPeeringType enumerates the values for express route circuit peering type.
type ExpressRouteCircuitPeeringType string

const (
	// AzurePrivatePeering ...
	AzurePrivatePeering ExpressRouteCircuitPeeringType = "AzurePrivatePeering"
	// AzurePublicPeering ...
	AzurePublicPeering ExpressRouteCircuitPeeringType = "AzurePublicPeering"
	// MicrosoftPeering ...
	MicrosoftPeering ExpressRouteCircuitPeeringType = "MicrosoftPeering"
)

// PossibleExpressRouteCircuitPeeringTypeValues returns an array of possible values for the ExpressRouteCircuitPeeringType const type.
func PossibleExpressRouteCircuitPeeringTypeValues() []ExpressRouteCircuitPeeringType {
	return []ExpressRouteCircuitPeeringType{AzurePrivatePeering, AzurePublicPeering, MicrosoftPeering}
}

// ExpressRouteCircuitSkuFamily enumerates the values for express route circuit sku family.
type ExpressRouteCircuitSkuFamily string

const (
	// MeteredData ...
	MeteredData ExpressRouteCircuitSkuFamily = "MeteredData"
	// UnlimitedData ...
	UnlimitedData ExpressRouteCircuitSkuFamily = "UnlimitedData"
)

// PossibleExpressRouteCircuitSkuFamilyValues returns an array of possible values for the ExpressRouteCircuitSkuFamily const type.
func PossibleExpressRouteCircuitSkuFamilyValues() []ExpressRouteCircuitSkuFamily {
	return []ExpressRouteCircuitSkuFamily{MeteredData, UnlimitedData}
}

// ExpressRouteCircuitSkuTier enumerates the values for express route circuit sku tier.
type ExpressRouteCircuitSkuTier string

const (
	// ExpressRouteCircuitSkuTierPremium ...
	ExpressRouteCircuitSkuTierPremium ExpressRouteCircuitSkuTier = "Premium"
	// ExpressRouteCircuitSkuTierStandard ...
	ExpressRouteCircuitSkuTierStandard ExpressRouteCircuitSkuTier = "Standard"
)

// PossibleExpressRouteCircuitSkuTierValues returns an array of possible values for the ExpressRouteCircuitSkuTier const type.
func PossibleExpressRouteCircuitSkuTierValues() []ExpressRouteCircuitSkuTier {
	return []ExpressRouteCircuitSkuTier{ExpressRouteCircuitSkuTierPremium, ExpressRouteCircuitSkuTierStandard}
}

// IPAllocationMethod enumerates the values for ip allocation method.
type IPAllocationMethod string

const (
	// Dynamic ...
	Dynamic IPAllocationMethod = "Dynamic"
	// Static ...
	Static IPAllocationMethod = "Static"
)

// PossibleIPAllocationMethodValues returns an array of possible values for the IPAllocationMethod const type.
func PossibleIPAllocationMethodValues() []IPAllocationMethod {
	return []IPAllocationMethod{Dynamic, Static}
}

// LoadDistribution enumerates the values for load distribution.
type LoadDistribution string

const (
	// Default ...
	Default LoadDistribution = "Default"
	// SourceIP ...
	SourceIP LoadDistribution = "SourceIP"
	// SourceIPProtocol ...
	SourceIPProtocol LoadDistribution = "SourceIPProtocol"
)

// PossibleLoadDistributionValues returns an array of possible values for the LoadDistribution const type.
func PossibleLoadDistributionValues() []LoadDistribution {
	return []LoadDistribution{Default, SourceIP, SourceIPProtocol}
}

// OperationStatus enumerates the values for operation status.
type OperationStatus string

const (
	// Failed ...
	Failed OperationStatus = "Failed"
	// InProgress ...
	InProgress OperationStatus = "InProgress"
	// Succeeded ...
	Succeeded OperationStatus = "Succeeded"
)

// PossibleOperationStatusValues returns an array of possible values for the OperationStatus const type.
func PossibleOperationStatusValues() []OperationStatus {
	return []OperationStatus{Failed, InProgress, Succeeded}
}

// ProbeProtocol enumerates the values for probe protocol.
type ProbeProtocol string

const (
	// ProbeProtocolHTTP ...
	ProbeProtocolHTTP ProbeProtocol = "Http"
	// ProbeProtocolTCP ...
	ProbeProtocolTCP ProbeProtocol = "Tcp"
)

// PossibleProbeProtocolValues returns an array of possible values for the ProbeProtocol const type.
func PossibleProbeProtocolValues() []ProbeProtocol {
	return []ProbeProtocol{ProbeProtocolHTTP, ProbeProtocolTCP}
}

// RouteNextHopType enumerates the values for route next hop type.
type RouteNextHopType string

const (
	// RouteNextHopTypeInternet ...
	RouteNextHopTypeInternet RouteNextHopType = "Internet"
	// RouteNextHopTypeNone ...
	RouteNextHopTypeNone RouteNextHopType = "None"
	// RouteNextHopTypeVirtualAppliance ...
	RouteNextHopTypeVirtualAppliance RouteNextHopType = "VirtualAppliance"
	// RouteNextHopTypeVirtualNetworkGateway ...
	RouteNextHopTypeVirtualNetworkGateway RouteNextHopType = "VirtualNetworkGateway"
	// RouteNextHopTypeVnetLocal ...
	RouteNextHopTypeVnetLocal RouteNextHopType = "VnetLocal"
)

// PossibleRouteNextHopTypeValues returns an array of possible values for the RouteNextHopType const type.
func PossibleRouteNextHopTypeValues() []RouteNextHopType {
	return []RouteNextHopType{RouteNextHopTypeInternet, RouteNextHopTypeNone, RouteNextHopTypeVirtualAppliance, RouteNextHopTypeVirtualNetworkGateway, RouteNextHopTypeVnetLocal}
}

// SecurityRuleAccess enumerates the values for security rule access.
type SecurityRuleAccess string

const (
	// Allow ...
	Allow SecurityRuleAccess = "Allow"
	// Deny ...
	Deny SecurityRuleAccess = "Deny"
)

// PossibleSecurityRuleAccessValues returns an array of possible values for the SecurityRuleAccess const type.
func PossibleSecurityRuleAccessValues() []SecurityRuleAccess {
	return []SecurityRuleAccess{Allow, Deny}
}

// SecurityRuleDirection enumerates the values for security rule direction.
type SecurityRuleDirection string

const (
	// Inbound ...
	Inbound SecurityRuleDirection = "Inbound"
	// Outbound ...
	Outbound SecurityRuleDirection = "Outbound"
)

// PossibleSecurityRuleDirectionValues returns an array of possible values for the SecurityRuleDirection const type.
func PossibleSecurityRuleDirectionValues() []SecurityRuleDirection {
	return []SecurityRuleDirection{Inbound, Outbound}
}

// SecurityRuleProtocol enumerates the values for security rule protocol.
type SecurityRuleProtocol string

const (
	// Asterisk ...
	Asterisk SecurityRuleProtocol = "*"
	// TCP ...
	TCP SecurityRuleProtocol = "Tcp"
	// UDP ...
	UDP SecurityRuleProtocol = "Udp"
)

// PossibleSecurityRuleProtocolValues returns an array of possible values for the SecurityRuleProtocol const type.
func PossibleSecurityRuleProtocolValues() []SecurityRuleProtocol {
	return []SecurityRuleProtocol{Asterisk, TCP, UDP}
}

// ServiceProviderProvisioningState enumerates the values for service provider provisioning state.
type ServiceProviderProvisioningState string

const (
	// Deprovisioning ...
	Deprovisioning ServiceProviderProvisioningState = "Deprovisioning"
	// NotProvisioned ...
	NotProvisioned ServiceProviderProvisioningState = "NotProvisioned"
	// Provisioned ...
	Provisioned ServiceProviderProvisioningState = "Provisioned"
	// Provisioning ...
	Provisioning ServiceProviderProvisioningState = "Provisioning"
)

// PossibleServiceProviderProvisioningStateValues returns an array of possible values for the ServiceProviderProvisioningState const type.
func PossibleServiceProviderProvisioningStateValues() []ServiceProviderProvisioningState {
	return []ServiceProviderProvisioningState{Deprovisioning, NotProvisioned, Provisioned, Provisioning}
}

// TransportProtocol enumerates the values for transport protocol.
type TransportProtocol string

const (
	// TransportProtocolTCP ...
	TransportProtocolTCP TransportProtocol = "Tcp"
	// TransportProtocolUDP ...
	TransportProtocolUDP TransportProtocol = "Udp"
)

// PossibleTransportProtocolValues returns an array of possible values for the TransportProtocol const type.
func PossibleTransportProtocolValues() []TransportProtocol {
	return []TransportProtocol{TransportProtocolTCP, TransportProtocolUDP}
}

// VirtualNetworkGatewayConnectionStatus enumerates the values for virtual network gateway connection status.
type VirtualNetworkGatewayConnectionStatus string

const (
	// Connected ...
	Connected VirtualNetworkGatewayConnectionStatus = "Connected"
	// Connecting ...
	Connecting VirtualNetworkGatewayConnectionStatus = "Connecting"
	// NotConnected ...
	NotConnected VirtualNetworkGatewayConnectionStatus = "NotConnected"
	// Unknown ...
	Unknown VirtualNetworkGatewayConnectionStatus = "Unknown"
)

// PossibleVirtualNetworkGatewayConnectionStatusValues returns an array of possible values for the VirtualNetworkGatewayConnectionStatus const type.
func PossibleVirtualNetworkGatewayConnectionStatusValues() []VirtualNetworkGatewayConnectionStatus {
	return []VirtualNetworkGatewayConnectionStatus{Connected, Connecting, NotConnected, Unknown}
}

// VirtualNetworkGatewayConnectionType enumerates the values for virtual network gateway connection type.
type VirtualNetworkGatewayConnectionType string

const (
	// ExpressRoute ...
	ExpressRoute VirtualNetworkGatewayConnectionType = "ExpressRoute"
	// IPsec ...
	IPsec VirtualNetworkGatewayConnectionType = "IPsec"
	// Vnet2Vnet ...
	Vnet2Vnet VirtualNetworkGatewayConnectionType = "Vnet2Vnet"
	// VPNClient ...
	VPNClient VirtualNetworkGatewayConnectionType = "VPNClient"
)

// PossibleVirtualNetworkGatewayConnectionTypeValues returns an array of possible values for the VirtualNetworkGatewayConnectionType const type.
func PossibleVirtualNetworkGatewayConnectionTypeValues() []VirtualNetworkGatewayConnectionType {
	return []VirtualNetworkGatewayConnectionType{ExpressRoute, IPsec, Vnet2Vnet, VPNClient}
}

// VirtualNetworkGatewayType enumerates the values for virtual network gateway type.
type VirtualNetworkGatewayType string

const (
	// VirtualNetworkGatewayTypeExpressRoute ...
	VirtualNetworkGatewayTypeExpressRoute VirtualNetworkGatewayType = "ExpressRoute"
	// VirtualNetworkGatewayTypeVpn ...
	VirtualNetworkGatewayTypeVpn VirtualNetworkGatewayType = "Vpn"
)

// PossibleVirtualNetworkGatewayTypeValues returns an array of possible values for the VirtualNetworkGatewayType const type.
func PossibleVirtualNetworkGatewayTypeValues() []VirtualNetworkGatewayType {
	return []VirtualNetworkGatewayType{VirtualNetworkGatewayTypeExpressRoute, VirtualNetworkGatewayTypeVpn}
}

// VpnType enumerates the values for vpn type.
type VpnType string

const (
	// PolicyBased ...
	PolicyBased VpnType = "PolicyBased"
	// RouteBased ...
	RouteBased VpnType = "RouteBased"
)

// PossibleVpnTypeValues returns an array of possible values for the VpnType const type.
func PossibleVpnTypeValues() []VpnType {
	return []VpnType{PolicyBased, RouteBased}
}
