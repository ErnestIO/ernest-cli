/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package view

import (
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"

	emodels "github.com/ernestio/ernest-go-sdk/models"
)

// PrintEnvInfo : Pretty print for build info
func PrintEnvInfo(env *emodels.Environment, build *emodels.Build) {
	fmt.Println("================\nPlatform Details\n================\n ")
	parts := strings.Split(env.Name, "/")
	fmt.Println("Name : " + parts[1])
	fmt.Println("Status : " + build.Status)
	fmt.Println("Project : " + parts[0])
	if env.Options != nil {
		if env.Options["sync_interval"] != nil {
			fmt.Println("Sync Schedule : " + env.Options["sync_interval"].(string))
		}
		if env.Options["sync_control"] != nil {
			fmt.Println("Sync Control : " + env.Options["sync_control"].(string))
		}
	}
	fmt.Println("Members:")
	for _, v := range env.Members {
		fmt.Printf("  %s (%s)\n", v.User, v.Role)
	}

	fmt.Println("Date : " + build.CreatedAt)

	if len(build.VPCs) > 0 {
		fmt.Println("\nVPCs:")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "ID", "Subnet"})
		for _, v := range build.VPCs {
			table.Append([]string{v.Name, v.ID, v.Subnet})
		}
		table.Render()
	}

	if len(build.ELBs) > 0 {
		fmt.Println("\nELBs:")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "DNS Name"})
		for _, v := range build.ELBs {
			table.Append([]string{v.Name, v.DNSName})
		}
		table.Render()
	}

	if len(build.Networks) > 0 {
		fmt.Println("\nNetworks:")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "ID", "Availability Zone"})
		for _, v := range build.Networks {
			table.Append([]string{v.Name, v.Subnet, v.AvailabilityZone})
		}
		table.Render()
	}

	if len(build.Instances) > 0 {
		fmt.Println("\nInstances:")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "ID", "Public IP", "Private IP"})
		for _, v := range build.Instances {
			table.Append([]string{v.Name, v.InstanceAWSID, v.PublicIP, v.IP})
		}
		table.Render()
	}

	if len(build.Nats) > 0 {
		fmt.Println("\nNAT gateways:")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "ID", "IP"})
		for _, v := range build.Nats {
			table.Append([]string{v.Name, v.NatGatewayAWSID, v.IP})
		}
		table.Render()
	}

	if len(build.SecurityGroups) > 0 {
		fmt.Println("\nSecurity groups:")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Group ID"})
		for _, v := range build.SecurityGroups {
			table.Append([]string{v.Name, v.SecurityGroupAWSID})
		}
		table.Render()
	}

	if len(build.RDSClusters) > 0 {
		fmt.Println("\nRDS Clusters:")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Endpoint"})
		for _, v := range build.RDSClusters {
			table.Append([]string{v.Name, v.Endpoint})
		}
		table.Render()
	}

	if len(build.RDSInstances) > 0 {
		fmt.Println("\nRDS Instances:")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Endpoint"})
		for _, v := range build.RDSInstances {
			table.Append([]string{v.Name, v.Endpoint})
		}
		table.Render()
	}

	if len(build.EBSVolumes) > 0 {
		fmt.Println("\nEBS Volumes:")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Volume ID"})
		for _, v := range build.EBSVolumes {
			table.Append([]string{v.Name, v.VolumeAWSID})
		}
		table.Render()
	}

	if len(build.LoadBalancers) > 0 {
		fmt.Println("\nLoad Balancers:")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "IP"})
		for _, v := range build.LoadBalancers {
			table.Append([]string{v.Name, v.PublicIP})
		}
		table.Render()
	}

	if len(build.VirtualMachines) > 0 {
		fmt.Println("\nVirtual Machines:")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Public IP", "Private IP"})
		for _, v := range build.VirtualMachines {
			table.Append([]string{v.Name, v.PublicIP, v.PrivateIP})
		}
		table.Render()
	}

	if len(build.SQLDatabases) > 0 {
		fmt.Println("\nSQL Databases:")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Server Name"})
		for _, v := range build.SQLDatabases {
			table.Append([]string{v.Name, v.ServerName})
		}
		table.Render()
	}

	if len(build.IamPolicies) > 0 {
		fmt.Println("\nIAM Policies:")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "ID", "Path"})
		for _, v := range build.IamPolicies {
			table.Append([]string{v.Name, v.ID, v.Path})
		}
		table.Render()
	}

	if len(build.IamRoles) > 0 {
		fmt.Println("\nIAM Roles:")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "ID", "Path", "Policies"})
		for _, v := range build.IamRoles {
			table.Append([]string{v.Name, v.ID, v.Path, strings.Join(v.Policies, ",")})
		}
		table.Render()
	}

	if len(build.IamInstanceProfiles) > 0 {
		fmt.Println("\nIAM Instance Profiles:")
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "ID", "Path", "Roles"})
		for _, v := range build.IamInstanceProfiles {
			table.Append([]string{v.Name, v.ID, v.Path, strings.Join(v.Roles, ",")})
		}
		table.Render()
	}
}
