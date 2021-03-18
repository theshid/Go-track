package diffsquares

import "math"

//SquareOfSum compute the square of sum of the first n natural numbers
func SquareOfSum(i int) int {
	inputToFloat := float64(i)
	sum := (inputToFloat / 2) * (inputToFloat + 1)
	return int(math.Pow(sum, 2))
}

//SumOfSquares compute the sum of squares of the first n natural numbers
func SumOfSquares(i int) int {
	inputToFloat := float64(i)
	return int((inputToFloat * (inputToFloat + 1) * (inputToFloat*2 + 1)) / 6)
}

//Difference compute the difference between the square of sum and the sum of squares of the first n natural numbers
func Difference(i int) int {
	return SquareOfSum(i) - SumOfSquares(i)
}
