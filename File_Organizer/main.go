package main

import (
	"flag"
	"fmt"
	"os"

	"File_Organizer/services"
)

func main() {
	// Define a command-line flag for the folder path
	sourceDir := flag.String("path", ".", "Path to the folder to organize")
	flag.Parse()

	// Expand the path to handle Git Bash
	expandedPath, err := services.ExpandPath(*sourceDir)

	if err != nil {
		fmt.Println("Error expanding path: ", err)
		return
	}

	// Read directory contents
	files, err := os.ReadDir(expandedPath)

	if err != nil {
		fmt.Println("Error reading directory: ", err)
		return
	}

	// Process each file
	for _, file := range files {
		if !file.IsDir() {
			services.OrganizeFile(expandedPath, file.Name())
		}
	}
	fmt.Println("File organization copmlete!")
}
