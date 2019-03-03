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
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var dateString string
var yesterday bool
var dateFormat = "2006-01-02"

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "List the entries from a specific day",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("log called")
	},
}

func init() {
	rootCmd.AddCommand(logCmd)

	logCmd.Flags().StringVarP(&dateString, "date", "d", time.Now().Format(dateFormat), "set the date to be logged in the format: YYYY-MM-DD")
	logCmd.Flags().BoolVarP(&yesterday, "yesterday", "y", false, "list the yesterday's entries")
}