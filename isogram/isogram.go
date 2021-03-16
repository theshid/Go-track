package isogram

import (
	"strings"
)

//IsIsogram is a function that determine if a given string have duplicate characters excluding hyphen and empty space
func IsIsogram(word string) bool {

	wordLowerCase := strings.ToLower(word)
	stringMap := make(map[rune]bool)

	for _, value := range wordLowerCase {

		if value == '-' || value == ' ' {
			continue
		}
		 if stringMap[value]{
		 	return false
		 }
		 stringMap[value]=true

	}

	return true
}



