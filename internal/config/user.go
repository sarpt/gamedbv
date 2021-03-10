package config

import (
	"encoding/json"
	"os"
	"path"

	jsonConfig "github.com/sarpt/gamedbv/internal/config/json"
)

const (
	userConfigFilename = "gamedbv.json"
)

func readUserConfig() (jsonConfig.Project, error) {
	jsonApp := jsonConfig.Project{}

	filename, err := userConfigPath()
	if err != nil {
		return jsonApp, err
	}

	configPayload, err := os.ReadFile(filename)
	if err != nil {
		return jsonApp, err
	}

	err = json.Unmarshal(configPayload, &jsonApp)

	return jsonApp, err
}

func userConfigPath() (string, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return path.Join(userConfigDir, userConfigFilename), nil
}

func userConfigExists() bool {
	filepath, err := userConfigPath()
	if err != nil {
		return false
	}

	_, err = os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

func writeUserConfig(project jsonConfig.Project) error {
	payload, err := json.Marshal(project)
	if err != nil {
		return err
	}

	path, err := userConfigPath()
	if err != nil {
		return err
	}

	err = os.WriteFile(path, payload, os.FileMode(0644))

	return err
}
