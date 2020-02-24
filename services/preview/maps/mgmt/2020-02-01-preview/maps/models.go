package maps

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
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/maps/mgmt/2020-02-01-preview/maps"

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// Application ...
	Application CreatedByType = "Application"
	// Key ...
	Key CreatedByType = "Key"
	// ManagedIdentity ...
	ManagedIdentity CreatedByType = "ManagedIdentity"
	// User ...
	User CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{Application, Key, ManagedIdentity, User}
}

// KeyType enumerates the values for key type.
type KeyType string

const (
	// Primary ...
	Primary KeyType = "primary"
	// Secondary ...
	Secondary KeyType = "secondary"
)

// PossibleKeyTypeValues returns an array of possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{Primary, Secondary}
}

// Account an Azure resource which represents access to a suite of Maps REST APIs.
type Account struct {
	autorest.Response `json:"-"`
	// Sku - READ-ONLY; The SKU of this account.
	Sku *Sku `json:"sku,omitempty"`
	// SystemData - READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData `json:"systemData,omitempty"`
	// Properties - READ-ONLY; The map account properties.
	Properties *AccountProperties `json:"properties,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Account.
func (a Account) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if a.Tags != nil {
		objectMap["tags"] = a.Tags
	}
	if a.Location != nil {
		objectMap["location"] = a.Location
	}
	return json.Marshal(objectMap)
}

// AccountCreateParameters parameters used to create a new Maps Account.
type AccountCreateParameters struct {
	// Location - The location of the resource.
	Location *string `json:"location,omitempty"`
	// Tags - Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]*string `json:"tags"`
	// Sku - The SKU of this account.
	Sku *Sku `json:"sku,omitempty"`
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
	if acp.Sku != nil {
		objectMap["sku"] = acp.Sku
	}
	return json.Marshal(objectMap)
}

// AccountKeys the set of keys which can be used to access the Maps REST APIs. Two keys are provided for
// key rotation without interruption.
type AccountKeys struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; The full Azure resource identifier of the Maps Account.
	ID *string `json:"id,omitempty"`
	// PrimaryKey - READ-ONLY; The primary key for accessing the Maps REST APIs.
	PrimaryKey *string `json:"primaryKey,omitempty"`
	// SecondaryKey - READ-ONLY; The secondary key for accessing the Maps REST APIs.
	SecondaryKey *string `json:"secondaryKey,omitempty"`
}

// AccountProperties additional Map account properties
type AccountProperties struct {
	// XMsClientID - A unique identifier for the maps account
	XMsClientID *string `json:"x-ms-client-id,omitempty"`
}

// Accounts a list of Maps Accounts.
type Accounts struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; a Maps Account.
	Value *[]Account `json:"value,omitempty"`
}

// AccountUpdateParameters parameters used to update an existing Maps Account.
type AccountUpdateParameters struct {
	// Tags - Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]*string `json:"tags"`
	// Sku - The SKU of this account.
	Sku *Sku `json:"sku,omitempty"`
}

// MarshalJSON is the custom marshaler for AccountUpdateParameters.
func (aup AccountUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if aup.Tags != nil {
		objectMap["tags"] = aup.Tags
	}
	if aup.Sku != nil {
		objectMap["sku"] = aup.Sku
	}
	return json.Marshal(objectMap)
}

// AzureEntityResource the resource model definition for a Azure Resource Manager resource with an etag.
type AzureEntityResource struct {
	// Etag - READ-ONLY; Resource Etag.
	Etag *string `json:"etag,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// ErrorAdditionalInfo the resource management error additional info.
type ErrorAdditionalInfo struct {
	// Type - READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty"`
	// Info - READ-ONLY; The additional info.
	Info interface{} `json:"info,omitempty"`
}

// ErrorResponse the resource management error response.
type ErrorResponse struct {
	// Code - READ-ONLY; The error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; The error message.
	Message *string `json:"message,omitempty"`
	// Target - READ-ONLY; The error target.
	Target *string `json:"target,omitempty"`
	// Details - READ-ONLY; The error details.
	Details *[]ErrorResponse `json:"details,omitempty"`
	// AdditionalInfo - READ-ONLY; The error additional info.
	AdditionalInfo *[]ErrorAdditionalInfo `json:"additionalInfo,omitempty"`
}

// KeySpecification whether the operation refers to the primary or secondary key.
type KeySpecification struct {
	// KeyType - Whether the operation refers to the primary or secondary key. Possible values include: 'Primary', 'Secondary'
	KeyType KeyType `json:"keyType,omitempty"`
}

// Operations the set of operations available for Maps.
type Operations struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; An operation available for Maps.
	Value *[]OperationsValueItem `json:"value,omitempty"`
}

