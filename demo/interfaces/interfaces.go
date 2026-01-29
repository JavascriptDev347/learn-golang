package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("cook chicken")
}

func (s Salad) PrepareDish() {
	fmt.Println("toss salad")
	fmt.Println("add dressing")
}

func prepareDish(dishes []Preparer) {
	fmt.Println("dishes")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]

		fmt.Printf("dishes %v--\n", dish)
		dish.PrepareDish()
	}

	fmt.Println("last dishes")
}

func main() {
	dishes := []Preparer{Chicken("KFC"), Salad("OLIVIE")}
	prepareDish(dishes)
}
