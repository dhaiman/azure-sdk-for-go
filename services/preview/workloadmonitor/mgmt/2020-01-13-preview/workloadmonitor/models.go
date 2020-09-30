package workloadmonitor

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
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/workloadmonitor/mgmt/2020-01-13-preview/workloadmonitor"

// DefaultError error body contract.
type DefaultError struct {
	// Error - Details about the error
	Error *DefaultErrorError `json:"error,omitempty"`
}

// DefaultErrorError details about the error
type DefaultErrorError struct {
	// Code - Service-defined error code. This code serves as a sub-status for the HTTP error code specified in the response.
	Code *string `json:"code,omitempty"`
	// Message - Human-readable representation of the error.
	Message *string `json:"message,omitempty"`
}

// Monitor information about a monitor.
type Monitor struct {
	autorest.Response `json:"-"`
	// ID - Arm ID of this monitor.
	ID *string `json:"id,omitempty"`
	// Name - Url-encoded monitor name.
	Name *string `json:"name,omitempty"`
	// Type - Type of ARM resource.
	Type *string `json:"type,omitempty"`
	// MonitorProperties - Properties of the monitor.
	*MonitorProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for Monitor.
func (mVar Monitor) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if mVar.ID != nil {
		objectMap["id"] = mVar.ID
	}
	if mVar.Name != nil {
		objectMap["name"] = mVar.Name
	}
	if mVar.Type != nil {
		objectMap["type"] = mVar.Type
	}
	if mVar.MonitorProperties != nil {
		objectMap["properties"] = mVar.MonitorProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Monitor struct.
func (mVar *Monitor) UnmarshalJSON(body []byte) error {
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
				mVar.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				mVar.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				mVar.Type = &typeVar
			}
		case "properties":
			if v != nil {
				var monitorProperties MonitorProperties
				err = json.Unmarshal(*v, &monitorProperties)
				if err != nil {
					return err
				}
				mVar.MonitorProperties = &monitorProperties
			}
		}
	}

	return nil
}

// MonitorList basic information about the current status of a monitor.
type MonitorList struct {
	autorest.Response `json:"-"`
	// Value - Array of monitors.
	Value *[]Monitor `json:"value,omitempty"`
	// NextLink - Link to next page if list is too long.
	NextLink *string `json:"nextLink,omitempty"`
}

