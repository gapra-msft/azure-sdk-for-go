//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf3574813e15bb33b3cb610f44edfcbebd8b1b23/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/preview/2021-12-01-preview/examples/WorkspacesSubscriptionList.json
func ExampleWorkspacesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacesClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.WorkspaceListResult = armoperationalinsights.WorkspaceListResult{
		// 	Value: []*armoperationalinsights.Workspace{
		// 		{
		// 			Name: to.Ptr("AzTest2170"),
		// 			Type: to.Ptr("Microsoft.OperationalInsights/workspaces"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000005/resourcegroups/oiautorest6685/providers/microsoft.operationalinsights/workspaces/aztest2170"),
		// 			Location: to.Ptr("australiasoutheast"),
		// 			Tags: map[string]*string{
		// 				"tag1": to.Ptr("val1"),
		// 			},
		// 			Properties: &armoperationalinsights.WorkspaceProperties{
		// 				CustomerID: to.Ptr("bc089d7b-485c-4aff-a71e-c00f362d8d2f"),
		// 				ProvisioningState: to.Ptr(armoperationalinsights.WorkspaceEntityStatusSucceeded),
		// 				PublicNetworkAccessForIngestion: to.Ptr(armoperationalinsights.PublicNetworkAccessTypeEnabled),
		// 				PublicNetworkAccessForQuery: to.Ptr(armoperationalinsights.PublicNetworkAccessTypeEnabled),
		// 				RetentionInDays: to.Ptr[int32](30),
		// 				SKU: &armoperationalinsights.WorkspaceSKU{
		// 					Name: to.Ptr(armoperationalinsights.WorkspaceSKUNameEnumPerGB2018),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf3574813e15bb33b3cb610f44edfcbebd8b1b23/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/preview/2021-12-01-preview/examples/WorkspacesListByResourceGroup.json
func ExampleWorkspacesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacesClient().NewListByResourceGroupPager("oiautorest6685", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.WorkspaceListResult = armoperationalinsights.WorkspaceListResult{
		// 	Value: []*armoperationalinsights.Workspace{
		// 		{
		// 			Name: to.Ptr("AzTest2170"),
		// 			Type: to.Ptr("Microsoft.OperationalInsights/workspaces"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000005/resourcegroups/oiautorest6685/providers/microsoft.operationalinsights/workspaces/aztest2170"),
		// 			Location: to.Ptr("australiasoutheast"),
		// 			Tags: map[string]*string{
		// 				"tag1": to.Ptr("val1"),
		// 			},
		// 			Properties: &armoperationalinsights.WorkspaceProperties{
		// 				CustomerID: to.Ptr("bc089d7b-485c-4aff-a71e-c00f362d8d2f"),
		// 				ProvisioningState: to.Ptr(armoperationalinsights.WorkspaceEntityStatusSucceeded),
		// 				PublicNetworkAccessForIngestion: to.Ptr(armoperationalinsights.PublicNetworkAccessTypeEnabled),
		// 				PublicNetworkAccessForQuery: to.Ptr(armoperationalinsights.PublicNetworkAccessTypeEnabled),
		// 				RetentionInDays: to.Ptr[int32](30),
		// 				SKU: &armoperationalinsights.WorkspaceSKU{
		// 					Name: to.Ptr(armoperationalinsights.WorkspaceSKUNameEnumPerGB2018),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf3574813e15bb33b3cb610f44edfcbebd8b1b23/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/preview/2021-12-01-preview/examples/WorkspacesCreate.json
func ExampleWorkspacesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspacesClient().BeginCreateOrUpdate(ctx, "oiautorest6685", "oiautorest6685", armoperationalinsights.Workspace{
		Location: to.Ptr("australiasoutheast"),
		Tags: map[string]*string{
			"tag1": to.Ptr("val1"),
		},
		Properties: &armoperationalinsights.WorkspaceProperties{
			RetentionInDays: to.Ptr[int32](30),
			SKU: &armoperationalinsights.WorkspaceSKU{
				Name: to.Ptr(armoperationalinsights.WorkspaceSKUNameEnumPerGB2018),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Workspace = armoperationalinsights.Workspace{
	// 	Name: to.Ptr("AzTest2170"),
	// 	Type: to.Ptr("Microsoft.OperationalInsights/workspaces"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000005/resourcegroups/oiautorest6685/providers/microsoft.operationalinsights/workspaces/aztest2170"),
	// 	Location: to.Ptr("australiasoutheast"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("val1"),
	// 	},
	// 	Properties: &armoperationalinsights.WorkspaceProperties{
	// 		CustomerID: to.Ptr("bc089d7b-485c-4aff-a71e-c00f362d8d2f"),
	// 		ProvisioningState: to.Ptr(armoperationalinsights.WorkspaceEntityStatusSucceeded),
	// 		PublicNetworkAccessForIngestion: to.Ptr(armoperationalinsights.PublicNetworkAccessTypeEnabled),
	// 		PublicNetworkAccessForQuery: to.Ptr(armoperationalinsights.PublicNetworkAccessTypeEnabled),
	// 		RetentionInDays: to.Ptr[int32](30),
	// 		SKU: &armoperationalinsights.WorkspaceSKU{
	// 			Name: to.Ptr(armoperationalinsights.WorkspaceSKUNameEnumPerGB2018),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf3574813e15bb33b3cb610f44edfcbebd8b1b23/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/preview/2021-12-01-preview/examples/WorkspacesDelete.json
func ExampleWorkspacesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWorkspacesClient().BeginDelete(ctx, "oiautorest6685", "oiautorest6685", &armoperationalinsights.WorkspacesClientBeginDeleteOptions{Force: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf3574813e15bb33b3cb610f44edfcbebd8b1b23/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/preview/2021-12-01-preview/examples/WorkspacesGet.json
func ExampleWorkspacesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().Get(ctx, "oiautorest6685", "oiautorest6685", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Workspace = armoperationalinsights.Workspace{
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf3574813e15bb33b3cb610f44edfcbebd8b1b23/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/preview/2021-12-01-preview/examples/WorkspacesUpdate.json
func ExampleWorkspacesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().Update(ctx, "oiautorest6685", "oiautorest6685", armoperationalinsights.WorkspacePatch{
		Properties: &armoperationalinsights.WorkspaceProperties{
			RetentionInDays: to.Ptr[int32](30),
			SKU: &armoperationalinsights.WorkspaceSKU{
				Name: to.Ptr(armoperationalinsights.WorkspaceSKUNameEnumPerGB2018),
			},
			WorkspaceCapping: &armoperationalinsights.WorkspaceCapping{
				DailyQuotaGb: to.Ptr[float64](-1),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Workspace = armoperationalinsights.Workspace{
	// 	Location: to.Ptr("australiasoutheast"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("val1"),
	// 	},
	// 	Properties: &armoperationalinsights.WorkspaceProperties{
	// 		PublicNetworkAccessForIngestion: to.Ptr(armoperationalinsights.PublicNetworkAccessTypeEnabled),
	// 		PublicNetworkAccessForQuery: to.Ptr(armoperationalinsights.PublicNetworkAccessTypeEnabled),
	// 		RetentionInDays: to.Ptr[int32](30),
	// 		SKU: &armoperationalinsights.WorkspaceSKU{
	// 			Name: to.Ptr(armoperationalinsights.WorkspaceSKUNameEnumPerGB2018),
	// 		},
	// 		WorkspaceCapping: &armoperationalinsights.WorkspaceCapping{
	// 			DailyQuotaGb: to.Ptr[float64](-1),
	// 			DataIngestionStatus: to.Ptr(armoperationalinsights.DataIngestionStatusRespectQuota),
	// 			QuotaNextResetTime: to.Ptr("Mon, 16 Nov 2020 15:00:00 GMT"),
	// 		},
	// 	},
	// }
}
