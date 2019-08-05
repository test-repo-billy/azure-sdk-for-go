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
	"github.com/Azure/azure-sdk-for-go/services/preview/securityinsight/mgmt/2017-08-01-preview/securityinsight"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result securityinsight.OperationsListPage, err error)
}

var _ OperationsClientAPI = (*securityinsight.OperationsClient)(nil)

// AlertRulesClientAPI contains the set of methods on the AlertRulesClient type.
type AlertRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string, alertRule securityinsight.BasicAlertRule) (result securityinsight.AlertRuleModel, err error)
	CreateOrUpdateAction(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string, actionID string, action securityinsight.Action) (result securityinsight.Action, err error)
	Delete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string) (result autorest.Response, err error)
	DeleteAction(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string, actionID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string) (result securityinsight.AlertRuleModel, err error)
	GetAction(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string, actionID string) (result securityinsight.Action, err error)
	List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result securityinsight.AlertRulesListPage, err error)
}

var _ AlertRulesClientAPI = (*securityinsight.AlertRulesClient)(nil)

// ActionsClientAPI contains the set of methods on the ActionsClient type.
type ActionsClientAPI interface {
	ListByAlertRule(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string) (result securityinsight.ActionsListPage, err error)
}

var _ ActionsClientAPI = (*securityinsight.ActionsClient)(nil)

// CasesClientAPI contains the set of methods on the CasesClient type.
type CasesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string, caseParameter securityinsight.Case) (result securityinsight.Case, err error)
	Delete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, caseID string) (result securityinsight.Case, err error)
	List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result securityinsight.CaseListPage, err error)
}

var _ CasesClientAPI = (*securityinsight.CasesClient)(nil)

// BookmarksClientAPI contains the set of methods on the BookmarksClient type.
type BookmarksClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, bookmarkID string, bookmark securityinsight.Bookmark) (result securityinsight.Bookmark, err error)
	Delete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, bookmarkID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, bookmarkID string) (result securityinsight.Bookmark, err error)
	List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result securityinsight.BookmarkListPage, err error)
}

var _ BookmarksClientAPI = (*securityinsight.BookmarksClient)(nil)

// DataConnectorsClientAPI contains the set of methods on the DataConnectorsClient type.
type DataConnectorsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, dataConnectorID string, dataConnector securityinsight.BasicDataConnector) (result securityinsight.DataConnectorModel, err error)
	Delete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, dataConnectorID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, dataConnectorID string) (result securityinsight.DataConnectorModel, err error)
	List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result securityinsight.DataConnectorListPage, err error)
}

var _ DataConnectorsClientAPI = (*securityinsight.DataConnectorsClient)(nil)

// EntitiesClientAPI contains the set of methods on the EntitiesClient type.
type EntitiesClientAPI interface {
	Expand(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, entityID string, parameters securityinsight.EntityExpandParameters) (result securityinsight.EntityExpandResponse, err error)
	Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, entityID string) (result securityinsight.EntityModel, err error)
	List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result securityinsight.EntityListPage, err error)
}

var _ EntitiesClientAPI = (*securityinsight.EntitiesClient)(nil)

// OfficeConsentsClientAPI contains the set of methods on the OfficeConsentsClient type.
type OfficeConsentsClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, consentID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, consentID string) (result securityinsight.OfficeConsent, err error)
	List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result securityinsight.OfficeConsentListPage, err error)
}

var _ OfficeConsentsClientAPI = (*securityinsight.OfficeConsentsClient)(nil)

// ProductSettingsClientAPI contains the set of methods on the ProductSettingsClient type.
type ProductSettingsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, settingsName string) (result securityinsight.SettingsModel, err error)
	Update(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, settingsName string, settings securityinsight.BasicSettings) (result securityinsight.SettingsModel, err error)
}

var _ ProductSettingsClientAPI = (*securityinsight.ProductSettingsClient)(nil)

// CasesAggregationsClientAPI contains the set of methods on the CasesAggregationsClient type.
type CasesAggregationsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, aggregationsName string) (result securityinsight.AggregationsModel, err error)
}

var _ CasesAggregationsClientAPI = (*securityinsight.CasesAggregationsClient)(nil)

// EntityQueriesClientAPI contains the set of methods on the EntityQueriesClient type.
type EntityQueriesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, entityQueryID string) (result securityinsight.EntityQuery, err error)
	List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result securityinsight.EntityQueryListPage, err error)
}

var _ EntityQueriesClientAPI = (*securityinsight.EntityQueriesClient)(nil)
