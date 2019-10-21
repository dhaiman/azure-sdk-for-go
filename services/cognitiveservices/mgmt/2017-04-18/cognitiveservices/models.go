package cognitiveservices

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
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/mgmt/2017-04-18/cognitiveservices"

// KeyName enumerates the values for key name.
type KeyName string

const (
	// Key1 ...
	Key1 KeyName = "Key1"
	// Key2 ...
	Key2 KeyName = "Key2"
)

// PossibleKeyNameValues returns an array of possible values for the KeyName const type.
func PossibleKeyNameValues() []KeyName {
	return []KeyName{Key1, Key2}
}

// NetworkRuleAction enumerates the values for network rule action.
type NetworkRuleAction string

const (
	// Allow ...
	Allow NetworkRuleAction = "Allow"
	// Deny ...
	Deny NetworkRuleAction = "Deny"
)

// PossibleNetworkRuleActionValues returns an array of possible values for the NetworkRuleAction const type.
func PossibleNetworkRuleActionValues() []NetworkRuleAction {
	return []NetworkRuleAction{Allow, Deny}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Moving ...
	Moving ProvisioningState = "Moving"
	// ResolvingDNS ...
	ResolvingDNS ProvisioningState = "ResolvingDNS"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Creating, Deleting, Failed, Moving, ResolvingDNS, Succeeded}
}

// QuotaUsageStatus enumerates the values for quota usage status.
type QuotaUsageStatus string

const (
	// Blocked ...
	Blocked QuotaUsageStatus = "Blocked"
	// Included ...
	Included QuotaUsageStatus = "Included"
	// InOverage ...
	InOverage QuotaUsageStatus = "InOverage"
	// Unknown ...
	Unknown QuotaUsageStatus = "Unknown"
)

// PossibleQuotaUsageStatusValues returns an array of possible values for the QuotaUsageStatus const type.
func PossibleQuotaUsageStatusValues() []QuotaUsageStatus {
	return []QuotaUsageStatus{Blocked, Included, InOverage, Unknown}
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

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// Free ...
	Free SkuTier = "Free"
	// Premium ...
	Premium SkuTier = "Premium"
	// Standard ...
	Standard SkuTier = "Standard"
)

// PossibleSkuTierValues returns an array of possible values for the SkuTier const type.
func PossibleSkuTierValues() []SkuTier {
	return []SkuTier{Free, Premium, Standard}
}

// UnitType enumerates the values for unit type.
type UnitType string

const (
	// Bytes ...
	Bytes UnitType = "Bytes"
	// BytesPerSecond ...
	BytesPerSecond UnitType = "BytesPerSecond"
	// Count ...
	Count UnitType = "Count"
	// CountPerSecond ...
	CountPerSecond UnitType = "CountPerSecond"
	// Milliseconds ...
	Milliseconds UnitType = "Milliseconds"
	// Percent ...
	Percent UnitType = "Percent"
	// Seconds ...
	Seconds UnitType = "Seconds"
)

// PossibleUnitTypeValues returns an array of possible values for the UnitType const type.
func PossibleUnitTypeValues() []UnitType {
	return []UnitType{Bytes, BytesPerSecond, Count, CountPerSecond, Milliseconds, Percent, Seconds}
}

// Account cognitive Services Account is an Azure resource representing the provisioned account, its type,
// location and SKU.
type Account struct {
	autorest.Response `json:"-"`
	// Etag - READ-ONLY; Entity Tag
	Etag *string `json:"etag,omitempty"`
	// ID - READ-ONLY; The id of the created account
	ID *string `json:"id,omitempty"`
	// Kind - The Kind of the resource.
	Kind *string `json:"kind,omitempty"`
	// Location - The location of the resource
	Location *string `json:"location,omitempty"`
	// Name - READ-ONLY; The name of the created account
	Name *string `json:"name,omitempty"`
	// Properties - Properties of Cognitive Services account.
	Properties *AccountProperties `json:"properties,omitempty"`
	// Sku - The SKU of Cognitive Services account.
	Sku *Sku `json:"sku,omitempty"`
	// Tags - Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]*string `json:"tags"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Account.
func (a Account) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if a.Kind != nil {
		objectMap["kind"] = a.Kind
	}
	if a.Location != nil {
		objectMap["location"] = a.Location
	}
	if a.Properties != nil {
		objectMap["properties"] = a.Properties
	}
	if a.Sku != nil {
		objectMap["sku"] = a.Sku
	}
	if a.Tags != nil {
		objectMap["tags"] = a.Tags
	}
	return json.Marshal(objectMap)
}

