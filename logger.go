package loghs

import "io"

const (
	_level   = "level"
	_message = "message"
)

// Logger logger
type Logger struct {
	out io.Writer
}

// New new
func New(w io.Writer) Logger {
	return Logger{out: w}
}

// Log non-level log
func (l *Logger) Log() *Entry {
	return newEntry(l.out, NoLevel)
}

// Info info
func (l *Logger) Info() *Entry {
	return newEntry(l.out, InfoLevel)
}

// Debug debug
func (l *Logger) Debug() *Entry {
	return newEntry(l.out, DebugLevel)
}

// Warn warn
func (l *Logger) Warn() *Entry {
	return newEntry(l.out, WarnLevel)
}

// Error error
func (l *Logger) Error() *Entry {
	return newEntry(l.out, ErrorLevel)
}

// Fatal fatal
func (l *Logger) Fatal() *Entry {
	return newEntry(l.out, FatalLevel)
}
