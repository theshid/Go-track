package isogram

import (
	"strings"
)

//IsIsogram is a function that determine if a given string have duplicate characters excluding hyphen and empty space
func IsIsogram(word string) bool {

	replacer := strings.NewReplacer("-", "", " ", "")
	trimWord := replacer.Replace(strings.ToLower(word))

	for _, value := range trimWord {

		if strings.Count(trimWord, string(value)) > 1 {
			return false
		}

	}

	return true
}
