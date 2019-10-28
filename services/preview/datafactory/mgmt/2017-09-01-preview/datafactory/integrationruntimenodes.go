package datafactory

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

// IntegrationRuntimeNodesClient is the the Azure Data Factory V2 management API provides a RESTful set of web services
// that interact with Azure Data Factory V2 services.
type IntegrationRuntimeNodesClient struct {
	BaseClient
}

// NewIntegrationRuntimeNodesClient creates an instance of the IntegrationRuntimeNodesClient client.
func NewIntegrationRuntimeNodesClient(subscriptionID string) IntegrationRuntimeNodesClient {
	return NewIntegrationRuntimeNodesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewIntegrationRuntimeNodesClientWithBaseURI creates an instance of the IntegrationRuntimeNodesClient client.
func NewIntegrationRuntimeNodesClientWithBaseURI(baseURI string, subscriptionID string) IntegrationRuntimeNodesClient {
	return IntegrationRuntimeNodesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Delete deletes a self-hosted integration runtime node.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// integrationRuntimeName - the integration runtime name.
// nodeName - the integration runtime node name.
func (client IntegrationRuntimeNodesClient) Delete(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, nodeName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationRuntimeNodesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: integrationRuntimeName,
			Constraints: []validation.Constraint{{Target: "integrationRuntimeName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "integrationRuntimeName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "integrationRuntimeName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: nodeName,
			Constraints: []validation.Constraint{{Target: "nodeName", Name: validation.MaxLength, Rule: 150, Chain: nil},
				{Target: "nodeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "nodeName", Name: validation.Pattern, Rule: `^[a-z0-9A-Z][a-z0-9A-Z_-]{0,149}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.IntegrationRuntimeNodesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, factoryName, integrationRuntimeName, nodeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.IntegrationRuntimeNodesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "datafactory.IntegrationRuntimeNodesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.IntegrationRuntimeNodesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client IntegrationRuntimeNodesClient) DeletePreparer(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, nodeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":            autorest.Encode("path", factoryName),
		"integrationRuntimeName": autorest.Encode("path", integrationRuntimeName),
		"nodeName":               autorest.Encode("path", nodeName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}/nodes/{nodeName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationRuntimeNodesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client IntegrationRuntimeNodesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetIPAddress get the IP address of self-hosted integration runtime node.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// integrationRuntimeName - the integration runtime name.
// nodeName - the integration runtime node name.
func (client IntegrationRuntimeNodesClient) GetIPAddress(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, nodeName string) (result IntegrationRuntimeNodeIPAddress, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationRuntimeNodesClient.GetIPAddress")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: integrationRuntimeName,
			Constraints: []validation.Constraint{{Target: "integrationRuntimeName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "integrationRuntimeName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "integrationRuntimeName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: nodeName,
			Constraints: []validation.Constraint{{Target: "nodeName", Name: validation.MaxLength, Rule: 150, Chain: nil},
				{Target: "nodeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "nodeName", Name: validation.Pattern, Rule: `^[a-z0-9A-Z][a-z0-9A-Z_-]{0,149}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.IntegrationRuntimeNodesClient", "GetIPAddress", err.Error())
	}

	req, err := client.GetIPAddressPreparer(ctx, resourceGroupName, factoryName, integrationRuntimeName, nodeName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.IntegrationRuntimeNodesClient", "GetIPAddress", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetIPAddressSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.IntegrationRuntimeNodesClient", "GetIPAddress", resp, "Failure sending request")
		return
	}

	result, err = client.GetIPAddressResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.IntegrationRuntimeNodesClient", "GetIPAddress", resp, "Failure responding to request")
	}

	return
}

// GetIPAddressPreparer prepares the GetIPAddress request.
func (client IntegrationRuntimeNodesClient) GetIPAddressPreparer(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, nodeName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":            autorest.Encode("path", factoryName),
		"integrationRuntimeName": autorest.Encode("path", integrationRuntimeName),
		"nodeName":               autorest.Encode("path", nodeName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}/nodes/{nodeName}/ipAddress", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetIPAddressSender sends the GetIPAddress request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationRuntimeNodesClient) GetIPAddressSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetIPAddressResponder handles the response to the GetIPAddress request. The method always
// closes the http.Response Body.
func (client IntegrationRuntimeNodesClient) GetIPAddressResponder(resp *http.Response) (result IntegrationRuntimeNodeIPAddress, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates a self-hosted integration runtime node.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// integrationRuntimeName - the integration runtime name.
// nodeName - the integration runtime node name.
// updateIntegrationRuntimeNodeRequest - the parameters for updating an integration runtime node.
func (client IntegrationRuntimeNodesClient) Update(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, nodeName string, updateIntegrationRuntimeNodeRequest UpdateIntegrationRuntimeNodeRequest) (result SelfHostedIntegrationRuntimeNode, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/IntegrationRuntimeNodesClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: integrationRuntimeName,
			Constraints: []validation.Constraint{{Target: "integrationRuntimeName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "integrationRuntimeName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "integrationRuntimeName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: nodeName,
			Constraints: []validation.Constraint{{Target: "nodeName", Name: validation.MaxLength, Rule: 150, Chain: nil},
				{Target: "nodeName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "nodeName", Name: validation.Pattern, Rule: `^[a-z0-9A-Z][a-z0-9A-Z_-]{0,149}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.IntegrationRuntimeNodesClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, factoryName, integrationRuntimeName, nodeName, updateIntegrationRuntimeNodeRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.IntegrationRuntimeNodesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.IntegrationRuntimeNodesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.IntegrationRuntimeNodesClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client IntegrationRuntimeNodesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, nodeName string, updateIntegrationRuntimeNodeRequest UpdateIntegrationRuntimeNodeRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":            autorest.Encode("path", factoryName),
		"integrationRuntimeName": autorest.Encode("path", integrationRuntimeName),
		"nodeName":               autorest.Encode("path", nodeName),
		"resourceGroupName":      autorest.Encode("path", resourceGroupName),
		"subscriptionId":         autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}/nodes/{nodeName}", pathParameters),
		autorest.WithJSON(updateIntegrationRuntimeNodeRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client IntegrationRuntimeNodesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client IntegrationRuntimeNodesClient) UpdateResponder(resp *http.Response) (result SelfHostedIntegrationRuntimeNode, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
