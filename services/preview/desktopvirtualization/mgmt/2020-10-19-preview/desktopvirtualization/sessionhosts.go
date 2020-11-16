package desktopvirtualization

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

// SessionHostsClient is the client for the SessionHosts methods of the Desktopvirtualization service.
type SessionHostsClient struct {
	BaseClient
}

// NewSessionHostsClient creates an instance of the SessionHostsClient client.
func NewSessionHostsClient(subscriptionID string) SessionHostsClient {
	return NewSessionHostsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSessionHostsClientWithBaseURI creates an instance of the SessionHostsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewSessionHostsClientWithBaseURI(baseURI string, subscriptionID string) SessionHostsClient {
	return SessionHostsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Delete remove a SessionHost.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// hostPoolName - the name of the host pool within the specified resource group
// sessionHostName - the name of the session host within the specified host pool
// force - force flag to force sessionHost deletion even when userSession exists.
func (client SessionHostsClient) Delete(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, force *bool) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SessionHostsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: hostPoolName,
			Constraints: []validation.Constraint{{Target: "hostPoolName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "hostPoolName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: sessionHostName,
			Constraints: []validation.Constraint{{Target: "sessionHostName", Name: validation.MaxLength, Rule: 48, Chain: nil},
				{Target: "sessionHostName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("desktopvirtualization.SessionHostsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, hostPoolName, sessionHostName, force)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.SessionHostsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "desktopvirtualization.SessionHostsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.SessionHostsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SessionHostsClient) DeletePreparer(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, force *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostPoolName":      autorest.Encode("path", hostPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sessionHostName":   autorest.Encode("path", sessionHostName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-19-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if force != nil {
		queryParameters["force"] = autorest.Encode("query", *force)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/sessionHosts/{sessionHostName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SessionHostsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SessionHostsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a session host.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// hostPoolName - the name of the host pool within the specified resource group
// sessionHostName - the name of the session host within the specified host pool
func (client SessionHostsClient) Get(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string) (result SessionHost, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SessionHostsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: hostPoolName,
			Constraints: []validation.Constraint{{Target: "hostPoolName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "hostPoolName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: sessionHostName,
			Constraints: []validation.Constraint{{Target: "sessionHostName", Name: validation.MaxLength, Rule: 48, Chain: nil},
				{Target: "sessionHostName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("desktopvirtualization.SessionHostsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, hostPoolName, sessionHostName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.SessionHostsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "desktopvirtualization.SessionHostsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.SessionHostsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SessionHostsClient) GetPreparer(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostPoolName":      autorest.Encode("path", hostPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sessionHostName":   autorest.Encode("path", sessionHostName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-19-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/sessionHosts/{sessionHostName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SessionHostsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SessionHostsClient) GetResponder(resp *http.Response) (result SessionHost, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list sessionHosts.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// hostPoolName - the name of the host pool within the specified resource group
func (client SessionHostsClient) List(ctx context.Context, resourceGroupName string, hostPoolName string) (result SessionHostListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SessionHostsClient.List")
		defer func() {
			sc := -1
			if result.shl.Response.Response != nil {
				sc = result.shl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: hostPoolName,
			Constraints: []validation.Constraint{{Target: "hostPoolName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "hostPoolName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("desktopvirtualization.SessionHostsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, hostPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.SessionHostsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.shl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "desktopvirtualization.SessionHostsClient", "List", resp, "Failure sending request")
		return
	}

	result.shl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.SessionHostsClient", "List", resp, "Failure responding to request")
	}
	if result.shl.hasNextLink() && result.shl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client SessionHostsClient) ListPreparer(ctx context.Context, resourceGroupName string, hostPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostPoolName":      autorest.Encode("path", hostPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-19-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/sessionHosts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SessionHostsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SessionHostsClient) ListResponder(resp *http.Response) (result SessionHostList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client SessionHostsClient) listNextResults(ctx context.Context, lastResults SessionHostList) (result SessionHostList, err error) {
	req, err := lastResults.sessionHostListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "desktopvirtualization.SessionHostsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "desktopvirtualization.SessionHostsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.SessionHostsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client SessionHostsClient) ListComplete(ctx context.Context, resourceGroupName string, hostPoolName string) (result SessionHostListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SessionHostsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, hostPoolName)
	return
}

// Update update a session host.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// hostPoolName - the name of the host pool within the specified resource group
// sessionHostName - the name of the session host within the specified host pool
// sessionHost - object containing SessionHost definitions.
func (client SessionHostsClient) Update(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, sessionHost *SessionHostPatch) (result SessionHost, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SessionHostsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: hostPoolName,
			Constraints: []validation.Constraint{{Target: "hostPoolName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "hostPoolName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: sessionHostName,
			Constraints: []validation.Constraint{{Target: "sessionHostName", Name: validation.MaxLength, Rule: 48, Chain: nil},
				{Target: "sessionHostName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("desktopvirtualization.SessionHostsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, hostPoolName, sessionHostName, sessionHost)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.SessionHostsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "desktopvirtualization.SessionHostsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "desktopvirtualization.SessionHostsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client SessionHostsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, sessionHost *SessionHostPatch) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostPoolName":      autorest.Encode("path", hostPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sessionHostName":   autorest.Encode("path", sessionHostName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-10-19-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/sessionHosts/{sessionHostName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if sessionHost != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(sessionHost))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client SessionHostsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client SessionHostsClient) UpdateResponder(resp *http.Response) (result SessionHost, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
