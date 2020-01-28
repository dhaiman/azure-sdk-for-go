package networkapi

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

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2017-03-01/network"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	CheckDNSNameAvailability(ctx context.Context, location string, domainNameLabel string) (result network.DNSNameAvailabilityResult, err error)
}

var _ BaseClientAPI = (*network.BaseClient)(nil)

// ApplicationGatewaysClientAPI contains the set of methods on the ApplicationGatewaysClient type.
type ApplicationGatewaysClientAPI interface {
	BackendHealth(ctx context.Context, resourceGroupName string, applicationGatewayName string, expand string) (result network.ApplicationGatewaysBackendHealthFuture, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, applicationGatewayName string, parameters network.ApplicationGateway) (result network.ApplicationGatewaysCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, applicationGatewayName string) (result network.ApplicationGatewaysDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, applicationGatewayName string) (result network.ApplicationGateway, err error)
	List(ctx context.Context, resourceGroupName string) (result network.ApplicationGatewayListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result network.ApplicationGatewayListResultIterator, err error)
	ListAll(ctx context.Context) (result network.ApplicationGatewayListResultPage, err error)
	ListAllComplete(ctx context.Context) (result network.ApplicationGatewayListResultIterator, err error)
	ListAvailableWafRuleSets(ctx context.Context) (result network.ApplicationGatewayAvailableWafRuleSetsResult, err error)
	Start(ctx context.Context, resourceGroupName string, applicationGatewayName string) (result network.ApplicationGatewaysStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, applicationGatewayName string) (result network.ApplicationGatewaysStopFuture, err error)
}

var _ ApplicationGatewaysClientAPI = (*network.ApplicationGatewaysClient)(nil)

// ExpressRouteCircuitAuthorizationsClientAPI contains the set of methods on the ExpressRouteCircuitAuthorizationsClient type.
type ExpressRouteCircuitAuthorizationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string, authorizationParameters network.ExpressRouteCircuitAuthorization) (result network.ExpressRouteCircuitAuthorizationsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string) (result network.ExpressRouteCircuitAuthorizationsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, circuitName string, authorizationName string) (result network.ExpressRouteCircuitAuthorization, err error)
	List(ctx context.Context, resourceGroupName string, circuitName string) (result network.AuthorizationListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, circuitName string) (result network.AuthorizationListResultIterator, err error)
}

var _ ExpressRouteCircuitAuthorizationsClientAPI = (*network.ExpressRouteCircuitAuthorizationsClient)(nil)

// ExpressRouteCircuitPeeringsClientAPI contains the set of methods on the ExpressRouteCircuitPeeringsClient type.
type ExpressRouteCircuitPeeringsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, peeringParameters network.ExpressRouteCircuitPeering) (result network.ExpressRouteCircuitPeeringsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, circuitName string, peeringName string) (result network.ExpressRouteCircuitPeeringsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, circuitName string, peeringName string) (result network.ExpressRouteCircuitPeering, err error)
	List(ctx context.Context, resourceGroupName string, circuitName string) (result network.ExpressRouteCircuitPeeringListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, circuitName string) (result network.ExpressRouteCircuitPeeringListResultIterator, err error)
}

var _ ExpressRouteCircuitPeeringsClientAPI = (*network.ExpressRouteCircuitPeeringsClient)(nil)

// ExpressRouteCircuitsClientAPI contains the set of methods on the ExpressRouteCircuitsClient type.
type ExpressRouteCircuitsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, parameters network.ExpressRouteCircuit) (result network.ExpressRouteCircuitsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, circuitName string) (result network.ExpressRouteCircuitsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, circuitName string) (result network.ExpressRouteCircuit, err error)
	GetPeeringStats(ctx context.Context, resourceGroupName string, circuitName string, peeringName string) (result network.ExpressRouteCircuitStats, err error)
	GetStats(ctx context.Context, resourceGroupName string, circuitName string) (result network.ExpressRouteCircuitStats, err error)
	List(ctx context.Context, resourceGroupName string) (result network.ExpressRouteCircuitListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result network.ExpressRouteCircuitListResultIterator, err error)
	ListAll(ctx context.Context) (result network.ExpressRouteCircuitListResultPage, err error)
	ListAllComplete(ctx context.Context) (result network.ExpressRouteCircuitListResultIterator, err error)
	ListArpTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string) (result network.ExpressRouteCircuitsListArpTableFuture, err error)
	ListRoutesTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string) (result network.ExpressRouteCircuitsListRoutesTableFuture, err error)
	ListRoutesTableSummary(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string) (result network.ExpressRouteCircuitsListRoutesTableSummaryFuture, err error)
}

