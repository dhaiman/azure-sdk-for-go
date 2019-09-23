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

package frontdoor

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2019-04-01/frontdoor"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ActionType = original.ActionType

const (
	Allow    ActionType = original.Allow
	Block    ActionType = original.Block
	Log      ActionType = original.Log
	Redirect ActionType = original.Redirect
)

type Availability = original.Availability

const (
	Available   Availability = original.Available
	Unavailable Availability = original.Unavailable
)

type BackendEnabledState = original.BackendEnabledState

const (
	Disabled BackendEnabledState = original.Disabled
	Enabled  BackendEnabledState = original.Enabled
)

type CertificateSource = original.CertificateSource

const (
	CertificateSourceAzureKeyVault CertificateSource = original.CertificateSourceAzureKeyVault
	CertificateSourceFrontDoor     CertificateSource = original.CertificateSourceFrontDoor
)

type CertificateType = original.CertificateType

const (
	Dedicated CertificateType = original.Dedicated
)

type CustomHTTPSProvisioningState = original.CustomHTTPSProvisioningState

const (
	CustomHTTPSProvisioningStateDisabled  CustomHTTPSProvisioningState = original.CustomHTTPSProvisioningStateDisabled
	CustomHTTPSProvisioningStateDisabling CustomHTTPSProvisioningState = original.CustomHTTPSProvisioningStateDisabling
	CustomHTTPSProvisioningStateEnabled   CustomHTTPSProvisioningState = original.CustomHTTPSProvisioningStateEnabled
	CustomHTTPSProvisioningStateEnabling  CustomHTTPSProvisioningState = original.CustomHTTPSProvisioningStateEnabling
	CustomHTTPSProvisioningStateFailed    CustomHTTPSProvisioningState = original.CustomHTTPSProvisioningStateFailed
)

type CustomHTTPSProvisioningSubstate = original.CustomHTTPSProvisioningSubstate

const (
	CertificateDeleted                            CustomHTTPSProvisioningSubstate = original.CertificateDeleted
	CertificateDeployed                           CustomHTTPSProvisioningSubstate = original.CertificateDeployed
	DeletingCertificate                           CustomHTTPSProvisioningSubstate = original.DeletingCertificate
	DeployingCertificate                          CustomHTTPSProvisioningSubstate = original.DeployingCertificate
	DomainControlValidationRequestApproved        CustomHTTPSProvisioningSubstate = original.DomainControlValidationRequestApproved
	DomainControlValidationRequestRejected        CustomHTTPSProvisioningSubstate = original.DomainControlValidationRequestRejected
	DomainControlValidationRequestTimedOut        CustomHTTPSProvisioningSubstate = original.DomainControlValidationRequestTimedOut
	IssuingCertificate                            CustomHTTPSProvisioningSubstate = original.IssuingCertificate
	PendingDomainControlValidationREquestApproval CustomHTTPSProvisioningSubstate = original.PendingDomainControlValidationREquestApproval
	SubmittingDomainControlValidationRequest      CustomHTTPSProvisioningSubstate = original.SubmittingDomainControlValidationRequest
)

type CustomRuleEnabledState = original.CustomRuleEnabledState

const (
	CustomRuleEnabledStateDisabled CustomRuleEnabledState = original.CustomRuleEnabledStateDisabled
	CustomRuleEnabledStateEnabled  CustomRuleEnabledState = original.CustomRuleEnabledStateEnabled
)

type DynamicCompressionEnabled = original.DynamicCompressionEnabled

const (
	DynamicCompressionEnabledDisabled DynamicCompressionEnabled = original.DynamicCompressionEnabledDisabled
	DynamicCompressionEnabledEnabled  DynamicCompressionEnabled = original.DynamicCompressionEnabledEnabled
)

type EnabledState = original.EnabledState

const (
	EnabledStateDisabled EnabledState = original.EnabledStateDisabled
	EnabledStateEnabled  EnabledState = original.EnabledStateEnabled
)

type EnforceCertificateNameCheckEnabledState = original.EnforceCertificateNameCheckEnabledState

