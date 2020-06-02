package loghs

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilepath(t *testing.T) {
	path := "a/b/c.log"
	base := "c"
	ext := ".log"
	assert.Equal(t, ext, filepath.Ext(path))
	assert.Equal(t, path, filepath.Join("a", "b", base)+ext)
}
