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

// InvoiceSectionsClient is the billing client provides access to billing resources for Azure subscriptions.
type InvoiceSectionsClient struct {
	BaseClient
}

// NewInvoiceSectionsClient creates an instance of the InvoiceSectionsClient client.
func NewInvoiceSectionsClient(subscriptionID string) InvoiceSectionsClient {
	return NewInvoiceSectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewInvoiceSectionsClientWithBaseURI creates an instance of the InvoiceSectionsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewInvoiceSectionsClientWithBaseURI(baseURI string, subscriptionID string) InvoiceSectionsClient {
	return InvoiceSectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create the operation to create a InvoiceSection.
// Parameters:
// billingAccountName - billing Account Id.
// parameters - parameters supplied to the Create InvoiceSection operation.
func (client InvoiceSectionsClient) Create(ctx context.Context, billingAccountName string, parameters InvoiceSectionCreationRequest) (result InvoiceSectionsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvoiceSectionsClient.Create")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, billingAccountName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client InvoiceSectionsClient) CreatePreparer(ctx context.Context, billingAccountName string, parameters InvoiceSectionCreationRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoiceSections", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client InvoiceSectionsClient) CreateSender(req *http.Request) (future InvoiceSectionsCreateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client InvoiceSectionsClient) CreateResponder(resp *http.Response) (result InvoiceSection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ElevateToBillingProfile elevates the caller's access to match their billing profile access.
// Parameters:
// billingAccountName - billing Account Id.
// invoiceSectionName - invoiceSection Id.
func (client InvoiceSectionsClient) ElevateToBillingProfile(ctx context.Context, billingAccountName string, invoiceSectionName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvoiceSectionsClient.ElevateToBillingProfile")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ElevateToBillingProfilePreparer(ctx, billingAccountName, invoiceSectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "ElevateToBillingProfile", nil, "Failure preparing request")
		return
	}

	resp, err := client.ElevateToBillingProfileSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "ElevateToBillingProfile", resp, "Failure sending request")
		return
	}

	result, err = client.ElevateToBillingProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "ElevateToBillingProfile", resp, "Failure responding to request")
	}

	return
}

