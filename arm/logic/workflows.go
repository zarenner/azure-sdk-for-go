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

// WorkflowsClient is the client for the Workflows methods of the Logic
// service.
type WorkflowsClient struct {
	ManagementClient
}

// NewWorkflowsClient creates an instance of the WorkflowsClient client.
func NewWorkflowsClient(subscriptionID string) WorkflowsClient {
	return NewWorkflowsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkflowsClientWithBaseURI creates an instance of the WorkflowsClient
// client.
func NewWorkflowsClientWithBaseURI(baseURI string, subscriptionID string) WorkflowsClient {
	return WorkflowsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a workflow.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. workflow is the workflow.
func (client WorkflowsClient) CreateOrUpdate(resourceGroupName string, workflowName string, workflow Workflow) (result Workflow, err error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, workflowName, workflow)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "CreateOrUpdate", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "CreateOrUpdate", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic/WorkflowsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client WorkflowsClient) CreateOrUpdatePreparer(resourceGroupName string, workflowName string, workflow Workflow) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}"),
		autorest.WithJSON(workflow),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client WorkflowsClient) CreateOrUpdateResponder(resp *http.Response) (result Workflow, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a workflow.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name.
func (client WorkflowsClient) Delete(resourceGroupName string, workflowName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, workflowName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic/WorkflowsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client WorkflowsClient) DeletePreparer(resourceGroupName string, workflowName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client WorkflowsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Disable disables a workflow.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name.
func (client WorkflowsClient) Disable(resourceGroupName string, workflowName string) (result autorest.Response, err error) {
	req, err := client.DisablePreparer(resourceGroupName, workflowName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "Disable", nil, "Failure preparing request")
	}

	resp, err := client.DisableSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "Disable", resp, "Failure sending request")
	}

	result, err = client.DisableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic/WorkflowsClient", "Disable", resp, "Failure responding to request")
	}

	return
}

// DisablePreparer prepares the Disable request.
func (client WorkflowsClient) DisablePreparer(resourceGroupName string, workflowName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/disable"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DisableSender sends the Disable request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsClient) DisableSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DisableResponder handles the response to the Disable request. The method always
// closes the http.Response Body.
func (client WorkflowsClient) DisableResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Enable enables a workflow.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name.
func (client WorkflowsClient) Enable(resourceGroupName string, workflowName string) (result autorest.Response, err error) {
	req, err := client.EnablePreparer(resourceGroupName, workflowName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "Enable", nil, "Failure preparing request")
	}

	resp, err := client.EnableSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "Enable", resp, "Failure sending request")
	}

	result, err = client.EnableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic/WorkflowsClient", "Enable", resp, "Failure responding to request")
	}

	return
}

// EnablePreparer prepares the Enable request.
func (client WorkflowsClient) EnablePreparer(resourceGroupName string, workflowName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/enable"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// EnableSender sends the Enable request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsClient) EnableSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// EnableResponder handles the response to the Enable request. The method always
// closes the http.Response Body.
func (client WorkflowsClient) EnableResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a workflow.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name.
func (client WorkflowsClient) Get(resourceGroupName string, workflowName string) (result Workflow, err error) {
	req, err := client.GetPreparer(resourceGroupName, workflowName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic/WorkflowsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkflowsClient) GetPreparer(resourceGroupName string, workflowName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
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
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkflowsClient) GetResponder(resp *http.Response) (result Workflow, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup gets a list of workflows by resource group.
//
// resourceGroupName is the resource group name. top is the number of items to
// be included in the result. filter is the filter to apply on the operation.
func (client WorkflowsClient) ListByResourceGroup(resourceGroupName string, top *int32, filter string) (result WorkflowListResult, err error) {
	req, err := client.ListByResourceGroupPreparer(resourceGroupName, top, filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "ListByResourceGroup", nil, "Failure preparing request")
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "ListByResourceGroup", resp, "Failure sending request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic/WorkflowsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client WorkflowsClient) ListByResourceGroupPreparer(resourceGroupName string, top *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
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
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client WorkflowsClient) ListByResourceGroupResponder(resp *http.Response) (result WorkflowListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroupNextResults retrieves the next set of results, if any.
func (client WorkflowsClient) ListByResourceGroupNextResults(lastResults WorkflowListResult) (result WorkflowListResult, err error) {
	req, err := lastResults.WorkflowListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "ListByResourceGroup", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "ListByResourceGroup", resp, "Failure sending next results request request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic/WorkflowsClient", "ListByResourceGroup", resp, "Failure responding to next results request request")
	}

	return
}

// ListBySubscription gets a list of workflows by subscription.
//
// top is the number of items to be included in the result. filter is the
// filter to apply on the operation.
func (client WorkflowsClient) ListBySubscription(top *int32, filter string) (result WorkflowListResult, err error) {
	req, err := client.ListBySubscriptionPreparer(top, filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "ListBySubscription", nil, "Failure preparing request")
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "ListBySubscription", resp, "Failure sending request")
	}

	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic/WorkflowsClient", "ListBySubscription", resp, "Failure responding to request")
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client WorkflowsClient) ListBySubscriptionPreparer(top *int32, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
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
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Logic/workflows"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client WorkflowsClient) ListBySubscriptionResponder(resp *http.Response) (result WorkflowListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySubscriptionNextResults retrieves the next set of results, if any.
func (client WorkflowsClient) ListBySubscriptionNextResults(lastResults WorkflowListResult) (result WorkflowListResult, err error) {
	req, err := lastResults.WorkflowListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "ListBySubscription", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "ListBySubscription", resp, "Failure sending next results request request")
	}

	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic/WorkflowsClient", "ListBySubscription", resp, "Failure responding to next results request request")
	}

	return
}

