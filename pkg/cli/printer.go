package cli

import (
	"fmt"

	"github.com/sarpt/gamedbv/pkg/progress"
)

// Printer is responsible for presenting information to the CLI
// It implements progress.Notifier and db.Results
type Printer struct {
	progress chan progress.Status
	errors   chan error
}

// NextStatus should be used for regular messages from function execution
func (printer Printer) NextStatus(status progress.Status) {
	printer.progress <- status
}

// NextError should be used for error from which program cannot recover
func (printer Printer) NextError(err error) {
	printer.errors <- err
}

// Close stops Printer from calling progressHandler and errorsHandler passed to NewCliPrinter
func (printer Printer) Close() {
	close(printer.progress)
	close(printer.errors)
}

// New initializes printer that will call progressHandler and errorsHandler when respective methods will be invoked
func New() Printer {
	printer := Printer{
		progress: make(chan progress.Status),
		errors:   make(chan error),
	}

	go progressReporter(printer.progress)
	go errorsReporter(printer.errors)

	return printer
}

func progressReporter(statuses <-chan progress.Status) {
	for status := range statuses {
		fmt.Println(fmt.Sprintf("%s: %s", status.Step, status.Message))
	}
}

func errorsReporter(errors <-chan error) {
	for err := range errors {
		panic(err)
	}
}
