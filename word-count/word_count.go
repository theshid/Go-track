package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(input string)Frequency{
	f := Frequency{}
	result:= strings.FieldsFunc(strings.ToLower(input), func(r rune) bool {
		return r != '\'' && unicode.IsPunct(r) || unicode.IsSymbol(r) || unicode.IsSpace(r)
	})

	for _, v := range result{
		if v[0]== '\'' {
			v = v[1:]
		}
		if v[len(v)-1] == '\''{
			v= v[:len(v)-1]
		}
		f[v]++
	}
	return f
}
