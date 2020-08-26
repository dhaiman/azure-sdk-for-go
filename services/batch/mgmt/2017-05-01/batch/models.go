package batch

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
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/batch/mgmt/2017-05-01/batch"

// AccountKeyType enumerates the values for account key type.
type AccountKeyType string

const (
	// Primary ...
	Primary AccountKeyType = "Primary"
	// Secondary ...
	Secondary AccountKeyType = "Secondary"
)

// PossibleAccountKeyTypeValues returns an array of possible values for the AccountKeyType const type.
func PossibleAccountKeyTypeValues() []AccountKeyType {
	return []AccountKeyType{Primary, Secondary}
}

// NameAvailabilityReason enumerates the values for name availability reason.
type NameAvailabilityReason string

const (
	// AlreadyExists ...
	AlreadyExists NameAvailabilityReason = "AlreadyExists"
	// Invalid ...
	Invalid NameAvailabilityReason = "Invalid"
)

// PossibleNameAvailabilityReasonValues returns an array of possible values for the NameAvailabilityReason const type.
func PossibleNameAvailabilityReasonValues() []NameAvailabilityReason {
	return []NameAvailabilityReason{AlreadyExists, Invalid}
}

// PackageState enumerates the values for package state.
type PackageState string

const (
	// Active ...
	Active PackageState = "active"
	// Pending ...
	Pending PackageState = "pending"
	// Unmapped ...
	Unmapped PackageState = "unmapped"
)

// PossiblePackageStateValues returns an array of possible values for the PackageState const type.
func PossiblePackageStateValues() []PackageState {
	return []PackageState{Active, Pending, Unmapped}
}

// PoolAllocationMode enumerates the values for pool allocation mode.
type PoolAllocationMode string

const (
	// BatchService ...
	BatchService PoolAllocationMode = "BatchService"
	// UserSubscription ...
	UserSubscription PoolAllocationMode = "UserSubscription"
)

// PossiblePoolAllocationModeValues returns an array of possible values for the PoolAllocationMode const type.
func PossiblePoolAllocationModeValues() []PoolAllocationMode {
	return []PoolAllocationMode{BatchService, UserSubscription}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateCancelled ...
	ProvisioningStateCancelled ProvisioningState = "Cancelled"
	// ProvisioningStateCreating ...
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleting ...
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed ...
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateInvalid ...
	ProvisioningStateInvalid ProvisioningState = "Invalid"
	// ProvisioningStateSucceeded ...
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{ProvisioningStateCancelled, ProvisioningStateCreating, ProvisioningStateDeleting, ProvisioningStateFailed, ProvisioningStateInvalid, ProvisioningStateSucceeded}
}

// Account contains information about an Azure Batch account.
type Account struct {
	autorest.Response `json:"-"`
	// AccountProperties - The properties associated with the account.
	*AccountProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; The ID of the resource.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
	// Location - READ-ONLY; The location of the resource.
	Location *string `json:"location,omitempty"`
	// Tags - READ-ONLY; The tags of the resource.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Account.
func (a Account) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if a.AccountProperties != nil {
		objectMap["properties"] = a.AccountProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Account struct.
func (a *Account) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var accountProperties AccountProperties
				err = json.Unmarshal(*v, &accountProperties)
				if err != nil {
					return err
				}
				a.AccountProperties = &accountProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				a.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				a.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				a.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				a.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				a.Tags = tags
			}
		}
	}

	return nil
}

// AccountCreateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type AccountCreateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *AccountCreateFuture) Result(client AccountClient) (a Account, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountCreateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("batch.AccountCreateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if a.Response.Response, err = future.GetResult(sender); err == nil && a.Response.Response.StatusCode != http.StatusNoContent {
		a, err = client.CreateResponder(a.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "batch.AccountCreateFuture", "Result", a.Response.Response, "Failure responding to request")
		}
	}
	return
}

// AccountCreateParameters parameters supplied to the Create operation.
type AccountCreateParameters struct {
	// Location - The region in which to create the account.
	Location *string `json:"location,omitempty"`
	// Tags - The user-specified tags associated with the account.
	Tags map[string]*string `json:"tags"`
	// AccountCreateProperties - The properties of the Batch account.
	*AccountCreateProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountCreateParameters.
func (acp AccountCreateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if acp.Location != nil {
		objectMap["location"] = acp.Location
	}
	if acp.Tags != nil {
		objectMap["tags"] = acp.Tags
	}
	if acp.AccountCreateProperties != nil {
		objectMap["properties"] = acp.AccountCreateProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for AccountCreateParameters struct.
func (acp *AccountCreateParameters) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				acp.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				acp.Tags = tags
			}
		case "properties":
			if v != nil {
				var accountCreateProperties AccountCreateProperties
				err = json.Unmarshal(*v, &accountCreateProperties)
				if err != nil {
					return err
				}
				acp.AccountCreateProperties = &accountCreateProperties
			}
		}
	}

	return nil
}

