package cmds

import (
	"os"
	"os/exec"
)

const (
	defaultLocalPath = "."
)

func createPlatformArguments(platforms []string) []string {
	args := []string{}

	for _, platform := range platforms {
		args = append(args, platformFlag, platform)
	}

	return args
}

func getPath(name string) string {
	found, err := exec.LookPath(name)
	if err == nil {
		return found
	}

	local, err := os.Getwd()
	if err == nil {
		return local
	}

	return defaultLocalPath
}
