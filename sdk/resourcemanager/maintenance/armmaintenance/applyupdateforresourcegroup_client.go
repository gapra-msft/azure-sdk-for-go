//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmaintenance

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

// ApplyUpdateForResourceGroupClient contains the methods for the ApplyUpdateForResourceGroup group.
// Don't use this type directly, use NewApplyUpdateForResourceGroupClient() instead.
type ApplyUpdateForResourceGroupClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewApplyUpdateForResourceGroupClient creates a new instance of ApplyUpdateForResourceGroupClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewApplyUpdateForResourceGroupClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ApplyUpdateForResourceGroupClient, error) {
	cl, err := arm.NewClient(moduleName+".ApplyUpdateForResourceGroupClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ApplyUpdateForResourceGroupClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Get Configuration records within a subscription and resource group
//
// Generated from API version 2022-07-01-preview
//   - resourceGroupName - Resource Group Name
//   - options - ApplyUpdateForResourceGroupClientListOptions contains the optional parameters for the ApplyUpdateForResourceGroupClient.NewListPager
//     method.
func (client *ApplyUpdateForResourceGroupClient) NewListPager(resourceGroupName string, options *ApplyUpdateForResourceGroupClientListOptions) *runtime.Pager[ApplyUpdateForResourceGroupClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplyUpdateForResourceGroupClientListResponse]{
		More: func(page ApplyUpdateForResourceGroupClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ApplyUpdateForResourceGroupClientListResponse) (ApplyUpdateForResourceGroupClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, options)
			if err != nil {
				return ApplyUpdateForResourceGroupClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ApplyUpdateForResourceGroupClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ApplyUpdateForResourceGroupClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ApplyUpdateForResourceGroupClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *ApplyUpdateForResourceGroupClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Maintenance/applyUpdates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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

// listHandleResponse handles the List response.
func (client *ApplyUpdateForResourceGroupClient) listHandleResponse(resp *http.Response) (ApplyUpdateForResourceGroupClientListResponse, error) {
	result := ApplyUpdateForResourceGroupClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListApplyUpdate); err != nil {
		return ApplyUpdateForResourceGroupClientListResponse{}, err
	}
	return result, nil
}
