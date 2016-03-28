package authorization

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

// ManagementLocksClient is the client for the ManagementLocks methods of the
// Authorization service.
type ManagementLocksClient struct {
	ManagementClient
}

// NewManagementLocksClient creates an instance of the ManagementLocksClient
// client.
func NewManagementLocksClient(subscriptionID string) ManagementLocksClient {
	return NewManagementLocksClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewManagementLocksClientWithBaseURI creates an instance of the
// ManagementLocksClient client.
func NewManagementLocksClientWithBaseURI(baseURI string, subscriptionID string) ManagementLocksClient {
	return ManagementLocksClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdateAtResourceGroupLevel create or update a management lock at
// the resource group level.
//
// resourceGroupName is the resource group name. lockName is the lock name.
// parameters is the management lock parameters.
func (client ManagementLocksClient) CreateOrUpdateAtResourceGroupLevel(resourceGroupName string, lockName string, parameters ManagementLockObject) (result ManagementLockObject, err error) {
	req, err := client.CreateOrUpdateAtResourceGroupLevelPreparer(resourceGroupName, lockName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "CreateOrUpdateAtResourceGroupLevel", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateAtResourceGroupLevelSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "CreateOrUpdateAtResourceGroupLevel", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateAtResourceGroupLevelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "CreateOrUpdateAtResourceGroupLevel", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateAtResourceGroupLevelPreparer prepares the CreateOrUpdateAtResourceGroupLevel request.
func (client ManagementLocksClient) CreateOrUpdateAtResourceGroupLevelPreparer(resourceGroupName string, lockName string, parameters ManagementLockObject) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"lockName":          url.QueryEscape(lockName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/locks/{lockName}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateAtResourceGroupLevelSender sends the CreateOrUpdateAtResourceGroupLevel request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementLocksClient) CreateOrUpdateAtResourceGroupLevelSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateAtResourceGroupLevelResponder handles the response to the CreateOrUpdateAtResourceGroupLevel request. The method always
// closes the http.Response Body.
func (client ManagementLocksClient) CreateOrUpdateAtResourceGroupLevelResponder(resp *http.Response) (result ManagementLockObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdateAtResourceLevel create or update a management lock at the
// resource level or any level below resource.
//
// resourceGroupName is the name of the resource group.
// resourceProviderNamespace is resource identity. parentResourcePath is
// resource identity. resourceType is resource identity. resourceName is
// resource identity. lockName is the name of lock. parameters is create or
// update management lock parameters.
func (client ManagementLocksClient) CreateOrUpdateAtResourceLevel(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, lockName string, parameters ManagementLockObject) (result ManagementLockObject, err error) {
	req, err := client.CreateOrUpdateAtResourceLevelPreparer(resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, lockName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "CreateOrUpdateAtResourceLevel", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateAtResourceLevelSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "CreateOrUpdateAtResourceLevel", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateAtResourceLevelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "CreateOrUpdateAtResourceLevel", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateAtResourceLevelPreparer prepares the CreateOrUpdateAtResourceLevel request.
func (client ManagementLocksClient) CreateOrUpdateAtResourceLevelPreparer(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, lockName string, parameters ManagementLockObject) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"lockName":                  url.QueryEscape(lockName),
		"parentResourcePath":        parentResourcePath,
		"resourceGroupName":         url.QueryEscape(resourceGroupName),
		"resourceName":              url.QueryEscape(resourceName),
		"resourceProviderNamespace": url.QueryEscape(resourceProviderNamespace),
		"resourceType":              resourceType,
		"subscriptionId":            url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/locks/{lockName}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateAtResourceLevelSender sends the CreateOrUpdateAtResourceLevel request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementLocksClient) CreateOrUpdateAtResourceLevelSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateAtResourceLevelResponder handles the response to the CreateOrUpdateAtResourceLevel request. The method always
// closes the http.Response Body.
func (client ManagementLocksClient) CreateOrUpdateAtResourceLevelResponder(resp *http.Response) (result ManagementLockObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdateAtSubscriptionLevel create or update a management lock at the
// subscription level.
//
// lockName is the name of lock. parameters is the management lock parameters.
func (client ManagementLocksClient) CreateOrUpdateAtSubscriptionLevel(lockName string, parameters ManagementLockObject) (result ManagementLockObject, err error) {
	req, err := client.CreateOrUpdateAtSubscriptionLevelPreparer(lockName, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "CreateOrUpdateAtSubscriptionLevel", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateAtSubscriptionLevelSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "CreateOrUpdateAtSubscriptionLevel", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateAtSubscriptionLevelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "CreateOrUpdateAtSubscriptionLevel", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateAtSubscriptionLevelPreparer prepares the CreateOrUpdateAtSubscriptionLevel request.
func (client ManagementLocksClient) CreateOrUpdateAtSubscriptionLevelPreparer(lockName string, parameters ManagementLockObject) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"lockName":       url.QueryEscape(lockName),
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/locks/{lockName}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateAtSubscriptionLevelSender sends the CreateOrUpdateAtSubscriptionLevel request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementLocksClient) CreateOrUpdateAtSubscriptionLevelSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateAtSubscriptionLevelResponder handles the response to the CreateOrUpdateAtSubscriptionLevel request. The method always
// closes the http.Response Body.
func (client ManagementLocksClient) CreateOrUpdateAtSubscriptionLevelResponder(resp *http.Response) (result ManagementLockObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteAtResourceGroupLevel deletes the management lock of a resource group.
//
// resourceGroup is the resource group names. lockName is the name of lock.
func (client ManagementLocksClient) DeleteAtResourceGroupLevel(resourceGroup string, lockName string) (result autorest.Response, err error) {
	req, err := client.DeleteAtResourceGroupLevelPreparer(resourceGroup, lockName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "DeleteAtResourceGroupLevel", nil, "Failure preparing request")
	}

	resp, err := client.DeleteAtResourceGroupLevelSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "DeleteAtResourceGroupLevel", resp, "Failure sending request")
	}

	result, err = client.DeleteAtResourceGroupLevelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "DeleteAtResourceGroupLevel", resp, "Failure responding to request")
	}

	return
}

// DeleteAtResourceGroupLevelPreparer prepares the DeleteAtResourceGroupLevel request.
func (client ManagementLocksClient) DeleteAtResourceGroupLevelPreparer(resourceGroup string, lockName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"lockName":       url.QueryEscape(lockName),
		"resourceGroup":  url.QueryEscape(resourceGroup),
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Authorization/locks/{lockName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteAtResourceGroupLevelSender sends the DeleteAtResourceGroupLevel request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementLocksClient) DeleteAtResourceGroupLevelSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteAtResourceGroupLevelResponder handles the response to the DeleteAtResourceGroupLevel request. The method always
// closes the http.Response Body.
func (client ManagementLocksClient) DeleteAtResourceGroupLevelResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteAtResourceLevel deletes the management lock of a resource or any
// level below resource.
//
// resourceGroupName is the name of the resource group.
// resourceProviderNamespace is resource identity. parentResourcePath is
// resource identity. resourceType is resource identity. resourceName is
// resource identity. lockName is the name of lock.
func (client ManagementLocksClient) DeleteAtResourceLevel(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, lockName string) (result autorest.Response, err error) {
	req, err := client.DeleteAtResourceLevelPreparer(resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, lockName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "DeleteAtResourceLevel", nil, "Failure preparing request")
	}

	resp, err := client.DeleteAtResourceLevelSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "DeleteAtResourceLevel", resp, "Failure sending request")
	}

	result, err = client.DeleteAtResourceLevelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "DeleteAtResourceLevel", resp, "Failure responding to request")
	}

	return
}

