//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/BillingAccountForecast.json
func ExampleForecastClient_Usage_billingAccountForecast() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewForecastClient().Usage(ctx, "providers/Microsoft.Billing/billingAccounts/12345:6789", armcostmanagement.ForecastDefinition{
		Type: to.Ptr(armcostmanagement.ForecastTypeUsage),
		Dataset: &armcostmanagement.ForecastDataset{
			Filter: &armcostmanagement.QueryFilter{
				And: []*armcostmanagement.QueryFilter{
					{
						Or: []*armcostmanagement.QueryFilter{
							{
								Dimensions: &armcostmanagement.QueryComparisonExpression{
									Name:     to.Ptr("ResourceLocation"),
									Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
									Values: []*string{
										to.Ptr("East US"),
										to.Ptr("West Europe")},
								},
							},
							{
								Tags: &armcostmanagement.QueryComparisonExpression{
									Name:     to.Ptr("Environment"),
									Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
									Values: []*string{
										to.Ptr("UAT"),
										to.Ptr("Prod")},
								},
							}},
					},
					{
						Dimensions: &armcostmanagement.QueryComparisonExpression{
							Name:     to.Ptr("ResourceGroup"),
							Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
							Values: []*string{
								to.Ptr("API")},
						},
					}},
			},
			Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
		},
		IncludeActualCost:       to.Ptr(false),
		IncludeFreshPartialCost: to.Ptr(false),
		Timeframe:               to.Ptr(armcostmanagement.ForecastTimeframeTypeMonthToDate),
	}, &armcostmanagement.ForecastClientUsageOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QueryResult = armcostmanagement.QueryResult{
	// 	Name: to.Ptr("ad67fd91-c131-4bda-9ba9-7187ecb1cebd"),
	// 	Type: to.Ptr("microsoft.costmanagement/Query"),
	// 	ID: to.Ptr("providers/Microsoft.Billing/billingAccounts/12345:6789/providers/Microsoft.CostManagement/query/ad67fd91-c131-4bda-9ba9-7187ecb1cebd"),
	// 	Properties: &armcostmanagement.QueryProperties{
	// 		Columns: []*armcostmanagement.QueryColumn{
	// 			{
	// 				Name: to.Ptr("PreTaxCost"),
	// 				Type: to.Ptr("Number"),
	// 			},
	// 			{
	// 				Name: to.Ptr("UsageDate"),
	// 				Type: to.Ptr("Number"),
	// 			},
	// 			{
	// 				Name: to.Ptr("CostStatus"),
	// 				Type: to.Ptr("String"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Currency"),
	// 				Type: to.Ptr("String"),
	// 		}},
	// 		Rows: [][]any{
	// 			[]any{
	// 				float64(2.10333307059661),
	// 				float64(20180331),
	// 				"Forecast",
	// 				"USD"},
	// 				[]any{
	// 					float64(218.68795741935486),
	// 					float64(20180331),
	// 					"Forecast",
	// 					"USD"},
	// 					[]any{
	// 						float64(0.14384913581657052),
	// 						float64(20180401),
	// 						"Forecast",
	// 						"USD"},
	// 						[]any{
	// 							float64(0.009865586851323632),
	// 							float64(20180429),
	// 							"Forecast",
	// 							"USD"}},
	// 						},
	// 					}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/BillingProfileForecast.json
func ExampleForecastClient_Usage_billingProfileForecast() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewForecastClient().Usage(ctx, "providers/Microsoft.Billing/billingAccounts/12345:6789/billingProfiles/13579", armcostmanagement.ForecastDefinition{
		Type: to.Ptr(armcostmanagement.ForecastTypeUsage),
		Dataset: &armcostmanagement.ForecastDataset{
			Filter: &armcostmanagement.QueryFilter{
				And: []*armcostmanagement.QueryFilter{
					{
						Or: []*armcostmanagement.QueryFilter{
							{
								Dimensions: &armcostmanagement.QueryComparisonExpression{
									Name:     to.Ptr("ResourceLocation"),
									Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
									Values: []*string{
										to.Ptr("East US"),
										to.Ptr("West Europe")},
								},
							},
							{
								Tags: &armcostmanagement.QueryComparisonExpression{
									Name:     to.Ptr("Environment"),
									Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
									Values: []*string{
										to.Ptr("UAT"),
										to.Ptr("Prod")},
								},
							}},
					},
					{
						Dimensions: &armcostmanagement.QueryComparisonExpression{
							Name:     to.Ptr("ResourceGroup"),
							Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
							Values: []*string{
								to.Ptr("API")},
						},
					}},
			},
			Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
		},
		IncludeActualCost:       to.Ptr(false),
		IncludeFreshPartialCost: to.Ptr(false),
		Timeframe:               to.Ptr(armcostmanagement.ForecastTimeframeTypeMonthToDate),
	}, &armcostmanagement.ForecastClientUsageOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QueryResult = armcostmanagement.QueryResult{
	// 	Name: to.Ptr("ad67fd91-c131-4bda-9ba9-7187ecb1cebd"),
	// 	Type: to.Ptr("microsoft.costmanagement/Query"),
	// 	ID: to.Ptr("providers/Microsoft.Billing/billingAccounts/12345:6789/billingProfiles/13579/providers/Microsoft.CostManagement/query/ad67fd91-c131-4bda-9ba9-7187ecb1cebd"),
	// 	Properties: &armcostmanagement.QueryProperties{
	// 		Columns: []*armcostmanagement.QueryColumn{
	// 			{
	// 				Name: to.Ptr("PreTaxCost"),
	// 				Type: to.Ptr("Number"),
	// 			},
	// 			{
	// 				Name: to.Ptr("UsageDate"),
	// 				Type: to.Ptr("Number"),
	// 			},
	// 			{
	// 				Name: to.Ptr("CostStatus"),
	// 				Type: to.Ptr("String"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Currency"),
	// 				Type: to.Ptr("String"),
	// 		}},
	// 		Rows: [][]any{
	// 			[]any{
	// 				float64(2.10333307059661),
	// 				float64(20180331),
	// 				"Forecast",
	// 				"USD"},
	// 				[]any{
	// 					float64(218.68795741935486),
	// 					float64(20180331),
	// 					"Forecast",
	// 					"USD"},
	// 					[]any{
	// 						float64(0.14384913581657052),
	// 						float64(20180401),
	// 						"Forecast",
	// 						"USD"},
	// 						[]any{
	// 							float64(0.009865586851323632),
	// 							float64(20180429),
	// 							"Forecast",
	// 							"USD"}},
	// 						},
	// 					}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/DepartmentForecast.json
func ExampleForecastClient_Usage_departmentForecast() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewForecastClient().Usage(ctx, "providers/Microsoft.Billing/billingAccounts/12345:6789/departments/123", armcostmanagement.ForecastDefinition{
		Type: to.Ptr(armcostmanagement.ForecastTypeUsage),
		Dataset: &armcostmanagement.ForecastDataset{
			Filter: &armcostmanagement.QueryFilter{
				And: []*armcostmanagement.QueryFilter{
					{
						Or: []*armcostmanagement.QueryFilter{
							{
								Dimensions: &armcostmanagement.QueryComparisonExpression{
									Name:     to.Ptr("ResourceLocation"),
									Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
									Values: []*string{
										to.Ptr("East US"),
										to.Ptr("West Europe")},
								},
							},
							{
								Tags: &armcostmanagement.QueryComparisonExpression{
									Name:     to.Ptr("Environment"),
									Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
									Values: []*string{
										to.Ptr("UAT"),
										to.Ptr("Prod")},
								},
							}},
					},
					{
						Dimensions: &armcostmanagement.QueryComparisonExpression{
							Name:     to.Ptr("ResourceGroup"),
							Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
							Values: []*string{
								to.Ptr("API")},
						},
					}},
			},
			Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
		},
		IncludeActualCost:       to.Ptr(false),
		IncludeFreshPartialCost: to.Ptr(false),
		Timeframe:               to.Ptr(armcostmanagement.ForecastTimeframeTypeMonthToDate),
	}, &armcostmanagement.ForecastClientUsageOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QueryResult = armcostmanagement.QueryResult{
	// 	Name: to.Ptr("ad67fd91-c131-4bda-9ba9-7187ecb1cebd"),
	// 	Type: to.Ptr("microsoft.costmanagement/Query"),
	// 	ID: to.Ptr("providers/Microsoft.Billing/billingAccounts/12345:6789/departments/123/providers/Microsoft.CostManagement/query/ad67fd91-c131-4bda-9ba9-7187ecb1cebd"),
	// 	Properties: &armcostmanagement.QueryProperties{
	// 		Columns: []*armcostmanagement.QueryColumn{
	// 			{
	// 				Name: to.Ptr("PreTaxCost"),
	// 				Type: to.Ptr("Number"),
	// 			},
	// 			{
	// 				Name: to.Ptr("UsageDate"),
	// 				Type: to.Ptr("Number"),
	// 			},
	// 			{
	// 				Name: to.Ptr("CostStatus"),
	// 				Type: to.Ptr("String"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Currency"),
	// 				Type: to.Ptr("String"),
	// 		}},
	// 		Rows: [][]any{
	// 			[]any{
	// 				float64(2.10333307059661),
	// 				float64(20180331),
	// 				"Forecast",
	// 				"USD"},
	// 				[]any{
	// 					float64(218.68795741935486),
	// 					float64(20180331),
	// 					"Forecast",
	// 					"USD"},
	// 					[]any{
	// 						float64(0.14384913581657052),
	// 						float64(20180401),
	// 						"Forecast",
	// 						"USD"},
	// 						[]any{
	// 							float64(0.009865586851323632),
	// 							float64(20180429),
	// 							"Forecast",
	// 							"USD"}},
	// 						},
	// 					}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/EnrollmentAccountForecast.json
func ExampleForecastClient_Usage_enrollmentAccountForecast() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewForecastClient().Usage(ctx, "providers/Microsoft.Billing/billingAccounts/12345:6789/enrollmentAccounts/456", armcostmanagement.ForecastDefinition{
		Type: to.Ptr(armcostmanagement.ForecastTypeUsage),
		Dataset: &armcostmanagement.ForecastDataset{
			Filter: &armcostmanagement.QueryFilter{
				And: []*armcostmanagement.QueryFilter{
					{
						Or: []*armcostmanagement.QueryFilter{
							{
								Dimensions: &armcostmanagement.QueryComparisonExpression{
									Name:     to.Ptr("ResourceLocation"),
									Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
									Values: []*string{
										to.Ptr("East US"),
										to.Ptr("West Europe")},
								},
							},
							{
								Tags: &armcostmanagement.QueryComparisonExpression{
									Name:     to.Ptr("Environment"),
									Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
									Values: []*string{
										to.Ptr("UAT"),
										to.Ptr("Prod")},
								},
							}},
					},
					{
						Dimensions: &armcostmanagement.QueryComparisonExpression{
							Name:     to.Ptr("ResourceGroup"),
							Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
							Values: []*string{
								to.Ptr("API")},
						},
					}},
			},
			Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
		},
		IncludeActualCost:       to.Ptr(false),
		IncludeFreshPartialCost: to.Ptr(false),
		Timeframe:               to.Ptr(armcostmanagement.ForecastTimeframeTypeMonthToDate),
	}, &armcostmanagement.ForecastClientUsageOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QueryResult = armcostmanagement.QueryResult{
	// 	Name: to.Ptr("ad67fd91-c131-4bda-9ba9-7187ecb1cebd"),
	// 	Type: to.Ptr("microsoft.costmanagement/Query"),
	// 	ID: to.Ptr("providers/Microsoft.Billing/billingAccounts/12345:6789/enrollmentAccounts/456/providers/Microsoft.CostManagement/query/ad67fd91-c131-4bda-9ba9-7187ecb1cebd"),
	// 	Properties: &armcostmanagement.QueryProperties{
	// 		Columns: []*armcostmanagement.QueryColumn{
	// 			{
	// 				Name: to.Ptr("PreTaxCost"),
	// 				Type: to.Ptr("Number"),
	// 			},
	// 			{
	// 				Name: to.Ptr("UsageDate"),
	// 				Type: to.Ptr("Number"),
	// 			},
	// 			{
	// 				Name: to.Ptr("CostStatus"),
	// 				Type: to.Ptr("String"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Currency"),
	// 				Type: to.Ptr("String"),
	// 		}},
	// 		Rows: [][]any{
	// 			[]any{
	// 				float64(2.10333307059661),
	// 				float64(20180331),
	// 				"Forecast",
	// 				"USD"},
	// 				[]any{
	// 					float64(218.68795741935486),
	// 					float64(20180331),
	// 					"Forecast",
	// 					"USD"},
	// 					[]any{
	// 						float64(0.14384913581657052),
	// 						float64(20180401),
	// 						"Forecast",
	// 						"USD"},
	// 						[]any{
	// 							float64(0.009865586851323632),
	// 							float64(20180429),
	// 							"Forecast",
	// 							"USD"}},
	// 						},
	// 					}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/InvoiceSectionForecast.json
func ExampleForecastClient_Usage_invoiceSectionForecast() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewForecastClient().Usage(ctx, "providers/Microsoft.Billing/billingAccounts/12345:6789/billingProfiles/13579/invoiceSections/9876", armcostmanagement.ForecastDefinition{
		Type: to.Ptr(armcostmanagement.ForecastTypeUsage),
		Dataset: &armcostmanagement.ForecastDataset{
			Filter: &armcostmanagement.QueryFilter{
				And: []*armcostmanagement.QueryFilter{
					{
						Or: []*armcostmanagement.QueryFilter{
							{
								Dimensions: &armcostmanagement.QueryComparisonExpression{
									Name:     to.Ptr("ResourceLocation"),
									Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
									Values: []*string{
										to.Ptr("East US"),
										to.Ptr("West Europe")},
								},
							},
							{
								Tags: &armcostmanagement.QueryComparisonExpression{
									Name:     to.Ptr("Environment"),
									Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
									Values: []*string{
										to.Ptr("UAT"),
										to.Ptr("Prod")},
								},
							}},
					},
					{
						Dimensions: &armcostmanagement.QueryComparisonExpression{
							Name:     to.Ptr("ResourceGroup"),
							Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
							Values: []*string{
								to.Ptr("API")},
						},
					}},
			},
			Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
		},
		IncludeActualCost:       to.Ptr(false),
		IncludeFreshPartialCost: to.Ptr(false),
		Timeframe:               to.Ptr(armcostmanagement.ForecastTimeframeTypeMonthToDate),
	}, &armcostmanagement.ForecastClientUsageOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QueryResult = armcostmanagement.QueryResult{
	// 	Name: to.Ptr("ad67fd91-c131-4bda-9ba9-7187ecb1cebd"),
	// 	Type: to.Ptr("microsoft.costmanagement/Query"),
	// 	ID: to.Ptr("providers/Microsoft.Billing/billingAccounts/12345:6789/billingProfiles/13579/invoiceSections/9876/providers/Microsoft.CostManagement/query/ad67fd91-c131-4bda-9ba9-7187ecb1cebd"),
	// 	Properties: &armcostmanagement.QueryProperties{
	// 		Columns: []*armcostmanagement.QueryColumn{
	// 			{
	// 				Name: to.Ptr("PreTaxCost"),
	// 				Type: to.Ptr("Number"),
	// 			},
	// 			{
	// 				Name: to.Ptr("UsageDate"),
	// 				Type: to.Ptr("Number"),
	// 			},
	// 			{
	// 				Name: to.Ptr("CostStatus"),
	// 				Type: to.Ptr("String"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Currency"),
	// 				Type: to.Ptr("String"),
	// 		}},
	// 		Rows: [][]any{
	// 			[]any{
	// 				float64(2.10333307059661),
	// 				float64(20180331),
	// 				"Forecast",
	// 				"USD"},
	// 				[]any{
	// 					float64(218.68795741935486),
	// 					float64(20180331),
	// 					"Forecast",
	// 					"USD"},
	// 					[]any{
	// 						float64(0.14384913581657052),
	// 						float64(20180401),
	// 						"Forecast",
	// 						"USD"},
	// 						[]any{
	// 							float64(0.009865586851323632),
	// 							float64(20180429),
	// 							"Forecast",
	// 							"USD"}},
	// 						},
	// 					}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/ResourceGroupForecast.json
func ExampleForecastClient_Usage_resourceGroupForecast() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewForecastClient().Usage(ctx, "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/ScreenSharingTest-peer", armcostmanagement.ForecastDefinition{
		Type: to.Ptr(armcostmanagement.ForecastTypeUsage),
		Dataset: &armcostmanagement.ForecastDataset{
			Filter: &armcostmanagement.QueryFilter{
				And: []*armcostmanagement.QueryFilter{
					{
						Or: []*armcostmanagement.QueryFilter{
							{
								Dimensions: &armcostmanagement.QueryComparisonExpression{
									Name:     to.Ptr("ResourceLocation"),
									Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
									Values: []*string{
										to.Ptr("East US"),
										to.Ptr("West Europe")},
								},
							},
							{
								Tags: &armcostmanagement.QueryComparisonExpression{
									Name:     to.Ptr("Environment"),
									Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
									Values: []*string{
										to.Ptr("UAT"),
										to.Ptr("Prod")},
								},
							}},
					},
					{
						Dimensions: &armcostmanagement.QueryComparisonExpression{
							Name:     to.Ptr("ResourceGroup"),
							Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
							Values: []*string{
								to.Ptr("API")},
						},
					}},
			},
			Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
		},
		IncludeActualCost:       to.Ptr(false),
		IncludeFreshPartialCost: to.Ptr(false),
		Timeframe:               to.Ptr(armcostmanagement.ForecastTimeframeTypeMonthToDate),
	}, &armcostmanagement.ForecastClientUsageOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QueryResult = armcostmanagement.QueryResult{
	// 	Name: to.Ptr("55312978-ba1b-415c-9304-cfd9c43c0481"),
	// 	Type: to.Ptr("microsoft.costmanagement/Query"),
	// 	ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/ScreenSharingTest-peer/providers/Microsoft.CostManagement/query/00000000-0000-0000-0000-000000000000"),
	// 	Properties: &armcostmanagement.QueryProperties{
	// 		Columns: []*armcostmanagement.QueryColumn{
	// 			{
	// 				Name: to.Ptr("PreTaxCost"),
	// 				Type: to.Ptr("Number"),
	// 			},
	// 			{
	// 				Name: to.Ptr("UsageDate"),
	// 				Type: to.Ptr("Number"),
	// 			},
	// 			{
	// 				Name: to.Ptr("CostStatus"),
	// 				Type: to.Ptr("String"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Currency"),
	// 				Type: to.Ptr("String"),
	// 		}},
	// 		Rows: [][]any{
	// 			[]any{
	// 				float64(2.10333307059661),
	// 				float64(20180331),
	// 				"Forecast",
	// 				"USD"}},
	// 			},
	// 		}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/SubscriptionForecast.json
func ExampleForecastClient_Usage_subscriptionForecast() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewForecastClient().Usage(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", armcostmanagement.ForecastDefinition{
		Type: to.Ptr(armcostmanagement.ForecastTypeUsage),
		Dataset: &armcostmanagement.ForecastDataset{
			Filter: &armcostmanagement.QueryFilter{
				And: []*armcostmanagement.QueryFilter{
					{
						Or: []*armcostmanagement.QueryFilter{
							{
								Dimensions: &armcostmanagement.QueryComparisonExpression{
									Name:     to.Ptr("ResourceLocation"),
									Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
									Values: []*string{
										to.Ptr("East US"),
										to.Ptr("West Europe")},
								},
							},
							{
								Tags: &armcostmanagement.QueryComparisonExpression{
									Name:     to.Ptr("Environment"),
									Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
									Values: []*string{
										to.Ptr("UAT"),
										to.Ptr("Prod")},
								},
							}},
					},
					{
						Dimensions: &armcostmanagement.QueryComparisonExpression{
							Name:     to.Ptr("ResourceGroup"),
							Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
							Values: []*string{
								to.Ptr("API")},
						},
					}},
			},
			Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
		},
		IncludeActualCost:       to.Ptr(false),
		IncludeFreshPartialCost: to.Ptr(false),
		Timeframe:               to.Ptr(armcostmanagement.ForecastTimeframeTypeMonthToDate),
	}, &armcostmanagement.ForecastClientUsageOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QueryResult = armcostmanagement.QueryResult{
	// 	Name: to.Ptr("55312978-ba1b-415c-9304-cfd9c43c0481"),
	// 	Type: to.Ptr("microsoft.costmanagement/Query"),
	// 	ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.CostManagement/query/00000000-0000-0000-0000-000000000000"),
	// 	Properties: &armcostmanagement.QueryProperties{
	// 		Columns: []*armcostmanagement.QueryColumn{
	// 			{
	// 				Name: to.Ptr("PreTaxCost"),
	// 				Type: to.Ptr("Number"),
	// 			},
	// 			{
	// 				Name: to.Ptr("UsageDate"),
	// 				Type: to.Ptr("Number"),
	// 			},
	// 			{
	// 				Name: to.Ptr("CostStatus"),
	// 				Type: to.Ptr("String"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Currency"),
	// 				Type: to.Ptr("String"),
	// 		}},
	// 		Rows: [][]any{
	// 			[]any{
	// 				float64(2.10333307059661),
	// 				float64(20180331),
	// 				"Forecast",
	// 				"USD"},
	// 				[]any{
	// 					float64(218.68795741935486),
	// 					float64(20180331),
	// 					"Forecast",
	// 					"USD"},
	// 					[]any{
	// 						float64(0.14384913581657052),
	// 						float64(20180401),
	// 						"Forecast",
	// 						"USD"},
	// 						[]any{
	// 							float64(0.009865586851323632),
	// 							float64(20180429),
	// 							"Forecast",
	// 							"USD"}},
	// 						},
	// 					}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/ExternalBillingAccountForecast.json
func ExampleForecastClient_ExternalCloudProviderUsage_externalBillingAccountForecast() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewForecastClient().ExternalCloudProviderUsage(ctx, armcostmanagement.ExternalCloudProviderTypeExternalBillingAccounts, "100", armcostmanagement.ForecastDefinition{
		Type: to.Ptr(armcostmanagement.ForecastTypeUsage),
		Dataset: &armcostmanagement.ForecastDataset{
			Filter: &armcostmanagement.QueryFilter{
				And: []*armcostmanagement.QueryFilter{
					{
						Or: []*armcostmanagement.QueryFilter{
							{
								Dimensions: &armcostmanagement.QueryComparisonExpression{
									Name:     to.Ptr("ResourceLocation"),
									Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
									Values: []*string{
										to.Ptr("East US"),
										to.Ptr("West Europe")},
								},
							},
							{
								Tags: &armcostmanagement.QueryComparisonExpression{
									Name:     to.Ptr("Environment"),
									Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
									Values: []*string{
										to.Ptr("UAT"),
										to.Ptr("Prod")},
								},
							}},
					},
					{
						Dimensions: &armcostmanagement.QueryComparisonExpression{
							Name:     to.Ptr("ResourceGroup"),
							Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
							Values: []*string{
								to.Ptr("API")},
						},
					}},
			},
			Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
		},
		Timeframe: to.Ptr(armcostmanagement.ForecastTimeframeTypeMonthToDate),
	}, &armcostmanagement.ForecastClientExternalCloudProviderUsageOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QueryResult = armcostmanagement.QueryResult{
	// 	Name: to.Ptr("6dc7b06a-d90a-4df5-b655-ce6cf1c0814d"),
	// 	Type: to.Ptr("Microsoft.CostManagement/query"),
	// 	ID: to.Ptr("providers/Microsoft.CostManagement/externalBillingAccounts/100/query/6dc7b06a-d90a-4df5-b655-ce6cf1c0814d"),
	// 	Properties: &armcostmanagement.QueryProperties{
	// 		Columns: []*armcostmanagement.QueryColumn{
	// 			{
	// 				Name: to.Ptr("PreTaxCost"),
	// 				Type: to.Ptr("Number"),
	// 			},
	// 			{
	// 				Name: to.Ptr("CostStatus"),
	// 				Type: to.Ptr("String"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Currency"),
	// 				Type: to.Ptr("String"),
	// 		}},
	// 		Rows: [][]any{
	// 			[]any{
	// 				float64(0),
	// 				"Forecast",
	// 				"USD"},
	// 				[]any{
	// 					float64(30.2572751438),
	// 					"Forecast",
	// 					"USD"},
	// 					[]any{
	// 						float64(0.07675760200000002),
	// 						"Forecast",
	// 						"USD"},
	// 						[]any{
	// 							float64(50.43096419040001),
	// 							"Forecast",
	// 							"USD"}},
	// 						},
	// 					}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/ExternalSubscriptionForecast.json
func ExampleForecastClient_ExternalCloudProviderUsage_externalSubscriptionForecast() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewForecastClient().ExternalCloudProviderUsage(ctx, armcostmanagement.ExternalCloudProviderTypeExternalSubscriptions, "100", armcostmanagement.ForecastDefinition{
		Type: to.Ptr(armcostmanagement.ForecastTypeUsage),
		Dataset: &armcostmanagement.ForecastDataset{
			Filter: &armcostmanagement.QueryFilter{
				And: []*armcostmanagement.QueryFilter{
					{
						Or: []*armcostmanagement.QueryFilter{
							{
								Dimensions: &armcostmanagement.QueryComparisonExpression{
									Name:     to.Ptr("ResourceLocation"),
									Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
									Values: []*string{
										to.Ptr("East US"),
										to.Ptr("West Europe")},
								},
							},
							{
								Tags: &armcostmanagement.QueryComparisonExpression{
									Name:     to.Ptr("Environment"),
									Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
									Values: []*string{
										to.Ptr("UAT"),
										to.Ptr("Prod")},
								},
							}},
					},
					{
						Dimensions: &armcostmanagement.QueryComparisonExpression{
							Name:     to.Ptr("ResourceGroup"),
							Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
							Values: []*string{
								to.Ptr("API")},
						},
					}},
			},
			Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
		},
		Timeframe: to.Ptr(armcostmanagement.ForecastTimeframeTypeMonthToDate),
	}, &armcostmanagement.ForecastClientExternalCloudProviderUsageOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QueryResult = armcostmanagement.QueryResult{
	// 	Name: to.Ptr("d99477af-7510-40ee-aca2-e59bdca0d10d"),
	// 	Type: to.Ptr("Microsoft.CostManagement/query"),
	// 	ID: to.Ptr("providers/Microsoft.CostManagement/externalSubscriptions/100/query/d99477af-7510-40ee-aca2-e59bdca0d10d"),
	// 	Properties: &armcostmanagement.QueryProperties{
	// 		Columns: []*armcostmanagement.QueryColumn{
	// 			{
	// 				Name: to.Ptr("PreTaxCost"),
	// 				Type: to.Ptr("Number"),
	// 			},
	// 			{
	// 				Name: to.Ptr("UsageDate"),
	// 				Type: to.Ptr("Number"),
	// 			},
	// 			{
	// 				Name: to.Ptr("CostStatus"),
	// 				Type: to.Ptr("String"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Currency"),
	// 				Type: to.Ptr("String"),
	// 		}},
	// 		Rows: [][]any{
	// 		},
	// 	},
	// }
}
