// Package twofer prints a message to share things
package twofer

import "fmt"

// ShareWith takes a name and prints a message. defaults to you if left blank
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
