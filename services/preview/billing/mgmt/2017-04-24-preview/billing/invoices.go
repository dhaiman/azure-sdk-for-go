package billing

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
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// InvoicesClient is the billing client provides access to billing resources for Azure Web-Direct subscriptions. Other
// subscription types which were not purchased directly through the Azure web portal are not supported through this
// preview API.
type InvoicesClient struct {
	BaseClient
}

// NewInvoicesClient creates an instance of the InvoicesClient client.
func NewInvoicesClient(subscriptionID string) InvoicesClient {
	return NewInvoicesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewInvoicesClientWithBaseURI creates an instance of the InvoicesClient client.
func NewInvoicesClientWithBaseURI(baseURI string, subscriptionID string) InvoicesClient {
	return InvoicesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a named invoice resource. When getting a single invoice, the downloadUrl property is expanded
// automatically.
// Parameters:
// invoiceName - the name of an invoice resource.
func (client InvoicesClient) Get(ctx context.Context, invoiceName string) (result Invoice, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvoicesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, invoiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client InvoicesClient) GetPreparer(ctx context.Context, invoiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"invoiceName":    autorest.Encode("path", invoiceName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Billing/invoices/{invoiceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client InvoicesClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client InvoicesClient) GetResponder(resp *http.Response) (result Invoice, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetLatest gets the most recent invoice. When getting a single invoice, the downloadUrl property is expanded
// automatically.
func (client InvoicesClient) GetLatest(ctx context.Context) (result Invoice, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvoicesClient.GetLatest")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetLatestPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "GetLatest", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetLatestSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "GetLatest", resp, "Failure sending request")
		return
	}

	result, err = client.GetLatestResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "GetLatest", resp, "Failure responding to request")
	}

	return
}

// GetLatestPreparer prepares the GetLatest request.
func (client InvoicesClient) GetLatestPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Billing/invoices/latest", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetLatestSender sends the GetLatest request. The method will close the
// http.Response Body if it receives an error.
func (client InvoicesClient) GetLatestSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetLatestResponder handles the response to the GetLatest request. The method always
// closes the http.Response Body.
func (client InvoicesClient) GetLatestResponder(resp *http.Response) (result Invoice, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the available invoices for a subscription in reverse chronological order beginning with the most recent
// invoice. In preview, invoices are available via this API only for invoice periods which end December 1, 2016 or
// later.
// Parameters:
// expand - may be used to expand the downloadUrl property within a list of invoices. This enables download
// links to be generated for multiple invoices at once. By default, downloadURLs are not included when listing
// invoices.
// filter - may be used to filter invoices by invoicePeriodEndDate. The filter supports 'eq', 'lt', 'gt', 'le',
// 'ge', and 'and'. It does not currently support 'ne', 'or', or 'not'.
// skiptoken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls.
// top - may be used to limit the number of results to the most recent N invoices.
func (client InvoicesClient) List(ctx context.Context, expand string, filter string, skiptoken string, top *int32) (result InvoicesListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvoicesClient.List")
		defer func() {
			sc := -1
			if result.ilr.Response.Response != nil {
				sc = result.ilr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(100), Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("billing.InvoicesClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, expand, filter, skiptoken, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ilr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "List", resp, "Failure sending request")
		return
	}

	result.ilr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client InvoicesClient) ListPreparer(ctx context.Context, expand string, filter string, skiptoken string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-24-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Billing/invoices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client InvoicesClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client InvoicesClient) ListResponder(resp *http.Response) (result InvoicesListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client InvoicesClient) listNextResults(ctx context.Context, lastResults InvoicesListResult) (result InvoicesListResult, err error) {
	req, err := lastResults.invoicesListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.InvoicesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.InvoicesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client InvoicesClient) ListComplete(ctx context.Context, expand string, filter string, skiptoken string, top *int32) (result InvoicesListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvoicesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, expand, filter, skiptoken, top)
	return
}
