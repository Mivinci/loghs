package buffer

import "sync"

// Pool is buffer pool
type Pool struct {
	p *sync.Pool
}

// NewPool constructs a new Pool.
func NewPool() Pool {
	return Pool{p: &sync.Pool{
		New: func() interface{} {
			return &Buffer{bs: make([]byte, 0, _size)}
		},
	}}
}

// Get retrieves a Buffer from the pool, creating one if necessary.
func (p Pool) Get() *Buffer {
	buf := p.p.Get().(*Buffer)
	buf.Reset()
	return buf
}

func (p Pool) put(buf *Buffer) {
	p.p.Put(buf)
}
