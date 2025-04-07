package logger

import (
	"fmt"
	"os"
)

// Logger defines logging methods that support formatted output and fatal exits.
// This allows custom logging behavior to be injected into components.
type Logger interface {
	Printf(format string, v ...any)
	Fatalf(format string, v ...any)
}

// defaultLogger is the standard implementation using fmt and os.
type defaultLogger struct{}

// New returns the default Logger implementation.
func New() Logger {
	return defaultLogger{}
}

// Printf writes a formatted message to standard output.
func (defaultLogger) Printf(format string, v ...any) {
	fmt.Printf(format+"\n", v...)
}

// Fatalf writes a formatted error message to standard error and exits.
func (defaultLogger) Fatalf(format string, v ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", v...)
	os.Exit(1)
}
