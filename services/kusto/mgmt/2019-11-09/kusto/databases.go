package kusto

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

// DatabasesClient is the the Azure Kusto management API provides a RESTful set of web services that interact with
// Azure Kusto services to manage your clusters and databases. The API enables you to create, update, and delete
// clusters and databases.
type DatabasesClient struct {
    BaseClient
}
// NewDatabasesClient creates an instance of the DatabasesClient client.
func NewDatabasesClient(subscriptionID string) DatabasesClient {
    return NewDatabasesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDatabasesClientWithBaseURI creates an instance of the DatabasesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewDatabasesClientWithBaseURI(baseURI string, subscriptionID string) DatabasesClient {
        return DatabasesClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// AddPrincipals add Database principals permissions.
    // Parameters:
        // resourceGroupName - the name of the resource group containing the Kusto cluster.
        // clusterName - the name of the Kusto cluster.
        // databaseName - the name of the database in the Kusto cluster.
        // databasePrincipalsToAdd - list of database principals to add.
func (client DatabasesClient) AddPrincipals(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, databasePrincipalsToAdd DatabasePrincipalListRequest) (result DatabasePrincipalListResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DatabasesClient.AddPrincipals")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.AddPrincipalsPreparer(ctx, resourceGroupName, clusterName, databaseName, databasePrincipalsToAdd)
    if err != nil {
    err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "AddPrincipals", nil , "Failure preparing request")
    return
    }

            resp, err := client.AddPrincipalsSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "AddPrincipals", resp, "Failure sending request")
            return
            }

            result, err = client.AddPrincipalsResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "AddPrincipals", resp, "Failure responding to request")
            }

    return
    }

    // AddPrincipalsPreparer prepares the AddPrincipals request.
    func (client DatabasesClient) AddPrincipalsPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, databasePrincipalsToAdd DatabasePrincipalListRequest) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "clusterName": autorest.Encode("path",clusterName),
            "databaseName": autorest.Encode("path",databaseName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-11-09"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/addPrincipals",pathParameters),
    autorest.WithJSON(databasePrincipalsToAdd),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // AddPrincipalsSender sends the AddPrincipals request. The method will close the
    // http.Response Body if it receives an error.
    func (client DatabasesClient) AddPrincipalsSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// AddPrincipalsResponder handles the response to the AddPrincipals request. The method always
