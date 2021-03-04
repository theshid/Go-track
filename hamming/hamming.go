package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1,errors.New("The words are not of same lengh")
	}
	var difference = 0
	for i:=0; i<len(b);i++{
		if a[i] != b[i] {
			difference +=1
		}
	}
	return difference,nil
}
