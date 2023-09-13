package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
}

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v, Welcome", name)
	return message
}
