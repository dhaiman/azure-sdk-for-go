package keyvault

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
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"github.com/satori/go.uuid"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2016-10-01/keyvault"

// AccessPolicyEntry an identity that have access to the key vault. All identities in the array must use
// the same tenant ID as the key vault's tenant ID.
type AccessPolicyEntry struct {
	// TenantID - The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantID *uuid.UUID `json:"tenantId,omitempty"`
	// ObjectID - The object ID of a user, service principal or security group in the Azure Active Directory tenant for the vault. The object ID must be unique for the list of access policies.
	ObjectID *string `json:"objectId,omitempty"`
	// ApplicationID -  Application ID of the client making request on behalf of a principal
	ApplicationID *uuid.UUID `json:"applicationId,omitempty"`
	// Permissions - Permissions the identity has for keys, secrets and certificates.
	Permissions *Permissions `json:"permissions,omitempty"`
}

// CheckNameAvailabilityResult the CheckNameAvailability operation response.
type CheckNameAvailabilityResult struct {
	autorest.Response `json:"-"`
	// NameAvailable - READ-ONLY; A boolean value that indicates whether the name is available for you to use. If true, the name is available. If false, the name has already been taken or is invalid and cannot be used.
	NameAvailable *bool `json:"nameAvailable,omitempty"`
	// Reason - READ-ONLY; The reason that a vault name could not be used. The Reason element is only returned if NameAvailable is false. Possible values include: 'AccountNameInvalid', 'AlreadyExists'
	Reason Reason `json:"reason,omitempty"`
	// Message - READ-ONLY; An error message explaining the Reason value in more detail.
	Message *string `json:"message,omitempty"`
}

// DeletedVault deleted vault information with extended details.
type DeletedVault struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; The resource ID for the deleted key vault.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the key vault.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The resource type of the key vault.
	Type *string `json:"type,omitempty"`
	// Properties - Properties of the vault
	Properties *DeletedVaultProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for DeletedVault.
func (dv DeletedVault) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if dv.Properties != nil {
		objectMap["properties"] = dv.Properties
	}
	return json.Marshal(objectMap)
}

// DeletedVaultListResult list of vaults
type DeletedVaultListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of deleted vaults.
	Value *[]DeletedVault `json:"value,omitempty"`
	// NextLink - The URL to get the next set of deleted vaults.
	NextLink *string `json:"nextLink,omitempty"`
}

