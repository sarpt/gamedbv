package cli

import "github.com/sarpt/gamedbv/pkg/progress"

// NewPrinter initializes printer that will call progressHandler and errorsHandler when respective methods will be invoked
func NewPrinter() Printer {
	printer := Printer{
		progress: make(chan progress.Status),
		errors:   make(chan error),
	}

	go progressReporter(printer.progress)
	go errorsReporter(printer.errors)

	return printer
}
