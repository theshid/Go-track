package etl

import "strings"

func Transform(oldSystem map[int][]string) map[string]int {

	result := make(map[string]int)

	for key, value := range oldSystem{
		for j:=0;j<len(value) ; j++  {
			result[strings.ToLower(value[j])]=key
		}

	}
	return result
}