package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Ahoy, world!"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestV3(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := V3(); got != want {
		t.Errorf("V3() = %q, want %q", got, want)
	}
}
