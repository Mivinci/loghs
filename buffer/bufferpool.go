package buffer

var (
	pool = NewPool()
	// Get retrieves a buffer from the pool
	Get = pool.Get
)
