package anagram

import "strings"

func Detect(input string, candidates []string) []string {

	var result []string
	inputToLower := strings.ToLower(input)
	countInput := count(inputToLower)
	for _,r := range candidates {
		wordCandidate := strings.ToLower(r)
		if len(inputToLower) != len(wordCandidate) || inputToLower == wordCandidate {
			continue
		}

		if  countInput == count(wordCandidate){
         result = append(result,r)
		}


	}
	return result
}


func count(s string)string {
	check := make([]byte,26)

	for i:=0;i<len(s) ;i++  {
		letter:= s[i]

		if letter < 'a' || letter > 'z' {
			continue
		}
		check[letter - 'a']++
	}
	return string(check)
}
