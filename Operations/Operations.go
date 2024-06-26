package Operations

import (
	"errors"
)

func ExecuteOperation(a int, b int, sign string) (int, error) {
	var result int
	switch sign {
	case "+": result = add(a, b)
	case "-": result = sub(a, b)
	case "*", "x": result = multiply(a, b)
	case "/": result = divide(a, b)
	default: return result, errors.New(unknown(sign)) 
	}
	return result, nil
}

// Add inputs a and b together
func add(a int, b int) int {
	result := a + b
	return result
}

func sub(a int, b int) int {
	result := a - b
	return result
}

func multiply(a int, b int) int {
	result := a * b
	return result
}

func divide(a int, b int) int {
	result := a / b
	return result
}

func unknown(sign string) string {
	return "No operation known for " + sign
	 
}