// OperationsValueItem ...
type OperationsValueItem struct {
	// Name - READ-ONLY; Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The human-readable description of the operation.
	Display *OperationsValueItemDisplay `json:"display,omitempty"`
	// Origin - READ-ONLY; The origin of the operation.
	Origin *string `json:"origin,omitempty"`
}

// OperationsValueItemDisplay the human-readable description of the operation.
type OperationsValueItemDisplay struct {
	// Provider - READ-ONLY; Service provider: Microsoft Maps.
	Provider *string `json:"provider,omitempty"`
	// Resource - READ-ONLY; Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
	// Operation - READ-ONLY; The action that users can perform, based on their permission level.
	Operation *string `json:"operation,omitempty"`
	// Description - READ-ONLY; The description of the operation.
	Description *string `json:"description,omitempty"`
}

// PrivateAtlas an Azure resource which represents which will provision the ability to create private
// location data.
type PrivateAtlas struct {
	autorest.Response `json:"-"`
	// Properties - The Private Atlas resource properties.
	Properties *PrivateAtlasProperties `json:"properties,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for PrivateAtlas.
func (pa PrivateAtlas) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if pa.Properties != nil {
		objectMap["properties"] = pa.Properties
	}
	if pa.Tags != nil {
		objectMap["tags"] = pa.Tags
	}
	if pa.Location != nil {
		objectMap["location"] = pa.Location
	}
	return json.Marshal(objectMap)
}

// PrivateAtlasCreateParameters parameters used to create a new Private Atlas resource.
type PrivateAtlasCreateParameters struct {
	// Location - The location of the resource.
	Location *string `json:"location,omitempty"`
	// Tags - Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for PrivateAtlasCreateParameters.
func (pacp PrivateAtlasCreateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if pacp.Location != nil {
		objectMap["location"] = pacp.Location
	}
	if pacp.Tags != nil {
		objectMap["tags"] = pacp.Tags
	}
	return json.Marshal(objectMap)
}

// PrivateAtlasList a list of Private Atlas resources.
type PrivateAtlasList struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; a Private Atlas.
	Value *[]PrivateAtlas `json:"value,omitempty"`
}

// PrivateAtlasProperties private Atlas resource properties
type PrivateAtlasProperties struct {
	// ProvisioningState - The state of the resource provisioning, terminal states: Succeeded, Failed, Canceled
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

// PrivateAtlasUpdateParameters parameters used to update an existing Private Atlas resource.
type PrivateAtlasUpdateParameters struct {
	// Tags - Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for PrivateAtlasUpdateParameters.
func (paup PrivateAtlasUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if paup.Tags != nil {
		objectMap["tags"] = paup.Tags
	}
	return json.Marshal(objectMap)
}

// ProxyResource the resource model definition for a ARM proxy resource. It will have everything other than
// required location and tags
type ProxyResource struct {
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// Resource ...
type Resource struct {
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// Sku the SKU of the Maps Account.
type Sku struct {
	// Name - The name of the SKU, in standard format (such as S0).
	Name *string `json:"name,omitempty"`
	// Tier - READ-ONLY; Gets the sku tier. This is based on the SKU name.
	Tier *string `json:"tier,omitempty"`
}

// SystemData metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// CreatedBy - The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`
	// CreatedByType - The type of identity that created the resource. Possible values include: 'User', 'Application', 'ManagedIdentity', 'Key'
	CreatedByType CreatedByType `json:"createdByType,omitempty"`
	// CreatedAt - The timestamp of resource creation (UTC).
	CreatedAt *date.Time `json:"createdAt,omitempty"`
	// LastModifiedBy - The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// LastModifiedByType - The type of identity that last modified the resource. Possible values include: 'User', 'Application', 'ManagedIdentity', 'Key'
	LastModifiedByType CreatedByType `json:"lastModifiedByType,omitempty"`
	// LastModifiedAt - The type of identity that last modified the resource.
	LastModifiedAt *date.Time `json:"lastModifiedAt,omitempty"`
}

// TrackedResource the resource model definition for a ARM tracked top level resource
type TrackedResource struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	if tr.Location != nil {
		objectMap["location"] = tr.Location
	}
	return json.Marshal(objectMap)
}
