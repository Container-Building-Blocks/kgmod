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
	"os"
	"os/exec"
	"regexp"
)

// GetCWD fetches the current working directory
func GetCWD() (cwd string) {
	cwd, err := os.Getwd()
	if err != nil {
		Errorf("error while finding the Current Working directory, %v", err)
	}
	return
}

// getGoVersion will return the Dockerfile compatible go version
func getGoVersion() string {
	cmdGoVersion := &exec.Cmd{
		Path:   GoExecPath,
		Args:   []string{GoExecPath, "version"},
		Stderr: os.Stdout,
	}
	logrus.Debugf("the command to be executed is %s ", cmdGoVersion.String())

	if byteOutput, err := cmdGoVersion.Output(); err != nil {
		logrus.Errorf("error while finding the version of go module: %v", err)
		os.Exit(1)
	} else {
		strVersion := string(byteOutput)
		goCompleteVersion := regexp.MustCompile(`go[0-9].[0-9][0-9].[0-9]+`)
		goVersionDocker := regexp.MustCompile(`[0-9].[0-9]+`)
		goCmplVer := goCompleteVersion.FindString(strVersion)
		return goVersionDocker.FindString(goCmplVer)
	}
	return ""
}
