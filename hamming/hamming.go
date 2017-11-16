// Package hamming provides utilites to calculate hamming codes
package hamming

import "fmt"

const testVersion = 6

// Distance calculates hamming distance between 2 strings.
// return error if strings are of unequal length
func Distance(a, b string) (int, error) {
	dist := 0
	if len(a) != len(b) {
		return -1, fmt.Errorf("strings should be of equal length")
	}
	for i := range a {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist, nil
}
