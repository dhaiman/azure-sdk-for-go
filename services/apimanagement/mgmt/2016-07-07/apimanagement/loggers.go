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

// LoggersClient is the apiManagement Client
type LoggersClient struct {
    BaseClient
}
// NewLoggersClient creates an instance of the LoggersClient client.
func NewLoggersClient(subscriptionID string) LoggersClient {
    return NewLoggersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLoggersClientWithBaseURI creates an instance of the LoggersClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewLoggersClientWithBaseURI(baseURI string, subscriptionID string) LoggersClient {
        return LoggersClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates or Updates a logger.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // loggerid - identifier of the logger.
        // parameters - create parameters.
func (client LoggersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, loggerid string, parameters LoggerCreateParameters) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/LoggersClient.CreateOrUpdate")
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
            { TargetValue: loggerid,
             Constraints: []validation.Constraint{	{Target: "loggerid", Name: validation.MaxLength, Rule: 256, Chain: nil },
            	{Target: "loggerid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil }}},
            { TargetValue: parameters,
             Constraints: []validation.Constraint{	{Target: "parameters.Type", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "parameters.Credentials", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.LoggersClient", "CreateOrUpdate", err.Error())
            }

                req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, loggerid, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.LoggersClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateOrUpdateSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "apimanagement.LoggersClient", "CreateOrUpdate", resp, "Failure sending request")
            return
            }

            result, err = client.CreateOrUpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.LoggersClient", "CreateOrUpdate", resp, "Failure responding to request")
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client LoggersClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, loggerid string, parameters LoggerCreateParameters) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "loggerid": autorest.Encode("path",loggerid),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-07-07"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers/{loggerid}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client LoggersClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client LoggersClient) CreateOrUpdateResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Delete deletes the specified logger.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // loggerid - identifier of the logger.
        // ifMatch - the entity state (Etag) version of the logger to delete. A value of "*" can be used for If-Match
        // to unconditionally apply the operation.
func (client LoggersClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, loggerid string, ifMatch string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/LoggersClient.Delete")
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
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.LoggersClient", "Delete", err.Error())
            }

                req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, loggerid, ifMatch)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.LoggersClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "apimanagement.LoggersClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.LoggersClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client LoggersClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, loggerid string, ifMatch string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "loggerid": autorest.Encode("path",loggerid),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-07-07"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers/{loggerid}",pathParameters),
    autorest.WithQueryParameters(queryParameters),
    autorest.WithHeader("If-Match", autorest.String(ifMatch)))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client LoggersClient) DeleteSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client LoggersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets the details of the logger specified by its identifier.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // loggerid - identifier of the logger.
func (client LoggersClient) Get(ctx context.Context, resourceGroupName string, serviceName string, loggerid string) (result LoggerResponse, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/LoggersClient.Get")
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
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.LoggersClient", "Get", err.Error())
            }

                req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, loggerid)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.LoggersClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "apimanagement.LoggersClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.LoggersClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client LoggersClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, loggerid string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "loggerid": autorest.Encode("path",loggerid),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-07-07"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers/{loggerid}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client LoggersClient) GetSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client LoggersClient) GetResponder(resp *http.Response) (result LoggerResponse, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByService lists a collection of loggers in the specified service instance.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // filter - | Field | Supported operators    | Supported functions                         |
        // |-------|------------------------|---------------------------------------------|
        // | id    | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
        // | type  | eq                     |                                             |
        // top - number of records to return.
        // skip - number of records to skip.
func (client LoggersClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result LoggerCollectionPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/LoggersClient.ListByService")
        defer func() {
            sc := -1
            if result.lc.Response.Response != nil {
                sc = result.lc.Response.Response.StatusCode
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
            return result, validation.NewError("apimanagement.LoggersClient", "ListByService", err.Error())
            }

                        result.fn = client.listByServiceNextResults
    req, err := client.ListByServicePreparer(ctx, resourceGroupName, serviceName, filter, top, skip)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.LoggersClient", "ListByService", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByServiceSender(req)
            if err != nil {
            result.lc.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "apimanagement.LoggersClient", "ListByService", resp, "Failure sending request")
            return
            }

            result.lc, err = client.ListByServiceResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.LoggersClient", "ListByService", resp, "Failure responding to request")
            }

    return
    }

    // ListByServicePreparer prepares the ListByService request.
    func (client LoggersClient) ListByServicePreparer(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (*http.Request, error) {
            pathParameters := map[string]interface{} {
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

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByServiceSender sends the ListByService request. The method will close the
    // http.Response Body if it receives an error.
    func (client LoggersClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client LoggersClient) ListByServiceResponder(resp *http.Response) (result LoggerCollection, err error) {
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
            func (client LoggersClient) listByServiceNextResults(ctx context.Context, lastResults LoggerCollection) (result LoggerCollection, err error) {
            req, err := lastResults.loggerCollectionPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "apimanagement.LoggersClient", "listByServiceNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByServiceSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "apimanagement.LoggersClient", "listByServiceNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByServiceResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.LoggersClient", "listByServiceNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByServiceComplete enumerates all values, automatically crossing page boundaries as required.
    func (client LoggersClient) ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result LoggerCollectionIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/LoggersClient.ListByService")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByService(ctx, resourceGroupName, serviceName, filter, top, skip)
                return
        }

// Update updates an existing logger.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // loggerid - identifier of the logger.
        // parameters - update parameters.
        // ifMatch - the entity state (Etag) version of the logger to update. A value of "*" can be used for If-Match
        // to unconditionally apply the operation.
func (client LoggersClient) Update(ctx context.Context, resourceGroupName string, serviceName string, loggerid string, parameters LoggerUpdateParameters, ifMatch string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/LoggersClient.Update")
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
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.LoggersClient", "Update", err.Error())
            }

                req, err := client.UpdatePreparer(ctx, resourceGroupName, serviceName, loggerid, parameters, ifMatch)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.LoggersClient", "Update", nil , "Failure preparing request")
    return
    }

            resp, err := client.UpdateSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "apimanagement.LoggersClient", "Update", resp, "Failure sending request")
            return
            }

            result, err = client.UpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.LoggersClient", "Update", resp, "Failure responding to request")
            }

    return
    }

    // UpdatePreparer prepares the Update request.
    func (client LoggersClient) UpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, loggerid string, parameters LoggerUpdateParameters, ifMatch string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "loggerid": autorest.Encode("path",loggerid),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2016-07-07"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/loggers/{loggerid}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters),
    autorest.WithHeader("If-Match", autorest.String(ifMatch)))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client LoggersClient) UpdateSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client LoggersClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

