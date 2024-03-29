package Concurrency

import (
	"fmt"
	"testing"
)

func TestCoSum(t *testing.T) {
	tests := []struct {
		name    string
		total   int
		threads int
	}{
		// TODO: Add test cases.
		{
			name:    "multi thread 8",
			total:   8000000,
			threads: 8,
		},
		{
			name:    "multi thread 16",
			total:   8000000,
			threads: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CoSum(tt.total, tt.threads)
			expected := tt.total * (tt.total - 1) / 2
			fmt.Println("-----expected---", expected)
			if result != expected {
				t.Errorf("Expected %d, but got %d", expected, result)
			}
		})
	}
}

func Test_singleTreadSum(t *testing.T) {
	tests := []struct {
		name  string
		total int
	}{
		// TODO: Add test cases.
		{
			name:  "single thread",
			total: 800000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := tt.total * (tt.total - 1) / 2

			if got := singleTreadSum(tt.total); got != expected {
				t.Errorf("singleTreadSum() = %v, want %v", got, expected)
			}
		})
	}
}
