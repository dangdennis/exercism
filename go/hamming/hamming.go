package hamming

import "errors"

// Distance returns the Hamming distance between two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Different lengths")
	}

	diff := 0
	for i := range a {
		if a[i] != b[i] {
			diff++
		}
	}

	return diff, nil
}
