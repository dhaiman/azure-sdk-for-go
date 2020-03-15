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

// APIVersionSetClient is the apiManagement Client
type APIVersionSetClient struct {
    BaseClient
}
// NewAPIVersionSetClient creates an instance of the APIVersionSetClient client.
func NewAPIVersionSetClient(subscriptionID string) APIVersionSetClient {
    return NewAPIVersionSetClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAPIVersionSetClientWithBaseURI creates an instance of the APIVersionSetClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewAPIVersionSetClientWithBaseURI(baseURI string, subscriptionID string) APIVersionSetClient {
        return APIVersionSetClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates or Updates a Api Version Set.
    // Parameters:
        // parameters - create or update parameters.
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // versionSetID - api Version Set identifier. Must be unique in the current API Management service instance.
        // ifMatch - the entity state (Etag) version of the user to update. A value of "*" can be used for If-Match to
        // unconditionally apply the operation.
func (client APIVersionSetClient) CreateOrUpdate(ctx context.Context, parameters APIVersionSetContract, resourceGroupName string, serviceName string, versionSetID string, ifMatch string) (result APIVersionSetContract, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/APIVersionSetClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: parameters,
             Constraints: []validation.Constraint{	{Target: "parameters.APIVersionSetContractProperties", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "parameters.APIVersionSetContractProperties.DisplayName", Name: validation.Null, Rule: true ,
            Chain: []validation.Constraint{	{Target: "parameters.APIVersionSetContractProperties.DisplayName", Name: validation.MaxLength, Rule: 100, Chain: nil },
            	{Target: "parameters.APIVersionSetContractProperties.DisplayName", Name: validation.MinLength, Rule: 1, Chain: nil },
            }},
            }}}},
            { TargetValue: serviceName,
             Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}},
            { TargetValue: versionSetID,
             Constraints: []validation.Constraint{	{Target: "versionSetID", Name: validation.MaxLength, Rule: 256, Chain: nil },
            	{Target: "versionSetID", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "versionSetID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.APIVersionSetClient", "CreateOrUpdate", err.Error())
            }

                req, err := client.CreateOrUpdatePreparer(ctx, parameters, resourceGroupName, serviceName, versionSetID, ifMatch)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.APIVersionSetClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            resp, err := client.CreateOrUpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "apimanagement.APIVersionSetClient", "CreateOrUpdate", resp, "Failure sending request")
            return
            }

            result, err = client.CreateOrUpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.APIVersionSetClient", "CreateOrUpdate", resp, "Failure responding to request")
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client APIVersionSetClient) CreateOrUpdatePreparer(ctx context.Context, parameters APIVersionSetContract, resourceGroupName string, serviceName string, versionSetID string, ifMatch string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "versionSetId": autorest.Encode("path",versionSetID),
            }

                        const APIVersion = "2017-03-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/api-version-sets/{versionSetId}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
            if len(ifMatch) > 0 {
            preparer = autorest.DecoratePreparer(preparer,
            autorest.WithHeader("If-Match",autorest.String(ifMatch)))
            }
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client APIVersionSetClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client APIVersionSetClient) CreateOrUpdateResponder(resp *http.Response) (result APIVersionSetContract, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete deletes specific Api Version Set.
    // Parameters:
        // ifMatch - the entity state (Etag) version of the Api Version Set to delete. A value of "*" can be used for
        // If-Match to unconditionally apply the operation.
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // versionSetID - api Version Set identifier. Must be unique in the current API Management service instance.
func (client APIVersionSetClient) Delete(ctx context.Context, ifMatch string, resourceGroupName string, serviceName string, versionSetID string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/APIVersionSetClient.Delete")
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
            { TargetValue: versionSetID,
             Constraints: []validation.Constraint{	{Target: "versionSetID", Name: validation.MaxLength, Rule: 256, Chain: nil },
            	{Target: "versionSetID", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "versionSetID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.APIVersionSetClient", "Delete", err.Error())
            }

                req, err := client.DeletePreparer(ctx, ifMatch, resourceGroupName, serviceName, versionSetID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.APIVersionSetClient", "Delete", nil , "Failure preparing request")
    return
    }

            resp, err := client.DeleteSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "apimanagement.APIVersionSetClient", "Delete", resp, "Failure sending request")
            return
            }

            result, err = client.DeleteResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.APIVersionSetClient", "Delete", resp, "Failure responding to request")
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client APIVersionSetClient) DeletePreparer(ctx context.Context, ifMatch string, resourceGroupName string, serviceName string, versionSetID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "versionSetId": autorest.Encode("path",versionSetID),
            }

                        const APIVersion = "2017-03-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/api-version-sets/{versionSetId}",pathParameters),
    autorest.WithQueryParameters(queryParameters),
    autorest.WithHeader("If-Match", autorest.String(ifMatch)))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client APIVersionSetClient) DeleteSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client APIVersionSetClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets the details of the Api Version Set specified by its identifier.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // versionSetID - api Version Set identifier. Must be unique in the current API Management service instance.
