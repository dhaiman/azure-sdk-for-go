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

// APIProductClient is the apiManagement Client
type APIProductClient struct {
    BaseClient
}
// NewAPIProductClient creates an instance of the APIProductClient client.
func NewAPIProductClient(subscriptionID string) APIProductClient {
    return NewAPIProductClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAPIProductClientWithBaseURI creates an instance of the APIProductClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewAPIProductClientWithBaseURI(baseURI string, subscriptionID string) APIProductClient {
        return APIProductClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// ListByApis lists all Products, which the API is part of.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // apiid - API identifier. Must be unique in the current API Management service instance.
        // filter - | Field | Supported operators    | Supported functions                         |
        // |-------|------------------------|---------------------------------------------|
        // | name  | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
        // top - number of records to return.
        // skip - number of records to skip.
func (client APIProductClient) ListByApis(ctx context.Context, resourceGroupName string, serviceName string, apiid string, filter string, top *int32, skip *int32) (result ProductCollectionPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/APIProductClient.ListByApis")
        defer func() {
            sc := -1
            if result.pc.Response.Response != nil {
                sc = result.pc.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: serviceName,
             Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}},
            { TargetValue: apiid,
             Constraints: []validation.Constraint{	{Target: "apiid", Name: validation.MaxLength, Rule: 80, Chain: nil },
            	{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "apiid", Name: validation.Pattern, Rule: `(^[\w]+$)|(^[\w][\w\-]+[\w]$)`, Chain: nil }}},
            { TargetValue: top,
             Constraints: []validation.Constraint{	{Target: "top", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil },
            }}}},
            { TargetValue: skip,
             Constraints: []validation.Constraint{	{Target: "skip", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil },
            }}}}}); err != nil {
            return result, validation.NewError("apimanagement.APIProductClient", "ListByApis", err.Error())
            }

                        result.fn = client.listByApisNextResults
    req, err := client.ListByApisPreparer(ctx, resourceGroupName, serviceName, apiid, filter, top, skip)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.APIProductClient", "ListByApis", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByApisSender(req)
            if err != nil {
            result.pc.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "apimanagement.APIProductClient", "ListByApis", resp, "Failure sending request")
            return
            }

            result.pc, err = client.ListByApisResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.APIProductClient", "ListByApis", resp, "Failure responding to request")
            }

    return
    }

    // ListByApisPreparer prepares the ListByApis request.
    func (client APIProductClient) ListByApisPreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, filter string, top *int32, skip *int32) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "apiId": autorest.Encode("path",apiid),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2018-01-01"
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

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/products",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByApisSender sends the ListByApis request. The method will close the
    // http.Response Body if it receives an error.
    func (client APIProductClient) ListByApisSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// ListByApisResponder handles the response to the ListByApis request. The method always
// closes the http.Response Body.
func (client APIProductClient) ListByApisResponder(resp *http.Response) (result ProductCollection, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listByApisNextResults retrieves the next set of results, if any.
            func (client APIProductClient) listByApisNextResults(ctx context.Context, lastResults ProductCollection) (result ProductCollection, err error) {
            req, err := lastResults.productCollectionPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "apimanagement.APIProductClient", "listByApisNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByApisSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "apimanagement.APIProductClient", "listByApisNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByApisResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.APIProductClient", "listByApisNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByApisComplete enumerates all values, automatically crossing page boundaries as required.
    func (client APIProductClient) ListByApisComplete(ctx context.Context, resourceGroupName string, serviceName string, apiid string, filter string, top *int32, skip *int32) (result ProductCollectionIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/APIProductClient.ListByApis")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByApis(ctx, resourceGroupName, serviceName, apiid, filter, top, skip)
                return
        }

