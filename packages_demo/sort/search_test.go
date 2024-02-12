package sort_demo

import (
	"reflect"
	"testing"
)

func TestSortSearchDemo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortSearchDemo()
		})
	}
}

func TestSortByKeyDemo(t *testing.T) {
	tests := []struct {
		name   string
		input  []Person
		output []Person
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			input: []Person{
				{Age: 25, Name: "25"},
				{Age: 22, Name: "22"},
				{Age: 30, Name: "30"},
			},
			output: []Person{
				{Age: 22, Name: "22"},
				{Age: 25, Name: "25"},
				{Age: 30, Name: "30"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortByKeyDemo(tt.input)
			if !reflect.DeepEqual(tt.input, tt.output) {
				t.Errorf("Expected %v, but got %v", tt.output, tt.input)
			}
		})
	}
}
