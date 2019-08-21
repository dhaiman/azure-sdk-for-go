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

// NewProductsClientWithBaseURI creates an instance of the ProductsClient client.
func NewProductsClientWithBaseURI(baseURI string, subscriptionID string) ProductsClient {
	return ProductsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get a single product by name.
// Parameters:
// billingAccountName - billing Account Id.
// invoiceSectionName - invoiceSection Id.
// productName - invoice Id.
func (client ProductsClient) Get(ctx context.Context, billingAccountName string, invoiceSectionName string, productName string) (result ProductSummary, err error) {
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
	req, err := client.GetPreparer(ctx, billingAccountName, invoiceSectionName, productName)
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
func (client ProductsClient) GetPreparer(ctx context.Context, billingAccountName string, invoiceSectionName string, productName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"invoiceSectionName": autorest.Encode("path", invoiceSectionName),
		"productName":        autorest.Encode("path", productName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoiceSections/{invoiceSectionName}/products/{productName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProductsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProductsClient) GetResponder(resp *http.Response) (result ProductSummary, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByBillingAccountName lists products by billing account name.
// Parameters:
// billingAccountName - billing Account Id.
// filter - may be used to filter by product type. The filter supports 'eq', 'lt', 'gt', 'le', 'ge', and 'and'.
// It does not currently support 'ne', 'or', or 'not'. Tag filter is a key value pair string where key and
// value is separated by a colon (:).
func (client ProductsClient) ListByBillingAccountName(ctx context.Context, billingAccountName string, filter string) (result ProductsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.ListByBillingAccountName")
		defer func() {
			sc := -1
			if result.plr.Response.Response != nil {
				sc = result.plr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByBillingAccountNameNextResults
	req, err := client.ListByBillingAccountNamePreparer(ctx, billingAccountName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ListByBillingAccountName", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingAccountNameSender(req)
	if err != nil {
		result.plr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ListByBillingAccountName", resp, "Failure sending request")
		return
	}

	result.plr, err = client.ListByBillingAccountNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ListByBillingAccountName", resp, "Failure responding to request")
	}

	return
}

// ListByBillingAccountNamePreparer prepares the ListByBillingAccountName request.
func (client ProductsClient) ListByBillingAccountNamePreparer(ctx context.Context, billingAccountName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
	}

	const APIVersion = "2018-11-01-preview"
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

// ListByBillingAccountNameSender sends the ListByBillingAccountName request. The method will close the
// http.Response Body if it receives an error.
func (client ProductsClient) ListByBillingAccountNameSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByBillingAccountNameResponder handles the response to the ListByBillingAccountName request. The method always
// closes the http.Response Body.
func (client ProductsClient) ListByBillingAccountNameResponder(resp *http.Response) (result ProductsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingAccountNameNextResults retrieves the next set of results, if any.
func (client ProductsClient) listByBillingAccountNameNextResults(ctx context.Context, lastResults ProductsListResult) (result ProductsListResult, err error) {
	req, err := lastResults.productsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.ProductsClient", "listByBillingAccountNameNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingAccountNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.ProductsClient", "listByBillingAccountNameNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingAccountNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "listByBillingAccountNameNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingAccountNameComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProductsClient) ListByBillingAccountNameComplete(ctx context.Context, billingAccountName string, filter string) (result ProductsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.ListByBillingAccountName")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingAccountName(ctx, billingAccountName, filter)
	return
}

// ListByInvoiceSectionName lists products by invoice section name.
// Parameters:
// billingAccountName - billing Account Id.
// invoiceSectionName - invoiceSection Id.
// filter - may be used to filter by product type. The filter supports 'eq', 'lt', 'gt', 'le', 'ge', and 'and'.
// It does not currently support 'ne', 'or', or 'not'. Tag filter is a key value pair string where key and
// value is separated by a colon (:).
func (client ProductsClient) ListByInvoiceSectionName(ctx context.Context, billingAccountName string, invoiceSectionName string, filter string) (result ProductsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.ListByInvoiceSectionName")
		defer func() {
			sc := -1
			if result.plr.Response.Response != nil {
				sc = result.plr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByInvoiceSectionNameNextResults
	req, err := client.ListByInvoiceSectionNamePreparer(ctx, billingAccountName, invoiceSectionName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ListByInvoiceSectionName", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByInvoiceSectionNameSender(req)
	if err != nil {
		result.plr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ListByInvoiceSectionName", resp, "Failure sending request")
		return
	}

	result.plr, err = client.ListByInvoiceSectionNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ListByInvoiceSectionName", resp, "Failure responding to request")
	}

	return
}

// ListByInvoiceSectionNamePreparer prepares the ListByInvoiceSectionName request.
func (client ProductsClient) ListByInvoiceSectionNamePreparer(ctx context.Context, billingAccountName string, invoiceSectionName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"invoiceSectionName": autorest.Encode("path", invoiceSectionName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoiceSections/{invoiceSectionName}/products", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByInvoiceSectionNameSender sends the ListByInvoiceSectionName request. The method will close the
// http.Response Body if it receives an error.
func (client ProductsClient) ListByInvoiceSectionNameSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByInvoiceSectionNameResponder handles the response to the ListByInvoiceSectionName request. The method always
// closes the http.Response Body.
func (client ProductsClient) ListByInvoiceSectionNameResponder(resp *http.Response) (result ProductsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByInvoiceSectionNameNextResults retrieves the next set of results, if any.
func (client ProductsClient) listByInvoiceSectionNameNextResults(ctx context.Context, lastResults ProductsListResult) (result ProductsListResult, err error) {
	req, err := lastResults.productsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.ProductsClient", "listByInvoiceSectionNameNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByInvoiceSectionNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.ProductsClient", "listByInvoiceSectionNameNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByInvoiceSectionNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "listByInvoiceSectionNameNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByInvoiceSectionNameComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProductsClient) ListByInvoiceSectionNameComplete(ctx context.Context, billingAccountName string, invoiceSectionName string, filter string) (result ProductsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.ListByInvoiceSectionName")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByInvoiceSectionName(ctx, billingAccountName, invoiceSectionName, filter)
	return
}

// listByInvoiceSectionNameNextResults retrieves the next set of results, if any.
func (client ProductsClient) listByInvoiceSectionNameNextResults(ctx context.Context, lastResults ProductsListResult) (result ProductsListResult, err error) {
	req, err := lastResults.productsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.ProductsClient", "listByInvoiceSectionNameNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByInvoiceSectionNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.ProductsClient", "listByInvoiceSectionNameNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByInvoiceSectionNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "listByInvoiceSectionNameNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByInvoiceSectionNameComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProductsClient) ListByInvoiceSectionNameComplete(ctx context.Context, billingAccountName string, invoiceSectionName string, filter string) (result ProductsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.ListByInvoiceSectionName")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByInvoiceSectionName(ctx, billingAccountName, invoiceSectionName, filter)
	return
}

// listByInvoiceSectionNameNextResults retrieves the next set of results, if any.
func (client ProductsClient) listByInvoiceSectionNameNextResults(ctx context.Context, lastResults ProductsListResult) (result ProductsListResult, err error) {
	req, err := lastResults.productsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.ProductsClient", "listByInvoiceSectionNameNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByInvoiceSectionNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.ProductsClient", "listByInvoiceSectionNameNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByInvoiceSectionNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "listByInvoiceSectionNameNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByInvoiceSectionNameComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProductsClient) ListByInvoiceSectionNameComplete(ctx context.Context, billingAccountName string, invoiceSectionName string, filter string) (result ProductsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.ListByInvoiceSectionName")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByInvoiceSectionName(ctx, billingAccountName, invoiceSectionName, filter)
	return
}

// Transfer the operation to transfer a Product to another invoice section.
// Parameters:
// billingAccountName - billing Account Id.
// invoiceSectionName - invoiceSection Id.
// productName - invoice Id.
// parameters - parameters supplied to the Transfer Product operation.
func (client ProductsClient) Transfer(ctx context.Context, billingAccountName string, invoiceSectionName string, productName string, parameters TransferProductRequestProperties) (result ProductSummary, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.Transfer")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.TransferPreparer(ctx, billingAccountName, invoiceSectionName, productName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "Transfer", nil, "Failure preparing request")
		return
	}

	resp, err := client.TransferSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "Transfer", resp, "Failure sending request")
		return
	}

	result, err = client.TransferResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "Transfer", resp, "Failure responding to request")
	}

	return
}

// TransferPreparer prepares the Transfer request.
func (client ProductsClient) TransferPreparer(ctx context.Context, billingAccountName string, invoiceSectionName string, productName string, parameters TransferProductRequestProperties) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"invoiceSectionName": autorest.Encode("path", invoiceSectionName),
		"productName":        autorest.Encode("path", productName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoiceSections/{invoiceSectionName}/products/{productName}/transfer", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TransferSender sends the Transfer request. The method will close the
// http.Response Body if it receives an error.
func (client ProductsClient) TransferSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// TransferResponder handles the response to the Transfer request. The method always
// closes the http.Response Body.
func (client ProductsClient) TransferResponder(resp *http.Response) (result ProductSummary, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateAutoRenewByBillingAccountName cancel auto renew for product by product id and billing account name
// Parameters:
// billingAccountName - billing Account Id.
// productName - invoice Id.
// body - update auto renew request parameters.
func (client ProductsClient) UpdateAutoRenewByBillingAccountName(ctx context.Context, billingAccountName string, productName string, body UpdateAutoRenewRequest) (result UpdateAutoRenewOperationSummary, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.UpdateAutoRenewByBillingAccountName")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateAutoRenewByBillingAccountNamePreparer(ctx, billingAccountName, productName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "UpdateAutoRenewByBillingAccountName", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateAutoRenewByBillingAccountNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "UpdateAutoRenewByBillingAccountName", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateAutoRenewByBillingAccountNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "UpdateAutoRenewByBillingAccountName", resp, "Failure responding to request")
	}

	return
}

// UpdateAutoRenewByBillingAccountNamePreparer prepares the UpdateAutoRenewByBillingAccountName request.
func (client ProductsClient) UpdateAutoRenewByBillingAccountNamePreparer(ctx context.Context, billingAccountName string, productName string, body UpdateAutoRenewRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"productName":        autorest.Encode("path", productName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/products/{productName}/updateAutoRenew", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateAutoRenewByBillingAccountNameSender sends the UpdateAutoRenewByBillingAccountName request. The method will close the
// http.Response Body if it receives an error.
func (client ProductsClient) UpdateAutoRenewByBillingAccountNameSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// UpdateAutoRenewByBillingAccountNameResponder handles the response to the UpdateAutoRenewByBillingAccountName request. The method always
// closes the http.Response Body.
func (client ProductsClient) UpdateAutoRenewByBillingAccountNameResponder(resp *http.Response) (result UpdateAutoRenewOperationSummary, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateAutoRenewByInvoiceSectionName cancel auto renew for product by product id and invoice section name
// Parameters:
// billingAccountName - billing Account Id.
// invoiceSectionName - invoiceSection Id.
// productName - invoice Id.
// body - update auto renew request parameters.
func (client ProductsClient) UpdateAutoRenewByInvoiceSectionName(ctx context.Context, billingAccountName string, invoiceSectionName string, productName string, body UpdateAutoRenewRequest) (result UpdateAutoRenewOperationSummary, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.UpdateAutoRenewByInvoiceSectionName")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateAutoRenewByInvoiceSectionNamePreparer(ctx, billingAccountName, invoiceSectionName, productName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "UpdateAutoRenewByInvoiceSectionName", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateAutoRenewByInvoiceSectionNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "UpdateAutoRenewByInvoiceSectionName", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateAutoRenewByInvoiceSectionNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "UpdateAutoRenewByInvoiceSectionName", resp, "Failure responding to request")
	}

	return
}

// UpdateAutoRenewByInvoiceSectionNamePreparer prepares the UpdateAutoRenewByInvoiceSectionName request.
func (client ProductsClient) UpdateAutoRenewByInvoiceSectionNamePreparer(ctx context.Context, billingAccountName string, invoiceSectionName string, productName string, body UpdateAutoRenewRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"invoiceSectionName": autorest.Encode("path", invoiceSectionName),
		"productName":        autorest.Encode("path", productName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoiceSections/{invoiceSectionName}/products/{productName}/updateAutoRenew", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateAutoRenewByInvoiceSectionNameSender sends the UpdateAutoRenewByInvoiceSectionName request. The method will close the
// http.Response Body if it receives an error.
func (client ProductsClient) UpdateAutoRenewByInvoiceSectionNameSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// UpdateAutoRenewByInvoiceSectionNameResponder handles the response to the UpdateAutoRenewByInvoiceSectionName request. The method always
// closes the http.Response Body.
func (client ProductsClient) UpdateAutoRenewByInvoiceSectionNameResponder(resp *http.Response) (result UpdateAutoRenewOperationSummary, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ValidateTransfer validates the transfer of products across invoice sections.
// Parameters:
// billingAccountName - billing Account Id.
// invoiceSectionName - invoiceSection Id.
// productName - invoice Id.
// parameters - parameters supplied to the Transfer Products operation.
func (client ProductsClient) ValidateTransfer(ctx context.Context, billingAccountName string, invoiceSectionName string, productName string, parameters TransferProductRequestProperties) (result ValidateProductTransferEligibilityResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProductsClient.ValidateTransfer")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ValidateTransferPreparer(ctx, billingAccountName, invoiceSectionName, productName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ValidateTransfer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ValidateTransferSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ValidateTransfer", resp, "Failure sending request")
		return
	}

	result, err = client.ValidateTransferResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.ProductsClient", "ValidateTransfer", resp, "Failure responding to request")
	}

	return
}

// ValidateTransferPreparer prepares the ValidateTransfer request.
func (client ProductsClient) ValidateTransferPreparer(ctx context.Context, billingAccountName string, invoiceSectionName string, productName string, parameters TransferProductRequestProperties) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"invoiceSectionName": autorest.Encode("path", invoiceSectionName),
		"productName":        autorest.Encode("path", productName),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoiceSections/{invoiceSectionName}/products/{productName}/validateTransferEligibility", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ValidateTransferSender sends the ValidateTransfer request. The method will close the
// http.Response Body if it receives an error.
func (client ProductsClient) ValidateTransferSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ValidateTransferResponder handles the response to the ValidateTransfer request. The method always
// closes the http.Response Body.
func (client ProductsClient) ValidateTransferResponder(resp *http.Response) (result ValidateProductTransferEligibilityResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
