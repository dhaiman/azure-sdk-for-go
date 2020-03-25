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

// ProtectionContainerRefreshOperationResultsClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type ProtectionContainerRefreshOperationResultsClient struct {
    BaseClient
}
// NewProtectionContainerRefreshOperationResultsClient creates an instance of the
// ProtectionContainerRefreshOperationResultsClient client.
func NewProtectionContainerRefreshOperationResultsClient(subscriptionID string) ProtectionContainerRefreshOperationResultsClient {
    return NewProtectionContainerRefreshOperationResultsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProtectionContainerRefreshOperationResultsClientWithBaseURI creates an instance of the
// ProtectionContainerRefreshOperationResultsClient client using a custom endpoint.  Use this when interacting with an
// Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewProtectionContainerRefreshOperationResultsClientWithBaseURI(baseURI string, subscriptionID string) ProtectionContainerRefreshOperationResultsClient {
        return ProtectionContainerRefreshOperationResultsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Get provides the result of the refresh operation triggered by the BeginRefresh operation.
    // Parameters:
        // vaultName - the name of the Recovery Services vault.
        // resourceGroupName - the name of the resource group associated with the Recovery Services vault.
        // fabricName - the fabric name associated with the container.
        // operationID - the operation ID used for this GET operation.
func (client ProtectionContainerRefreshOperationResultsClient) Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, operationID string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ProtectionContainerRefreshOperationResultsClient.Get")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.GetPreparer(ctx, vaultName, resourceGroupName, fabricName, operationID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "backup.ProtectionContainerRefreshOperationResultsClient", "Get", nil , "Failure preparing request")
    return
    }

            resp, err := client.GetSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "backup.ProtectionContainerRefreshOperationResultsClient", "Get", resp, "Failure sending request")
            return
            }

            result, err = client.GetResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "backup.ProtectionContainerRefreshOperationResultsClient", "Get", resp, "Failure responding to request")
            }

    return
    }

    // GetPreparer prepares the Get request.
    func (client ProtectionContainerRefreshOperationResultsClient) GetPreparer(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, operationID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "fabricName": autorest.Encode("path",fabricName),
            "operationId": autorest.Encode("path",operationID),
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
    autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/operationResults/{operationId}",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client ProtectionContainerRefreshOperationResultsClient) GetSender(req *http.Request) (*http.Response, error) {
            return client.Send(req, azure.DoRetryWithRegistration(client.Client))
            }

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProtectionContainerRefreshOperationResultsClient) GetResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
    autorest.ByClosing())
    result.Response = resp
        return
    }

