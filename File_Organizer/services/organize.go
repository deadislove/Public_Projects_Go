package services

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// OrganizeFile moves files into folders based on their extensions
func OrganizeFile(sourceDir, fileName string) {
	// get file extension
	ext := strings.ToLower(filepath.Ext(fileName))

	if ext == "" {
		ext = "unknown"
	} else {
		ext = ext[1:] //remove the dot (.)
	}

	//Create destination folder
	destDir := filepath.Join(sourceDir, ext)
	err := os.MkdirAll(destDir, os.ModePerm)

	if err != nil {
		fmt.Println("Error creating directory.", err)
		return
	}

	// Move file to the new fodler
	sourcePath := filepath.Join(sourceDir, fileName)
	destPath := filepath.Join(destDir, fileName)

	err = os.Rename(sourcePath, destPath)

	if err != nil {
		fmt.Println("Error moving file: ", err)
		return
	}

	fmt.Println("Moved %s to %s\n", fileName, destDir)

}
