package luhn

import (
	"strings"
	"unicode"
)

// Valid check if the given string satisfies the luhn algorithm
func Valid(input string) bool {
	trimInput := []rune(strings.ReplaceAll(input, " ", ""))
	count := 0
	if len(trimInput) <= 1 {
		return false
	}
	isSecondDigit := false
	for i := len(trimInput) - 1; i >= 0; i-- {
		value := trimInput[i]
		if !unicode.IsDigit(value) {
			return false
		}
		digit := int(value - '0')
		if digit >= 0 && digit <= 9 {
			if isSecondDigit {
				digit = digit * 2
				if digit > 9 {
					digit -= 9
				}
			}
			count += digit
			isSecondDigit = !isSecondDigit
		}
	}
	return count%10 == 0
}