func (client APIVersionSetClient) Get(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string) (result APIVersionSetContract, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/APIVersionSetClient.Get")
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
            { TargetValue: versionSetID,
             Constraints: []validation.Constraint{	{Target: "versionSetID", Name: validation.MaxLength, Rule: 256, Chain: nil },
            	{Target: "versionSetID", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "versionSetID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.APIVersionSetClient", "Get", err.Error())
            }

                req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, versionSetID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.APIVersionSetClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "apimanagement.APIVersionSetClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.APIVersionSetClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client APIVersionSetClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "versionSetId": autorest.Encode("path",versionSetID),
            }

                        const APIVersion = "2017-03-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/api-version-sets/{versionSetId}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client APIVersionSetClient) GetSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client APIVersionSetClient) GetResponder(resp *http.Response) (result APIVersionSetContract, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// GetEntityTag gets the entity state (Etag) version of the Api Version Set specified by its identifier.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // versionSetID - api Version Set identifier. Must be unique in the current API Management service instance.
func (client APIVersionSetClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/APIVersionSetClient.GetEntityTag")
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
            { TargetValue: versionSetID,
             Constraints: []validation.Constraint{	{Target: "versionSetID", Name: validation.MaxLength, Rule: 256, Chain: nil },
            	{Target: "versionSetID", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "versionSetID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.APIVersionSetClient", "GetEntityTag", err.Error())
            }

                req, err := client.GetEntityTagPreparer(ctx, resourceGroupName, serviceName, versionSetID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.APIVersionSetClient", "GetEntityTag", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetEntityTagSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "apimanagement.APIVersionSetClient", "GetEntityTag", resp, "Failure sending request")
            return
            }

            result, err = client.GetEntityTagResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.APIVersionSetClient", "GetEntityTag", resp, "Failure responding to request")
            }

    return
    }

    // GetEntityTagPreparer prepares the GetEntityTag request.
    func (client APIVersionSetClient) GetEntityTagPreparer(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "versionSetId": autorest.Encode("path",versionSetID),
            }

                        const APIVersion = "2017-03-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsHead(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/api-version-sets/{versionSetId}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetEntityTagSender sends the GetEntityTag request. The method will close the
    // http.Response Body if it receives an error.
    func (client APIVersionSetClient) GetEntityTagSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// GetEntityTagResponder handles the response to the GetEntityTag request. The method always
// closes the http.Response Body.
func (client APIVersionSetClient) GetEntityTagResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// ListByService lists a collection of API Version Sets in the specified service instance.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // filter - | Field            | Supported operators    | Supported functions               |
        // |------------------|------------------------|-----------------------------------|
        // | id               | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
        // | firstName        | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
        // | lastName         | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
        // | email            | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
        // | state            | eq                     | N/A                               |
        // | registrationDate | ge, le, eq, ne, gt, lt | N/A                               |
        // | note             | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
        // top - number of records to return.
        // skip - number of records to skip.
func (client APIVersionSetClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result APIVersionSetCollectionPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/APIVersionSetClient.ListByService")
        defer func() {
            sc := -1
            if result.avsc.Response.Response != nil {
                sc = result.avsc.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: top,
             Constraints: []validation.Constraint{	{Target: "top", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil },
            }}}},
            { TargetValue: skip,
             Constraints: []validation.Constraint{	{Target: "skip", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil },
            }}}},
            { TargetValue: serviceName,
             Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil },
            	{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.APIVersionSetClient", "ListByService", err.Error())
            }

                        result.fn = client.listByServiceNextResults
    req, err := client.ListByServicePreparer(ctx, resourceGroupName, serviceName, filter, top, skip)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.APIVersionSetClient", "ListByService", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByServiceSender(req)
            if err != nil {
            result.avsc.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "apimanagement.APIVersionSetClient", "ListByService", resp, "Failure sending request")
            return
            }

            result.avsc, err = client.ListByServiceResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.APIVersionSetClient", "ListByService", resp, "Failure responding to request")
            }

    return
    }

    // ListByServicePreparer prepares the ListByService request.
    func (client APIVersionSetClient) ListByServicePreparer(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-03-01"
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
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/api-version-sets",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByServiceSender sends the ListByService request. The method will close the
    // http.Response Body if it receives an error.
    func (client APIVersionSetClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client APIVersionSetClient) ListByServiceResponder(resp *http.Response) (result APIVersionSetCollection, err error) {
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
            func (client APIVersionSetClient) listByServiceNextResults(ctx context.Context, lastResults APIVersionSetCollection) (result APIVersionSetCollection, err error) {
            req, err := lastResults.aPIVersionSetCollectionPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "apimanagement.APIVersionSetClient", "listByServiceNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByServiceSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "apimanagement.APIVersionSetClient", "listByServiceNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByServiceResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.APIVersionSetClient", "listByServiceNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByServiceComplete enumerates all values, automatically crossing page boundaries as required.
    func (client APIVersionSetClient) ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result APIVersionSetCollectionIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/APIVersionSetClient.ListByService")
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

// Update updates the details of the Api VersionSet specified by its identifier.
    // Parameters:
        // parameters - update parameters.
        // ifMatch - the entity state (Etag) version of the user to update. A value of "*" can be used for If-Match to
        // unconditionally apply the operation.
        // resourceGroupName - the name of the resource group.
        // serviceName - the name of the API Management service.
        // versionSetID - api Version Set identifier. Must be unique in the current API Management service instance.
func (client APIVersionSetClient) Update(ctx context.Context, parameters APIVersionSetUpdateParameters, ifMatch string, resourceGroupName string, serviceName string, versionSetID string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/APIVersionSetClient.Update")
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
            { TargetValue: versionSetID,
             Constraints: []validation.Constraint{	{Target: "versionSetID", Name: validation.MaxLength, Rule: 256, Chain: nil },
            	{Target: "versionSetID", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "versionSetID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("apimanagement.APIVersionSetClient", "Update", err.Error())
            }

                req, err := client.UpdatePreparer(ctx, parameters, ifMatch, resourceGroupName, serviceName, versionSetID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "apimanagement.APIVersionSetClient", "Update", nil , "Failure preparing request")
    return
    }

            resp, err := client.UpdateSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "apimanagement.APIVersionSetClient", "Update", resp, "Failure sending request")
            return
            }

            result, err = client.UpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "apimanagement.APIVersionSetClient", "Update", resp, "Failure responding to request")
            }

    return
    }

    // UpdatePreparer prepares the Update request.
    func (client APIVersionSetClient) UpdatePreparer(ctx context.Context, parameters APIVersionSetUpdateParameters, ifMatch string, resourceGroupName string, serviceName string, versionSetID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "serviceName": autorest.Encode("path",serviceName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "versionSetId": autorest.Encode("path",versionSetID),
            }

                        const APIVersion = "2017-03-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/api-version-sets/{versionSetId}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters),
    autorest.WithHeader("If-Match", autorest.String(ifMatch)))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client APIVersionSetClient) UpdateSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client APIVersionSetClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

