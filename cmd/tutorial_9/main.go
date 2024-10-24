package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main() {
	// var c = make(chan int) // chan for channel, then the type it holds
	// <- special syntax to add data
	// c <- 1
	// var i int = <-c // we extracted data from the channel
	// fmt.Println(i)

	// CHANNEL vs BUFFERED CHANNEL

	// var c = make(chan int)
	// go process(c)
	// for i := range c {
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second * 1)
	// }

	// c = make(chan int, 5)
	// go process(c)
	// for i := range c {
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second * 1)
	// }

	// ----------------------
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go CheckTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func process(c chan int) {
	defer close(c) // do this unless you want a deadlock
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Exiting process")
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}
func CheckTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofuPrice = rand.Float32() * 20
		if tofuPrice <= MAX_TOFU_PRICE {
			tofuChannel <- website
			break
		}
	}
}
func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("The chicken is available at %s\n", website)
	case website := <-tofuChannel:
		fmt.Printf("The tofu is available at %s\n", website)

	}
}
