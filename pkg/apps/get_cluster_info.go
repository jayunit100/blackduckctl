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

func RunClusterCommandListHubs() error {
	err := util.RunKubeCmd("get", "hubs", "--all-namespaces")
	return err
}

func RunClustercommandStatus() error {
	e1 := util.RunKubeCmd("cluster-info")
	return e1
	// return both errors.. clumsy but comprehensive.
}

var ClusterCommand = &cobra.Command{
	Use:   "cluster",
	Short: "tells you what cluster your on",
	Args: func(cmd *cobra.Command, args []string) error {
		if args[0] == "list" || args[0] == "status" {
			return fmt.Errorf("Require an action: list, or status.")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "status" {
			RunClustercommandStatus()
		} else {
			RunClusterCommandListHubs()
		}
	},
}

// implementing init is important ! thats how cobra knows to bind your 'app' to top level command.
func init() {
	RootCmd.AddCommand(ClusterCommand)
}
