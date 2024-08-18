package main

import (
    "fmt"

    "example.com/greetings"
)

func main() {
    message := greetings.Hello("Itanú")
    fmt.Println(message)
}