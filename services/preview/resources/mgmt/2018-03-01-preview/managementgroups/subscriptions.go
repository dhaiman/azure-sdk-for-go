package managementgroups

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

// SubscriptionsClient is the the Azure Management Groups API enables consolidation of multiple subscriptions/resources
// into an organizational hierarchy and centrally manage access control, policies, alerting and reporting for those
// resources.
type SubscriptionsClient struct {
	BaseClient
}

// NewSubscriptionsClient creates an instance of the SubscriptionsClient client.
func NewSubscriptionsClient() SubscriptionsClient {
	return NewSubscriptionsClientWithBaseURI(DefaultBaseURI)
}

// NewSubscriptionsClientWithBaseURI creates an instance of the SubscriptionsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewSubscriptionsClientWithBaseURI(baseURI string) SubscriptionsClient {
	return SubscriptionsClient{NewWithBaseURI(baseURI)}
}

// Create associates existing subscription with the management group.
// Parameters:
// groupID - management Group ID.
// subscriptionID - subscription ID.
// cacheControl - indicates that the request shouldn't utilize any caches.
func (client SubscriptionsClient) Create(ctx context.Context, groupID string, subscriptionID string, cacheControl string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.Create")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreatePreparer(ctx, groupID, subscriptionID, cacheControl)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.SubscriptionsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "managementgroups.SubscriptionsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.SubscriptionsClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client SubscriptionsClient) CreatePreparer(ctx context.Context, groupID string, subscriptionID string, cacheControl string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId":        autorest.Encode("path", groupID),
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2018-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{groupId}/subscriptions/{subscriptionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(cacheControl) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String(cacheControl)))
	} else {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String("no-cache")))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) CreateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) CreateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete de-associates subscription from the management group.
// Parameters:
// groupID - management Group ID.
// subscriptionID - subscription ID.
// cacheControl - indicates that the request shouldn't utilize any caches.
func (client SubscriptionsClient) Delete(ctx context.Context, groupID string, subscriptionID string, cacheControl string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SubscriptionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, groupID, subscriptionID, cacheControl)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.SubscriptionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "managementgroups.SubscriptionsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.SubscriptionsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SubscriptionsClient) DeletePreparer(ctx context.Context, groupID string, subscriptionID string, cacheControl string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupId":        autorest.Encode("path", groupID),
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2018-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{groupId}/subscriptions/{subscriptionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(cacheControl) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String(cacheControl)))
	} else {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Cache-Control", autorest.String("no-cache")))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SubscriptionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SubscriptionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
