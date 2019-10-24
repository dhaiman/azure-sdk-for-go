package eventgridapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/eventgrid/mgmt/2020-01-01-preview/eventgrid"
)

// DomainsClientAPI contains the set of methods on the DomainsClient type.
type DomainsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, domainName string, domainInfo eventgrid.Domain) (result eventgrid.DomainsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, domainName string) (result eventgrid.DomainsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, domainName string) (result eventgrid.Domain, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32) (result eventgrid.DomainsListResultPage, err error)
	ListBySubscription(ctx context.Context, filter string, top *int32) (result eventgrid.DomainsListResultPage, err error)
	ListSharedAccessKeys(ctx context.Context, resourceGroupName string, domainName string) (result eventgrid.DomainSharedAccessKeys, err error)
	RegenerateKey(ctx context.Context, resourceGroupName string, domainName string, regenerateKeyRequest eventgrid.DomainRegenerateKeyRequest) (result eventgrid.DomainSharedAccessKeys, err error)
	Update(ctx context.Context, resourceGroupName string, domainName string, domainUpdateParameters eventgrid.DomainUpdateParameters) (result eventgrid.DomainsUpdateFuture, err error)
}

var _ DomainsClientAPI = (*eventgrid.DomainsClient)(nil)

// DomainTopicsClientAPI contains the set of methods on the DomainTopicsClient type.
type DomainTopicsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string) (result eventgrid.DomainTopicsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string) (result eventgrid.DomainTopicsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, domainName string, domainTopicName string) (result eventgrid.DomainTopic, err error)
	ListByDomain(ctx context.Context, resourceGroupName string, domainName string, filter string, top *int32) (result eventgrid.DomainTopicsListResultPage, err error)
}

var _ DomainTopicsClientAPI = (*eventgrid.DomainTopicsClient)(nil)

// EventSubscriptionsClientAPI contains the set of methods on the EventSubscriptionsClient type.
type EventSubscriptionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, scope string, eventSubscriptionName string, eventSubscriptionInfo eventgrid.EventSubscription) (result eventgrid.EventSubscriptionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, scope string, eventSubscriptionName string) (result eventgrid.EventSubscriptionsDeleteFuture, err error)
	Get(ctx context.Context, scope string, eventSubscriptionName string) (result eventgrid.EventSubscription, err error)
	GetFullURL(ctx context.Context, scope string, eventSubscriptionName string) (result eventgrid.EventSubscriptionFullURL, err error)
	ListByDomainTopic(ctx context.Context, resourceGroupName string, domainName string, topicName string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultPage, err error)
	ListByResource(ctx context.Context, resourceGroupName string, providerNamespace string, resourceTypeName string, resourceName string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultPage, err error)
	ListGlobalByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultPage, err error)
	ListGlobalByResourceGroupForTopicType(ctx context.Context, resourceGroupName string, topicTypeName string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultPage, err error)
	ListGlobalBySubscription(ctx context.Context, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultPage, err error)
	ListGlobalBySubscriptionForTopicType(ctx context.Context, topicTypeName string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultPage, err error)
	ListRegionalByResourceGroup(ctx context.Context, resourceGroupName string, location string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultPage, err error)
	ListRegionalByResourceGroupForTopicType(ctx context.Context, resourceGroupName string, location string, topicTypeName string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultPage, err error)
	ListRegionalBySubscription(ctx context.Context, location string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultPage, err error)
	ListRegionalBySubscriptionForTopicType(ctx context.Context, location string, topicTypeName string, filter string, top *int32) (result eventgrid.EventSubscriptionsListResultPage, err error)
	Update(ctx context.Context, scope string, eventSubscriptionName string, eventSubscriptionUpdateParameters eventgrid.EventSubscriptionUpdateParameters) (result eventgrid.EventSubscriptionsUpdateFuture, err error)
}

var _ EventSubscriptionsClientAPI = (*eventgrid.EventSubscriptionsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result eventgrid.OperationsListResult, err error)
}

var _ OperationsClientAPI = (*eventgrid.OperationsClient)(nil)

// TopicsClientAPI contains the set of methods on the TopicsClient type.
type TopicsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, topicName string, topicInfo eventgrid.Topic) (result eventgrid.TopicsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, topicName string) (result eventgrid.TopicsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, topicName string) (result eventgrid.Topic, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32) (result eventgrid.TopicsListResultPage, err error)
	ListBySubscription(ctx context.Context, filter string, top *int32) (result eventgrid.TopicsListResultPage, err error)
	ListEventTypes(ctx context.Context, resourceGroupName string, providerNamespace string, resourceTypeName string, resourceName string) (result eventgrid.EventTypesListResult, err error)
	ListSharedAccessKeys(ctx context.Context, resourceGroupName string, topicName string) (result eventgrid.TopicSharedAccessKeys, err error)
	RegenerateKey(ctx context.Context, resourceGroupName string, topicName string, regenerateKeyRequest eventgrid.TopicRegenerateKeyRequest) (result eventgrid.TopicSharedAccessKeys, err error)
	Update(ctx context.Context, resourceGroupName string, topicName string, topicUpdateParameters eventgrid.TopicUpdateParameters) (result eventgrid.TopicsUpdateFuture, err error)
}

var _ TopicsClientAPI = (*eventgrid.TopicsClient)(nil)

// TopicTypesClientAPI contains the set of methods on the TopicTypesClient type.
type TopicTypesClientAPI interface {
	Get(ctx context.Context, topicTypeName string) (result eventgrid.TopicTypeInfo, err error)
	List(ctx context.Context) (result eventgrid.TopicTypesListResult, err error)
	ListEventTypes(ctx context.Context, topicTypeName string) (result eventgrid.EventTypesListResult, err error)
}

var _ TopicTypesClientAPI = (*eventgrid.TopicTypesClient)(nil)
