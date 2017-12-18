/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package command

// CmdProject subcommand
import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	h "github.com/ernestio/ernest-cli/helper"
	"github.com/ernestio/ernest-cli/model"
	"github.com/ernestio/ernest-cli/view"
	"github.com/fatih/color"
	"github.com/urfave/cli"

	emodels "github.com/ernestio/ernest-go-sdk/models"
)

// ListEnvs ...
var ListEnvs = cli.Command{
	Name:        "list",
	Usage:       h.T("envs.list.usage"),
	ArgsUsage:   h.T("envs.list.args"),
	Description: h.T("envs.list.description"),
	Action: func(c *cli.Context) error {
		client := esetup(c, AuthUsersValidation)
		envs := client.Environment().ListAll()
		view.PrintEnvList(envs)

		return nil
	},
}

// UpdateEnv :  Creates an empty environment based on a name and a project name
var UpdateEnv = cli.Command{
	Name:        "update",
	Aliases:     []string{"a"},
	Usage:       h.T("envs.update.usage"),
	ArgsUsage:   h.T("envs.update.args"),
	Description: h.T("envs.update.description"),
	Flags: append([]cli.Flag{
		stringFlagND("sync_interval", "sets the automatic sync interval. Accepts cron syntax, i.e. '@every 1d', '@weekly' or '0 0 * * * *' (Daily at midnight)"),
		stringFlagND("submissions", "allows user build submissions from users that have only read only permission to an environment. Options are 'enable' or 'disable'"),
	}, AllProviderFlags...),
	Action: func(c *cli.Context) error {
		paramsLenValidation(c, 2, "envs.update.args")
		client := esetup(c, AuthUsersValidation)
		env := client.Environment().Get(c.Args()[0], c.Args()[1])
		env.Credentials = ProviderFlagsToSlice(c)
		env.Options = MapEnvOptions(c, env.Options)
		client.Environment().Update(c.Args()[0], env)
		color.Green("Environment successfully updated")

		return nil
	},
}

// CreateEnv :  Creates an empty environment based on a name and a project name
var CreateEnv = cli.Command{
	Name:        "create",
	Aliases:     []string{"a"},
	Usage:       h.T("envs.create.usage"),
	ArgsUsage:   h.T("envs.create.args"),
	Description: h.T("envs.create.description"),
	Flags: append([]cli.Flag{
		stringFlagND("credentials", "will override project information"),
		stringFlagND("sync_interval", "sets the automatic sync interval. Accepts cron syntax, i.e. '@every 1d', '@weekly' or '0 0 * * * *' (Daily at midnight)"),
		stringFlagND("submissions", "allows user build submissions from users that have only read only permission to an environment. Options are 'enable' or 'disable'"),
	}, AllProviderFlags...),
	Action: func(c *cli.Context) error {
		paramsLenValidation(c, 2, "envs.create.args")
		client := esetup(c, AuthUsersValidation)

		env := emodels.Environment{
			Name:        c.Args()[1],
			Project:     c.Args()[0],
			Credentials: ProviderFlagsToSlice(c),
			Options:     MapEnvOptions(c, nil),
		}
		client.Environment().Create(c.Args()[0], &env)
		color.Green("Environment successfully created")

		return nil
	},
}

func mapDefinition(c *cli.Context) *model.Definition {
	file := "ernest.yml"
	if len(c.Args()) == 1 {
		file = c.Args()[0]
	}
	payload, err := ioutil.ReadFile(file)
	if err != nil {
		h.PrintError("You should specify a valid template path or store an ernest.yml on the current folder")
	}
	def := model.Definition{}
	if err := def.Load(payload); err != nil {
		h.PrintError("Could not process definition yaml")
	}
	if err := def.LoadFileImports(); err != nil {
		h.PrintError(err.Error())
	}

	return &def
}

