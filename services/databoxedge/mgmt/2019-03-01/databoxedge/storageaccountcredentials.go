package databoxedge

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

// StorageAccountCredentialsClient is the client for the StorageAccountCredentials methods of the Databoxedge service.
type StorageAccountCredentialsClient struct {
    BaseClient
}
// NewStorageAccountCredentialsClient creates an instance of the StorageAccountCredentialsClient client.
func NewStorageAccountCredentialsClient(subscriptionID string) StorageAccountCredentialsClient {
    return NewStorageAccountCredentialsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewStorageAccountCredentialsClientWithBaseURI creates an instance of the StorageAccountCredentialsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
    func NewStorageAccountCredentialsClientWithBaseURI(baseURI string, subscriptionID string) StorageAccountCredentialsClient {
        return StorageAccountCredentialsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates or updates the storage account credential.
    // Parameters:
        // deviceName - the device name.
        // name - the storage account credential name.
        // storageAccountCredential - the storage account credential.
        // resourceGroupName - the resource group name.
func (client StorageAccountCredentialsClient) CreateOrUpdate(ctx context.Context, deviceName string, name string, storageAccountCredential StorageAccountCredential, resourceGroupName string) (result StorageAccountCredentialsCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/StorageAccountCredentialsClient.CreateOrUpdate")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
            if err := validation.Validate([]validation.Validation{
            { TargetValue: storageAccountCredential,
             Constraints: []validation.Constraint{	{Target: "storageAccountCredential.StorageAccountCredentialProperties", Name: validation.Null, Rule: true ,
            Chain: []validation.Constraint{	{Target: "storageAccountCredential.StorageAccountCredentialProperties.Alias", Name: validation.Null, Rule: true, Chain: nil },
            	{Target: "storageAccountCredential.StorageAccountCredentialProperties.AccountKey", Name: validation.Null, Rule: false ,
            Chain: []validation.Constraint{	{Target: "storageAccountCredential.StorageAccountCredentialProperties.AccountKey.Value", Name: validation.Null, Rule: true, Chain: nil },
            }},
            }}}}}); err != nil {
            return result, validation.NewError("databoxedge.StorageAccountCredentialsClient", "CreateOrUpdate", err.Error())
            }

                req, err := client.CreateOrUpdatePreparer(ctx, deviceName, name, storageAccountCredential, resourceGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "databoxedge.StorageAccountCredentialsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

            result, err = client.CreateOrUpdateSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "databoxedge.StorageAccountCredentialsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client StorageAccountCredentialsClient) CreateOrUpdatePreparer(ctx context.Context, deviceName string, name string, storageAccountCredential StorageAccountCredential, resourceGroupName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "deviceName": autorest.Encode("path",deviceName),
            "name": autorest.Encode("path",name),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-03-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPut(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials/{name}",pathParameters),
    autorest.WithJSON(storageAccountCredential),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client StorageAccountCredentialsClient) CreateOrUpdateSender(req *http.Request) (future StorageAccountCredentialsCreateOrUpdateFuture, err error) {
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
func (client StorageAccountCredentialsClient) CreateOrUpdateResponder(resp *http.Response) (result StorageAccountCredential, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Delete deletes the storage account credential.
    // Parameters:
        // deviceName - the device name.
        // name - the storage account credential name.
        // resourceGroupName - the resource group name.
func (client StorageAccountCredentialsClient) Delete(ctx context.Context, deviceName string, name string, resourceGroupName string) (result StorageAccountCredentialsDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/StorageAccountCredentialsClient.Delete")
        defer func() {
            sc := -1
            if result.Response() != nil {
                sc = result.Response().StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.DeletePreparer(ctx, deviceName, name, resourceGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "databoxedge.StorageAccountCredentialsClient", "Delete", nil , "Failure preparing request")
    return
    }

            result, err = client.DeleteSender(req)
            if err != nil {
            err = autorest.NewErrorWithError(err, "databoxedge.StorageAccountCredentialsClient", "Delete", result.Response(), "Failure sending request")
            return
            }

    return
    }

    // DeletePreparer prepares the Delete request.
    func (client StorageAccountCredentialsClient) DeletePreparer(ctx context.Context, deviceName string, name string, resourceGroupName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "deviceName": autorest.Encode("path",deviceName),
            "name": autorest.Encode("path",name),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-03-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials/{name}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client StorageAccountCredentialsClient) DeleteSender(req *http.Request) (future StorageAccountCredentialsDeleteFuture, err error) {
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
func (client StorageAccountCredentialsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Get gets the properties of the specified storage account credential.
    // Parameters:
        // deviceName - the device name.
        // name - the storage account credential name.
        // resourceGroupName - the resource group name.
func (client StorageAccountCredentialsClient) Get(ctx context.Context, deviceName string, name string, resourceGroupName string) (result StorageAccountCredential, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/StorageAccountCredentialsClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, deviceName, name, resourceGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "databoxedge.StorageAccountCredentialsClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "databoxedge.StorageAccountCredentialsClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "databoxedge.StorageAccountCredentialsClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client StorageAccountCredentialsClient) GetPreparer(ctx context.Context, deviceName string, name string, resourceGroupName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "deviceName": autorest.Encode("path",deviceName),
            "name": autorest.Encode("path",name),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-03-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials/{name}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client StorageAccountCredentialsClient) GetSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client StorageAccountCredentialsClient) GetResponder(resp *http.Response) (result StorageAccountCredential, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// ListByDataBoxEdgeDevice sends the list by data box edge device request.
    // Parameters:
        // deviceName - the device name.
        // resourceGroupName - the resource group name.
func (client StorageAccountCredentialsClient) ListByDataBoxEdgeDevice(ctx context.Context, deviceName string, resourceGroupName string) (result StorageAccountCredentialListPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/StorageAccountCredentialsClient.ListByDataBoxEdgeDevice")
        defer func() {
            sc := -1
            if result.sacl.Response.Response != nil {
                sc = result.sacl.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listByDataBoxEdgeDeviceNextResults
    req, err := client.ListByDataBoxEdgeDevicePreparer(ctx, deviceName, resourceGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "databoxedge.StorageAccountCredentialsClient", "ListByDataBoxEdgeDevice", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListByDataBoxEdgeDeviceSender(req)
            if err != nil {
            result.sacl.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "databoxedge.StorageAccountCredentialsClient", "ListByDataBoxEdgeDevice", resp, "Failure sending request")
            return
            }

            result.sacl, err = client.ListByDataBoxEdgeDeviceResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "databoxedge.StorageAccountCredentialsClient", "ListByDataBoxEdgeDevice", resp, "Failure responding to request")
            }

    return
    }

    // ListByDataBoxEdgeDevicePreparer prepares the ListByDataBoxEdgeDevice request.
    func (client StorageAccountCredentialsClient) ListByDataBoxEdgeDevicePreparer(ctx context.Context, deviceName string, resourceGroupName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "deviceName": autorest.Encode("path",deviceName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2019-03-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/storageAccountCredentials",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByDataBoxEdgeDeviceSender sends the ListByDataBoxEdgeDevice request. The method will close the
    // http.Response Body if it receives an error.
    func (client StorageAccountCredentialsClient) ListByDataBoxEdgeDeviceSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// ListByDataBoxEdgeDeviceResponder handles the response to the ListByDataBoxEdgeDevice request. The method always
// closes the http.Response Body.
func (client StorageAccountCredentialsClient) ListByDataBoxEdgeDeviceResponder(resp *http.Response) (result StorageAccountCredentialList, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

            // listByDataBoxEdgeDeviceNextResults retrieves the next set of results, if any.
            func (client StorageAccountCredentialsClient) listByDataBoxEdgeDeviceNextResults(ctx context.Context, lastResults StorageAccountCredentialList) (result StorageAccountCredentialList, err error) {
            req, err := lastResults.storageAccountCredentialListPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "databoxedge.StorageAccountCredentialsClient", "listByDataBoxEdgeDeviceNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByDataBoxEdgeDeviceSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "databoxedge.StorageAccountCredentialsClient", "listByDataBoxEdgeDeviceNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByDataBoxEdgeDeviceResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "databoxedge.StorageAccountCredentialsClient", "listByDataBoxEdgeDeviceNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListByDataBoxEdgeDeviceComplete enumerates all values, automatically crossing page boundaries as required.
    func (client StorageAccountCredentialsClient) ListByDataBoxEdgeDeviceComplete(ctx context.Context, deviceName string, resourceGroupName string) (result StorageAccountCredentialListIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/StorageAccountCredentialsClient.ListByDataBoxEdgeDevice")
            defer func() {
                sc := -1
                if result.Response().Response.Response != nil {
                    sc = result.page.Response().Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
     }
        result.page, err = client.ListByDataBoxEdgeDevice(ctx, deviceName, resourceGroupName)
                return
        }

