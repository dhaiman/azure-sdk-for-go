package portal

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
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/portal/mgmt/2015-08-01-preview/portal"

// Dashboard the shared dashboard resource definition.
type Dashboard struct {
	autorest.Response `json:"-"`
	// DashboardProperties - The shared dashboard properties.
	*DashboardProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Resource Id
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type
	Type *string `json:"type,omitempty"`
	// Location - Resource location
	Location *string `json:"location,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Dashboard.
func (d Dashboard) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if d.DashboardProperties != nil {
		objectMap["properties"] = d.DashboardProperties
	}
	if d.Location != nil {
		objectMap["location"] = d.Location
	}
	if d.Tags != nil {
		objectMap["tags"] = d.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Dashboard struct.
func (d *Dashboard) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var dashboardProperties DashboardProperties
				err = json.Unmarshal(*v, &dashboardProperties)
				if err != nil {
					return err
				}
				d.DashboardProperties = &dashboardProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				d.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				d.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				d.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				d.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				d.Tags = tags
			}
		}
	}

	return nil
}

// DashboardLens a dashboard lens.
type DashboardLens struct {
	// Order - The lens order.
	Order *int32 `json:"order,omitempty"`
	// Parts - The dashboard parts.
	Parts map[string]*DashboardParts `json:"parts"`
	// Metadata - The dashboard len's metadata.
	Metadata map[string]interface{} `json:"metadata"`
}

// MarshalJSON is the custom marshaler for DashboardLens.
func (dl DashboardLens) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if dl.Order != nil {
		objectMap["order"] = dl.Order
	}
	if dl.Parts != nil {
		objectMap["parts"] = dl.Parts
	}
	if dl.Metadata != nil {
		objectMap["metadata"] = dl.Metadata
	}
	return json.Marshal(objectMap)
}

// DashboardListResult list of dashboards.
type DashboardListResult struct {
	autorest.Response `json:"-"`
	// Value - The array of custom resource provider manifests.
	Value *[]Dashboard `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// DashboardListResultIterator provides access to a complete listing of Dashboard values.
type DashboardListResultIterator struct {
	i    int
	page DashboardListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *DashboardListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DashboardListResultIterator.NextWithContext")
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
func (iter *DashboardListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter DashboardListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter DashboardListResultIterator) Response() DashboardListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter DashboardListResultIterator) Value() Dashboard {
	if !iter.page.NotDone() {
		return Dashboard{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the DashboardListResultIterator type.
func NewDashboardListResultIterator(page DashboardListResultPage) DashboardListResultIterator {
	return DashboardListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (dlr DashboardListResult) IsEmpty() bool {
	return dlr.Value == nil || len(*dlr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (dlr DashboardListResult) hasNextLink() bool {
	return dlr.NextLink != nil && len(*dlr.NextLink) != 0
}

// dashboardListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (dlr DashboardListResult) dashboardListResultPreparer(ctx context.Context) (*http.Request, error) {
	if !dlr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(dlr.NextLink)))
}

// DashboardListResultPage contains a page of Dashboard values.
type DashboardListResultPage struct {
	fn  func(context.Context, DashboardListResult) (DashboardListResult, error)
	dlr DashboardListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *DashboardListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DashboardListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.dlr)
		if err != nil {
			return err
		}
		page.dlr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *DashboardListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page DashboardListResultPage) NotDone() bool {
	return !page.dlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page DashboardListResultPage) Response() DashboardListResult {
	return page.dlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page DashboardListResultPage) Values() []Dashboard {
	if page.dlr.IsEmpty() {
		return nil
	}
	return *page.dlr.Value
}

// Creates a new instance of the DashboardListResultPage type.
func NewDashboardListResultPage(cur DashboardListResult, getNextPage func(context.Context, DashboardListResult) (DashboardListResult, error)) DashboardListResultPage {
	return DashboardListResultPage{
		fn:  getNextPage,
		dlr: cur,
	}
}

// DashboardParts a dashboard part.
type DashboardParts struct {
	// Position - The dashboard's part position.
	Position *DashboardPartsPosition `json:"position,omitempty"`
	// Metadata - The dashboard part's metadata.
	Metadata map[string]interface{} `json:"metadata"`
}

// MarshalJSON is the custom marshaler for DashboardParts.
func (dp DashboardParts) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if dp.Position != nil {
		objectMap["position"] = dp.Position
	}
	if dp.Metadata != nil {
		objectMap["metadata"] = dp.Metadata
	}
	return json.Marshal(objectMap)
}

// DashboardPartsPosition the dashboard's part position.
type DashboardPartsPosition struct {
	// X - The dashboard's part x coordinate.
	X *int32 `json:"x,omitempty"`
	// Y - The dashboard's part y coordinate.
	Y *int32 `json:"y,omitempty"`
	// RowSpan - The dashboard's part row span.
	RowSpan *int32 `json:"rowSpan,omitempty"`
	// ColSpan - The dashboard's part column span.
	ColSpan *int32 `json:"colSpan,omitempty"`
	// Metadata - The dashboard part's metadata.
	Metadata map[string]interface{} `json:"metadata"`
}

// MarshalJSON is the custom marshaler for DashboardPartsPosition.
func (dp DashboardPartsPosition) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if dp.X != nil {
		objectMap["x"] = dp.X
	}
	if dp.Y != nil {
		objectMap["y"] = dp.Y
	}
	if dp.RowSpan != nil {
		objectMap["rowSpan"] = dp.RowSpan
	}
	if dp.ColSpan != nil {
		objectMap["colSpan"] = dp.ColSpan
	}
	if dp.Metadata != nil {
		objectMap["metadata"] = dp.Metadata
	}
	return json.Marshal(objectMap)
}

// DashboardProperties the shared dashboard properties.
type DashboardProperties struct {
	// Lenses - The dashboard lenses.
	Lenses map[string]*DashboardLens `json:"lenses"`
	// Metadata - The dashboard metadata.
	Metadata map[string]interface{} `json:"metadata"`
}

// MarshalJSON is the custom marshaler for DashboardProperties.
func (dp DashboardProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if dp.Lenses != nil {
		objectMap["lenses"] = dp.Lenses
	}
	if dp.Metadata != nil {
		objectMap["metadata"] = dp.Metadata
	}
	return json.Marshal(objectMap)
}

// ErrorDefinition error definition.
type ErrorDefinition struct {
	// Code - READ-ONLY; Service specific error code which serves as the substatus for the HTTP error code.
	Code *string `json:"code,omitempty"`
	// Message - READ-ONLY; Description of the error.
	Message *string `json:"message,omitempty"`
	// Details - READ-ONLY; Internal error details.
	Details *[]ErrorDefinition `json:"details,omitempty"`
}

// ErrorResponse error response.
type ErrorResponse struct {
	// Error - The error details.
	Error *ErrorDefinition `json:"error,omitempty"`
}

// PatchableDashboard the shared dashboard resource definition.
type PatchableDashboard struct {
	// DashboardProperties - The shared dashboard properties.
	*DashboardProperties `json:"properties,omitempty"`
	// Tags - Resource tags
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for PatchableDashboard.
func (pd PatchableDashboard) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if pd.DashboardProperties != nil {
		objectMap["properties"] = pd.DashboardProperties
	}
	if pd.Tags != nil {
		objectMap["tags"] = pd.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for PatchableDashboard struct.
func (pd *PatchableDashboard) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var dashboardProperties DashboardProperties
				err = json.Unmarshal(*v, &dashboardProperties)
				if err != nil {
					return err
				}
				pd.DashboardProperties = &dashboardProperties
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				pd.Tags = tags
			}
		}
	}

	return nil
}

