package main

import (
	"fmt"
	bf "github.com/jdxyw/bloomfilter-go"
)

func main() {
	b, _ := bf.NewBloomFilter(100000, 0.01)

	b.Set([]byte("Java"))
	b.Set([]byte("Python"))
	b.Set([]byte("Go"))
	b.Set([]byte("C++"))

	if b.Check([]byte("Python")) == true {
		fmt.Println("The Python is in this bloomfilter.")
	}
}