// AccountAPIProperties the api properties for special APIs.
type AccountAPIProperties struct {
	// QnaRuntimeEndpoint - (QnAMaker Only) The runtime endpoint of QnAMaker.
	QnaRuntimeEndpoint *string `json:"qnaRuntimeEndpoint,omitempty"`
	// StatisticsEnabled - (Bing Search Only) The flag to enable statistics of Bing Search.
	StatisticsEnabled *bool `json:"statisticsEnabled,omitempty"`
	// EventHubConnectionString - (Personalization Only) The flag to enable statistics of Bing Search.
	EventHubConnectionString *string `json:"eventHubConnectionString,omitempty"`
	// StorageAccountConnectionString - (Personalization Only) The storage account connection string.
	StorageAccountConnectionString *string `json:"storageAccountConnectionString,omitempty"`
}

// AccountEnumerateSkusResult the list of cognitive services accounts operation response.
type AccountEnumerateSkusResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; Gets the list of Cognitive Services accounts and their properties.
	Value *[]ResourceAndSku `json:"value,omitempty"`
}

// AccountKeys the access keys for the cognitive services account.
type AccountKeys struct {
	autorest.Response `json:"-"`
	// Key1 - Gets the value of key 1.
	Key1 *string `json:"key1,omitempty"`
	// Key2 - Gets the value of key 2.
	Key2 *string `json:"key2,omitempty"`
}

// AccountListResult the list of cognitive services accounts operation response.
type AccountListResult struct {
	autorest.Response `json:"-"`
	// NextLink - The link used to get the next page of accounts.
	NextLink *string `json:"nextLink,omitempty"`
	// Value - READ-ONLY; Gets the list of Cognitive Services accounts and their properties.
	Value *[]Account `json:"value,omitempty"`
}

