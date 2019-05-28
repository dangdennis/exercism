package isogram

import "unicode"

// IsIsogram determines whether a word is an isogram or not
func IsIsogram(word string) bool {
	found := make(map[rune]bool)

	for _, r := range word {
		l := unicode.ToUpper(r)

		if l == ' ' || l == '-' {
			continue
		}

		if _, ok := found[l]; !ok {
			found[l] = true
		} else {
			return false
		}
	}

	return true
}