// MonitorListIterator provides access to a complete listing of Monitor values.
type MonitorListIterator struct {
	i    int
	page MonitorListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *MonitorListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorListIterator.NextWithContext")
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
func (iter *MonitorListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter MonitorListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter MonitorListIterator) Response() MonitorList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter MonitorListIterator) Value() Monitor {
	if !iter.page.NotDone() {
		return Monitor{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the MonitorListIterator type.
func NewMonitorListIterator(page MonitorListPage) MonitorListIterator {
	return MonitorListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (ml MonitorList) IsEmpty() bool {
	return ml.Value == nil || len(*ml.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (ml MonitorList) hasNextLink() bool {
	return ml.NextLink != nil && len(*ml.NextLink) != 0
}

// monitorListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ml MonitorList) monitorListPreparer(ctx context.Context) (*http.Request, error) {
	if !ml.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ml.NextLink)))
}

// MonitorListPage contains a page of Monitor values.
type MonitorListPage struct {
	fn func(context.Context, MonitorList) (MonitorList, error)
	ml MonitorList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *MonitorListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.ml)
		if err != nil {
			return err
		}
		page.ml = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *MonitorListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page MonitorListPage) NotDone() bool {
	return !page.ml.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page MonitorListPage) Response() MonitorList {
	return page.ml
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page MonitorListPage) Values() []Monitor {
	if page.ml.IsEmpty() {
		return nil
	}
	return *page.ml.Value
}

// Creates a new instance of the MonitorListPage type.
func NewMonitorListPage(getNextPage func(context.Context, MonitorList) (MonitorList, error)) MonitorListPage {
	return MonitorListPage{fn: getNextPage}
}

// MonitorProperties properties of the monitor.
type MonitorProperties struct {
	// MonitorName - Human-readable name of this monitor.
	MonitorName *string `json:"monitorName,omitempty"`
	// MonitorType - Type of this monitor.
	MonitorType *string `json:"monitorType,omitempty"`
	// MonitoredObject - Dynamic monitored object of this monitor.
	MonitoredObject *string `json:"monitoredObject,omitempty"`
	// ParentMonitorName - Name of this monitor's parent.
	ParentMonitorName *string `json:"parentMonitorName,omitempty"`
	// PreviousMonitorState - Current health state of this monitor. Possible values include: 'Healthy', 'Critical', 'Warning', 'Unknown'
	PreviousMonitorState HealthState `json:"previousMonitorState,omitempty"`
	// CurrentMonitorState - Current health state of this monitor. Possible values include: 'Healthy', 'Critical', 'Warning', 'Unknown'
	CurrentMonitorState HealthState `json:"currentMonitorState,omitempty"`
	// EvaluationTimestamp - Timestamp that this monitor was last evaluated.
	EvaluationTimestamp *string `json:"evaluationTimestamp,omitempty"`
	// CurrentStateFirstObservedTimestamp - Timestamp of this monitor's last state change.
	CurrentStateFirstObservedTimestamp *string `json:"currentStateFirstObservedTimestamp,omitempty"`
	// LastReportedTimestamp - Timestamp of this monitor's last reported state.
	LastReportedTimestamp *string `json:"lastReportedTimestamp,omitempty"`
	// Evidence - Evidence of this monitor's last state change.
	Evidence interface{} `json:"evidence,omitempty"`
	// MonitorConfiguration - Configuration settings at the time of this monitor's last state change.
	MonitorConfiguration interface{} `json:"monitorConfiguration,omitempty"`
}

// MonitorStateChange information about a state transition of a monitor.
type MonitorStateChange struct {
	autorest.Response `json:"-"`
	// ID - Arm ID of this monitor.
	ID *string `json:"id,omitempty"`
	// Name - Timestamp of state change (in unix).
	Name *string `json:"name,omitempty"`
	// Type - Type of ARM resource.
	Type *string `json:"type,omitempty"`
	// MonitorStateChangeProperties - Properties of the monitor.
	*MonitorStateChangeProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for MonitorStateChange.
func (msc MonitorStateChange) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if msc.ID != nil {
		objectMap["id"] = msc.ID
	}
	if msc.Name != nil {
		objectMap["name"] = msc.Name
	}
	if msc.Type != nil {
		objectMap["type"] = msc.Type
	}
	if msc.MonitorStateChangeProperties != nil {
		objectMap["properties"] = msc.MonitorStateChangeProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for MonitorStateChange struct.
func (msc *MonitorStateChange) UnmarshalJSON(body []byte) error {
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
				msc.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				msc.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				msc.Type = &typeVar
			}
		case "properties":
			if v != nil {
				var monitorStateChangeProperties MonitorStateChangeProperties
				err = json.Unmarshal(*v, &monitorStateChangeProperties)
				if err != nil {
					return err
				}
				msc.MonitorStateChangeProperties = &monitorStateChangeProperties
			}
		}
	}

	return nil
}

// MonitorStateChangeList the monitor history of a monitor
type MonitorStateChangeList struct {
	autorest.Response `json:"-"`
	// Value - Array of state change transitions.
	Value *[]MonitorStateChange `json:"value,omitempty"`
	// NextLink - Link to next page if list is too long.
	NextLink *string `json:"nextLink,omitempty"`
}

