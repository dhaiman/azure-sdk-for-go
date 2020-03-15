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

// SubscriptionsClient is the apiManagement Client
type SubscriptionsClient struct {
    BaseClient
}
// NewSubscriptionsClient creates an instance of the SubscriptionsClient client.
func NewSubscriptionsClient(subscriptionID string) SubscriptionsClient {
    return NewSubscriptionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSubscriptionsClientWithBaseURI creates an instance of the SubscriptionsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) SubscriptionsClient {
        return SubscriptionsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates or updates the subscription of specified user to the specified product.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // sid - subscription entity Identifier. The entity represents the association between a user and a product in
        // API Management.
        // parameters - create parameters.
func (client SubscriptionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, sid string, parameters SubscriptionCreateParameters) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SubscriptionsClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: serviceName,
             Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}},
            { TargetValue: sid,
             Constraints: []validation.Constraint{	{Target: "sid", Name: validation.MaxLength, Rule: 256, Chain: nil },
            	{Target: "sid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil }}},
            { TargetValue: parameters,
             Constraints: []validation.Constraint{	{Target: "parameters.UserID", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "parameters.ProductID", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "parameters.Name", Name: validation.Null, Rule: true ,
            Chain: []validation.Constraint{	{Target: "parameters.Name", Name: validation.MaxLength, Rule: 100, Chain: nil },
            	{Target: "parameters.Name", Name: validation.MinLength, Rule: 1, Chain: nil },
            }},
            	{Target: "parameters.PrimaryKey", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "parameters.PrimaryKey", Name: validation.MaxLength, Rule: 256, Chain: nil },
            	{Target: "parameters.PrimaryKey", Name: validation.MinLength, Rule: 1, Chain: nil },
            }},
            	{Target: "parameters.SecondaryKey", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "parameters.SecondaryKey", Name: validation.MaxLength, Rule: 256, Chain: nil },
            	{Target: "parameters.SecondaryKey", Name: validation.MinLength, Rule: 1, Chain: nil },
            }}}}}); err != nil {
            return result, validation.NewError("apimanagement.SubscriptionsClient", "CreateOrUpdate", err.Error())
            }

                req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, sid, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateOrUpdateSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "CreateOrUpdate", resp, "Failure sending request")
            return
            }

            result, err = client.CreateOrUpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "CreateOrUpdate", resp, "Failure responding to request")
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client SubscriptionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, sid string, parameters SubscriptionCreateParameters) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "sid": autorest.Encode("path",sid),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-10-10"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client SubscriptionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) CreateOrUpdateResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Delete deletes the specified subscription.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // sid - subscription entity Identifier. The entity represents the association between a user and a product in
        // API Management.
        // ifMatch - eTag of the Subscription Entity. ETag should match the current entity state from the header
        // response of the GET request or it should be * for unconditional update.
func (client SubscriptionsClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, sid string, ifMatch string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SubscriptionsClient.Delete")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: serviceName,
             Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}},
            { TargetValue: sid,
             Constraints: []validation.Constraint{	{Target: "sid", Name: validation.MaxLength, Rule: 256, Chain: nil },
            	{Target: "sid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.SubscriptionsClient", "Delete", err.Error())
            }

                req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, sid, ifMatch)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client SubscriptionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, sid string, ifMatch string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "sid": autorest.Encode("path",sid),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-10-10"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}",pathParameters),
    autorest.WithQueryParameters(queryParameters),
    autorest.WithHeader("If-Match", autorest.String(ifMatch)))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client SubscriptionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets the specified Subscription entity.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // sid - subscription entity Identifier. The entity represents the association between a user and a product in
        // API Management.
func (client SubscriptionsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, sid string) (result SubscriptionContract, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SubscriptionsClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: serviceName,
             Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}},
            { TargetValue: sid,
             Constraints: []validation.Constraint{	{Target: "sid", Name: validation.MaxLength, Rule: 256, Chain: nil },
            	{Target: "sid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.SubscriptionsClient", "Get", err.Error())
            }

                req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, sid)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client SubscriptionsClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, sid string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "sid": autorest.Encode("path",sid),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-10-10"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client SubscriptionsClient) GetSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) GetResponder(resp *http.Response) (result SubscriptionContract, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List lists all subscriptions of the API Management service instance.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // filter - | Field        | Supported operators    | Supported functions                         |
        // |--------------|------------------------|---------------------------------------------|
        // | id           | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
        // | name         | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
        // | stateComment | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
        // | userId       | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
        // | productId    | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
        // | state        | eq                     |                                             |
        // top - number of records to return.
        // skip - number of records to skip.
func (client SubscriptionsClient) List(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result SubscriptionCollectionPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SubscriptionsClient.List")
        defer func() {
            sc := -1
            if result.sc.Response.Response != nil {
                sc = result.sc.Response.Response.StatusCode
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
            return result, validation.NewError("apimanagement.SubscriptionsClient", "List", err.Error())
            }

                        result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName, serviceName, filter, top, skip)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.sc.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "List", resp, "Failure sending request")
            return
            }

            result.sc, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client SubscriptionsClient) ListPreparer(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-10-10"
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
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client SubscriptionsClient) ListSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) ListResponder(resp *http.Response) (result SubscriptionCollection, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client SubscriptionsClient) listNextResults(ctx context.Context, lastResults SubscriptionCollection) (result SubscriptionCollection, err error) {
            req, err := lastResults.subscriptionCollectionPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListComplete enumerates all values, automatically crossing page boundaries as required.
    func (client SubscriptionsClient) ListComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result SubscriptionCollectionIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/SubscriptionsClient.List")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.List(ctx, resourceGroupName, serviceName, filter, top, skip)
                return
        }

