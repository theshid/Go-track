package main

import (
	"fmt"
	"space/space_age"
	"strconv"
)

func main() {
	// Println function is used to
	// display output in the next line
	fmt.Println("Enter seconds: ")

	// var then variable name then variable type
	var seconds string

	// Taking input from user
	fmt.Scanln(&seconds)
	fmt.Println("Enter name of planet: ")
	var planet string
	fmt.Scanln(&planet)

	number, err := strconv.ParseFloat(seconds, 64)
	if err != nil {
		return 0, err
	}
	fmt.Printf("your age on planet %s is:%0.2f", planet, space_age.Age(planet, number))

	// Print function is used to
	// display output in the same line
	fmt.Print("Your Full Name is: ")
}
