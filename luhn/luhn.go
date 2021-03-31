package luhn

import (
	"strings"
)

// Valid check if the given string satisfies the luhn algorithm
func Valid(input string) bool {
	trimInput := strings.ReplaceAll(input, " ", "")
	count := 0
	if len(trimInput) <= 1 {
		return false
	}
	checkLengthParity := len(trimInput) % 2

	for i, r := range trimInput {
		digit := int(r - '0')
		if digit < 0 || digit > 9 {
			return false
		}
		if i%2 == checkLengthParity {
			digit = digit * 2
			if digit > 9 {
				digit -= 9
			}
		}
		count += digit

	}
	return count%10 == 0
}
