package romannumerals

import "fmt"

func ToRomanNumeral(number int) (result string, err error) {
	if number < 1 || number > 3000 {
		return result, fmt.Errorf("number is out of range")
	}

	numbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < len(numbers); i++ {
		for number >= numbers[i] {
			number -= numbers[i]
			result += romans[i]
		}
	}
	return result,nil
}
