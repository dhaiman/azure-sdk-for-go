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

// RoleDefinitionsClient is the billing client provides access to billing resources for Azure subscriptions.
type RoleDefinitionsClient struct {
	BaseClient
}

// NewRoleDefinitionsClient creates an instance of the RoleDefinitionsClient client.
func NewRoleDefinitionsClient(subscriptionID string, subscriptionID1 string) RoleDefinitionsClient {
	return NewRoleDefinitionsClientWithBaseURI(DefaultBaseURI, subscriptionID, subscriptionID1)
}

// NewRoleDefinitionsClientWithBaseURI creates an instance of the RoleDefinitionsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRoleDefinitionsClientWithBaseURI(baseURI string, subscriptionID string, subscriptionID1 string) RoleDefinitionsClient {
	return RoleDefinitionsClient{NewWithBaseURI(baseURI, subscriptionID, subscriptionID1)}
}

// GetByBillingAccount gets the definition for a role on a billing account. The operation is supported for billing
// accounts with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingRoleDefinitionName - the ID that uniquely identifies a role definition.
func (client RoleDefinitionsClient) GetByBillingAccount(ctx context.Context, billingAccountName string, billingRoleDefinitionName string) (result RoleDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleDefinitionsClient.GetByBillingAccount")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByBillingAccountPreparer(ctx, billingAccountName, billingRoleDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "GetByBillingAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByBillingAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "GetByBillingAccount", resp, "Failure sending request")
		return
	}

	result, err = client.GetByBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "GetByBillingAccount", resp, "Failure responding to request")
	}

	return
}

// GetByBillingAccountPreparer prepares the GetByBillingAccount request.
func (client RoleDefinitionsClient) GetByBillingAccountPreparer(ctx context.Context, billingAccountName string, billingRoleDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName":        autorest.Encode("path", billingAccountName),
		"billingRoleDefinitionName": autorest.Encode("path", billingRoleDefinitionName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingRoleDefinitions/{billingRoleDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByBillingAccountSender sends the GetByBillingAccount request. The method will close the
// http.Response Body if it receives an error.
func (client RoleDefinitionsClient) GetByBillingAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetByBillingAccountResponder handles the response to the GetByBillingAccount request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) GetByBillingAccountResponder(resp *http.Response) (result RoleDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByBillingProfile gets the definition for a role on a billing profile. The operation is supported for billing
// accounts with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
// billingRoleDefinitionName - the ID that uniquely identifies a role definition.
func (client RoleDefinitionsClient) GetByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string, billingRoleDefinitionName string) (result RoleDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleDefinitionsClient.GetByBillingProfile")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByBillingProfilePreparer(ctx, billingAccountName, billingProfileName, billingRoleDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "GetByBillingProfile", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByBillingProfileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "GetByBillingProfile", resp, "Failure sending request")
		return
	}

	result, err = client.GetByBillingProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "GetByBillingProfile", resp, "Failure responding to request")
	}

	return
}

// GetByBillingProfilePreparer prepares the GetByBillingProfile request.
func (client RoleDefinitionsClient) GetByBillingProfilePreparer(ctx context.Context, billingAccountName string, billingProfileName string, billingRoleDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName":        autorest.Encode("path", billingAccountName),
		"billingProfileName":        autorest.Encode("path", billingProfileName),
		"billingRoleDefinitionName": autorest.Encode("path", billingRoleDefinitionName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/billingRoleDefinitions/{billingRoleDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByBillingProfileSender sends the GetByBillingProfile request. The method will close the
// http.Response Body if it receives an error.
func (client RoleDefinitionsClient) GetByBillingProfileSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetByBillingProfileResponder handles the response to the GetByBillingProfile request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) GetByBillingProfileResponder(resp *http.Response) (result RoleDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByInvoiceSection gets the definition for a role on an invoice section. The operation is supported only for
// billing accounts with agreement type Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
// invoiceSectionName - the ID that uniquely identifies an invoice section.
// billingRoleDefinitionName - the ID that uniquely identifies a role definition.
func (client RoleDefinitionsClient) GetByInvoiceSection(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, billingRoleDefinitionName string) (result RoleDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleDefinitionsClient.GetByInvoiceSection")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByInvoiceSectionPreparer(ctx, billingAccountName, billingProfileName, invoiceSectionName, billingRoleDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "GetByInvoiceSection", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByInvoiceSectionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "GetByInvoiceSection", resp, "Failure sending request")
		return
	}

	result, err = client.GetByInvoiceSectionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "GetByInvoiceSection", resp, "Failure responding to request")
	}

	return
}

