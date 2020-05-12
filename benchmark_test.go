package loghs

import (
	"testing"
)

const (
	fakeMessage = "Test logging, but use a somewhat realistic message length."
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
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Info().Msg(fakeMessage)
		}
	})
}

func BenchmarkLogEmpty(b *testing.B) {
	log := New(nil)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Log().Msg(fakeMessage)
		}
	})
}
