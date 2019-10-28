package datafactoryapi

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
	"github.com/Azure/azure-sdk-for-go/services/datafactory/mgmt/2018-06-01/datafactory"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result datafactory.OperationListResponsePage, err error)
}

var _ OperationsClientAPI = (*datafactory.OperationsClient)(nil)

// FactoriesClientAPI contains the set of methods on the FactoriesClient type.
type FactoriesClientAPI interface {
	ConfigureFactoryRepo(ctx context.Context, locationID string, factoryRepoUpdate datafactory.FactoryRepoUpdate) (result datafactory.Factory, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, factory datafactory.Factory, ifMatch string) (result datafactory.Factory, err error)
	Delete(ctx context.Context, resourceGroupName string, factoryName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, factoryName string, ifNoneMatch string) (result datafactory.Factory, err error)
	GetDataPlaneAccess(ctx context.Context, resourceGroupName string, factoryName string, policy datafactory.UserAccessPolicy) (result datafactory.AccessPolicyResponse, err error)
	GetGitHubAccessToken(ctx context.Context, resourceGroupName string, factoryName string, gitHubAccessTokenRequest datafactory.GitHubAccessTokenRequest) (result datafactory.GitHubAccessTokenResponse, err error)
	List(ctx context.Context) (result datafactory.FactoryListResponsePage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result datafactory.FactoryListResponsePage, err error)
	Update(ctx context.Context, resourceGroupName string, factoryName string, factoryUpdateParameters datafactory.FactoryUpdateParameters) (result datafactory.Factory, err error)
}

var _ FactoriesClientAPI = (*datafactory.FactoriesClient)(nil)

// ExposureControlClientAPI contains the set of methods on the ExposureControlClient type.
type ExposureControlClientAPI interface {
	GetFeatureValue(ctx context.Context, locationID string, exposureControlRequest datafactory.ExposureControlRequest) (result datafactory.ExposureControlResponse, err error)
	GetFeatureValueByFactory(ctx context.Context, resourceGroupName string, factoryName string, exposureControlRequest datafactory.ExposureControlRequest) (result datafactory.ExposureControlResponse, err error)
}

var _ ExposureControlClientAPI = (*datafactory.ExposureControlClient)(nil)

// IntegrationRuntimesClientAPI contains the set of methods on the IntegrationRuntimesClient type.
type IntegrationRuntimesClientAPI interface {
	CreateLinkedIntegrationRuntime(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, createLinkedIntegrationRuntimeRequest datafactory.CreateLinkedIntegrationRuntimeRequest) (result datafactory.IntegrationRuntimeStatusResponse, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, integrationRuntime datafactory.IntegrationRuntimeResource, ifMatch string) (result datafactory.IntegrationRuntimeResource, err error)
	Delete(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, ifNoneMatch string) (result datafactory.IntegrationRuntimeResource, err error)
	GetConnectionInfo(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string) (result datafactory.IntegrationRuntimeConnectionInfo, err error)
	GetMonitoringData(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string) (result datafactory.IntegrationRuntimeMonitoringData, err error)
	GetStatus(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string) (result datafactory.IntegrationRuntimeStatusResponse, err error)
	ListAuthKeys(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string) (result datafactory.IntegrationRuntimeAuthKeys, err error)
	ListByFactory(ctx context.Context, resourceGroupName string, factoryName string) (result datafactory.IntegrationRuntimeListResponsePage, err error)
	RegenerateAuthKey(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, regenerateKeyParameters datafactory.IntegrationRuntimeRegenerateKeyParameters) (result datafactory.IntegrationRuntimeAuthKeys, err error)
	RemoveLinks(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, linkedIntegrationRuntimeRequest datafactory.LinkedIntegrationRuntimeRequest) (result autorest.Response, err error)
	Start(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string) (result datafactory.IntegrationRuntimesStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string) (result datafactory.IntegrationRuntimesStopFuture, err error)
	SyncCredentials(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string) (result autorest.Response, err error)
	Update(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, updateIntegrationRuntimeRequest datafactory.UpdateIntegrationRuntimeRequest) (result datafactory.IntegrationRuntimeResource, err error)
	Upgrade(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string) (result autorest.Response, err error)
}

var _ IntegrationRuntimesClientAPI = (*datafactory.IntegrationRuntimesClient)(nil)

