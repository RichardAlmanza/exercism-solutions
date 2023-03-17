// twofer is used to share the cookie
package twofer

import "fmt"

// ShareWith returns a string telling with whom I'm sharing.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