// AccountCreateProperties the properties of a Batch account.
type AccountCreateProperties struct {
	// AutoStorage - The properties related to the auto-storage account.
	AutoStorage *AutoStorageBaseProperties `json:"autoStorage,omitempty"`
	// PoolAllocationMode - The pool allocation mode also affects how clients may authenticate to the Batch Service API. If the mode is BatchService, clients may authenticate using access keys or Azure Active Directory. If the mode is UserSubscription, clients must use Azure Active Directory. The default is BatchService. Possible values include: 'BatchService', 'UserSubscription'
	PoolAllocationMode PoolAllocationMode `json:"poolAllocationMode,omitempty"`
	// KeyVaultReference - A reference to the Azure key vault associated with the Batch account.
	KeyVaultReference *KeyVaultReference `json:"keyVaultReference,omitempty"`
}

// AccountDeleteFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type AccountDeleteFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *AccountDeleteFuture) Result(client AccountClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.AccountDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("batch.AccountDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// AccountKeys a set of Azure Batch account keys.
type AccountKeys struct {
	autorest.Response `json:"-"`
	// AccountName - READ-ONLY; The Batch account name.
	AccountName *string `json:"accountName,omitempty"`
	// Primary - READ-ONLY; The primary key associated with the account.
	Primary *string `json:"primary,omitempty"`
	// Secondary - READ-ONLY; The secondary key associated with the account.
	Secondary *string `json:"secondary,omitempty"`
}

// AccountListResult values returned by the List operation.
type AccountListResult struct {
	autorest.Response `json:"-"`
	// Value - The collection of Batch accounts returned by the listing operation.
	Value *[]Account `json:"value,omitempty"`
	// NextLink - The continuation token.
	NextLink *string `json:"nextLink,omitempty"`
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

// AccountProperties account specific properties.
type AccountProperties struct {
	// AccountEndpoint - READ-ONLY; The account endpoint used to interact with the Batch service.
	AccountEndpoint *string `json:"accountEndpoint,omitempty"`
	// ProvisioningState - READ-ONLY; The provisioned state of the resource. Possible values include: 'ProvisioningStateInvalid', 'ProvisioningStateCreating', 'ProvisioningStateDeleting', 'ProvisioningStateSucceeded', 'ProvisioningStateFailed', 'ProvisioningStateCancelled'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// PoolAllocationMode - READ-ONLY; Possible values include: 'BatchService', 'UserSubscription'
	PoolAllocationMode PoolAllocationMode `json:"poolAllocationMode,omitempty"`
	// KeyVaultReference - READ-ONLY
	KeyVaultReference *KeyVaultReference `json:"keyVaultReference,omitempty"`
	// AutoStorage - READ-ONLY
	AutoStorage *AutoStorageProperties `json:"autoStorage,omitempty"`
	// DedicatedCoreQuota - READ-ONLY
	DedicatedCoreQuota *int32 `json:"dedicatedCoreQuota,omitempty"`
	// LowPriorityCoreQuota - READ-ONLY
	LowPriorityCoreQuota *int32 `json:"lowPriorityCoreQuota,omitempty"`
	// PoolQuota - READ-ONLY
	PoolQuota *int32 `json:"poolQuota,omitempty"`
	// ActiveJobAndJobScheduleQuota - READ-ONLY
	ActiveJobAndJobScheduleQuota *int32 `json:"activeJobAndJobScheduleQuota,omitempty"`
}

// AccountRegenerateKeyParameters parameters supplied to the RegenerateKey operation.
type AccountRegenerateKeyParameters struct {
	// KeyName - The type of account key to regenerate. Possible values include: 'Primary', 'Secondary'
	KeyName AccountKeyType `json:"keyName,omitempty"`
}

// AccountUpdateParameters parameters for updating an Azure Batch account.
type AccountUpdateParameters struct {
	// Tags - The user-specified tags associated with the account.
	Tags map[string]*string `json:"tags"`
	// AccountUpdateProperties - The properties of the account.
	*AccountUpdateProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountUpdateParameters.
func (aup AccountUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if aup.Tags != nil {
		objectMap["tags"] = aup.Tags
	}
	if aup.AccountUpdateProperties != nil {
		objectMap["properties"] = aup.AccountUpdateProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for AccountUpdateParameters struct.
func (aup *AccountUpdateParameters) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				aup.Tags = tags
			}
		case "properties":
			if v != nil {
				var accountUpdateProperties AccountUpdateProperties
				err = json.Unmarshal(*v, &accountUpdateProperties)
				if err != nil {
					return err
				}
				aup.AccountUpdateProperties = &accountUpdateProperties
			}
		}
	}

	return nil
}

