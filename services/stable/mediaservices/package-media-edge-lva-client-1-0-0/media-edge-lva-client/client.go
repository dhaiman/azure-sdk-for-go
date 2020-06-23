// Package mediaedgelvaclient implements the Azure ARM Mediaedgelvaclient service API version 1.0.0.
//
// Direct Methods for Live Video Analytics on IoT Edge.
package mediaedgelvaclient

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
	// DefaultBaseURI is the default URI used for the service Mediaedgelvaclient
	DefaultBaseURI = ""
)

// BaseClient is the base client for Mediaedgelvaclient.
type BaseClient struct {
	autorest.Client
	BaseURI string
}

// New creates an instance of the BaseClient client.
func New() BaseClient {
	return NewWithBaseURI(DefaultBaseURI)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string) BaseClient {
	return BaseClient{
		Client:  autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI: baseURI,
	}
}

// GraphInstanceActivate activates a graph instance with the name provided, if it exists.
// Parameters:
// name - name of the graph instance to be activated.
func (client BaseClient) GraphInstanceActivate(ctx context.Context, name string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GraphInstanceActivate")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GraphInstanceActivatePreparer(ctx, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphInstanceActivate", nil, "Failure preparing request")
		return
	}

	resp, err := client.GraphInstanceActivateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphInstanceActivate", resp, "Failure sending request")
		return
	}

	result, err = client.GraphInstanceActivateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphInstanceActivate", resp, "Failure responding to request")
	}

	return
}

// GraphInstanceActivatePreparer prepares the GraphInstanceActivate request.
func (client BaseClient) GraphInstanceActivatePreparer(ctx context.Context, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name": autorest.Encode("path", name),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/api/GraphInstance/{name}/activate", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GraphInstanceActivateSender sends the GraphInstanceActivate request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GraphInstanceActivateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GraphInstanceActivateResponder handles the response to the GraphInstanceActivate request. The method always
// closes the http.Response Body.
func (client BaseClient) GraphInstanceActivateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GraphInstanceDeactivate deactivates a graph instance with the name provided, if it is active.
// Parameters:
// name - name of the graph instance to be deactivated.
func (client BaseClient) GraphInstanceDeactivate(ctx context.Context, name string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GraphInstanceDeactivate")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GraphInstanceDeactivatePreparer(ctx, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphInstanceDeactivate", nil, "Failure preparing request")
		return
	}

	resp, err := client.GraphInstanceDeactivateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphInstanceDeactivate", resp, "Failure sending request")
		return
	}

	result, err = client.GraphInstanceDeactivateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphInstanceDeactivate", resp, "Failure responding to request")
	}

	return
}

// GraphInstanceDeactivatePreparer prepares the GraphInstanceDeactivate request.
func (client BaseClient) GraphInstanceDeactivatePreparer(ctx context.Context, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name": autorest.Encode("path", name),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/api/GraphInstance/{name}/deactivate", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GraphInstanceDeactivateSender sends the GraphInstanceDeactivate request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GraphInstanceDeactivateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GraphInstanceDeactivateResponder handles the response to the GraphInstanceDeactivate request. The method always
// closes the http.Response Body.
func (client BaseClient) GraphInstanceDeactivateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GraphInstanceDelete deletes a graph instance.
// Parameters:
// name - name of the graph instance to be deleted.
func (client BaseClient) GraphInstanceDelete(ctx context.Context, name string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GraphInstanceDelete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GraphInstanceDeletePreparer(ctx, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphInstanceDelete", nil, "Failure preparing request")
		return
	}

	resp, err := client.GraphInstanceDeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphInstanceDelete", resp, "Failure sending request")
		return
	}

	result, err = client.GraphInstanceDeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphInstanceDelete", resp, "Failure responding to request")
	}

	return
}

