package json

import (
	// Needed to read embedded default json config.
	_ "embed"
	"encoding/json"
)

//go:embed "default-config.json"
var defaultPayload []byte

// Default returns embedded in a binary during compile-time default config as a Project instance.
func Default() (Project, error) {
	defaultApp := Project{}

	err := json.Unmarshal(defaultPayload, &defaultApp)
	return defaultApp, err
}
