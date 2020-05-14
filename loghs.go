package loghs

import (
	"fmt"
	"os"
)

const depth = 2

var h = New(os.Stdout)

// Info writes a info-level log to the stdout
func Info(s string) {
	h.Info().Time("").Caller(depth).Message(s)
}

// Error writes a error-level log to the stdout
func Error(s string) {
	h.Error().Time("").Caller(depth).Message(s)
}

// Fatal writes a fatal-level log to the stdout then exit with code 1
func Fatal(s string) {
	h.Fatal().Time("").Caller(depth).Message(s)
	os.Exit(1)
}

// Infof writes a info-level log to the stdout
func Infof(format string, args ...interface{}) {
	Info(fmt.Sprintf(format, args...))
}

// Errorf writes a error-level log to the stdout
func Errorf(format string, args ...interface{}) {
	Info(fmt.Sprintf(format, args...))
}

// Fatalf writes a fatal-level log to the stdout then exit with code 1
func Fatalf(format string, args ...interface{}) {
	Fatal(fmt.Sprintf(format, args...))
	os.Exit(1)
}
