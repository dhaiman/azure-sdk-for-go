// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package policyinsights

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/policyinsights/mgmt/2019-10-01/policyinsights"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type PolicyStatesResource = original.PolicyStatesResource

const (
	Default PolicyStatesResource = original.Default
	Latest  PolicyStatesResource = original.Latest
)

type BaseClient = original.BaseClient
type ComplianceDetail = original.ComplianceDetail
type ErrorDefinition = original.ErrorDefinition
type ErrorResponse = original.ErrorResponse
type ExpressionEvaluationDetails = original.ExpressionEvaluationDetails
type IfNotExistsEvaluationDetails = original.IfNotExistsEvaluationDetails
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationsClient = original.OperationsClient
type OperationsListResults = original.OperationsListResults
type PolicyAssignmentSummary = original.PolicyAssignmentSummary
type PolicyDefinitionSummary = original.PolicyDefinitionSummary
type PolicyDetails = original.PolicyDetails
type PolicyEvaluationDetails = original.PolicyEvaluationDetails
type PolicyEvent = original.PolicyEvent
type PolicyEventsClient = original.PolicyEventsClient
type PolicyEventsQueryResults = original.PolicyEventsQueryResults
type PolicyGroupSummary = original.PolicyGroupSummary
type PolicyState = original.PolicyState
type PolicyStatesClient = original.PolicyStatesClient
type PolicyStatesQueryResults = original.PolicyStatesQueryResults
type PolicyTrackedResource = original.PolicyTrackedResource
type PolicyTrackedResourcesClient = original.PolicyTrackedResourcesClient
type PolicyTrackedResourcesQueryResults = original.PolicyTrackedResourcesQueryResults
type PolicyTrackedResourcesQueryResultsIterator = original.PolicyTrackedResourcesQueryResultsIterator
type PolicyTrackedResourcesQueryResultsPage = original.PolicyTrackedResourcesQueryResultsPage
type QueryFailure = original.QueryFailure
type QueryFailureError = original.QueryFailureError
type Remediation = original.Remediation
type RemediationDeployment = original.RemediationDeployment
type RemediationDeploymentSummary = original.RemediationDeploymentSummary
type RemediationDeploymentsListResult = original.RemediationDeploymentsListResult
type RemediationDeploymentsListResultIterator = original.RemediationDeploymentsListResultIterator
type RemediationDeploymentsListResultPage = original.RemediationDeploymentsListResultPage
type RemediationFilters = original.RemediationFilters
type RemediationListResult = original.RemediationListResult
type RemediationListResultIterator = original.RemediationListResultIterator
type RemediationListResultPage = original.RemediationListResultPage
type RemediationProperties = original.RemediationProperties
type RemediationsClient = original.RemediationsClient
type String = original.String
type SummarizeResults = original.SummarizeResults
type Summary = original.Summary
type SummaryResults = original.SummaryResults
type TrackedResourceModificationDetails = original.TrackedResourceModificationDetails
type TypedErrorInfo = original.TypedErrorInfo

func New() BaseClient {
	return original.New()
}
func NewOperationsClient() OperationsClient {
	return original.NewOperationsClient()
}
func NewOperationsClientWithBaseURI(baseURI string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI)
}
func NewPolicyEventsClient() PolicyEventsClient {
	return original.NewPolicyEventsClient()
}
func NewPolicyEventsClientWithBaseURI(baseURI string) PolicyEventsClient {
	return original.NewPolicyEventsClientWithBaseURI(baseURI)
}
func NewPolicyStatesClient() PolicyStatesClient {
	return original.NewPolicyStatesClient()
}
func NewPolicyStatesClientWithBaseURI(baseURI string) PolicyStatesClient {
	return original.NewPolicyStatesClientWithBaseURI(baseURI)
}
func NewPolicyTrackedResourcesClient() PolicyTrackedResourcesClient {
	return original.NewPolicyTrackedResourcesClient()
}
func NewPolicyTrackedResourcesClientWithBaseURI(baseURI string) PolicyTrackedResourcesClient {
	return original.NewPolicyTrackedResourcesClientWithBaseURI(baseURI)
}
func NewPolicyTrackedResourcesQueryResultsIterator(page PolicyTrackedResourcesQueryResultsPage) PolicyTrackedResourcesQueryResultsIterator {
	return original.NewPolicyTrackedResourcesQueryResultsIterator(page)
}
func NewPolicyTrackedResourcesQueryResultsPage(getNextPage func(context.Context, PolicyTrackedResourcesQueryResults) (PolicyTrackedResourcesQueryResults, error)) PolicyTrackedResourcesQueryResultsPage {
	return original.NewPolicyTrackedResourcesQueryResultsPage(getNextPage)
}
func NewRemediationDeploymentsListResultIterator(page RemediationDeploymentsListResultPage) RemediationDeploymentsListResultIterator {
	return original.NewRemediationDeploymentsListResultIterator(page)
}
func NewRemediationDeploymentsListResultPage(getNextPage func(context.Context, RemediationDeploymentsListResult) (RemediationDeploymentsListResult, error)) RemediationDeploymentsListResultPage {
	return original.NewRemediationDeploymentsListResultPage(getNextPage)
}
func NewRemediationListResultIterator(page RemediationListResultPage) RemediationListResultIterator {
	return original.NewRemediationListResultIterator(page)
}
func NewRemediationListResultPage(getNextPage func(context.Context, RemediationListResult) (RemediationListResult, error)) RemediationListResultPage {
	return original.NewRemediationListResultPage(getNextPage)
}
func NewRemediationsClient() RemediationsClient {
	return original.NewRemediationsClient()
}
func NewRemediationsClientWithBaseURI(baseURI string) RemediationsClient {
	return original.NewRemediationsClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossiblePolicyStatesResourceValues() []PolicyStatesResource {
	return original.PossiblePolicyStatesResourceValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
