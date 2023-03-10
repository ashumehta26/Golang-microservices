package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	names := []string{"Ashu", "Sam", "John"}
	// Get a greeting message and print it.
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