// closes the http.Response Body.
func (client DatabasesClient) AddPrincipalsResponder(resp *http.Response) (result DatabasePrincipalListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// CheckNameAvailability checks that the database name is valid and is not already in use.
    // Parameters:
        // resourceGroupName - the name of the resource group containing the Kusto cluster.
        // clusterName - the name of the Kusto cluster.
        // resourceName - the name of the resource.
func (client DatabasesClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, clusterName string, resourceName CheckNameRequest) (result CheckNameResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DatabasesClient.CheckNameAvailability")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: resourceName,
             Constraints: []validation.Constraint{	{Target: "resourceName.Name", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
            return result, validation.NewError("kusto.DatabasesClient", "CheckNameAvailability", err.Error())
            }

                req, err := client.CheckNameAvailabilityPreparer(ctx, resourceGroupName, clusterName, resourceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "CheckNameAvailability", nil , "Failure preparing request")
    return
    }

            resp, err := client.CheckNameAvailabilitySender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "CheckNameAvailability", resp, "Failure sending request")
            return
            }

            result, err = client.CheckNameAvailabilityResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "CheckNameAvailability", resp, "Failure responding to request")
            }

    return
    }

    // CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
    func (client DatabasesClient) CheckNameAvailabilityPreparer(ctx context.Context, resourceGroupName string, clusterName string, resourceName CheckNameRequest) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "clusterName": autorest.Encode("path",clusterName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-11-09"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/checkNameAvailability",pathParameters),
    autorest.WithJSON(resourceName),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
    // http.Response Body if it receives an error.
    func (client DatabasesClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client DatabasesClient) CheckNameAvailabilityResponder(resp *http.Response) (result CheckNameResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// CreateOrUpdate creates or updates a database.
    // Parameters:
        // resourceGroupName - the name of the resource group containing the Kusto cluster.
        // clusterName - the name of the Kusto cluster.
        // databaseName - the name of the database in the Kusto cluster.
        // parameters - the database parameters supplied to the CreateOrUpdate operation.
func (client DatabasesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters BasicDatabase) (result DatabasesCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DatabasesClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, clusterName, databaseName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            result, err = client.CreateOrUpdateSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client DatabasesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters BasicDatabase) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "clusterName": autorest.Encode("path",clusterName),
            "databaseName": autorest.Encode("path",databaseName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-11-09"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client DatabasesClient) CreateOrUpdateSender(req *http.Request) (future DatabasesCreateOrUpdateFuture, err error) {
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
func (client DatabasesClient) CreateOrUpdateResponder(resp *http.Response) (result DatabaseModel, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated,http.StatusAccepted),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete deletes the database with the given name.
    // Parameters:
        // resourceGroupName - the name of the resource group containing the Kusto cluster.
        // clusterName - the name of the Kusto cluster.
        // databaseName - the name of the database in the Kusto cluster.
func (client DatabasesClient) Delete(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result DatabasesDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DatabasesClient.Delete")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.DeletePreparer(ctx, resourceGroupName, clusterName, databaseName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "Delete", nil , "Failure preparing request")
    return
    }

            result, err = client.DeleteSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "Delete", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client DatabasesClient) DeletePreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "clusterName": autorest.Encode("path",clusterName),
            "databaseName": autorest.Encode("path",databaseName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-11-09"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client DatabasesClient) DeleteSender(req *http.Request) (future DatabasesDeleteFuture, err error) {
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
func (client DatabasesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get returns a database.
    // Parameters:
        // resourceGroupName - the name of the resource group containing the Kusto cluster.
        // clusterName - the name of the Kusto cluster.
        // databaseName - the name of the database in the Kusto cluster.
func (client DatabasesClient) Get(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result DatabaseModel, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DatabasesClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, resourceGroupName, clusterName, databaseName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client DatabasesClient) GetPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "clusterName": autorest.Encode("path",clusterName),
            "databaseName": autorest.Encode("path",databaseName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-11-09"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client DatabasesClient) GetSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DatabasesClient) GetResponder(resp *http.Response) (result DatabaseModel, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByCluster returns the list of databases of the given Kusto cluster.
    // Parameters:
        // resourceGroupName - the name of the resource group containing the Kusto cluster.
        // clusterName - the name of the Kusto cluster.
func (client DatabasesClient) ListByCluster(ctx context.Context, resourceGroupName string, clusterName string) (result DatabaseListResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DatabasesClient.ListByCluster")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListByClusterPreparer(ctx, resourceGroupName, clusterName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "ListByCluster", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByClusterSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "ListByCluster", resp, "Failure sending request")
            return
            }

            result, err = client.ListByClusterResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "ListByCluster", resp, "Failure responding to request")
            }

    return
    }

    // ListByClusterPreparer prepares the ListByCluster request.
    func (client DatabasesClient) ListByClusterPreparer(ctx context.Context, resourceGroupName string, clusterName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "clusterName": autorest.Encode("path",clusterName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-11-09"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByClusterSender sends the ListByCluster request. The method will close the
    // http.Response Body if it receives an error.
    func (client DatabasesClient) ListByClusterSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// ListByClusterResponder handles the response to the ListByCluster request. The method always
// closes the http.Response Body.
func (client DatabasesClient) ListByClusterResponder(resp *http.Response) (result DatabaseListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListPrincipals returns a list of database principals of the given Kusto cluster and database.
    // Parameters:
        // resourceGroupName - the name of the resource group containing the Kusto cluster.
        // clusterName - the name of the Kusto cluster.
        // databaseName - the name of the database in the Kusto cluster.
func (client DatabasesClient) ListPrincipals(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result DatabasePrincipalListResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DatabasesClient.ListPrincipals")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListPrincipalsPreparer(ctx, resourceGroupName, clusterName, databaseName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "ListPrincipals", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListPrincipalsSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "ListPrincipals", resp, "Failure sending request")
            return
            }

            result, err = client.ListPrincipalsResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "ListPrincipals", resp, "Failure responding to request")
            }

    return
    }

    // ListPrincipalsPreparer prepares the ListPrincipals request.
    func (client DatabasesClient) ListPrincipalsPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "clusterName": autorest.Encode("path",clusterName),
            "databaseName": autorest.Encode("path",databaseName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-11-09"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/listPrincipals",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListPrincipalsSender sends the ListPrincipals request. The method will close the
    // http.Response Body if it receives an error.
    func (client DatabasesClient) ListPrincipalsSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// ListPrincipalsResponder handles the response to the ListPrincipals request. The method always
// closes the http.Response Body.
func (client DatabasesClient) ListPrincipalsResponder(resp *http.Response) (result DatabasePrincipalListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// RemovePrincipals remove Database principals permissions.
    // Parameters:
        // resourceGroupName - the name of the resource group containing the Kusto cluster.
        // clusterName - the name of the Kusto cluster.
        // databaseName - the name of the database in the Kusto cluster.
        // databasePrincipalsToRemove - list of database principals to remove.
func (client DatabasesClient) RemovePrincipals(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, databasePrincipalsToRemove DatabasePrincipalListRequest) (result DatabasePrincipalListResult, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DatabasesClient.RemovePrincipals")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.RemovePrincipalsPreparer(ctx, resourceGroupName, clusterName, databaseName, databasePrincipalsToRemove)
    if err != nil {
    err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "RemovePrincipals", nil , "Failure preparing request")
    return
    }

            resp, err := client.RemovePrincipalsSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "RemovePrincipals", resp, "Failure sending request")
            return
            }

            result, err = client.RemovePrincipalsResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "RemovePrincipals", resp, "Failure responding to request")
            }

    return
    }

    // RemovePrincipalsPreparer prepares the RemovePrincipals request.
    func (client DatabasesClient) RemovePrincipalsPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, databasePrincipalsToRemove DatabasePrincipalListRequest) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "clusterName": autorest.Encode("path",clusterName),
            "databaseName": autorest.Encode("path",databaseName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-11-09"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/removePrincipals",pathParameters),
    autorest.WithJSON(databasePrincipalsToRemove),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // RemovePrincipalsSender sends the RemovePrincipals request. The method will close the
    // http.Response Body if it receives an error.
    func (client DatabasesClient) RemovePrincipalsSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// RemovePrincipalsResponder handles the response to the RemovePrincipals request. The method always
