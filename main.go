package main

import "log"

func main() {
	sum := calculator(1, "+", 5)
	min := calculator(1, "-", 5)
	zarb := calculator(1, "*", 5)
	taqsim := calculator(100, "/", 5)
	log.Println("sum=====> ", sum)
	log.Println("min=====> ", min)
	log.Println("zarb====> ", zarb)
	log.Println("taqsim==> ", taqsim)
}

func calculator(first int, amal string, second int) int {
	switch amal {
	case "+":
		return first + second
	case "-":
		return first - second
	case "*":
		return first * second
	case "/":
		return first / second
	default:
		return first + second
	}
}
