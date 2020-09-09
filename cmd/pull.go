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

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pulls the basic recommended configurations from GitHub",
	Long: `Pulls the recommended configurations for the kgmod.yaml file from the GitHub's kgmod
repository and stores in the current working directory as .kgmod.yaml.
This file can also be pulled from a private location as well.'

For example(provide the raw URL of the file):

kgmod pull -l=https://example.com/kgmod.yaml

⚠️ By default the kgmod.yaml file is pulled from github.com/container-building-blocks/kgmod/kgmod.yaml
`,
	Run: func(cmd *cobra.Command, args []string) {
		location, err := cmd.Flags().GetString("location")
		if err != nil {
			utils.Errorf("error while reading input for the location flag %v", err)
		} else if len(location) > 0 && location != "" {
			utils.ConfigFileLocation = location
		}
		utils.DownloadFile(utils.ConfigFileLocation)
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
	pullCmd.Flags().StringP("location", "l", "", "Provide a location to pull the config file from.")
}
