package diffsquares

//SquareOfSum compute the square of sum of the first n natural numbers
func SquareOfSum(i int) int {
	ret := (i * (i + 1)) / 2
	return ret * ret
}

//SumOfSquares compute the sum of squares of the first n natural numbers
func SumOfSquares(i int) int {
	return (i * (i + 1) * (2*i + 1)) / 6
}

//Difference compute the difference between the square of sum and the sum of squares of the first n natural numbers
func Difference(i int) int {
	return SquareOfSum(i) - SumOfSquares(i)
}