// closes the http.Response Body.
func (client DatabasesClient) RemovePrincipalsResponder(resp *http.Response) (result DatabasePrincipalListResult, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Update updates a database.
    // Parameters:
        // resourceGroupName - the name of the resource group containing the Kusto cluster.
        // clusterName - the name of the Kusto cluster.
        // databaseName - the name of the database in the Kusto cluster.
        // parameters - the database parameters supplied to the Update operation.
func (client DatabasesClient) Update(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters BasicDatabase) (result DatabasesUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/DatabasesClient.Update")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.UpdatePreparer(ctx, resourceGroupName, clusterName, databaseName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "Update", nil , "Failure preparing request")
    return
    }

            result, err = client.UpdateSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "kusto.DatabasesClient", "Update", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // UpdatePreparer prepares the Update request.
    func (client DatabasesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters BasicDatabase) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "clusterName": autorest.Encode("path",clusterName),
            "databaseName": autorest.Encode("path",databaseName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-11-09"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPatch(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client DatabasesClient) UpdateSender(req *http.Request) (future DatabasesUpdateFuture, err error) {
            var resp *http.Response
            resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
            if err != nil {
            return
            }
            future.Future, err = azure.NewFutureFromResponse(resp)
            return
            }

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client DatabasesClient) UpdateResponder(resp *http.Response) (result DatabaseModel, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated,http.StatusAccepted),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

