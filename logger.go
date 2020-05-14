package loghs

import "io"

const (
	_level   = "level"
	_message = "message"
)

var (
	r Render
)

// Logger logger
type Logger struct {
	out   io.Writer
	level Level
}

// New new
func New(w io.Writer) Logger {
	return Logger{out: w, level: NoLevel}
}

func (l *Logger) output(level Level, message string, fields ...Field) {
	r.Render(l.out, append(fields, String(_level, level.String()), String(_message, message)))
}

// Info info
func (l *Logger) Info(message string, fields ...Field) {
	l.output(InfoLevel, message, fields...)
}
