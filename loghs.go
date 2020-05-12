package loghs

import "os"

// Default is default logger with stdout
var (
	Default    = New(os.Stdout)
	callerSkip = CallerSkip + 1
)

// Info writes a info level log to the stdout
func Info(s string) {
	Default.Info().Time(TimeFormat).Caller(callerSkip).Msg(s)
}
