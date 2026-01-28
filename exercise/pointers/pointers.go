//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

type Item struct {
	Name string
	chip bool
}

func deactivate(item *Item) {
	item.chip = false
}

func checkout(products []Item) {
	for index, _ := range products {
		deactivate(&products[index])
	}
}

func main() {
	shirt := Item{"Shirt", true}
	pants := Item{"Pants", true}
	shoes := Item{"Shoes", true}
	hat := Item{"Hat", true}

	cart := []Item{shirt, pants, shoes, hat}
	fmt.Println("Initial items:", cart)

	deactivate(&cart[1])
	fmt.Println("After deactivating one:", cart)

	checkout(cart)
	fmt.Println("After checkout:", cart)
}
