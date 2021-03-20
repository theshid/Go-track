package pangram

import "strings"

func IsPangram(input string) bool {
	newWord := strings.ToLower(input)

	for l := 'a'; l <= 'z'; l++ {
		if strings.IndexRune(newWord, l) == -1 {
			return false
		}
	}
	return true
}
