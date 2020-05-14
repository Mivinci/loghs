package loghs

import "io"

// Render entry
type Render interface {
	Render(io.Writer, []Field)
}
