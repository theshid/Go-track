package isogram

import (
	"strings"
)

//IsIsogram is a function that determine if a given string have duplicate characters excluding hyphen and empty space
func IsIsogram(word string) bool {
	wordInRune := []rune(strings.ToLower(word))
	hyphenAndEmptySpace := []rune("- ")
	stringMap := make(map[rune]int)
	for i := 0; i < len(word); i++ {

		if wordInRune[i] == hyphenAndEmptySpace[0] || wordInRune[i] == hyphenAndEmptySpace[1] {
			continue
		}

		stringMap[wordInRune[i]]++
		if stringMap[wordInRune[i]] > 1 {
			return false
		}

	}

	return true
}
