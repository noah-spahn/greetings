package main

import (
	"fmt"
	"github.com/noah-spahn/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
