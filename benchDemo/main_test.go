package main

import "testing"

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkFib10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}
