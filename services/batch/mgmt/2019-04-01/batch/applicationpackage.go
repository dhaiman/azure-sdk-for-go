package batch

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

// ApplicationPackageClient is the client for the ApplicationPackage methods of the Batch service.
type ApplicationPackageClient struct {
	BaseClient
}

// NewApplicationPackageClient creates an instance of the ApplicationPackageClient client.
func NewApplicationPackageClient(subscriptionID string) ApplicationPackageClient {
	return NewApplicationPackageClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewApplicationPackageClientWithBaseURI creates an instance of the ApplicationPackageClient client.
func NewApplicationPackageClientWithBaseURI(baseURI string, subscriptionID string) ApplicationPackageClient {
	return ApplicationPackageClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Activate activates the specified application package.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// applicationName - the name of the application. This must be unique within the account.
// versionName - the version of the application.
// parameters - the parameters for the request.
func (client ApplicationPackageClient) Activate(ctx context.Context, resourceGroupName string, accountName string, applicationName string, versionName string, parameters ActivateApplicationPackageParameters) (result ApplicationPackage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationPackageClient.Activate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: applicationName,
			Constraints: []validation.Constraint{{Target: "applicationName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applicationName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "applicationName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]+$`, Chain: nil}}},
		{TargetValue: versionName,
			Constraints: []validation.Constraint{{Target: "versionName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "versionName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "versionName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-][a-zA-Z0-9_.-]*$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Format", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.ApplicationPackageClient", "Activate", err.Error())
	}

	req, err := client.ActivatePreparer(ctx, resourceGroupName, accountName, applicationName, versionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.ApplicationPackageClient", "Activate", nil, "Failure preparing request")
		return
	}

	resp, err := client.ActivateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.ApplicationPackageClient", "Activate", resp, "Failure sending request")
		return
	}

	result, err = client.ActivateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.ApplicationPackageClient", "Activate", resp, "Failure responding to request")
	}

	return
}

// ActivatePreparer prepares the Activate request.
func (client ApplicationPackageClient) ActivatePreparer(ctx context.Context, resourceGroupName string, accountName string, applicationName string, versionName string, parameters ActivateApplicationPackageParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"applicationName":   autorest.Encode("path", applicationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"versionName":       autorest.Encode("path", versionName),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}/versions/{versionName}/activate", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ActivateSender sends the Activate request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationPackageClient) ActivateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ActivateResponder handles the response to the Activate request. The method always
// closes the http.Response Body.
func (client ApplicationPackageClient) ActivateResponder(resp *http.Response) (result ApplicationPackage, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Create creates an application package record.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// applicationName - the name of the application. This must be unique within the account.
// versionName - the version of the application.
// parameters - the parameters for the request.
func (client ApplicationPackageClient) Create(ctx context.Context, resourceGroupName string, accountName string, applicationName string, versionName string, parameters *ApplicationPackage) (result ApplicationPackage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationPackageClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: applicationName,
			Constraints: []validation.Constraint{{Target: "applicationName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applicationName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "applicationName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]+$`, Chain: nil}}},
		{TargetValue: versionName,
			Constraints: []validation.Constraint{{Target: "versionName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "versionName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "versionName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.ApplicationPackageClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, accountName, applicationName, versionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.ApplicationPackageClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.ApplicationPackageClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.ApplicationPackageClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ApplicationPackageClient) CreatePreparer(ctx context.Context, resourceGroupName string, accountName string, applicationName string, versionName string, parameters *ApplicationPackage) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"applicationName":   autorest.Encode("path", applicationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"versionName":       autorest.Encode("path", versionName),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}/versions/{versionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if parameters != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(parameters))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationPackageClient) CreateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ApplicationPackageClient) CreateResponder(resp *http.Response) (result ApplicationPackage, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an application package record and its associated binary file.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// applicationName - the name of the application. This must be unique within the account.
// versionName - the version of the application.
func (client ApplicationPackageClient) Delete(ctx context.Context, resourceGroupName string, accountName string, applicationName string, versionName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationPackageClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: applicationName,
			Constraints: []validation.Constraint{{Target: "applicationName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applicationName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "applicationName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]+$`, Chain: nil}}},
		{TargetValue: versionName,
			Constraints: []validation.Constraint{{Target: "versionName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "versionName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "versionName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.ApplicationPackageClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, accountName, applicationName, versionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.ApplicationPackageClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "batch.ApplicationPackageClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.ApplicationPackageClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ApplicationPackageClient) DeletePreparer(ctx context.Context, resourceGroupName string, accountName string, applicationName string, versionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"applicationName":   autorest.Encode("path", applicationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"versionName":       autorest.Encode("path", versionName),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}/versions/{versionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationPackageClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ApplicationPackageClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about the specified application package.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// applicationName - the name of the application. This must be unique within the account.
// versionName - the version of the application.
func (client ApplicationPackageClient) Get(ctx context.Context, resourceGroupName string, accountName string, applicationName string, versionName string) (result ApplicationPackage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationPackageClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: applicationName,
			Constraints: []validation.Constraint{{Target: "applicationName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applicationName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "applicationName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]+$`, Chain: nil}}},
		{TargetValue: versionName,
			Constraints: []validation.Constraint{{Target: "versionName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "versionName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "versionName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-][a-zA-Z0-9_.-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.ApplicationPackageClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, accountName, applicationName, versionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.ApplicationPackageClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.ApplicationPackageClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.ApplicationPackageClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplicationPackageClient) GetPreparer(ctx context.Context, resourceGroupName string, accountName string, applicationName string, versionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"applicationName":   autorest.Encode("path", applicationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"versionName":       autorest.Encode("path", versionName),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}/versions/{versionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationPackageClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplicationPackageClient) GetResponder(resp *http.Response) (result ApplicationPackage, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all of the application packages in the specified application.
// Parameters:
// resourceGroupName - the name of the resource group that contains the Batch account.
// accountName - the name of the Batch account.
// applicationName - the name of the application. This must be unique within the account.
// maxresults - the maximum number of items to return in the response.
func (client ApplicationPackageClient) List(ctx context.Context, resourceGroupName string, accountName string, applicationName string, maxresults *int32) (result ListApplicationPackagesResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationPackageClient.List")
		defer func() {
			sc := -1
			if result.lapr.Response.Response != nil {
				sc = result.lapr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: applicationName,
			Constraints: []validation.Constraint{{Target: "applicationName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "applicationName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "applicationName", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("batch.ApplicationPackageClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, accountName, applicationName, maxresults)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.ApplicationPackageClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lapr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batch.ApplicationPackageClient", "List", resp, "Failure sending request")
		return
	}

	result.lapr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.ApplicationPackageClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ApplicationPackageClient) ListPreparer(ctx context.Context, resourceGroupName string, accountName string, applicationName string, maxresults *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"applicationName":   autorest.Encode("path", applicationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if maxresults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxresults)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationName}/versions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationPackageClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ApplicationPackageClient) ListResponder(resp *http.Response) (result ListApplicationPackagesResult, err error) {
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
func (client ApplicationPackageClient) listNextResults(ctx context.Context, lastResults ListApplicationPackagesResult) (result ListApplicationPackagesResult, err error) {
	req, err := lastResults.listApplicationPackagesResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batch.ApplicationPackageClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batch.ApplicationPackageClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batch.ApplicationPackageClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ApplicationPackageClient) ListComplete(ctx context.Context, resourceGroupName string, accountName string, applicationName string, maxresults *int32) (result ListApplicationPackagesResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationPackageClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, accountName, applicationName, maxresults)
	return
}
