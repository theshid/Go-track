package isogram

import (
	"strings"
	"unicode"
)

//IsIsogram is a function that determine if a given string have duplicate characters excluding hyphen and empty space
func IsIsogram(word string) bool {
	wordInRune := []rune(strings.ToLower(word))
	hyphen := []rune("-")
	emptySpace := []rune(" ")
	stringMap := make(map[rune]int)
	for i := 0; i < len(word); i++ {

		if unicode.ToLower(wordInRune[i]) == hyphen[0] || unicode.ToLower(wordInRune[i]) == emptySpace[0] {
			continue
		}

		stringMap[wordInRune[i]]++
		if stringMap[wordInRune[i]] > 1 {
			return false
		}

	}

	return true
}