var _ ExpressRouteCircuitsClientAPI = (*network.ExpressRouteCircuitsClient)(nil)

// ExpressRouteServiceProvidersClientAPI contains the set of methods on the ExpressRouteServiceProvidersClient type.
type ExpressRouteServiceProvidersClientAPI interface {
	List(ctx context.Context) (result network.ExpressRouteServiceProviderListResultPage, err error)
	ListComplete(ctx context.Context) (result network.ExpressRouteServiceProviderListResultIterator, err error)
}

var _ ExpressRouteServiceProvidersClientAPI = (*network.ExpressRouteServiceProvidersClient)(nil)

// LoadBalancersClientAPI contains the set of methods on the LoadBalancersClient type.
type LoadBalancersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters network.LoadBalancer) (result network.LoadBalancersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, loadBalancerName string) (result network.LoadBalancersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, loadBalancerName string, expand string) (result network.LoadBalancer, err error)
	List(ctx context.Context, resourceGroupName string) (result network.LoadBalancerListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result network.LoadBalancerListResultIterator, err error)
	ListAll(ctx context.Context) (result network.LoadBalancerListResultPage, err error)
	ListAllComplete(ctx context.Context) (result network.LoadBalancerListResultIterator, err error)
}

var _ LoadBalancersClientAPI = (*network.LoadBalancersClient)(nil)

// InterfacesClientAPI contains the set of methods on the InterfacesClient type.
type InterfacesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, networkInterfaceName string, parameters network.Interface) (result network.InterfacesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, networkInterfaceName string) (result network.InterfacesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, networkInterfaceName string, expand string) (result network.Interface, err error)
	GetEffectiveRouteTable(ctx context.Context, resourceGroupName string, networkInterfaceName string) (result network.InterfacesGetEffectiveRouteTableFuture, err error)
	GetVirtualMachineScaleSetNetworkInterface(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, expand string) (result network.Interface, err error)
	List(ctx context.Context, resourceGroupName string) (result network.InterfaceListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result network.InterfaceListResultIterator, err error)
	ListAll(ctx context.Context) (result network.InterfaceListResultPage, err error)
	ListAllComplete(ctx context.Context) (result network.InterfaceListResultIterator, err error)
	ListEffectiveNetworkSecurityGroups(ctx context.Context, resourceGroupName string, networkInterfaceName string) (result network.InterfacesListEffectiveNetworkSecurityGroupsFuture, err error)
	ListVirtualMachineScaleSetNetworkInterfaces(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string) (result network.InterfaceListResultPage, err error)
	ListVirtualMachineScaleSetNetworkInterfacesComplete(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string) (result network.InterfaceListResultIterator, err error)
	ListVirtualMachineScaleSetVMNetworkInterfaces(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string) (result network.InterfaceListResultPage, err error)
	ListVirtualMachineScaleSetVMNetworkInterfacesComplete(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string) (result network.InterfaceListResultIterator, err error)
}

var _ InterfacesClientAPI = (*network.InterfacesClient)(nil)

// SecurityGroupsClientAPI contains the set of methods on the SecurityGroupsClient type.
type SecurityGroupsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters network.SecurityGroup) (result network.SecurityGroupsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, networkSecurityGroupName string) (result network.SecurityGroupsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, expand string) (result network.SecurityGroup, err error)
	List(ctx context.Context, resourceGroupName string) (result network.SecurityGroupListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result network.SecurityGroupListResultIterator, err error)
	ListAll(ctx context.Context) (result network.SecurityGroupListResultPage, err error)
	ListAllComplete(ctx context.Context) (result network.SecurityGroupListResultIterator, err error)
}

var _ SecurityGroupsClientAPI = (*network.SecurityGroupsClient)(nil)

// SecurityRulesClientAPI contains the set of methods on the SecurityRulesClient type.
type SecurityRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, securityRuleParameters network.SecurityRule) (result network.SecurityRulesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string) (result network.SecurityRulesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string) (result network.SecurityRule, err error)
	List(ctx context.Context, resourceGroupName string, networkSecurityGroupName string) (result network.SecurityRuleListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, networkSecurityGroupName string) (result network.SecurityRuleListResultIterator, err error)
}

