//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnewrelicobservability

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// TagRulesClient contains the methods for the TagRules group.
// Don't use this type directly, use NewTagRulesClient() instead.
type TagRulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTagRulesClient creates a new instance of TagRulesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTagRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TagRulesClient, error) {
	cl, err := arm.NewClient(moduleName+".TagRulesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TagRulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a TagRule
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Name of the Monitors resource
//   - ruleSetName - Name of the TagRule
//   - resource - Resource create parameters.
//   - options - TagRulesClientBeginCreateOrUpdateOptions contains the optional parameters for the TagRulesClient.BeginCreateOrUpdate
//     method.
func (client *TagRulesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, resource TagRule, options *TagRulesClientBeginCreateOrUpdateOptions) (*runtime.Poller[TagRulesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, monitorName, ruleSetName, resource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TagRulesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[TagRulesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create a TagRule
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-01-preview
func (client *TagRulesClient) createOrUpdate(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, resource TagRule, options *TagRulesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, monitorName, ruleSetName, resource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TagRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, resource TagRule, options *TagRulesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/NewRelic.Observability/monitors/{monitorName}/tagRules/{ruleSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, resource)
}

// BeginDelete - Delete a TagRule
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Name of the Monitors resource
//   - ruleSetName - Name of the TagRule
//   - options - TagRulesClientBeginDeleteOptions contains the optional parameters for the TagRulesClient.BeginDelete method.
func (client *TagRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *TagRulesClientBeginDeleteOptions) (*runtime.Poller[TagRulesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, monitorName, ruleSetName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TagRulesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[TagRulesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete a TagRule
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-01-preview
func (client *TagRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *TagRulesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, monitorName, ruleSetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TagRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *TagRulesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/NewRelic.Observability/monitors/{monitorName}/tagRules/{ruleSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a TagRule
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Name of the Monitors resource
//   - ruleSetName - Name of the TagRule
//   - options - TagRulesClientGetOptions contains the optional parameters for the TagRulesClient.Get method.
func (client *TagRulesClient) Get(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *TagRulesClientGetOptions) (TagRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, monitorName, ruleSetName, options)
	if err != nil {
		return TagRulesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TagRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TagRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TagRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *TagRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/NewRelic.Observability/monitors/{monitorName}/tagRules/{ruleSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TagRulesClient) getHandleResponse(resp *http.Response) (TagRulesClientGetResponse, error) {
	result := TagRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagRule); err != nil {
		return TagRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByNewRelicMonitorResourcePager - List TagRule resources by NewRelicMonitorResource
//
// Generated from API version 2022-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Name of the Monitors resource
//   - options - TagRulesClientListByNewRelicMonitorResourceOptions contains the optional parameters for the TagRulesClient.NewListByNewRelicMonitorResourcePager
//     method.
func (client *TagRulesClient) NewListByNewRelicMonitorResourcePager(resourceGroupName string, monitorName string, options *TagRulesClientListByNewRelicMonitorResourceOptions) *runtime.Pager[TagRulesClientListByNewRelicMonitorResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[TagRulesClientListByNewRelicMonitorResourceResponse]{
		More: func(page TagRulesClientListByNewRelicMonitorResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TagRulesClientListByNewRelicMonitorResourceResponse) (TagRulesClientListByNewRelicMonitorResourceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByNewRelicMonitorResourceCreateRequest(ctx, resourceGroupName, monitorName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TagRulesClientListByNewRelicMonitorResourceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return TagRulesClientListByNewRelicMonitorResourceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TagRulesClientListByNewRelicMonitorResourceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByNewRelicMonitorResourceHandleResponse(resp)
		},
	})
}

// listByNewRelicMonitorResourceCreateRequest creates the ListByNewRelicMonitorResource request.
func (client *TagRulesClient) listByNewRelicMonitorResourceCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *TagRulesClientListByNewRelicMonitorResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/NewRelic.Observability/monitors/{monitorName}/tagRules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByNewRelicMonitorResourceHandleResponse handles the ListByNewRelicMonitorResource response.
func (client *TagRulesClient) listByNewRelicMonitorResourceHandleResponse(resp *http.Response) (TagRulesClientListByNewRelicMonitorResourceResponse, error) {
	result := TagRulesClientListByNewRelicMonitorResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagRuleListResult); err != nil {
		return TagRulesClientListByNewRelicMonitorResourceResponse{}, err
	}
	return result, nil
}

// Update - Update a TagRule
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Name of the Monitors resource
//   - ruleSetName - Name of the TagRule
//   - properties - The resource properties to be updated.
//   - options - TagRulesClientUpdateOptions contains the optional parameters for the TagRulesClient.Update method.
func (client *TagRulesClient) Update(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, properties TagRuleUpdate, options *TagRulesClientUpdateOptions) (TagRulesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, monitorName, ruleSetName, properties, options)
	if err != nil {
		return TagRulesClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TagRulesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TagRulesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *TagRulesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, properties TagRuleUpdate, options *TagRulesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/NewRelic.Observability/monitors/{monitorName}/tagRules/{ruleSetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if ruleSetName == "" {
		return nil, errors.New("parameter ruleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ruleSetName}", url.PathEscape(ruleSetName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, properties)
}

// updateHandleResponse handles the Update response.
func (client *TagRulesClient) updateHandleResponse(resp *http.Response) (TagRulesClientUpdateResponse, error) {
	result := TagRulesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagRule); err != nil {
		return TagRulesClientUpdateResponse{}, err
	}
	return result, nil
}
