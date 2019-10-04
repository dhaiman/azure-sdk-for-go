package locks

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
const fqdn = "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2015-01-01/locks"

// LockLevel enumerates the values for lock level.
type LockLevel string

const (
	// CanNotDelete ...
	CanNotDelete LockLevel = "CanNotDelete"
	// NotSpecified ...
	NotSpecified LockLevel = "NotSpecified"
	// ReadOnly ...
	ReadOnly LockLevel = "ReadOnly"
)

// PossibleLockLevelValues returns an array of possible values for the LockLevel const type.
func PossibleLockLevelValues() []LockLevel {
	return []LockLevel{CanNotDelete, NotSpecified, ReadOnly}
}

// ManagementLockListResult list of management locks.
type ManagementLockListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of locks.
	Value *[]ManagementLockObject `json:"value,omitempty"`
	// NextLink - The URL to get the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// ManagementLockListResultIterator provides access to a complete listing of ManagementLockObject values.
type ManagementLockListResultIterator struct {
	i    int
	page ManagementLockListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ManagementLockListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagementLockListResultIterator.NextWithContext")
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
func (iter *ManagementLockListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ManagementLockListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ManagementLockListResultIterator) Response() ManagementLockListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ManagementLockListResultIterator) Value() ManagementLockObject {
	if !iter.page.NotDone() {
		return ManagementLockObject{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ManagementLockListResultIterator type.
func NewManagementLockListResultIterator(page ManagementLockListResultPage) ManagementLockListResultIterator {
	return ManagementLockListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (mllr ManagementLockListResult) IsEmpty() bool {
	return mllr.Value == nil || len(*mllr.Value) == 0
}

// managementLockListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (mllr ManagementLockListResult) managementLockListResultPreparer(ctx context.Context) (*http.Request, error) {
	if mllr.NextLink == nil || len(to.String(mllr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(mllr.NextLink)))
}

// ManagementLockListResultPage contains a page of ManagementLockObject values.
type ManagementLockListResultPage struct {
	fn   func(context.Context, ManagementLockListResult) (ManagementLockListResult, error)
	mllr ManagementLockListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ManagementLockListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagementLockListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.mllr)
	if err != nil {
		return err
	}
	page.mllr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ManagementLockListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ManagementLockListResultPage) NotDone() bool {
	return !page.mllr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ManagementLockListResultPage) Response() ManagementLockListResult {
	return page.mllr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ManagementLockListResultPage) Values() []ManagementLockObject {
	if page.mllr.IsEmpty() {
		return nil
	}
	return *page.mllr.Value
}

// Creates a new instance of the ManagementLockListResultPage type.
func NewManagementLockListResultPage(getNextPage func(context.Context, ManagementLockListResult) (ManagementLockListResult, error)) ManagementLockListResultPage {
	return ManagementLockListResultPage{fn: getNextPage}
}

// ManagementLockObject management lock information.
type ManagementLockObject struct {
	autorest.Response `json:"-"`
	// ManagementLockProperties - The properties of the lock.
	*ManagementLockProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; The Id of the lock.
	ID *string `json:"id,omitempty"`
	// Type - READ-ONLY; The type of the lock.
	Type *string `json:"type,omitempty"`
	// Name - The name of the lock.
	Name *string `json:"name,omitempty"`
}

// MarshalJSON is the custom marshaler for ManagementLockObject.
func (mlo ManagementLockObject) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if mlo.ManagementLockProperties != nil {
		objectMap["properties"] = mlo.ManagementLockProperties
	}
	if mlo.Name != nil {
		objectMap["name"] = mlo.Name
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ManagementLockObject struct.
func (mlo *ManagementLockObject) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var managementLockProperties ManagementLockProperties
				err = json.Unmarshal(*v, &managementLockProperties)
				if err != nil {
					return err
				}
				mlo.ManagementLockProperties = &managementLockProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				mlo.ID = &ID
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				mlo.Type = &typeVar
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				mlo.Name = &name
			}
		}
	}

	return nil
}

// ManagementLockProperties the management lock properties.
type ManagementLockProperties struct {
	// Level - The lock level of the management lock. Possible values include: 'NotSpecified', 'CanNotDelete', 'ReadOnly'
	Level LockLevel `json:"level,omitempty"`
	// Notes - The notes of the management lock.
	Notes *string `json:"notes,omitempty"`
}
