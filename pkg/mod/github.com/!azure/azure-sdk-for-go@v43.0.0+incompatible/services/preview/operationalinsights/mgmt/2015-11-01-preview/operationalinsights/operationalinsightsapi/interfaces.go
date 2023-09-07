package operationalinsightsapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/operationalinsights"
	"github.com/Azure/go-autorest/autorest"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	GetAsyncOperationsStatus(ctx context.Context, location string, asyncOperationID string) (result operationalinsights.OperationStatus, err error)
}

var _ BaseClientAPI = (*operationalinsights.BaseClient)(nil)

// LinkedServicesClientAPI contains the set of methods on the LinkedServicesClient type.
type LinkedServicesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, linkedServiceName string, parameters operationalinsights.LinkedService) (result operationalinsights.LinkedService, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, linkedServiceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, linkedServiceName string) (result operationalinsights.LinkedService, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.LinkedServiceListResult, err error)
}

var _ LinkedServicesClientAPI = (*operationalinsights.LinkedServicesClient)(nil)

// DataSourcesClientAPI contains the set of methods on the DataSourcesClient type.
type DataSourcesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceName string, parameters operationalinsights.DataSource) (result operationalinsights.DataSource, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceName string) (result operationalinsights.DataSource, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, filter string, skiptoken string) (result operationalinsights.DataSourceListResultPage, err error)
	ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string, filter string, skiptoken string) (result operationalinsights.DataSourceListResultIterator, err error)
}

var _ DataSourcesClientAPI = (*operationalinsights.DataSourcesClient)(nil)

// WorkspacesClientAPI contains the set of methods on the WorkspacesClient type.
type WorkspacesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, parameters operationalinsights.Workspace) (result operationalinsights.WorkspacesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string) (result autorest.Response, err error)
	DisableIntelligencePack(ctx context.Context, resourceGroupName string, workspaceName string, intelligencePackName string) (result autorest.Response, err error)
	EnableIntelligencePack(ctx context.Context, resourceGroupName string, workspaceName string, intelligencePackName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.Workspace, err error)
	GetSharedKeys(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.SharedKeys, err error)
	List(ctx context.Context) (result operationalinsights.WorkspaceListResult, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result operationalinsights.WorkspaceListResult, err error)
	ListIntelligencePacks(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.ListIntelligencePack, err error)
	ListManagementGroups(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.WorkspaceListManagementGroupsResult, err error)
	ListUsages(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.WorkspaceListUsagesResult, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, parameters operationalinsights.Workspace) (result operationalinsights.Workspace, err error)
}

var _ WorkspacesClientAPI = (*operationalinsights.WorkspacesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result operationalinsights.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result operationalinsights.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*operationalinsights.OperationsClient)(nil)