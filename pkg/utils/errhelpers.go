package utils

import (
	"fmt"
	"log"
	"net/http"
)

func IDCreationErr(err error, w http.ResponseWriter) bool {
	text := "Error creating ID"
	if err != nil {
		log.Println(text, err)
		http.Error(w, text, http.StatusInternalServerError)
		return true
	}

	return false
}

func CreationErr(err error, w http.ResponseWriter) bool {
	text := "Error in creating new data to database"
	if err != nil {
		log.Println(text, err)
		http.Error(w, text, http.StatusInternalServerError)
		return true
	}

	return false
}

func DecodingErr(err error, endpoint string, w http.ResponseWriter) bool {
	if err != nil {
		log.Printf("Error in %s: %d", endpoint, err)
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return true
	}

	return false
}

func ItemNotFound(err error, item string, w http.ResponseWriter) bool {
	if err != nil {
		http.Error(w, fmt.Sprintf("%s not found", item), http.StatusNotFound)
		return true
	}

	return false
}

func MuxVarsNotProvided(isOk bool, value string, itemName string, w http.ResponseWriter) bool {
	if !isOk || value == "" {
		http.Error(w, fmt.Sprintf("%s not provided", itemName), http.StatusBadRequest)
		return true
	}
	return false
}

func ItemFetchError(err error, item string, w http.ResponseWriter) bool {
	text := fmt.Sprintf("Error getting %s", item)
	if err != nil {
		log.Println(text, err)
		http.Error(w, text, http.StatusInternalServerError)
		return true
	}

	return false
}

func CookieError(err error, w http.ResponseWriter) bool {
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
		return true
	}

	return false
}