// GraphInstanceDeletePreparer prepares the GraphInstanceDelete request.
func (client BaseClient) GraphInstanceDeletePreparer(ctx context.Context, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name": autorest.Encode("path", name),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/api/GraphInstance/{name}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GraphInstanceDeleteSender sends the GraphInstanceDelete request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GraphInstanceDeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GraphInstanceDeleteResponder handles the response to the GraphInstanceDelete request. The method always
// closes the http.Response Body.
func (client BaseClient) GraphInstanceDeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GraphInstanceGet retrieves a  graph instance by name. If a graph instance with that name has been previously
// created, then the direct method call will return the JSON representation of that instance.
// Parameters:
// name - name of the graph instance to be retrieved.
func (client BaseClient) GraphInstanceGet(ctx context.Context, name string) (result MediaGraphInstance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GraphInstanceGet")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GraphInstanceGetPreparer(ctx, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphInstanceGet", nil, "Failure preparing request")
		return
	}

	resp, err := client.GraphInstanceGetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphInstanceGet", resp, "Failure sending request")
		return
	}

	result, err = client.GraphInstanceGetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphInstanceGet", resp, "Failure responding to request")
	}

	return
}

// GraphInstanceGetPreparer prepares the GraphInstanceGet request.
func (client BaseClient) GraphInstanceGetPreparer(ctx context.Context, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name": autorest.Encode("path", name),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/api/GraphInstance/{name}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GraphInstanceGetSender sends the GraphInstanceGet request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GraphInstanceGetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GraphInstanceGetResponder handles the response to the GraphInstanceGet request. The method always
// closes the http.Response Body.
func (client BaseClient) GraphInstanceGetResponder(resp *http.Response) (result MediaGraphInstance, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GraphInstanceList retrieves a list of graph instances, if any, that have been created in the Live Video Analytics
// IoT Edge module.
func (client BaseClient) GraphInstanceList(ctx context.Context) (result MediaGraphInstanceCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GraphInstanceList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GraphInstanceListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphInstanceList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GraphInstanceListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphInstanceList", resp, "Failure sending request")
		return
	}

	result, err = client.GraphInstanceListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphInstanceList", resp, "Failure responding to request")
	}

	return
}

// GraphInstanceListPreparer prepares the GraphInstanceList request.
func (client BaseClient) GraphInstanceListPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/api/GraphInstance"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GraphInstanceListSender sends the GraphInstanceList request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GraphInstanceListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GraphInstanceListResponder handles the response to the GraphInstanceList request. The method always
// closes the http.Response Body.
func (client BaseClient) GraphInstanceListResponder(resp *http.Response) (result MediaGraphInstanceCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GraphInstanceSet creates or updates a graph instance.
func (client BaseClient) GraphInstanceSet(ctx context.Context, instance MediaGraphInstance) (result MediaGraphInstance, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GraphInstanceSet")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: instance,
			Constraints: []validation.Constraint{{Target: "instance.Name", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("mediaedgelvaclient.BaseClient", "GraphInstanceSet", err.Error())
	}

	req, err := client.GraphInstanceSetPreparer(ctx, instance)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphInstanceSet", nil, "Failure preparing request")
		return
	}

	resp, err := client.GraphInstanceSetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphInstanceSet", resp, "Failure sending request")
		return
	}

	result, err = client.GraphInstanceSetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphInstanceSet", resp, "Failure responding to request")
	}

	return
}

// GraphInstanceSetPreparer prepares the GraphInstanceSet request.
func (client BaseClient) GraphInstanceSetPreparer(ctx context.Context, instance MediaGraphInstance) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/api/GraphInstance"),
		autorest.WithJSON(instance))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GraphInstanceSetSender sends the GraphInstanceSet request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GraphInstanceSetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GraphInstanceSetResponder handles the response to the GraphInstanceSet request. The method always
// closes the http.Response Body.
func (client BaseClient) GraphInstanceSetResponder(resp *http.Response) (result MediaGraphInstance, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GraphTopologyDelete deletes a graph topology with the given name. This method should be called after all instances
// of the topology have been stopped and deleted.
// Parameters:
// name - name of the graph topology to be deleted.
func (client BaseClient) GraphTopologyDelete(ctx context.Context, name string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GraphTopologyDelete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GraphTopologyDeletePreparer(ctx, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphTopologyDelete", nil, "Failure preparing request")
		return
	}

	resp, err := client.GraphTopologyDeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphTopologyDelete", resp, "Failure sending request")
		return
	}

	result, err = client.GraphTopologyDeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphTopologyDelete", resp, "Failure responding to request")
	}

	return
}