// AccountUpdateProperties the properties of a Batch account.
type AccountUpdateProperties struct {
	// AutoStorage - The properties related to the auto-storage account.
	AutoStorage *AutoStorageBaseProperties `json:"autoStorage,omitempty"`
}

// ActivateApplicationPackageParameters parameters for an activating an application package.
type ActivateApplicationPackageParameters struct {
	// Format - The format of the application package binary file.
	Format *string `json:"format,omitempty"`
}

// Application contains information about an application in a Batch account.
type Application struct {
	autorest.Response `json:"-"`
	// ID - A string that uniquely identifies the application within the account.
	ID *string `json:"id,omitempty"`
	// DisplayName - The display name for the application.
	DisplayName *string `json:"displayName,omitempty"`
	// Packages - The list of packages under this application.
	Packages *[]ApplicationPackage `json:"packages,omitempty"`
	// AllowUpdates - A value indicating whether packages within the application may be overwritten using the same version string.
	AllowUpdates *bool `json:"allowUpdates,omitempty"`
	// DefaultVersion - The package to use if a client requests the application but does not specify a version.
	DefaultVersion *string `json:"defaultVersion,omitempty"`
}

// ApplicationCreateParameters parameters for adding an Application.
type ApplicationCreateParameters struct {
	// AllowUpdates - A value indicating whether packages within the application may be overwritten using the same version string.
	AllowUpdates *bool `json:"allowUpdates,omitempty"`
	// DisplayName - The display name for the application.
	DisplayName *string `json:"displayName,omitempty"`
}

// ApplicationPackage an application package which represents a particular version of an application.
type ApplicationPackage struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; The ID of the application.
	ID *string `json:"id,omitempty"`
	// Version - READ-ONLY; The version of the application package.
	Version *string `json:"version,omitempty"`
	// State - READ-ONLY; The current state of the application package. Possible values include: 'Pending', 'Active', 'Unmapped'
	State PackageState `json:"state,omitempty"`
	// Format - READ-ONLY; The format of the application package, if the package is active.
	Format *string `json:"format,omitempty"`
	// StorageURL - READ-ONLY; The URL for the application package in Azure Storage.
	StorageURL *string `json:"storageUrl,omitempty"`
	// StorageURLExpiry - READ-ONLY; The UTC time at which the Azure Storage URL will expire.
	StorageURLExpiry *date.Time `json:"storageUrlExpiry,omitempty"`
	// LastActivationTime - READ-ONLY; The time at which the package was last activated, if the package is active.
	LastActivationTime *date.Time `json:"lastActivationTime,omitempty"`
}

// ApplicationUpdateParameters parameters for an update application request.
type ApplicationUpdateParameters struct {
	// AllowUpdates - A value indicating whether packages within the application may be overwritten using the same version string.
	AllowUpdates *bool `json:"allowUpdates,omitempty"`
	// DefaultVersion - The package to use if a client requests the application but does not specify a version.
	DefaultVersion *string `json:"defaultVersion,omitempty"`
	// DisplayName - The display name for the application.
	DisplayName *string `json:"displayName,omitempty"`
}

// AutoStorageBaseProperties the properties related to the auto-storage account.
type AutoStorageBaseProperties struct {
	// StorageAccountID - The resource ID of the storage account to be used for auto-storage account.
	StorageAccountID *string `json:"storageAccountId,omitempty"`
}

// AutoStorageProperties contains information about the auto-storage account associated with a Batch
// account.
type AutoStorageProperties struct {
	// LastKeySync - The UTC time at which storage keys were last synchronized with the Batch account.
	LastKeySync *date.Time `json:"lastKeySync,omitempty"`
	// StorageAccountID - The resource ID of the storage account to be used for auto-storage account.
	StorageAccountID *string `json:"storageAccountId,omitempty"`
}

// CheckNameAvailabilityParameters parameters for a check name availability request.
type CheckNameAvailabilityParameters struct {
	// Name - The name to check for availability
	Name *string `json:"name,omitempty"`
	// Type - The resource type. Must be set to Microsoft.Batch/batchAccounts
	Type *string `json:"type,omitempty"`
}

// CheckNameAvailabilityResult the CheckNameAvailability operation response.
type CheckNameAvailabilityResult struct {
	autorest.Response `json:"-"`
	// NameAvailable - READ-ONLY; Gets a boolean value that indicates whether the name is available for you to use. If true, the name is available. If false, the name has already been taken or invalid and cannot be used.
	NameAvailable *bool `json:"nameAvailable,omitempty"`
	// Reason - READ-ONLY; Gets the reason that a Batch account name could not be used. The Reason element is only returned if NameAvailable is false. Possible values include: 'Invalid', 'AlreadyExists'
	Reason NameAvailabilityReason `json:"reason,omitempty"`
	// Message - READ-ONLY; Gets an error message explaining the Reason value in more detail.
	Message *string `json:"message,omitempty"`
}

