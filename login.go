package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(cookiesFilePath string, credentialsFilePath string) {
	errorOutput := log.New(os.Stderr, "", 0)

	if len(cookiesFilePath) == 0 || len(credentialsFilePath) == 0 {
		errorOutput.Println("Empty file paths for credentials or cookie files")
		return
	}

	credentialsFile, err := os.Open(credentialsFilePath)
	if err != nil {
		errorOutput.Println("Credentials file " + credentialsFilePath + " cannot be read")
		return
	}

	var credentials *Credentials
	json.NewDecoder(credentialsFile).Decode(&credentials)

	if len(credentials.Username) == 0 || len(credentials.Password) == 0 {
		errorOutput.Println("Credentials are empty")
		return
	}

	scraper := twitterscraper.New()

	err = scraper.Login(credentials.Username, credentials.Password)
	if err != nil || !scraper.IsLoggedIn() {
		errorOutput.Println("Cannot login with credentials from file " + credentialsFilePath)
		return
	}

	cookiesJSON, err := json.Marshal(scraper.GetCookies())
	if err != nil {
		errorOutput.Println("Cannot serialize cookies to JSON")
		return
	}

	cookiesFile, err := os.Create(cookiesFilePath)
	if err != nil {
		errorOutput.Println("Cannot create file with cookies")
		return
	}

	_, err = cookiesFile.Write(cookiesJSON)
	if err != nil {
		errorOutput.Println("Cannot write file with cookies")
		return
	}

	fmt.Println("Logged in and saved cookies to file " + cookiesFilePath)
}
