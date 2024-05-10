package main

import "fmt"

func main() {
	var number1 int = 16
	var number2 int = 3
	var result, remainder int = intDivision(number1, number2)
	fmt.Printf("The result Of Integer Divison Is %v With Remainder %v \n",result, remainder)
}
func intDivision(numerotar int, denominator int) (int, int) {
	var result int = numerotar / denominator
	var remainder int = numerotar % denominator
	return result, remainder
}
