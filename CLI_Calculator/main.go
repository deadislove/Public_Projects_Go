package main

import (
	"cli-calculator/calculator" // Adjusted import path
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Define command line flags
	num1 := flag.String("n1", "", "First number")
	num2 := flag.String("n2", "", "Second number")
	op := flag.String("op", "", "Operation (add, sub, mul, div)")
	flag.Parse()

	// Ensure numbers are provided
	if *num1 == "" || *num2 == "" {
		fmt.Println("Usage: -n1 <number1> -n2 <number2> -op <operation>")
		os.Exit(1)
	}

	// Convert inputs to floats
	n1, err1 := strconv.ParseFloat(*num1, 64)
	n2, err2 := strconv.ParseFloat(*num2, 64)

	if err1 != nil || err2 != nil {
		log.Fatalf("Invalid number format. Please provide valid numbers.")
	}

	// Create a new Calculator instance
	calc := calculator.Calculator{
		Num1: n1,
		Num2: n2,
	}

	// Perform the operation
	switch *op {
	case "add":
		fmt.Printf("Result: %f\n", calc.Add())
	case "sub":
		fmt.Printf("Result: %f\n", calc.Subtract())
	case "mul":
		fmt.Printf("Result: %f\n", calc.Multiply())
	case "div":
		result, err := calc.Divide()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		fmt.Printf("Result: %f\n", result)
	default:
		log.Fatalf("Invalid operation. Use one of: add, sub, mul, div.")
	}

}
