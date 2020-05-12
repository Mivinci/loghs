package loghs

import (
	"io"
)

// Flags
const (
	Color int8 = 1 << iota
	Caller

	Std = Color | Caller
)

// Time formats
const (
	timeFormat      = "2006/01/02 15:04:05"
	timeMicroFormat = "2006/01/02 15:04:05.999999"
)

// Logger is log handler
type Logger struct {
	out   io.Writer
	level Level
}

// New creates a logger with given output writer
func New(w io.Writer) Logger {
	return Logger{out: w}
}

// Info sends a info level log to writer
func (l *Logger) Info(msg string) {
	l.newEntry(InfoLevel).Msg(msg)
}

func (l *Logger) newEntry(level Level) *Entry {
	return newEntry(l.out, level)
}
