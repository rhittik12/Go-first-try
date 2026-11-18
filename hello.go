package main

import (
	"fmt"
	"log"
	"os"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	name := "Gladys"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	message, err := greetings.Hello(name)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
