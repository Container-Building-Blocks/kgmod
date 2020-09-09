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
	"io"
	"net/http"
	"os"
)

// DownloadFile downloads the file from the specified URL and writes
// it to the current working directory
func DownloadFile(url string) error {
	response, err := http.Get(url)
	if err != nil {
		Errorf("error occurred while pulling the config file from github: %v", err)
	}
	defer response.Body.Close()
	out, err := os.Create(GetCWD() + "/.kgmod.yaml")
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, response.Body)
	return err
}
