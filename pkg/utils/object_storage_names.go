package utils

import "fmt"

func UserProfilePicName(userId string) string {
	return fmt.Sprintf("user/%s", userId)
}

func GroupProfilePicName(groupId string) string {
	return fmt.Sprintf("group/%s", groupId)
}