// MonitorStateChangeListIterator provides access to a complete listing of MonitorStateChange values.
type MonitorStateChangeListIterator struct {
	i    int
	page MonitorStateChangeListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *MonitorStateChangeListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorStateChangeListIterator.NextWithContext")
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
func (iter *MonitorStateChangeListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter MonitorStateChangeListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter MonitorStateChangeListIterator) Response() MonitorStateChangeList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter MonitorStateChangeListIterator) Value() MonitorStateChange {
	if !iter.page.NotDone() {
		return MonitorStateChange{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the MonitorStateChangeListIterator type.
func NewMonitorStateChangeListIterator(page MonitorStateChangeListPage) MonitorStateChangeListIterator {
	return MonitorStateChangeListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (mscl MonitorStateChangeList) IsEmpty() bool {
	return mscl.Value == nil || len(*mscl.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (mscl MonitorStateChangeList) hasNextLink() bool {
	return mscl.NextLink != nil && len(*mscl.NextLink) != 0
}

// monitorStateChangeListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (mscl MonitorStateChangeList) monitorStateChangeListPreparer(ctx context.Context) (*http.Request, error) {
	if !mscl.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(mscl.NextLink)))
}

// MonitorStateChangeListPage contains a page of MonitorStateChange values.
type MonitorStateChangeListPage struct {
	fn   func(context.Context, MonitorStateChangeList) (MonitorStateChangeList, error)
	mscl MonitorStateChangeList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *MonitorStateChangeListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MonitorStateChangeListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.mscl)
		if err != nil {
			return err
		}
		page.mscl = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *MonitorStateChangeListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page MonitorStateChangeListPage) NotDone() bool {
	return !page.mscl.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page MonitorStateChangeListPage) Response() MonitorStateChangeList {
	return page.mscl
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page MonitorStateChangeListPage) Values() []MonitorStateChange {
	if page.mscl.IsEmpty() {
		return nil
	}
	return *page.mscl.Value
}

// Creates a new instance of the MonitorStateChangeListPage type.
func NewMonitorStateChangeListPage(getNextPage func(context.Context, MonitorStateChangeList) (MonitorStateChangeList, error)) MonitorStateChangeListPage {
	return MonitorStateChangeListPage{fn: getNextPage}
}

// MonitorStateChangeProperties properties of the monitor.
type MonitorStateChangeProperties struct {
	// MonitorName - Human-readable name of this monitor.
	MonitorName *string `json:"monitorName,omitempty"`
	// MonitorType - Type of this monitor.
	MonitorType *string `json:"monitorType,omitempty"`
	// MonitoredObject - Dynamic monitored object of this monitor.
	MonitoredObject *string `json:"monitoredObject,omitempty"`
	// EvaluationTimestamp - Timestamp of that this event ocurred.
	EvaluationTimestamp *string `json:"evaluationTimestamp,omitempty"`
	// CurrentStateFirstObservedTimestamp - Timestamp of that this health state first ocurred.
	CurrentStateFirstObservedTimestamp *string `json:"currentStateFirstObservedTimestamp,omitempty"`
	// PreviousMonitorState - Previous health state. Possible values include: 'Healthy', 'Critical', 'Warning', 'Unknown'
	PreviousMonitorState HealthState `json:"previousMonitorState,omitempty"`
	// CurrentMonitorState - New health state. Possible values include: 'Healthy', 'Critical', 'Warning', 'Unknown'
	CurrentMonitorState HealthState `json:"currentMonitorState,omitempty"`
	// Evidence - Evidence of this monitor's last state change.
	Evidence interface{} `json:"evidence,omitempty"`
	// MonitorConfiguration - Configuration settings at the time of this monitor's last state change.
	MonitorConfiguration interface{} `json:"monitorConfiguration,omitempty"`
}

// Operation operation supported by the resource provider.
type Operation struct {
	// Name - Name of the operation.
	Name *string `json:"name,omitempty"`
	// Display - The properties of the resource operation.
	Display *OperationDisplay `json:"display,omitempty"`
	// Origin - The origin of the operation.
	Origin *string `json:"origin,omitempty"`
}

// OperationDisplay the properties of the resource operation.
type OperationDisplay struct {
	// Provider - Provider name of this operation.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource name of this operation.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation name of the operation.
	Operation *string `json:"operation,omitempty"`
	// Description - Description of the operation.
	Description *string `json:"description,omitempty"`
}

// OperationList list of possible operations.
type OperationList struct {
	autorest.Response `json:"-"`
	// Value - Array of possible operations.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - Link to next page if list is too long.
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

// hasNextLink returns true if the NextLink is not empty.
func (ol OperationList) hasNextLink() bool {
	return ol.NextLink != nil && len(*ol.NextLink) != 0
}

// operationListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ol OperationList) operationListPreparer(ctx context.Context) (*http.Request, error) {
	if !ol.hasNextLink() {
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
	for {
		next, err := page.fn(ctx, page.ol)
		if err != nil {
			return err
		}
		page.ol = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
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
