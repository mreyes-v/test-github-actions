package calculator

// BasicCalculator implements the Calculator interface
type BasicCalculator struct{}

// Add performs addition operation
func (bc *BasicCalculator) Add(a, b float64) float64 {
	return a + b
}

// Subtract performs subtraction operation
func (bc *BasicCalculator) Subtract(a, b float64) float64 {
	return a - b
}

// Multiply performs multiplication operation
func (bc *BasicCalculator) Multiply(a, b float64) float64 {
	return a * b
}
