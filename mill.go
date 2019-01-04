package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Printf("%d\n", time.Now().Nanosecond)
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().UnixNano() / int64(time.Millisecond))
}
