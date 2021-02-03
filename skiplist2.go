package main

import (
	"fmt"
	"github.com/jdxyw/skiplist-go"
)

type IntCmp struct {}

func (IntCmp) Compare(rhs, lhs interface{}) int {
	rhsint := rhs.(int)
	lhsint := lhs.(int)

	switch result := rhsint-lhsint; {
	case result == 0:
		return 0
	case result > 0:
		return 1
	default:
		return -1
	}
}

func (IntCmp) Name() string {
	return "Int64Comparator"
}

func main() {
	// We implement a Int64 Comaparator and pass it.
	// This skiplist would be use int64 as the key/value type.
	s := skiplist.NewSkiplist(10, IntCmp{})

	// Use the Set to insert/update element in this list.
	// The `value` could be nil.
	s.Set(111, 123)
	s.Set(222, 234)
	s.Set(333, 345)
	s.Set(444, 456)
	s.Set(555, 567)


	fmt.Printf("The length of this skiplist is %v.\n", s.Len())

	// Get value from a skiplist.
	val, _ := s.Get(111)
	fmt.Printf("The value of the key 111 in this skiplist is %v.\n", val.(int))

	if _, err := s.Get(666); err != nil {
		fmt.Printf("The key 666 is not in this skiplist is.\n")
	}

	if s.Contains(444) == true {
		fmt.Println("The key 444 is exist in this skiplist.")
	}

	// Remove one key from the skiplist
	s.Delete(444)
	if s.Contains(444) == false {
		fmt.Println("The key 444 has been removed from this skiplist.")
	}
}
