package main

import (
	"fmt"
	"log"

	"github.com/mi11km/playground/pkg/greetings"

	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Mike")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello, world!", message)
	fmt.Println(quote.Go())
}
