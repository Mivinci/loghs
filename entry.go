package loghs

import (
	"fmt"
	"io"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var (
	enc = plainEncoder{}

	mu sync.RWMutex

	pool = &sync.Pool{
		New: func() interface{} {
			return &Entry{
				buf: make([]byte, 0, 512), // half KB
			}
		},
	}
)

// Entry Entry
type Entry struct {
	level Level
	out   io.Writer
	buf   []byte
}

func newEntry(out io.Writer, level Level) *Entry {
	e := pool.Get().(*Entry)
	e.out = out
	e.level = level
	e.buf = e.buf[:0]
	if level != NoLevel {
		e.buf = enc.String(e.buf, level.String())
	}
	return e
}

// Message message
func (e *Entry) Message(s string) {
	if e == nil {
		return
	}
	// append message string
	e.buf = enc.Suffix(e.buf, s)
	e.write()
}

func (e *Entry) write() {
	if _, err := e.out.Write(e.buf); err != nil {
		fmt.Fprintln(e.out, "loghs: could not write a log")
	}
	pool.Put(e)
}

// Time time
func (e *Entry) Time(format string) *Entry {
	e.buf = enc.Time(e.buf, time.Now(), format)
	return e
}

// TimeUnix timestamp microseconds
func (e *Entry) TimeUnix() *Entry {
	e.buf = enc.Int64(e.buf, time.Now().Unix())
	return e
}

// Caller caller
func (e *Entry) Caller(depth int) *Entry {
	_, file, line, ok := runtime.Caller(depth)
	if !ok {
		return e
	}
	// Seems like I can't make this faster
	e.buf = enc.Suffix(e.buf, file+":"+strconv.Itoa(line))
	return e
}

// String string
func (e *Entry) String(s string) *Entry {
	e.buf = enc.Suffix(e.buf, s)
	return e
}

// Int64 int64
func (e *Entry) Int64(i int64) *Entry {
	e.buf = enc.Int64(e.buf, i)
	return e
}
