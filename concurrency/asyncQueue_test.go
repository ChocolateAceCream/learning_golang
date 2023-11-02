package Concurrency

import (
	"testing"
)

func BenchmarkServe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MyFunction()
	}
}
