/*
Package cmd holds the docker command. This subcommand deals with the instantiation
of the Dockerfile
*/
/*
Copyright © 2020 Karthikeyan Govindaraj <github.gkarthiks@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"kgmod/utils"
)

// dockerCmd represents the docker command
var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "A command to create a skeleton Dockerfile",
	Long: `A command that created a skeleton Dockerfile for you to start work with.
This is similar to the --docker flag with the init command, but you need to pass the app name here. For example:

kgmod docker --app-name kgmod-svc`,
	Run: func(cmd *cobra.Command, args []string) {
		appName, _ := cmd.Flags().GetString("app-name")
		utils.CreateDockerfile(appName)
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
	dockerCmd.Flags().StringP("app-name", "a", "", "Provide the application name")
	dockerCmd.MarkFlagRequired("app-name")
}
