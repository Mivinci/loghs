package loghs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	filenameFormat = "%s.%s%s"
	timeFormat     = "200601021504"
)

// File file
// Path: a/b.c
// base: b
// ext:  .c
type File struct {
	Path    string
	MaxSize int64

	dir  string
	base string
	ext  string

	size int64
	file *os.File
	mu   sync.Mutex
}

// NewFile new
func NewFile(path string, maxSize int64) *File {
	dir := filepath.Dir(path)
	ext := filepath.Ext(path)
	if ext == "" {
		ext = ".log"
	}
	base := strings.TrimSuffix(filepath.Base(path), ext)
	return &File{
		Path:    path,
		MaxSize: maxSize,
		base:    base,
		ext:     ext,
		dir:     dir,
	}
}

func (f *File) Write(p []byte) (n int, err error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.write(p)
}

func (f *File) write(p []byte) (n int, err error) {
	plen := int64(len(p))
	if plen > f.MaxSize {
		return 0, fmt.Errorf("write length %d exceeds maximum file size %d", plen, f.MaxSize)
	}
	if f.file == nil {
		err = f.open()
		return
	}
	if plen+f.size > f.MaxSize {
		if err = f.rotate(); err != nil {
			return
		}
	}
	n, err = f.file.Write(p)
	n, err = f.file.WriteString("\n")
	f.size += plen
	return
}

// Rotate rotates
func (f *File) Rotate() error {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.rotate()
}

func (f *File) rotate() error {
	if err := f.close(); err != nil {
		return err
	}
	if err := f.backup(); err != nil {
		return err
	}
	if err := f.open(); err != nil {
		return err
	}
	return nil
}

func (f *File) backup() error {
	if err := os.Rename(f.filename(), f.backupFilename()); err != nil {
		return fmt.Errorf("rename file error(%v)", err)
	}
	return nil
}

func (f *File) open() error {
	file, err := os.OpenFile(f.filename(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("open new file error(%v)", err)
	}
	f.file = file
	f.size = 0
	return nil
}

func (f *File) close() error {
	if f.file == nil {
		return nil
	}
	err := f.file.Close()
	f.file = nil
	return err
}

func (f *File) backupFilename() string {
	return fmt.Sprintf(filenameFormat, filepath.Join(f.dir, f.base), time.Now().Format(timeFormat), f.ext)
}

func (f *File) filename() string {
	return filepath.Join(f.dir, f.base) + f.ext
}
