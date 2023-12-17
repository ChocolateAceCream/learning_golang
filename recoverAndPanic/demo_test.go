package recoverAndPanic

import "testing"

func TestDemo(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		// TODO: Add test cases.
		{name: "demoTest1", expected: "new error"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Demo()
			if err.Error() != tt.expected {
				t.Errorf("Expected %s, but got %v", tt.expected, err)
			}
		})
	}
}
