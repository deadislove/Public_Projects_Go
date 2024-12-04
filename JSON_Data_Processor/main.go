package main

import (
	"JSON_Data_Processor/services"
	"flag"
	"fmt"
	"log"
	"os"
)

func printHelp() {
	// Print the help message
	fmt.Println("Usage: go run main.go -path <json file path> [-output <output file path>]")
	fmt.Println("\nFlags:")
	fmt.Println("  -path <path>   Path to the JSON file to process")
	fmt.Println("  -output <path> Optional path to save the pretty-printed JSON output (default: output.json)")
	fmt.Println("  -h, -help      Show this help message")
}

func main() {
	// Define the -path flag to get the JSON file path from the user
	filePath := flag.String("path", "", "Path to the JSON file")
	outputPath := flag.String("output", "output.json", "Optional path to save the pretty-printed JSON output (default: output.json)")
	helpFlag := flag.Bool("help", false, "Show help message")

	// Parse the flags
	flag.Parse()

	// Display the help message if -h or -help is provided
	if *helpFlag {
		printHelp()
		//flag.Usage()
		os.Exit(0) // Exit after showing help
	}

	// Ensure the user provides a file path
	if *filePath == "" {
		log.Fatal("Please provide a JSON file path using the -path flag")
	}

	// Read and process the JSON file
	jsonData, err := services.ReadJSONFile(*filePath)

	if err != nil {
		log.Fatal("Error reading JSON file: %v", err)
	}

	// Print the JSON data in a beautiful format
	prettyJSON := services.PrettyPrintJSON(jsonData)

	// If the user provided an output path, save the JSON to that file
	if *outputPath != "" {
		err := os.WriteFile(*outputPath, []byte(prettyJSON), 0644)
		if err != nil {
			log.Fatalf("Error writing to output file: %v", err)
		}
		fmt.Printf("Output saved to %s\n", *outputPath)
	} else {
		// Otherwise, print the JSON to the console
		fmt.Println(prettyJSON)
	}
}
