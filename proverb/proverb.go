// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"
const sentenceOne = "For want of a %s the %s was lost."
const sentenceLast = "And all for the want of a %s."
// Function that generate a proverb providing the input
func Proverb(rhyme []string) []string {

	var proverb []string

		for i := range rhyme {
			if i<len(rhyme)-1 {
				proverb = append(proverb,fmt.Sprintf(sentenceOne,rhyme[i],rhyme[i+1]) )
			} else {
				proverb = append(proverb,fmt.Sprintf(sentenceLast,rhyme[0]))
			}

		}


	return proverb
}
