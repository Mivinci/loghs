package buffer

import (
	"strconv"
	"time"
)

const _size = 1024

// Buffer is buffer
type Buffer struct {
	bs []byte
}

// Bytes returns the bytes value
func (b *Buffer) Bytes() []byte {
	return b.bs
}

// AppendByte writes a byte to thr buffer
func (b *Buffer) AppendByte(v byte) {
	b.bs = append(b.bs, v)
}

// AppendString writes a string to the buffer.
func (b *Buffer) AppendString(s string) {
	b.bs = append(b.bs, s...)
}

// AppendInt appends an integer to the underlying buffer (assuming base 10).
func (b *Buffer) AppendInt(i int64) {
	b.bs = strconv.AppendInt(b.bs, i, 10)
}

// AppendTime appends the time formatted using the specified layout.
func (b *Buffer) AppendTime(t time.Time, layout string) {
	b.bs = t.AppendFormat(b.bs, layout)
}

// Reset .
func (b *Buffer) Reset() {
	b.bs = b.bs[:0]
}
