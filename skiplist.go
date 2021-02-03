package main

import (
	"fmt"
	"github.com/jdxyw/skiplist-go"
)

func main() {
	// If you pass the nil to `cmp` parameter, which would use the default comparactor (Bytes wise).
	s := skiplist.NewSkiplist(10, nil)

	// Use the Set to insert/update element in this list.
	// The `value` could be nil.
	s.Set([]byte("Hello"), []byte("world"))
	s.Set([]byte("Python"), []byte("Perl"))
	s.Set([]byte("PHP"), []byte("C++"))
	s.Set([]byte("PyTorch"), []byte("Tensorflow"))
	s.Set([]byte("Java"), nil)

	fmt.Printf("The length of this skiplist is %v.\n", s.Len())

	// Get value from a skiplist.
	val, _ := s.Get([]byte("PyTorch"))
	fmt.Printf("The value of the key PyTorch in this skiplist is %v.\n", string(val.([]byte)))

	if _, err := s.Get([]byte("IBM")); err != nil {
		fmt.Printf("The key IBM is not in this skiplist is.\n")
	}

	if s.Contains([]byte("Hello")) == true {
		fmt.Println("The key Hello is exist in this skiplist.")
	}

	// Remove one key from the skiplist
	s.Delete([]byte("Hello"))
	if s.Contains([]byte("Hello")) == false {
		fmt.Println("The key Hello has been removed from this skiplist.")
	}
}
