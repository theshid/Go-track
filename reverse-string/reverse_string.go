package reverse

// String reverses the given string.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]

	}
	return string(runes)
}

//My Solution
/*func Reverse(word string)string{

	if len(word) ==0 {
		return word
	}
	var reversedString strings.Builder
	var wordInRune  = []rune(word)
	reversedString.WriteRune(wordInRune[0])
	return Reverse(string(wordInRune[1:])) + reversedString.String()
}*/

