// Package raindrops converts a number to string depending on factors
package raindrops

import (
	"strconv"
)

const testVersion = 3

// Convert takes an integer as an input and returns a string based on
// factors 3, 5 and 7
// here naive string concat is more efficient as compared to bytes.Buffer
// as we are dealing with short strings. Refer - http://herman.asia/efficient-string-concatenation-in-go
func Convert(i int) string {
	res := ""
	if i%3 == 0 {
		res += "Pling"
	}
	if i%5 == 0 {
		res += "Plang"
	}
	if i%7 == 0 {
		res += "Plong"
	}
	if res == "" {
		return strconv.Itoa(i)
	}
	return res
}