const (
	EnforceCertificateNameCheckEnabledStateDisabled EnforceCertificateNameCheckEnabledState = original.EnforceCertificateNameCheckEnabledStateDisabled
	EnforceCertificateNameCheckEnabledStateEnabled  EnforceCertificateNameCheckEnabledState = original.EnforceCertificateNameCheckEnabledStateEnabled
)

type ForwardingProtocol = original.ForwardingProtocol

const (
	HTTPOnly     ForwardingProtocol = original.HTTPOnly
	HTTPSOnly    ForwardingProtocol = original.HTTPSOnly
	MatchRequest ForwardingProtocol = original.MatchRequest
)

type ManagedRuleEnabledState = original.ManagedRuleEnabledState

const (
	ManagedRuleEnabledStateDisabled ManagedRuleEnabledState = original.ManagedRuleEnabledStateDisabled
	ManagedRuleEnabledStateEnabled  ManagedRuleEnabledState = original.ManagedRuleEnabledStateEnabled
)

type MatchVariable = original.MatchVariable

const (
	Cookies       MatchVariable = original.Cookies
	PostArgs      MatchVariable = original.PostArgs
	QueryString   MatchVariable = original.QueryString
	RemoteAddr    MatchVariable = original.RemoteAddr
	RequestBody   MatchVariable = original.RequestBody
	RequestHeader MatchVariable = original.RequestHeader
	RequestMethod MatchVariable = original.RequestMethod
	RequestURI    MatchVariable = original.RequestURI
	SocketAddr    MatchVariable = original.SocketAddr
)

type NetworkOperationStatus = original.NetworkOperationStatus

const (
	Failed     NetworkOperationStatus = original.Failed
	InProgress NetworkOperationStatus = original.InProgress
	Succeeded  NetworkOperationStatus = original.Succeeded
)

type OdataType = original.OdataType

const (
	OdataTypeMicrosoftAzureFrontDoorModelsFrontdoorForwardingConfiguration OdataType = original.OdataTypeMicrosoftAzureFrontDoorModelsFrontdoorForwardingConfiguration
	OdataTypeMicrosoftAzureFrontDoorModelsFrontdoorRedirectConfiguration   OdataType = original.OdataTypeMicrosoftAzureFrontDoorModelsFrontdoorRedirectConfiguration
	OdataTypeRouteConfiguration                                            OdataType = original.OdataTypeRouteConfiguration
)

type Operator = original.Operator

const (
	Any                Operator = original.Any
	BeginsWith         Operator = original.BeginsWith
	Contains           Operator = original.Contains
	EndsWith           Operator = original.EndsWith
	Equal              Operator = original.Equal
	GeoMatch           Operator = original.GeoMatch
	GreaterThan        Operator = original.GreaterThan
	GreaterThanOrEqual Operator = original.GreaterThanOrEqual
	IPMatch            Operator = original.IPMatch
	LessThan           Operator = original.LessThan
	LessThanOrEqual    Operator = original.LessThanOrEqual
	RegEx              Operator = original.RegEx
)

type PolicyEnabledState = original.PolicyEnabledState

const (
	PolicyEnabledStateDisabled PolicyEnabledState = original.PolicyEnabledStateDisabled
	PolicyEnabledStateEnabled  PolicyEnabledState = original.PolicyEnabledStateEnabled
)

type PolicyMode = original.PolicyMode

const (
	Detection  PolicyMode = original.Detection
	Prevention PolicyMode = original.Prevention
)

type PolicyResourceState = original.PolicyResourceState

const (
	PolicyResourceStateCreating  PolicyResourceState = original.PolicyResourceStateCreating
	PolicyResourceStateDeleting  PolicyResourceState = original.PolicyResourceStateDeleting
	PolicyResourceStateDisabled  PolicyResourceState = original.PolicyResourceStateDisabled
	PolicyResourceStateDisabling PolicyResourceState = original.PolicyResourceStateDisabling
	PolicyResourceStateEnabled   PolicyResourceState = original.PolicyResourceStateEnabled
	PolicyResourceStateEnabling  PolicyResourceState = original.PolicyResourceStateEnabling
)

type Protocol = original.Protocol

const (
	HTTP  Protocol = original.HTTP
	HTTPS Protocol = original.HTTPS
)

type Query = original.Query

