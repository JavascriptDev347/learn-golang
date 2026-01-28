//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	Name  string
	Price float64
	Count int
}

func printDisplay(products [3]Product) {
	var total float64
	var count int
	for _, product := range products {
		total += product.Price * float64(product.Count)
		count += product.Count
		fmt.Println("Product name is ", product.Name, " and Price is ", product.Price)
	}
	fmt.Println("Total is ", total)
	fmt.Println("Count is ", count)
}

func main() {
	products := [...]Product{
		{
			Name:  "Olma",
			Price: 100,
			Count: 3,
		},
		{
			Name:  "Anor",
			Price: 12,
			Count: 3,
		},
		{
			Name:  "Banan",
			Price: 10,
			Count: 3,
		},
	}
	printDisplay(products)
}
