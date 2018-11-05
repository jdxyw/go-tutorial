package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(1)
	fmt.Println(rand.Intn(4))
	fmt.Println(rand.Intn(4))
	fmt.Println(rand.Intn(4))
}