// DeleteAtResourceLevelPreparer prepares the DeleteAtResourceLevel request.
func (client ManagementLocksClient) DeleteAtResourceLevelPreparer(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, lockName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"lockName":                  url.QueryEscape(lockName),
		"parentResourcePath":        parentResourcePath,
		"resourceGroupName":         url.QueryEscape(resourceGroupName),
		"resourceName":              url.QueryEscape(resourceName),
		"resourceProviderNamespace": url.QueryEscape(resourceProviderNamespace),
		"resourceType":              resourceType,
		"subscriptionId":            url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/locks/{lockName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteAtResourceLevelSender sends the DeleteAtResourceLevel request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementLocksClient) DeleteAtResourceLevelSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteAtResourceLevelResponder handles the response to the DeleteAtResourceLevel request. The method always
// closes the http.Response Body.
func (client ManagementLocksClient) DeleteAtResourceLevelResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteAtSubscriptionLevel deletes the management lock of a subscription.
//
// lockName is the name of lock.
func (client ManagementLocksClient) DeleteAtSubscriptionLevel(lockName string) (result autorest.Response, err error) {
	req, err := client.DeleteAtSubscriptionLevelPreparer(lockName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "DeleteAtSubscriptionLevel", nil, "Failure preparing request")
	}

	resp, err := client.DeleteAtSubscriptionLevelSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "DeleteAtSubscriptionLevel", resp, "Failure sending request")
	}

	result, err = client.DeleteAtSubscriptionLevelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "DeleteAtSubscriptionLevel", resp, "Failure responding to request")
	}

	return
}

