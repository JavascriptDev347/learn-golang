package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)

	shoppingList["eggs"] = 14
	shoppingList["coffee"] = 12
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 1

	fmt.Println("MAP: ", shoppingList)

	delete(shoppingList, "coffee")
	fmt.Println("MAP 2: ", shoppingList)

	cereal, found := shoppingList["cereal"]
	if !found {
		fmt.Println("Cereal not found")
	} else {
		fmt.Println("Cereal: ", cereal)
	}

}
