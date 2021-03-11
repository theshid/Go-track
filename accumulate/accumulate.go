package accumulate
//Accumulate is a function that does operation on a collection of strings
func Accumulate(words []string, operand func(cond string) string) (result[]string) {

	result = make([]string, len(words))

	for i, letter := range words {
		result[i] = operand(letter)
	}

	return
}
