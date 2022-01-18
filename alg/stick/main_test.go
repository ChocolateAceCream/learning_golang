package main

import "testing"

func BenchmarkSingleThread(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SingleThread()
	}
}

func BenchmarkMultiThread(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MultiThread()
	}
}
