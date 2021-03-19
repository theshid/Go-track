package reverse

import "strings"

func Reverse(word string)string{

	if len(word) ==0 {
		return word
	}
	var reversedString strings.Builder
	var wordInRune  = []rune(word)
	reversedString.WriteRune(wordInRune[0])
	return Reverse(string(wordInRune[1:])) + reversedString.String()
}
