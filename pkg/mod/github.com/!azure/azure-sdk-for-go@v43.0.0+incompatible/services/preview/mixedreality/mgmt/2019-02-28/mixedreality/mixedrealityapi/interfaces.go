package mixedrealityapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/mixedreality/mgmt/2019-02-28/mixedreality"
	"github.com/Azure/go-autorest/autorest"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	CheckNameAvailabilityLocal(ctx context.Context, location string, checkNameAvailability mixedreality.CheckNameAvailabilityRequest) (result mixedreality.CheckNameAvailabilityResponse, err error)
}

var _ BaseClientAPI = (*mixedreality.BaseClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result mixedreality.OperationListPage, err error)
	ListComplete(ctx context.Context) (result mixedreality.OperationListIterator, err error)
}

var _ OperationsClientAPI = (*mixedreality.OperationsClient)(nil)

// SpatialAnchorsAccountsClientAPI contains the set of methods on the SpatialAnchorsAccountsClient type.
type SpatialAnchorsAccountsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, spatialAnchorsAccountName string, spatialAnchorsAccount mixedreality.SpatialAnchorsAccount) (result mixedreality.SpatialAnchorsAccount, err error)
	Delete(ctx context.Context, resourceGroupName string, spatialAnchorsAccountName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, spatialAnchorsAccountName string) (result mixedreality.SpatialAnchorsAccount, err error)
	GetKeys(ctx context.Context, resourceGroupName string, spatialAnchorsAccountName string) (result mixedreality.SpatialAnchorsAccountKeys, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result mixedreality.SpatialAnchorsAccountListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result mixedreality.SpatialAnchorsAccountListIterator, err error)
	ListBySubscription(ctx context.Context) (result mixedreality.SpatialAnchorsAccountListPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result mixedreality.SpatialAnchorsAccountListIterator, err error)
	RegenerateKeys(ctx context.Context, resourceGroupName string, spatialAnchorsAccountName string, spatialAnchorsAccountKeyRegenerate mixedreality.SpatialAnchorsAccountKeyRegenerateRequest) (result mixedreality.SpatialAnchorsAccountKeys, err error)
	Update(ctx context.Context, resourceGroupName string, spatialAnchorsAccountName string, spatialAnchorsAccount mixedreality.SpatialAnchorsAccount) (result mixedreality.SpatialAnchorsAccount, err error)
}

var _ SpatialAnchorsAccountsClientAPI = (*mixedreality.SpatialAnchorsAccountsClient)(nil)