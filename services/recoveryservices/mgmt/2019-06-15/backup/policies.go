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

// PoliciesClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type PoliciesClient struct {
    BaseClient
}
// NewPoliciesClient creates an instance of the PoliciesClient client.
func NewPoliciesClient(subscriptionID string) PoliciesClient {
    return NewPoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPoliciesClientWithBaseURI creates an instance of the PoliciesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewPoliciesClientWithBaseURI(baseURI string, subscriptionID string) PoliciesClient {
        return PoliciesClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// List lists of backup policies associated with Recovery Services Vault. API provides pagination parameters to fetch
// scoped results.
    // Parameters:
        // vaultName - the name of the recovery services vault.
        // resourceGroupName - the name of the resource group where the recovery services vault is present.
        // filter - oData filter options.
func (client PoliciesClient) List(ctx context.Context, vaultName string, resourceGroupName string, filter string) (result ProtectionPolicyResourceListPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PoliciesClient.List")
        defer func() {
            sc := -1
            if result.pprl.Response.Response != nil {
                sc = result.pprl.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
                result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, vaultName, resourceGroupName, filter)
    if err != nil {
    err = autorest.NewErrorWithError(err, "backup.PoliciesClient", "List", nil , "Failure preparing request")
    return
    }

            resp, err := client.ListSender(req)
            if err != nil {
            result.pprl.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "backup.PoliciesClient", "List", resp, "Failure sending request")
            return
            }

            result.pprl, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "backup.PoliciesClient", "List", resp, "Failure responding to request")
            }

    return
    }

    // ListPreparer prepares the List request.
    func (client PoliciesClient) ListPreparer(ctx context.Context, vaultName string, resourceGroupName string, filter string) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "resourceGroupName": autorest.Encode("path",resourceGroupName),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            "vaultName": autorest.Encode("path",vaultName),
            }

                        const APIVersion = "2019-06-15"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }
            if len(filter) > 0 {
            queryParameters["$filter"] = autorest.Encode("query",filter)
            }

        preparer := autorest.CreatePreparer(
    autorest.AsGet(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupPolicies",pathParameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client PoliciesClient) ListSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client PoliciesClient) ListResponder(resp *http.Response) (result ProtectionPolicyResourceList, err error) {
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
            func (client PoliciesClient) listNextResults(ctx context.Context, lastResults ProtectionPolicyResourceList) (result ProtectionPolicyResourceList, err error) {
            req, err := lastResults.protectionPolicyResourceListPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "backup.PoliciesClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "backup.PoliciesClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "backup.PoliciesClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

    // ListComplete enumerates all values, automatically crossing page boundaries as required.
    func (client PoliciesClient) ListComplete(ctx context.Context, vaultName string, resourceGroupName string, filter string) (result ProtectionPolicyResourceListIterator, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/PoliciesClient.List")
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