var _ SecurityRulesClientAPI = (*network.SecurityRulesClient)(nil)

// WatchersClientAPI contains the set of methods on the WatchersClient type.
type WatchersClientAPI interface {
	CheckConnectivity(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.ConnectivityParameters) (result network.WatchersCheckConnectivityFuture, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.Watcher) (result network.Watcher, err error)
	Delete(ctx context.Context, resourceGroupName string, networkWatcherName string) (result network.WatchersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, networkWatcherName string) (result network.Watcher, err error)
	GetFlowLogStatus(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.FlowLogStatusParameters) (result network.WatchersGetFlowLogStatusFuture, err error)
	GetNextHop(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.NextHopParameters) (result network.WatchersGetNextHopFuture, err error)
	GetTopology(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.TopologyParameters) (result network.Topology, err error)
	GetTroubleshooting(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.TroubleshootingParameters) (result network.WatchersGetTroubleshootingFuture, err error)
	GetTroubleshootingResult(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.QueryTroubleshootingParameters) (result network.WatchersGetTroubleshootingResultFuture, err error)
	GetVMSecurityRules(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.SecurityGroupViewParameters) (result network.WatchersGetVMSecurityRulesFuture, err error)
	List(ctx context.Context, resourceGroupName string) (result network.WatcherListResult, err error)
	ListAll(ctx context.Context) (result network.WatcherListResult, err error)
	SetFlowLogConfiguration(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.FlowLogInformation) (result network.WatchersSetFlowLogConfigurationFuture, err error)
	VerifyIPFlow(ctx context.Context, resourceGroupName string, networkWatcherName string, parameters network.VerificationIPFlowParameters) (result network.WatchersVerifyIPFlowFuture, err error)
}

var _ WatchersClientAPI = (*network.WatchersClient)(nil)

// PacketCapturesClientAPI contains the set of methods on the PacketCapturesClient type.
type PacketCapturesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, parameters network.PacketCapture) (result network.PacketCapturesCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string) (result network.PacketCapturesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string) (result network.PacketCaptureResult, err error)
	GetStatus(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string) (result network.PacketCapturesGetStatusFuture, err error)
	List(ctx context.Context, resourceGroupName string, networkWatcherName string) (result network.PacketCaptureListResult, err error)
	Stop(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string) (result network.PacketCapturesStopFuture, err error)
}

var _ PacketCapturesClientAPI = (*network.PacketCapturesClient)(nil)

// PublicIPAddressesClientAPI contains the set of methods on the PublicIPAddressesClient type.
type PublicIPAddressesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters network.PublicIPAddress) (result network.PublicIPAddressesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, publicIPAddressName string) (result network.PublicIPAddressesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, publicIPAddressName string, expand string) (result network.PublicIPAddress, err error)
	GetVirtualMachineScaleSetPublicIPAddress(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, IPConfigurationName string, publicIPAddressName string, expand string) (result network.PublicIPAddress, err error)
	List(ctx context.Context, resourceGroupName string) (result network.PublicIPAddressListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result network.PublicIPAddressListResultIterator, err error)
	ListAll(ctx context.Context) (result network.PublicIPAddressListResultPage, err error)
	ListAllComplete(ctx context.Context) (result network.PublicIPAddressListResultIterator, err error)
	ListVirtualMachineScaleSetPublicIPAddresses(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string) (result network.PublicIPAddressListResultPage, err error)
	ListVirtualMachineScaleSetPublicIPAddressesComplete(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string) (result network.PublicIPAddressListResultIterator, err error)
	ListVirtualMachineScaleSetVMPublicIPAddresses(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, IPConfigurationName string) (result network.PublicIPAddressListResultPage, err error)
	ListVirtualMachineScaleSetVMPublicIPAddressesComplete(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, IPConfigurationName string) (result network.PublicIPAddressListResultIterator, err error)
}

var _ PublicIPAddressesClientAPI = (*network.PublicIPAddressesClient)(nil)

// RouteFiltersClientAPI contains the set of methods on the RouteFiltersClient type.
type RouteFiltersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters network.RouteFilter) (result network.RouteFiltersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, routeFilterName string) (result network.RouteFiltersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, routeFilterName string, expand string) (result network.RouteFilter, err error)
	List(ctx context.Context) (result network.RouteFilterListResultPage, err error)
	ListComplete(ctx context.Context) (result network.RouteFilterListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result network.RouteFilterListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result network.RouteFilterListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters network.PatchRouteFilter) (result network.RouteFiltersUpdateFuture, err error)
}