// CloudError an error response from the Batch service.
type CloudError struct {
	Error *CloudErrorBody `json:"error,omitempty"`
}

// CloudErrorBody an error response from the Batch service.
type CloudErrorBody struct {
	// Code - An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
	Code *string `json:"code,omitempty"`
	// Message - A message describing the error, intended to be suitable for display in a user interface.
	Message *string `json:"message,omitempty"`
	// Target - The target of the particular error. For example, the name of the property in error.
	Target *string `json:"target,omitempty"`
	// Details - A list of additional details about the error.
	Details *[]CloudErrorBody `json:"details,omitempty"`
}

// KeyVaultReference identifies the Azure key vault associated with a Batch account.
type KeyVaultReference struct {
	// ID - The resource ID of the Azure key vault associated with the Batch account.
	ID *string `json:"id,omitempty"`
	// URL - The URL of the Azure key vault associated with the Batch account.
	URL *string `json:"url,omitempty"`
}

// ListApplicationsResult the result of performing list applications.
type ListApplicationsResult struct {
	autorest.Response `json:"-"`
	// Value - The list of applications.
	Value *[]Application `json:"value,omitempty"`
	// NextLink - The URL to get the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// ListApplicationsResultIterator provides access to a complete listing of Application values.
type ListApplicationsResultIterator struct {
	i    int
	page ListApplicationsResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ListApplicationsResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListApplicationsResultIterator.NextWithContext")
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
func (iter *ListApplicationsResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ListApplicationsResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ListApplicationsResultIterator) Response() ListApplicationsResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ListApplicationsResultIterator) Value() Application {
	if !iter.page.NotDone() {
		return Application{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ListApplicationsResultIterator type.
func NewListApplicationsResultIterator(page ListApplicationsResultPage) ListApplicationsResultIterator {
	return ListApplicationsResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (lar ListApplicationsResult) IsEmpty() bool {
	return lar.Value == nil || len(*lar.Value) == 0
}

// listApplicationsResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (lar ListApplicationsResult) listApplicationsResultPreparer(ctx context.Context) (*http.Request, error) {
	if lar.NextLink == nil || len(to.String(lar.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(lar.NextLink)))
}

// ListApplicationsResultPage contains a page of Application values.
type ListApplicationsResultPage struct {
	fn  func(context.Context, ListApplicationsResult) (ListApplicationsResult, error)
	lar ListApplicationsResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ListApplicationsResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListApplicationsResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.lar)
	if err != nil {
		return err
	}
	page.lar = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ListApplicationsResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ListApplicationsResultPage) NotDone() bool {
	return !page.lar.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ListApplicationsResultPage) Response() ListApplicationsResult {
	return page.lar
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ListApplicationsResultPage) Values() []Application {
	if page.lar.IsEmpty() {
		return nil
	}
	return *page.lar.Value
}

// Creates a new instance of the ListApplicationsResultPage type.
func NewListApplicationsResultPage(getNextPage func(context.Context, ListApplicationsResult) (ListApplicationsResult, error)) ListApplicationsResultPage {
	return ListApplicationsResultPage{fn: getNextPage}
}

// LocationQuota quotas associated with a Batch region for a particular subscription.
type LocationQuota struct {
	autorest.Response `json:"-"`
	// AccountQuota - READ-ONLY; The number of Batch accounts that may be created under the subscription in the specified region.
	AccountQuota *int32 `json:"accountQuota,omitempty"`
}

// Operation ...
type Operation struct {
	// Name - This is of the format {provider}/{resource}/{operation}
	Name       *string           `json:"name,omitempty"`
	Display    *OperationDisplay `json:"display,omitempty"`
	Origin     *string           `json:"origin,omitempty"`
	Properties interface{}       `json:"properties,omitempty"`
}

// OperationDisplay ...
type OperationDisplay struct {
	Provider *string `json:"provider,omitempty"`
	// Operation - For example: read, write, delete, or listKeys/action
	Operation   *string `json:"operation,omitempty"`
	Resource    *string `json:"resource,omitempty"`
	Description *string `json:"description,omitempty"`
}

// OperationListResult ...
type OperationListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Operation `json:"value,omitempty"`
	NextLink          *string      `json:"nextLink,omitempty"`
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

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if olr.NextLink == nil || len(to.String(olr.NextLink)) < 1 {
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
	next, err := page.fn(ctx, page.olr)
	if err != nil {
		return err
	}
	page.olr = next
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
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return OperationListResultPage{fn: getNextPage}
}

// Resource a definition of an Azure resource.
type Resource struct {
	// ID - READ-ONLY; The ID of the resource.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty"`
	// Location - READ-ONLY; The location of the resource.
	Location *string `json:"location,omitempty"`
	// Tags - READ-ONLY; The tags of the resource.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}
