package features

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
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/to"
    "github.com/Azure/go-autorest/tracing"
    "net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2015-12-01/features"

            // Operation microsoft.Features operation
            type Operation struct {
            // Name - Operation name: {provider}/{resource}/{operation}
            Name *string `json:"name,omitempty"`
            // Display - The object that represents the operation.
            Display *OperationDisplay `json:"display,omitempty"`
            }

            // OperationDisplay the object that represents the operation.
            type OperationDisplay struct {
            // Provider - Service provider: Microsoft.Features
            Provider *string `json:"provider,omitempty"`
            // Resource - Resource on which the operation is performed: Profile, endpoint, etc.
            Resource *string `json:"resource,omitempty"`
            // Operation - Operation type: Read, write, delete, etc.
            Operation *string `json:"operation,omitempty"`
            }

            // OperationListResult result of the request to list Microsoft.Features operations. It contains a list of
            // operations and a URL link to get the next set of results.
            type OperationListResult struct {
            autorest.Response `json:"-"`
            // Value - List of Microsoft.Features operations.
            Value *[]Operation `json:"value,omitempty"`
            // NextLink - URL to get the next set of operation list results if there are any.
            NextLink *string `json:"nextLink,omitempty"`
            }

            // OperationListResultIterator provides access to a complete listing of Operation values.
            type OperationListResultIterator struct {
                i int
                page OperationListResultPage
            }
        // NextWithContext advances to the next value.  If there was an error making
        // the request the iterator does not advance and the error is returned.
        func (iter * OperationListResultIterator) NextWithContext(ctx context.Context) (err error) {
        if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/OperationListResultIterator.NextWithContext")
        defer func() {
        sc := -1
        if iter.Response().Response.Response != nil {
        sc = iter.Response().Response.Response.StatusCode
        }
        tracing.EndSpan(ctx, sc, err)
        }()
        }
        iter.i++
        if iter.i < len(iter. page.Values()) {
        return nil
        }
        err = iter.page.NextWithContext(ctx)
        if err != nil {
        iter. i--
        return err
        }
        iter.i = 0
        return nil
        }
        // Next advances to the next value.  If there was an error making
        // the request the iterator does not advance and the error is returned.
        // Deprecated: Use NextWithContext() instead.
        func (iter * OperationListResultIterator) Next() error {
        return iter.NextWithContext(context.Background())
        }
        // NotDone returns true if the enumeration should be started or is not yet complete.
        func (iter OperationListResultIterator) NotDone() bool {
        return iter.page.NotDone() && iter.i < len(iter. page.Values())
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
        func NewOperationListResultIterator (page OperationListResultPage) OperationListResultIterator {
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
                    autorest.WithBaseURL(to.String( olr.NextLink)));
                    }

            // OperationListResultPage contains a page of Operation values.
            type OperationListResultPage struct {
                fn func(context.Context, OperationListResult) (OperationListResult, error)
                olr OperationListResult
            }

        // NextWithContext advances to the next page of values.  If there was an error making
        // the request the page does not advance and the error is returned.
        func (page * OperationListResultPage) NextWithContext(ctx context.Context) (err error) {
        if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/OperationListResultPage.NextWithContext")
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
        func (page * OperationListResultPage) Next() error {
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
        func NewOperationListResultPage (getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
            return OperationListResultPage{fn: getNextPage}
        }

            // OperationsListResult list of previewed features.
            type OperationsListResult struct {
            autorest.Response `json:"-"`
            // Value - The array of features.
            Value *[]Result `json:"value,omitempty"`
            // NextLink - The URL to use for getting the next set of results.
            NextLink *string `json:"nextLink,omitempty"`
            }

            // OperationsListResultIterator provides access to a complete listing of Result values.
            type OperationsListResultIterator struct {
                i int
                page OperationsListResultPage
            }
        // NextWithContext advances to the next value.  If there was an error making
        // the request the iterator does not advance and the error is returned.
        func (iter * OperationsListResultIterator) NextWithContext(ctx context.Context) (err error) {
        if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/OperationsListResultIterator.NextWithContext")
        defer func() {
        sc := -1
        if iter.Response().Response.Response != nil {
        sc = iter.Response().Response.Response.StatusCode
        }
        tracing.EndSpan(ctx, sc, err)
        }()
        }
        iter.i++
        if iter.i < len(iter. page.Values()) {
        return nil
        }
        err = iter.page.NextWithContext(ctx)
        if err != nil {
        iter. i--
        return err
        }
        iter.i = 0
        return nil
        }
        // Next advances to the next value.  If there was an error making
        // the request the iterator does not advance and the error is returned.
        // Deprecated: Use NextWithContext() instead.
        func (iter * OperationsListResultIterator) Next() error {
        return iter.NextWithContext(context.Background())
        }
        // NotDone returns true if the enumeration should be started or is not yet complete.
        func (iter OperationsListResultIterator) NotDone() bool {
        return iter.page.NotDone() && iter.i < len(iter. page.Values())
        }
        // Response returns the raw server response from the last page request.
        func (iter OperationsListResultIterator) Response() OperationsListResult {
        return iter.page.Response()
        }
        // Value returns the current value or a zero-initialized value if the
        // iterator has advanced beyond the end of the collection.
        func (iter OperationsListResultIterator) Value() Result {
        if !iter.page.NotDone() {
        return Result{}
        }
        return iter.page.Values()[iter.i]
        }
        // Creates a new instance of the OperationsListResultIterator type.
        func NewOperationsListResultIterator (page OperationsListResultPage) OperationsListResultIterator {
            return OperationsListResultIterator{page: page}
        }


                // IsEmpty returns true if the ListResult contains no values.
                func (olr OperationsListResult) IsEmpty() bool {
                return olr.Value == nil || len(*olr.Value) == 0
                }

                    // operationsListResultPreparer prepares a request to retrieve the next set of results.
                    // It returns nil if no more results exist.
                    func (olr OperationsListResult) operationsListResultPreparer(ctx context.Context) (*http.Request, error) {
                    if olr.NextLink == nil || len(to.String(olr.NextLink)) < 1 {
                    return nil, nil
                    }
                    return autorest.Prepare((&http.Request{}).WithContext(ctx),
                    autorest.AsJSON(),
                    autorest.AsGet(),
                    autorest.WithBaseURL(to.String( olr.NextLink)));
                    }

            // OperationsListResultPage contains a page of Result values.
            type OperationsListResultPage struct {
                fn func(context.Context, OperationsListResult) (OperationsListResult, error)
                olr OperationsListResult
            }

        // NextWithContext advances to the next page of values.  If there was an error making
        // the request the page does not advance and the error is returned.
        func (page * OperationsListResultPage) NextWithContext(ctx context.Context) (err error) {
        if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/OperationsListResultPage.NextWithContext")
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
        func (page * OperationsListResultPage) Next() error {
        return page.NextWithContext(context.Background())
        }
        // NotDone returns true if the page enumeration should be started or is not yet complete.
        func (page OperationsListResultPage) NotDone() bool {
        return !page.olr.IsEmpty()
        }
        // Response returns the raw server response from the last page request.
        func (page OperationsListResultPage) Response() OperationsListResult {
        return page.olr
        }
        // Values returns the slice of values for the current page or nil if there are no values.
        func (page OperationsListResultPage) Values() []Result {
        if page.olr.IsEmpty() {
        return nil
        }
        return *page.olr.Value
        }
        // Creates a new instance of the OperationsListResultPage type.
        func NewOperationsListResultPage (getNextPage func(context.Context, OperationsListResult) (OperationsListResult, error)) OperationsListResultPage {
            return OperationsListResultPage{fn: getNextPage}
        }

            // Properties information about feature.
            type Properties struct {
            // State - The registration state of the feature for the subscription.
            State *string `json:"state,omitempty"`
            }

            // Result previewed feature information.
            type Result struct {
            autorest.Response `json:"-"`
            // Name - The name of the feature.
            Name *string `json:"name,omitempty"`
            // Properties - Properties of the previewed feature.
            Properties *Properties `json:"properties,omitempty"`
            // ID - The resource ID of the feature.
            ID *string `json:"id,omitempty"`
            // Type - The resource type of the feature.
            Type *string `json:"type,omitempty"`
            }

