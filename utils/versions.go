package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

// GetVersion prints the version kgmod and all dependent binaries
func GetVersion() {
	fmt.Printf("======================== kgmod Version ========================\n%s\n\n", BuildVersion)

	fmt.Printf("======================== GoLang Version ========================\n%s\n", getVersionOnBinaries("go", GoExecPath))

	DockerPath = GetPath("docker")
	fmt.Printf("======================== Docker Version ========================\n%s\n", getVersionOnBinaries("docker", DockerPath))

	HelmPath = GetPath("helm")
	fmt.Printf("======================== Helm Version ========================\n%s\n", getVersionOnBinaries("helm", HelmPath))

}

func getVersionOnBinaries(commandBinaryName, commandBinaryPath string) string {
	version := &exec.Cmd{
		Path:   commandBinaryPath,
		Args:   []string{commandBinaryPath, "version"},
		Stderr: os.Stdout,
	}
	logrus.Debugf("the command to be executed is %s ", version.String())

	if byteOutput, err := version.Output(); err != nil {
		logrus.Errorf("error while finding the version of %s: %v", commandBinaryName, err)
		os.Exit(1)
	} else {
		return string(byteOutput)
	}
	return ""
}
