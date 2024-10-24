package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type T interface {
	int | int32 | int64 | float32 | float64
}
type contactInfo struct {
	Name  string
	Email string
}
type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

func main() {
	println(add(1, 2))
	println(add(int32(1), int32(2)))
	println(add(int64(1), int64(2)))
	println(add(float32(1), float32(2)))
	println(add(float64(1), float64(2)))

	println(sumSlice([]int{1, 2, 3}))
	println(sumSlice([]float32{1, 2, 3}))

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("cmd/tutorial_10/purchaseInfo.json")
	fmt.Printf("%+v\n", purchases)

}

func add[t T](a t, b t) t {
	return a + b
}

func sumSlice[T int | float32 | float64](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return []T{}
	}
	var loaded = []T{}
	json.Unmarshal(data, &loaded)
	return loaded
}
