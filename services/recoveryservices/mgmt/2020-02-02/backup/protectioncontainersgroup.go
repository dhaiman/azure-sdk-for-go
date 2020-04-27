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

// ProtectionContainersGroupClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type ProtectionContainersGroupClient struct {
	BaseClient
}

// NewProtectionContainersGroupClient creates an instance of the ProtectionContainersGroupClient client.
func NewProtectionContainersGroupClient(subscriptionID string) ProtectionContainersGroupClient {
	return NewProtectionContainersGroupClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProtectionContainersGroupClientWithBaseURI creates an instance of the ProtectionContainersGroupClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewProtectionContainersGroupClientWithBaseURI(baseURI string, subscriptionID string) ProtectionContainersGroupClient {
	return ProtectionContainersGroupClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists the containers registered to Recovery Services Vault.
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
// filter - oData filter options.
func (client ProtectionContainersGroupClient) List(ctx context.Context, vaultName string, resourceGroupName string, filter string) (result ProtectionContainerResourceListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProtectionContainersGroupClient.List")
		defer func() {
			sc := -1
			if result.pcrl.Response.Response != nil {
				sc = result.pcrl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, vaultName, resourceGroupName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ProtectionContainersGroupClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.pcrl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "backup.ProtectionContainersGroupClient", "List", resp, "Failure sending request")
		return
	}

	result.pcrl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ProtectionContainersGroupClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ProtectionContainersGroupClient) ListPreparer(ctx context.Context, vaultName string, resourceGroupName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupProtectionContainers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ProtectionContainersGroupClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ProtectionContainersGroupClient) ListResponder(resp *http.Response) (result ProtectionContainerResourceList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ProtectionContainersGroupClient) listNextResults(ctx context.Context, lastResults ProtectionContainerResourceList) (result ProtectionContainerResourceList, err error) {
	req, err := lastResults.protectionContainerResourceListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "backup.ProtectionContainersGroupClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "backup.ProtectionContainersGroupClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ProtectionContainersGroupClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProtectionContainersGroupClient) ListComplete(ctx context.Context, vaultName string, resourceGroupName string, filter string) (result ProtectionContainerResourceListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProtectionContainersGroupClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, vaultName, resourceGroupName, filter)
	return
}
