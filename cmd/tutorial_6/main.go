package main

import (
	"fmt"
	"unsafe"
)

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner // or just owner and it is named as the type name as well in this case
	int             // this makes a field called int of type int
}
type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}
type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}
func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) bool {
	if e.milesLeft() >= miles {
		return true
	} else {
		return false
	}
}

func main() {
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 10, ownerInfo: owner{"Alex"}, int: 10} // or we can omit field names and they'll be populated in order
	myEngine.mpg = 20                                                                           // we can also directly access like this
	fmt.Println(myEngine)

	// inline struct
	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{25, 15} // we can also populate inline
	fmt.Println(myEngine2)
	fmt.Println(unsafe.Sizeof(struct {
		a int32
		b int32
		c int32
		int8
	}{1, 1, 1, 1})) // 16 bytes; the int32+int8 pair is being padded to reach 8 bytes which is the words of my machine (64-bit system)
	//! methods - functions related to a struct
	fmt.Printf("Total miles left %v\n", myEngine.milesLeft())
}
