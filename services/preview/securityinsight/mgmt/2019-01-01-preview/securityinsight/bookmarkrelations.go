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

// BookmarkRelationsClient is the API spec for Microsoft.SecurityInsights (Azure Security Insights) resource provider
type BookmarkRelationsClient struct {
    BaseClient
}
// NewBookmarkRelationsClient creates an instance of the BookmarkRelationsClient client.
func NewBookmarkRelationsClient(subscriptionID string) BookmarkRelationsClient {
    return NewBookmarkRelationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBookmarkRelationsClientWithBaseURI creates an instance of the BookmarkRelationsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
    func NewBookmarkRelationsClientWithBaseURI(baseURI string, subscriptionID string) BookmarkRelationsClient {
        return BookmarkRelationsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdateRelation creates the bookmark relation.
    // Parameters:
        // resourceGroupName - the name of the resource group within the user's subscription. The name is case
        // insensitive.
        // operationalInsightsResourceProvider - the namespace of workspaces resource provider-
        // Microsoft.OperationalInsights.
        // workspaceName - the name of the workspace.
        // bookmarkID - bookmark ID
        // relationName - relation Name
        // relation - the relation model
func (client BookmarkRelationsClient) CreateOrUpdateRelation(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, bookmarkID string, relationName string, relation Relation) (result Relation, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/BookmarkRelationsClient.CreateOrUpdateRelation")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
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
        	{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil }}},
        { TargetValue: relation,
         Constraints: []validation.Constraint{	{Target: "relation.RelationProperties", Name: validation.Null, Rule: false ,
        Chain: []validation.Constraint{	{Target: "relation.RelationProperties.RelatedResourceID", Name: validation.Null, Rule: true, Chain: nil },
        }}}}}); err != nil {
        return result, validation.NewError("securityinsight.BookmarkRelationsClient", "CreateOrUpdateRelation", err.Error())
        }

        req, err := client.CreateOrUpdateRelationPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, bookmarkID, relationName, relation)
    if err != nil {
    err = autorest.NewErrorWithError(err, "securityinsight.BookmarkRelationsClient", "CreateOrUpdateRelation", nil , "Failure preparing request")
    return
    }

        resp, err := client.CreateOrUpdateRelationSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "securityinsight.BookmarkRelationsClient", "CreateOrUpdateRelation", resp, "Failure sending request")
        return
        }

        result, err = client.CreateOrUpdateRelationResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "securityinsight.BookmarkRelationsClient", "CreateOrUpdateRelation", resp, "Failure responding to request")
        }

    return
}

    // CreateOrUpdateRelationPreparer prepares the CreateOrUpdateRelation request.
    func (client BookmarkRelationsClient) CreateOrUpdateRelationPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, bookmarkID string, relationName string, relation Relation) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "bookmarkId": autorest.Encode("path",bookmarkID),
        "operationalInsightsResourceProvider": autorest.Encode("path",operationalInsightsResourceProvider),
        "relationName": autorest.Encode("path",relationName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        "workspaceName": autorest.Encode("path",workspaceName),
        }

            const APIVersion = "2019-01-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/bookmarks/{bookmarkId}/relations/{relationName}",pathParameters),
autorest.WithJSON(relation),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateRelationSender sends the CreateOrUpdateRelation request. The method will close the
    // http.Response Body if it receives an error.
    func (client BookmarkRelationsClient) CreateOrUpdateRelationSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // CreateOrUpdateRelationResponder handles the response to the CreateOrUpdateRelation request. The method always
    // closes the http.Response Body.
    func (client BookmarkRelationsClient) CreateOrUpdateRelationResponder(resp *http.Response) (result Relation, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// DeleteRelation delete the bookmark relation.
    // Parameters:
        // resourceGroupName - the name of the resource group within the user's subscription. The name is case
        // insensitive.
        // operationalInsightsResourceProvider - the namespace of workspaces resource provider-
        // Microsoft.OperationalInsights.
        // workspaceName - the name of the workspace.
        // bookmarkID - bookmark ID
        // relationName - relation Name
func (client BookmarkRelationsClient) DeleteRelation(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, bookmarkID string, relationName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/BookmarkRelationsClient.DeleteRelation")
        defer func() {
            sc := -1
        if result.Response != nil {
        sc = result.Response.StatusCode
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
        return result, validation.NewError("securityinsight.BookmarkRelationsClient", "DeleteRelation", err.Error())
        }

        req, err := client.DeleteRelationPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, bookmarkID, relationName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "securityinsight.BookmarkRelationsClient", "DeleteRelation", nil , "Failure preparing request")
    return
    }

        resp, err := client.DeleteRelationSender(req)
        if err != nil {
        result.Response = resp
        err = autorest.NewErrorWithError(err, "securityinsight.BookmarkRelationsClient", "DeleteRelation", resp, "Failure sending request")
        return
        }

        result, err = client.DeleteRelationResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "securityinsight.BookmarkRelationsClient", "DeleteRelation", resp, "Failure responding to request")
        }

    return
}

    // DeleteRelationPreparer prepares the DeleteRelation request.
    func (client BookmarkRelationsClient) DeleteRelationPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, bookmarkID string, relationName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "bookmarkId": autorest.Encode("path",bookmarkID),
        "operationalInsightsResourceProvider": autorest.Encode("path",operationalInsightsResourceProvider),
        "relationName": autorest.Encode("path",relationName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        "workspaceName": autorest.Encode("path",workspaceName),
        }

            const APIVersion = "2019-01-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/bookmarks/{bookmarkId}/relations/{relationName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteRelationSender sends the DeleteRelation request. The method will close the
    // http.Response Body if it receives an error.
    func (client BookmarkRelationsClient) DeleteRelationSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // DeleteRelationResponder handles the response to the DeleteRelation request. The method always
    // closes the http.Response Body.
    func (client BookmarkRelationsClient) DeleteRelationResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// GetRelation gets a bookmark relation.
    // Parameters:
        // resourceGroupName - the name of the resource group within the user's subscription. The name is case
        // insensitive.
        // operationalInsightsResourceProvider - the namespace of workspaces resource provider-
        // Microsoft.OperationalInsights.
        // workspaceName - the name of the workspace.
        // bookmarkID - bookmark ID
        // relationName - relation Name
