package main

func double(x int) int {
	return x * x
}

func add(a, b int) int {
	return a + b
}

func greet() {
	println("Hello World")
}

func main() {
	greet()
	dozen := double(5)
	println("Dozen is ", dozen)

	bakersDozen := add(dozen, 10)
	println("BakersDozen is ", bakersDozen)

	anotherBakersDozen := add(double(6), dozen)
	println("anotherBakersDozen is ", anotherBakersDozen)
}
