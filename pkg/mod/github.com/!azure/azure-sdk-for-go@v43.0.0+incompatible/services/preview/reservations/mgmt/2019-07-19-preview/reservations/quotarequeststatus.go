package reservations

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

// QuotaRequestStatusClient is the client for the QuotaRequestStatus methods of the Reservations service.
type QuotaRequestStatusClient struct {
	BaseClient
}

// NewQuotaRequestStatusClient creates an instance of the QuotaRequestStatusClient client.
func NewQuotaRequestStatusClient() QuotaRequestStatusClient {
	return NewQuotaRequestStatusClientWithBaseURI(DefaultBaseURI)
}

// NewQuotaRequestStatusClientWithBaseURI creates an instance of the QuotaRequestStatusClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewQuotaRequestStatusClientWithBaseURI(baseURI string) QuotaRequestStatusClient {
	return QuotaRequestStatusClient{NewWithBaseURI(baseURI)}
}

// Get gets the QuotaRequest details and status by the quota request Id for the resources for the resource provider at
// a specific location. The requestId is returned as response to the Put requests for serviceLimits.
// Parameters:
// subscriptionID - azure subscription id.
// providerID - azure resource Provider id.
// location - azure region.
// ID - quota Request id.
func (client QuotaRequestStatusClient) Get(ctx context.Context, subscriptionID string, providerID string, location string, ID string) (result QuotaRequestDetails, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QuotaRequestStatusClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, subscriptionID, providerID, location, ID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.QuotaRequestStatusClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "reservations.QuotaRequestStatusClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.QuotaRequestStatusClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client QuotaRequestStatusClient) GetPreparer(ctx context.Context, subscriptionID string, providerID string, location string, ID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id":             autorest.Encode("path", ID),
		"location":       autorest.Encode("path", location),
		"providerId":     autorest.Encode("path", providerID),
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2019-07-19-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/resourceProviders/{providerId}/locations/{location}/serviceLimitsRequests/{id}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client QuotaRequestStatusClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client QuotaRequestStatusClient) GetResponder(resp *http.Response) (result QuotaRequestDetails, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List for the specified location and Resource provider gets the current quota requests under the subscription over
// the time period of one year ago from now to one year back. oData filter can be used to select quota requests.
// Parameters:
// subscriptionID - azure subscription id.
// providerID - azure resource Provider id.
// location - azure region.
// filter - | Field                    | Supported operators
// |---------------------|------------------------
//
// |requestSubmitTime | ge, le, eq, gt, lt
// top - number of records to return.
// skiptoken - skiptoken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skiptoken parameter that
// specifies a starting point to use for subsequent calls
func (client QuotaRequestStatusClient) List(ctx context.Context, subscriptionID string, providerID string, location string, filter string, top *int32, skiptoken string) (result QuotaRequestDetailsListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QuotaRequestStatusClient.List")
		defer func() {
			sc := -1
			if result.qrdl.Response.Response != nil {
				sc = result.qrdl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("reservations.QuotaRequestStatusClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, subscriptionID, providerID, location, filter, top, skiptoken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.QuotaRequestStatusClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.qrdl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "reservations.QuotaRequestStatusClient", "List", resp, "Failure sending request")
		return
	}

	result.qrdl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.QuotaRequestStatusClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client QuotaRequestStatusClient) ListPreparer(ctx context.Context, subscriptionID string, providerID string, location string, filter string, top *int32, skiptoken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"providerId":     autorest.Encode("path", providerID),
		"subscriptionId": autorest.Encode("path", subscriptionID),
	}

	const APIVersion = "2019-07-19-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Capacity/resourceProviders/{providerId}/locations/{location}/serviceLimitsRequests", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client QuotaRequestStatusClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client QuotaRequestStatusClient) ListResponder(resp *http.Response) (result QuotaRequestDetailsList, err error) {
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
func (client QuotaRequestStatusClient) listNextResults(ctx context.Context, lastResults QuotaRequestDetailsList) (result QuotaRequestDetailsList, err error) {
	req, err := lastResults.quotaRequestDetailsListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "reservations.QuotaRequestStatusClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "reservations.QuotaRequestStatusClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservations.QuotaRequestStatusClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client QuotaRequestStatusClient) ListComplete(ctx context.Context, subscriptionID string, providerID string, location string, filter string, top *int32, skiptoken string) (result QuotaRequestDetailsListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QuotaRequestStatusClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, subscriptionID, providerID, location, filter, top, skiptoken)
	return
}
