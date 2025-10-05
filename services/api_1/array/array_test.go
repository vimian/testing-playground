package array

import (
	"testing"
)

func TestFindMin(t *testing.T) {
	tests := []struct {
		name      string
		input     []float64
		output    float64
		expectErr bool
	}{
		{"empty slice", []float64{}, 0, true},
		{"single element", []float64{42}, 41, false},
		{"multiple elements", []float64{42, 13, 7, 99}, 7, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FindMin(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err != nil)
			}
			if err == nil && result != tt.output {
				t.Errorf("expected %v, got %v", tt.output, result)
			}
		})
	}
}
