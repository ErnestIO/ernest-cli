/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package command

import (
	"runtime"

	"github.com/ernestio/ernest-cli/helper"
	"github.com/urfave/cli"

	h "github.com/ernestio/ernest-cli/helper"
)

// CmdLog : Preferences setup
var CmdLog = cli.Command{
	Name:        "log",
	Usage:       h.T("log.usage"),
	ArgsUsage:   h.T("log.args"),
	Description: h.T("log.description"),
	Flags: []cli.Flag{
		tBoolFlag("log.flags.raw"),
	},
	Action: func(c *cli.Context) error {
		client := esetup(c, AuthUsersValidation)

		if c.Bool("raw") {
			_ = helper.PrintRawLogs(client.Logger().Stream())
		} else {
			_ = helper.PrintLogs(client.Logger().Stream())
		}

		runtime.Goexit()

		return nil
	},
}