var _ RouteFiltersClientAPI = (*network.RouteFiltersClient)(nil)

// RouteFilterRulesClientAPI contains the set of methods on the RouteFilterRulesClient type.
type RouteFilterRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, routeFilterRuleParameters network.RouteFilterRule) (result network.RouteFilterRulesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string) (result network.RouteFilterRulesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string) (result network.RouteFilterRule, err error)
	ListByRouteFilter(ctx context.Context, resourceGroupName string, routeFilterName string) (result network.RouteFilterRuleListResultPage, err error)
	ListByRouteFilterComplete(ctx context.Context, resourceGroupName string, routeFilterName string) (result network.RouteFilterRuleListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, routeFilterName string, ruleName string, routeFilterRuleParameters network.PatchRouteFilterRule) (result network.RouteFilterRulesUpdateFuture, err error)
}

var _ RouteFilterRulesClientAPI = (*network.RouteFilterRulesClient)(nil)

// RouteTablesClientAPI contains the set of methods on the RouteTablesClient type.
type RouteTablesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, routeTableName string, parameters network.RouteTable) (result network.RouteTablesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, routeTableName string) (result network.RouteTablesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, routeTableName string, expand string) (result network.RouteTable, err error)
	List(ctx context.Context, resourceGroupName string) (result network.RouteTableListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result network.RouteTableListResultIterator, err error)
	ListAll(ctx context.Context) (result network.RouteTableListResultPage, err error)
	ListAllComplete(ctx context.Context) (result network.RouteTableListResultIterator, err error)
}

var _ RouteTablesClientAPI = (*network.RouteTablesClient)(nil)

// RoutesClientAPI contains the set of methods on the RoutesClient type.
type RoutesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, routeTableName string, routeName string, routeParameters network.Route) (result network.RoutesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, routeTableName string, routeName string) (result network.RoutesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, routeTableName string, routeName string) (result network.Route, err error)
	List(ctx context.Context, resourceGroupName string, routeTableName string) (result network.RouteListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, routeTableName string) (result network.RouteListResultIterator, err error)
}

var _ RoutesClientAPI = (*network.RoutesClient)(nil)

// BgpServiceCommunitiesClientAPI contains the set of methods on the BgpServiceCommunitiesClient type.
type BgpServiceCommunitiesClientAPI interface {
	List(ctx context.Context) (result network.BgpServiceCommunityListResultPage, err error)
	ListComplete(ctx context.Context) (result network.BgpServiceCommunityListResultIterator, err error)
}

var _ BgpServiceCommunitiesClientAPI = (*network.BgpServiceCommunitiesClient)(nil)

// UsagesClientAPI contains the set of methods on the UsagesClient type.
type UsagesClientAPI interface {
	List(ctx context.Context, location string) (result network.UsagesListResultPage, err error)
	ListComplete(ctx context.Context, location string) (result network.UsagesListResultIterator, err error)
}

var _ UsagesClientAPI = (*network.UsagesClient)(nil)

// VirtualNetworksClientAPI contains the set of methods on the VirtualNetworksClient type.
type VirtualNetworksClientAPI interface {
	CheckIPAddressAvailability(ctx context.Context, resourceGroupName string, virtualNetworkName string, IPAddress string) (result network.IPAddressAvailabilityResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, parameters network.VirtualNetwork) (result network.VirtualNetworksCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, virtualNetworkName string) (result network.VirtualNetworksDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, expand string) (result network.VirtualNetwork, err error)
	List(ctx context.Context, resourceGroupName string) (result network.VirtualNetworkListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result network.VirtualNetworkListResultIterator, err error)
	ListAll(ctx context.Context) (result network.VirtualNetworkListResultPage, err error)
	ListAllComplete(ctx context.Context) (result network.VirtualNetworkListResultIterator, err error)
	ListUsage(ctx context.Context, resourceGroupName string, virtualNetworkName string) (result network.VirtualNetworkListUsageResultPage, err error)
	ListUsageComplete(ctx context.Context, resourceGroupName string, virtualNetworkName string) (result network.VirtualNetworkListUsageResultIterator, err error)
}

var _ VirtualNetworksClientAPI = (*network.VirtualNetworksClient)(nil)