// ResourceProviderOperation supported operations of this resource provider.
type ResourceProviderOperation struct {
	// Name - Operation name, in format of {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// IsDataAction - Indicates whether the operation applies to data-plane.
	IsDataAction *string `json:"isDataAction,omitempty"`
	// Display - Display metadata associated with the operation.
	Display *ResourceProviderOperationDisplay `json:"display,omitempty"`
}

// ResourceProviderOperationDisplay display metadata associated with the operation.
type ResourceProviderOperationDisplay struct {
	// Provider - Resource provider: Microsoft Custom Providers.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
	// Operation - Type of operation: get, read, delete, etc.
	Operation *string `json:"operation,omitempty"`
	// Description - Description of this operation.
	Description *string `json:"description,omitempty"`
}

// ResourceProviderOperationList results of the request to list operations.
type ResourceProviderOperationList struct {
	autorest.Response `json:"-"`
	// Value - List of operations supported by this resource provider.
	Value *[]ResourceProviderOperation `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// ResourceProviderOperationListIterator provides access to a complete listing of ResourceProviderOperation
// values.
type ResourceProviderOperationListIterator struct {
	i    int
	page ResourceProviderOperationListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ResourceProviderOperationListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceProviderOperationListIterator.NextWithContext")
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
func (iter *ResourceProviderOperationListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ResourceProviderOperationListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ResourceProviderOperationListIterator) Response() ResourceProviderOperationList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ResourceProviderOperationListIterator) Value() ResourceProviderOperation {
	if !iter.page.NotDone() {
		return ResourceProviderOperation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ResourceProviderOperationListIterator type.
func NewResourceProviderOperationListIterator(page ResourceProviderOperationListPage) ResourceProviderOperationListIterator {
	return ResourceProviderOperationListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (rpol ResourceProviderOperationList) IsEmpty() bool {
	return rpol.Value == nil || len(*rpol.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (rpol ResourceProviderOperationList) hasNextLink() bool {
	return rpol.NextLink != nil && len(*rpol.NextLink) != 0
}

// resourceProviderOperationListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (rpol ResourceProviderOperationList) resourceProviderOperationListPreparer(ctx context.Context) (*http.Request, error) {
	if !rpol.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(rpol.NextLink)))
}

// ResourceProviderOperationListPage contains a page of ResourceProviderOperation values.
type ResourceProviderOperationListPage struct {
	fn   func(context.Context, ResourceProviderOperationList) (ResourceProviderOperationList, error)
	rpol ResourceProviderOperationList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ResourceProviderOperationListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ResourceProviderOperationListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.rpol)
		if err != nil {
			return err
		}
		page.rpol = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ResourceProviderOperationListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ResourceProviderOperationListPage) NotDone() bool {
	return !page.rpol.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ResourceProviderOperationListPage) Response() ResourceProviderOperationList {
	return page.rpol
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ResourceProviderOperationListPage) Values() []ResourceProviderOperation {
	if page.rpol.IsEmpty() {
		return nil
	}
	return *page.rpol.Value
}

// Creates a new instance of the ResourceProviderOperationListPage type.
func NewResourceProviderOperationListPage(cur ResourceProviderOperationList, getNextPage func(context.Context, ResourceProviderOperationList) (ResourceProviderOperationList, error)) ResourceProviderOperationListPage {
	return ResourceProviderOperationListPage{
		fn:   getNextPage,
		rpol: cur,
	}
}
