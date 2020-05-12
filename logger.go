package loghs

import (
	"io"
	"io/ioutil"
)

// Time formats
const (
	CallerSkip      = 2
	TimeFormat      = "2006/01/02 15:04:05"
	TimeMicroFormat = "2006/01/02 15:04:05.999999"
)

// Logger is log handler
type Logger struct {
	out   io.Writer
	level Level
	entry *Entry
}

// New creates a logger with given output writer
func New(w io.Writer) Logger {
	if w == nil {
		w = ioutil.Discard
	}
	return Logger{out: w, level: NoLevel}
}

func (l *Logger) newEntry(level Level) *Entry {
	e := newEntry(l.out, level)
	l.entry = e
	if level != NoLevel {
		enc.String(&e.buf, level.String())
	}
	return e
}

// Log sends a log with no level
func (l Logger) Log() *Entry {
	return l.newEntry(NoLevel)
}

// With returns entry
func (l Logger) With() *Entry {
	return l.entry
}

// Info sends a info level log to writer
func (l *Logger) Info() *Entry {
	return l.newEntry(InfoLevel)
}
