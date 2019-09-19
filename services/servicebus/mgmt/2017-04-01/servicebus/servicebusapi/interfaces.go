package servicebusapi

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
	"github.com/Azure/azure-sdk-for-go/services/servicebus/mgmt/2017-04-01/servicebus"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result servicebus.OperationListResultPage, err error)
}

var _ OperationsClientAPI = (*servicebus.OperationsClient)(nil)

// NamespacesClientAPI contains the set of methods on the NamespacesClient type.
type NamespacesClientAPI interface {
	CheckNameAvailabilityMethod(ctx context.Context, parameters servicebus.CheckNameAvailability) (result servicebus.CheckNameAvailabilityResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, parameters servicebus.SBNamespace) (result servicebus.NamespacesCreateOrUpdateFuture, err error)
	CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string, parameters servicebus.SBAuthorizationRule) (result servicebus.SBAuthorizationRule, err error)
	CreateOrUpdateNetworkRuleSet(ctx context.Context, resourceGroupName string, namespaceName string, parameters servicebus.NetworkRuleSet) (result servicebus.NetworkRuleSet, err error)
	Delete(ctx context.Context, resourceGroupName string, namespaceName string) (result servicebus.NamespacesDeleteFuture, err error)
	DeleteAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, namespaceName string) (result servicebus.SBNamespace, err error)
	GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string) (result servicebus.SBAuthorizationRule, err error)
	GetNetworkRuleSet(ctx context.Context, resourceGroupName string, namespaceName string) (result servicebus.NetworkRuleSet, err error)
	List(ctx context.Context) (result servicebus.SBNamespaceListResultPage, err error)
	ListAuthorizationRules(ctx context.Context, resourceGroupName string, namespaceName string) (result servicebus.SBAuthorizationRuleListResultPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result servicebus.SBNamespaceListResultPage, err error)
	ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string) (result servicebus.AccessKeys, err error)
	ListNetworkRuleSets(ctx context.Context, resourceGroupName string, namespaceName string) (result servicebus.NetworkRuleSetListResultPage, err error)
	Migrate(ctx context.Context, resourceGroupName string, namespaceName string, parameters servicebus.SBNamespaceMigrate) (result autorest.Response, err error)
	RegenerateKeys(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string, parameters servicebus.RegenerateAccessKeyParameters) (result servicebus.AccessKeys, err error)
	Update(ctx context.Context, resourceGroupName string, namespaceName string, parameters servicebus.SBNamespaceUpdateParameters) (result servicebus.SBNamespace, err error)
}

var _ NamespacesClientAPI = (*servicebus.NamespacesClient)(nil)

// DisasterRecoveryConfigsClientAPI contains the set of methods on the DisasterRecoveryConfigsClient type.
type DisasterRecoveryConfigsClientAPI interface {
	BreakPairing(ctx context.Context, resourceGroupName string, namespaceName string, alias string) (result autorest.Response, err error)
	CheckNameAvailabilityMethod(ctx context.Context, resourceGroupName string, namespaceName string, parameters servicebus.CheckNameAvailability) (result servicebus.CheckNameAvailabilityResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, alias string, parameters servicebus.ArmDisasterRecovery) (result servicebus.ArmDisasterRecovery, err error)
	Delete(ctx context.Context, resourceGroupName string, namespaceName string, alias string) (result autorest.Response, err error)
	FailOver(ctx context.Context, resourceGroupName string, namespaceName string, alias string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, namespaceName string, alias string) (result servicebus.ArmDisasterRecovery, err error)
	GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, alias string, authorizationRuleName string) (result servicebus.SBAuthorizationRule, err error)
	List(ctx context.Context, resourceGroupName string, namespaceName string) (result servicebus.ArmDisasterRecoveryListResultPage, err error)
	ListAuthorizationRules(ctx context.Context, resourceGroupName string, namespaceName string, alias string) (result servicebus.SBAuthorizationRuleListResultPage, err error)
	ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, alias string, authorizationRuleName string) (result servicebus.AccessKeys, err error)
}

var _ DisasterRecoveryConfigsClientAPI = (*servicebus.DisasterRecoveryConfigsClient)(nil)

// MigrationConfigsClientAPI contains the set of methods on the MigrationConfigsClient type.
type MigrationConfigsClientAPI interface {
	CompleteMigration(ctx context.Context, resourceGroupName string, namespaceName string) (result autorest.Response, err error)
	CreateAndStartMigration(ctx context.Context, resourceGroupName string, namespaceName string, parameters servicebus.MigrationConfigProperties) (result servicebus.MigrationConfigsCreateAndStartMigrationFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, namespaceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, namespaceName string) (result servicebus.MigrationConfigProperties, err error)
	List(ctx context.Context, resourceGroupName string, namespaceName string) (result servicebus.MigrationConfigListResultPage, err error)
	Revert(ctx context.Context, resourceGroupName string, namespaceName string) (result autorest.Response, err error)
}

