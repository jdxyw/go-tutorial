package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start goruntine")

	go func() {
		defer wg.Done()
		for c := 0; c < 2; c++ {
			fmt.Println(c)
		}
	}()

	go func() {
		defer wg.Done()
		for c := 10; c < 13; c++ {
			fmt.Println(c)
		}
	}()

	fmt.Println("waiting for gorountune finish")
	wg.Wait()
}
