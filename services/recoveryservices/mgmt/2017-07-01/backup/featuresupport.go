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

// FeatureSupportClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type FeatureSupportClient struct {
    BaseClient
}
// NewFeatureSupportClient creates an instance of the FeatureSupportClient client.
func NewFeatureSupportClient(subscriptionID string) FeatureSupportClient {
    return NewFeatureSupportClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFeatureSupportClientWithBaseURI creates an instance of the FeatureSupportClient client.
    func NewFeatureSupportClientWithBaseURI(baseURI string, subscriptionID string) FeatureSupportClient {
        return FeatureSupportClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Validate sends the validate request.
    // Parameters:
        // azureRegion - azure region to hit Api
        // parameters - feature support request object
func (client FeatureSupportClient) Validate(ctx context.Context, azureRegion string, parameters BasicFeatureSupportRequest) (result AzureVMResourceFeatureSupportResponse, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/FeatureSupportClient.Validate")
        defer func() {
            sc := -1
            if result.Response.Response != nil {
                sc = result.Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        req, err := client.ValidatePreparer(ctx, azureRegion, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "backup.FeatureSupportClient", "Validate", nil , "Failure preparing request")
    return
    }

            resp, err := client.ValidateSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            err = autorest.NewErrorWithError(err, "backup.FeatureSupportClient", "Validate", resp, "Failure sending request")
            return
            }

            result, err = client.ValidateResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "backup.FeatureSupportClient", "Validate", resp, "Failure responding to request")
            }

    return
    }

    // ValidatePreparer prepares the Validate request.
    func (client FeatureSupportClient) ValidatePreparer(ctx context.Context, azureRegion string, parameters BasicFeatureSupportRequest) (*http.Request, error) {
            pathParameters := map[string]interface{} {
            "azureRegion": autorest.Encode("path",azureRegion),
            "subscriptionId": autorest.Encode("path",client.SubscriptionID),
            }

                        const APIVersion = "2017-07-01"
        queryParameters := map[string]interface{} {
        "api-version": APIVersion,
        }

        preparer := autorest.CreatePreparer(
    autorest.AsContentType("application/json; charset=utf-8"),
    autorest.AsPost(),
    autorest.WithBaseURL(client.BaseURI),
    autorest.WithPathParameters("/Subscriptions/{subscriptionId}/providers/Microsoft.RecoveryServices/locations/{azureRegion}/backupValidateFeatures",pathParameters),
    autorest.WithJSON(parameters),
    autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ValidateSender sends the Validate request. The method will close the
    // http.Response Body if it receives an error.
    func (client FeatureSupportClient) ValidateSender(req *http.Request) (*http.Response, error) {
        sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
            return autorest.SendWithSender(client, req, sd...)
            }

// ValidateResponder handles the response to the Validate request. The method always
// closes the http.Response Body.
func (client FeatureSupportClient) ValidateResponder(resp *http.Response) (result AzureVMResourceFeatureSupportResponse, err error) {
    err = autorest.Respond(
    resp,
    client.ByInspecting(),
    azure.WithErrorUnlessStatusCode(http.StatusOK),
    autorest.ByUnmarshallingJSON(&result),
    autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
        return
    }

