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
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// BackupsClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type BackupsClient struct {
	BaseClient
}

// NewBackupsClient creates an instance of the BackupsClient client.
func NewBackupsClient(subscriptionID string) BackupsClient {
	return NewBackupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBackupsClientWithBaseURI creates an instance of the BackupsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewBackupsClientWithBaseURI(baseURI string, subscriptionID string) BackupsClient {
	return BackupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Trigger triggers the backup job for the specified backup item. This is an asynchronous operation. To know the status
// of the operation, call GetProtectedItemOperationResult API.
// Parameters:
// vaultName - the name of the Recovery Services vault.
// resourceGroupName - the name of the resource group associated with the Recovery Services vault.
// fabricName - the fabric name associated with the backup item.
// containerName - the container name associated with the backup item.
// protectedItemName - the name of backup item used in this POST operation.
// resourceBackupRequest - the resource backup request.
func (client BackupsClient) Trigger(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, resourceBackupRequest RequestResource) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BackupsClient.Trigger")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.TriggerPreparer(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, resourceBackupRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.BackupsClient", "Trigger", nil, "Failure preparing request")
		return
	}

	resp, err := client.TriggerSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "backup.BackupsClient", "Trigger", resp, "Failure sending request")
		return
	}

	result, err = client.TriggerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.BackupsClient", "Trigger", resp, "Failure responding to request")
	}

	return
}

// TriggerPreparer prepares the Trigger request.
func (client BackupsClient) TriggerPreparer(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, resourceBackupRequest RequestResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerName":     autorest.Encode("path", containerName),
		"fabricName":        autorest.Encode("path", fabricName),
		"protectedItemName": autorest.Encode("path", protectedItemName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/backup", pathParameters),
		autorest.WithJSON(resourceBackupRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TriggerSender sends the Trigger request. The method will close the
// http.Response Body if it receives an error.
func (client BackupsClient) TriggerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// TriggerResponder handles the response to the Trigger request. The method always
// closes the http.Response Body.
func (client BackupsClient) TriggerResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