// DeletedVaultListResultIterator provides access to a complete listing of DeletedVault values.
type DeletedVaultListResultIterator struct {
	i    int
	page DeletedVaultListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *DeletedVaultListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeletedVaultListResultIterator.NextWithContext")
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
func (iter *DeletedVaultListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter DeletedVaultListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter DeletedVaultListResultIterator) Response() DeletedVaultListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter DeletedVaultListResultIterator) Value() DeletedVault {
	if !iter.page.NotDone() {
		return DeletedVault{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the DeletedVaultListResultIterator type.
func NewDeletedVaultListResultIterator(page DeletedVaultListResultPage) DeletedVaultListResultIterator {
	return DeletedVaultListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (dvlr DeletedVaultListResult) IsEmpty() bool {
	return dvlr.Value == nil || len(*dvlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (dvlr DeletedVaultListResult) hasNextLink() bool {
	return dvlr.NextLink != nil && len(*dvlr.NextLink) != 0
}

// deletedVaultListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (dvlr DeletedVaultListResult) deletedVaultListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !dvlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(dvlr.NextLink)))
}

// DeletedVaultListResultPage contains a page of DeletedVault values.
type DeletedVaultListResultPage struct {
	fn   func(context.Context, DeletedVaultListResult) (DeletedVaultListResult, error)
	dvlr DeletedVaultListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *DeletedVaultListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeletedVaultListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.dvlr)
		if err != nil {
			return err
		}
		page.dvlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *DeletedVaultListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page DeletedVaultListResultPage) NotDone() bool {
	return !page.dvlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page DeletedVaultListResultPage) Response() DeletedVaultListResult {
	return page.dvlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page DeletedVaultListResultPage) Values() []DeletedVault {
	if page.dvlr.IsEmpty() {
		return nil
	}
	return *page.dvlr.Value
}

// Creates a new instance of the DeletedVaultListResultPage type.
func NewDeletedVaultListResultPage(cur DeletedVaultListResult, getNextPage func(context.Context, DeletedVaultListResult) (DeletedVaultListResult, error)) DeletedVaultListResultPage {
	return DeletedVaultListResultPage{
		fn:   getNextPage,
		dvlr: cur,
	}
}

// DeletedVaultProperties properties of the deleted vault.
type DeletedVaultProperties struct {
	// VaultID - READ-ONLY; The resource id of the original vault.
	VaultID *string `json:"vaultId,omitempty"`
	// Location - READ-ONLY; The location of the original vault.
	Location *string `json:"location,omitempty"`
	// DeletionDate - READ-ONLY; The deleted date.
	DeletionDate *date.Time `json:"deletionDate,omitempty"`
	// ScheduledPurgeDate - READ-ONLY; The scheduled purged date.
	ScheduledPurgeDate *date.Time `json:"scheduledPurgeDate,omitempty"`
	// Tags - READ-ONLY; Tags of the original vault.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for DeletedVaultProperties.
func (dvp DeletedVaultProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// LogSpecification log specification of operation.
type LogSpecification struct {
	// Name - Name of log specification.
	Name *string `json:"name,omitempty"`
	// DisplayName - Display name of log specification.
	DisplayName *string `json:"displayName,omitempty"`
	// BlobDuration - Blob duration of specification.
	BlobDuration *string `json:"blobDuration,omitempty"`
}

// Operation key Vault REST API operation definition.
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - Display metadata associated with the operation.
	Display *OperationDisplay `json:"display,omitempty"`
	// Origin - The origin of operations.
	Origin *string `json:"origin,omitempty"`
	// OperationProperties - Properties of operation, include metric specifications.
	*OperationProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if o.Name != nil {
		objectMap["name"] = o.Name
	}
	if o.Display != nil {
		objectMap["display"] = o.Display
	}
	if o.Origin != nil {
		objectMap["origin"] = o.Origin
	}
	if o.OperationProperties != nil {
		objectMap["properties"] = o.OperationProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Operation struct.
func (o *Operation) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				o.Name = &name
			}
		case "display":
			if v != nil {
				var display OperationDisplay
				err = json.Unmarshal(*v, &display)
				if err != nil {
					return err
				}
				o.Display = &display
			}
		case "origin":
			if v != nil {
				var origin string
				err = json.Unmarshal(*v, &origin)
				if err != nil {
					return err
				}
				o.Origin = &origin
			}
		case "properties":
			if v != nil {
				var operationProperties OperationProperties
				err = json.Unmarshal(*v, &operationProperties)
				if err != nil {
					return err
				}
				o.OperationProperties = &operationProperties
			}
		}
	}

	return nil
}

// OperationDisplay display metadata associated with the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft Key Vault.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Type of operation: get, read, delete, etc.
	Operation *string `json:"operation,omitempty"`
	// Description - Description of operation.
	Description *string `json:"description,omitempty"`
}

// OperationListResult result of the request to list Storage operations. It contains a list of operations
// and a URL link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of Storage operations supported by the Storage resource provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - The URL to get the next set of operations.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListResultIterator provides access to a complete listing of Operation values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultIterator.NextWithContext")
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
func (iter *OperationListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationListResultIterator type.
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return OperationListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (olr OperationListResult) hasNextLink() bool {
	return olr.NextLink != nil && len(*olr.NextLink) != 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !olr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of Operation values.
type OperationListResultPage struct {
	fn  func(context.Context, OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.olr)
		if err != nil {
			return err
		}
		page.olr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []Operation {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// Creates a new instance of the OperationListResultPage type.
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return OperationListResultPage{
		fn:  getNextPage,
		olr: cur,
	}
}

// OperationProperties properties of operation, include metric specifications.
type OperationProperties struct {
	// ServiceSpecification - One property of operation, include metric specifications.
	ServiceSpecification *ServiceSpecification `json:"serviceSpecification,omitempty"`
}

// Permissions permissions the identity has for keys, secrets, certificates and storage.
type Permissions struct {
	// Keys - Permissions to keys
	Keys *[]KeyPermissions `json:"keys,omitempty"`
	// Secrets - Permissions to secrets
	Secrets *[]SecretPermissions `json:"secrets,omitempty"`
	// Certificates - Permissions to certificates
	Certificates *[]CertificatePermissions `json:"certificates,omitempty"`
	// Storage - Permissions to storage accounts
	Storage *[]StoragePermissions `json:"storage,omitempty"`
}

// Resource key Vault resource
type Resource struct {
	// ID - READ-ONLY; The Azure Resource Manager resource ID for the key vault.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the key vault.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The resource type of the key vault.
	Type *string `json:"type,omitempty"`
	// Location - The supported Azure location where the key vault should be created.
	Location *string `json:"location,omitempty"`
	// Tags - The tags that will be assigned to the key vault.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// ResourceListResult list of vault resources.
type ResourceListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of vault resources.
	Value *[]Resource `json:"value,omitempty"`
	// NextLink - The URL to get the next set of vault resources.
	NextLink *string `json:"nextLink,omitempty"`
}

// ResourceListResultIterator provides access to a complete listing of Resource values.
type ResourceListResultIterator struct {
	i    int
	page ResourceListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ResourceListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceListResultIterator.NextWithContext")
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
func (iter *ResourceListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ResourceListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ResourceListResultIterator) Response() ResourceListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ResourceListResultIterator) Value() Resource {
	if !iter.page.NotDone() {
		return Resource{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ResourceListResultIterator type.
func NewResourceListResultIterator(page ResourceListResultPage) ResourceListResultIterator {
	return ResourceListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (rlr ResourceListResult) IsEmpty() bool {
	return rlr.Value == nil || len(*rlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (rlr ResourceListResult) hasNextLink() bool {
	return rlr.NextLink != nil && len(*rlr.NextLink) != 0
}

// resourceListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (rlr ResourceListResult) resourceListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !rlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(rlr.NextLink)))
}

// ResourceListResultPage contains a page of Resource values.
type ResourceListResultPage struct {
	fn  func(context.Context, ResourceListResult) (ResourceListResult, error)
	rlr ResourceListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ResourceListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.rlr)
		if err != nil {
			return err
		}
		page.rlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ResourceListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ResourceListResultPage) NotDone() bool {
	return !page.rlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ResourceListResultPage) Response() ResourceListResult {
	return page.rlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ResourceListResultPage) Values() []Resource {
	if page.rlr.IsEmpty() {
		return nil
	}
	return *page.rlr.Value
}

// Creates a new instance of the ResourceListResultPage type.
func NewResourceListResultPage(cur ResourceListResult, getNextPage func(context.Context, ResourceListResult) (ResourceListResult, error)) ResourceListResultPage {
	return ResourceListResultPage{
		fn:  getNextPage,
		rlr: cur,
	}
}

// ServiceSpecification one property of operation, include log specifications.
type ServiceSpecification struct {
	// LogSpecifications - Log specifications of operation.
	LogSpecifications *[]LogSpecification `json:"logSpecifications,omitempty"`
}

// Sku SKU details
type Sku struct {
	// Family - SKU family name
	Family *string `json:"family,omitempty"`
	// Name - SKU name to specify whether the key vault is a standard vault or a premium vault. Possible values include: 'Standard', 'Premium'
	Name SkuName `json:"name,omitempty"`
}

// Vault resource information with extended details.
type Vault struct {
	autorest.Response `json:"-"`
	// Properties - Properties of the vault
	Properties *VaultProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; The Azure Resource Manager resource ID for the key vault.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the key vault.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The resource type of the key vault.
	Type *string `json:"type,omitempty"`
	// Location - The supported Azure location where the key vault should be created.
	Location *string `json:"location,omitempty"`
	// Tags - The tags that will be assigned to the key vault.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Vault.
func (vVar Vault) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if vVar.Properties != nil {
		objectMap["properties"] = vVar.Properties
	}
	if vVar.Location != nil {
		objectMap["location"] = vVar.Location
	}
	if vVar.Tags != nil {
		objectMap["tags"] = vVar.Tags
	}
	return json.Marshal(objectMap)
}

// VaultAccessPolicyParameters parameters for updating the access policy in a vault
type VaultAccessPolicyParameters struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; The resource id of the access policy.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The resource name of the access policy.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The resource name of the access policy.
	Type *string `json:"type,omitempty"`
	// Location - READ-ONLY; The resource type of the access policy.
	Location *string `json:"location,omitempty"`
	// Properties - Properties of the access policy
	Properties *VaultAccessPolicyProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for VaultAccessPolicyParameters.
func (vapp VaultAccessPolicyParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if vapp.Properties != nil {
		objectMap["properties"] = vapp.Properties
	}
	return json.Marshal(objectMap)
}

// VaultAccessPolicyProperties properties of the vault access policy
type VaultAccessPolicyProperties struct {
	// AccessPolicies - An array of 0 to 16 identities that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant ID.
	AccessPolicies *[]AccessPolicyEntry `json:"accessPolicies,omitempty"`
}

// VaultCheckNameAvailabilityParameters the parameters used to check the availability of the vault name.
type VaultCheckNameAvailabilityParameters struct {
	// Name - The vault name.
	Name *string `json:"name,omitempty"`
	// Type - The type of resource, Microsoft.KeyVault/vaults
	Type *string `json:"type,omitempty"`
}

// VaultCreateOrUpdateParameters parameters for creating or updating a vault
type VaultCreateOrUpdateParameters struct {
	// Location - The supported Azure location where the key vault should be created.
	Location *string `json:"location,omitempty"`
	// Tags - The tags that will be assigned to the key vault.
	Tags map[string]*string `json:"tags"`
	// Properties - Properties of the vault
	Properties *VaultProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for VaultCreateOrUpdateParameters.
func (vcoup VaultCreateOrUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if vcoup.Location != nil {
		objectMap["location"] = vcoup.Location
	}
	if vcoup.Tags != nil {
		objectMap["tags"] = vcoup.Tags
	}
	if vcoup.Properties != nil {
		objectMap["properties"] = vcoup.Properties
	}
	return json.Marshal(objectMap)
}

// VaultListResult list of vaults
type VaultListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of vaults.
	Value *[]Vault `json:"value,omitempty"`
	// NextLink - The URL to get the next set of vaults.
	NextLink *string `json:"nextLink,omitempty"`
}

// VaultListResultIterator provides access to a complete listing of Vault values.
type VaultListResultIterator struct {
	i    int
	page VaultListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *VaultListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VaultListResultIterator.NextWithContext")
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
func (iter *VaultListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter VaultListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter VaultListResultIterator) Response() VaultListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter VaultListResultIterator) Value() Vault {
	if !iter.page.NotDone() {
		return Vault{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the VaultListResultIterator type.
func NewVaultListResultIterator(page VaultListResultPage) VaultListResultIterator {
	return VaultListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (vlr VaultListResult) IsEmpty() bool {
	return vlr.Value == nil || len(*vlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (vlr VaultListResult) hasNextLink() bool {
	return vlr.NextLink != nil && len(*vlr.NextLink) != 0
}

// vaultListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (vlr VaultListResult) vaultListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !vlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(vlr.NextLink)))
}

// VaultListResultPage contains a page of Vault values.
type VaultListResultPage struct {
	fn  func(context.Context, VaultListResult) (VaultListResult, error)
	vlr VaultListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *VaultListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VaultListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.vlr)
		if err != nil {
			return err
		}
		page.vlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *VaultListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page VaultListResultPage) NotDone() bool {
	return !page.vlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page VaultListResultPage) Response() VaultListResult {
	return page.vlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page VaultListResultPage) Values() []Vault {
	if page.vlr.IsEmpty() {
		return nil
	}
	return *page.vlr.Value
}

// Creates a new instance of the VaultListResultPage type.
func NewVaultListResultPage(cur VaultListResult, getNextPage func(context.Context, VaultListResult) (VaultListResult, error)) VaultListResultPage {
	return VaultListResultPage{
		fn:  getNextPage,
		vlr: cur,
	}
}

// VaultPatchParameters parameters for creating or updating a vault
type VaultPatchParameters struct {
	// Tags - The tags that will be assigned to the key vault.
	Tags map[string]*string `json:"tags"`
	// Properties - Properties of the vault
	Properties *VaultPatchProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for VaultPatchParameters.
func (vpp VaultPatchParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if vpp.Tags != nil {
		objectMap["tags"] = vpp.Tags
	}
	if vpp.Properties != nil {
		objectMap["properties"] = vpp.Properties
	}
	return json.Marshal(objectMap)
}

// VaultPatchProperties properties of the vault
type VaultPatchProperties struct {
	// TenantID - The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantID *uuid.UUID `json:"tenantId,omitempty"`
	// Sku - SKU details
	Sku *Sku `json:"sku,omitempty"`
	// AccessPolicies - An array of 0 to 16 identities that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant ID.
	AccessPolicies *[]AccessPolicyEntry `json:"accessPolicies,omitempty"`
	// EnabledForDeployment - Property to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault.
	EnabledForDeployment *bool `json:"enabledForDeployment,omitempty"`
	// EnabledForDiskEncryption - Property to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys.
	EnabledForDiskEncryption *bool `json:"enabledForDiskEncryption,omitempty"`
	// EnabledForTemplateDeployment - Property to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault.
	EnabledForTemplateDeployment *bool `json:"enabledForTemplateDeployment,omitempty"`
	// EnableSoftDelete - Property specifying whether recoverable deletion ('soft' delete) is enabled for this key vault. The property may not be set to false.
	EnableSoftDelete *bool `json:"enableSoftDelete,omitempty"`
	// CreateMode - The vault's create mode to indicate whether the vault need to be recovered or not. Possible values include: 'CreateModeRecover', 'CreateModeDefault'
	CreateMode CreateMode `json:"createMode,omitempty"`
	// EnablePurgeProtection - Property specifying whether protection against purge is enabled for this vault; it is only effective if soft delete is also enabled. Once activated, the property may no longer be reset to false.
	EnablePurgeProtection *bool `json:"enablePurgeProtection,omitempty"`
}

// VaultProperties properties of the vault
type VaultProperties struct {
	// TenantID - The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.
	TenantID *uuid.UUID `json:"tenantId,omitempty"`
	// Sku - SKU details
	Sku *Sku `json:"sku,omitempty"`
	// AccessPolicies - An array of 0 to 16 identities that have access to the key vault. All identities in the array must use the same tenant ID as the key vault's tenant ID. When `createMode` is set to `recover`, access policies are not required. Otherwise, access policies are required.
	AccessPolicies *[]AccessPolicyEntry `json:"accessPolicies,omitempty"`
	// VaultURI - The URI of the vault for performing operations on keys and secrets.
	VaultURI *string `json:"vaultUri,omitempty"`
	// EnabledForDeployment - Property to specify whether Azure Virtual Machines are permitted to retrieve certificates stored as secrets from the key vault.
	EnabledForDeployment *bool `json:"enabledForDeployment,omitempty"`
	// EnabledForDiskEncryption - Property to specify whether Azure Disk Encryption is permitted to retrieve secrets from the vault and unwrap keys.
	EnabledForDiskEncryption *bool `json:"enabledForDiskEncryption,omitempty"`
	// EnabledForTemplateDeployment - Property to specify whether Azure Resource Manager is permitted to retrieve secrets from the key vault.
	EnabledForTemplateDeployment *bool `json:"enabledForTemplateDeployment,omitempty"`
	// EnableSoftDelete - Property specifying whether recoverable deletion is enabled for this key vault. Setting this property to true activates the soft delete feature, whereby vaults or vault entities can be recovered after deletion. Enabling this functionality is irreversible - that is, the property does not accept false as its value.
	EnableSoftDelete *bool `json:"enableSoftDelete,omitempty"`
	// CreateMode - The vault's create mode to indicate whether the vault need to be recovered or not. Possible values include: 'CreateModeRecover', 'CreateModeDefault'
	CreateMode CreateMode `json:"createMode,omitempty"`
	// EnablePurgeProtection - Property specifying whether protection against purge is enabled for this vault. Setting this property to true activates protection against purge for this vault and its content - only the Key Vault service may initiate a hard, irrecoverable deletion. The setting is effective only if soft delete is also enabled. Enabling this functionality is irreversible - that is, the property does not accept false as its value.
	EnablePurgeProtection *bool `json:"enablePurgeProtection,omitempty"`
}

// VaultsPurgeDeletedFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type VaultsPurgeDeletedFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *VaultsPurgeDeletedFuture) Result(client VaultsClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.VaultsPurgeDeletedFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("keyvault.VaultsPurgeDeletedFuture")
		return
	}
	ar.Response = future.Response()
	return
}
