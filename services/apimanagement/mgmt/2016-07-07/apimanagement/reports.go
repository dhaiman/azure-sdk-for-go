package apimanagement

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
    "github.com/Azure/go-autorest/autorest/validation"
)

// ReportsClient is the apiManagement Client
type ReportsClient struct {
    BaseClient
}
// NewReportsClient creates an instance of the ReportsClient client.
func NewReportsClient(subscriptionID string) ReportsClient {
    return NewReportsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewReportsClientWithBaseURI creates an instance of the ReportsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewReportsClientWithBaseURI(baseURI string, subscriptionID string) ReportsClient {
        return ReportsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// ListByService lists report records.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // aggregation - report aggregation.
        // filter - the filter to apply on the operation.
        // top - number of records to return.
        // skip - number of records to skip.
        // interval - by time interval. This value is only applicable to ByTime aggregation. Interval must be multiple
        // of 15 minutes and may not be zero. The value should be in ISO  8601 format
        // (http://en.wikipedia.org/wiki/ISO_8601#Durations).This code can be used to convert TimSpan to a valid
        // interval string: XmlConvert.ToString(new TimeSpan(hours, minutes, seconds))
func (client ReportsClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string, aggregation ReportsAggregation, filter string, top *int32, skip *int32, interval string) (result ReportCollectionPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ReportsClient.ListByService")
        defer func() {
            sc := -1
            if result.rc.Response.Response != nil {
                sc = result.rc.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: serviceName,
             Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}},
            { TargetValue: top,
             Constraints: []validation.Constraint{	{Target: "top", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil },
            }}}},
            { TargetValue: skip,
             Constraints: []validation.Constraint{	{Target: "skip", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil },
            }}}}}); err != nil {
            return result, validation.NewError("apimanagement.ReportsClient", "ListByService", err.Error())
            }

                        result.fn = client.listByServiceNextResults
    req, err := client.ListByServicePreparer(ctx, resourceGroupName, serviceName, aggregation, filter, top, skip, interval)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByService", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByServiceSender(req)
            if err != nil {
            result.rc.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByService", resp, "Failure sending request")
            return
            }

            result.rc, err = client.ListByServiceResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "ListByService", resp, "Failure responding to request")
            }

    return
    }

    // ListByServicePreparer prepares the ListByService request.
    func (client ReportsClient) ListByServicePreparer(ctx context.Context, resourceGroupName string, serviceName string, aggregation ReportsAggregation, filter string, top *int32, skip *int32, interval string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "aggregation": autorest.Encode("path",aggregation),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-07-07"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }
            if len(filter) > 0 {
            queryParameters["$filter"] = autorest.Encode("query",filter)
            }
            if top != nil {
            queryParameters["$top"] = autorest.Encode("query",*top)
            }
            if skip != nil {
            queryParameters["$skip"] = autorest.Encode("query",*skip)
            }
            if len(interval) > 0 {
            queryParameters["interval"] = autorest.Encode("query",interval)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/reports/{aggregation}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByServiceSender sends the ListByService request. The method will close the
    // http.Response Body if it receives an error.
    func (client ReportsClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client ReportsClient) ListByServiceResponder(resp *http.Response) (result ReportCollection, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listByServiceNextResults retrieves the next set of results, if any.
            func (client ReportsClient) listByServiceNextResults(ctx context.Context, lastResults ReportCollection) (result ReportCollection, err error) {
            req, err := lastResults.reportCollectionPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listByServiceNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByServiceSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listByServiceNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByServiceResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.ReportsClient", "listByServiceNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByServiceComplete enumerates all values, automatically crossing page boundaries as required.
    func (client ReportsClient) ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, aggregation ReportsAggregation, filter string, top *int32, skip *int32, interval string) (result ReportCollectionIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/ReportsClient.ListByService")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByService(ctx, resourceGroupName, serviceName, aggregation, filter, top, skip, interval)
                return
        }

