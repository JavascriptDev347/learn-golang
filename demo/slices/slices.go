package main

import "fmt"

func printSlice(title string, slice []string) {
	fmt.Println("TITLE", title)

	for _, value := range slice {
		element := value
		fmt.Println(element)
	}

}
func main() {
	route := []string{"GET", "POST", "PUT", "DELETE"}

	route = append(route, "OPTIONS")
	printSlice("API", route)

	route = route[2:]
	printSlice("API from 2 element", route)
}
