//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armconfidentialledger

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

// Client contains the methods for the ConfidentialLedger group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	internal       *arm.Client
	subscriptionID string
}

// NewClient creates a new instance of Client with the specified values.
//   - subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*Client, error) {
	cl, err := arm.NewClient(moduleName+".Client", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &Client{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckNameAvailability - To check whether a resource name is available.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-13
//   - nameAvailabilityRequest - Name availability request payload.
//   - options - ClientCheckNameAvailabilityOptions contains the optional parameters for the Client.CheckNameAvailability method.
func (client *Client) CheckNameAvailability(ctx context.Context, nameAvailabilityRequest CheckNameAvailabilityRequest, options *ClientCheckNameAvailabilityOptions) (ClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, nameAvailabilityRequest, options)
	if err != nil {
		return ClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *Client) checkNameAvailabilityCreateRequest(ctx context.Context, nameAvailabilityRequest CheckNameAvailabilityRequest, options *ClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ConfidentialLedger/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-13")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, nameAvailabilityRequest)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *Client) checkNameAvailabilityHandleResponse(resp *http.Response) (ClientCheckNameAvailabilityResponse, error) {
	result := ClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResponse); err != nil {
		return ClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}
