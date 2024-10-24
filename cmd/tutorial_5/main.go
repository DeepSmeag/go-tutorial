package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	// alternatively, cast to array of runes
	var myString2 = []rune("résumé")
	var indexed2 = myString2[1]
	fmt.Printf("%v, %T\n", indexed2, indexed2)
	fmt.Printf("%v, %T\n", myString2, myString2)

	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "p", "t", "i", "o", "n"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i] // this creates a new string each time; way inefficient
	}
	fmt.Printf("\n%v\n", catStr)

	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	fmt.Printf("\n%v\n", strBuilder.String())
}
