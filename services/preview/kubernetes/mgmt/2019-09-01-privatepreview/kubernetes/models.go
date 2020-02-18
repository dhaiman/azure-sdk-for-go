package kubernetes

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
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/kubernetes/mgmt/2019-09-01-privatepreview/kubernetes"

// AzureEnvironment enumerates the values for azure environment.
type AzureEnvironment string

const (
	// AzureChinaCloud ...
	AzureChinaCloud AzureEnvironment = "AzureChinaCloud"
	// AzureGermanCloud ...
	AzureGermanCloud AzureEnvironment = "AzureGermanCloud"
	// AzurePublicCloud ...
	AzurePublicCloud AzureEnvironment = "AzurePublicCloud"
	// AzureUSGovernmentCloud ...
	AzureUSGovernmentCloud AzureEnvironment = "AzureUSGovernmentCloud"
)

// PossibleAzureEnvironmentValues returns an array of possible values for the AzureEnvironment const type.
func PossibleAzureEnvironmentValues() []AzureEnvironment {
	return []AzureEnvironment{AzureChinaCloud, AzureGermanCloud, AzurePublicCloud, AzureUSGovernmentCloud}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// None ...
	None ResourceIdentityType = "None"
	// SystemAssigned ...
	SystemAssigned ResourceIdentityType = "SystemAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{None, SystemAssigned}
}

// Type enumerates the values for type.
type Type string

const (
	// MicrosoftKubernetesconnectedClusters ...
	MicrosoftKubernetesconnectedClusters Type = "Microsoft.Kubernetes/connectedClusters"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{MicrosoftKubernetesconnectedClusters}
}

// AccessProfile ...
type AccessProfile struct {
	// KubeConfig - Base64 Encoded kubeconfig for accessing target cluster
	KubeConfig       *string                           `json:"kubeConfig,omitempty"`
	AadAccessProfile *ConnectedClusterAADAccessProfile `json:"aadAccessProfile,omitempty"`
}

// ConnectedCluster ...
type ConnectedCluster struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Type of the resource requested. Possible values include: 'MicrosoftKubernetesconnectedClusters'
	Type Type `json:"type,omitempty"`
	// Location - Location of the cluster
	Location *string `json:"location,omitempty"`
	// Tags - Connected Cluster Resource Tags.
	Tags map[string]*string `json:"tags"`
	// Identity - The identity of the connected cluster, if configured.
	Identity                    *ConnectedClusterIdentity `json:"identity,omitempty"`
	*ConnectedClusterProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for ConnectedCluster.
func (cc ConnectedCluster) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cc.Location != nil {
		objectMap["location"] = cc.Location
	}
	if cc.Tags != nil {
		objectMap["tags"] = cc.Tags
	}
	if cc.Identity != nil {
		objectMap["identity"] = cc.Identity
	}
	if cc.ConnectedClusterProperties != nil {
		objectMap["properties"] = cc.ConnectedClusterProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ConnectedCluster struct.
func (cc *ConnectedCluster) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				cc.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				cc.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar Type
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				cc.Type = typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				cc.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				cc.Tags = tags
			}
		case "identity":
			if v != nil {
				var identity ConnectedClusterIdentity
				err = json.Unmarshal(*v, &identity)
				if err != nil {
					return err
				}
				cc.Identity = &identity
			}
		case "properties":
			if v != nil {
				var connectedClusterProperties ConnectedClusterProperties
				err = json.Unmarshal(*v, &connectedClusterProperties)
				if err != nil {
					return err
				}
				cc.ConnectedClusterProperties = &connectedClusterProperties
			}
		}
	}

	return nil
}

// ConnectedClusterAADAccessProfile ...
type ConnectedClusterAADAccessProfile struct {
	// EndPoint - The endpoint to connect to the cluster
	EndPoint *string `json:"endPoint,omitempty"`
	// EndpointAuthorityData - The end point authority data to connect to the cluster
	EndpointAuthorityData *string `json:"endpointAuthorityData,omitempty"`
	// APIServerID - AAD API Server Id
	APIServerID *string `json:"apiServerId,omitempty"`
	// ClientID - Id of the client
	ClientID *string `json:"clientId,omitempty"`
	// Environment - The azure environment. Possible values include: 'AzurePublicCloud', 'AzureUSGovernmentCloud', 'AzureChinaCloud', 'AzureGermanCloud'
	Environment AzureEnvironment `json:"environment,omitempty"`
	// TenantID - AAD Tenant Id
	TenantID *string `json:"tenantId,omitempty"`
}

