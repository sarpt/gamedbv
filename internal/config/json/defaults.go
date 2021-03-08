package json

import (
	// Needed to read embedded default json config.
	_ "embed"
	"encoding/json"
)

//go:embed "default-config.json"
var defaultAppPayload []byte

// DefaultApp returns embedded in a binary during compile-time a default config as a App instance.
func DefaultApp() (App, error) {
	defaultApp := App{}

	err := json.Unmarshal(defaultAppPayload, &defaultApp)
	return defaultApp, err
}