var _ MigrationConfigsClientAPI = (*servicebus.MigrationConfigsClient)(nil)

// QueuesClientAPI contains the set of methods on the QueuesClient type.
type QueuesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, parameters servicebus.SBQueue) (result servicebus.SBQueue, err error)
	CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string, parameters servicebus.SBAuthorizationRule) (result servicebus.SBAuthorizationRule, err error)
	Delete(ctx context.Context, resourceGroupName string, namespaceName string, queueName string) (result autorest.Response, err error)
	DeleteAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, namespaceName string, queueName string) (result servicebus.SBQueue, err error)
	GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string) (result servicebus.SBAuthorizationRule, err error)
	ListAuthorizationRules(ctx context.Context, resourceGroupName string, namespaceName string, queueName string) (result servicebus.SBAuthorizationRuleListResultPage, err error)
	ListByNamespace(ctx context.Context, resourceGroupName string, namespaceName string, skip *int32, top *int32) (result servicebus.SBQueueListResultPage, err error)
	ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string) (result servicebus.AccessKeys, err error)
	RegenerateKeys(ctx context.Context, resourceGroupName string, namespaceName string, queueName string, authorizationRuleName string, parameters servicebus.RegenerateAccessKeyParameters) (result servicebus.AccessKeys, err error)
}

var _ QueuesClientAPI = (*servicebus.QueuesClient)(nil)

// TopicsClientAPI contains the set of methods on the TopicsClient type.
type TopicsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, parameters servicebus.SBTopic) (result servicebus.SBTopic, err error)
	CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, parameters servicebus.SBAuthorizationRule) (result servicebus.SBAuthorizationRule, err error)
	Delete(ctx context.Context, resourceGroupName string, namespaceName string, topicName string) (result autorest.Response, err error)
	DeleteAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, namespaceName string, topicName string) (result servicebus.SBTopic, err error)
	GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string) (result servicebus.SBAuthorizationRule, err error)
	ListAuthorizationRules(ctx context.Context, resourceGroupName string, namespaceName string, topicName string) (result servicebus.SBAuthorizationRuleListResultPage, err error)
	ListByNamespace(ctx context.Context, resourceGroupName string, namespaceName string, skip *int32, top *int32) (result servicebus.SBTopicListResultPage, err error)
	ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string) (result servicebus.AccessKeys, err error)
	RegenerateKeys(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, parameters servicebus.RegenerateAccessKeyParameters) (result servicebus.AccessKeys, err error)
}

var _ TopicsClientAPI = (*servicebus.TopicsClient)(nil)

// SubscriptionsClientAPI contains the set of methods on the SubscriptionsClient type.
type SubscriptionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string, parameters servicebus.SBSubscription) (result servicebus.SBSubscription, err error)
	Delete(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string) (result servicebus.SBSubscription, err error)
	ListByTopic(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, skip *int32, top *int32) (result servicebus.SBSubscriptionListResultPage, err error)
}

var _ SubscriptionsClientAPI = (*servicebus.SubscriptionsClient)(nil)

// RulesClientAPI contains the set of methods on the RulesClient type.
type RulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string, ruleName string, parameters servicebus.Rule) (result servicebus.Rule, err error)
	Delete(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string, ruleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string, ruleName string) (result servicebus.Rule, err error)
	ListBySubscriptions(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, subscriptionName string, skip *int32, top *int32) (result servicebus.RuleListResultPage, err error)
}

var _ RulesClientAPI = (*servicebus.RulesClient)(nil)

// RegionsClientAPI contains the set of methods on the RegionsClient type.
type RegionsClientAPI interface {
	ListBySku(ctx context.Context, sku string) (result servicebus.PremiumMessagingRegionsListResultPage, err error)
}

var _ RegionsClientAPI = (*servicebus.RegionsClient)(nil)

// PremiumMessagingRegionsClientAPI contains the set of methods on the PremiumMessagingRegionsClient type.
type PremiumMessagingRegionsClientAPI interface {
	List(ctx context.Context) (result servicebus.PremiumMessagingRegionsListResultPage, err error)
}

var _ PremiumMessagingRegionsClientAPI = (*servicebus.PremiumMessagingRegionsClient)(nil)

// EventHubsClientAPI contains the set of methods on the EventHubsClient type.
type EventHubsClientAPI interface {
	ListByNamespace(ctx context.Context, resourceGroupName string, namespaceName string) (result servicebus.EventHubListResultPage, err error)
}

var _ EventHubsClientAPI = (*servicebus.EventHubsClient)(nil)
