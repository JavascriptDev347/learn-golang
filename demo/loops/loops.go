package main

func main() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
		println("Sum is ", sum)
	}

	if sum > 10 {
		sum -= 5
		println("Decrement sum is ", sum)
	}

}