const (
	StripAll  Query = original.StripAll
	StripNone Query = original.StripNone
)

type RedirectProtocol = original.RedirectProtocol

const (
	RedirectProtocolHTTPOnly     RedirectProtocol = original.RedirectProtocolHTTPOnly
	RedirectProtocolHTTPSOnly    RedirectProtocol = original.RedirectProtocolHTTPSOnly
	RedirectProtocolMatchRequest RedirectProtocol = original.RedirectProtocolMatchRequest
)

type RedirectType = original.RedirectType

const (
	Found             RedirectType = original.Found
	Moved             RedirectType = original.Moved
	PermanentRedirect RedirectType = original.PermanentRedirect
	TemporaryRedirect RedirectType = original.TemporaryRedirect
)

type ResourceState = original.ResourceState

const (
	ResourceStateCreating  ResourceState = original.ResourceStateCreating
	ResourceStateDeleting  ResourceState = original.ResourceStateDeleting
	ResourceStateDisabled  ResourceState = original.ResourceStateDisabled
	ResourceStateDisabling ResourceState = original.ResourceStateDisabling
	ResourceStateEnabled   ResourceState = original.ResourceStateEnabled
	ResourceStateEnabling  ResourceState = original.ResourceStateEnabling
)

type ResourceType = original.ResourceType

const (
	MicrosoftNetworkfrontDoors                  ResourceType = original.MicrosoftNetworkfrontDoors
	MicrosoftNetworkfrontDoorsfrontendEndpoints ResourceType = original.MicrosoftNetworkfrontDoorsfrontendEndpoints
)

type RoutingRuleEnabledState = original.RoutingRuleEnabledState

const (
	RoutingRuleEnabledStateDisabled RoutingRuleEnabledState = original.RoutingRuleEnabledStateDisabled
	RoutingRuleEnabledStateEnabled  RoutingRuleEnabledState = original.RoutingRuleEnabledStateEnabled
)

type RuleType = original.RuleType

const (
	MatchRule     RuleType = original.MatchRule
	RateLimitRule RuleType = original.RateLimitRule
)

type SessionAffinityEnabledState = original.SessionAffinityEnabledState

const (
	SessionAffinityEnabledStateDisabled SessionAffinityEnabledState = original.SessionAffinityEnabledStateDisabled
	SessionAffinityEnabledStateEnabled  SessionAffinityEnabledState = original.SessionAffinityEnabledStateEnabled
)

type TLSProtocolType = original.TLSProtocolType

const (
	ServerNameIndication TLSProtocolType = original.ServerNameIndication
)

type TransformType = original.TransformType

const (
	Lowercase   TransformType = original.Lowercase
	RemoveNulls TransformType = original.RemoveNulls
	Trim        TransformType = original.Trim
	Uppercase   TransformType = original.Uppercase
	URLDecode   TransformType = original.URLDecode
	URLEncode   TransformType = original.URLEncode
)

