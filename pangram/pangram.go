// Package pangram contains utility methods to check whether a string is a pangram
package pangram

import (
	"strings"
)

const testVersion = 2

func isAlphabet(r rune) bool {
	c := int(r)
	if (c >= int('a') && c <= int('z')) ||
		(c >= int('A') && c <= int('Z')) {
		return true
	}
	return false
}

// IsPangram return whether a string is a pangram
func IsPangram(s string) bool {
	bv := make([]bool, 26)
	for _, c := range strings.ToLower(s) {
		if isAlphabet(c) {
			bv[int(c)-int('a')] = true
		}
	}
	for _, bit := range bv {
		if bit == false {
			return false
		}
	}
	return true
}
