// Package cdn implements the Azure ARM Cdn service API version 2017-04-02.
//
// Cdn Management Client
package cdn

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
    "github.com/Azure/go-autorest/autorest/validation"
    "github.com/Azure/go-autorest/tracing"
    "net/http"
)

const (
// DefaultBaseURI is the default URI used for the service Cdn
DefaultBaseURI = "https://management.azure.com")

// BaseClient is the base client for Cdn.
type BaseClient struct {
    autorest.Client
    BaseURI string
            SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string)BaseClient {
    return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
    return BaseClient{
        Client: autorest.NewClientWithUserAgent(UserAgent()),
        BaseURI: baseURI,
                SubscriptionID: subscriptionID,
    }
}

    // CheckNameAvailability check the availability of a resource name. This is needed for resources where name is globally
    // unique, such as a CDN endpoint.
        // Parameters:
            // checkNameAvailabilityInput - input to check.
    func (client BaseClient) CheckNameAvailability(ctx context.Context, checkNameAvailabilityInput CheckNameAvailabilityInput) (result CheckNameAvailabilityOutput, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.CheckNameAvailability")
            defer func() {
                sc := -1
                if result.Response.Response != nil {
                    sc = result.Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
                if err := validation.Validate([]validation.Validation{
                { TargetValue: checkNameAvailabilityInput,
                 Constraints: []validation.Constraint{	{Target: "checkNameAvailabilityInput.Name", Name: validation.Null, Rule: true, Chain: nil },
                	{Target: "checkNameAvailabilityInput.Type", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
                return result, validation.NewError("cdn.BaseClient", "CheckNameAvailability", err.Error())
                }

                    req, err := client.CheckNameAvailabilityPreparer(ctx, checkNameAvailabilityInput)
        if err != nil {
        err = autorest.NewErrorWithError(err, "cdn.BaseClient", "CheckNameAvailability", nil , "Failure preparing request")
        return
        }

                resp, err := client.CheckNameAvailabilitySender(req)
                if err != nil {
                result.Response = autorest.Response{Response: resp}
                err = autorest.NewErrorWithError(err, "cdn.BaseClient", "CheckNameAvailability", resp, "Failure sending request")
                return
                }

                result, err = client.CheckNameAvailabilityResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "cdn.BaseClient", "CheckNameAvailability", resp, "Failure responding to request")
                }

        return
        }

        // CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
        func (client BaseClient) CheckNameAvailabilityPreparer(ctx context.Context, checkNameAvailabilityInput CheckNameAvailabilityInput) (*http.Request, error) {
                        const APIVersion = "2017-04-02"
            queryParameters := map[string]interface{} {
            "api-version": APIVersion,
            }

            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json; charset=utf-8"),
        autorest.AsPost(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPath("/providers/Microsoft.Cdn/checkNameAvailability"),
        autorest.WithJSON(checkNameAvailabilityInput),
        autorest.WithQueryParameters(queryParameters))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
    // closes the http.Response Body.
    func (client BaseClient) CheckNameAvailabilityResponder(resp *http.Response) (result CheckNameAvailabilityOutput, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByUnmarshallingJSON(&result),
        autorest.ByClosing())
        result.Response = autorest.Response{Response: resp}
            return
        }

    // ValidateProbe check if the probe path is a valid path and the file can be accessed. Probe path is the path to a file
    // hosted on the origin server to help accelerate the delivery of dynamic content via the CDN endpoint. This path is
    // relative to the origin path specified in the endpoint configuration.
        // Parameters:
            // validateProbeInput - input to check.
    func (client BaseClient) ValidateProbe(ctx context.Context, validateProbeInput ValidateProbeInput) (result ValidateProbeOutput, err error) {
        if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BaseClient.ValidateProbe")
            defer func() {
                sc := -1
                if result.Response.Response != nil {
                    sc = result.Response.Response.StatusCode
                }
                tracing.EndSpan(ctx, sc, err)
            }()
        }
                if err := validation.Validate([]validation.Validation{
                { TargetValue: validateProbeInput,
                 Constraints: []validation.Constraint{	{Target: "validateProbeInput.ProbeURL", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
                return result, validation.NewError("cdn.BaseClient", "ValidateProbe", err.Error())
                }

                    req, err := client.ValidateProbePreparer(ctx, validateProbeInput)
        if err != nil {
        err = autorest.NewErrorWithError(err, "cdn.BaseClient", "ValidateProbe", nil , "Failure preparing request")
        return
        }

                resp, err := client.ValidateProbeSender(req)
                if err != nil {
                result.Response = autorest.Response{Response: resp}
                err = autorest.NewErrorWithError(err, "cdn.BaseClient", "ValidateProbe", resp, "Failure sending request")
                return
                }

                result, err = client.ValidateProbeResponder(resp)
                if err != nil {
                err = autorest.NewErrorWithError(err, "cdn.BaseClient", "ValidateProbe", resp, "Failure responding to request")
                }

        return
        }

        // ValidateProbePreparer prepares the ValidateProbe request.
        func (client BaseClient) ValidateProbePreparer(ctx context.Context, validateProbeInput ValidateProbeInput) (*http.Request, error) {
                pathParameters := map[string]interface{} {
                "subscriptionId": autorest.Encode("path",client.SubscriptionID),
                }

                            const APIVersion = "2017-04-02"
            queryParameters := map[string]interface{} {
            "api-version": APIVersion,
            }

            preparer := autorest.CreatePreparer(
        autorest.AsContentType("application/json; charset=utf-8"),
        autorest.AsPost(),
        autorest.WithBaseURL(client.BaseURI),
        autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/validateProbe",pathParameters),
        autorest.WithJSON(validateProbeInput),
        autorest.WithQueryParameters(queryParameters))
        return preparer.Prepare((&http.Request{}).WithContext(ctx))
        }

        // ValidateProbeSender sends the ValidateProbe request. The method will close the
        // http.Response Body if it receives an error.
        func (client BaseClient) ValidateProbeSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ValidateProbeResponder handles the response to the ValidateProbe request. The method always
    // closes the http.Response Body.
    func (client BaseClient) ValidateProbeResponder(resp *http.Response) (result ValidateProbeOutput, err error) {
        err = autorest.Respond(
        resp,
        client.ByInspecting(),
        azure.WithErrorUnlessStatusCode(http.StatusOK),
        autorest.ByUnmarshallingJSON(&result),
        autorest.ByClosing())
        result.Response = autorest.Response{Response: resp}
            return
        }

