package calculator

import (
	"errors"
)

// Calculator struct
type Calculator struct {
	Num1 float64
	Num2 float64
}

// Add method
func (c *Calculator) Add() float64 {
	return c.Num1 + c.Num2
}

// Subtract method
func (c *Calculator) Subtract() float64 {
	return c.Num1 - c.Num2
}

// Multiply method
func (c *Calculator) Multiply() float64 {
	return c.Num1 * c.Num2
}

// Divide method with error handling
func (c *Calculator) Divide() (float64, error) {
	if c.Num2 == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return c.Num1 / c.Num2, nil
}
