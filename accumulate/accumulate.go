// Package accumulate provides functions to apply a function to input array
package accumulate

const testVersion = 1

// Accumulate applies a function to an array of strings and returns result
func Accumulate(inp []string, f func(string) string) []string {
	var res []string
	for _, s := range inp {
		res = append(res, f(s))
	}
	return res
}
