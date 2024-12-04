package services

import (
	"errors"
	"math/rand"
	"time"
)

// GeneratePassword ensures at least one uppercase and one lowercase letter
func GeneratePassword(policy Policy, recentPasswords []string) (string, error) {
	var charSet string
	var uppercaseSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var lowercaseSet = "abcdefghijklmnopqrstuvwxyz"

	// Build character set based on policy
	if policy.IncludeLowercase {
		charSet += lowercaseSet
	}
	if policy.IncludeUppercase {
		charSet += uppercaseSet
	}
	if policy.IncludeNumbers {
		charSet += "0123456789"
	}
	if policy.IncludeSpecial {
		charSet += policy.SpecialChars
	}

	// Validate character set
	if charSet == "" {
		return "", errors.New("character set is empty; check your policy settings")
	}

	// Ensure the random seed is unique for each run
	rand.Seed(time.Now().UnixNano())

	// Ensure the password length is sufficient
	if policy.PasswordLength < 2 {
		return "", errors.New("password length must be at least 2")
	}

	// Generate password with mandatory characters
	password := make([]byte, policy.PasswordLength)
	password[0] = uppercaseSet[rand.Intn(len(uppercaseSet))] // Ensure at least one uppercase
	password[1] = lowercaseSet[rand.Intn(len(lowercaseSet))] // Ensure at least one lowercase

	// Fill the rest of the password
	for i := 2; i < policy.PasswordLength; i++ {
		password[i] = charSet[rand.Intn(len(charSet))]
	}

	// Shuffle the password to avoid fixed positions for mandatory characters
	rand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	// Convert to string and ensure uniqueness
	passwordStr := string(password)
	if contains(recentPasswords, passwordStr) {
		return GeneratePassword(policy, recentPasswords) // Recursive retry if duplicate
	}

	return passwordStr, nil
}

// Helper function to check if a password exists in the recent list
func contains(passwords []string, newPass string) bool {
	for _, p := range passwords {
		if p == newPass {
			return true
		}
	}
	return false
}
