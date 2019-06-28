package authorization

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

// GlobalAdministratorClient is the client for the GlobalAdministrator methods of the Authorization service.
type GlobalAdministratorClient struct {
	BaseClient
}

// NewGlobalAdministratorClient creates an instance of the GlobalAdministratorClient client.
func NewGlobalAdministratorClient(subscriptionID string) GlobalAdministratorClient {
	return NewGlobalAdministratorClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGlobalAdministratorClientWithBaseURI creates an instance of the GlobalAdministratorClient client.
func NewGlobalAdministratorClientWithBaseURI(baseURI string, subscriptionID string) GlobalAdministratorClient {
	return GlobalAdministratorClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ElevateAccess elevates access for a Global Administrator.
func (client GlobalAdministratorClient) ElevateAccess(ctx context.Context) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GlobalAdministratorClient.ElevateAccess")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ElevateAccessPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.GlobalAdministratorClient", "ElevateAccess", nil, "Failure preparing request")
		return
	}

	resp, err := client.ElevateAccessSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "authorization.GlobalAdministratorClient", "ElevateAccess", resp, "Failure sending request")
		return
	}

	result, err = client.ElevateAccessResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.GlobalAdministratorClient", "ElevateAccess", resp, "Failure responding to request")
	}

	return
}

// ElevateAccessPreparer prepares the ElevateAccess request.
func (client GlobalAdministratorClient) ElevateAccessPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2015-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Authorization/elevateAccess"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ElevateAccessSender sends the ElevateAccess request. The method will close the
// http.Response Body if it receives an error.
func (client GlobalAdministratorClient) ElevateAccessSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ElevateAccessResponder handles the response to the ElevateAccess request. The method always
// closes the http.Response Body.
func (client GlobalAdministratorClient) ElevateAccessResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
