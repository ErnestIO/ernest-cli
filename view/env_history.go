/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package view

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"

	emodels "github.com/ernestio/ernest-go-sdk/models"
)

// PrintEnvHistory : Pretty print for build history
func PrintEnvHistory(name string, builds []*emodels.Build) {
	if len(builds) == 0 {
		fmt.Println("\nThere are no registered builds for this environment")
		fmt.Println("")
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Name", "Type", "Status", "Version", "User"})
		num := len(builds) + 1
		for _, b := range builds {
			num = num - 1
			id := strconv.Itoa(num)
			table.Append([]string{id, name, b.Type, b.Status, b.CreatedAt, b.Username})
		}
		table.Render()
	}
}
