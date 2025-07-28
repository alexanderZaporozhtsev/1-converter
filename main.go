package main

import "fmt"

func main() {
	const USDToEuro float64 = 0.8525
	const USDToRub float64 = 79.40

	amount := getUserInput()

	eurosToRubles := amount / USDToEuro * USDToRub

	fmt.Print(eurosToRubles)
}

func getUserInput() float64 {
	var amount float64
	fmt.Scan(&amount)

	return amount
}

func convert(amount float64, convertFrom string, convertTo string) {}
