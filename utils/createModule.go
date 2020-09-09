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

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"strings"
)

// CreateModule creates the module with the specified package name
func CreateModule(pkgName string) bool {
	logrus.Debugf("Creating module on package name %s", pkgName)
	cmdGoModInit := &exec.Cmd{
		Path:   GoExecPath,
		Args:   []string{GoExecPath, "mod", "init", pkgName},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}
	logrus.Debugf("the command to be executed is %s ", cmdGoModInit.String())

	if err := cmdGoModInit.Run(); err != nil {
		logrus.Errorf("error while initializing the go module: %v", err)
		return false
	}
	return true
}

// GetGoModules downloads the modules specified already in the config file
func GetGoModules() {
	if len(viper.GetStringSlice("modules")) != 0 {
		modules := viper.GetStringSlice("modules")
		downloadModules(modules)
	}
}

// GetGoAddModules downloads the additional modules specified via flag
func GetGoAddModules(additionalModules []string) {
	if len(additionalModules) != 0 {
		downloadModules(additionalModules)
	}
}

// downloadModules downloads the modules
func downloadModules(modules []string) {
	for _, mod := range modules {
		cmdGoModGet := &exec.Cmd{
			Path:   GoExecPath,
			Args:   []string{GoExecPath, "get", strings.TrimSpace(mod)},
			Stdout: os.Stdout,
			Stderr: os.Stdout,
		}
		logrus.Debugf("the command to be executed is %s ", cmdGoModGet.String())

		if err := cmdGoModGet.Run(); err != nil {
			logrus.Fatalf("following error occurred while getting the go module %s: %v", mod, err)
		}
	}
}
