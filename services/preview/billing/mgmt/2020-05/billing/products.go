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

// ProductsClient is the billing client provides access to billing resources for Azure subscriptions.
type ProductsClient struct {
	BaseClient
}

// NewProductsClient creates an instance of the ProductsClient client.
func NewProductsClient(subscriptionID string) ProductsClient {
	return NewProductsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProductsClientWithBaseURI creates an instance of the ProductsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewProductsClientWithBaseURI(baseURI string, subscriptionID string) ProductsClient {
	return ProductsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a product by ID. The operation is supported only for billing accounts with agreement type Microsoft
// Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// productName - the ID that uniquely identifies a product.
func (client ProductsClient) Get(ctx context.Context, billingAccountName string, productName string) (result Product, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, billingAccountName, productName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProductsClient) GetPreparer(ctx context.Context, billingAccountName string, productName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"productName":        autorest.Encode("path", productName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/products/{productName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProductsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProductsClient) GetResponder(resp *http.Response) (result Product, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByBillingAccount lists the products for a billing account. These don't include products billed based on usage.
// The operation is supported for billing accounts with agreement type Microsoft Customer Agreement or Microsoft
// Partner Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// filter - may be used to filter by product type. The filter supports 'eq', 'lt', 'gt', 'le', 'ge', and 'and'.
// It does not currently support 'ne', 'or', or 'not'. Tag filter is a key value pair string where key and
// value are separated by a colon (:).
func (client ProductsClient) ListByBillingAccount(ctx context.Context, billingAccountName string, filter string) (result ProductsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.ListByBillingAccount")
		defer func() {
			sc := -1
			if result.plr.Response.Response != nil {
				sc = result.plr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByBillingAccountNextResults
	req, err := client.ListByBillingAccountPreparer(ctx, billingAccountName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ListByBillingAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingAccountSender(req)
	if err != nil {
		result.plr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ListByBillingAccount", resp, "Failure sending request")
		return
	}

	result.plr, err = client.ListByBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ListByBillingAccount", resp, "Failure responding to request")
	}

	return
}

// ListByBillingAccountPreparer prepares the ListByBillingAccount request.
func (client ProductsClient) ListByBillingAccountPreparer(ctx context.Context, billingAccountName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/products", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingAccountSender sends the ListByBillingAccount request. The method will close the
// http.Response Body if it receives an error.
func (client ProductsClient) ListByBillingAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingAccountResponder handles the response to the ListByBillingAccount request. The method always
// closes the http.Response Body.
func (client ProductsClient) ListByBillingAccountResponder(resp *http.Response) (result ProductsListResult, err error) {
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
func (client ProductsClient) listByBillingAccountNextResults(ctx context.Context, lastResults ProductsListResult) (result ProductsListResult, err error) {
	req, err := lastResults.productsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.ProductsClient", "listByBillingAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.ProductsClient", "listByBillingAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "listByBillingAccountNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProductsClient) ListByBillingAccountComplete(ctx context.Context, billingAccountName string, filter string) (result ProductsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.ListByBillingAccount")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingAccount(ctx, billingAccountName, filter)
	return
}

// ListByBillingProfile lists the products for a billing profile. These don't include products billed based on usage.
// The operation is supported for billing accounts with agreement type Microsoft Customer Agreement or Microsoft
// Partner Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
// filter - may be used to filter by product type. The filter supports 'eq', 'lt', 'gt', 'le', 'ge', and 'and'.
// It does not currently support 'ne', 'or', or 'not'. Tag filter is a key value pair string where key and
// value are separated by a colon (:).
func (client ProductsClient) ListByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string, filter string) (result ProductsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.ListByBillingProfile")
		defer func() {
			sc := -1
			if result.plr.Response.Response != nil {
				sc = result.plr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByBillingProfileNextResults
	req, err := client.ListByBillingProfilePreparer(ctx, billingAccountName, billingProfileName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ListByBillingProfile", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingProfileSender(req)
	if err != nil {
		result.plr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ListByBillingProfile", resp, "Failure sending request")
		return
	}

	result.plr, err = client.ListByBillingProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ListByBillingProfile", resp, "Failure responding to request")
	}

	return
}

// ListByBillingProfilePreparer prepares the ListByBillingProfile request.
func (client ProductsClient) ListByBillingProfilePreparer(ctx context.Context, billingAccountName string, billingProfileName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/products", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingProfileSender sends the ListByBillingProfile request. The method will close the
// http.Response Body if it receives an error.
func (client ProductsClient) ListByBillingProfileSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingProfileResponder handles the response to the ListByBillingProfile request. The method always
// closes the http.Response Body.
func (client ProductsClient) ListByBillingProfileResponder(resp *http.Response) (result ProductsListResult, err error) {
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
func (client ProductsClient) listByBillingProfileNextResults(ctx context.Context, lastResults ProductsListResult) (result ProductsListResult, err error) {
	req, err := lastResults.productsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.ProductsClient", "listByBillingProfileNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingProfileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.ProductsClient", "listByBillingProfileNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "listByBillingProfileNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingProfileComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProductsClient) ListByBillingProfileComplete(ctx context.Context, billingAccountName string, billingProfileName string, filter string) (result ProductsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.ListByBillingProfile")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingProfile(ctx, billingAccountName, billingProfileName, filter)
	return
}

// ListByCustomer lists the products for a customer. These don't include products billed based on usage.The operation
// is supported only for billing accounts with agreement type Microsoft Partner Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// customerName - the ID that uniquely identifies a customer.
func (client ProductsClient) ListByCustomer(ctx context.Context, billingAccountName string, customerName string) (result ProductsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.ListByCustomer")
		defer func() {
			sc := -1
			if result.plr.Response.Response != nil {
				sc = result.plr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByCustomerNextResults
	req, err := client.ListByCustomerPreparer(ctx, billingAccountName, customerName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ListByCustomer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByCustomerSender(req)
	if err != nil {
		result.plr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ListByCustomer", resp, "Failure sending request")
		return
	}

	result.plr, err = client.ListByCustomerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ListByCustomer", resp, "Failure responding to request")
	}

	return
}

// ListByCustomerPreparer prepares the ListByCustomer request.
func (client ProductsClient) ListByCustomerPreparer(ctx context.Context, billingAccountName string, customerName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"customerName":       autorest.Encode("path", customerName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}/products", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByCustomerSender sends the ListByCustomer request. The method will close the
// http.Response Body if it receives an error.
func (client ProductsClient) ListByCustomerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByCustomerResponder handles the response to the ListByCustomer request. The method always
// closes the http.Response Body.
func (client ProductsClient) ListByCustomerResponder(resp *http.Response) (result ProductsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByCustomerNextResults retrieves the next set of results, if any.
func (client ProductsClient) listByCustomerNextResults(ctx context.Context, lastResults ProductsListResult) (result ProductsListResult, err error) {
	req, err := lastResults.productsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.ProductsClient", "listByCustomerNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByCustomerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.ProductsClient", "listByCustomerNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByCustomerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "listByCustomerNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByCustomerComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProductsClient) ListByCustomerComplete(ctx context.Context, billingAccountName string, customerName string) (result ProductsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.ListByCustomer")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByCustomer(ctx, billingAccountName, customerName)
	return
}

// ListByInvoiceSection lists the products for an invoice section. These don't include products billed based on usage.
// The operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
// invoiceSectionName - the ID that uniquely identifies an invoice section.
// filter - may be used to filter by product type. The filter supports 'eq', 'lt', 'gt', 'le', 'ge', and 'and'.
// It does not currently support 'ne', 'or', or 'not'. Tag filter is a key value pair string where key and
// value are separated by a colon (:).
func (client ProductsClient) ListByInvoiceSection(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, filter string) (result ProductsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.ListByInvoiceSection")
		defer func() {
			sc := -1
			if result.plr.Response.Response != nil {
				sc = result.plr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByInvoiceSectionNextResults
	req, err := client.ListByInvoiceSectionPreparer(ctx, billingAccountName, billingProfileName, invoiceSectionName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ListByInvoiceSection", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByInvoiceSectionSender(req)
	if err != nil {
		result.plr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ListByInvoiceSection", resp, "Failure sending request")
		return
	}

	result.plr, err = client.ListByInvoiceSectionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ListByInvoiceSection", resp, "Failure responding to request")
	}

	return
}

// ListByInvoiceSectionPreparer prepares the ListByInvoiceSection request.
func (client ProductsClient) ListByInvoiceSectionPreparer(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
		"invoiceSectionName": autorest.Encode("path", invoiceSectionName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/products", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByInvoiceSectionSender sends the ListByInvoiceSection request. The method will close the
// http.Response Body if it receives an error.
func (client ProductsClient) ListByInvoiceSectionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByInvoiceSectionResponder handles the response to the ListByInvoiceSection request. The method always
// closes the http.Response Body.
func (client ProductsClient) ListByInvoiceSectionResponder(resp *http.Response) (result ProductsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByInvoiceSectionNextResults retrieves the next set of results, if any.
func (client ProductsClient) listByInvoiceSectionNextResults(ctx context.Context, lastResults ProductsListResult) (result ProductsListResult, err error) {
	req, err := lastResults.productsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.ProductsClient", "listByInvoiceSectionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByInvoiceSectionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.ProductsClient", "listByInvoiceSectionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByInvoiceSectionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "listByInvoiceSectionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByInvoiceSectionComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProductsClient) ListByInvoiceSectionComplete(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, filter string) (result ProductsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.ListByInvoiceSection")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByInvoiceSection(ctx, billingAccountName, billingProfileName, invoiceSectionName, filter)
	return
}

// Move moves a product's charges to a new invoice section. The new invoice section must belong to the same billing
// profile as the existing invoice section. This operation is supported only for products that are purchased with a
// recurring charge and for billing accounts with agreement type Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// productName - the ID that uniquely identifies a product.
// parameters - request parameters that are provided to the transfer product operation.
func (client ProductsClient) Move(ctx context.Context, billingAccountName string, productName string, parameters TransferProductRequestProperties) (result Product, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.Move")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.MovePreparer(ctx, billingAccountName, productName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "Move", nil, "Failure preparing request")
		return
	}

	resp, err := client.MoveSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "Move", resp, "Failure sending request")
		return
	}

	result, err = client.MoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "Move", resp, "Failure responding to request")
	}

	return
}

// MovePreparer prepares the Move request.
func (client ProductsClient) MovePreparer(ctx context.Context, billingAccountName string, productName string, parameters TransferProductRequestProperties) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"productName":        autorest.Encode("path", productName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/products/{productName}/move", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// MoveSender sends the Move request. The method will close the
// http.Response Body if it receives an error.
func (client ProductsClient) MoveSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// MoveResponder handles the response to the Move request. The method always
// closes the http.Response Body.
func (client ProductsClient) MoveResponder(resp *http.Response) (result Product, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ValidateMove validates if a product's charges can be moved to a new invoice section. This operation is supported
// only for products that are purchased with a recurring charge and for billing accounts with agreement type Microsoft
// Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// productName - the ID that uniquely identifies a product.
// parameters - request parameters that are provided to the validate transfer eligibility operation.
func (client ProductsClient) ValidateMove(ctx context.Context, billingAccountName string, productName string, parameters TransferProductRequestProperties) (result ValidateProductTransferEligibilityResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.ValidateMove")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ValidateMovePreparer(ctx, billingAccountName, productName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ValidateMove", nil, "Failure preparing request")
		return
	}

	resp, err := client.ValidateMoveSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ValidateMove", resp, "Failure sending request")
		return
	}

	result, err = client.ValidateMoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ValidateMove", resp, "Failure responding to request")
	}

	return
}

// ValidateMovePreparer prepares the ValidateMove request.
func (client ProductsClient) ValidateMovePreparer(ctx context.Context, billingAccountName string, productName string, parameters TransferProductRequestProperties) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"productName":        autorest.Encode("path", productName),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/products/{productName}/validateMoveEligibility", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ValidateMoveSender sends the ValidateMove request. The method will close the
// http.Response Body if it receives an error.
func (client ProductsClient) ValidateMoveSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ValidateMoveResponder handles the response to the ValidateMove request. The method always
// closes the http.Response Body.
func (client ProductsClient) ValidateMoveResponder(resp *http.Response) (result ValidateProductTransferEligibilityResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
