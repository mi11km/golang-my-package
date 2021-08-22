package main

import (
	"fmt"

	"github.com/mi11km/playground/pkg/greetings"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, world!", greetings.Hello("Mike"))
	fmt.Println(quote.Go())
}
