package main

import (
	"fmt"
	"githubActionsTest/calculator"
)

func main() {
	basicCalculator := calculator.BasicCalculator{}

	a, b := 10.5, 5.5

	// Addition
	sum := basicCalculator.Add(a, b)
	fmt.Printf("Addition result: %.2f\n", sum)

	// Subtraction
	diff := basicCalculator.Subtract(a, b)
	fmt.Printf("Subtraction result: %.2f\n", diff)

	// Multiplication
	product := basicCalculator.Multiply(a, b)
	fmt.Printf("Multiplication result: %.2f\n", product)
}
