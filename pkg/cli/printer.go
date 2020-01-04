package cli

import "fmt"

// Printer is responsible for presenting information to the CLI
type Printer struct {
	progress chan string
	errors   chan error
}

// NextProgress should be used for regular messages from function execution
func (printer Printer) NextProgress(message string) {
	printer.progress <- message
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
		progress: make(chan string),
		errors:   make(chan error),
	}

	go progressReporter(printer.progress)
	go errorsReporter(printer.errors)

	return printer
}

func progressReporter(progress <-chan string) {
	for message := range progress {
		fmt.Println(message)
	}
}

func errorsReporter(errors <-chan error) {
	for err := range errors {
		panic(err)
	}
}
