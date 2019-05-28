package luhn

import "strings"

// Valid determines whether or not a number is valid per the Luhn formula.
func Valid(str string) bool {
	if len(str) <= 1 {
		return false
	}

	str = strings.TrimSpace(str)

	return true
}
