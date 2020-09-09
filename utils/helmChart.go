/*
Package utils holds the utility files and methods for all the commands
and subcommands og kgmod
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
package utils

import chartUtil "helm.sh/helm/v3/pkg/chartutil"

// CreateHelmChart creates the helm chart on specified name
func CreateHelmChart(appName string) {
	_, err := chartUtil.Create(appName, GetCWD())
	if err != nil {
		Errorf("error while creating helm chart, %v", err)
	}
}
