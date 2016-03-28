package logic

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.14.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
	"net/url"
)

// WorkflowRunActionsClient is the client for the WorkflowRunActions methods
// of the Logic service.
type WorkflowRunActionsClient struct {
	ManagementClient
}

// NewWorkflowRunActionsClient creates an instance of the
// WorkflowRunActionsClient client.
func NewWorkflowRunActionsClient(subscriptionID string) WorkflowRunActionsClient {
	return NewWorkflowRunActionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkflowRunActionsClientWithBaseURI creates an instance of the
// WorkflowRunActionsClient client.
func NewWorkflowRunActionsClientWithBaseURI(baseURI string, subscriptionID string) WorkflowRunActionsClient {
	return WorkflowRunActionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a workflow run action.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. runName is the workflow run name. actionName is the workflow action
// name.
func (client WorkflowRunActionsClient) Get(resourceGroupName string, workflowName string, runName string, actionName string) (result WorkflowRunAction, err error) {
	req, err := client.GetPreparer(resourceGroupName, workflowName, runName, actionName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowRunActionsClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowRunActionsClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic/WorkflowRunActionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkflowRunActionsClient) GetPreparer(resourceGroupName string, workflowName string, runName string, actionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"actionName":        url.QueryEscape(actionName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"runName":           url.QueryEscape(runName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions/{actionName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowRunActionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkflowRunActionsClient) GetResponder(resp *http.Response) (result WorkflowRunAction, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of workflow run actions.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. runName is the workflow run name. top is the number of items to be
// included in the result. filter is the filter to apply on the operation.
func (client WorkflowRunActionsClient) List(resourceGroupName string, workflowName string, runName string, top *int32, filter string) (result WorkflowRunActionListResult, err error) {
	req, err := client.ListPreparer(resourceGroupName, workflowName, runName, top, filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowRunActionsClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowRunActionsClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic/WorkflowRunActionsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client WorkflowRunActionsClient) ListPreparer(resourceGroupName string, workflowName string, runName string, top *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"runName":           url.QueryEscape(runName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = top
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/actions"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowRunActionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client WorkflowRunActionsClient) ListResponder(resp *http.Response) (result WorkflowRunActionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client WorkflowRunActionsClient) ListNextResults(lastResults WorkflowRunActionListResult) (result WorkflowRunActionListResult, err error) {
	req, err := lastResults.WorkflowRunActionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowRunActionsClient", "List", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowRunActionsClient", "List", resp, "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic/WorkflowRunActionsClient", "List", resp, "Failure responding to next results request request")
	}

	return
}
