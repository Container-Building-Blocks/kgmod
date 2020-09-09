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

import (
	"github.com/asaskevich/govalidator"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"strings"
)

// ReplaceGoModules will replaces the given go module
func ReplaceGoModules() {
	if len(viper.GetStringSlice("replace")) != 0 {
		modules := viper.GetStringSlice("replace")
		replaceModules(modules)
	}
}

// ReplaceGoAddModules replaces the additional go modules from the dynamic flag
func ReplaceGoAddModules(replacingModules []string) {
	if len(replacingModules) != 0 {
		replaceModules(replacingModules)
	}
}

// replaceModules actual implementation of the replace module logic
func replaceModules(modules []string) {
	for _, mod := range modules {
		if strings.Contains(mod, "=") {
			replaceComponents := strings.Split(mod, "=")
			if govalidator.IsURL(replaceComponents[0]) {
				execReplaceModCmd(replaceComponents)
			} else {
				replaceModuleError()
			}
		} else {
			replaceModuleError()
		}
	}
}

func execReplaceModCmd(replaceComponents []string) {
	cmdGoModEditReplace := &exec.Cmd{
		Path:   GoExecPath,
		Args:   []string{GoExecPath, "mod", "edit", "-replace=" + strings.TrimSpace(replaceComponents[0]) + "=" + strings.TrimSpace(replaceComponents[1])},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}
	logrus.Debugf("the command to be executed is %s ", cmdGoModEditReplace.String())

	if err := cmdGoModEditReplace.Run(); err != nil {
		logrus.Fatalf("following error occurred while replacing the go module %s to %s:\n %v", replaceComponents[0], replaceComponents[1], err)
	}
}

// replaceModuleError defines the error when the encountered URL is not valid
// or any other issues
func replaceModuleError() {
	logrus.Errorf(`incorrect replace definition. define like below in the YAML file
replace:
  - github.com/gkarthiks/k8s-discovery=/Users/gkarthiks/k8s-discovery

					or
kgmod init -p <pkg-name> -r github.com/gkarthiks/k8s-discovery=/Users/gkarthiks/k8s-discovery

in the command line flags.
`)
}

// CheckModuleInit validates if the mod has been already initialized
func CheckModuleInit() bool {
	path, err := os.Getwd()
	if err != nil {
		Errorf("error while getting the current working directory: %v", err)
	}
	if _, err := os.Stat(path + "go.mod"); err == nil {
		return true
	}
	return false
}
