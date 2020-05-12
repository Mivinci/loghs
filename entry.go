package loghs

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var entryPool = &sync.Pool{
	New: func() interface{} {
		return &Entry{
			buf: make([]byte, 0, 512), // 512 Bytes
		}
	},
}

func putEntry(e *Entry) {
	// 64 KB
	if cap(e.buf) > (1 << 16) {
		return
	}
	entryPool.Put(e)
}

// Entry is header of a log
type Entry struct {
	buf   []byte
	out   io.Writer
	level Level
}

func newEntry(w io.Writer, level Level) *Entry {
	e := entryPool.Get().(*Entry)
	e.buf = e.buf[:0]
	e.out = w
	e.level = level
	return e
}

// Msg sends a message to the writer
func (e *Entry) Msg(msg string) {
	if e == nil {
		return
	}
	e.msg(msg)
}

func (e *Entry) msg(msg string) {
	if msg != "" {
		enc.Space(&e.buf)
		enc.String(&e.buf, msg)
	}

	if err := e.write(); err != nil {
		fmt.Fprintf(os.Stderr, "loghs: counld not write entry: %v\n", err)
	}
}

func (e *Entry) write() (err error) {
	if e == nil {
		return
	}
	if e.out != nil {
		_, err = e.out.Write(e.buf) // TODO
	}
	putEntry(e)
	return
}

// Caller sends file:line of the caller to the writer
func (e *Entry) Caller(skip int) *Entry {
	if e == nil {
		return e
	}
	return e.caller(skip)
}

func (e *Entry) caller(skip int) *Entry {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return e
	}
	enc.Space(&e.buf)
	enc.String(&e.buf, file+":"+strconv.Itoa(line))
	return e
}

// Time sends formated time to writer
func (e *Entry) Time(format string) *Entry {
	if e == nil {
		return e
	}
	enc.Space(&e.buf)
	enc.Time(&e.buf, time.Now(), format)
	return e
}

// TimeUnix sends timestamp to writer
func (e *Entry) TimeUnix() *Entry {
	if e == nil {
		return e
	}
	enc.Space(&e.buf)
	enc.TimeUnix(&e.buf, time.Now())
	return e
}
