package encode

import (
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(input string) string {
	result := ""

	for len(input) > 0 {
		letter := string(input[0])
		inputLen := len(input)
		input = strings.TrimLeft(input, letter)
		if n := inputLen - len(input); n > 1 {
			result += strconv.Itoa(n)
			result += letter
		} else {
			result += letter
		}
	}
	return result
}

func RunLengthDecode(input string) string {
	result := ""
	for len(input) > 0 {
		i := strings.IndexFunc(input, func(r rune) bool {
			return !unicode.IsDigit(r)
		})
		if i > 0 {
			number, _ := strconv.Atoi(input[:i])
			result +=strings.Repeat(string(input[i]),number)
			input = input[i+1:]
		}else{
			result+=string(input[i])
			input = input[i+1:]
		}
	}
	return result
}
