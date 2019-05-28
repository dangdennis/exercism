package diffsquares

// Difference returns the difference between SquareOfSum and SumOfSquare
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

// SquareOfSum returns the square of sums
func SquareOfSum(n int) int {
	num := 0
	for i := 1; i < n+1; i++ {
		num += i
	}
	return num * num

	// The mathematical O(1) way:
	// r := n * (n + 1) / 2
	// return r * r
}

// SumOfSquares returns the sum of squares
func SumOfSquares(n int) int {
	num := 0
	for i := 1; i < n+1; i++ {
		num += i * i
	}
	return num

	// The mathematical O(1) way:
	// return n * (n + 1) * (2*n + 1) / 6
}