// ConnectedClusterAADProfile ...
type ConnectedClusterAADProfile struct {
	// TenantID - The aad tenant id which is configured on target K8s cluster
	TenantID *string `json:"tenantId,omitempty"`
	// ClientAppID - The client app id configured on target K8 cluster
	ClientAppID *string `json:"clientAppId,omitempty"`
	// ServerAppID - The server app id to access AD server
	ServerAppID *string `json:"serverAppId,omitempty"`
}

// ConnectedClusterAccessProfile ...
type ConnectedClusterAccessProfile struct {
	autorest.Response `json:"-"`
	// Type - Type of the resource on which the credentials are requested
	Type *string `json:"type,omitempty"`
	// Location - Location of the cluster
	Location       *string `json:"location,omitempty"`
	*AccessProfile `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for ConnectedClusterAccessProfile.
func (ccap ConnectedClusterAccessProfile) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ccap.Type != nil {
		objectMap["type"] = ccap.Type
	}
	if ccap.Location != nil {
		objectMap["location"] = ccap.Location
	}
	if ccap.AccessProfile != nil {
		objectMap["properties"] = ccap.AccessProfile
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ConnectedClusterAccessProfile struct.
func (ccap *ConnectedClusterAccessProfile) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				ccap.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				ccap.Location = &location
			}
		case "properties":
			if v != nil {
				var accessProfile AccessProfile
				err = json.Unmarshal(*v, &accessProfile)
				if err != nil {
					return err
				}
				ccap.AccessProfile = &accessProfile
			}
		}
	}

	return nil
}

// ConnectedClusterIdentity identity for the virtual machine.
type ConnectedClusterIdentity struct {
	// PrincipalID - READ-ONLY; The principal id of connected cluster identity. This property will only be provided for a system assigned identity.
	PrincipalID *string `json:"principalId,omitempty"`
	// TenantID - READ-ONLY; The tenant id associated with the connected cluster. This property will only be provided for a system assigned identity.
	TenantID *string `json:"tenantId,omitempty"`
	// Type - The type of identity used for the connected cluster. The type 'SystemAssigned includes a system created identity. The type 'None' will remove any identities from the virtual machine. Possible values include: 'SystemAssigned', 'None'
	Type ResourceIdentityType `json:"type,omitempty"`
}

// ConnectedClusterList the paginated list of connected Clusters
type ConnectedClusterList struct {
	autorest.Response `json:"-"`
	// Value - The list of connected clusters
	Value *[]ConnectedCluster `json:"value,omitempty"`
	// NextLink - The link to fetch the next page of connected cluster
	NextLink *string `json:"nextLink,omitempty"`
}

// ConnectedClusterListIterator provides access to a complete listing of ConnectedCluster values.
type ConnectedClusterListIterator struct {
	i    int
	page ConnectedClusterListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ConnectedClusterListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConnectedClusterListIterator.NextWithContext")
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
func (iter *ConnectedClusterListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ConnectedClusterListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ConnectedClusterListIterator) Response() ConnectedClusterList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ConnectedClusterListIterator) Value() ConnectedCluster {
	if !iter.page.NotDone() {
		return ConnectedCluster{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ConnectedClusterListIterator type.
func NewConnectedClusterListIterator(page ConnectedClusterListPage) ConnectedClusterListIterator {
	return ConnectedClusterListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (ccl ConnectedClusterList) IsEmpty() bool {
	return ccl.Value == nil || len(*ccl.Value) == 0
}

// connectedClusterListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ccl ConnectedClusterList) connectedClusterListPreparer(ctx context.Context) (*http.Request, error) {
	if ccl.NextLink == nil || len(to.String(ccl.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ccl.NextLink)))
}

// ConnectedClusterListPage contains a page of ConnectedCluster values.
type ConnectedClusterListPage struct {
	fn  func(context.Context, ConnectedClusterList) (ConnectedClusterList, error)
	ccl ConnectedClusterList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ConnectedClusterListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ConnectedClusterListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.ccl)
	if err != nil {
		return err
	}
	page.ccl = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ConnectedClusterListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ConnectedClusterListPage) NotDone() bool {
	return !page.ccl.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ConnectedClusterListPage) Response() ConnectedClusterList {
	return page.ccl
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ConnectedClusterListPage) Values() []ConnectedCluster {
	if page.ccl.IsEmpty() {
		return nil
	}
	return *page.ccl.Value
}

// Creates a new instance of the ConnectedClusterListPage type.
func NewConnectedClusterListPage(getNextPage func(context.Context, ConnectedClusterList) (ConnectedClusterList, error)) ConnectedClusterListPage {
	return ConnectedClusterListPage{fn: getNextPage}
}

// ConnectedClusterPatch object containing updates for patch operations.
type ConnectedClusterPatch struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for ConnectedClusterPatch.
func (ccp ConnectedClusterPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ccp.Tags != nil {
		objectMap["tags"] = ccp.Tags
	}
	return json.Marshal(objectMap)
}

// ConnectedClusterProperties ...
type ConnectedClusterProperties struct {
	// AgentPublicKeyCertificate - Base64 encoded public certificate used by the agent to do the initial handshake to the backend services in Azure.
	AgentPublicKeyCertificate *string                     `json:"agentPublicKeyCertificate,omitempty"`
	AadProfile                *ConnectedClusterAADProfile `json:"aadProfile,omitempty"`
	// KubernetesVersion - The Kubernetes version of the connected cluster resource
	KubernetesVersion *string `json:"kubernetesVersion,omitempty"`
	// TotalNodeCount - Number of nodes present in the connected cluster resource
	TotalNodeCount *int32 `json:"totalNodeCount,omitempty"`
	// AgentVersion - Version of the agent running on the connected cluster resource
	AgentVersion *string `json:"agentVersion,omitempty"`
}

// ErrorDetails the error response details containing error code and error message
type ErrorDetails struct {
	// Code - READ-ONLY; The error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; The error message.
	Message *string `json:"message,omitempty"`
}

// ErrorResponse the error response that indicates why an operation has failed.
type ErrorResponse struct {
	Error *ErrorDetails `json:"error,omitempty"`
}

// Operation the Connected cluster API operation
type Operation struct {
	// Name - READ-ONLY; Operation name: {Microsoft.Kubernetes}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - READ-ONLY; The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.connectedClusters
	Provider *string `json:"provider,omitempty"`
	// Resource - Connected Cluster Resource on which the operation is performed
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
	// Description - Description of the operation.
	Description *string `json:"description,omitempty"`
}

// OperationList the paginated list of connected cluster API operations.
type OperationList struct {
	autorest.Response `json:"-"`
	// Value - READ-ONLY; The list of connected cluster API operations.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - The link to fetch the next page of connected cluster API operations.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListIterator provides access to a complete listing of Operation values.
type OperationListIterator struct {
	i    int
	page OperationListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListIterator.NextWithContext")
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
func (iter *OperationListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListIterator) Response() OperationList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationListIterator type.
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return OperationListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (ol OperationList) IsEmpty() bool {
	return ol.Value == nil || len(*ol.Value) == 0
}

// operationListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ol OperationList) operationListPreparer(ctx context.Context) (*http.Request, error) {
	if ol.NextLink == nil || len(to.String(ol.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ol.NextLink)))
}

// OperationListPage contains a page of Operation values.
type OperationListPage struct {
	fn func(context.Context, OperationList) (OperationList, error)
	ol OperationList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.ol)
	if err != nil {
		return err
	}
	page.ol = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListPage) NotDone() bool {
	return !page.ol.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListPage) Response() OperationList {
	return page.ol
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListPage) Values() []Operation {
	if page.ol.IsEmpty() {
		return nil
	}
	return *page.ol.Value
}

// Creates a new instance of the OperationListPage type.
func NewOperationListPage(getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return OperationListPage{fn: getNextPage}
}
