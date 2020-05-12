package loghs

import (
	"testing"
)

type nopOutput struct{}

func (w *nopOutput) Write(b []byte) (n int, err error) {
	// return the actual length in order to `AddPrinter(...)` to be work with io.MultiWriter
	return len(b), nil
}

// IsNop defines this wrriter as a nop writer.
func (w *nopOutput) IsNop() bool {
	return true
}

func BenchmarkInfo(b *testing.B) {
	log := New(&nopOutput{})
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		log.Info("hhh")
	}
}