// ApplyEnv command
// Applies changes described on a YAML file to an env
var ApplyEnv = cli.Command{
	Name:        "apply",
	Aliases:     []string{"a"},
	Usage:       h.T("envs.apply.usage"),
	ArgsUsage:   h.T("envs.apply.args"),
	Description: h.T("envs.apply.description"),
	Flags: append([]cli.Flag{
		boolFlag("dry", "print the changes to be applied on an environment intead of applying them"),
		stringFlagND("credentials", "will override project information"),
	}, AllProviderFlags...),
	Action: func(c *cli.Context) error {
		paramsLenValidation(c, 1, "envs.apply.args")
		client := esetup(c, AuthUsersValidation)
		def := mapDefinition(c)

		if _, err := client.Cli().Environments.Get(def.Project, def.Name); err != nil {
			env := emodels.Environment{
				Name:        def.Name,
				Project:     def.Project,
				Credentials: ProviderFlagsToSlice(c),
				Options:     MapEnvOptions(c, nil),
			}
			client.Environment().Create(def.Project, &env)
		}
		payload, err := def.Save()
		if err != nil {
			h.PrintError("Could not finalize definition yaml")
		}
		if c.Bool("dry") == true {
			view.EnvDryII(*client.Build().Dry(payload))
			return nil
		}
		build := client.Build().Create(payload)
		if build.Status == "submitted" {
			color.Green("Build has been succesfully submitted and is awaiting approval.")
			os.Exit(0)
		}

		h.MonitorizeII(client.Build().Stream(build.ID))
		view.PrintEnvInfoII(
			client.Project().Get(def.Project),
			client.Environment().Get(def.Project, def.Name),
			client.Build().Get(def.Project, def.Name, build.GetID()),
		)

		return nil
	},
}

// SyncEnv command
// Syncs an environment with the cloud provider
var SyncEnv = cli.Command{
	Name:        "sync",
	Aliases:     []string{"s"},
	Usage:       h.T("envs.sync.usage"),
	ArgsUsage:   h.T("envs.sync.args"),
	Description: h.T("envs.sync.description"),
	Flags:       append([]cli.Flag{}),
	Action: func(c *cli.Context) error {
		paramsLenValidation(c, 2, "envs.sync.args")
		client := esetup(c, AuthUsersValidation)
		client.Environment().Sync(c.Args()[0], c.Args()[1])

		return nil
	},
}

func buildIDFromIndex(builds []*emodels.Build, index string) *emodels.Build {
	num, _ := strconv.Atoi(index)
	if num < 1 || num > len(builds) {
		h.PrintError("Invalid build ID")
	}
	num = len(builds) - num
	return builds[num]
}

// ReviewEnv command
// Approval for outstanding build submissions
var ReviewEnv = cli.Command{
	Name:        "review",
	Aliases:     []string{"rev"},
	Usage:       h.T("envs.review.usage"),
	ArgsUsage:   h.T("envs.review.args"),
	Description: h.T("envs.review.description"),
	Flags: append([]cli.Flag{
		boolFlag("accept, a", "Accept Sync changes"),
		boolFlag("reject, r", "Reject Sync changes"),
	}),
	Action: func(c *cli.Context) error {
		paramsLenValidation(c, 2, "envs.review.args")
		client := esetup(c, AuthUsersValidation)

		project := c.Args()[0]
		env := c.Args()[1]

		var resolution string
		if c.Bool("accept") {
			resolution = "submission-accepted"
		}
		if c.Bool("reject") {
			resolution = "submission-rejected"
		}

		if resolution == "" {
			builds := client.Build().List(project, env)

			b1 := buildIDFromIndex(builds, strconv.Itoa(len(builds)-1))
			b2 := buildIDFromIndex(builds, strconv.Itoa(len(builds)))
			def1 := client.Build().Definition(project, env, b1.ID)
			def2 := client.Build().Definition(project, env, b2.ID)

			view.PrintEnvDiff(b1.ID, b2.ID, []byte(def1), []byte(def2))

			return nil
		}
		client.Environment().Resolve(project, env, resolution)
		return nil
	},
}

// ResolveEnv command
// Resolves an issue that requires user input
var ResolveEnv = cli.Command{
	Name:        "resolve",
	Aliases:     []string{"r"},
	Usage:       h.T("envs.resolve.usage"),
	ArgsUsage:   h.T("envs.resolve.args"),
	Description: h.T("envs.resolve.description"),
	Flags: append([]cli.Flag{
		boolFlag("accept, a", "Accept Sync changes"),
		boolFlag("reject, r", "Reject Sync changes"),
		boolFlag("ignore, i", "Ignore Sync changes"),
	}),
	Action: func(c *cli.Context) error {
		paramsLenValidation(c, 2, "envs.resolve.args")
		client := esetup(c, AuthUsersValidation)

		var resolution string
		if c.Bool("accept") {
			resolution = "accept-changes"
		}
		if c.Bool("reject") {
			resolution = "reject-changes"
		}
		if c.Bool("ignore") {
			resolution = "ignore-changes"
		}

		if resolution == "" {
			h.PrintError("You should specify a valid resolution [accept|reject|ignore]")
		}
		client.Environment().Resolve(c.Args()[0], c.Args()[1], resolution)
		return nil
	},
}

