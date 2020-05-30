package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sarpt/gamedbv/internal/progress"
)

// JSONPrinter is responsible for presenting information to the stdout and stderr as a JSON
// It implements progress.Notifier and db.Results
type JSONPrinter struct{}

// NewJSONPrinter initializes printer
func NewJSONPrinter() JSONPrinter {
	printer := JSONPrinter{}

	return printer
}

// NextStatus should be used for regular messages from function execution
func (printer JSONPrinter) NextStatus(status progress.Status) {
	payload, err := json.Marshal(status)
	if err != nil {
		printer.NextError(err)
		return
	}

	fmt.Fprintf(os.Stdout, string(payload))
}

// NextError should be used for error from which program cannot recover
func (printer JSONPrinter) NextError(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
}
