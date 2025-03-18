package main

import (
	"math/rand"
	"time"
)

func randInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min+1)
}

func main() {
	fmt.Println("The random number is:", randInt(1, 10))
}
