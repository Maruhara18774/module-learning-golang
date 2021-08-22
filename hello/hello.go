package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Jan", "Samantha", "Darrin"}

	message, err := greetings.Hello("Lin")
	if err != nil {
		log.Fatal(err)
	}

	message2, err2 := greetings.Hellos(names)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(message)
	for _, result := range message2 {
		fmt.Println(result)
	}

}