// ElevateToBillingProfilePreparer prepares the ElevateToBillingProfile request.
func (client InvoiceSectionsClient) ElevateToBillingProfilePreparer(ctx context.Context, billingAccountName string, invoiceSectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"invoiceSectionName": autorest.Encode("path", invoiceSectionName),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoiceSections/{invoiceSectionName}/elevate", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ElevateToBillingProfileSender sends the ElevateToBillingProfile request. The method will close the
// http.Response Body if it receives an error.
func (client InvoiceSectionsClient) ElevateToBillingProfileSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ElevateToBillingProfileResponder handles the response to the ElevateToBillingProfile request. The method always
// closes the http.Response Body.
func (client InvoiceSectionsClient) ElevateToBillingProfileResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the InvoiceSection by id.
// Parameters:
// billingAccountName - billing Account Id.
// invoiceSectionName - invoiceSection Id.
// expand - may be used to expand the billingProfiles.
func (client InvoiceSectionsClient) Get(ctx context.Context, billingAccountName string, invoiceSectionName string, expand string) (result InvoiceSection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvoiceSectionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, billingAccountName, invoiceSectionName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client InvoiceSectionsClient) GetPreparer(ctx context.Context, billingAccountName string, invoiceSectionName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"invoiceSectionName": autorest.Encode("path", invoiceSectionName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoiceSections/{invoiceSectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client InvoiceSectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client InvoiceSectionsClient) GetResponder(resp *http.Response) (result InvoiceSection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByBillingAccountName lists all invoice sections for which a user has access.
// Parameters:
// billingAccountName - billing Account Id.
// expand - may be used to expand the billingProfiles.
func (client InvoiceSectionsClient) ListByBillingAccountName(ctx context.Context, billingAccountName string, expand string) (result InvoiceSectionListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvoiceSectionsClient.ListByBillingAccountName")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByBillingAccountNamePreparer(ctx, billingAccountName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "ListByBillingAccountName", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingAccountNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "ListByBillingAccountName", resp, "Failure sending request")
		return
	}

	result, err = client.ListByBillingAccountNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "ListByBillingAccountName", resp, "Failure responding to request")
	}

	return
}

// ListByBillingAccountNamePreparer prepares the ListByBillingAccountName request.
func (client InvoiceSectionsClient) ListByBillingAccountNamePreparer(ctx context.Context, billingAccountName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoiceSections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingAccountNameSender sends the ListByBillingAccountName request. The method will close the
// http.Response Body if it receives an error.
func (client InvoiceSectionsClient) ListByBillingAccountNameSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingAccountNameResponder handles the response to the ListByBillingAccountName request. The method always
// closes the http.Response Body.
func (client InvoiceSectionsClient) ListByBillingAccountNameResponder(resp *http.Response) (result InvoiceSectionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByBillingProfileName lists all invoice sections under a billing profile for which a user has access.
// Parameters:
// billingAccountName - billing Account Id.
// billingProfileName - billing Profile Id.
func (client InvoiceSectionsClient) ListByBillingProfileName(ctx context.Context, billingAccountName string, billingProfileName string) (result InvoiceSectionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvoiceSectionsClient.ListByBillingProfileName")
		defer func() {
			sc := -1
			if result.islr.Response.Response != nil {
				sc = result.islr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByBillingProfileNameNextResults
	req, err := client.ListByBillingProfileNamePreparer(ctx, billingAccountName, billingProfileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "ListByBillingProfileName", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingProfileNameSender(req)
	if err != nil {
		result.islr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "ListByBillingProfileName", resp, "Failure sending request")
		return
	}

	result.islr, err = client.ListByBillingProfileNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "ListByBillingProfileName", resp, "Failure responding to request")
	}
	if result.islr.hasNextLink() && result.islr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByBillingProfileNamePreparer prepares the ListByBillingProfileName request.
func (client InvoiceSectionsClient) ListByBillingProfileNamePreparer(ctx context.Context, billingAccountName string, billingProfileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingProfileNameSender sends the ListByBillingProfileName request. The method will close the
// http.Response Body if it receives an error.
func (client InvoiceSectionsClient) ListByBillingProfileNameSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingProfileNameResponder handles the response to the ListByBillingProfileName request. The method always
// closes the http.Response Body.
func (client InvoiceSectionsClient) ListByBillingProfileNameResponder(resp *http.Response) (result InvoiceSectionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingProfileNameNextResults retrieves the next set of results, if any.
func (client InvoiceSectionsClient) listByBillingProfileNameNextResults(ctx context.Context, lastResults InvoiceSectionListResult) (result InvoiceSectionListResult, err error) {
	req, err := lastResults.invoiceSectionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "listByBillingProfileNameNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingProfileNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "listByBillingProfileNameNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingProfileNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "listByBillingProfileNameNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingProfileNameComplete enumerates all values, automatically crossing page boundaries as required.
func (client InvoiceSectionsClient) ListByBillingProfileNameComplete(ctx context.Context, billingAccountName string, billingProfileName string) (result InvoiceSectionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvoiceSectionsClient.ListByBillingProfileName")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingProfileName(ctx, billingAccountName, billingProfileName)
	return
}

// ListByCreateSubscriptionPermission lists all invoiceSections with create subscription permission for a user.
// Parameters:
// billingAccountName - billing Account Id.
// expand - may be used to expand the billingProfiles.
func (client InvoiceSectionsClient) ListByCreateSubscriptionPermission(ctx context.Context, billingAccountName string, expand string) (result InvoiceSectionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvoiceSectionsClient.ListByCreateSubscriptionPermission")
		defer func() {
			sc := -1
			if result.islr.Response.Response != nil {
				sc = result.islr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByCreateSubscriptionPermissionNextResults
	req, err := client.ListByCreateSubscriptionPermissionPreparer(ctx, billingAccountName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "ListByCreateSubscriptionPermission", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByCreateSubscriptionPermissionSender(req)
	if err != nil {
		result.islr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "ListByCreateSubscriptionPermission", resp, "Failure sending request")
		return
	}

	result.islr, err = client.ListByCreateSubscriptionPermissionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "ListByCreateSubscriptionPermission", resp, "Failure responding to request")
	}
	if result.islr.hasNextLink() && result.islr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByCreateSubscriptionPermissionPreparer prepares the ListByCreateSubscriptionPermission request.
func (client InvoiceSectionsClient) ListByCreateSubscriptionPermissionPreparer(ctx context.Context, billingAccountName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/listInvoiceSectionsWithCreateSubscriptionPermission", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByCreateSubscriptionPermissionSender sends the ListByCreateSubscriptionPermission request. The method will close the
// http.Response Body if it receives an error.
func (client InvoiceSectionsClient) ListByCreateSubscriptionPermissionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByCreateSubscriptionPermissionResponder handles the response to the ListByCreateSubscriptionPermission request. The method always
// closes the http.Response Body.
func (client InvoiceSectionsClient) ListByCreateSubscriptionPermissionResponder(resp *http.Response) (result InvoiceSectionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByCreateSubscriptionPermissionNextResults retrieves the next set of results, if any.
func (client InvoiceSectionsClient) listByCreateSubscriptionPermissionNextResults(ctx context.Context, lastResults InvoiceSectionListResult) (result InvoiceSectionListResult, err error) {
	req, err := lastResults.invoiceSectionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "listByCreateSubscriptionPermissionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByCreateSubscriptionPermissionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "listByCreateSubscriptionPermissionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByCreateSubscriptionPermissionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "listByCreateSubscriptionPermissionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByCreateSubscriptionPermissionComplete enumerates all values, automatically crossing page boundaries as required.
func (client InvoiceSectionsClient) ListByCreateSubscriptionPermissionComplete(ctx context.Context, billingAccountName string, expand string) (result InvoiceSectionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvoiceSectionsClient.ListByCreateSubscriptionPermission")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByCreateSubscriptionPermission(ctx, billingAccountName, expand)
	return
}

// Update the operation to update a InvoiceSection.
// Parameters:
// billingAccountName - billing Account Id.
// invoiceSectionName - invoiceSection Id.
// parameters - parameters supplied to the Create InvoiceSection operation.
func (client InvoiceSectionsClient) Update(ctx context.Context, billingAccountName string, invoiceSectionName string, parameters InvoiceSection) (result InvoiceSectionsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvoiceSectionsClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, billingAccountName, invoiceSectionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client InvoiceSectionsClient) UpdatePreparer(ctx context.Context, billingAccountName string, invoiceSectionName string, parameters InvoiceSection) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"invoiceSectionName": autorest.Encode("path", invoiceSectionName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoiceSections/{invoiceSectionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client InvoiceSectionsClient) UpdateSender(req *http.Request) (future InvoiceSectionsUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client InvoiceSectionsClient) UpdateResponder(resp *http.Response) (result InvoiceSection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
