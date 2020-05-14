package loghs

import (
	"io"
	"os"
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

type noRender struct{}

func (r *noRender) Render(out io.Writer, fields []Field) {}

func BenchmarkFieldLog(b *testing.B) {
	r = &noRender{}
	log := New(os.Stdout)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Info("")
		}
	})
}

// func BenchmarkInfo(b *testing.B) {
// 	log := New(&nopOutput{})
// 	b.ResetTimer()
// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			log.Info().Msg(fakeMessage)
// 		}
// 	})
// }

// func BenchmarkLogEmpty(b *testing.B) {
// 	log := New(nil)
// 	b.ResetTimer()
// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			log.Log().Msg(fakeMessage)
// 		}
// 	})
// }
