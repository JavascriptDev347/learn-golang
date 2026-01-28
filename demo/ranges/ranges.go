package main

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, element := range slice {
		println(index, element)
	}
}
