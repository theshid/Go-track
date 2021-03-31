package luhn

import (
	"strings"
)

// Valid check if the given string satisfies the luhn algorithm
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	count := 0
	if len(input) <= 1 {
		return false
	}
	double := len(input)%2 == 0

	for _, r := range input {
		digit := int(r - '0')
		if digit < 0 || digit > 9 {
			return false
		}
		if double {
			digit = digit * 2
			if digit > 9 {
				digit -= 9
			}
		}
		count += digit
		double = !double

	}
	return count%10 == 0
}
