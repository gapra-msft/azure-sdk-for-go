// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresources

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// ProviderResourceTypesClient contains the methods for the ProviderResourceTypes group.
// Don't use this type directly, use NewProviderResourceTypesClient() instead.
type ProviderResourceTypesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewProviderResourceTypesClient creates a new instance of ProviderResourceTypesClient with the specified values.
func NewProviderResourceTypesClient(con *armcore.Connection, subscriptionID string) *ProviderResourceTypesClient {
	return &ProviderResourceTypesClient{con: con, subscriptionID: subscriptionID}
}

// List - List the resource types for a specified resource provider.
// If the operation fails it returns the *CloudError error type.
func (client *ProviderResourceTypesClient) List(ctx context.Context, resourceProviderNamespace string, options *ProviderResourceTypesListOptions) (ProviderResourceTypeListResultResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceProviderNamespace, options)
	if err != nil {
		return ProviderResourceTypeListResultResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ProviderResourceTypeListResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ProviderResourceTypeListResultResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ProviderResourceTypesClient) listCreateRequest(ctx context.Context, resourceProviderNamespace string, options *ProviderResourceTypesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/{resourceProviderNamespace}/resourceTypes"
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2021-04-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ProviderResourceTypesClient) listHandleResponse(resp *azcore.Response) (ProviderResourceTypeListResultResponse, error) {
	var val *ProviderResourceTypeListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ProviderResourceTypeListResultResponse{}, err
	}
	return ProviderResourceTypeListResultResponse{RawResponse: resp.Response, ProviderResourceTypeListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *ProviderResourceTypesClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
