package cli

import (
	"fmt"
	"os"

	"github.com/sarpt/gamedbv/pkg/progress"
)

// TextPrinter is responsible for presenting information to the CLI
// It implements progress.Notifier and db.Results
type TextPrinter struct{}

// NewTextPrinter initializes printer
func NewTextPrinter() TextPrinter {
	printer := TextPrinter{}

	return printer
}

// NextStatus should be used for regular messages from function execution
func (printer TextPrinter) NextStatus(status progress.Status) {
	fmt.Println(fmt.Sprintf("%s: %s", status.Step, status.Message))
}

// NextError should be used for error from which program cannot recover
func (printer TextPrinter) NextError(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
}

// Close stops Printer
func (printer TextPrinter) Close() {}
