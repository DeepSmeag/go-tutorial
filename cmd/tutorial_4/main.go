package main

import "fmt"

func main() {
	var intArr [3]int32
	fmt.Println(intArr)      // print the entire thing
	fmt.Println(intArr[0])   // print single elem
	fmt.Println(intArr[1:3]) // print a slice of the array
	fmt.Println(&intArr[0])  // print the memory address of the first element
	fmt.Println(&intArr)     //! would expect address of array right? but no
	fmt.Println(len(intArr)) // print the length of the array
	fmt.Println(cap(intArr)) // print the capacity of the array
	var arrPtr *[3]int32 = &intArr
	fmt.Println(arrPtr) //! still can't show address of array; can only say address of 1st element is that of the array
	// var arr2 [3]int32 = [3]int32{1, 2, 3} // to assign directly
	// arr2 = append(arr2, 4) doesn't work because arr2 is an array, not a slice
	var slice []int32 = []int32{1, 2, 3}
	slice = append(slice, 4)
	fmt.Println(slice)

	var intSlice []int32 = []int32{4, 5, 6}
	var intSlice2 []int32 = []int32{7, 8, 9}
	intSlice = append(intSlice, intSlice2...) // spread operator (in JS/TS it's before the object name)
	fmt.Println(intSlice)
	var intSlice3 []int32 = make([]int32, 4, 6)
	fmt.Println(len(intSlice3))
	fmt.Println(cap(intSlice3))
	intSlice3 = append(intSlice3, 1)
	fmt.Println(len(intSlice3))
	fmt.Println(cap(intSlice3))

	//! Maps
	var myMap map[string]int32 = map[string]int32{"one": 1, "two": 2} // key is in brackets; valye is after that
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Not existing"]) // will retrieve default of value type
	var age, ok = myMap2["Sarah"]
	fmt.Println(age, ok) // ok to check if key exists
	delete(myMap2, "Adam")
	fmt.Println(myMap2)
	myMap2["Adam"] = 25
	fmt.Println(myMap2)

	for key, value := range myMap2 {
		fmt.Println(key, value)
	}
	// a technical while loop
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	// an infinite loop
	for {
		fmt.Println("infinite")
		i++
		if i == 15 {
			break
		}
	}
	// traditional for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i) // i is a local variable; doesn't affect the outside i
	}
	fmt.Printf("i is %d after loop\n", i)
	// for j := 0; j < 10; j++ {
	// 	fmt.Println(j) ; don't need to use var
	// }

}
