package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Sally", "Dorothy"}

	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for name, greeting := range message {
		fmt.Println(name, ":", greeting)
	}

	//for name, greeting := range message {
	//fmt.Println(name, ":", greeting)
	//}
}
