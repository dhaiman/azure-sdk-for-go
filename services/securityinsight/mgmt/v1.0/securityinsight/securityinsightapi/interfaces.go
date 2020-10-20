package securityinsightapi

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
	"github.com/Azure/azure-sdk-for-go/services/securityinsight/mgmt/v1.0/securityinsight"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result securityinsight.OperationsListPage, err error)
	ListComplete(ctx context.Context) (result securityinsight.OperationsListIterator, err error)
}

var _ OperationsClientAPI = (*securityinsight.OperationsClient)(nil)

// AlertRulesClientAPI contains the set of methods on the AlertRulesClient type.
type AlertRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, alertRule securityinsight.BasicAlertRule) (result securityinsight.AlertRuleModel, err error)
	CreateOrUpdateAction(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, actionID string, action securityinsight.ActionRequest) (result securityinsight.ActionResponse, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string) (result autorest.Response, err error)
	DeleteAction(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, actionID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string) (result securityinsight.AlertRuleModel, err error)
	GetAction(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string, actionID string) (result securityinsight.ActionResponse, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.AlertRulesListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.AlertRulesListIterator, err error)
}

var _ AlertRulesClientAPI = (*securityinsight.AlertRulesClient)(nil)

// ActionsClientAPI contains the set of methods on the ActionsClient type.
type ActionsClientAPI interface {
	ListByAlertRule(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string) (result securityinsight.ActionsListPage, err error)
	ListByAlertRuleComplete(ctx context.Context, resourceGroupName string, workspaceName string, ruleID string) (result securityinsight.ActionsListIterator, err error)
}

var _ ActionsClientAPI = (*securityinsight.ActionsClient)(nil)

// AlertRuleTemplatesClientAPI contains the set of methods on the AlertRuleTemplatesClient type.
type AlertRuleTemplatesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, workspaceName string, alertRuleTemplateID string) (result securityinsight.AlertRuleTemplateModel, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.AlertRuleTemplatesListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.AlertRuleTemplatesListIterator, err error)
}

var _ AlertRuleTemplatesClientAPI = (*securityinsight.AlertRuleTemplatesClient)(nil)

// BookmarksClientAPI contains the set of methods on the BookmarksClient type.
type BookmarksClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, bookmarkID string, bookmark securityinsight.Bookmark) (result securityinsight.Bookmark, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, bookmarkID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, bookmarkID string) (result securityinsight.Bookmark, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.BookmarkListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.BookmarkListIterator, err error)
}

var _ BookmarksClientAPI = (*securityinsight.BookmarksClient)(nil)

// DataConnectorsClientAPI contains the set of methods on the DataConnectorsClient type.
type DataConnectorsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string, dataConnector securityinsight.BasicDataConnector) (result securityinsight.DataConnectorModel, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, dataConnectorID string) (result securityinsight.DataConnectorModel, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.DataConnectorListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result securityinsight.DataConnectorListIterator, err error)
}

var _ DataConnectorsClientAPI = (*securityinsight.DataConnectorsClient)(nil)

// IncidentsClientAPI contains the set of methods on the IncidentsClient type.
type IncidentsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incident securityinsight.Incident) (result securityinsight.Incident, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string) (result securityinsight.Incident, err error)
	List(ctx context.Context, resourceGroupName string, workspaceName string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.IncidentListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, workspaceName string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.IncidentListIterator, err error)
}

var _ IncidentsClientAPI = (*securityinsight.IncidentsClient)(nil)

// IncidentCommentsClientAPI contains the set of methods on the IncidentCommentsClient type.
type IncidentCommentsClientAPI interface {
	CreateComment(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incidentCommentID string, incidentComment securityinsight.IncidentComment) (result securityinsight.IncidentComment, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incidentCommentID string) (result securityinsight.IncidentComment, err error)
	ListByIncident(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.IncidentCommentListPage, err error)
	ListByIncidentComplete(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, filter string, orderby string, top *int32, skipToken string) (result securityinsight.IncidentCommentListIterator, err error)
}

var _ IncidentCommentsClientAPI = (*securityinsight.IncidentCommentsClient)(nil)
