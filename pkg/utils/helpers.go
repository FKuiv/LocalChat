package utils

import "github.com/FKuiv/LocalChat/pkg/models"

// Helper function to check if a slice contains a string
func SliceContainsStr(slice []string, str string) bool {
	for _, value := range slice {
		if value == str {
			return true
		}
	}
	return false
}

// Helper function to check if a slice of Users contains a User with a certain ID
func ContainsUser(users []*models.User, id string) bool {
	for _, user := range users {
		if user.ID == id {
			return true
		}
	}
	return false
}
