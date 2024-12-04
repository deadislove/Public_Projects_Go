package main

import (
	"Password_Generator/services"
	"fmt"
	"log"
)

func main() {
	// Load the policy configuration
	policy, err := services.LoadPolicy("config.json")
	if err != nil {
		log.Fatalf("Failed to load policy: %v", err)
	}

	// Initialize the database
	db := services.InitializeDB("data/passwords.db")
	defer db.Close()

	// Retrieve the recent passwords
	recentPasswords, err := services.GetRecentPasswords(db)
	if err != nil {
		log.Fatalf("Failed to retrieve recent passwords: %v", err)
	}

	// Generate the password ensuring it's unique
	password, err := services.GeneratePassword(policy, recentPasswords)
	if err != nil {
		log.Fatalf("Error generating password: %v", err)
	}

	// Save the new password in the database
	services.SavePassword(db, password)

	fmt.Printf("Generated Password: %s\n", password)
}
