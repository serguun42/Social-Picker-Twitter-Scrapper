package main

import (
	"log"
	"os"
)

func main() {
	errorOutput := log.New(os.Stderr, "", 0)

	if len(os.Args) < 2 {
		errorOutput.Println("No method was passed as CLI argument")
		return
	}

	method := os.Args[1] // login | getTweet

	if method == "login" {
		if len(os.Args) < 4 {
			errorOutput.Println("No cookies file is present or no credentials file is present")
			return
		}

		cookiesFilePath := os.Args[2]
		credentialsFilePath := os.Args[3]

		login(cookiesFilePath, credentialsFilePath)
	} else if method == "getTweet" {
		if len(os.Args) < 4 {
			errorOutput.Println("No cookies file is present or tweet id was not passed as an argument")
			return
		}

		cookiesFilePath := os.Args[2]
		tweetId := os.Args[3]

		getTweet(cookiesFilePath, tweetId)
	} else {
		errorOutput.Println("No such method (from CLI arguments): " + method)
	}
}
