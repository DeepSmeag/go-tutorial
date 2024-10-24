package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("Value of p points to: %v\n", *p)
	fmt.Printf("Value of i: %v\n", i)
	p = &i
	*p = 10
	fmt.Printf("Value of p points to: %v\n", *p)
	fmt.Printf("Value of i: %v\n", i)

	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("Memory location of thing1 arr is %p\n", &thing1)
	var result [5]float64 = square(thing1)
	fmt.Println(thing1)
	fmt.Println(result)

	var resultSlice []float64 = squareSlice(thing1[:])
	// fmt.Printf("Memory location of resultSlice arr is %p\n", &resultSlice)
	fmt.Println(thing1)
	fmt.Println(resultSlice)
	// so it's different, which means we have copied the slice
	var slice1 []float64 = []float64{1, 2, 3, 4, 5}
	fmt.Printf("Memory location of slice1 arr is %p\n", slice1)
	var resultSlice2 []float64 = squareSlice(slice1)
	// fmt.Printf("Memory location of resultSlice2 arr is %p\n", &resultSlice2)
	fmt.Println(slice1)
	fmt.Println(resultSlice2)
	// and even if we pass the slice itself and not a new object, we get this; so it's pass by value
	var map1 = map[int]float64{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
	fmt.Printf("Memory location of map1 arr is %p\n", map1)
	var resultMap map[int]float64 = squareMap(map1)
	// fmt.Printf("Memory location of resultMap arr is %p\n", &resultMap)
	fmt.Println(map1)
	fmt.Println(resultMap)

	// pass by pointer
	fmt.Printf("Memory location of thing1 arr is %p\n", &thing1)
	result = squarePointer(&thing1)
	// fmt.Printf("Memory location of result arr is %p\n", &result)
	fmt.Println(thing1)
	fmt.Println(result)

}

func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("Memory location of thing2 arr is %p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}
func squareSlice(thing2 []float64) []float64 {
	fmt.Printf("Memory location of thing2 arr is %p\n", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}
func squareMap(thing2 map[int]float64) map[int]float64 {
	fmt.Printf("Memory location of thing2 arr is %p\n", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}
func squarePointer(thing2 *[5]float64) [5]float64 {
	fmt.Printf("Memory location of thing2 arr is %p\n", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
