package main

import "testing"

func BenchmarkStringBuilder(b *testing.B) {
	StringBuilder()
}

func BenchmarkStringConcatenation(b *testing.B) {
	StringConcatenation()
}