func (client BookmarkRelationsClient) GetRelation(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, bookmarkID string, relationName string) (result Relation, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/BookmarkRelationsClient.GetRelation")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
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
        return result, validation.NewError("securityinsight.BookmarkRelationsClient", "GetRelation", err.Error())
        }

        req, err := client.GetRelationPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, bookmarkID, relationName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "securityinsight.BookmarkRelationsClient", "GetRelation", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetRelationSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "securityinsight.BookmarkRelationsClient", "GetRelation", resp, "Failure sending request")
        return
        }

        result, err = client.GetRelationResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "securityinsight.BookmarkRelationsClient", "GetRelation", resp, "Failure responding to request")
        }

    return
}

    // GetRelationPreparer prepares the GetRelation request.
    func (client BookmarkRelationsClient) GetRelationPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, bookmarkID string, relationName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "bookmarkId": autorest.Encode("path",bookmarkID),
        "operationalInsightsResourceProvider": autorest.Encode("path",operationalInsightsResourceProvider),
        "relationName": autorest.Encode("path",relationName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        "workspaceName": autorest.Encode("path",workspaceName),
        }

            const APIVersion = "2019-01-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/bookmarks/{bookmarkId}/relations/{relationName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetRelationSender sends the GetRelation request. The method will close the
    // http.Response Body if it receives an error.
    func (client BookmarkRelationsClient) GetRelationSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // GetRelationResponder handles the response to the GetRelation request. The method always
    // closes the http.Response Body.
    func (client BookmarkRelationsClient) GetRelationResponder(resp *http.Response) (result Relation, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// List gets all bookmark relations.
    // Parameters:
        // resourceGroupName - the name of the resource group within the user's subscription. The name is case
        // insensitive.
        // operationalInsightsResourceProvider - the namespace of workspaces resource provider-
        // Microsoft.OperationalInsights.
        // workspaceName - the name of the workspace.
        // bookmarkID - bookmark ID
        // filter - filters the results, based on a Boolean condition. Optional.
        // orderby - sorts the results. Optional.
        // top - returns only the first n results. Optional.
        // skipToken - skiptoken is only used if a previous operation returned a partial result. If a previous response
        // contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
        // specifies a starting point to use for subsequent calls. Optional.
func (client BookmarkRelationsClient) List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, bookmarkID string, filter string, orderby string, top *int32, skipToken string) (result RelationListPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/BookmarkRelationsClient.List")
        defer func() {
            sc := -1
        if result.rl.Response.Response != nil {
        sc = result.rl.Response.Response.StatusCode
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
        return result, validation.NewError("securityinsight.BookmarkRelationsClient", "List", err.Error())
        }

            result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, bookmarkID, filter, orderby, top, skipToken)
    if err != nil {
    err = autorest.NewErrorWithError(err, "securityinsight.BookmarkRelationsClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.rl.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "securityinsight.BookmarkRelationsClient", "List", resp, "Failure sending request")
        return
        }

        result.rl, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "securityinsight.BookmarkRelationsClient", "List", resp, "Failure responding to request")
        }
            if result.rl.hasNextLink() && result.rl.IsEmpty() {
            err = result.NextWithContext(ctx)
            }

    return
}

    // ListPreparer prepares the List request.
    func (client BookmarkRelationsClient) ListPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, bookmarkID string, filter string, orderby string, top *int32, skipToken string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "bookmarkId": autorest.Encode("path",bookmarkID),
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
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/bookmarks/{bookmarkId}/relations",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client BookmarkRelationsClient) ListSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client BookmarkRelationsClient) ListResponder(resp *http.Response) (result RelationList, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client BookmarkRelationsClient) listNextResults(ctx context.Context, lastResults RelationList) (result RelationList, err error) {
            req, err := lastResults.relationListPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "securityinsight.BookmarkRelationsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "securityinsight.BookmarkRelationsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "securityinsight.BookmarkRelationsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListComplete enumerates all values, automatically crossing page boundaries as required.
            func (client BookmarkRelationsClient) ListComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, bookmarkID string, filter string, orderby string, top *int32, skipToken string) (result RelationListIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BookmarkRelationsClient.List")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.List(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, bookmarkID, filter, orderby, top, skipToken)
                            return
            }

