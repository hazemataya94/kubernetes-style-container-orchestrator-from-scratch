package main

import (
	"fmt"
	"golang/modules/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Hazem")
	fmt.Println(message)
}
