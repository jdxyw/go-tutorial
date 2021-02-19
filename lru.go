package main

import (
	"fmt"
	"github.com/jdxyw/lru-go"
)

func main() {
	cache := lru.NewCache(100)
	cache.Add("Go", 1)
	val, _ := cache.Get("Go")

	fmt.Printf("The value for Go is %v.", val)
}
