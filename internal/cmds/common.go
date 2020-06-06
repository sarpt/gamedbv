package cmds

import (
	"os"
	"os/exec"
	"path"
)

const (
	defaultLocalPath = "."
	longArgPrefix    = "--"
)

func createPlatformsArguments(platforms []string) []string {
	return createMultipleFlagArguments(PlatformFlag, platforms)
}

func createJSONArgument(isSet bool) []string {
	if isSet {
		return []string{longArgument(JSONFlag)}
	}

	return []string{}
}

func createMultipleFlagArguments(flagName string, values []string) []string {
	args := []string{}

	for _, val := range values {
		args = append(args, longArgument(flagName), val)
	}

	return args
}

func longArgument(flagName string) string {
	return longArgPrefix + flagName
}

func getParentDirPath(name string) string {
	found, err := exec.LookPath(name)
	if err == nil {
		return path.Dir(found)
	}

	local, err := os.Getwd()
	if err == nil {
		return local
	}

	return defaultLocalPath
}
