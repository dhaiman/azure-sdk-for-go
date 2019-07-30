package datamigration

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

// FilesClient is the data Migration Client
type FilesClient struct {
	BaseClient
}

// NewFilesClient creates an instance of the FilesClient client.
func NewFilesClient(subscriptionID string) FilesClient {
	return NewFilesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFilesClientWithBaseURI creates an instance of the FilesClient client.
func NewFilesClientWithBaseURI(baseURI string, subscriptionID string) FilesClient {
	return FilesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the PUT method creates a new file or updates an existing one.
// Parameters:
// parameters - information about the file
// groupName - name of the resource group
// serviceName - name of the service
// projectName - name of the project
// fileName - name of the File
func (client FilesClient) CreateOrUpdate(ctx context.Context, parameters ProjectFile, groupName string, serviceName string, projectName string, fileName string) (result ProjectFile, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FilesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, parameters, groupName, serviceName, projectName, fileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.FilesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datamigration.FilesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.FilesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client FilesClient) CreateOrUpdatePreparer(ctx context.Context, parameters ProjectFile, groupName string, serviceName string, projectName string, fileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fileName":       autorest.Encode("path", fileName),
		"groupName":      autorest.Encode("path", groupName),
		"projectName":    autorest.Encode("path", projectName),
		"serviceName":    autorest.Encode("path", serviceName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/files/{fileName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client FilesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client FilesClient) CreateOrUpdateResponder(resp *http.Response) (result ProjectFile, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete this method deletes a file.
// Parameters:
// groupName - name of the resource group
// serviceName - name of the service
// projectName - name of the project
// fileName - name of the File
func (client FilesClient) Delete(ctx context.Context, groupName string, serviceName string, projectName string, fileName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FilesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, groupName, serviceName, projectName, fileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.FilesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "datamigration.FilesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.FilesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client FilesClient) DeletePreparer(ctx context.Context, groupName string, serviceName string, projectName string, fileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fileName":       autorest.Encode("path", fileName),
		"groupName":      autorest.Encode("path", groupName),
		"projectName":    autorest.Encode("path", projectName),
		"serviceName":    autorest.Encode("path", serviceName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/files/{fileName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client FilesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client FilesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the files resource is a nested, proxy-only resource representing a file stored under the project resource. This
// method retrieves information about a file.
// Parameters:
// groupName - name of the resource group
// serviceName - name of the service
// projectName - name of the project
// fileName - name of the File
func (client FilesClient) Get(ctx context.Context, groupName string, serviceName string, projectName string, fileName string) (result ProjectFile, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FilesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, groupName, serviceName, projectName, fileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.FilesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datamigration.FilesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.FilesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client FilesClient) GetPreparer(ctx context.Context, groupName string, serviceName string, projectName string, fileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fileName":       autorest.Encode("path", fileName),
		"groupName":      autorest.Encode("path", groupName),
		"projectName":    autorest.Encode("path", projectName),
		"serviceName":    autorest.Encode("path", serviceName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/files/{fileName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client FilesClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client FilesClient) GetResponder(resp *http.Response) (result ProjectFile, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the project resource is a nested resource representing a stored migration project. This method returns a list
// of files owned by a project resource.
// Parameters:
// groupName - name of the resource group
// serviceName - name of the service
// projectName - name of the project
func (client FilesClient) List(ctx context.Context, groupName string, serviceName string, projectName string) (result FileListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FilesClient.List")
		defer func() {
			sc := -1
			if result.fl.Response.Response != nil {
				sc = result.fl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, groupName, serviceName, projectName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.FilesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.fl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datamigration.FilesClient", "List", resp, "Failure sending request")
		return
	}

	result.fl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.FilesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client FilesClient) ListPreparer(ctx context.Context, groupName string, serviceName string, projectName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"groupName":      autorest.Encode("path", groupName),
		"projectName":    autorest.Encode("path", projectName),
		"serviceName":    autorest.Encode("path", serviceName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/files", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client FilesClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client FilesClient) ListResponder(resp *http.Response) (result FileList, err error) {
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
func (client FilesClient) listNextResults(ctx context.Context, lastResults FileList) (result FileList, err error) {
	req, err := lastResults.fileListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datamigration.FilesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datamigration.FilesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.FilesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client FilesClient) ListComplete(ctx context.Context, groupName string, serviceName string, projectName string) (result FileListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FilesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, groupName, serviceName, projectName)
	return
}

// Read this method is used for requesting storage information using which contents of the file can be downloaded.
// Parameters:
// groupName - name of the resource group
// serviceName - name of the service
// projectName - name of the project
// fileName - name of the File
func (client FilesClient) Read(ctx context.Context, groupName string, serviceName string, projectName string, fileName string) (result FileStorageInfo, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FilesClient.Read")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ReadPreparer(ctx, groupName, serviceName, projectName, fileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.FilesClient", "Read", nil, "Failure preparing request")
		return
	}

	resp, err := client.ReadSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datamigration.FilesClient", "Read", resp, "Failure sending request")
		return
	}

	result, err = client.ReadResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.FilesClient", "Read", resp, "Failure responding to request")
	}

	return
}

// ReadPreparer prepares the Read request.
func (client FilesClient) ReadPreparer(ctx context.Context, groupName string, serviceName string, projectName string, fileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fileName":       autorest.Encode("path", fileName),
		"groupName":      autorest.Encode("path", groupName),
		"projectName":    autorest.Encode("path", projectName),
		"serviceName":    autorest.Encode("path", serviceName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/files/{fileName}/read", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ReadSender sends the Read request. The method will close the
// http.Response Body if it receives an error.
func (client FilesClient) ReadSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ReadResponder handles the response to the Read request. The method always
// closes the http.Response Body.
func (client FilesClient) ReadResponder(resp *http.Response) (result FileStorageInfo, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ReadWrite this method is used for requesting information for reading and writing the file content.
// Parameters:
// groupName - name of the resource group
// serviceName - name of the service
// projectName - name of the project
// fileName - name of the File
func (client FilesClient) ReadWrite(ctx context.Context, groupName string, serviceName string, projectName string, fileName string) (result FileStorageInfo, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FilesClient.ReadWrite")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ReadWritePreparer(ctx, groupName, serviceName, projectName, fileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.FilesClient", "ReadWrite", nil, "Failure preparing request")
		return
	}

	resp, err := client.ReadWriteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datamigration.FilesClient", "ReadWrite", resp, "Failure sending request")
		return
	}

	result, err = client.ReadWriteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.FilesClient", "ReadWrite", resp, "Failure responding to request")
	}

	return
}

// ReadWritePreparer prepares the ReadWrite request.
func (client FilesClient) ReadWritePreparer(ctx context.Context, groupName string, serviceName string, projectName string, fileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fileName":       autorest.Encode("path", fileName),
		"groupName":      autorest.Encode("path", groupName),
		"projectName":    autorest.Encode("path", projectName),
		"serviceName":    autorest.Encode("path", serviceName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/files/{fileName}/readwrite", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ReadWriteSender sends the ReadWrite request. The method will close the
// http.Response Body if it receives an error.
func (client FilesClient) ReadWriteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ReadWriteResponder handles the response to the ReadWrite request. The method always
// closes the http.Response Body.
func (client FilesClient) ReadWriteResponder(resp *http.Response) (result FileStorageInfo, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update this method updates an existing file.
// Parameters:
// parameters - information about the file
// groupName - name of the resource group
// serviceName - name of the service
// projectName - name of the project
// fileName - name of the File
func (client FilesClient) Update(ctx context.Context, parameters ProjectFile, groupName string, serviceName string, projectName string, fileName string) (result ProjectFile, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/FilesClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, parameters, groupName, serviceName, projectName, fileName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.FilesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datamigration.FilesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datamigration.FilesClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client FilesClient) UpdatePreparer(ctx context.Context, parameters ProjectFile, groupName string, serviceName string, projectName string, fileName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fileName":       autorest.Encode("path", fileName),
		"groupName":      autorest.Encode("path", groupName),
		"projectName":    autorest.Encode("path", projectName),
		"serviceName":    autorest.Encode("path", serviceName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.DataMigration/services/{serviceName}/projects/{projectName}/files/{fileName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client FilesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client FilesClient) UpdateResponder(resp *http.Response) (result ProjectFile, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