// DestroyEnv command
var DestroyEnv = cli.Command{
	Name:        "delete",
	Aliases:     []string{"destroy", "d"},
	Usage:       h.T("envs.destroy.usage"),
	ArgsUsage:   h.T("envs.destroy.args"),
	Description: h.T("envs.destroy.description"),
	Flags: []cli.Flag{
		boolFlag("force,f", "Hard ernest env removal."),
		boolFlag("yes,y", "Destroy an environment without prompting confirmation."),
	},
	Action: func(c *cli.Context) error {
		paramsLenValidation(c, 2, "envs.destroy.args")
		client := esetup(c, AuthUsersValidation)

		if c.Bool("force") {
			client.Environment().ForceDeletion(c.Args()[0], c.Args()[1])
		} else {
			if c.Bool("yes") {
				client.Environment().Delete(c.Args()[0], c.Args()[1])
			} else {
				fmt.Print("Do you really want to destroy this environment? (Y/n) ")
				if askForConfirmation() == false {
					return nil
				}
				client.Environment().Delete(c.Args()[0], c.Args()[1])
			}
		}
		color.Green("Environment successfully removed")
		return nil
	},
}

// HistoryEnv command
// Shows the history of an env, a list of builds
var HistoryEnv = cli.Command{
	Name:        "history",
	Usage:       h.T("envs.history.usage"),
	ArgsUsage:   h.T("envs.history.args"),
	Description: h.T("envs.history.description"),
	Action: func(c *cli.Context) error {
		paramsLenValidation(c, 2, "envs.history.args")
		client := esetup(c, AuthUsersValidation)
		envs := client.Build().List(c.Args()[0], c.Args()[1])
		view.PrintEnvHistory(c.Args()[1], envs)
		return nil
	},
}

// ResetEnv command
var ResetEnv = cli.Command{
	Name:        "reset",
	Usage:       h.T("envs.reset.usage"),
	ArgsUsage:   h.T("envs.reset.args"),
	Description: h.T("envs.reset.description"),
	Action: func(c *cli.Context) error {
		paramsLenValidation(c, 2, "envs.reset.args")
		client := esetup(c, AuthUsersValidation)
		client.Environment().Reset(c.Args()[0], c.Args()[1])
		color.Red("You've successfully resetted the environment '" + c.Args()[0] + " / " + c.Args()[1] + "'")

		return nil
	},
}

// RevertEnv command
var RevertEnv = cli.Command{
	Name:        "revert",
	Usage:       h.T("envs.revert.usage"),
	ArgsUsage:   h.T("envs.revert.args"),
	Description: h.T("envs.revert.description"),
	Flags: []cli.Flag{
		boolFlag("dry", "print the changes to be applied on an environment intead of applying them"),
	},
	Action: func(c *cli.Context) error {
		paramsLenValidation(c, 3, "envs.revert.args")
		client := esetup(c, AuthUsersValidation)
		build := client.Build().Get(c.Args()[0], c.Args()[1], c.Args()[2])
		def := client.Build().Definition(c.Args()[0], c.Args()[1], c.Args()[2])

		if c.Bool("dry") == true {
			view.EnvDryII(*client.Build().Dry([]byte(def)))
		} else {
			client.Build().Create([]byte(def))
			if build.Status == "submitted" {
				color.Green("Build has been succesfully submitted and is awaiting approval.")
				os.Exit(0)
			}

			h.MonitorizeII(client.Build().Stream(build.ID))
		}

		return nil
	},
}

