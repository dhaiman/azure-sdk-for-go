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
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
    "context"
    "github.com/Azure/go-autorest/tracing"
)

// InvoicesClient is the billing client provides access to billing resources for Azure subscriptions.
type InvoicesClient struct {
    BaseClient
}
// NewInvoicesClient creates an instance of the InvoicesClient client.
func NewInvoicesClient(subscriptionID string) InvoicesClient {
    return NewInvoicesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewInvoicesClientWithBaseURI creates an instance of the InvoicesClient client.
    func NewInvoicesClientWithBaseURI(baseURI string, subscriptionID string) InvoicesClient {
        return InvoicesClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Get get the invoice by name.
    // Parameters:
        // billingAccountName - billing Account Id.
        // billingProfileName - billing Profile Id.
        // invoiceName - invoice Id.
func (client InvoicesClient) Get(ctx context.Context, billingAccountName string, billingProfileName string, invoiceName string) (result InvoiceSummary, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/InvoicesClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, billingAccountName, billingProfileName, invoiceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "Get", nil , "Failure preparing request")
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
    func (client InvoicesClient) GetPreparer(ctx context.Context, billingAccountName string, billingProfileName string, invoiceName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "billingAccountName": autorest.Encode("path",billingAccountName),
            "billingProfileName": autorest.Encode("path",billingProfileName),
            "invoiceName": autorest.Encode("path",invoiceName),
            }

                        const APIVersion = "2018-11-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoices/{invoiceName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client InvoicesClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client InvoicesClient) GetResponder(resp *http.Response) (result InvoiceSummary, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByBillingAccountName list of invoices for a billing account.
    // Parameters:
        // billingAccountName - billing Account Id.
        // periodStartDate - invoice period start date.
        // periodEndDate - invoice period end date.
func (client InvoicesClient) ListByBillingAccountName(ctx context.Context, billingAccountName string, periodStartDate string, periodEndDate string) (result InvoiceListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/InvoicesClient.ListByBillingAccountName")
        defer func() {
            sc := -1
            if result.ilr.Response.Response != nil {
                sc = result.ilr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listByBillingAccountNameNextResults
    req, err := client.ListByBillingAccountNamePreparer(ctx, billingAccountName, periodStartDate, periodEndDate)
    if err != nil {
    err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "ListByBillingAccountName", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByBillingAccountNameSender(req)
            if err != nil {
            result.ilr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "ListByBillingAccountName", resp, "Failure sending request")
            return
            }

            result.ilr, err = client.ListByBillingAccountNameResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "ListByBillingAccountName", resp, "Failure responding to request")
            }

    return
    }

    // ListByBillingAccountNamePreparer prepares the ListByBillingAccountName request.
    func (client InvoicesClient) ListByBillingAccountNamePreparer(ctx context.Context, billingAccountName string, periodStartDate string, periodEndDate string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "billingAccountName": autorest.Encode("path",billingAccountName),
            }

                        const APIVersion = "2018-11-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        "periodEndDate": autorest.Encode("query",periodEndDate),
        "periodStartDate": autorest.Encode("query",periodStartDate),
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoices",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByBillingAccountNameSender sends the ListByBillingAccountName request. The method will close the
    // http.Response Body if it receives an error.
    func (client InvoicesClient) ListByBillingAccountNameSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByBillingAccountNameResponder handles the response to the ListByBillingAccountName request. The method always
// closes the http.Response Body.
func (client InvoicesClient) ListByBillingAccountNameResponder(resp *http.Response) (result InvoiceListResult, err error) {
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
            func (client InvoicesClient) listByBillingAccountNameNextResults(ctx context.Context, lastResults InvoiceListResult) (result InvoiceListResult, err error) {
            req, err := lastResults.invoiceListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "billing.InvoicesClient", "listByBillingAccountNameNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByBillingAccountNameSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "billing.InvoicesClient", "listByBillingAccountNameNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByBillingAccountNameResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "listByBillingAccountNameNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByBillingAccountNameComplete enumerates all values, automatically crossing page boundaries as required.
    func (client InvoicesClient) ListByBillingAccountNameComplete(ctx context.Context, billingAccountName string, periodStartDate string, periodEndDate string) (result InvoiceListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/InvoicesClient.ListByBillingAccountName")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByBillingAccountName(ctx, billingAccountName, periodStartDate, periodEndDate)
                return
        }

// listByBillingAccountNameNextResults retrieves the next set of results, if any.
func (client InvoicesClient) listByBillingAccountNameNextResults(ctx context.Context, lastResults InvoiceListResult) (result InvoiceListResult, err error) {
	req, err := lastResults.invoiceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.InvoicesClient", "listByBillingAccountNameNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingAccountNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.InvoicesClient", "listByBillingAccountNameNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingAccountNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "listByBillingAccountNameNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingAccountNameComplete enumerates all values, automatically crossing page boundaries as required.
func (client InvoicesClient) ListByBillingAccountNameComplete(ctx context.Context, billingAccountName string, periodStartDate string, periodEndDate string) (result InvoiceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvoicesClient.ListByBillingAccountName")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingAccountName(ctx, billingAccountName, periodStartDate, periodEndDate)
	return
}

// ListByBillingProfile list of invoices for a billing profile.
    // Parameters:
        // billingAccountName - billing Account Id.
        // billingProfileName - billing Profile Id.
        // periodStartDate - invoice period start date.
        // periodEndDate - invoice period end date.
func (client InvoicesClient) ListByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string, periodStartDate string, periodEndDate string) (result InvoiceListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/InvoicesClient.ListByBillingProfile")
        defer func() {
            sc := -1
            if result.ilr.Response.Response != nil {
                sc = result.ilr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listByBillingProfileNextResults
    req, err := client.ListByBillingProfilePreparer(ctx, billingAccountName, billingProfileName, periodStartDate, periodEndDate)
    if err != nil {
    err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "ListByBillingProfile", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByBillingProfileSender(req)
            if err != nil {
            result.ilr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "ListByBillingProfile", resp, "Failure sending request")
            return
            }

            result.ilr, err = client.ListByBillingProfileResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "ListByBillingProfile", resp, "Failure responding to request")
            }

    return
    }

    // ListByBillingProfilePreparer prepares the ListByBillingProfile request.
    func (client InvoicesClient) ListByBillingProfilePreparer(ctx context.Context, billingAccountName string, billingProfileName string, periodStartDate string, periodEndDate string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "billingAccountName": autorest.Encode("path",billingAccountName),
            "billingProfileName": autorest.Encode("path",billingProfileName),
            }

                        const APIVersion = "2018-11-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        "periodEndDate": autorest.Encode("query",periodEndDate),
        "periodStartDate": autorest.Encode("query",periodStartDate),
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoices",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByBillingProfileSender sends the ListByBillingProfile request. The method will close the
    // http.Response Body if it receives an error.
    func (client InvoicesClient) ListByBillingProfileSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByBillingProfileResponder handles the response to the ListByBillingProfile request. The method always
// closes the http.Response Body.
func (client InvoicesClient) ListByBillingProfileResponder(resp *http.Response) (result InvoiceListResult, err error) {
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
            func (client InvoicesClient) listByBillingProfileNextResults(ctx context.Context, lastResults InvoiceListResult) (result InvoiceListResult, err error) {
            req, err := lastResults.invoiceListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "billing.InvoicesClient", "listByBillingProfileNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByBillingProfileSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "billing.InvoicesClient", "listByBillingProfileNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByBillingProfileResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "billing.InvoicesClient", "listByBillingProfileNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByBillingProfileComplete enumerates all values, automatically crossing page boundaries as required.
    func (client InvoicesClient) ListByBillingProfileComplete(ctx context.Context, billingAccountName string, billingProfileName string, periodStartDate string, periodEndDate string) (result InvoiceListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/InvoicesClient.ListByBillingProfile")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByBillingProfile(ctx, billingAccountName, billingProfileName, periodStartDate, periodEndDate)
                return
        }

