package apps

// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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

import (
	"github.com/spf13/cobra"
)

func Operator(args []string) error {
	if args[0] == "init" {
		// deployer, _ := integration.NewDeployerWithDefaultKubeconfig()
		// deployer.Run()
	}
	return nil
}

var OperatorCommand = &cobra.Command{
	Use:   "operator initialize",
	Short: "initialize the blackduck operator on your cluster",
	Run: func(cmd *cobra.Command, args []string) {
		if err := Operator(args); err != nil {
			cmd.Help()
		}
	},
}

// implementing init is important ! thats how cobra knows to bind your 'app' to top level command.
func init() {
	RootCmd.AddCommand(ClusterCommand)
}
