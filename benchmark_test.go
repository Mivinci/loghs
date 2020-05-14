package loghs

import (
	"io/ioutil"
	"log"
	"testing"
)

const (
	fakeMessage = "Test logging, but use a somewhat realistic message length."
)

func BenchmarkEmpty(b *testing.B) {
	logger := New(ioutil.Discard)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Log().Message(fakeMessage)
		}
	})
}

func BenchmarkInfoTime(b *testing.B) {
	logger := New(ioutil.Discard)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().Time("").Message(fakeMessage)
		}
	})
}

func BenchmarkStdLogInfoTime(b *testing.B) {
	// r = &noRender{}
	logger := log.New(ioutil.Discard, "INFO ", log.LstdFlags)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Printf(fakeMessage)
		}
	})
}