// RegeneratePrimaryKey regenerates primary key of existing subscription of the API Management service instance.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // sid - subscription entity Identifier. The entity represents the association between a user and a product in
        // API Management.
func (client SubscriptionsClient) RegeneratePrimaryKey(ctx context.Context, resourceGroupName string, serviceName string, sid string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SubscriptionsClient.RegeneratePrimaryKey")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: serviceName,
             Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}},
            { TargetValue: sid,
             Constraints: []validation.Constraint{	{Target: "sid", Name: validation.MaxLength, Rule: 256, Chain: nil },
            	{Target: "sid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.SubscriptionsClient", "RegeneratePrimaryKey", err.Error())
            }

                req, err := client.RegeneratePrimaryKeyPreparer(ctx, resourceGroupName, serviceName, sid)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "RegeneratePrimaryKey", nil , "Failure preparing request")
    return
    }

            resp, err := client.RegeneratePrimaryKeySender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "RegeneratePrimaryKey", resp, "Failure sending request")
            return
            }

            result, err = client.RegeneratePrimaryKeyResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "RegeneratePrimaryKey", resp, "Failure responding to request")
            }

    return
    }

    // RegeneratePrimaryKeyPreparer prepares the RegeneratePrimaryKey request.
    func (client SubscriptionsClient) RegeneratePrimaryKeyPreparer(ctx context.Context, resourceGroupName string, serviceName string, sid string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "sid": autorest.Encode("path",sid),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-10-10"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}/regeneratePrimaryKey",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // RegeneratePrimaryKeySender sends the RegeneratePrimaryKey request. The method will close the
    // http.Response Body if it receives an error.
    func (client SubscriptionsClient) RegeneratePrimaryKeySender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// RegeneratePrimaryKeyResponder handles the response to the RegeneratePrimaryKey request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) RegeneratePrimaryKeyResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// RegenerateSecondaryKey regenerates secondary key of existing subscription of the API Management service instance.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // sid - subscription entity Identifier. The entity represents the association between a user and a product in
        // API Management.
func (client SubscriptionsClient) RegenerateSecondaryKey(ctx context.Context, resourceGroupName string, serviceName string, sid string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SubscriptionsClient.RegenerateSecondaryKey")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: serviceName,
             Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}},
            { TargetValue: sid,
             Constraints: []validation.Constraint{	{Target: "sid", Name: validation.MaxLength, Rule: 256, Chain: nil },
            	{Target: "sid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.SubscriptionsClient", "RegenerateSecondaryKey", err.Error())
            }

                req, err := client.RegenerateSecondaryKeyPreparer(ctx, resourceGroupName, serviceName, sid)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "RegenerateSecondaryKey", nil , "Failure preparing request")
    return
    }

            resp, err := client.RegenerateSecondaryKeySender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "RegenerateSecondaryKey", resp, "Failure sending request")
            return
            }

            result, err = client.RegenerateSecondaryKeyResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "RegenerateSecondaryKey", resp, "Failure responding to request")
            }

    return
    }

    // RegenerateSecondaryKeyPreparer prepares the RegenerateSecondaryKey request.
    func (client SubscriptionsClient) RegenerateSecondaryKeyPreparer(ctx context.Context, resourceGroupName string, serviceName string, sid string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "sid": autorest.Encode("path",sid),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-10-10"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}/regenerateSecondaryKey",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // RegenerateSecondaryKeySender sends the RegenerateSecondaryKey request. The method will close the
    // http.Response Body if it receives an error.
    func (client SubscriptionsClient) RegenerateSecondaryKeySender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// RegenerateSecondaryKeyResponder handles the response to the RegenerateSecondaryKey request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) RegenerateSecondaryKeyResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Update updates the details of a subscription specified by its identifier.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // sid - subscription entity Identifier. The entity represents the association between a user and a product in
        // API Management.
        // parameters - update parameters.
        // ifMatch - eTag of the Subscription Entity. ETag should match the current entity state from the header
        // response of the GET request or it should be * for unconditional update.
func (client SubscriptionsClient) Update(ctx context.Context, resourceGroupName string, serviceName string, sid string, parameters SubscriptionUpdateParameters, ifMatch string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SubscriptionsClient.Update")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: serviceName,
             Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}},
            { TargetValue: sid,
             Constraints: []validation.Constraint{	{Target: "sid", Name: validation.MaxLength, Rule: 256, Chain: nil },
            	{Target: "sid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.SubscriptionsClient", "Update", err.Error())
            }

                req, err := client.UpdatePreparer(ctx, resourceGroupName, serviceName, sid, parameters, ifMatch)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "Update", nil , "Failure preparing request")
    return
    }

            resp, err := client.UpdateSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "Update", resp, "Failure sending request")
            return
            }

            result, err = client.UpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.SubscriptionsClient", "Update", resp, "Failure responding to request")
            }

    return
    }

    // UpdatePreparer prepares the Update request.
    func (client SubscriptionsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, sid string, parameters SubscriptionUpdateParameters, ifMatch string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "sid": autorest.Encode("path",sid),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-10-10"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/subscriptions/{sid}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters),
    autorest.WithHeader("If-Match", autorest.String(ifMatch)))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client SubscriptionsClient) UpdateSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

