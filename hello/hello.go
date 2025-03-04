package main

import (
    "fmt"
	"log"
	"github.com/theyvraj/go_start/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(1)
	names := []string{"Alice", "Bob", "Carol"}
	messages, err := greetings.Hellos(names)
	message, err := greetings.Hello("Devraj")
	if err != nil {
		log.Fatal(err)
	}    
    fmt.Println(message)
	fmt.Println(messages)
}