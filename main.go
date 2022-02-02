package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println("Call code in an external package " + quote.Go())
}
