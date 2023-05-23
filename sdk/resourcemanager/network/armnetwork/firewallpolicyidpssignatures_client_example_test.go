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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/FirewallPolicyQuerySignatureOverrides.json
func ExampleFirewallPolicyIdpsSignaturesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFirewallPolicyIdpsSignaturesClient().List(ctx, "rg1", "firewallPolicy", armnetwork.IDPSQueryObject{
		Filters: []*armnetwork.FilterItems{
			{
				Field: to.Ptr("Mode"),
				Values: []*string{
					to.Ptr("Deny")},
			}},
		OrderBy: &armnetwork.OrderBy{
			Field: to.Ptr("severity"),
			Order: to.Ptr(armnetwork.FirewallPolicyIDPSQuerySortOrderAscending),
		},
		ResultsPerPage: to.Ptr[int32](20),
		Search:         to.Ptr(""),
		Skip:           to.Ptr[int32](0),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QueryResults = armnetwork.QueryResults{
	// 	MatchingRecordsCount: to.Ptr[int64](2),
	// 	Signatures: []*armnetwork.SingleQueryResult{
	// 		{
	// 			Description: to.Ptr("P2P Phatbot Control Connection"),
	// 			DestinationPorts: []*string{
	// 				to.Ptr("any")},
	// 				Direction: to.Ptr(armnetwork.FirewallPolicyIDPSSignatureDirectionTwo),
	// 				Group: to.Ptr("A Network Trojan was detected"),
	// 				InheritedFromParentPolicy: to.Ptr(false),
	// 				LastUpdated: to.Ptr("2010-07-30T00:00:00"),
	// 				Mode: to.Ptr(armnetwork.FirewallPolicyIDPSSignatureModeTwo),
	// 				Severity: to.Ptr(armnetwork.FirewallPolicyIDPSSignatureSeverityOne),
	// 				SignatureID: to.Ptr[int32](2000015),
	// 				SourcePorts: []*string{
	// 					to.Ptr("any")},
	// 					Protocol: to.Ptr("tcp"),
	// 				},
	// 				{
	// 					Description: to.Ptr("WEB_SERVER SQL sp_delete_alert attempt"),
	// 					DestinationPorts: []*string{
	// 						to.Ptr("any")},
	// 						Direction: to.Ptr(armnetwork.FirewallPolicyIDPSSignatureDirectionOne),
	// 						Group: to.Ptr("Attempted User Privilege Gain"),
	// 						InheritedFromParentPolicy: to.Ptr(false),
	// 						LastUpdated: to.Ptr("2019-09-27T00:00:00"),
	// 						Mode: to.Ptr(armnetwork.FirewallPolicyIDPSSignatureModeTwo),
	// 						Severity: to.Ptr(armnetwork.FirewallPolicyIDPSSignatureSeverityOne),
	// 						SignatureID: to.Ptr[int32](2000106),
	// 						SourcePorts: []*string{
	// 							to.Ptr("any")},
	// 							Protocol: to.Ptr("http"),
	// 					}},
	// 				}
}
