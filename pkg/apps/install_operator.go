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
	"fmt"

	"github.com/jayunit100/blackduckctl/pkg/util"
	"github.com/spf13/cobra"
)

var InstallOperatorCommand = &cobra.Command{
	Use:   "install-operator",
	Short: "Install the blackduck operator!",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			fmt.Printf("No password provided, will use defaults for secrets!")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		util.RunKubeCmd("blarg")
	},
}

// implementing init is important ! thats how cobra knows to bind your 'app' to top level command.
func init() {
	RootCmd.AddCommand(InstallOperatorCommand)
}