// IntegrationRuntimeObjectMetadataClientAPI contains the set of methods on the IntegrationRuntimeObjectMetadataClient type.
type IntegrationRuntimeObjectMetadataClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, getMetadataRequest *datafactory.GetSsisObjectMetadataRequest) (result datafactory.SsisObjectMetadataListResponse, err error)
	Refresh(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string) (result datafactory.IntegrationRuntimeObjectMetadataRefreshFuture, err error)
}

var _ IntegrationRuntimeObjectMetadataClientAPI = (*datafactory.IntegrationRuntimeObjectMetadataClient)(nil)

// IntegrationRuntimeNodesClientAPI contains the set of methods on the IntegrationRuntimeNodesClient type.
type IntegrationRuntimeNodesClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, nodeName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, nodeName string) (result datafactory.SelfHostedIntegrationRuntimeNode, err error)
	GetIPAddress(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, nodeName string) (result datafactory.IntegrationRuntimeNodeIPAddress, err error)
	Update(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, nodeName string, updateIntegrationRuntimeNodeRequest datafactory.UpdateIntegrationRuntimeNodeRequest) (result datafactory.SelfHostedIntegrationRuntimeNode, err error)
}

var _ IntegrationRuntimeNodesClientAPI = (*datafactory.IntegrationRuntimeNodesClient)(nil)

// LinkedServicesClientAPI contains the set of methods on the LinkedServicesClient type.
type LinkedServicesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, linkedServiceName string, linkedService datafactory.LinkedServiceResource, ifMatch string) (result datafactory.LinkedServiceResource, err error)
	Delete(ctx context.Context, resourceGroupName string, factoryName string, linkedServiceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, factoryName string, linkedServiceName string, ifNoneMatch string) (result datafactory.LinkedServiceResource, err error)
	ListByFactory(ctx context.Context, resourceGroupName string, factoryName string) (result datafactory.LinkedServiceListResponsePage, err error)
}

var _ LinkedServicesClientAPI = (*datafactory.LinkedServicesClient)(nil)

// DatasetsClientAPI contains the set of methods on the DatasetsClient type.
type DatasetsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, datasetName string, dataset datafactory.DatasetResource, ifMatch string) (result datafactory.DatasetResource, err error)
	Delete(ctx context.Context, resourceGroupName string, factoryName string, datasetName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, factoryName string, datasetName string, ifNoneMatch string) (result datafactory.DatasetResource, err error)
	ListByFactory(ctx context.Context, resourceGroupName string, factoryName string) (result datafactory.DatasetListResponsePage, err error)
}

var _ DatasetsClientAPI = (*datafactory.DatasetsClient)(nil)

// PipelinesClientAPI contains the set of methods on the PipelinesClient type.
type PipelinesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, pipeline datafactory.PipelineResource, ifMatch string) (result datafactory.PipelineResource, err error)
	CreateRun(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, referencePipelineRunID string, isRecovery *bool, startActivityName string, parameters map[string]interface{}) (result datafactory.CreateRunResponse, err error)
	Delete(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, ifNoneMatch string) (result datafactory.PipelineResource, err error)
	ListByFactory(ctx context.Context, resourceGroupName string, factoryName string) (result datafactory.PipelineListResponsePage, err error)
}

var _ PipelinesClientAPI = (*datafactory.PipelinesClient)(nil)

// PipelineRunsClientAPI contains the set of methods on the PipelineRunsClient type.
type PipelineRunsClientAPI interface {
	Cancel(ctx context.Context, resourceGroupName string, factoryName string, runID string, isRecursive *bool) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, factoryName string, runID string) (result datafactory.PipelineRun, err error)
	QueryByFactory(ctx context.Context, resourceGroupName string, factoryName string, filterParameters datafactory.RunFilterParameters) (result datafactory.PipelineRunsQueryResponse, err error)
}

var _ PipelineRunsClientAPI = (*datafactory.PipelineRunsClient)(nil)

// ActivityRunsClientAPI contains the set of methods on the ActivityRunsClient type.
type ActivityRunsClientAPI interface {
	QueryByPipelineRun(ctx context.Context, resourceGroupName string, factoryName string, runID string, filterParameters datafactory.RunFilterParameters) (result datafactory.ActivityRunsQueryResponse, err error)
}

var _ ActivityRunsClientAPI = (*datafactory.ActivityRunsClient)(nil)

