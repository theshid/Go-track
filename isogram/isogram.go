package isogram

import (
	"strings"
)

//IsIsogram is a function that determine if a given string have duplicate characters excluding hyphen and empty space
func IsIsogram(word string) bool {

	repeated := map[rune]bool{}

	for _, value := range strings.ToLower(word) {

		if value == '-' || value == ' ' {
			continue
		}
		 if repeated[value]{
		 	return false
		 }
		 repeated[value]=true

	}

	return true
}




/*func IsIsogram(word string) bool {

	replacer := strings.NewReplacer("-", "", " ", "")
	trimWord := replacer.Replace(strings.ToLower(word))

	for _, value := range trimWord {

		if strings.Count(trimWord, string(value)) > 1 {
			return false
		}

	}

	return true
}*/
