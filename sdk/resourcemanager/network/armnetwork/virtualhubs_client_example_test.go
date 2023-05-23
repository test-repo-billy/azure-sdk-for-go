//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/VirtualHubGet.json
func ExampleVirtualHubsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualHubsClient().Get(ctx, "rg1", "virtualHub1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualHub = armnetwork.VirtualHub{
	// 	Name: to.Ptr("virtualHub1"),
	// 	Type: to.Ptr("Microsoft.Network/virtualHubs"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1"),
	// 	Location: to.Ptr("West US"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.VirtualHubProperties{
	// 		AddressPrefix: to.Ptr("10.10.1.0/24"),
	// 		AllowBranchToBranchTraffic: to.Ptr(false),
	// 		HubRoutingPreference: to.Ptr(armnetwork.HubRoutingPreferenceExpressRoute),
	// 		PreferredRoutingGateway: to.Ptr(armnetwork.PreferredRoutingGatewayExpressRoute),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RoutingState: to.Ptr(armnetwork.RoutingStateProvisioned),
	// 		SKU: to.Ptr("Basic"),
	// 		VirtualHubRouteTableV2S: []*armnetwork.VirtualHubRouteTableV2{
	// 		},
	// 		VirtualRouterAsn: to.Ptr[int64](65515),
	// 		VirtualRouterAutoScaleConfiguration: &armnetwork.VirtualRouterAutoScaleConfiguration{
	// 			MinCapacity: to.Ptr[int32](2),
	// 		},
	// 		VirtualRouterIPs: []*string{
	// 			to.Ptr("10.10.1.12"),
	// 			to.Ptr("10.10.1.13")},
	// 			VirtualWan: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualWans/virtualWan1"),
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/VirtualHubPut.json
func ExampleVirtualHubsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualHubsClient().BeginCreateOrUpdate(ctx, "rg1", "virtualHub2", armnetwork.VirtualHub{
		Location: to.Ptr("West US"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
		},
		Properties: &armnetwork.VirtualHubProperties{
			AddressPrefix: to.Ptr("10.168.0.0/24"),
			SKU:           to.Ptr("Basic"),
			VirtualWan: &armnetwork.SubResource{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualWans/virtualWan1"),
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
	// res.VirtualHub = armnetwork.VirtualHub{
	// 	Name: to.Ptr("virtualHub2"),
	// 	Type: to.Ptr("Microsoft.Network/virtualHubs"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub2"),
	// 	Location: to.Ptr("West US"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.VirtualHubProperties{
	// 		AddressPrefix: to.Ptr("10.168.0.0/24"),
	// 		AllowBranchToBranchTraffic: to.Ptr(false),
	// 		HubRoutingPreference: to.Ptr(armnetwork.HubRoutingPreferenceExpressRoute),
	// 		PreferredRoutingGateway: to.Ptr(armnetwork.PreferredRoutingGatewayExpressRoute),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		RoutingState: to.Ptr(armnetwork.RoutingStateProvisioned),
	// 		SKU: to.Ptr("Basic"),
	// 		VirtualHubRouteTableV2S: []*armnetwork.VirtualHubRouteTableV2{
	// 		},
	// 		VirtualRouterAsn: to.Ptr[int64](65515),
	// 		VirtualRouterAutoScaleConfiguration: &armnetwork.VirtualRouterAutoScaleConfiguration{
	// 			MinCapacity: to.Ptr[int32](2),
	// 		},
	// 		VirtualRouterIPs: []*string{
	// 			to.Ptr("10.10.1.12"),
	// 			to.Ptr("10.10.1.13")},
	// 			VirtualWan: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualWans/virtualWan1"),
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/VirtualHubUpdateTags.json
func ExampleVirtualHubsClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualHubsClient().UpdateTags(ctx, "rg1", "virtualHub2", armnetwork.TagsObject{
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
			"key2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualHub = armnetwork.VirtualHub{
	// 	Name: to.Ptr("virtualHub2"),
	// 	Type: to.Ptr("Microsoft.Network/virtualHubs"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub2"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key2": to.Ptr("value2"),
	// 	},
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.VirtualHubProperties{
	// 		AddressPrefix: to.Ptr("10.168.0.0/24"),
	// 		AllowBranchToBranchTraffic: to.Ptr(false),
	// 		HubRoutingPreference: to.Ptr(armnetwork.HubRoutingPreferenceExpressRoute),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		SKU: to.Ptr("Basic"),
	// 		VirtualHubRouteTableV2S: []*armnetwork.VirtualHubRouteTableV2{
	// 		},
	// 		VirtualRouterAsn: to.Ptr[int64](65515),
	// 		VirtualRouterAutoScaleConfiguration: &armnetwork.VirtualRouterAutoScaleConfiguration{
	// 			MinCapacity: to.Ptr[int32](2),
	// 		},
	// 		VirtualRouterIPs: []*string{
	// 			to.Ptr("10.10.1.12"),
	// 			to.Ptr("10.10.1.13")},
	// 			VirtualWan: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualWans/virtualWan1"),
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/VirtualHubDelete.json
func ExampleVirtualHubsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualHubsClient().BeginDelete(ctx, "rg1", "virtualHub1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/VirtualHubListByResourceGroup.json
func ExampleVirtualHubsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualHubsClient().NewListByResourceGroupPager("rg1", nil)
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
		// page.ListVirtualHubsResult = armnetwork.ListVirtualHubsResult{
		// 	Value: []*armnetwork.VirtualHub{
		// 		{
		// 			Name: to.Ptr("virtualHub1"),
		// 			Type: to.Ptr("Microsoft.Network/virtualHubs"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1"),
		// 			Location: to.Ptr("West US"),
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.VirtualHubProperties{
		// 				AddressPrefix: to.Ptr("10.10.1.0/24"),
		// 				AllowBranchToBranchTraffic: to.Ptr(false),
		// 				HubRoutingPreference: to.Ptr(armnetwork.HubRoutingPreferenceExpressRoute),
		// 				PreferredRoutingGateway: to.Ptr(armnetwork.PreferredRoutingGatewayExpressRoute),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				RoutingState: to.Ptr(armnetwork.RoutingStateProvisioned),
		// 				SKU: to.Ptr("Basic"),
		// 				VirtualHubRouteTableV2S: []*armnetwork.VirtualHubRouteTableV2{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub2/routeTables/virtualHubRouteTable2"),
		// 						Name: to.Ptr("rt2a"),
		// 						Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 						Properties: &armnetwork.VirtualHubRouteTableV2Properties{
		// 							AttachedConnections: []*string{
		// 								to.Ptr("All_Vnets")},
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								Routes: []*armnetwork.VirtualHubRouteV2{
		// 									{
		// 										DestinationType: to.Ptr("CIDR"),
		// 										Destinations: []*string{
		// 											to.Ptr("20.10.0.0/16"),
		// 											to.Ptr("20.20.0.0/16")},
		// 											NextHopType: to.Ptr("IPAddress"),
		// 											NextHops: []*string{
		// 												to.Ptr("10.0.0.68")},
		// 											},
		// 											{
		// 												DestinationType: to.Ptr("CIDR"),
		// 												Destinations: []*string{
		// 													to.Ptr("0.0.0.0/0")},
		// 													NextHopType: to.Ptr("IPAddress"),
		// 													NextHops: []*string{
		// 														to.Ptr("10.0.0.68")},
		// 												}},
		// 											},
		// 									}},
		// 									VirtualRouterAsn: to.Ptr[int64](65515),
		// 									VirtualRouterAutoScaleConfiguration: &armnetwork.VirtualRouterAutoScaleConfiguration{
		// 										MinCapacity: to.Ptr[int32](2),
		// 									},
		// 									VirtualRouterIPs: []*string{
		// 										to.Ptr("10.10.1.12"),
		// 										to.Ptr("10.10.1.13")},
		// 										VirtualWan: &armnetwork.SubResource{
		// 											ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualWans/virtualWan1"),
		// 										},
		// 									},
		// 								},
		// 								{
		// 									Name: to.Ptr("virtualHub2"),
		// 									Type: to.Ptr("Microsoft.Network/virtualHubs"),
		// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub2"),
		// 									Location: to.Ptr("East US"),
		// 									Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 									Properties: &armnetwork.VirtualHubProperties{
		// 										AddressPrefix: to.Ptr("210.10.1.0/24"),
		// 										AllowBranchToBranchTraffic: to.Ptr(false),
		// 										HubRoutingPreference: to.Ptr(armnetwork.HubRoutingPreferenceExpressRoute),
		// 										PreferredRoutingGateway: to.Ptr(armnetwork.PreferredRoutingGatewayExpressRoute),
		// 										ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 										RoutingState: to.Ptr(armnetwork.RoutingStateProvisioned),
		// 										SKU: to.Ptr("Basic"),
		// 										VirtualHubRouteTableV2S: []*armnetwork.VirtualHubRouteTableV2{
		// 											{
		// 												ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub2/routeTables/virtualHubRouteTable2"),
		// 												Name: to.Ptr("rt2a"),
		// 												Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 												Properties: &armnetwork.VirtualHubRouteTableV2Properties{
		// 													AttachedConnections: []*string{
		// 														to.Ptr("All_Vnets")},
		// 														ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 														Routes: []*armnetwork.VirtualHubRouteV2{
		// 															{
		// 																DestinationType: to.Ptr("CIDR"),
		// 																Destinations: []*string{
		// 																	to.Ptr("20.10.0.0/16"),
		// 																	to.Ptr("20.20.0.0/16")},
		// 																	NextHopType: to.Ptr("IPAddress"),
		// 																	NextHops: []*string{
		// 																		to.Ptr("10.0.0.68")},
		// 																	},
		// 																	{
		// 																		DestinationType: to.Ptr("CIDR"),
		// 																		Destinations: []*string{
		// 																			to.Ptr("0.0.0.0/0")},
		// 																			NextHopType: to.Ptr("IPAddress"),
		// 																			NextHops: []*string{
		// 																				to.Ptr("10.0.0.68")},
		// 																		}},
		// 																	},
		// 															}},
		// 															VirtualRouterAsn: to.Ptr[int64](65515),
		// 															VirtualRouterAutoScaleConfiguration: &armnetwork.VirtualRouterAutoScaleConfiguration{
		// 																MinCapacity: to.Ptr[int32](2),
		// 															},
		// 															VirtualRouterIPs: []*string{
		// 																to.Ptr("10.10.1.12"),
		// 																to.Ptr("10.10.1.13")},
		// 																VirtualWan: &armnetwork.SubResource{
		// 																	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualWans/virtualWan1"),
		// 																},
		// 															},
		// 													}},
		// 												}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/VirtualHubList.json
func ExampleVirtualHubsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualHubsClient().NewListPager(nil)
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
		// page.ListVirtualHubsResult = armnetwork.ListVirtualHubsResult{
		// 	Value: []*armnetwork.VirtualHub{
		// 		{
		// 			Name: to.Ptr("virtualHub1"),
		// 			Type: to.Ptr("Microsoft.Network/virtualHubs"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1"),
		// 			Location: to.Ptr("West US"),
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.VirtualHubProperties{
		// 				AddressPrefix: to.Ptr("10.10.1.0/24"),
		// 				AllowBranchToBranchTraffic: to.Ptr(false),
		// 				HubRoutingPreference: to.Ptr(armnetwork.HubRoutingPreferenceExpressRoute),
		// 				PreferredRoutingGateway: to.Ptr(armnetwork.PreferredRoutingGatewayExpressRoute),
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				RoutingState: to.Ptr(armnetwork.RoutingStateProvisioned),
		// 				SKU: to.Ptr("Basic"),
		// 				VirtualHubRouteTableV2S: []*armnetwork.VirtualHubRouteTableV2{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeTables/virtualHubRouteTable1"),
		// 						Name: to.Ptr("rt2a"),
		// 						Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 						Properties: &armnetwork.VirtualHubRouteTableV2Properties{
		// 							AttachedConnections: []*string{
		// 								to.Ptr("All_Vnets")},
		// 								ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 								Routes: []*armnetwork.VirtualHubRouteV2{
		// 									{
		// 										DestinationType: to.Ptr("CIDR"),
		// 										Destinations: []*string{
		// 											to.Ptr("20.10.0.0/16"),
		// 											to.Ptr("20.20.0.0/16")},
		// 											NextHopType: to.Ptr("IPAddress"),
		// 											NextHops: []*string{
		// 												to.Ptr("10.0.0.68")},
		// 											},
		// 											{
		// 												DestinationType: to.Ptr("CIDR"),
		// 												Destinations: []*string{
		// 													to.Ptr("0.0.0.0/0")},
		// 													NextHopType: to.Ptr("IPAddress"),
		// 													NextHops: []*string{
		// 														to.Ptr("10.0.0.68")},
		// 												}},
		// 											},
		// 									}},
		// 									VirtualRouterAsn: to.Ptr[int64](65515),
		// 									VirtualRouterAutoScaleConfiguration: &armnetwork.VirtualRouterAutoScaleConfiguration{
		// 										MinCapacity: to.Ptr[int32](2),
		// 									},
		// 									VirtualRouterIPs: []*string{
		// 										to.Ptr("10.10.1.12"),
		// 										to.Ptr("10.10.1.13")},
		// 										VirtualWan: &armnetwork.SubResource{
		// 											ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualWans/virtualWan1"),
		// 										},
		// 									},
		// 								},
		// 								{
		// 									Name: to.Ptr("virtualHub2"),
		// 									Type: to.Ptr("Microsoft.Network/virtualHubs"),
		// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub2"),
		// 									Location: to.Ptr("East US"),
		// 									Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 									Properties: &armnetwork.VirtualHubProperties{
		// 										AddressPrefix: to.Ptr("210.10.1.0/24"),
		// 										AllowBranchToBranchTraffic: to.Ptr(false),
		// 										HubRoutingPreference: to.Ptr(armnetwork.HubRoutingPreferenceExpressRoute),
		// 										PreferredRoutingGateway: to.Ptr(armnetwork.PreferredRoutingGatewayExpressRoute),
		// 										ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 										RoutingState: to.Ptr(armnetwork.RoutingStateProvisioned),
		// 										SKU: to.Ptr("Basic"),
		// 										VirtualHubRouteTableV2S: []*armnetwork.VirtualHubRouteTableV2{
		// 											{
		// 												ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub2/routeTables/virtualHubRouteTable2"),
		// 												Name: to.Ptr("rt2a"),
		// 												Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 												Properties: &armnetwork.VirtualHubRouteTableV2Properties{
		// 													AttachedConnections: []*string{
		// 														to.Ptr("All_Vnets")},
		// 														ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 														Routes: []*armnetwork.VirtualHubRouteV2{
		// 															{
		// 																DestinationType: to.Ptr("CIDR"),
		// 																Destinations: []*string{
		// 																	to.Ptr("20.10.0.0/16"),
		// 																	to.Ptr("20.20.0.0/16")},
		// 																	NextHopType: to.Ptr("IPAddress"),
		// 																	NextHops: []*string{
		// 																		to.Ptr("10.0.0.68")},
		// 																	},
		// 																	{
		// 																		DestinationType: to.Ptr("CIDR"),
		// 																		Destinations: []*string{
		// 																			to.Ptr("0.0.0.0/0")},
		// 																			NextHopType: to.Ptr("IPAddress"),
		// 																			NextHops: []*string{
		// 																				to.Ptr("10.0.0.68")},
		// 																		}},
		// 																	},
		// 															}},
		// 															VirtualRouterAsn: to.Ptr[int64](65515),
		// 															VirtualRouterAutoScaleConfiguration: &armnetwork.VirtualRouterAutoScaleConfiguration{
		// 																MinCapacity: to.Ptr[int32](2),
		// 															},
		// 															VirtualRouterIPs: []*string{
		// 																to.Ptr("10.10.1.12"),
		// 																to.Ptr("10.10.1.13")},
		// 																VirtualWan: &armnetwork.SubResource{
		// 																	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualWans/virtualWan1"),
		// 																},
		// 															},
		// 													}},
		// 												}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/EffectiveRoutesListForConnection.json
func ExampleVirtualHubsClient_BeginGetEffectiveVirtualHubRoutes_effectiveRoutesForAConnectionResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualHubsClient().BeginGetEffectiveVirtualHubRoutes(ctx, "rg1", "virtualHub1", &armnetwork.VirtualHubsClientBeginGetEffectiveVirtualHubRoutesOptions{EffectiveRoutesParameters: &armnetwork.EffectiveRoutesParameters{
		ResourceID:             to.Ptr("/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteGateways/expressRouteGatewayName/expressRouteConnections/connectionName"),
		VirtualWanResourceType: to.Ptr("ExpressRouteConnection"),
	},
	})
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
	// res.VirtualHubEffectiveRouteList = armnetwork.VirtualHubEffectiveRouteList{
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/EffectiveRoutesListForRouteTable.json
func ExampleVirtualHubsClient_BeginGetEffectiveVirtualHubRoutes_effectiveRoutesForARouteTableResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualHubsClient().BeginGetEffectiveVirtualHubRoutes(ctx, "rg1", "virtualHub1", &armnetwork.VirtualHubsClientBeginGetEffectiveVirtualHubRoutesOptions{EffectiveRoutesParameters: &armnetwork.EffectiveRoutesParameters{
		ResourceID:             to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
		VirtualWanResourceType: to.Ptr("RouteTable"),
	},
	})
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
	// res.VirtualHubEffectiveRouteList = armnetwork.VirtualHubEffectiveRouteList{
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/EffectiveRoutesListForVirtualHub.json
func ExampleVirtualHubsClient_BeginGetEffectiveVirtualHubRoutes_effectiveRoutesForTheVirtualHub() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualHubsClient().BeginGetEffectiveVirtualHubRoutes(ctx, "rg1", "virtualHub1", &armnetwork.VirtualHubsClientBeginGetEffectiveVirtualHubRoutesOptions{EffectiveRoutesParameters: nil})
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
	// res.VirtualHubEffectiveRouteList = armnetwork.VirtualHubEffectiveRouteList{
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/GetInboundRoutes.json
func ExampleVirtualHubsClient_BeginGetInboundRoutes() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualHubsClient().BeginGetInboundRoutes(ctx, "rg1", "virtualHub1", armnetwork.GetInboundRoutesParameters{
		ConnectionType: to.Ptr("ExpressRouteConnection"),
		ResourceURI:    to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteGateways/exrGw1/expressRouteConnections/exrConn1"),
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
	// res.EffectiveRouteMapRouteList = armnetwork.EffectiveRouteMapRouteList{
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/GetOutboundRoutes.json
func ExampleVirtualHubsClient_BeginGetOutboundRoutes() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualHubsClient().BeginGetOutboundRoutes(ctx, "rg1", "virtualHub1", armnetwork.GetOutboundRoutesParameters{
		ConnectionType: to.Ptr("ExpressRouteConnection"),
		ResourceURI:    to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/expressRouteGateways/exrGw1/expressRouteConnections/exrConn1"),
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
	// res.EffectiveRouteMapRouteList = armnetwork.EffectiveRouteMapRouteList{
	// }
}
