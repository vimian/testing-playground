package array

import "fmt"

// FindMin returns the minimum value in a slice of float64 numbers.
func FindMin(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("empty slice")
	}
	min := numbers[0]
	for _, n := range numbers[1:] {
		if n < min {
			min = n
		}
	}
	return min, nil
}
