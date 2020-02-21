package compute

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

// ContainerServicesClient is the compute Client
type ContainerServicesClient struct {
	BaseClient
}

// NewContainerServicesClient creates an instance of the ContainerServicesClient client.
func NewContainerServicesClient(subscriptionID string) ContainerServicesClient {
	return NewContainerServicesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewContainerServicesClientWithBaseURI creates an instance of the ContainerServicesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewContainerServicesClientWithBaseURI(baseURI string, subscriptionID string) ContainerServicesClient {
	return ContainerServicesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a container service with the specified configuration of orchestrator, masters, and
// agents.
// Parameters:
// resourceGroupName - the name of the resource group.
// containerServiceName - the name of the container service in the specified subscription and resource group.
// parameters - parameters supplied to the Create or Update a Container Service operation.
func (client ContainerServicesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, containerServiceName string, parameters ContainerService) (result ContainerServicesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainerServicesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ContainerServiceProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.ContainerServiceProperties.CustomProfile", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "parameters.ContainerServiceProperties.CustomProfile.Orchestrator", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "parameters.ContainerServiceProperties.ServicePrincipalProfile", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.ContainerServiceProperties.ServicePrincipalProfile.ClientID", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.ContainerServiceProperties.ServicePrincipalProfile.Secret", Name: validation.Null, Rule: true, Chain: nil},
						}},
					{Target: "parameters.ContainerServiceProperties.MasterProfile", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.ContainerServiceProperties.MasterProfile.DNSPrefix", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "parameters.ContainerServiceProperties.AgentPoolProfiles", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.ContainerServiceProperties.WindowsProfile", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.ContainerServiceProperties.WindowsProfile.AdminUsername", Name: validation.Null, Rule: true,
							Chain: []validation.Constraint{{Target: "parameters.ContainerServiceProperties.WindowsProfile.AdminUsername", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+([._]?[a-zA-Z0-9]+)*$`, Chain: nil}}},
							{Target: "parameters.ContainerServiceProperties.WindowsProfile.AdminPassword", Name: validation.Null, Rule: true, Chain: nil},
						}},
					{Target: "parameters.ContainerServiceProperties.LinuxProfile", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.ContainerServiceProperties.LinuxProfile.AdminUsername", Name: validation.Null, Rule: true,
							Chain: []validation.Constraint{{Target: "parameters.ContainerServiceProperties.LinuxProfile.AdminUsername", Name: validation.Pattern, Rule: `^[a-z][a-z0-9_-]*$`, Chain: nil}}},
							{Target: "parameters.ContainerServiceProperties.LinuxProfile.SSH", Name: validation.Null, Rule: true,
								Chain: []validation.Constraint{{Target: "parameters.ContainerServiceProperties.LinuxProfile.SSH.PublicKeys", Name: validation.Null, Rule: true, Chain: nil}}},
						}},
					{Target: "parameters.ContainerServiceProperties.DiagnosticsProfile", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.ContainerServiceProperties.DiagnosticsProfile.VMDiagnostics", Name: validation.Null, Rule: true,
							Chain: []validation.Constraint{{Target: "parameters.ContainerServiceProperties.DiagnosticsProfile.VMDiagnostics.Enabled", Name: validation.Null, Rule: true, Chain: nil}}},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("compute.ContainerServicesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, containerServiceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.ContainerServicesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.ContainerServicesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ContainerServicesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, containerServiceName string, parameters ContainerService) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerServiceName": autorest.Encode("path", containerServiceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerServicesClient) CreateOrUpdateSender(req *http.Request) (future ContainerServicesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ContainerServicesClient) CreateOrUpdateResponder(resp *http.Response) (result ContainerService, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified container service in the specified subscription and resource group. The operation does
// not delete other resources created as part of creating a container service, including storage accounts, VMs, and
// availability sets. All the other resources created with the container service are part of the same resource group
// and can be deleted individually.
// Parameters:
// resourceGroupName - the name of the resource group.
// containerServiceName - the name of the container service in the specified subscription and resource group.
func (client ContainerServicesClient) Delete(ctx context.Context, resourceGroupName string, containerServiceName string) (result ContainerServicesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainerServicesClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, containerServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.ContainerServicesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.ContainerServicesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ContainerServicesClient) DeletePreparer(ctx context.Context, resourceGroupName string, containerServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerServiceName": autorest.Encode("path", containerServiceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerServicesClient) DeleteSender(req *http.Request) (future ContainerServicesDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ContainerServicesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the properties of the specified container service in the specified subscription and resource group. The
// operation returns the properties including state, orchestrator, number of masters and agents, and FQDNs of masters
// and agents.
// Parameters:
// resourceGroupName - the name of the resource group.
// containerServiceName - the name of the container service in the specified subscription and resource group.
func (client ContainerServicesClient) Get(ctx context.Context, resourceGroupName string, containerServiceName string) (result ContainerService, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainerServicesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, containerServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.ContainerServicesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.ContainerServicesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.ContainerServicesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ContainerServicesClient) GetPreparer(ctx context.Context, resourceGroupName string, containerServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerServiceName": autorest.Encode("path", containerServiceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerServicesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ContainerServicesClient) GetResponder(resp *http.Response) (result ContainerService, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of container services in the specified subscription. The operation returns properties of each
// container service including state, orchestrator, number of masters and agents, and FQDNs of masters and agents.
func (client ContainerServicesClient) List(ctx context.Context) (result ContainerServiceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainerServicesClient.List")
		defer func() {
			sc := -1
			if result.cslr.Response.Response != nil {
				sc = result.cslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.ContainerServicesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.cslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.ContainerServicesClient", "List", resp, "Failure sending request")
		return
	}

	result.cslr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.ContainerServicesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ContainerServicesClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/containerServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerServicesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ContainerServicesClient) ListResponder(resp *http.Response) (result ContainerServiceListResult, err error) {
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
func (client ContainerServicesClient) listNextResults(ctx context.Context, lastResults ContainerServiceListResult) (result ContainerServiceListResult, err error) {
	req, err := lastResults.containerServiceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.ContainerServicesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.ContainerServicesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.ContainerServicesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ContainerServicesClient) ListComplete(ctx context.Context) (result ContainerServiceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainerServicesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup gets a list of container services in the specified subscription and resource group. The
// operation returns properties of each container service including state, orchestrator, number of masters and agents,
// and FQDNs of masters and agents.
// Parameters:
// resourceGroupName - the name of the resource group.
func (client ContainerServicesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ContainerServiceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainerServicesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.cslr.Response.Response != nil {
				sc = result.cslr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.ContainerServicesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.cslr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.ContainerServicesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.cslr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.ContainerServicesClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ContainerServicesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-01-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerServicesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ContainerServicesClient) ListByResourceGroupResponder(resp *http.Response) (result ContainerServiceListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client ContainerServicesClient) listByResourceGroupNextResults(ctx context.Context, lastResults ContainerServiceListResult) (result ContainerServiceListResult, err error) {
	req, err := lastResults.containerServiceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.ContainerServicesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.ContainerServicesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.ContainerServicesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ContainerServicesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result ContainerServiceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ContainerServicesClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}
