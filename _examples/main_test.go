package main

import (
	"testing"

	"github.com/mivinci/loghs"
)

func BenchmarkInfoTime(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			loghs.Info("hello")
		}
	})
}
