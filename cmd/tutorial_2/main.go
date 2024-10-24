package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// fmt.Println("Hello, World!")
	var intNum int16 = 2
	fmt.Println(intNum)

	var floatNum float32 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString string = "Hello, World!γ"
	fmt.Println(myString)

	//! len(string) gives no. of bytes, not of characters...
	fmt.Println(len(myString))
	//! So strings are encoded using UTF-8; when characters are outside the original ASCII set, they take more than one byte; so can't do string len like that
	//! To get no. of characters, use the following:
	fmt.Println(utf8.RuneCountInString(myString))

	//! Rune = character in Go
	var myRune rune = 'a'
	fmt.Println(myRune)         // prints a's ASCII code
	fmt.Println(string(myRune)) // prints 'a'

	var myBoolean bool = false
	println(myBoolean)

	//! Default values: 0 for all numeric types & runes; for strings "", for bools false

	//! can omit type (or include for clarity) if assigning right away
	var myVar = "text" // definitely a string
	fmt.Println(myVar)

	//! can also use := to declare and assign a variable
	myVar2 := "text2"
	fmt.Println(myVar2)
	// or we can chain
	myVar3, myVar4 := "text3", "text4"
	fmt.Println(myVar3, myVar4)
	//! Good practice to add the type when it's not obvious; I'd say add it anyway, doesn't hurt

	//! Constants
	const myConst string = "const value" // they also need to be initialized at declaration time
	fmt.Println(myConst)
}
