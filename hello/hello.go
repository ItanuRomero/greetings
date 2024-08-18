package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    message, error := greetings.Hello("Itanú")
    if error != nil {
        log.Fatal(error)
    }
    
    fmt.Println(message, "\n")

    names := []string{"Gladys", "Itanú", "Romero"}

    messages, err := greetings.Hellos(names)

    if err != nil {
        log.Fatal(err)
    }

    for index, message := range messages {
        fmt.Println(index, message)
    }
}