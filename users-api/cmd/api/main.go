package main

import (
	"fmt"
)

const message = "Users API starting upt... 🚀"

func main() {
	response := fmt.Sprintf("Hello, World, %s", message)
	fmt.Println(response)
}
