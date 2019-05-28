package raindrops

import "strconv"

// Convert transform a number into a string of Pling, Plang, or Plong, or a mix of them
func Convert(num int) string {
	isPling := false
	isPlang := false
	isPlong := false

	if num%3 == 0 {
		isPling = true
	}

	if num%5 == 0 {
		isPlang = true
	}

	if num%7 == 0 {
		isPlong = true
	}

	var res string

	if isPling {
		res += "Pling"
	}

	if isPlang {
		res += "Plang"
	}

	if isPlong {
		res += "Plong"
	}

	if res == "" {
		return strconv.Itoa(num)
	}

	return res
}
