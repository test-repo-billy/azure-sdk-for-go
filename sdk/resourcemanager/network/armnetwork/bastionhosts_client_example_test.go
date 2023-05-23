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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/BastionHostDelete.json
func ExampleBastionHostsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBastionHostsClient().BeginDelete(ctx, "rg1", "bastionhosttenant", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/BastionHostGet.json
func ExampleBastionHostsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBastionHostsClient().Get(ctx, "rg1", "bastionhosttenant'", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BastionHost = armnetwork.BastionHost{
	// 	Name: to.Ptr("bastionhost'"),
	// 	Type: to.Ptr("Microsoft.Network/bastionHosts"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/bastionHosts/bastionhosttenant'"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 	},
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.BastionHostPropertiesFormat{
	// 		DisableCopyPaste: to.Ptr(false),
	// 		DNSName: to.Ptr("bst-9d89d361-100e-4c01-b92d-466548c476dc.bastion.azure.com"),
	// 		EnableIPConnect: to.Ptr(false),
	// 		EnableKerberos: to.Ptr(false),
	// 		EnableShareableLink: to.Ptr(false),
	// 		EnableTunneling: to.Ptr(false),
	// 		IPConfigurations: []*armnetwork.BastionHostIPConfiguration{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/bastionHosts/bastionhosttenant/bastionHostIpConfigurations/bastionHostIpConfiguration"),
	// 				Name: to.Ptr("bastionHostIpConfiguration"),
	// 				Type: to.Ptr("Microsoft.Network/bastionHosts/bastionHostIpConfigurations"),
	// 				Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 				Properties: &armnetwork.BastionHostIPConfigurationPropertiesFormat{
	// 					PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					PublicIPAddress: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName"),
	// 					},
	// 					Subnet: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/BastionHostSubnet"),
	// 					},
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ScaleUnits: to.Ptr[int32](2),
	// 	},
	// 	SKU: &armnetwork.SKU{
	// 		Name: to.Ptr(armnetwork.BastionHostSKUNameStandard),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/BastionHostPut.json
func ExampleBastionHostsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBastionHostsClient().BeginCreateOrUpdate(ctx, "rg1", "bastionhosttenant", armnetwork.BastionHost{
		Properties: &armnetwork.BastionHostPropertiesFormat{
			IPConfigurations: []*armnetwork.BastionHostIPConfiguration{
				{
					Name: to.Ptr("bastionHostIpConfiguration"),
					Properties: &armnetwork.BastionHostIPConfigurationPropertiesFormat{
						PublicIPAddress: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName"),
						},
						Subnet: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/BastionHostSubnet"),
						},
					},
				}},
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
	// res.BastionHost = armnetwork.BastionHost{
	// 	Name: to.Ptr("bastionhost"),
	// 	Type: to.Ptr("Microsoft.Network/bastionHosts"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/bastionHosts/bastionhosttenant'"),
	// 	Location: to.Ptr("West US"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.BastionHostPropertiesFormat{
	// 		DisableCopyPaste: to.Ptr(false),
	// 		DNSName: to.Ptr("bst-9d89d361-100e-4c01-b92d-466548c476dc.bastion.azure.com"),
	// 		EnableIPConnect: to.Ptr(false),
	// 		EnableKerberos: to.Ptr(false),
	// 		EnableShareableLink: to.Ptr(false),
	// 		EnableTunneling: to.Ptr(false),
	// 		IPConfigurations: []*armnetwork.BastionHostIPConfiguration{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/bastionHosts/bastionhosttenant/bastionHostIpConfigurations/bastionHostIpConfiguration"),
	// 				Name: to.Ptr("bastionHostIpConfiguration"),
	// 				Type: to.Ptr("Microsoft.Network/bastionHosts/bastionHostIpConfigurations"),
	// 				Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 				Properties: &armnetwork.BastionHostIPConfigurationPropertiesFormat{
	// 					PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					PublicIPAddress: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName"),
	// 					},
	// 					Subnet: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/BastionHostSubnet"),
	// 					},
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ScaleUnits: to.Ptr[int32](2),
	// 	},
	// 	SKU: &armnetwork.SKU{
	// 		Name: to.Ptr(armnetwork.BastionHostSKUNameStandard),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/BastionHostPatch.json
func ExampleBastionHostsClient_BeginUpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBastionHostsClient().BeginUpdateTags(ctx, "rg1", "bastionhosttenant", armnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
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
	// res.BastionHost = armnetwork.BastionHost{
	// 	Name: to.Ptr("bastionhosttenant"),
	// 	Type: to.Ptr("Microsoft.Network/bastionHosts"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/bastionHosts/bastionhosttenant'"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.BastionHostPropertiesFormat{
	// 		DisableCopyPaste: to.Ptr(false),
	// 		DNSName: to.Ptr("bst-9d89d361-100e-4c01-b92d-466548c476dc.bastion.azure.com"),
	// 		EnableIPConnect: to.Ptr(false),
	// 		EnableKerberos: to.Ptr(false),
	// 		EnableShareableLink: to.Ptr(false),
	// 		EnableTunneling: to.Ptr(false),
	// 		IPConfigurations: []*armnetwork.BastionHostIPConfiguration{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/bastionHosts/bastionhosttenant/bastionHostIpConfigurations/bastionHostIpConfiguration"),
	// 				Name: to.Ptr("bastionHostIpConfiguration"),
	// 				Type: to.Ptr("Microsoft.Network/bastionHosts/bastionHostIpConfigurations"),
	// 				Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 				Properties: &armnetwork.BastionHostIPConfigurationPropertiesFormat{
	// 					PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					PublicIPAddress: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName"),
	// 					},
	// 					Subnet: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/BastionHostSubnet"),
	// 					},
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ScaleUnits: to.Ptr[int32](2),
	// 	},
	// 	SKU: &armnetwork.SKU{
	// 		Name: to.Ptr(armnetwork.BastionHostSKUNameStandard),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/BastionHostListBySubscription.json
func ExampleBastionHostsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBastionHostsClient().NewListPager(nil)
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
		// page.BastionHostListResult = armnetwork.BastionHostListResult{
		// 	Value: []*armnetwork.BastionHost{
		// 		{
		// 			Name: to.Ptr("bastionhost'"),
		// 			Type: to.Ptr("Microsoft.Network/bastionHosts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/bastionHosts/bastionhosttenant'"),
		// 			Location: to.Ptr("West US"),
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.BastionHostPropertiesFormat{
		// 				DisableCopyPaste: to.Ptr(false),
		// 				DNSName: to.Ptr("bst-9d89d361-100e-4c01-b92d-466548c476dc.bastion.azure.com"),
		// 				EnableIPConnect: to.Ptr(false),
		// 				EnableKerberos: to.Ptr(false),
		// 				EnableShareableLink: to.Ptr(false),
		// 				EnableTunneling: to.Ptr(false),
		// 				IPConfigurations: []*armnetwork.BastionHostIPConfiguration{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/bastionHosts/bastionhosttenant/bastionHostIpConfigurations/bastionHostIpConfiguration"),
		// 						Name: to.Ptr("bastionHostIpConfiguration"),
		// 						Type: to.Ptr("Microsoft.Network/bastionHosts/bastionHostIpConfigurations"),
		// 						Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 						Properties: &armnetwork.BastionHostIPConfigurationPropertiesFormat{
		// 							PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							PublicIPAddress: &armnetwork.SubResource{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName"),
		// 							},
		// 							Subnet: &armnetwork.SubResource{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/BastionHostSubnet"),
		// 							},
		// 						},
		// 				}},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				ScaleUnits: to.Ptr[int32](2),
		// 			},
		// 			SKU: &armnetwork.SKU{
		// 				Name: to.Ptr(armnetwork.BastionHostSKUNameStandard),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/BastionHostListByResourceGroup.json
func ExampleBastionHostsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBastionHostsClient().NewListByResourceGroupPager("rg1", nil)
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
		// page.BastionHostListResult = armnetwork.BastionHostListResult{
		// 	Value: []*armnetwork.BastionHost{
		// 		{
		// 			Name: to.Ptr("bastionhost'"),
		// 			Type: to.Ptr("Microsoft.Network/bastionHosts"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/bastionHosts/bastionhosttenant'"),
		// 			Location: to.Ptr("West US"),
		// 			Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 			Properties: &armnetwork.BastionHostPropertiesFormat{
		// 				DisableCopyPaste: to.Ptr(false),
		// 				DNSName: to.Ptr("bst-9d89d361-100e-4c01-b92d-466548c476dc.bastion.azure.com"),
		// 				EnableIPConnect: to.Ptr(false),
		// 				EnableKerberos: to.Ptr(false),
		// 				EnableShareableLink: to.Ptr(false),
		// 				EnableTunneling: to.Ptr(false),
		// 				IPConfigurations: []*armnetwork.BastionHostIPConfiguration{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/bastionHosts/bastionhosttenant/bastionHostIpConfigurations/bastionHostIpConfiguration"),
		// 						Name: to.Ptr("bastionHostIpConfiguration"),
		// 						Type: to.Ptr("Microsoft.Network/bastionHosts/bastionHostIpConfigurations"),
		// 						Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
		// 						Properties: &armnetwork.BastionHostIPConfigurationPropertiesFormat{
		// 							PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							PublicIPAddress: &armnetwork.SubResource{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName"),
		// 							},
		// 							Subnet: &armnetwork.SubResource{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/BastionHostSubnet"),
		// 							},
		// 						},
		// 				}},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				ScaleUnits: to.Ptr[int32](2),
		// 			},
		// 			SKU: &armnetwork.SKU{
		// 				Name: to.Ptr(armnetwork.BastionHostSKUNameStandard),
		// 			},
		// 	}},
		// }
	}
}
