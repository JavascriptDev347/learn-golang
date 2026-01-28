package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	fmt.Println("incrementing pointer 1: ", counter.hits)
	counter.hits += 1
	fmt.Println("incrementing pointer 2: ", counter.hits)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func main() {
	counter := Counter{
		hits: 0,
	}

	hello := "Hello "
	world := "World"
	replace(&hello, world, &counter)

	phrase := []string{hello, world}
	fmt.Println("P1: ", phrase)

	replace(&phrase[1], "Go", &counter)
	fmt.Println("P2: ", phrase)
}
