package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Function to read json data from a file
func ReadJSONFile(filePath string) (interface{}, error) {
	// Read the JSON file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON into an interface{}
	var jsonData interface{}
	err = json.Unmarshal(data, &jsonData)

	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

// Function to print the JSON data in a beautiful format
func PrettyPrintJSON(data interface{}) string {
	// Marshal the data into pretty JSON with indentation
	prettyJSON, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		log.Fatal("Error formatting JSON: ", err)
		return ""
	}

	// Print the pretty JSON
	fmt.Println(string(prettyJSON))
	return string(prettyJSON)
}