// GetByInvoiceSectionPreparer prepares the GetByInvoiceSection request.
func (client RoleDefinitionsClient) GetByInvoiceSectionPreparer(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, billingRoleDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName":        autorest.Encode("path", billingAccountName),
		"billingProfileName":        autorest.Encode("path", billingProfileName),
		"billingRoleDefinitionName": autorest.Encode("path", billingRoleDefinitionName),
		"invoiceSectionName":        autorest.Encode("path", invoiceSectionName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/billingRoleDefinitions/{billingRoleDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByInvoiceSectionSender sends the GetByInvoiceSection request. The method will close the
// http.Response Body if it receives an error.
func (client RoleDefinitionsClient) GetByInvoiceSectionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetByInvoiceSectionResponder handles the response to the GetByInvoiceSection request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) GetByInvoiceSectionResponder(resp *http.Response) (result RoleDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByBillingAccount lists the role definitions for a billing account. The operation is supported for billing
// accounts with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
func (client RoleDefinitionsClient) ListByBillingAccount(ctx context.Context, billingAccountName string) (result RoleDefinitionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleDefinitionsClient.ListByBillingAccount")
		defer func() {
			sc := -1
			if result.rdlr.Response.Response != nil {
				sc = result.rdlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByBillingAccountNextResults
	req, err := client.ListByBillingAccountPreparer(ctx, billingAccountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "ListByBillingAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingAccountSender(req)
	if err != nil {
		result.rdlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "ListByBillingAccount", resp, "Failure sending request")
		return
	}

	result.rdlr, err = client.ListByBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "ListByBillingAccount", resp, "Failure responding to request")
	}
	if result.rdlr.hasNextLink() && result.rdlr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByBillingAccountPreparer prepares the ListByBillingAccount request.
func (client RoleDefinitionsClient) ListByBillingAccountPreparer(ctx context.Context, billingAccountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingRoleDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingAccountSender sends the ListByBillingAccount request. The method will close the
// http.Response Body if it receives an error.
func (client RoleDefinitionsClient) ListByBillingAccountSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingAccountResponder handles the response to the ListByBillingAccount request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) ListByBillingAccountResponder(resp *http.Response) (result RoleDefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingAccountNextResults retrieves the next set of results, if any.
func (client RoleDefinitionsClient) listByBillingAccountNextResults(ctx context.Context, lastResults RoleDefinitionListResult) (result RoleDefinitionListResult, err error) {
	req, err := lastResults.roleDefinitionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "listByBillingAccountNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "listByBillingAccountNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "listByBillingAccountNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingAccountComplete enumerates all values, automatically crossing page boundaries as required.
func (client RoleDefinitionsClient) ListByBillingAccountComplete(ctx context.Context, billingAccountName string) (result RoleDefinitionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleDefinitionsClient.ListByBillingAccount")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingAccount(ctx, billingAccountName)
	return
}

// ListByBillingProfile lists the role definitions for a billing profile. The operation is supported for billing
// accounts with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
func (client RoleDefinitionsClient) ListByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string) (result RoleDefinitionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleDefinitionsClient.ListByBillingProfile")
		defer func() {
			sc := -1
			if result.rdlr.Response.Response != nil {
				sc = result.rdlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByBillingProfileNextResults
	req, err := client.ListByBillingProfilePreparer(ctx, billingAccountName, billingProfileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "ListByBillingProfile", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByBillingProfileSender(req)
	if err != nil {
		result.rdlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "ListByBillingProfile", resp, "Failure sending request")
		return
	}

	result.rdlr, err = client.ListByBillingProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "ListByBillingProfile", resp, "Failure responding to request")
	}
	if result.rdlr.hasNextLink() && result.rdlr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByBillingProfilePreparer prepares the ListByBillingProfile request.
func (client RoleDefinitionsClient) ListByBillingProfilePreparer(ctx context.Context, billingAccountName string, billingProfileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/billingRoleDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByBillingProfileSender sends the ListByBillingProfile request. The method will close the
// http.Response Body if it receives an error.
func (client RoleDefinitionsClient) ListByBillingProfileSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByBillingProfileResponder handles the response to the ListByBillingProfile request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) ListByBillingProfileResponder(resp *http.Response) (result RoleDefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByBillingProfileNextResults retrieves the next set of results, if any.
func (client RoleDefinitionsClient) listByBillingProfileNextResults(ctx context.Context, lastResults RoleDefinitionListResult) (result RoleDefinitionListResult, err error) {
	req, err := lastResults.roleDefinitionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "listByBillingProfileNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingProfileSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "listByBillingProfileNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingProfileResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "listByBillingProfileNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingProfileComplete enumerates all values, automatically crossing page boundaries as required.
func (client RoleDefinitionsClient) ListByBillingProfileComplete(ctx context.Context, billingAccountName string, billingProfileName string) (result RoleDefinitionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleDefinitionsClient.ListByBillingProfile")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingProfile(ctx, billingAccountName, billingProfileName)
	return
}

// ListByInvoiceSection lists the role definitions for an invoice section. The operation is supported for billing
// accounts with agreement type Microsoft Partner Agreement or Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
// invoiceSectionName - the ID that uniquely identifies an invoice section.
func (client RoleDefinitionsClient) ListByInvoiceSection(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (result RoleDefinitionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleDefinitionsClient.ListByInvoiceSection")
		defer func() {
			sc := -1
			if result.rdlr.Response.Response != nil {
				sc = result.rdlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByInvoiceSectionNextResults
	req, err := client.ListByInvoiceSectionPreparer(ctx, billingAccountName, billingProfileName, invoiceSectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "ListByInvoiceSection", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByInvoiceSectionSender(req)
	if err != nil {
		result.rdlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "ListByInvoiceSection", resp, "Failure sending request")
		return
	}

	result.rdlr, err = client.ListByInvoiceSectionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "ListByInvoiceSection", resp, "Failure responding to request")
	}
	if result.rdlr.hasNextLink() && result.rdlr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByInvoiceSectionPreparer prepares the ListByInvoiceSection request.
func (client RoleDefinitionsClient) ListByInvoiceSectionPreparer(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
		"invoiceSectionName": autorest.Encode("path", invoiceSectionName),
	}

	const APIVersion = "2020-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/billingRoleDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByInvoiceSectionSender sends the ListByInvoiceSection request. The method will close the
// http.Response Body if it receives an error.
func (client RoleDefinitionsClient) ListByInvoiceSectionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByInvoiceSectionResponder handles the response to the ListByInvoiceSection request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) ListByInvoiceSectionResponder(resp *http.Response) (result RoleDefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByInvoiceSectionNextResults retrieves the next set of results, if any.
func (client RoleDefinitionsClient) listByInvoiceSectionNextResults(ctx context.Context, lastResults RoleDefinitionListResult) (result RoleDefinitionListResult, err error) {
	req, err := lastResults.roleDefinitionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "listByInvoiceSectionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByInvoiceSectionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "listByInvoiceSectionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByInvoiceSectionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.RoleDefinitionsClient", "listByInvoiceSectionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByInvoiceSectionComplete enumerates all values, automatically crossing page boundaries as required.
func (client RoleDefinitionsClient) ListByInvoiceSectionComplete(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (result RoleDefinitionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleDefinitionsClient.ListByInvoiceSection")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByInvoiceSection(ctx, billingAccountName, billingProfileName, invoiceSectionName)
	return
}
