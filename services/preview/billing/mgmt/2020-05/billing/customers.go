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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// CustomersClient is the billing client provides access to billing resources for Azure subscriptions.
type CustomersClient struct {
	BaseClient
}

// NewCustomersClient creates an instance of the CustomersClient client.
func NewCustomersClient(subscriptionID string) CustomersClient {
	return NewCustomersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCustomersClientWithBaseURI creates an instance of the CustomersClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewCustomersClientWithBaseURI(baseURI string, subscriptionID string) CustomersClient {
	return CustomersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a customer by its ID. The operation is supported only for billing accounts with agreement type Microsoft
// Partner Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// customerName - the ID that uniquely identifies a customer.
// expand - may be used to expand enabledAzurePlans and resellers
func (client CustomersClient) Get(ctx context.Context, billingAccountName string, customerName string, expand string) (result Customer, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, billingAccountName, customerName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.CustomersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.CustomersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.CustomersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client CustomersClient) GetPreparer(ctx context.Context, billingAccountName string, customerName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"customerName":       autorest.Encode("path", customerName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CustomersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CustomersClient) GetResponder(resp *http.Response) (result Customer, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByBillingAccount lists the customers that are billed to a billing account. The operation is supported only for
// billing accounts with agreement type Microsoft Partner Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// search - used for searching customers by their name. Any customer with name containing the search text will
// be included in the response
// filter - may be used to filter the list of customers.
func (client CustomersClient) ListByBillingAccount(ctx context.Context, billingAccountName string, search string, filter string) (result CustomerListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomersClient.ListByBillingAccount")
		defer func() {
			sc := -1
			if result.clr.Response.Response != nil {
				sc = result.clr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByBillingAccountNextResults
	req, err := client.ListByBillingAccountPreparer(ctx, billingAccountName, search, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.CustomersClient", "ListByBillingAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingAccountSender(req)
	if err != nil {
		result.clr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.CustomersClient", "ListByBillingAccount", resp, "Failure sending request")
		return
	}

	result.clr, err = client.ListByBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.CustomersClient", "ListByBillingAccount", resp, "Failure responding to request")
	}

	return
}

// ListByBillingAccountPreparer prepares the ListByBillingAccount request.
func (client CustomersClient) ListByBillingAccountPreparer(ctx context.Context, billingAccountName string, search string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(search) > 0 {
		queryParameters["$search"] = autorest.Encode("query", search)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingAccountSender sends the ListByBillingAccount request. The method will close the
// http.Response Body if it receives an error.
func (client CustomersClient) ListByBillingAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingAccountResponder handles the response to the ListByBillingAccount request. The method always
// closes the http.Response Body.
func (client CustomersClient) ListByBillingAccountResponder(resp *http.Response) (result CustomerListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingAccountNextResults retrieves the next set of results, if any.
func (client CustomersClient) listByBillingAccountNextResults(ctx context.Context, lastResults CustomerListResult) (result CustomerListResult, err error) {
	req, err := lastResults.customerListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.CustomersClient", "listByBillingAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.CustomersClient", "listByBillingAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.CustomersClient", "listByBillingAccountNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client CustomersClient) ListByBillingAccountComplete(ctx context.Context, billingAccountName string, search string, filter string) (result CustomerListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomersClient.ListByBillingAccount")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingAccount(ctx, billingAccountName, search, filter)
	return
}

// ListByBillingProfile lists the customers that are billed to a billing profile. The operation is supported only for
// billing accounts with agreement type Microsoft Partner Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
// search - used for searching customers by their name. Any customer with name containing the search text will
// be included in the response
// filter - may be used to filter the list of customers.
func (client CustomersClient) ListByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string, search string, filter string) (result CustomerListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomersClient.ListByBillingProfile")
		defer func() {
			sc := -1
			if result.clr.Response.Response != nil {
				sc = result.clr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByBillingProfileNextResults
	req, err := client.ListByBillingProfilePreparer(ctx, billingAccountName, billingProfileName, search, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.CustomersClient", "ListByBillingProfile", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingProfileSender(req)
	if err != nil {
		result.clr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.CustomersClient", "ListByBillingProfile", resp, "Failure sending request")
		return
	}

	result.clr, err = client.ListByBillingProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.CustomersClient", "ListByBillingProfile", resp, "Failure responding to request")
	}

	return
}

// ListByBillingProfilePreparer prepares the ListByBillingProfile request.
func (client CustomersClient) ListByBillingProfilePreparer(ctx context.Context, billingAccountName string, billingProfileName string, search string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(search) > 0 {
		queryParameters["$search"] = autorest.Encode("query", search)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/customers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingProfileSender sends the ListByBillingProfile request. The method will close the
// http.Response Body if it receives an error.
func (client CustomersClient) ListByBillingProfileSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingProfileResponder handles the response to the ListByBillingProfile request. The method always
// closes the http.Response Body.
func (client CustomersClient) ListByBillingProfileResponder(resp *http.Response) (result CustomerListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingProfileNextResults retrieves the next set of results, if any.
func (client CustomersClient) listByBillingProfileNextResults(ctx context.Context, lastResults CustomerListResult) (result CustomerListResult, err error) {
	req, err := lastResults.customerListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.CustomersClient", "listByBillingProfileNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingProfileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.CustomersClient", "listByBillingProfileNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.CustomersClient", "listByBillingProfileNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingProfileComplete enumerates all values, automatically crossing page boundaries as required.
func (client CustomersClient) ListByBillingProfileComplete(ctx context.Context, billingAccountName string, billingProfileName string, search string, filter string) (result CustomerListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomersClient.ListByBillingProfile")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingProfile(ctx, billingAccountName, billingProfileName, search, filter)
	return
}
