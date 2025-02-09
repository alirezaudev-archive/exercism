package twofer

import "fmt"

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	if name == "Do-yun" || name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
