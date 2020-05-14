package loghs

import (
	"strconv"
	"time"
)

type plainEncoder struct{}

var _ Encoder = plainEncoder{}

// String string
func (plainEncoder) String(buf []byte, s string) []byte {
	return append(buf, s...)
}

// Byte byte
func (plainEncoder) Byte(buf []byte, b byte) []byte {
	return append(buf, b)
}

// Int64 int64
func (plainEncoder) Int64(buf []byte, i int64) []byte {
	return strconv.AppendInt(append(buf, ' '), i, 10)
}

// Suffix prefix
func (plainEncoder) Suffix(buf []byte, s string) []byte {
	return append(append(buf, ' '), s...)
}

// Time time
func (plainEncoder) Time(buf []byte, t time.Time, format string) []byte {
	if format == "" {
		format = "2006/01/02 15:04:05"
	}
	return t.AppendFormat(append(buf, ' '), format)
}
