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

// RecoveryPointsClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type RecoveryPointsClient struct {
	BaseClient
}

// NewRecoveryPointsClient creates an instance of the RecoveryPointsClient client.
func NewRecoveryPointsClient(subscriptionID string) RecoveryPointsClient {
	return NewRecoveryPointsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRecoveryPointsClientWithBaseURI creates an instance of the RecoveryPointsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRecoveryPointsClientWithBaseURI(baseURI string, subscriptionID string) RecoveryPointsClient {
	return RecoveryPointsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetAccessToken sends the get access token request.
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
// fabricName - fabric name associated with the container.
// containerName - name of the container.
// protectedItemName - name of the Protected Item.
// recoveryPointID - recovery Point Id
func (client RecoveryPointsClient) GetAccessToken(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string) (result CrrAccessTokenResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecoveryPointsClient.GetAccessToken")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetAccessTokenPreparer(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, recoveryPointID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.RecoveryPointsClient", "GetAccessToken", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAccessTokenSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "backup.RecoveryPointsClient", "GetAccessToken", resp, "Failure sending request")
		return
	}

	result, err = client.GetAccessTokenResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.RecoveryPointsClient", "GetAccessToken", resp, "Failure responding to request")
	}

	return
}

// GetAccessTokenPreparer prepares the GetAccessToken request.
func (client RecoveryPointsClient) GetAccessTokenPreparer(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerName":     autorest.Encode("path", containerName),
		"fabricName":        autorest.Encode("path", fabricName),
		"protectedItemName": autorest.Encode("path", protectedItemName),
		"recoveryPointId":   autorest.Encode("path", recoveryPointID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2018-12-20"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/recoveryPoints/{recoveryPointId}/accessToken", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAccessTokenSender sends the GetAccessToken request. The method will close the
// http.Response Body if it receives an error.
func (client RecoveryPointsClient) GetAccessTokenSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetAccessTokenResponder handles the response to the GetAccessToken request. The method always
// closes the http.Response Body.
func (client RecoveryPointsClient) GetAccessTokenResponder(resp *http.Response) (result CrrAccessTokenResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