// SubnetsClientAPI contains the set of methods on the SubnetsClient type.
type SubnetsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters network.Subnet) (result network.SubnetsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string) (result network.SubnetsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, expand string) (result network.Subnet, err error)
	List(ctx context.Context, resourceGroupName string, virtualNetworkName string) (result network.SubnetListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, virtualNetworkName string) (result network.SubnetListResultIterator, err error)
}

var _ SubnetsClientAPI = (*network.SubnetsClient)(nil)

// VirtualNetworkPeeringsClientAPI contains the set of methods on the VirtualNetworkPeeringsClient type.
type VirtualNetworkPeeringsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string, virtualNetworkPeeringParameters network.VirtualNetworkPeering) (result network.VirtualNetworkPeeringsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string) (result network.VirtualNetworkPeeringsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, virtualNetworkPeeringName string) (result network.VirtualNetworkPeering, err error)
	List(ctx context.Context, resourceGroupName string, virtualNetworkName string) (result network.VirtualNetworkPeeringListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, virtualNetworkName string) (result network.VirtualNetworkPeeringListResultIterator, err error)
}

var _ VirtualNetworkPeeringsClientAPI = (*network.VirtualNetworkPeeringsClient)(nil)

// VirtualNetworkGatewaysClientAPI contains the set of methods on the VirtualNetworkGatewaysClient type.
type VirtualNetworkGatewaysClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, parameters network.VirtualNetworkGateway) (result network.VirtualNetworkGatewaysCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string) (result network.VirtualNetworkGatewaysDeleteFuture, err error)
	Generatevpnclientpackage(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, parameters network.VpnClientParameters) (result network.VirtualNetworkGatewaysGeneratevpnclientpackageFuture, err error)
	Get(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string) (result network.VirtualNetworkGateway, err error)
	GetAdvertisedRoutes(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, peer string) (result network.VirtualNetworkGatewaysGetAdvertisedRoutesFuture, err error)
	GetBgpPeerStatus(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, peer string) (result network.VirtualNetworkGatewaysGetBgpPeerStatusFuture, err error)
	GetLearnedRoutes(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string) (result network.VirtualNetworkGatewaysGetLearnedRoutesFuture, err error)
	List(ctx context.Context, resourceGroupName string) (result network.VirtualNetworkGatewayListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result network.VirtualNetworkGatewayListResultIterator, err error)
	Reset(ctx context.Context, resourceGroupName string, virtualNetworkGatewayName string, gatewayVip string) (result network.VirtualNetworkGatewaysResetFuture, err error)
}

var _ VirtualNetworkGatewaysClientAPI = (*network.VirtualNetworkGatewaysClient)(nil)

// VirtualNetworkGatewayConnectionsClientAPI contains the set of methods on the VirtualNetworkGatewayConnectionsClient type.
type VirtualNetworkGatewayConnectionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters network.VirtualNetworkGatewayConnection) (result network.VirtualNetworkGatewayConnectionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (result network.VirtualNetworkGatewayConnectionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (result network.VirtualNetworkGatewayConnection, err error)
	GetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string) (result network.ConnectionSharedKey, err error)
	List(ctx context.Context, resourceGroupName string) (result network.VirtualNetworkGatewayConnectionListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result network.VirtualNetworkGatewayConnectionListResultIterator, err error)
	ResetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters network.ConnectionResetSharedKey) (result network.VirtualNetworkGatewayConnectionsResetSharedKeyFuture, err error)
	SetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters network.ConnectionSharedKey) (result network.VirtualNetworkGatewayConnectionsSetSharedKeyFuture, err error)
}

var _ VirtualNetworkGatewayConnectionsClientAPI = (*network.VirtualNetworkGatewayConnectionsClient)(nil)

// LocalNetworkGatewaysClientAPI contains the set of methods on the LocalNetworkGatewaysClient type.
type LocalNetworkGatewaysClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters network.LocalNetworkGateway) (result network.LocalNetworkGatewaysCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, localNetworkGatewayName string) (result network.LocalNetworkGatewaysDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, localNetworkGatewayName string) (result network.LocalNetworkGateway, err error)
	List(ctx context.Context, resourceGroupName string) (result network.LocalNetworkGatewayListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result network.LocalNetworkGatewayListResultIterator, err error)
}

var _ LocalNetworkGatewaysClientAPI = (*network.LocalNetworkGatewaysClient)(nil)
