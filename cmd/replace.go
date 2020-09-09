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
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"kgmod/utils"
)

// replaceCmd represents the replace command
var replaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "A dynamic replace command to replace the source module",
	Long: `A dynamic command to replace the existing modules in the go.mod file.
For example:

kgmod replace -r github.com/gkarthiks/k8s-discovery=/Users/gkarthiks/k8s-discovery

will add a replace statement to the k8s-discovery package pointing to the local folder`,
	Run: func(cmd *cobra.Command, args []string) {
		replacingModules, _ := cmd.Flags().GetString("r")
		if utils.CheckModuleInit() {
			logrus.Println(replacingModules)
		} else {
			utils.Error("error! no module is initiated, initiate the module first")
		}
	},
}

func init() {
	rootCmd.AddCommand(replaceCmd)
	replaceCmd.Flags().StringP("r", "r", "", "Provide additional dynamic modules for the replacement of source modules (comma separated)")
	replaceCmd.MarkFlagRequired("r")
}
