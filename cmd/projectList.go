// Copyright © 2019 Lucas dos Santos Abreu <lucas.s.abreu@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"io"
	"os"

	"github.com/lucassabreu/clockify-cli/api"
	"github.com/lucassabreu/clockify-cli/api/dto"
	"github.com/lucassabreu/clockify-cli/reports"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// projectListCmd represents the projectList command
var projectListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List projects on Clockify and project links",
	RunE: withClockifyClient(func(cmd *cobra.Command, args []string, c *api.Client) error {
		format, _ := cmd.Flags().GetString("format")
		quiet, _ := cmd.Flags().GetBool("quiet")
		name, _ := cmd.Flags().GetString("name")
		archived, _ := cmd.Flags().GetBool("archived")

		projects, err := c.GetProjects(api.GetProjectsParam{
			Workspace:       viper.GetString("workspace"),
			Name:            name,
			Archived:        archived,
			PaginationParam: api.PaginationParam{AllPages: true},
		})

		if err != nil {
			return err
		}

		var reportFn func([]dto.Project, io.Writer) error

		reportFn = reports.ProjectPrint
		if format != "" {
			reportFn = reports.ProjectPrintWithTemplate(format)
		}

		if quiet {
			reportFn = reports.ProjectPrintQuietly
		}

		return reportFn(projects, os.Stdout)
	}),
}

func init() {
	projectCmd.AddCommand(projectListCmd)

	projectListCmd.Flags().StringP("name", "n", "", "will be used to filter the tag by name")
	projectListCmd.Flags().BoolP("archived", "", false, "list only archived projects")
	projectListCmd.Flags().StringP("format", "f", "", "golang text/template format to be applied on each Project")
	projectListCmd.Flags().BoolP("quiet", "q", false, "only display ids")
}
