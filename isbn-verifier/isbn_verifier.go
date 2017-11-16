// Package isbn checks for a valid ISBN string
package isbn

import (
	"strings"
	"unicode"
)

// IsValidISBN checks for a valid ISBN ignoring dashes
func IsValidISBN(s string) bool {
	s = strings.Replace(s, "-", "", -1)
	s = strings.ToLower(s)
	if len(s) < 10 {
		return false
	}
	sum := 0
	for i := 0; i < len(s)-1; i++ {
		if !unicode.IsDigit(rune(s[i])) {
			return false
		}
		num := int(rune(s[i]) - '0')
		sum += num * (10 - i)
	}
	lastChar := s[len(s)-1]
	if rune(lastChar) == 'x' {
		sum += 10
	} else if !unicode.IsDigit(rune(lastChar)) {
		return false
	} else {
		sum += int(rune(lastChar) - '0')
	}

	if sum%11 != 0 {
		return false
	}
	return true
}
