package backup

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
)

// ProtectionContainersClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type ProtectionContainersClient struct {
    BaseClient
}
// NewProtectionContainersClient creates an instance of the ProtectionContainersClient client.
func NewProtectionContainersClient(subscriptionID string) ProtectionContainersClient {
    return NewProtectionContainersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProtectionContainersClientWithBaseURI creates an instance of the ProtectionContainersClient client.
    func NewProtectionContainersClientWithBaseURI(baseURI string, subscriptionID string) ProtectionContainersClient {
        return ProtectionContainersClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Get gets details of the specific container registered to your Recovery Services vault.
    // Parameters:
        // vaultName - the name of the Recovery Services vault.
        // resourceGroupName - the name of the resource group associated with the Recovery Services vault.
        // fabricName - the fabric name associated with the container.
        // containerName - the container name used for this GET operation.
func (client ProtectionContainersClient) Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string) (result ProtectionContainerResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ProtectionContainersClient.Get")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, vaultName, resourceGroupName, fabricName, containerName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "backup.ProtectionContainersClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "backup.ProtectionContainersClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "backup.ProtectionContainersClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client ProtectionContainersClient) GetPreparer(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "containerName": autorest.Encode("path",containerName),
            "fabricName": autorest.Encode("path",fabricName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "vaultName": autorest.Encode("path",vaultName),
            }

                        const APIVersion = "2016-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client ProtectionContainersClient) GetSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProtectionContainersClient) GetResponder(resp *http.Response) (result ProtectionContainerResource, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// List lists the containers registered to the Recovery Services vault.
    // Parameters:
        // vaultName - the name of the Recovery Services vault.
        // resourceGroupName - the name of the resource group associated with the Recovery Services vault.
        // filter - the following equation is used to sort or filter the containers registered to the vault.
        // providerType eq {AzureIaasVM, MAB, DPM, AzureBackupServer, AzureSql} and status eq {Unknown, NotRegistered,
        // Registered, Registering} and friendlyName eq {containername} and backupManagementType eq {AzureIaasVM, MAB,
        // DPM, AzureBackupServer, AzureSql}.
func (client ProtectionContainersClient) List(ctx context.Context, vaultName string, resourceGroupName string, filter string) (result ProtectionContainerResourceList, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ProtectionContainersClient.List")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ListPreparer(ctx, vaultName, resourceGroupName, filter)
    if err != nil {
    err = autorest.NewErrorWithError(err, "backup.ProtectionContainersClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "backup.ProtectionContainersClient", "List", resp, "Failure sending request")
            return
            }

            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "backup.ProtectionContainersClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client ProtectionContainersClient) ListPreparer(ctx context.Context, vaultName string, resourceGroupName string, filter string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "vaultName": autorest.Encode("path",vaultName),
            }

                        const APIVersion = "2016-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }
            if len(filter) > 0 {
            queryParameters["$filter"] = autorest.Encode("query",filter)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupProtectionContainers",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client ProtectionContainersClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ProtectionContainersClient) ListResponder(resp *http.Response) (result ProtectionContainerResourceList, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

// Refresh discovers the containers in the subscription that can be protected in a Recovery Services vault. This is an
// asynchronous operation. To learn the status of the operation, use the GetRefreshOperationResult API.
    // Parameters:
        // vaultName - the name of the Recovery Services vault.
        // resourceGroupName - the name of the resource group associated with the Recovery Services vault.
        // fabricName - the fabric name associated with the container.
func (client ProtectionContainersClient) Refresh(ctx context.Context, vaultName string, resourceGroupName string, fabricName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ProtectionContainersClient.Refresh")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.RefreshPreparer(ctx, vaultName, resourceGroupName, fabricName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "backup.ProtectionContainersClient", "Refresh", nil , "Failure preparing request")
    return
    }

            resp, err := client.RefreshSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "backup.ProtectionContainersClient", "Refresh", resp, "Failure sending request")
            return
            }

            result, err = client.RefreshResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "backup.ProtectionContainersClient", "Refresh", resp, "Failure responding to request")
            }

    return
    }

    // RefreshPreparer prepares the Refresh request.
    func (client ProtectionContainersClient) RefreshPreparer(ctx context.Context, vaultName string, resourceGroupName string, fabricName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "fabricName": autorest.Encode("path",fabricName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "vaultName": autorest.Encode("path",vaultName),
            }

                        const APIVersion = "2016-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/refreshContainers",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // RefreshSender sends the Refresh request. The method will close the
    // http.Response Body if it receives an error.
    func (client ProtectionContainersClient) RefreshSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// RefreshResponder handles the response to the Refresh request. The method always
// closes the http.Response Body.
func (client ProtectionContainersClient) RefreshResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Unregister unregisters the given container from your Recovery Services vault.
    // Parameters:
        // resourceGroupName - the name of the resource group associated with the Recovery Services vault.
        // vaultName - the name of the Recovery Services vault.
        // identityName - name of the protection container to unregister.
func (client ProtectionContainersClient) Unregister(ctx context.Context, resourceGroupName string, vaultName string, identityName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ProtectionContainersClient.Unregister")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.UnregisterPreparer(ctx, resourceGroupName, vaultName, identityName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "backup.ProtectionContainersClient", "Unregister", nil , "Failure preparing request")
    return
    }

            resp, err := client.UnregisterSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "backup.ProtectionContainersClient", "Unregister", resp, "Failure sending request")
            return
            }

            result, err = client.UnregisterResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "backup.ProtectionContainersClient", "Unregister", resp, "Failure responding to request")
            }

    return
    }

    // UnregisterPreparer prepares the Unregister request.
    func (client ProtectionContainersClient) UnregisterPreparer(ctx context.Context, resourceGroupName string, vaultName string, identityName string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "identityName": autorest.Encode("path",identityName),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "vaultName": autorest.Encode("path",vaultName),
            }

                        const APIVersion = "2016-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsDelete(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/registeredIdentities/{identityName}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UnregisterSender sends the Unregister request. The method will close the
    // http.Response Body if it receives an error.
    func (client ProtectionContainersClient) UnregisterSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// UnregisterResponder handles the response to the Unregister request. The method always
// closes the http.Response Body.
func (client ProtectionContainersClient) UnregisterResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

