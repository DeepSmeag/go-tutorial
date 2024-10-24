package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello, World!"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 3
	var result, remainder, err = intDivision(numerator, denominator)
	//! Pretty string with function i/o as well; I can't return 2 things and not extract both
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result is %v!", result)
	} else {
		fmt.Printf("The result is %v with remainder %v!\n", result, remainder) //! interesting choice, printf doesn't end the line
	}
	fmt.Println(6 & 1)

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result is %v!", result)
	default:
		fmt.Printf("The result is %v with remainder %v!\n", result, remainder)
	}

}
func printMe(printValue string) {
	fmt.Println(printValue)
}
func intDivision(numerator int, denominator int) (int, int, error) {
	//! Error types; any function that can return an error should return it as the last return value
	var err error // default is nil
	if denominator == 0 {
		err = errors.New("Division by zero!")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
