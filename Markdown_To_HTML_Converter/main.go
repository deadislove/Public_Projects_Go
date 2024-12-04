package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/russross/blackfriday/v2"
)

func main() {
	// Define command-line flags for input and output file paths
	inputPath := flag.String("input", "", "Path to the input Markdown file")
	outputPath := flag.String("output", "output.html", "Path to the output HTML file (default: output.html)")

	// Parse the flags
	flag.Parse()

	// Check if the input file path is provided
	if *inputPath == "" {
		fmt.Println("Error: Input file path is required.")
		fmt.Println("Usage: go run main.go -input=\"C:\\path\\to\\input.md\" [-output=\"C:\\path\\to\\output.html\"]")
		os.Exit(1)
	}

	// Normalize paths for cross-platform compatibility
	inputFile := filepath.FromSlash(*inputPath)
	outputFile := filepath.FromSlash(*outputPath)

	// Read the Markdown file
	markdownContent, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", inputFile, err)
		os.Exit(1)
	}

	// Convert Markdown to HTML
	htmlContent := blackfriday.Run(markdownContent)

	// Write the HTML content to the output file
	err = ioutil.WriteFile(outputFile, htmlContent, 0644)
	if err != nil {
		fmt.Printf("Error writing file %s: %v\n", outputFile, err)
		os.Exit(1)
	}

	fmt.Printf("Conversion successful! HTML saved to %s\n", outputFile)
}
