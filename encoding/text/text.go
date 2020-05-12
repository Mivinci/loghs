package text

import (
	"strconv"
	"time"
)

// Encoder empty struct
type Encoder struct{}

// String appends a string to the writer
func (Encoder) String(buf *[]byte, s string) {
	*buf = append(*buf, s...)
}

// Byte appends a simple byte to the writer
func (Encoder) Byte(buf *[]byte, b byte) {
	*buf = append(*buf, b)
}

// Int64 appends converted int64 to the writer
func (Encoder) Int64(buf *[]byte, i int64) {
	*buf = strconv.AppendInt(*buf, i, 10)
}

// Time appends formated time to the writer
func (e Encoder) Time(buf *[]byte, t time.Time, format string) {
	*buf = t.AppendFormat(*buf, format)
}

// TimeUnix appends timestamp to the writer
func (e Encoder) TimeUnix(buf *[]byte, t time.Time) {
	e.Int64(buf, t.Unix())
}

// Space appends space to the writer
func (Encoder) Space(buf *[]byte) {
	*buf = append(*buf, ' ')
}

// LineBreaker appends space to the writer
func (Encoder) LineBreaker(buf *[]byte) {
	*buf = append(*buf, '\n')
}
