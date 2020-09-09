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
	"kgmod/tmpl"
	"os"
)

// CreateDockerfile creates the docker file based on the specified
// go tmpl in the current directory
func CreateDockerfile(pkgName string) {
	dockerTmplDef := make(map[string]string)
	dockerTmplDef["AppName"] = pkgName
	dockerTmplDef["GoVersion"] = getGoVersion()
	dockerTmplDef["AlpineVersion"] = "3.10"
	out, err := os.Create(GetCWD() + "/Dockerfile")
	if err != nil {
		logrus.Error(err)
	}
	defer out.Close()
	tmpl.Dockerfile_Tmpl.ExecuteTemplate(out, "dockerfile", dockerTmplDef)
}