// DeleteAtSubscriptionLevelPreparer prepares the DeleteAtSubscriptionLevel request.
func (client ManagementLocksClient) DeleteAtSubscriptionLevelPreparer(lockName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"lockName":       url.QueryEscape(lockName),
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/locks/{lockName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteAtSubscriptionLevelSender sends the DeleteAtSubscriptionLevel request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementLocksClient) DeleteAtSubscriptionLevelSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteAtSubscriptionLevelResponder handles the response to the DeleteAtSubscriptionLevel request. The method always
// closes the http.Response Body.
func (client ManagementLocksClient) DeleteAtSubscriptionLevelResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the management lock of a scope.
//
// lockName is name of the management lock.
func (client ManagementLocksClient) Get(lockName string) (result ManagementLockObject, err error) {
	req, err := client.GetPreparer(lockName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ManagementLocksClient) GetPreparer(lockName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"lockName":       url.QueryEscape(lockName),
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/locks/{lockName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementLocksClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ManagementLocksClient) GetResponder(resp *http.Response) (result ManagementLockObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAtResourceGroupLevel gets all the management locks of a resource group.
//
// resourceGroupName is resource group name. filter is the filter to apply on
// the operation.
func (client ManagementLocksClient) ListAtResourceGroupLevel(resourceGroupName string, filter string) (result ManagementLockListResult, err error) {
	req, err := client.ListAtResourceGroupLevelPreparer(resourceGroupName, filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListAtResourceGroupLevel", nil, "Failure preparing request")
	}

	resp, err := client.ListAtResourceGroupLevelSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListAtResourceGroupLevel", resp, "Failure sending request")
	}

	result, err = client.ListAtResourceGroupLevelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListAtResourceGroupLevel", resp, "Failure responding to request")
	}

	return
}

// ListAtResourceGroupLevelPreparer prepares the ListAtResourceGroupLevel request.
func (client ManagementLocksClient) ListAtResourceGroupLevelPreparer(resourceGroupName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/locks"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListAtResourceGroupLevelSender sends the ListAtResourceGroupLevel request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementLocksClient) ListAtResourceGroupLevelSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListAtResourceGroupLevelResponder handles the response to the ListAtResourceGroupLevel request. The method always
// closes the http.Response Body.
func (client ManagementLocksClient) ListAtResourceGroupLevelResponder(resp *http.Response) (result ManagementLockListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAtResourceGroupLevelNextResults retrieves the next set of results, if any.
func (client ManagementLocksClient) ListAtResourceGroupLevelNextResults(lastResults ManagementLockListResult) (result ManagementLockListResult, err error) {
	req, err := lastResults.ManagementLockListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListAtResourceGroupLevel", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListAtResourceGroupLevelSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListAtResourceGroupLevel", resp, "Failure sending next results request request")
	}

	result, err = client.ListAtResourceGroupLevelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListAtResourceGroupLevel", resp, "Failure responding to next results request request")
	}

	return
}

// ListAtResourceLevel gets all the management locks of a resource or any
// level below resource.
//
// resourceGroupName is the name of the resource group. The name is case
// insensitive. resourceProviderNamespace is resource identity.
// parentResourcePath is resource identity. resourceType is resource
// identity. resourceName is resource identity. filter is the filter to apply
// on the operation.
func (client ManagementLocksClient) ListAtResourceLevel(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (result ManagementLockListResult, err error) {
	req, err := client.ListAtResourceLevelPreparer(resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListAtResourceLevel", nil, "Failure preparing request")
	}

	resp, err := client.ListAtResourceLevelSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListAtResourceLevel", resp, "Failure sending request")
	}

	result, err = client.ListAtResourceLevelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListAtResourceLevel", resp, "Failure responding to request")
	}

	return
}

