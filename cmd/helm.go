/*
Package cmd holds the helm command. This subcommand deals with the instantiation
of the helm chart
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

// helmCmd represents the helm command
var helmCmd = &cobra.Command{
	Use:   "helm",
	Short: "Creates the skeleton helm chart for the application",
	Long: `Creates the skeleton helm project in the current working directory. This is similar to the --helm flag in the init command, but you need to pass the app name. For example:

kgmod helm --app-name kgmod-svc`,
	Run: func(cmd *cobra.Command, args []string) {
		appName, _ := cmd.Flags().GetString("app-name")
		utils.CreateHelmChart(appName)
	},
}

func init() {
	rootCmd.AddCommand(helmCmd)
	helmCmd.Flags().StringP("app-name", "a", "", "Provide the application name")
	helmCmd.MarkFlagRequired("app-name")
}
