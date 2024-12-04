package services

import (
	"encoding/json"
	"io/ioutil"
)

// Policy defines the structure for the password policy settings
type Policy struct {
	PasswordLength   int    `json:"password_length"`
	IncludeUppercase bool   `json:"include_uppercase"`
	IncludeLowercase bool   `json:"include_lowercase"`
	IncludeNumbers   bool   `json:"include_numbers"`
	IncludeSpecial   bool   `json:"include_special"`
	CustomCharRange  string `json:"custom_char_range"`
	SpecialChars     string `json:"special_characters"`
}

// LoadPolicy reads the JSON configuration file and returns the policy
func LoadPolicy(path string) (Policy, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return Policy{}, err
	}

	var policy Policy
	err = json.Unmarshal(data, &policy)
	return policy, err
}
