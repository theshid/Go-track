package isogram

import "unicode"

//IsIsogram is a function that determine if a given string have duplicate characters excluding hyphen and empty space
func IsIsogram(word string) bool {
	wordInRune := []rune(word)
	hyphen := []rune("-")
	emptySpace := []rune(" ")
	for i := 0; i < len(word); i++ {

		for j := i + 1; j < len(word); j++ {
			if unicode.ToLower(wordInRune[i]) == unicode.ToLower(wordInRune[j]) {
				if unicode.ToLower(wordInRune[i]) == hyphen[0] || unicode.ToLower(wordInRune[i]) == emptySpace[0] {
					continue
				}
				return false
			}
		}
	}
	return true
}
