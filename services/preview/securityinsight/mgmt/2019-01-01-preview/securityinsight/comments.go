package securityinsight

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

// CommentsClient is the API spec for Microsoft.SecurityInsights (Azure Security Insights) resource provider
type CommentsClient struct {
    BaseClient
}
// NewCommentsClient creates an instance of the CommentsClient client.
func NewCommentsClient(subscriptionID string) CommentsClient {
    return NewCommentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCommentsClientWithBaseURI creates an instance of the CommentsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewCommentsClientWithBaseURI(baseURI string, subscriptionID string) CommentsClient {
        return CommentsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// ListByCase gets all case comments.
    // Parameters:
        // resourceGroupName - the name of the resource group within the user's subscription. The name is case
        // insensitive.
        // operationalInsightsResourceProvider - the namespace of workspaces resource provider-
        // Microsoft.OperationalInsights.
        // workspaceName - the name of the workspace.
        // caseID - case ID
        // filter - filters the results, based on a Boolean condition. Optional.
        // orderby - sorts the results. Optional.
        // top - returns only the first n results. Optional.
        // skipToken - skiptoken is only used if a previous operation returned a partial result. If a previous response
        // contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
        // specifies a starting point to use for subsequent calls. Optional.
func (client CommentsClient) ListByCase(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, filter string, orderby string, top *int32, skipToken string) (result CaseCommentListPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CommentsClient.ListByCase")
        defer func() {
            sc := -1
        if result.ccl.Response.Response != nil {
        sc = result.ccl.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: client.SubscriptionID,
         Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil }}},
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
        { TargetValue: workspaceName,
         Constraints: []validation.Constraint{	{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil },
        	{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil }}}}); err != nil {
        return result, validation.NewError("securityinsight.CommentsClient", "ListByCase", err.Error())
        }

            result.fn = client.listByCaseNextResults
    req, err := client.ListByCasePreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, caseID, filter, orderby, top, skipToken)
    if err != nil {
    err = autorest.NewErrorWithError(err, "securityinsight.CommentsClient", "ListByCase", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListByCaseSender(req)
        if err != nil {
        result.ccl.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "securityinsight.CommentsClient", "ListByCase", resp, "Failure sending request")
        return
        }

        result.ccl, err = client.ListByCaseResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "securityinsight.CommentsClient", "ListByCase", resp, "Failure responding to request")
        }
            if result.ccl.hasNextLink() && result.ccl.IsEmpty() {
            err = result.NextWithContext(ctx)
            }

    return
}

    // ListByCasePreparer prepares the ListByCase request.
    func (client CommentsClient) ListByCasePreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, filter string, orderby string, top *int32, skipToken string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "caseId": autorest.Encode("path",caseID),
        "operationalInsightsResourceProvider": autorest.Encode("path",operationalInsightsResourceProvider),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        "workspaceName": autorest.Encode("path",workspaceName),
        }

            const APIVersion = "2019-01-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }
        if len(filter) > 0 {
        queryParameters["$filter"] = autorest.Encode("query",filter)
        }
        if len(orderby) > 0 {
        queryParameters["$orderby"] = autorest.Encode("query",orderby)
        }
        if top != nil {
        queryParameters["$top"] = autorest.Encode("query",*top)
        }
        if len(skipToken) > 0 {
        queryParameters["$skipToken"] = autorest.Encode("query",skipToken)
        }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/cases/{caseId}/comments",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByCaseSender sends the ListByCase request. The method will close the
    // http.Response Body if it receives an error.
    func (client CommentsClient) ListByCaseSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // ListByCaseResponder handles the response to the ListByCase request. The method always
    // closes the http.Response Body.
    func (client CommentsClient) ListByCaseResponder(resp *http.Response) (result CaseCommentList, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listByCaseNextResults retrieves the next set of results, if any.
            func (client CommentsClient) listByCaseNextResults(ctx context.Context, lastResults CaseCommentList) (result CaseCommentList, err error) {
            req, err := lastResults.caseCommentListPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "securityinsight.CommentsClient", "listByCaseNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByCaseSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "securityinsight.CommentsClient", "listByCaseNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByCaseResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "securityinsight.CommentsClient", "listByCaseNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListByCaseComplete enumerates all values, automatically crossing page boundaries as required.
            func (client CommentsClient) ListByCaseComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, filter string, orderby string, top *int32, skipToken string) (result CaseCommentListIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/CommentsClient.ListByCase")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.ListByCase(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, caseID, filter, orderby, top, skipToken)
                            return
            }

