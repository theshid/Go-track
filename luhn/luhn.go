package luhn

import (
	"strings"
)

// Valid check if the given string satifies the luhn algorithm
func Valid(input string) bool {
	trimInput := strings.ReplaceAll(input, " ", "")
	count := 0
	if len(trimInput) <= 1 {
		return false
	}

	for i := len(trimInput) - 1; i >= 0; i-- {
		value:= trimInput[i]
		if value >= '0' && value <='9' {
			digit := int(value - '0')
			if (i % 2) == (len(trimInput) % 2) {

				if digit > 4 {
					count += (digit * 2) - 9
					continue
				}
				count += digit * 2
				continue
			}
			count += digit
		} else{
			return false
		}

	}

	return count%10 == 0

}