// DefinitionEnv command
// Shows the current definition of an environment by its name
var DefinitionEnv = cli.Command{
	Name:        "definition",
	Aliases:     []string{"s"},
	Usage:       h.T("envs.definition.usage"),
	ArgsUsage:   h.T("envs.definition.args"),
	Description: h.T("envs.definition.description"),
	Flags: []cli.Flag{
		stringFlag("build", "", "Build ID"),
	},
	Action: func(c *cli.Context) error {
		paramsLenValidation(c, 2, "envs.definition.args")
		client := esetup(c, AuthUsersValidation)

		build := client.Build().BuildByPosition(c.Args()[0], c.Args()[1], c.String("build"))
		def := client.Build().Definition(c.Args()[0], c.Args()[1], build.ID)
		fmt.Println(def)

		return nil
	},
}

// InfoEnv : Shows detailed information of a env by its name
var InfoEnv = cli.Command{
	Name:        "info",
	Aliases:     []string{"i"},
	Usage:       h.T("envs.info.usage"),
	ArgsUsage:   h.T("envs.info.args"),
	Description: h.T("envs.info.description"),
	Flags: []cli.Flag{
		stringFlag("build", "", "Build ID"),
	},
	Action: func(c *cli.Context) error {
		paramsLenValidation(c, 2, "envs.info.args")
		client := esetup(c, AuthUsersValidation)

		build := client.Build().BuildByPosition(c.Args()[0], c.Args()[1], c.String("build"))
		build = client.Build().Get(c.Args()[0], c.Args()[1], build.ID)
		env := client.Environment().Get(c.Args()[0], c.Args()[1])
		project := client.Project().Get(c.Args()[0])
		view.PrintEnvInfoII(project, env, build)

		return nil
	},
}

// DiffEnv : Shows detailed information of an env by its name
var DiffEnv = cli.Command{
	Name:        "diff",
	Aliases:     []string{"i"},
	Usage:       h.T("envs.diff.usage"),
	ArgsUsage:   h.T("envs.diff.args"),
	Description: h.T("envs.diff.description"),
	Action: func(c *cli.Context) error {
		paramsLenValidation(c, 4, "envs.diff.args")
		client := esetup(c, AuthUsersValidation)

		build1 := client.Build().BuildByPosition(c.Args()[0], c.Args()[1], c.Args()[2])
		build2 := client.Build().BuildByPosition(c.Args()[0], c.Args()[1], c.Args()[3])
		def1 := client.Build().Definition(c.Args()[0], c.Args()[1], build1.GetID())
		def2 := client.Build().Definition(c.Args()[0], c.Args()[1], build2.GetID())
		view.PrintEnvDiff(build1.GetID(), build2.GetID(), []byte(def1), []byte(def2))

		return nil
	},
}

// ImportEnv : Shows detailed information of an env by its name
var ImportEnv = cli.Command{
	Name:        "import",
	Aliases:     []string{"i"},
	Usage:       h.T("envs.import.usage"),
	ArgsUsage:   h.T("envs.import.args"),
	Description: h.T("envs.import.description"),
	Flags: []cli.Flag{
		stringFlag("project", "", "Project name"),
		stringFlag("filters", "", "Import filters comma delimited list"),
	},
	Action: func(c *cli.Context) error {
		paramsLenValidation(c, 2, "envs.import.args")
		client := esetup(c, AuthUsersValidation)

		var filters []string
		if c.String("filters") != "" {
			filters = strings.Split(c.String("filters"), ",")
		}

		a := client.Environment().Import(c.Args()[0], c.Args()[1], filters)
		h.MonitorizeII(client.Build().Stream(a.ResourceID))

		return nil
	},
}

// posString returns the first index of element in slice.
// If slice does not contain element, returns -1.
func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}

// containsString returns true iff slice contains element
func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}

// CmdEnv ...
var CmdEnv = cli.Command{
	Name:    "environment",
	Aliases: []string{"env"},
	Usage:   "Environment related subcommands",
	Subcommands: []cli.Command{
		ListEnvs,
		CreateEnv,
		UpdateEnv,
		ApplyEnv,
		DestroyEnv,
		HistoryEnv,
		ResetEnv,
		RevertEnv,
		DefinitionEnv,
		InfoEnv,
		MonitorEnv,
		DiffEnv,
		ImportEnv,
		SyncEnv,
		ResolveEnv,
		ReviewEnv,
		ScheduleEnv,
	},
}