// ListAtResourceLevelPreparer prepares the ListAtResourceLevel request.
func (client ManagementLocksClient) ListAtResourceLevelPreparer(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentResourcePath":        parentResourcePath,
		"resourceGroupName":         url.QueryEscape(resourceGroupName),
		"resourceName":              url.QueryEscape(resourceName),
		"resourceProviderNamespace": url.QueryEscape(resourceProviderNamespace),
		"resourceType":              resourceType,
		"subscriptionId":            url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/locks"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListAtResourceLevelSender sends the ListAtResourceLevel request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementLocksClient) ListAtResourceLevelSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListAtResourceLevelResponder handles the response to the ListAtResourceLevel request. The method always
// closes the http.Response Body.
func (client ManagementLocksClient) ListAtResourceLevelResponder(resp *http.Response) (result ManagementLockListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAtResourceLevelNextResults retrieves the next set of results, if any.
func (client ManagementLocksClient) ListAtResourceLevelNextResults(lastResults ManagementLockListResult) (result ManagementLockListResult, err error) {
	req, err := lastResults.ManagementLockListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListAtResourceLevel", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListAtResourceLevelSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListAtResourceLevel", resp, "Failure sending next results request request")
	}

	result, err = client.ListAtResourceLevelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListAtResourceLevel", resp, "Failure responding to next results request request")
	}

	return
}

// ListAtSubscriptionLevel gets all the management locks of a subscription.
//
// filter is the filter to apply on the operation.
func (client ManagementLocksClient) ListAtSubscriptionLevel(filter string) (result ManagementLockListResult, err error) {
	req, err := client.ListAtSubscriptionLevelPreparer(filter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListAtSubscriptionLevel", nil, "Failure preparing request")
	}

	resp, err := client.ListAtSubscriptionLevelSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListAtSubscriptionLevel", resp, "Failure sending request")
	}

	result, err = client.ListAtSubscriptionLevelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListAtSubscriptionLevel", resp, "Failure responding to request")
	}

	return
}

// ListAtSubscriptionLevelPreparer prepares the ListAtSubscriptionLevel request.
func (client ManagementLocksClient) ListAtSubscriptionLevelPreparer(filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/locks"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListAtSubscriptionLevelSender sends the ListAtSubscriptionLevel request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementLocksClient) ListAtSubscriptionLevelSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListAtSubscriptionLevelResponder handles the response to the ListAtSubscriptionLevel request. The method always
// closes the http.Response Body.
func (client ManagementLocksClient) ListAtSubscriptionLevelResponder(resp *http.Response) (result ManagementLockListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAtSubscriptionLevelNextResults retrieves the next set of results, if any.
func (client ManagementLocksClient) ListAtSubscriptionLevelNextResults(lastResults ManagementLockListResult) (result ManagementLockListResult, err error) {
	req, err := lastResults.ManagementLockListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListAtSubscriptionLevel", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListAtSubscriptionLevelSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListAtSubscriptionLevel", resp, "Failure sending next results request request")
	}

	result, err = client.ListAtSubscriptionLevelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListAtSubscriptionLevel", resp, "Failure responding to next results request request")
	}

	return
}

// ListNext get a list of management locks at resource level or below.
//
// nextLink is nextLink from the previous successful call to List operation.
func (client ManagementLocksClient) ListNext(nextLink string) (result ManagementLockListResult, err error) {
	req, err := client.ListNextPreparer(nextLink)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListNext", nil, "Failure preparing request")
	}

	resp, err := client.ListNextSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListNext", resp, "Failure sending request")
	}

	result, err = client.ListNextResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListNext", resp, "Failure responding to request")
	}

	return
}

// ListNextPreparer prepares the ListNext request.
func (client ManagementLocksClient) ListNextPreparer(nextLink string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nextLink":       nextLink,
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/{nextLink}"),
		autorest.WithPathParameters(pathParameters))
}

// ListNextSender sends the ListNext request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementLocksClient) ListNextSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListNextResponder handles the response to the ListNext request. The method always
// closes the http.Response Body.
func (client ManagementLocksClient) ListNextResponder(resp *http.Response) (result ManagementLockListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextNextResults retrieves the next set of results, if any.
func (client ManagementLocksClient) ListNextNextResults(lastResults ManagementLockListResult) (result ManagementLockListResult, err error) {
	req, err := lastResults.ManagementLockListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListNext", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListNextSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListNext", resp, "Failure sending next results request request")
	}

	result, err = client.ListNextResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization/ManagementLocksClient", "ListNext", resp, "Failure responding to next results request request")
	}

	return
}