// AccountListResultIterator provides access to a complete listing of Account values.
type AccountListResultIterator struct {
	i    int
	page AccountListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *AccountListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *AccountListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter AccountListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter AccountListResultIterator) Response() AccountListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter AccountListResultIterator) Value() Account {
	if !iter.page.NotDone() {
		return Account{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the AccountListResultIterator type.
func NewAccountListResultIterator(page AccountListResultPage) AccountListResultIterator {
	return AccountListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (alr AccountListResult) IsEmpty() bool {
	return alr.Value == nil || len(*alr.Value) == 0
}

// accountListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (alr AccountListResult) accountListResultPreparer(ctx context.Context) (*http.Request, error) {
	if alr.NextLink == nil || len(to.String(alr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(alr.NextLink)))
}

// AccountListResultPage contains a page of Account values.
type AccountListResultPage struct {
	fn  func(context.Context, AccountListResult) (AccountListResult, error)
	alr AccountListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *AccountListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AccountListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.alr)
	if err != nil {
		return err
	}
	page.alr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *AccountListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page AccountListResultPage) NotDone() bool {
	return !page.alr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page AccountListResultPage) Response() AccountListResult {
	return page.alr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page AccountListResultPage) Values() []Account {
	if page.alr.IsEmpty() {
		return nil
	}
	return *page.alr.Value
}

// Creates a new instance of the AccountListResultPage type.
func NewAccountListResultPage(getNextPage func(context.Context, AccountListResult) (AccountListResult, error)) AccountListResultPage {
	return AccountListResultPage{fn: getNextPage}
}

// AccountProperties properties of Cognitive Services account.
type AccountProperties struct {
	// ProvisioningState - READ-ONLY; Gets the status of the cognitive services account at the time the operation was called. Possible values include: 'Creating', 'ResolvingDNS', 'Moving', 'Deleting', 'Succeeded', 'Failed'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// Endpoint - READ-ONLY; Endpoint of the created account.
	Endpoint *string `json:"endpoint,omitempty"`
	// InternalID - READ-ONLY; The internal identifier.
	InternalID *string `json:"internalId,omitempty"`
	// CustomSubDomainName - Optional subdomain name used for token-based authentication.
	CustomSubDomainName *string `json:"customSubDomainName,omitempty"`
	// NetworkAcls - A collection of rules governing the accessibility from specific network locations.
	NetworkAcls *NetworkRuleSet `json:"networkAcls,omitempty"`
	// APIProperties - The api properties for special APIs.
	APIProperties *AccountAPIProperties `json:"apiProperties,omitempty"`
}

// CheckDomainAvailabilityParameter check Domain availability parameter.
type CheckDomainAvailabilityParameter struct {
	// SubdomainName - The subdomain name to use.
	SubdomainName *string `json:"subdomainName,omitempty"`
	// Type - The Type of the resource.
	Type *string `json:"type,omitempty"`
}

// CheckDomainAvailabilityResult check Domain availability result.
type CheckDomainAvailabilityResult struct {
	autorest.Response `json:"-"`
	// IsSubdomainAvailable - Indicates the given SKU is available or not.
	IsSubdomainAvailable *bool `json:"isSubdomainAvailable,omitempty"`
	// Reason - Reason why the SKU is not available.
	Reason *string `json:"reason,omitempty"`
	// SubdomainName - The subdomain name to use.
	SubdomainName *string `json:"subdomainName,omitempty"`
	// Type - The Type of the resource.
	Type *string `json:"type,omitempty"`
}

// CheckSkuAvailabilityParameter check SKU availability parameter.
type CheckSkuAvailabilityParameter struct {
	// Skus - The SKU of the resource.
	Skus *[]string `json:"skus,omitempty"`
	// Kind - The Kind of the resource.
	Kind *string `json:"kind,omitempty"`
	// Type - The Type of the resource.
	Type *string `json:"type,omitempty"`
}

// CheckSkuAvailabilityResult check SKU availability result.
type CheckSkuAvailabilityResult struct {
	// Kind - The Kind of the resource.
	Kind *string `json:"kind,omitempty"`
	// Type - The Type of the resource.
	Type *string `json:"type,omitempty"`
	// SkuName - The SKU of Cognitive Services account.
	SkuName *string `json:"skuName,omitempty"`
	// SkuAvailable - Indicates the given SKU is available or not.
	SkuAvailable *bool `json:"skuAvailable,omitempty"`
	// Reason - Reason why the SKU is not available.
	Reason *string `json:"reason,omitempty"`
	// Message - Additional error message.
	Message *string `json:"message,omitempty"`
}

// CheckSkuAvailabilityResultList check SKU availability result list.
type CheckSkuAvailabilityResultList struct {
	autorest.Response `json:"-"`
	// Value - Check SKU availability result list.
	Value *[]CheckSkuAvailabilityResult `json:"value,omitempty"`
}

// Error cognitive Services error object.
type Error struct {
	// Error - The error body.
	Error *ErrorBody `json:"error,omitempty"`
}

// ErrorBody cognitive Services error body.
type ErrorBody struct {
	// Code - error code
	Code *string `json:"code,omitempty"`
	// Message - error message
	Message *string `json:"message,omitempty"`
}

// IPRule a rule governing the accessibility from a specific ip address or ip range.
type IPRule struct {
	// Value - An IPv4 address range in CIDR notation, such as '124.56.78.91' (simple IP address) or '124.56.78.0/24' (all addresses that start with 124.56.78).
	Value *string `json:"value,omitempty"`
}

// MetricName a metric name.
type MetricName struct {
	// Value - READ-ONLY; The name of the metric.
	Value *string `json:"value,omitempty"`
	// LocalizedValue - READ-ONLY; The friendly name of the metric.
	LocalizedValue *string `json:"localizedValue,omitempty"`
}

// NetworkRuleSet a set of rules governing the network accessibility.
type NetworkRuleSet struct {
	// DefaultAction - The default action when no rule from ipRules and from virtualNetworkRules match. This is only used after the bypass property has been evaluated. Possible values include: 'Allow', 'Deny'
	DefaultAction NetworkRuleAction `json:"defaultAction,omitempty"`
	// IPRules - The list of IP address rules.
	IPRules *[]IPRule `json:"ipRules,omitempty"`
	// VirtualNetworkRules - The list of virtual network rules.
	VirtualNetworkRules *[]VirtualNetworkRule `json:"virtualNetworkRules,omitempty"`
}

// OperationDisplayInfo the operation supported by Cognitive Services.
type OperationDisplayInfo struct {
	// Description - The description of the operation.
	Description *string `json:"description,omitempty"`
	// Operation - The action that users can perform, based on their permission level.
	Operation *string `json:"operation,omitempty"`
	// Provider - Service provider: Microsoft Cognitive Services.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
}

// OperationEntity the operation supported by Cognitive Services.
type OperationEntity struct {
	// Name - Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The operation supported by Cognitive Services.
	Display *OperationDisplayInfo `json:"display,omitempty"`
	// Origin - The origin of the operation.
	Origin *string `json:"origin,omitempty"`
	// Properties - Additional properties.
	Properties interface{} `json:"properties,omitempty"`
}

// OperationEntityListResult the list of cognitive services accounts operation response.
type OperationEntityListResult struct {
	autorest.Response `json:"-"`
	// NextLink - The link used to get the next page of operations.
	NextLink *string `json:"nextLink,omitempty"`
	// Value - The list of operations.
	Value *[]OperationEntity `json:"value,omitempty"`
}

// OperationEntityListResultIterator provides access to a complete listing of OperationEntity values.
type OperationEntityListResultIterator struct {
	i    int
	page OperationEntityListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationEntityListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationEntityListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationEntityListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationEntityListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationEntityListResultIterator) Response() OperationEntityListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationEntityListResultIterator) Value() OperationEntity {
	if !iter.page.NotDone() {
		return OperationEntity{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationEntityListResultIterator type.
func NewOperationEntityListResultIterator(page OperationEntityListResultPage) OperationEntityListResultIterator {
	return OperationEntityListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (oelr OperationEntityListResult) IsEmpty() bool {
	return oelr.Value == nil || len(*oelr.Value) == 0
}

// operationEntityListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (oelr OperationEntityListResult) operationEntityListResultPreparer(ctx context.Context) (*http.Request, error) {
	if oelr.NextLink == nil || len(to.String(oelr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(oelr.NextLink)))
}

// OperationEntityListResultPage contains a page of OperationEntity values.
type OperationEntityListResultPage struct {
	fn   func(context.Context, OperationEntityListResult) (OperationEntityListResult, error)
	oelr OperationEntityListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationEntityListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationEntityListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.oelr)
	if err != nil {
		return err
	}
	page.oelr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationEntityListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationEntityListResultPage) NotDone() bool {
	return !page.oelr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationEntityListResultPage) Response() OperationEntityListResult {
	return page.oelr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationEntityListResultPage) Values() []OperationEntity {
	if page.oelr.IsEmpty() {
		return nil
	}
	return *page.oelr.Value
}

// Creates a new instance of the OperationEntityListResultPage type.
func NewOperationEntityListResultPage(getNextPage func(context.Context, OperationEntityListResult) (OperationEntityListResult, error)) OperationEntityListResultPage {
	return OperationEntityListResultPage{fn: getNextPage}
}

// RegenerateKeyParameters regenerate key parameters.
type RegenerateKeyParameters struct {
	// KeyName - key name to generate (Key1|Key2). Possible values include: 'Key1', 'Key2'
	KeyName KeyName `json:"keyName,omitempty"`
}

// ResourceAndSku cognitive Services resource type and SKU.
type ResourceAndSku struct {
	// ResourceType - Resource Namespace and Type
	ResourceType *string `json:"resourceType,omitempty"`
	// Sku - The SKU of Cognitive Services account.
	Sku *Sku `json:"sku,omitempty"`
}

// ResourceSku describes an available Cognitive Services SKU.
type ResourceSku struct {
	// ResourceType - READ-ONLY; The type of resource the SKU applies to.
	ResourceType *string `json:"resourceType,omitempty"`
	// Name - READ-ONLY; The name of SKU.
	Name *string `json:"name,omitempty"`
	// Tier - READ-ONLY; Specifies the tier of Cognitive Services account.
	Tier *string `json:"tier,omitempty"`
	// Kind - READ-ONLY; The Kind of resources that are supported in this SKU.
	Kind *string `json:"kind,omitempty"`
	// Locations - READ-ONLY; The set of locations that the SKU is available.
	Locations *[]string `json:"locations,omitempty"`
	// Restrictions - READ-ONLY; The restrictions because of which SKU cannot be used. This is empty if there are no restrictions.
	Restrictions *[]ResourceSkuRestrictions `json:"restrictions,omitempty"`
}

// ResourceSkuRestrictionInfo ...
type ResourceSkuRestrictionInfo struct {
	// Locations - READ-ONLY; Locations where the SKU is restricted
	Locations *[]string `json:"locations,omitempty"`
	// Zones - READ-ONLY; List of availability zones where the SKU is restricted.
	Zones *[]string `json:"zones,omitempty"`
}

// ResourceSkuRestrictions describes restrictions of a SKU.
type ResourceSkuRestrictions struct {
	// Type - READ-ONLY; The type of restrictions. Possible values include: 'Location', 'Zone'
	Type ResourceSkuRestrictionsType `json:"type,omitempty"`
	// Values - READ-ONLY; The value of restrictions. If the restriction type is set to location. This would be different locations where the SKU is restricted.
	Values *[]string `json:"values,omitempty"`
	// RestrictionInfo - READ-ONLY; The information about the restriction where the SKU cannot be used.
	RestrictionInfo *ResourceSkuRestrictionInfo `json:"restrictionInfo,omitempty"`
	// ReasonCode - READ-ONLY; The reason for restriction. Possible values include: 'QuotaID', 'NotAvailableForSubscription'
	ReasonCode ResourceSkuRestrictionsReasonCode `json:"reasonCode,omitempty"`
}

// ResourceSkusResult the Get Skus operation response.
type ResourceSkusResult struct {
	autorest.Response `json:"-"`
	// Value - The list of skus available for the subscription.
	Value *[]ResourceSku `json:"value,omitempty"`
	// NextLink - The uri to fetch the next page of Skus.
	NextLink *string `json:"nextLink,omitempty"`
}

// ResourceSkusResultIterator provides access to a complete listing of ResourceSku values.
type ResourceSkusResultIterator struct {
	i    int
	page ResourceSkusResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ResourceSkusResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceSkusResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *ResourceSkusResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ResourceSkusResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ResourceSkusResultIterator) Response() ResourceSkusResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ResourceSkusResultIterator) Value() ResourceSku {
	if !iter.page.NotDone() {
		return ResourceSku{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ResourceSkusResultIterator type.
func NewResourceSkusResultIterator(page ResourceSkusResultPage) ResourceSkusResultIterator {
	return ResourceSkusResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (rsr ResourceSkusResult) IsEmpty() bool {
	return rsr.Value == nil || len(*rsr.Value) == 0
}

// resourceSkusResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (rsr ResourceSkusResult) resourceSkusResultPreparer(ctx context.Context) (*http.Request, error) {
	if rsr.NextLink == nil || len(to.String(rsr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(rsr.NextLink)))
}

// ResourceSkusResultPage contains a page of ResourceSku values.
type ResourceSkusResultPage struct {
	fn  func(context.Context, ResourceSkusResult) (ResourceSkusResult, error)
	rsr ResourceSkusResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ResourceSkusResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceSkusResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.rsr)
	if err != nil {
		return err
	}
	page.rsr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ResourceSkusResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ResourceSkusResultPage) NotDone() bool {
	return !page.rsr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ResourceSkusResultPage) Response() ResourceSkusResult {
	return page.rsr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ResourceSkusResultPage) Values() []ResourceSku {
	if page.rsr.IsEmpty() {
		return nil
	}
	return *page.rsr.Value
}

// Creates a new instance of the ResourceSkusResultPage type.
func NewResourceSkusResultPage(getNextPage func(context.Context, ResourceSkusResult) (ResourceSkusResult, error)) ResourceSkusResultPage {
	return ResourceSkusResultPage{fn: getNextPage}
}

// Sku the SKU of the cognitive services account.
type Sku struct {
	// Name - Gets or sets the sku name. Required for account creation, optional for update.
	Name *string `json:"name,omitempty"`
	// Tier - READ-ONLY; Gets the sku tier. This is based on the SKU name. Possible values include: 'Free', 'Standard', 'Premium'
	Tier SkuTier `json:"tier,omitempty"`
}

// Usage the usage data for a usage request.
type Usage struct {
	// Unit - The unit of the metric. Possible values include: 'Count', 'Bytes', 'Seconds', 'Percent', 'CountPerSecond', 'BytesPerSecond', 'Milliseconds'
	Unit UnitType `json:"unit,omitempty"`
	// Name - READ-ONLY; The name information for the metric.
	Name *MetricName `json:"name,omitempty"`
	// QuotaPeriod - READ-ONLY; The quota period used to summarize the usage values.
	QuotaPeriod *string `json:"quotaPeriod,omitempty"`
	// Limit - READ-ONLY; Maximum value for this metric.
	Limit *float64 `json:"limit,omitempty"`
	// CurrentValue - READ-ONLY; Current value for this metric.
	CurrentValue *float64 `json:"currentValue,omitempty"`
	// NextResetTime - READ-ONLY; Next reset time for current quota.
	NextResetTime *string `json:"nextResetTime,omitempty"`
	// Status - Cognitive Services account quota usage status. Possible values include: 'Included', 'Blocked', 'InOverage', 'Unknown'
	Status QuotaUsageStatus `json:"status,omitempty"`
}

// UsagesResult the response to a list usage request.
type UsagesResult struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; The list of usages for Cognitive Service account.
	Value *[]Usage `json:"value,omitempty"`
}

// VirtualNetworkRule a rule governing the accessibility from a specific virtual network.
type VirtualNetworkRule struct {
	// ID - Full resource id of a vnet subnet, such as '/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1'.
	ID *string `json:"id,omitempty"`
	// State - Gets the state of virtual network rule.
	State *string `json:"state,omitempty"`
	// IgnoreMissingVnetServiceEndpoint - Ignore missing vnet service endpoint or not.
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty"`
}
