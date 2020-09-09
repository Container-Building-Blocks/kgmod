/*
Package cmd holds the initialization command. This subcommand deals with the instantiation
of the kgmod cli process
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
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"kgmod/utils"
	"strings"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the go mod project with kgmod standards",
	Long:  `Initialize the go mod project with kgmod standards.`,
	Run: func(cmd *cobra.Command, args []string) {
		pkgName, _ := cmd.Flags().GetString("pkg-name")
		if utils.CreateModule(pkgName) {
			utils.GetGoModules()
		}
		additionalModules, err := cmd.Flags().GetString("module")
		if err != nil {
			logrus.Errorf("error while getting input for modules flag: %v", err)
		} else if additionalModules != "" {
			slicedAddtnlModules := strings.Split(additionalModules, ",")
			utils.GetGoAddModules(slicedAddtnlModules)
		}

		utils.ReplaceGoModules()
		replaceModules, err := cmd.Flags().GetString("replace")
		if err != nil {
			logrus.Errorf("error while parsing the input for replace flag: %v", err)
		} else if replaceModules != "" {
			slicedReplacingModules := strings.Split(replaceModules, ",")
			utils.ReplaceGoAddModules(slicedReplacingModules)
		}

		dockerFile, err := cmd.Flags().GetBool("docker")
		if err != nil {
			utils.Error("error occurred while parsing the docker flag")
		} else if dockerFile || viper.GetBool("docker") {
			utils.CreateDockerfile(pkgName)
		}

		chartHelm, err := cmd.Flags().GetBool("chart-helm")
		if err != nil {
			utils.Error("error occurred while parsing the chart-helm flag")
		} else if chartHelm || viper.GetBool("chart-helm") {
			utils.CreateHelmChart(pkgName)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().SortFlags = false
	initCmd.Flags().StringP("pkg-name", "p", "", "Provide your module name")
	_ = initCmd.MarkFlagRequired("pkg-name")
	initCmd.Flags().StringP("module", "m", "", "Provide additional dynamic modules' name to be added (comma separated)")
	initCmd.Flags().StringP("replace", "r", "", "Provide additional dynamic modules for the replacement of source modules (comma separated)")
	initCmd.Flags().BoolP("docker", "d", false, "Creates a Dockerfile when enabled")
	initCmd.Flags().BoolP("chart-helm", "c", false, "Creates a skeleton helm project in the current working directory")
}
