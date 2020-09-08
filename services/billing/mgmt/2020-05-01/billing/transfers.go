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

// TransfersClient is the billing client provides access to billing resources for Azure subscriptions.
type TransfersClient struct {
	BaseClient
}

// NewTransfersClient creates an instance of the TransfersClient client.
func NewTransfersClient(subscriptionID string, subscriptionID1 string) TransfersClient {
	return NewTransfersClientWithBaseURI(DefaultBaseURI, subscriptionID, subscriptionID1)
}

// NewTransfersClientWithBaseURI creates an instance of the TransfersClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewTransfersClientWithBaseURI(baseURI string, subscriptionID string, subscriptionID1 string) TransfersClient {
	return TransfersClient{NewWithBaseURI(baseURI, subscriptionID, subscriptionID1)}
}

// Cancel cancels a transfer request. The operation is supported only for billing accounts with agreement type
// Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
// invoiceSectionName - the ID that uniquely identifies an invoice section.
// transferName - the ID that uniquely identifies a transfer request.
func (client TransfersClient) Cancel(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, transferName string) (result TransferDetails, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TransfersClient.Cancel")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CancelPreparer(ctx, billingAccountName, billingProfileName, invoiceSectionName, transferName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.TransfersClient", "Cancel", nil, "Failure preparing request")
		return
	}

	resp, err := client.CancelSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.TransfersClient", "Cancel", resp, "Failure sending request")
		return
	}

	result, err = client.CancelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.TransfersClient", "Cancel", resp, "Failure responding to request")
	}

	return
}

// CancelPreparer prepares the Cancel request.
func (client TransfersClient) CancelPreparer(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, transferName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
		"invoiceSectionName": autorest.Encode("path", invoiceSectionName),
		"transferName":       autorest.Encode("path", transferName),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/transfers/{transferName}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client TransfersClient) CancelSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client TransfersClient) CancelResponder(resp *http.Response) (result TransferDetails, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a transfer request by ID. The operation is supported only for billing accounts with agreement type
// Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
// invoiceSectionName - the ID that uniquely identifies an invoice section.
// transferName - the ID that uniquely identifies a transfer request.
func (client TransfersClient) Get(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, transferName string) (result TransferDetails, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TransfersClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, billingAccountName, billingProfileName, invoiceSectionName, transferName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.TransfersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.TransfersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.TransfersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client TransfersClient) GetPreparer(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, transferName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
		"invoiceSectionName": autorest.Encode("path", invoiceSectionName),
		"transferName":       autorest.Encode("path", transferName),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/transfers/{transferName}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client TransfersClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client TransfersClient) GetResponder(resp *http.Response) (result TransferDetails, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Initiate sends a request to a user in another billing account to transfer billing ownership of their subscriptions.
// The operation is supported only for billing accounts with agreement type Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
// invoiceSectionName - the ID that uniquely identifies an invoice section.
// parameters - request parameters that are provided to the initiate transfer operation.
func (client TransfersClient) Initiate(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, parameters InitiateTransferRequest) (result TransferDetails, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TransfersClient.Initiate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.InitiatePreparer(ctx, billingAccountName, billingProfileName, invoiceSectionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.TransfersClient", "Initiate", nil, "Failure preparing request")
		return
	}

	resp, err := client.InitiateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.TransfersClient", "Initiate", resp, "Failure sending request")
		return
	}

	result, err = client.InitiateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.TransfersClient", "Initiate", resp, "Failure responding to request")
	}

	return
}

// InitiatePreparer prepares the Initiate request.
func (client TransfersClient) InitiatePreparer(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, parameters InitiateTransferRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
		"invoiceSectionName": autorest.Encode("path", invoiceSectionName),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/initiateTransfer", pathParameters),
		autorest.WithJSON(parameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// InitiateSender sends the Initiate request. The method will close the
// http.Response Body if it receives an error.
func (client TransfersClient) InitiateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// InitiateResponder handles the response to the Initiate request. The method always
// closes the http.Response Body.
func (client TransfersClient) InitiateResponder(resp *http.Response) (result TransferDetails, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the transfer requests for an invoice section. The operation is supported only for billing accounts with
// agreement type Microsoft Customer Agreement.
// Parameters:
// billingAccountName - the ID that uniquely identifies a billing account.
// billingProfileName - the ID that uniquely identifies a billing profile.
// invoiceSectionName - the ID that uniquely identifies an invoice section.
func (client TransfersClient) List(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (result TransferDetailsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TransfersClient.List")
		defer func() {
			sc := -1
			if result.tdlr.Response.Response != nil {
				sc = result.tdlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, billingAccountName, billingProfileName, invoiceSectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.TransfersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.tdlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.TransfersClient", "List", resp, "Failure sending request")
		return
	}

	result.tdlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.TransfersClient", "List", resp, "Failure responding to request")
	}
	if result.tdlr.hasNextLink() && result.tdlr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client TransfersClient) ListPreparer(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"billingProfileName": autorest.Encode("path", billingProfileName),
		"invoiceSectionName": autorest.Encode("path", invoiceSectionName),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/transfers", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client TransfersClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client TransfersClient) ListResponder(resp *http.Response) (result TransferDetailsListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client TransfersClient) listNextResults(ctx context.Context, lastResults TransferDetailsListResult) (result TransferDetailsListResult, err error) {
	req, err := lastResults.transferDetailsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.TransfersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.TransfersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.TransfersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client TransfersClient) ListComplete(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string) (result TransferDetailsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TransfersClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, billingAccountName, billingProfileName, invoiceSectionName)
	return
}
