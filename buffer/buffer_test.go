package buffer

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkBuffers(b *testing.B) {
	str := strings.Repeat("a", 1024)
	slice := make([]byte, 1024)
	buf := bytes.NewBuffer(slice)
	custom := NewPool().Get()
	b.Run("BytesBuffer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			buf.WriteString(str)
			buf.Reset()
		}
	})
	b.Run("CustomBuffer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			custom.AppendString(str)
			custom.Reset()
		}
	})
}
