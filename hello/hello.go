package main

import (
    "fmt"
	"github.com/theyvraj/go_start/greetings"
)

func main() {
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}