// Run runs a workflow. This method handles polling itself for long-running
// operations. User can cancel polling for long-running calls by passing
// cancel channel as an argument to this method (cancel <-chan struct{}).
// This channel won't cancel the operation, it will only stop the polling.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. parameters is the parameters.
func (client WorkflowsClient) Run(resourceGroupName string, workflowName string, parameters RunWorkflowParameters, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.RunPreparer(resourceGroupName, workflowName, parameters, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "Run", nil, "Failure preparing request")
	}

	resp, err := client.RunSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "Run", resp, "Failure sending request")
	}

	result, err = client.RunResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic/WorkflowsClient", "Run", resp, "Failure responding to request")
	}

	return
}

// RunPreparer prepares the Run request.
func (client WorkflowsClient) RunPreparer(resourceGroupName string, workflowName string, parameters RunWorkflowParameters, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{Cancel: cancel}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/run"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// RunSender sends the Run request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsClient) RunSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(autorest.DefaultPollingDelay))
}

// RunResponder handles the response to the Run request. The method always
// closes the http.Response Body.
func (client WorkflowsClient) RunResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update updates a workflow.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. workflow is the workflow.
func (client WorkflowsClient) Update(resourceGroupName string, workflowName string, workflow Workflow) (result Workflow, err error) {
	req, err := client.UpdatePreparer(resourceGroupName, workflowName, workflow)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "Update", nil, "Failure preparing request")
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "Update", resp, "Failure sending request")
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic/WorkflowsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client WorkflowsClient) UpdatePreparer(resourceGroupName string, workflowName string, workflow Workflow) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}"),
		autorest.WithJSON(workflow),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client WorkflowsClient) UpdateResponder(resp *http.Response) (result Workflow, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Validate validates a workflow.
//
// resourceGroupName is the resource group name. workflowName is the workflow
// name. workflow is the workflow.
func (client WorkflowsClient) Validate(resourceGroupName string, workflowName string, workflow Workflow) (result autorest.Response, err error) {
	req, err := client.ValidatePreparer(resourceGroupName, workflowName, workflow)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "Validate", nil, "Failure preparing request")
	}

	resp, err := client.ValidateSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "logic/WorkflowsClient", "Validate", resp, "Failure sending request")
	}

	result, err = client.ValidateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logic/WorkflowsClient", "Validate", resp, "Failure responding to request")
	}

	return
}

// ValidatePreparer prepares the Validate request.
func (client WorkflowsClient) ValidatePreparer(resourceGroupName string, workflowName string, workflow Workflow) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
		"workflowName":      url.QueryEscape(workflowName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/validate"),
		autorest.WithJSON(workflow),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ValidateSender sends the Validate request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsClient) ValidateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ValidateResponder handles the response to the Validate request. The method always
// closes the http.Response Body.
func (client WorkflowsClient) ValidateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
