package netapp

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

// PoolsClient is the microsoft NetApp Azure Resource Provider specification
type PoolsClient struct {
    BaseClient
}
// NewPoolsClient creates an instance of the PoolsClient client.
func NewPoolsClient(subscriptionID string) PoolsClient {
    return NewPoolsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPoolsClientWithBaseURI creates an instance of the PoolsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewPoolsClientWithBaseURI(baseURI string, subscriptionID string) PoolsClient {
        return PoolsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate create or Update a capacity pool
    // Parameters:
        // body - capacity pool object supplied in the body of the operation.
        // resourceGroupName - the name of the resource group.
        // accountName - the name of the NetApp account
        // poolName - the name of the capacity pool
func (client PoolsClient) CreateOrUpdate(ctx context.Context, body CapacityPool, resourceGroupName string, accountName string, poolName string) (result PoolsCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PoolsClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: body,
             Constraints: []validation.Constraint{	{Target: "body.Location", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "body.PoolProperties", Name: validation.Null, Rule: true ,
            Chain: []validation.Constraint{	{Target: "body.PoolProperties.PoolID", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "body.PoolProperties.PoolID", Name: validation.MaxLength, Rule: 36, Chain: nil },
            	{Target: "body.PoolProperties.PoolID", Name: validation.MinLength, Rule: 36, Chain: nil },
            	{Target: "body.PoolProperties.PoolID", Name: validation.Pattern, Rule: `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`, Chain: nil },
            }},
            	{Target: "body.PoolProperties.Size", Name: validation.Null, Rule: true ,
            Chain: []validation.Constraint{	{Target: "body.PoolProperties.Size", Name: validation.InclusiveMaximum, Rule: int64(549755813888000), Chain: nil },
            	{Target: "body.PoolProperties.Size", Name: validation.InclusiveMinimum, Rule: int64(4398046511104), Chain: nil },
            }},
            }}}},
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("netapp.PoolsClient", "CreateOrUpdate", err.Error())
            }

                req, err := client.CreateOrUpdatePreparer(ctx, body, resourceGroupName, accountName, poolName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "netapp.PoolsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            result, err = client.CreateOrUpdateSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "netapp.PoolsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client PoolsClient) CreateOrUpdatePreparer(ctx context.Context, body CapacityPool, resourceGroupName string, accountName string, poolName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "accountName": autorest.Encode("path",accountName),
            "poolName": autorest.Encode("path",poolName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-08-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

                    body.ID = nil
                body.Name = nil
                body.Type = nil
    preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}",pathParameters),
    autorest.WithJSON(body),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client PoolsClient) CreateOrUpdateSender(req *http.Request) (future PoolsCreateOrUpdateFuture, err error) {
            var resp *http.Response
            resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client PoolsClient) CreateOrUpdateResponder(resp *http.Response) (result CapacityPool, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated,http.StatusAccepted),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete delete the specified capacity pool
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // accountName - the name of the NetApp account
        // poolName - the name of the capacity pool
func (client PoolsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, poolName string) (result PoolsDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PoolsClient.Delete")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("netapp.PoolsClient", "Delete", err.Error())
            }

                req, err := client.DeletePreparer(ctx, resourceGroupName, accountName, poolName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "netapp.PoolsClient", "Delete", nil , "Failure preparing request")
    return
    }

            result, err = client.DeleteSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "netapp.PoolsClient", "Delete", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client PoolsClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string, poolName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "accountName": autorest.Encode("path",accountName),
            "poolName": autorest.Encode("path",poolName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-08-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client PoolsClient) DeleteSender(req *http.Request) (future PoolsDeleteFuture, err error) {
            var resp *http.Response
            resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client PoolsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get get details of the specified capacity pool
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // accountName - the name of the NetApp account
        // poolName - the name of the capacity pool
func (client PoolsClient) Get(ctx context.Context, resourceGroupName string, accountName string, poolName string) (result CapacityPool, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PoolsClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("netapp.PoolsClient", "Get", err.Error())
            }

                req, err := client.GetPreparer(ctx, resourceGroupName, accountName, poolName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "netapp.PoolsClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "netapp.PoolsClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "netapp.PoolsClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client PoolsClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, poolName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "accountName": autorest.Encode("path",accountName),
            "poolName": autorest.Encode("path",poolName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-08-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client PoolsClient) GetSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client PoolsClient) GetResponder(resp *http.Response) (result CapacityPool, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List list all capacity pools in the NetApp Account
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // accountName - the name of the NetApp account
func (client PoolsClient) List(ctx context.Context, resourceGroupName string, accountName string) (result CapacityPoolList, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PoolsClient.List")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("netapp.PoolsClient", "List", err.Error())
            }

                req, err := client.ListPreparer(ctx, resourceGroupName, accountName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "netapp.PoolsClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "netapp.PoolsClient", "List", resp, "Failure sending request")
            return
            }

            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "netapp.PoolsClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client PoolsClient) ListPreparer(ctx context.Context, resourceGroupName string, accountName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "accountName": autorest.Encode("path",accountName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-08-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client PoolsClient) ListSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PoolsClient) ListResponder(resp *http.Response) (result CapacityPoolList, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Update patch the specified capacity pool
    // Parameters:
        // body - capacity pool object supplied in the body of the operation.
        // resourceGroupName - the name of the resource group.
        // accountName - the name of the NetApp account
        // poolName - the name of the capacity pool
func (client PoolsClient) Update(ctx context.Context, body CapacityPoolPatch, resourceGroupName string, accountName string, poolName string) (result CapacityPool, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PoolsClient.Update")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceGroupName,
             Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
            	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}}}); err != nil {
            return result, validation.NewError("netapp.PoolsClient", "Update", err.Error())
            }

                req, err := client.UpdatePreparer(ctx, body, resourceGroupName, accountName, poolName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "netapp.PoolsClient", "Update", nil , "Failure preparing request")
    return
    }

            resp, err := client.UpdateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "netapp.PoolsClient", "Update", resp, "Failure sending request")
            return
            }

            result, err = client.UpdateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "netapp.PoolsClient", "Update", resp, "Failure responding to request")
            }

    return
    }

    // UpdatePreparer prepares the Update request.
    func (client PoolsClient) UpdatePreparer(ctx context.Context, body CapacityPoolPatch, resourceGroupName string, accountName string, poolName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "accountName": autorest.Encode("path",accountName),
            "poolName": autorest.Encode("path",poolName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-08-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

                    body.ID = nil
                body.Name = nil
                body.Type = nil
    preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}",pathParameters),
    autorest.WithJSON(body),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client PoolsClient) UpdateSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client PoolsClient) UpdateResponder(resp *http.Response) (result CapacityPool, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

