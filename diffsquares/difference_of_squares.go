package diffsquares

func SquareOfSums(n int) (result int) {

	result = sum(n)
	result *= result
	return

}

func SumOfSquares(n int) (result int) {
	if n != 0 {
		result = n*n + SumOfSquares(n-1)
	}
	return
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

func sum(n int) (result int) {
	if n != 0 {
		result = n + sum(n-1)
	}
	return
}
