package isogram

import (
	"strings"
)

//IsIsogram is a function that determine if a given string have duplicate characters excluding hyphen and empty space
func IsIsogram(word string) bool {

	replacer := strings.NewReplacer("-", "", " ", "")
	trimWord := replacer.Replace(strings.ToLower(word))

	wordInRune := []rune(trimWord)
	stringMap := make(map[rune]int)

	for i := 0; i < len(trimWord); i++ {
		stringMap[wordInRune[i]]++
		if stringMap[wordInRune[i]] > 1 {
			return false
		}

	}

	return true
}
