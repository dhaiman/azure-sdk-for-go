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

// TransactionsClient is the billing client provides access to billing resources for Azure subscriptions.
type TransactionsClient struct {
    BaseClient
}
// NewTransactionsClient creates an instance of the TransactionsClient client.
func NewTransactionsClient(subscriptionID string) TransactionsClient {
    return NewTransactionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTransactionsClientWithBaseURI creates an instance of the TransactionsClient client.
    func NewTransactionsClientWithBaseURI(baseURI string, subscriptionID string) TransactionsClient {
        return TransactionsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// ListByBillingAccountName lists the transactions by billing account name for given start and end date.
    // Parameters:
        // billingAccountName - billing Account Id.
        // startDate - start date
        // endDate - end date
        // filter - may be used to filter by transaction kind. The filter supports 'eq', 'lt', 'gt', 'le', 'ge', and
        // 'and'. It does not currently support 'ne', 'or', or 'not'. Tag filter is a key value pair string where key
        // and value is separated by a colon (:).
func (client TransactionsClient) ListByBillingAccountName(ctx context.Context, billingAccountName string, startDate string, endDate string, filter string) (result TransactionsListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/TransactionsClient.ListByBillingAccountName")
        defer func() {
            sc := -1
            if result.tlr.Response.Response != nil {
                sc = result.tlr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listByBillingAccountNameNextResults
    req, err := client.ListByBillingAccountNamePreparer(ctx, billingAccountName, startDate, endDate, filter)
    if err != nil {
    err = autorest.NewErrorWithError(err, "billing.TransactionsClient", "ListByBillingAccountName", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByBillingAccountNameSender(req)
            if err != nil {
            result.tlr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "billing.TransactionsClient", "ListByBillingAccountName", resp, "Failure sending request")
            return
            }

            result.tlr, err = client.ListByBillingAccountNameResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "billing.TransactionsClient", "ListByBillingAccountName", resp, "Failure responding to request")
            }

    return
    }

    // ListByBillingAccountNamePreparer prepares the ListByBillingAccountName request.
    func (client TransactionsClient) ListByBillingAccountNamePreparer(ctx context.Context, billingAccountName string, startDate string, endDate string, filter string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "billingAccountName": autorest.Encode("path",billingAccountName),
            }

                        const APIVersion = "2018-11-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        "endDate": autorest.Encode("query",endDate),
        "startDate": autorest.Encode("query",startDate),
        }
            if len(filter) > 0 {
            queryParameters["$filter"] = autorest.Encode("query",filter)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/transactions",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByBillingAccountNameSender sends the ListByBillingAccountName request. The method will close the
    // http.Response Body if it receives an error.
    func (client TransactionsClient) ListByBillingAccountNameSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByBillingAccountNameResponder handles the response to the ListByBillingAccountName request. The method always
// closes the http.Response Body.
func (client TransactionsClient) ListByBillingAccountNameResponder(resp *http.Response) (result TransactionsListResult, err error) {
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
            func (client TransactionsClient) listByBillingAccountNameNextResults(ctx context.Context, lastResults TransactionsListResult) (result TransactionsListResult, err error) {
            req, err := lastResults.transactionsListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "billing.TransactionsClient", "listByBillingAccountNameNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByBillingAccountNameSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "billing.TransactionsClient", "listByBillingAccountNameNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByBillingAccountNameResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "billing.TransactionsClient", "listByBillingAccountNameNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByBillingAccountNameComplete enumerates all values, automatically crossing page boundaries as required.
    func (client TransactionsClient) ListByBillingAccountNameComplete(ctx context.Context, billingAccountName string, startDate string, endDate string, filter string) (result TransactionsListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/TransactionsClient.ListByBillingAccountName")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByBillingAccountName(ctx, billingAccountName, startDate, endDate, filter)
                return
        }

// ListByBillingProfileName lists the transactions by billing profile name for given start date and end date.
    // Parameters:
        // billingAccountName - billing Account Id.
        // billingProfileName - billing Profile Id.
        // startDate - start date
        // endDate - end date
        // filter - may be used to filter by transaction kind. The filter supports 'eq', 'lt', 'gt', 'le', 'ge', and
        // 'and'. It does not currently support 'ne', 'or', or 'not'. Tag filter is a key value pair string where key
        // and value is separated by a colon (:).
func (client TransactionsClient) ListByBillingProfileName(ctx context.Context, billingAccountName string, billingProfileName string, startDate string, endDate string, filter string) (result TransactionsListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/TransactionsClient.ListByBillingProfileName")
        defer func() {
            sc := -1
            if result.tlr.Response.Response != nil {
                sc = result.tlr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listByBillingProfileNameNextResults
    req, err := client.ListByBillingProfileNamePreparer(ctx, billingAccountName, billingProfileName, startDate, endDate, filter)
    if err != nil {
    err = autorest.NewErrorWithError(err, "billing.TransactionsClient", "ListByBillingProfileName", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByBillingProfileNameSender(req)
            if err != nil {
            result.tlr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "billing.TransactionsClient", "ListByBillingProfileName", resp, "Failure sending request")
            return
            }

            result.tlr, err = client.ListByBillingProfileNameResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "billing.TransactionsClient", "ListByBillingProfileName", resp, "Failure responding to request")
            }

    return
    }

    // ListByBillingProfileNamePreparer prepares the ListByBillingProfileName request.
    func (client TransactionsClient) ListByBillingProfileNamePreparer(ctx context.Context, billingAccountName string, billingProfileName string, startDate string, endDate string, filter string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "billingAccountName": autorest.Encode("path",billingAccountName),
            "billingProfileName": autorest.Encode("path",billingProfileName),
            }

                        const APIVersion = "2018-11-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        "endDate": autorest.Encode("query",endDate),
        "startDate": autorest.Encode("query",startDate),
        }
            if len(filter) > 0 {
            queryParameters["$filter"] = autorest.Encode("query",filter)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/transactions",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByBillingProfileNameSender sends the ListByBillingProfileName request. The method will close the
    // http.Response Body if it receives an error.
    func (client TransactionsClient) ListByBillingProfileNameSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByBillingProfileNameResponder handles the response to the ListByBillingProfileName request. The method always
// closes the http.Response Body.
func (client TransactionsClient) ListByBillingProfileNameResponder(resp *http.Response) (result TransactionsListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listByBillingProfileNameNextResults retrieves the next set of results, if any.
            func (client TransactionsClient) listByBillingProfileNameNextResults(ctx context.Context, lastResults TransactionsListResult) (result TransactionsListResult, err error) {
            req, err := lastResults.transactionsListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "billing.TransactionsClient", "listByBillingProfileNameNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByBillingProfileNameSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "billing.TransactionsClient", "listByBillingProfileNameNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByBillingProfileNameResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "billing.TransactionsClient", "listByBillingProfileNameNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByBillingProfileNameComplete enumerates all values, automatically crossing page boundaries as required.
    func (client TransactionsClient) ListByBillingProfileNameComplete(ctx context.Context, billingAccountName string, billingProfileName string, startDate string, endDate string, filter string) (result TransactionsListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/TransactionsClient.ListByBillingProfileName")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByBillingProfileName(ctx, billingAccountName, billingProfileName, startDate, endDate, filter)
                return
        }

// ListByCustomerName lists the transactions by invoice section name for given start date and end date.
    // Parameters:
        // billingAccountName - billing Account Id.
        // customerName - customer Id.
        // startDate - start date
        // endDate - end date
        // filter - may be used to filter by transaction kind. The filter supports 'eq', 'lt', 'gt', 'le', 'ge', and
        // 'and'. It does not currently support 'ne', 'or', or 'not'. Tag filter is a key value pair string where key
        // and value is separated by a colon (:).
func (client TransactionsClient) ListByCustomerName(ctx context.Context, billingAccountName string, customerName string, startDate string, endDate string, filter string) (result TransactionsListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/TransactionsClient.ListByCustomerName")
        defer func() {
            sc := -1
            if result.tlr.Response.Response != nil {
                sc = result.tlr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listByCustomerNameNextResults
    req, err := client.ListByCustomerNamePreparer(ctx, billingAccountName, customerName, startDate, endDate, filter)
    if err != nil {
    err = autorest.NewErrorWithError(err, "billing.TransactionsClient", "ListByCustomerName", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByCustomerNameSender(req)
            if err != nil {
            result.tlr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "billing.TransactionsClient", "ListByCustomerName", resp, "Failure sending request")
            return
            }

            result.tlr, err = client.ListByCustomerNameResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "billing.TransactionsClient", "ListByCustomerName", resp, "Failure responding to request")
            }

    return
    }

    // ListByCustomerNamePreparer prepares the ListByCustomerName request.
    func (client TransactionsClient) ListByCustomerNamePreparer(ctx context.Context, billingAccountName string, customerName string, startDate string, endDate string, filter string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "billingAccountName": autorest.Encode("path",billingAccountName),
            "customerName": autorest.Encode("path",customerName),
            }

                        const APIVersion = "2018-11-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        "endDate": autorest.Encode("query",endDate),
        "startDate": autorest.Encode("query",startDate),
        }
            if len(filter) > 0 {
            queryParameters["$filter"] = autorest.Encode("query",filter)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}/transactions",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByCustomerNameSender sends the ListByCustomerName request. The method will close the
    // http.Response Body if it receives an error.
    func (client TransactionsClient) ListByCustomerNameSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByCustomerNameResponder handles the response to the ListByCustomerName request. The method always
// closes the http.Response Body.
func (client TransactionsClient) ListByCustomerNameResponder(resp *http.Response) (result TransactionsListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listByCustomerNameNextResults retrieves the next set of results, if any.
            func (client TransactionsClient) listByCustomerNameNextResults(ctx context.Context, lastResults TransactionsListResult) (result TransactionsListResult, err error) {
            req, err := lastResults.transactionsListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "billing.TransactionsClient", "listByCustomerNameNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByCustomerNameSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "billing.TransactionsClient", "listByCustomerNameNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByCustomerNameResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "billing.TransactionsClient", "listByCustomerNameNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByCustomerNameComplete enumerates all values, automatically crossing page boundaries as required.
    func (client TransactionsClient) ListByCustomerNameComplete(ctx context.Context, billingAccountName string, customerName string, startDate string, endDate string, filter string) (result TransactionsListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/TransactionsClient.ListByCustomerName")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByCustomerName(ctx, billingAccountName, customerName, startDate, endDate, filter)
                return
        }

// listByBillingProfileNameNextResults retrieves the next set of results, if any.
func (client TransactionsClient) listByBillingProfileNameNextResults(ctx context.Context, lastResults TransactionsListResult) (result TransactionsListResult, err error) {
	req, err := lastResults.transactionsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.TransactionsClient", "listByBillingProfileNameNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByBillingProfileNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.TransactionsClient", "listByBillingProfileNameNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByBillingProfileNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.TransactionsClient", "listByBillingProfileNameNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByBillingProfileNameComplete enumerates all values, automatically crossing page boundaries as required.
func (client TransactionsClient) ListByBillingProfileNameComplete(ctx context.Context, billingAccountName string, billingProfileName string, startDate string, endDate string, filter string) (result TransactionsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TransactionsClient.ListByBillingProfileName")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByBillingProfileName(ctx, billingAccountName, billingProfileName, startDate, endDate, filter)
	return
}

// ListByCustomerName lists the transactions by invoice section name for given start date and end date.
// Parameters:
// billingAccountName - billing Account Id.
// customerName - customer Id.
// startDate - start date
// endDate - end date
// filter - may be used to filter by transaction kind. The filter supports 'eq', 'lt', 'gt', 'le', 'ge', and
// 'and'. It does not currently support 'ne', 'or', or 'not'. Tag filter is a key value pair string where key
// and value is separated by a colon (:).
func (client TransactionsClient) ListByCustomerName(ctx context.Context, billingAccountName string, customerName string, startDate string, endDate string, filter string) (result TransactionsListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TransactionsClient.ListByCustomerName")
		defer func() {
			sc := -1
			if result.tlr.Response.Response != nil {
				sc = result.tlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByCustomerNameNextResults
	req, err := client.ListByCustomerNamePreparer(ctx, billingAccountName, customerName, startDate, endDate, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.TransactionsClient", "ListByCustomerName", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByCustomerNameSender(req)
	if err != nil {
		result.tlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.TransactionsClient", "ListByCustomerName", resp, "Failure sending request")
		return
	}

	result.tlr, err = client.ListByCustomerNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.TransactionsClient", "ListByCustomerName", resp, "Failure responding to request")
	}

	return
}

// ListByCustomerNamePreparer prepares the ListByCustomerName request.
func (client TransactionsClient) ListByCustomerNamePreparer(ctx context.Context, billingAccountName string, customerName string, startDate string, endDate string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
		"customerName":       autorest.Encode("path", customerName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"endDate":     autorest.Encode("query", endDate),
		"startDate":   autorest.Encode("query", startDate),
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}/transactions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByCustomerNameSender sends the ListByCustomerName request. The method will close the
// http.Response Body if it receives an error.
func (client TransactionsClient) ListByCustomerNameSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByCustomerNameResponder handles the response to the ListByCustomerName request. The method always
// closes the http.Response Body.
func (client TransactionsClient) ListByCustomerNameResponder(resp *http.Response) (result TransactionsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByCustomerNameNextResults retrieves the next set of results, if any.
func (client TransactionsClient) listByCustomerNameNextResults(ctx context.Context, lastResults TransactionsListResult) (result TransactionsListResult, err error) {
	req, err := lastResults.transactionsListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "billing.TransactionsClient", "listByCustomerNameNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByCustomerNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "billing.TransactionsClient", "listByCustomerNameNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByCustomerNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.TransactionsClient", "listByCustomerNameNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByCustomerNameComplete enumerates all values, automatically crossing page boundaries as required.
func (client TransactionsClient) ListByCustomerNameComplete(ctx context.Context, billingAccountName string, customerName string, startDate string, endDate string, filter string) (result TransactionsListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TransactionsClient.ListByCustomerName")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByCustomerName(ctx, billingAccountName, customerName, startDate, endDate, filter)
	return
}

// ListByInvoiceSectionName lists the transactions by invoice section name for given start date and end date.
    // Parameters:
        // billingAccountName - billing Account Id.
        // invoiceSectionName - invoiceSection Id.
        // startDate - start date
        // endDate - end date
        // filter - may be used to filter by transaction kind. The filter supports 'eq', 'lt', 'gt', 'le', 'ge', and
        // 'and'. It does not currently support 'ne', 'or', or 'not'. Tag filter is a key value pair string where key
        // and value is separated by a colon (:).
func (client TransactionsClient) ListByInvoiceSectionName(ctx context.Context, billingAccountName string, invoiceSectionName string, startDate string, endDate string, filter string) (result TransactionsListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/TransactionsClient.ListByInvoiceSectionName")
        defer func() {
            sc := -1
            if result.tlr.Response.Response != nil {
                sc = result.tlr.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listByInvoiceSectionNameNextResults
    req, err := client.ListByInvoiceSectionNamePreparer(ctx, billingAccountName, invoiceSectionName, startDate, endDate, filter)
    if err != nil {
    err = autorest.NewErrorWithError(err, "billing.TransactionsClient", "ListByInvoiceSectionName", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByInvoiceSectionNameSender(req)
            if err != nil {
            result.tlr.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "billing.TransactionsClient", "ListByInvoiceSectionName", resp, "Failure sending request")
            return
            }

            result.tlr, err = client.ListByInvoiceSectionNameResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "billing.TransactionsClient", "ListByInvoiceSectionName", resp, "Failure responding to request")
            }

    return
    }

    // ListByInvoiceSectionNamePreparer prepares the ListByInvoiceSectionName request.
    func (client TransactionsClient) ListByInvoiceSectionNamePreparer(ctx context.Context, billingAccountName string, invoiceSectionName string, startDate string, endDate string, filter string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "billingAccountName": autorest.Encode("path",billingAccountName),
            "invoiceSectionName": autorest.Encode("path",invoiceSectionName),
            }

                        const APIVersion = "2018-11-01-preview"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        "endDate": autorest.Encode("query",endDate),
        "startDate": autorest.Encode("query",startDate),
        }
            if len(filter) > 0 {
            queryParameters["$filter"] = autorest.Encode("query",filter)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/invoiceSections/{invoiceSectionName}/transactions",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByInvoiceSectionNameSender sends the ListByInvoiceSectionName request. The method will close the
    // http.Response Body if it receives an error.
    func (client TransactionsClient) ListByInvoiceSectionNameSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListByInvoiceSectionNameResponder handles the response to the ListByInvoiceSectionName request. The method always
// closes the http.Response Body.
func (client TransactionsClient) ListByInvoiceSectionNameResponder(resp *http.Response) (result TransactionsListResult, err error) {
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
            func (client TransactionsClient) listByInvoiceSectionNameNextResults(ctx context.Context, lastResults TransactionsListResult) (result TransactionsListResult, err error) {
            req, err := lastResults.transactionsListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "billing.TransactionsClient", "listByInvoiceSectionNameNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByInvoiceSectionNameSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "billing.TransactionsClient", "listByInvoiceSectionNameNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByInvoiceSectionNameResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "billing.TransactionsClient", "listByInvoiceSectionNameNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByInvoiceSectionNameComplete enumerates all values, automatically crossing page boundaries as required.
    func (client TransactionsClient) ListByInvoiceSectionNameComplete(ctx context.Context, billingAccountName string, invoiceSectionName string, startDate string, endDate string, filter string) (result TransactionsListResultIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/TransactionsClient.ListByInvoiceSectionName")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByInvoiceSectionName(ctx, billingAccountName, invoiceSectionName, startDate, endDate, filter)
                return
        }

