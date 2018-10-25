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
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

func runCmd(cmd string, args ...string) error {
	cmd2 := exec.Command(cmd, args...)
	stdoutErr, err := cmd2.CombinedOutput()
	fmt.Printf("%s\n", stdoutErr)
	if err != nil {
		fmt.Printf("Error running command !!!")
		return err
	}
	fmt.Printf("%s\n", stdoutErr)
	time.Sleep(1 * time.Second)
	return nil
}

func RunClusterCommand(args []string) error {
	if args == nil || len(args) == 0 {
		return fmt.Errorf("No subcommand provided !")
	} else if args[0] == "status" {
		e := runCmd("kubectl", "cluster-info")
		if e != nil {
			runCmd("oc", "status")
		}
	} else if args[0] == "list" {
		if args[1] == "hubs" {
			runCmd("kubectl", "get", "hubs", "--all-namespaces")
		} else {
			fmt.Println("Error: Did you mean 'list hubs' ?")
		}
	} else {
		return fmt.Errorf("Invalid subcommand provided.")
	}
	return nil
}

var ClusterCommand = &cobra.Command{
	Use:   "cluster",
	Short: "tells you what cluster your on",
	Run: func(cmd *cobra.Command, args []string) {
		if err := RunClusterCommand(args); err != nil {
			cmd.Help()
		}
	},
}

// implementing init is important ! thats how cobra knows to bind your 'app' to top level command.
func init() {
	RootCmd.AddCommand(ClusterCommand)
}
