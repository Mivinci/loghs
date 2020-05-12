package loghs

import (
	"fmt"
	"io"
	"os"
	"sync"
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

// caller .
type caller struct {
	pc      uintptr
	file    string
	line    int
	defined bool
}

// newCaller new a caller
func newCaller(pc uintptr, file string, line int, ok bool) caller {
	return caller{pc, file, line, ok}
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
		enc.AppendString(&e.buf, msg)
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
