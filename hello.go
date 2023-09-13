package main

import (
	"fmt"

	twoone "learn.com/golearntwo"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	message := twoone.Hello("Gladys")
	fmt.Println(message)
}
