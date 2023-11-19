package utils

import (
	"log"
	"net/http"
)

func IDErr(err error, w http.ResponseWriter) {
	text := "Error creating ID"
	if err != nil {
		log.Println(text, err)
		http.Error(w, text, http.StatusInternalServerError)
		return
	}
}

func CreationErr(err error, w http.ResponseWriter) {
	text := "Error in creating new data to database"
	if err != nil {
		log.Println(text, err)
		http.Error(w, text, http.StatusInternalServerError)
		return
	}

}

func DecodingErr(err error, w http.ResponseWriter, endpoint string) {
	if err != nil {
		log.Println("Error in /register", err)
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
}
