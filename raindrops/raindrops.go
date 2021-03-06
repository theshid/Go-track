package raindrops

import "strconv"

const factor3 = "Pling"
const factor5 = "Plang"
const factor7 = "Plong"

//Functions that convert a number into raindrop sound
func Convert(number int) string {
	result := ""
	if number%3 != 0 && number%5 != 0 && number%7 != 0 {
		result = strconv.Itoa(number)
	}
	if number%3 == 0 {
		result += factor3
	}
	if number%5 == 0 {
		result += factor5
	}
	if number%7 == 0 {
		result += factor7
	}

	return result
}