type AzureAsyncOperationResult = original.AzureAsyncOperationResult
type Backend = original.Backend
type BackendPool = original.BackendPool
type BackendPoolListResult = original.BackendPoolListResult
type BackendPoolListResultIterator = original.BackendPoolListResultIterator
type BackendPoolListResultPage = original.BackendPoolListResultPage
type BackendPoolProperties = original.BackendPoolProperties
type BackendPoolUpdateParameters = original.BackendPoolUpdateParameters
type BackendPoolsClient = original.BackendPoolsClient
type BackendPoolsCreateOrUpdateFuture = original.BackendPoolsCreateOrUpdateFuture
type BackendPoolsDeleteFuture = original.BackendPoolsDeleteFuture
type BackendPoolsSettings = original.BackendPoolsSettings
type BaseClient = original.BaseClient
type BasicRouteConfiguration = original.BasicRouteConfiguration
type CacheConfiguration = original.CacheConfiguration
type CertificateSourceParameters = original.CertificateSourceParameters
type CheckNameAvailabilityInput = original.CheckNameAvailabilityInput
type CheckNameAvailabilityOutput = original.CheckNameAvailabilityOutput
type CustomHTTPSConfiguration = original.CustomHTTPSConfiguration
type CustomRule = original.CustomRule
type CustomRuleList = original.CustomRuleList
type EndpointsClient = original.EndpointsClient
type EndpointsPurgeContentFuture = original.EndpointsPurgeContentFuture
type Error = original.Error
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type ForwardingConfiguration = original.ForwardingConfiguration
type FrontDoor = original.FrontDoor
type FrontDoorsClient = original.FrontDoorsClient
type FrontDoorsCreateOrUpdateFutureType = original.FrontDoorsCreateOrUpdateFutureType
type FrontDoorsDeleteFutureType = original.FrontDoorsDeleteFutureType
type FrontendEndpoint = original.FrontendEndpoint
type FrontendEndpointLink = original.FrontendEndpointLink
type FrontendEndpointProperties = original.FrontendEndpointProperties
type FrontendEndpointUpdateParameters = original.FrontendEndpointUpdateParameters
type FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink = original.FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink
type FrontendEndpointsClient = original.FrontendEndpointsClient
type FrontendEndpointsCreateOrUpdateFuture = original.FrontendEndpointsCreateOrUpdateFuture
type FrontendEndpointsDeleteFuture = original.FrontendEndpointsDeleteFuture
type FrontendEndpointsDisableHTTPSFuture = original.FrontendEndpointsDisableHTTPSFuture
type FrontendEndpointsEnableHTTPSFuture = original.FrontendEndpointsEnableHTTPSFuture
type FrontendEndpointsListResult = original.FrontendEndpointsListResult
type FrontendEndpointsListResultIterator = original.FrontendEndpointsListResultIterator
type FrontendEndpointsListResultPage = original.FrontendEndpointsListResultPage
type HealthProbeSettingsClient = original.HealthProbeSettingsClient
type HealthProbeSettingsCreateOrUpdateFuture = original.HealthProbeSettingsCreateOrUpdateFuture
type HealthProbeSettingsDeleteFuture = original.HealthProbeSettingsDeleteFuture
type HealthProbeSettingsListResult = original.HealthProbeSettingsListResult
type HealthProbeSettingsListResultIterator = original.HealthProbeSettingsListResultIterator
type HealthProbeSettingsListResultPage = original.HealthProbeSettingsListResultPage
type HealthProbeSettingsModel = original.HealthProbeSettingsModel
type HealthProbeSettingsProperties = original.HealthProbeSettingsProperties
type HealthProbeSettingsUpdateParameters = original.HealthProbeSettingsUpdateParameters
type KeyVaultCertificateSourceParameters = original.KeyVaultCertificateSourceParameters
type KeyVaultCertificateSourceParametersVault = original.KeyVaultCertificateSourceParametersVault
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type LoadBalancingSettingsClient = original.LoadBalancingSettingsClient
type LoadBalancingSettingsCreateOrUpdateFuture = original.LoadBalancingSettingsCreateOrUpdateFuture
type LoadBalancingSettingsDeleteFuture = original.LoadBalancingSettingsDeleteFuture
type LoadBalancingSettingsListResult = original.LoadBalancingSettingsListResult
type LoadBalancingSettingsListResultIterator = original.LoadBalancingSettingsListResultIterator
type LoadBalancingSettingsListResultPage = original.LoadBalancingSettingsListResultPage
type LoadBalancingSettingsModel = original.LoadBalancingSettingsModel
type LoadBalancingSettingsProperties = original.LoadBalancingSettingsProperties
type LoadBalancingSettingsUpdateParameters = original.LoadBalancingSettingsUpdateParameters
type ManagedRuleDefinition = original.ManagedRuleDefinition
type ManagedRuleGroupDefinition = original.ManagedRuleGroupDefinition
type ManagedRuleGroupOverride = original.ManagedRuleGroupOverride
type ManagedRuleOverride = original.ManagedRuleOverride
type ManagedRuleSet = original.ManagedRuleSet
type ManagedRuleSetDefinition = original.ManagedRuleSetDefinition
type ManagedRuleSetDefinitionList = original.ManagedRuleSetDefinitionList
type ManagedRuleSetDefinitionListIterator = original.ManagedRuleSetDefinitionListIterator
type ManagedRuleSetDefinitionListPage = original.ManagedRuleSetDefinitionListPage
type ManagedRuleSetDefinitionProperties = original.ManagedRuleSetDefinitionProperties
type ManagedRuleSetList = original.ManagedRuleSetList
type ManagedRuleSetsClient = original.ManagedRuleSetsClient
type MatchCondition = original.MatchCondition
type PoliciesClient = original.PoliciesClient
type PoliciesCreateOrUpdateFuture = original.PoliciesCreateOrUpdateFuture
type PoliciesDeleteFuture = original.PoliciesDeleteFuture
type PolicySettings = original.PolicySettings
type Properties = original.Properties
type PurgeParameters = original.PurgeParameters
type RedirectConfiguration = original.RedirectConfiguration
type Resource = original.Resource
type RouteConfiguration = original.RouteConfiguration
type RoutingRule = original.RoutingRule
type RoutingRuleListResult = original.RoutingRuleListResult
type RoutingRuleListResultIterator = original.RoutingRuleListResultIterator
type RoutingRuleListResultPage = original.RoutingRuleListResultPage
type RoutingRuleProperties = original.RoutingRuleProperties
type RoutingRuleUpdateParameters = original.RoutingRuleUpdateParameters
type RoutingRulesClient = original.RoutingRulesClient
type RoutingRulesCreateOrUpdateFuture = original.RoutingRulesCreateOrUpdateFuture
type RoutingRulesDeleteFuture = original.RoutingRulesDeleteFuture
type SubResource = original.SubResource
type TagsObject = original.TagsObject
type UpdateParameters = original.UpdateParameters
type ValidateCustomDomainInput = original.ValidateCustomDomainInput
type ValidateCustomDomainOutput = original.ValidateCustomDomainOutput
type WebApplicationFirewallPolicy = original.WebApplicationFirewallPolicy
type WebApplicationFirewallPolicyList = original.WebApplicationFirewallPolicyList
type WebApplicationFirewallPolicyListIterator = original.WebApplicationFirewallPolicyListIterator
type WebApplicationFirewallPolicyListPage = original.WebApplicationFirewallPolicyListPage
type WebApplicationFirewallPolicyProperties = original.WebApplicationFirewallPolicyProperties

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewBackendPoolListResultIterator(page BackendPoolListResultPage) BackendPoolListResultIterator {
	return original.NewBackendPoolListResultIterator(page)
}
func NewBackendPoolListResultPage(getNextPage func(context.Context, BackendPoolListResult) (BackendPoolListResult, error)) BackendPoolListResultPage {
	return original.NewBackendPoolListResultPage(getNextPage)
}
func NewBackendPoolsClient(subscriptionID string) BackendPoolsClient {
	return original.NewBackendPoolsClient(subscriptionID)
}
func NewBackendPoolsClientWithBaseURI(baseURI string, subscriptionID string) BackendPoolsClient {
	return original.NewBackendPoolsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEndpointsClient(subscriptionID string) EndpointsClient {
	return original.NewEndpointsClient(subscriptionID)
}
func NewEndpointsClientWithBaseURI(baseURI string, subscriptionID string) EndpointsClient {
	return original.NewEndpointsClientWithBaseURI(baseURI, subscriptionID)
}
func NewFrontDoorsClient(subscriptionID string) FrontDoorsClient {
	return original.NewFrontDoorsClient(subscriptionID)
}
func NewFrontDoorsClientWithBaseURI(baseURI string, subscriptionID string) FrontDoorsClient {
	return original.NewFrontDoorsClientWithBaseURI(baseURI, subscriptionID)
}
func NewFrontendEndpointsClient(subscriptionID string) FrontendEndpointsClient {
	return original.NewFrontendEndpointsClient(subscriptionID)
}
func NewFrontendEndpointsClientWithBaseURI(baseURI string, subscriptionID string) FrontendEndpointsClient {
	return original.NewFrontendEndpointsClientWithBaseURI(baseURI, subscriptionID)
}
func NewFrontendEndpointsListResultIterator(page FrontendEndpointsListResultPage) FrontendEndpointsListResultIterator {
	return original.NewFrontendEndpointsListResultIterator(page)
}
func NewFrontendEndpointsListResultPage(getNextPage func(context.Context, FrontendEndpointsListResult) (FrontendEndpointsListResult, error)) FrontendEndpointsListResultPage {
	return original.NewFrontendEndpointsListResultPage(getNextPage)
}
func NewHealthProbeSettingsClient(subscriptionID string) HealthProbeSettingsClient {
	return original.NewHealthProbeSettingsClient(subscriptionID)
}
func NewHealthProbeSettingsClientWithBaseURI(baseURI string, subscriptionID string) HealthProbeSettingsClient {
	return original.NewHealthProbeSettingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewHealthProbeSettingsListResultIterator(page HealthProbeSettingsListResultPage) HealthProbeSettingsListResultIterator {
	return original.NewHealthProbeSettingsListResultIterator(page)
}
func NewHealthProbeSettingsListResultPage(getNextPage func(context.Context, HealthProbeSettingsListResult) (HealthProbeSettingsListResult, error)) HealthProbeSettingsListResultPage {
	return original.NewHealthProbeSettingsListResultPage(getNextPage)
}
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return original.NewListResultIterator(page)
}
func NewListResultPage(getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return original.NewListResultPage(getNextPage)
}
func NewLoadBalancingSettingsClient(subscriptionID string) LoadBalancingSettingsClient {
	return original.NewLoadBalancingSettingsClient(subscriptionID)
}
func NewLoadBalancingSettingsClientWithBaseURI(baseURI string, subscriptionID string) LoadBalancingSettingsClient {
	return original.NewLoadBalancingSettingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewLoadBalancingSettingsListResultIterator(page LoadBalancingSettingsListResultPage) LoadBalancingSettingsListResultIterator {
	return original.NewLoadBalancingSettingsListResultIterator(page)
}
func NewLoadBalancingSettingsListResultPage(getNextPage func(context.Context, LoadBalancingSettingsListResult) (LoadBalancingSettingsListResult, error)) LoadBalancingSettingsListResultPage {
	return original.NewLoadBalancingSettingsListResultPage(getNextPage)
}
func NewManagedRuleSetDefinitionListIterator(page ManagedRuleSetDefinitionListPage) ManagedRuleSetDefinitionListIterator {
	return original.NewManagedRuleSetDefinitionListIterator(page)
}
func NewManagedRuleSetDefinitionListPage(getNextPage func(context.Context, ManagedRuleSetDefinitionList) (ManagedRuleSetDefinitionList, error)) ManagedRuleSetDefinitionListPage {
	return original.NewManagedRuleSetDefinitionListPage(getNextPage)
}
func NewManagedRuleSetsClient(subscriptionID string) ManagedRuleSetsClient {
	return original.NewManagedRuleSetsClient(subscriptionID)
}
func NewManagedRuleSetsClientWithBaseURI(baseURI string, subscriptionID string) ManagedRuleSetsClient {
	return original.NewManagedRuleSetsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPoliciesClient(subscriptionID string) PoliciesClient {
	return original.NewPoliciesClient(subscriptionID)
}
func NewPoliciesClientWithBaseURI(baseURI string, subscriptionID string) PoliciesClient {
	return original.NewPoliciesClientWithBaseURI(baseURI, subscriptionID)
}
func NewRoutingRuleListResultIterator(page RoutingRuleListResultPage) RoutingRuleListResultIterator {
	return original.NewRoutingRuleListResultIterator(page)
}
func NewRoutingRuleListResultPage(getNextPage func(context.Context, RoutingRuleListResult) (RoutingRuleListResult, error)) RoutingRuleListResultPage {
	return original.NewRoutingRuleListResultPage(getNextPage)
}
func NewRoutingRulesClient(subscriptionID string) RoutingRulesClient {
	return original.NewRoutingRulesClient(subscriptionID)
}
func NewRoutingRulesClientWithBaseURI(baseURI string, subscriptionID string) RoutingRulesClient {
	return original.NewRoutingRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWebApplicationFirewallPolicyListIterator(page WebApplicationFirewallPolicyListPage) WebApplicationFirewallPolicyListIterator {
	return original.NewWebApplicationFirewallPolicyListIterator(page)
}
func NewWebApplicationFirewallPolicyListPage(getNextPage func(context.Context, WebApplicationFirewallPolicyList) (WebApplicationFirewallPolicyList, error)) WebApplicationFirewallPolicyListPage {
	return original.NewWebApplicationFirewallPolicyListPage(getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleActionTypeValues() []ActionType {
	return original.PossibleActionTypeValues()
}
func PossibleAvailabilityValues() []Availability {
	return original.PossibleAvailabilityValues()
}
func PossibleBackendEnabledStateValues() []BackendEnabledState {
	return original.PossibleBackendEnabledStateValues()
}
func PossibleCertificateSourceValues() []CertificateSource {
	return original.PossibleCertificateSourceValues()
}
func PossibleCertificateTypeValues() []CertificateType {
	return original.PossibleCertificateTypeValues()
}
func PossibleCustomHTTPSProvisioningStateValues() []CustomHTTPSProvisioningState {
	return original.PossibleCustomHTTPSProvisioningStateValues()
}
func PossibleCustomHTTPSProvisioningSubstateValues() []CustomHTTPSProvisioningSubstate {
	return original.PossibleCustomHTTPSProvisioningSubstateValues()
}
func PossibleCustomRuleEnabledStateValues() []CustomRuleEnabledState {
	return original.PossibleCustomRuleEnabledStateValues()
}
func PossibleDynamicCompressionEnabledValues() []DynamicCompressionEnabled {
	return original.PossibleDynamicCompressionEnabledValues()
}
func PossibleEnabledStateValues() []EnabledState {
	return original.PossibleEnabledStateValues()
}
func PossibleEnforceCertificateNameCheckEnabledStateValues() []EnforceCertificateNameCheckEnabledState {
	return original.PossibleEnforceCertificateNameCheckEnabledStateValues()
}
func PossibleForwardingProtocolValues() []ForwardingProtocol {
	return original.PossibleForwardingProtocolValues()
}
func PossibleManagedRuleEnabledStateValues() []ManagedRuleEnabledState {
	return original.PossibleManagedRuleEnabledStateValues()
}
func PossibleMatchVariableValues() []MatchVariable {
	return original.PossibleMatchVariableValues()
}
func PossibleNetworkOperationStatusValues() []NetworkOperationStatus {
	return original.PossibleNetworkOperationStatusValues()
}
func PossibleOdataTypeValues() []OdataType {
	return original.PossibleOdataTypeValues()
}
func PossibleOperatorValues() []Operator {
	return original.PossibleOperatorValues()
}
func PossiblePolicyEnabledStateValues() []PolicyEnabledState {
	return original.PossiblePolicyEnabledStateValues()
}
func PossiblePolicyModeValues() []PolicyMode {
	return original.PossiblePolicyModeValues()
}
func PossiblePolicyResourceStateValues() []PolicyResourceState {
	return original.PossiblePolicyResourceStateValues()
}
func PossibleProtocolValues() []Protocol {
	return original.PossibleProtocolValues()
}
func PossibleQueryValues() []Query {
	return original.PossibleQueryValues()
}
func PossibleRedirectProtocolValues() []RedirectProtocol {
	return original.PossibleRedirectProtocolValues()
}
func PossibleRedirectTypeValues() []RedirectType {
	return original.PossibleRedirectTypeValues()
}
func PossibleResourceStateValues() []ResourceState {
	return original.PossibleResourceStateValues()
}
func PossibleResourceTypeValues() []ResourceType {
	return original.PossibleResourceTypeValues()
}
func PossibleRoutingRuleEnabledStateValues() []RoutingRuleEnabledState {
	return original.PossibleRoutingRuleEnabledStateValues()
}
func PossibleRuleTypeValues() []RuleType {
	return original.PossibleRuleTypeValues()
}
func PossibleSessionAffinityEnabledStateValues() []SessionAffinityEnabledState {
	return original.PossibleSessionAffinityEnabledStateValues()
}
func PossibleTLSProtocolTypeValues() []TLSProtocolType {
	return original.PossibleTLSProtocolTypeValues()
}
func PossibleTransformTypeValues() []TransformType {
	return original.PossibleTransformTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
