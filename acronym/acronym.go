// Package acronym has utility method to generate acronym from string
package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 3

// Abbreviate splits string based on ' ' or '-' and generate an acronym
func Abbreviate(s string) string {
	res, parts := "", strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '-'
	})
	for _, part := range parts {
		res += string(unicode.ToUpper(rune(part[0])))
	}
	return res
}
