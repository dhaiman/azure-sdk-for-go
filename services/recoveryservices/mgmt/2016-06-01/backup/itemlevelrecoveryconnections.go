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

// ItemLevelRecoveryConnectionsClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type ItemLevelRecoveryConnectionsClient struct {
    BaseClient
}
// NewItemLevelRecoveryConnectionsClient creates an instance of the ItemLevelRecoveryConnectionsClient client.
func NewItemLevelRecoveryConnectionsClient(subscriptionID string) ItemLevelRecoveryConnectionsClient {
    return NewItemLevelRecoveryConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewItemLevelRecoveryConnectionsClientWithBaseURI creates an instance of the ItemLevelRecoveryConnectionsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
    func NewItemLevelRecoveryConnectionsClientWithBaseURI(baseURI string, subscriptionID string) ItemLevelRecoveryConnectionsClient {
        return ItemLevelRecoveryConnectionsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Provision provisions a script which invokes an iSCSI connection to the backup data. Executing this script opens File
// Explorer which displays the recoverable files and folders. This is an asynchronous operation. To get the
// provisioning status, call GetProtectedItemOperationResult API.
    // Parameters:
        // vaultName - the name of the Recovery Services vault.
        // resourceGroupName - the name of the resource group associated with the Recovery Services vault.
        // fabricName - the fabric name associated with the backup items.
        // containerName - the container name associated with the backup items.
        // protectedItemName - the name of the backup item whose files or folders are to be restored.
        // recoveryPointID - the recovery point ID for backup data. The iSCSI connection will be provisioned for this
        // backup data.
        // resourceILRRequest - the resource Item Level Recovery (ILR) request.
func (client ItemLevelRecoveryConnectionsClient) Provision(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, resourceILRRequest ILRRequestResource) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ItemLevelRecoveryConnectionsClient.Provision")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ProvisionPreparer(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, recoveryPointID, resourceILRRequest)
    if err != nil {
    err = autorest.NewErrorWithError(err, "backup.ItemLevelRecoveryConnectionsClient", "Provision", nil , "Failure preparing request")
    return
    }

            resp, err := client.ProvisionSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "backup.ItemLevelRecoveryConnectionsClient", "Provision", resp, "Failure sending request")
            return
            }

            result, err = client.ProvisionResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "backup.ItemLevelRecoveryConnectionsClient", "Provision", resp, "Failure responding to request")
            }

    return
    }

    // ProvisionPreparer prepares the Provision request.
    func (client ItemLevelRecoveryConnectionsClient) ProvisionPreparer(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, resourceILRRequest ILRRequestResource) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "containerName": autorest.Encode("path",containerName),
            "fabricName": autorest.Encode("path",fabricName),
            "protectedItemName": autorest.Encode("path",protectedItemName),
            "recoveryPointId": autorest.Encode("path",recoveryPointID),
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "vaultName": autorest.Encode("path",vaultName),
            }

                        const APIVersion = "2016-06-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/recoveryPoints/{recoveryPointId}/provisionInstantItemRecovery",pathParameters),
    autorest.WithJSON(resourceILRRequest),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ProvisionSender sends the Provision request. The method will close the
    // http.Response Body if it receives an error.
    func (client ItemLevelRecoveryConnectionsClient) ProvisionSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ProvisionResponder handles the response to the Provision request. The method always
// closes the http.Response Body.
func (client ItemLevelRecoveryConnectionsClient) ProvisionResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
    autorest.ByClosing())
    result.Response = resp
        return
    }

// Revoke revokes an iSCSI connection which can be used to download a script. Executing this script opens a file
// explorer displaying all recoverable files and folders. This is an asynchronous operation.
    // Parameters:
        // vaultName - the name of the Recovery Services vault.
        // resourceGroupName - the name of the resource group associated with the Recovery Services vault.
        // fabricName - the fabric name associated with the backup items. The value allowed is Azure.
        // containerName - the container name associated with the backup items.
        // protectedItemName - the name of the backup items whose files or folders will be restored.
        // recoveryPointID - the string that identifies the recovery point. The iSCSI connection will be revoked for
        // this protected data.
func (client ItemLevelRecoveryConnectionsClient) Revoke(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ItemLevelRecoveryConnectionsClient.Revoke")
        defer func() {
            sc := -1
            if result.Response != nil {
                sc = result.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.RevokePreparer(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, recoveryPointID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "backup.ItemLevelRecoveryConnectionsClient", "Revoke", nil , "Failure preparing request")
    return
    }

            resp, err := client.RevokeSender(req)
            if err != nil {
            result.Response = resp
            err = autorest.NewErrorWithError(err, "backup.ItemLevelRecoveryConnectionsClient", "Revoke", resp, "Failure sending request")
            return
            }

            result, err = client.RevokeResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "backup.ItemLevelRecoveryConnectionsClient", "Revoke", resp, "Failure responding to request")
            }

    return
    }

    // RevokePreparer prepares the Revoke request.
    func (client ItemLevelRecoveryConnectionsClient) RevokePreparer(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "containerName": autorest.Encode("path",containerName),
            "fabricName": autorest.Encode("path",fabricName),
            "protectedItemName": autorest.Encode("path",protectedItemName),
            "recoveryPointId": autorest.Encode("path",recoveryPointID),
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
    autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/recoveryPoints/{recoveryPointId}/revokeInstantItemRecovery",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // RevokeSender sends the Revoke request. The method will close the
    // http.Response Body if it receives an error.
    func (client ItemLevelRecoveryConnectionsClient) RevokeSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// RevokeResponder handles the response to the Revoke request. The method always
// closes the http.Response Body.
func (client ItemLevelRecoveryConnectionsClient) RevokeResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
    autorest.ByClosing())
    result.Response = resp
        return
    }

