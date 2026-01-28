package main

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 9, 7, 8
	if quiz1 > quiz2 {
		println("quiz 1 is higher than quiz 2")
	} else if quiz1 < quiz2 {
		println("quiz 1 is lower than quiz 2")
	} else {
		println("quiz 2 and quiz 1 same score")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		println("Acceptable score")
	} else {
		println("Not acceptable score")
	}
}
