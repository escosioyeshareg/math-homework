package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Generate a random integer between 1 and 100 (inclusive)
    var randomNumber int = rand.Intn(100) + 1

    fmt.Println("Random number:", randomNumber)
}