// GraphTopologyDeletePreparer prepares the GraphTopologyDelete request.
func (client BaseClient) GraphTopologyDeletePreparer(ctx context.Context, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name": autorest.Encode("path", name),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/api/GraphTopology/{name}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GraphTopologyDeleteSender sends the GraphTopologyDelete request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GraphTopologyDeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GraphTopologyDeleteResponder handles the response to the GraphTopologyDelete request. The method always
// closes the http.Response Body.
func (client BaseClient) GraphTopologyDeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GraphTopologyGet retrieves a graph topology by name. If a topology with that name has been previously set, then the
// call will return the JSON representation of that topology.
// Parameters:
// name - name of the graph topology to be retrieved.
func (client BaseClient) GraphTopologyGet(ctx context.Context, name string) (result MediaGraphTopology, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GraphTopologyGet")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GraphTopologyGetPreparer(ctx, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphTopologyGet", nil, "Failure preparing request")
		return
	}

	resp, err := client.GraphTopologyGetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphTopologyGet", resp, "Failure sending request")
		return
	}

	result, err = client.GraphTopologyGetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphTopologyGet", resp, "Failure responding to request")
	}

	return
}

// GraphTopologyGetPreparer prepares the GraphTopologyGet request.
func (client BaseClient) GraphTopologyGetPreparer(ctx context.Context, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name": autorest.Encode("path", name),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/api/GraphTopology/{name}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GraphTopologyGetSender sends the GraphTopologyGet request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GraphTopologyGetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GraphTopologyGetResponder handles the response to the GraphTopologyGet request. The method always
// closes the http.Response Body.
func (client BaseClient) GraphTopologyGetResponder(resp *http.Response) (result MediaGraphTopology, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotFound),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GraphTopologyList retrieves a list of graph topologies that have been added to the Live Video Analytics IoT Edge
// module.
func (client BaseClient) GraphTopologyList(ctx context.Context) (result MediaGraphTopologyCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GraphTopologyList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GraphTopologyListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphTopologyList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GraphTopologyListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphTopologyList", resp, "Failure sending request")
		return
	}

	result, err = client.GraphTopologyListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphTopologyList", resp, "Failure responding to request")
	}

	return
}

// GraphTopologyListPreparer prepares the GraphTopologyList request.
func (client BaseClient) GraphTopologyListPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/api/GraphTopology"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GraphTopologyListSender sends the GraphTopologyList request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GraphTopologyListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GraphTopologyListResponder handles the response to the GraphTopologyList request. The method always
// closes the http.Response Body.
func (client BaseClient) GraphTopologyListResponder(resp *http.Response) (result MediaGraphTopologyCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GraphTopologySet creates a new topology or updates an existing one, with the given name. A topology is a blueprint
// defining what nodes are in a graph, and how they are interconnected.
func (client BaseClient) GraphTopologySet(ctx context.Context, topology MediaGraphTopology) (result MediaGraphTopology, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GraphTopologySet")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: topology,
			Constraints: []validation.Constraint{{Target: "topology.Name", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("mediaedgelvaclient.BaseClient", "GraphTopologySet", err.Error())
	}

	req, err := client.GraphTopologySetPreparer(ctx, topology)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphTopologySet", nil, "Failure preparing request")
		return
	}

	resp, err := client.GraphTopologySetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphTopologySet", resp, "Failure sending request")
		return
	}

	result, err = client.GraphTopologySetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mediaedgelvaclient.BaseClient", "GraphTopologySet", resp, "Failure responding to request")
	}

	return
}

// GraphTopologySetPreparer prepares the GraphTopologySet request.
func (client BaseClient) GraphTopologySetPreparer(ctx context.Context, topology MediaGraphTopology) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/api/GraphTopology"),
		autorest.WithJSON(topology))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GraphTopologySetSender sends the GraphTopologySet request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GraphTopologySetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GraphTopologySetResponder handles the response to the GraphTopologySet request. The method always
// closes the http.Response Body.
func (client BaseClient) GraphTopologySetResponder(resp *http.Response) (result MediaGraphTopology, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