// TriggersClientAPI contains the set of methods on the TriggersClient type.
type TriggersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, trigger datafactory.TriggerResource, ifMatch string) (result datafactory.TriggerResource, err error)
	Delete(ctx context.Context, resourceGroupName string, factoryName string, triggerName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, ifNoneMatch string) (result datafactory.TriggerResource, err error)
	GetEventSubscriptionStatus(ctx context.Context, resourceGroupName string, factoryName string, triggerName string) (result datafactory.TriggerSubscriptionOperationStatus, err error)
	ListByFactory(ctx context.Context, resourceGroupName string, factoryName string) (result datafactory.TriggerListResponsePage, err error)
	Start(ctx context.Context, resourceGroupName string, factoryName string, triggerName string) (result datafactory.TriggersStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, factoryName string, triggerName string) (result datafactory.TriggersStopFuture, err error)
	SubscribeToEvents(ctx context.Context, resourceGroupName string, factoryName string, triggerName string) (result datafactory.TriggersSubscribeToEventsFuture, err error)
	UnsubscribeFromEvents(ctx context.Context, resourceGroupName string, factoryName string, triggerName string) (result datafactory.TriggersUnsubscribeFromEventsFuture, err error)
}

var _ TriggersClientAPI = (*datafactory.TriggersClient)(nil)

// TriggerRunsClientAPI contains the set of methods on the TriggerRunsClient type.
type TriggerRunsClientAPI interface {
	QueryByFactory(ctx context.Context, resourceGroupName string, factoryName string, filterParameters datafactory.RunFilterParameters) (result datafactory.TriggerRunsQueryResponse, err error)
	Rerun(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, runID string) (result autorest.Response, err error)
}

var _ TriggerRunsClientAPI = (*datafactory.TriggerRunsClient)(nil)

// RerunTriggersClientAPI contains the set of methods on the RerunTriggersClient type.
type RerunTriggersClientAPI interface {
	Cancel(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, rerunTriggerName string) (result datafactory.RerunTriggersCancelFuture, err error)
	Create(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, rerunTriggerName string, rerunTumblingWindowTriggerActionParameters datafactory.RerunTumblingWindowTriggerActionParameters) (result datafactory.TriggerResource, err error)
	ListByTrigger(ctx context.Context, resourceGroupName string, factoryName string, triggerName string) (result datafactory.RerunTriggerListResponsePage, err error)
	Start(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, rerunTriggerName string) (result datafactory.RerunTriggersStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, rerunTriggerName string) (result datafactory.RerunTriggersStopFuture, err error)
}

var _ RerunTriggersClientAPI = (*datafactory.RerunTriggersClient)(nil)

// DataFlowsClientAPI contains the set of methods on the DataFlowsClient type.
type DataFlowsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, dataFlowName string, dataFlow datafactory.DataFlowResource, ifMatch string) (result datafactory.DataFlowResource, err error)
	Delete(ctx context.Context, resourceGroupName string, factoryName string, dataFlowName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, factoryName string, dataFlowName string, ifNoneMatch string) (result datafactory.DataFlowResource, err error)
	ListByFactory(ctx context.Context, resourceGroupName string, factoryName string) (result datafactory.DataFlowListResponsePage, err error)
}

var _ DataFlowsClientAPI = (*datafactory.DataFlowsClient)(nil)

// DataFlowDebugSessionClientAPI contains the set of methods on the DataFlowDebugSessionClient type.
type DataFlowDebugSessionClientAPI interface {
	AddDataFlow(ctx context.Context, resourceGroupName string, factoryName string, request datafactory.DataFlowDebugPackage) (result datafactory.AddDataFlowToDebugSessionResponse, err error)
	Create(ctx context.Context, resourceGroupName string, factoryName string, request datafactory.CreateDataFlowDebugSessionRequest) (result datafactory.DataFlowDebugSessionCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, factoryName string, request datafactory.DeleteDataFlowDebugSessionRequest) (result autorest.Response, err error)
	ExecuteCommand(ctx context.Context, resourceGroupName string, factoryName string, request datafactory.DataFlowDebugCommandRequest) (result datafactory.DataFlowDebugSessionExecuteCommandFuture, err error)
	QueryByFactory(ctx context.Context, resourceGroupName string, factoryName string) (result datafactory.QueryDataFlowDebugSessionsResponsePage, err error)
}

var _ DataFlowDebugSessionClientAPI = (*datafactory.DataFlowDebugSessionClient)(nil)
