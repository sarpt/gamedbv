package cmds

import (
	"fmt"
	"strings"
)

const (
	stringSeparator = ", "
)

// MultipleFlag can be used with std flag package to be used while setting process arguments
type MultipleFlag struct {
	values      []string
	validValues map[string]bool
}

// Strings returns string representation of values that were set
func (f *MultipleFlag) String() string {
	return strings.Join(f.values, stringSeparator)
}

// Valid returns whether the value is in valid form
// Always returns true if no valid values were provided
func (f MultipleFlag) Valid(value string) bool {
	if len(f.validValues) == 0 {
		return true
	}

	_, ok := f.validValues[value]

	return ok
}

// Values returns set values of the flag
func (f MultipleFlag) Values() []string {
	return f.values
}

// Set is used during setting of flags when used with flag package
// When the value is invalid, error is being returned
func (f *MultipleFlag) Set(value string) error {
	if !f.Valid(value) {
		return fmt.Errorf("value %s is not valid", value)
	}

	f.values = append(f.values, value)

	return nil
}
