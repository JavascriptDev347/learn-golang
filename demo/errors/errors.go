package main

import (
	"errors"
	"fmt"
)

type Stuff struct {
	values []int
}

func (s *Stuff) Get(index int) (int, error) {

	if index >= len(s.values) {
		return 0, errors.New("no element at the index")
	}

	return s.values[index], nil
}

func main() {
	s := Stuff{values: []int{1, 2, 3}}
	value, err := s.Get(3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value:", value)
	}
}
