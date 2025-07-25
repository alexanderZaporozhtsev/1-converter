package main

import "fmt"

func main() {
	const USDToEuro float64 = 0.8525
	const USDToRub float64 = 79.40

	euros := 150.0

	eurosToRubles := euros / USDToEuro * USDToRub

	fmt.Print(eurosToRubles)
}
