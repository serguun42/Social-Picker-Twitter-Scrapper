package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

type Media struct {
	Type        string `json:"type"`
	ExternalURL string `json:"externalUrl"`
	Original    string `json:"original"`
}

type SocialPost struct {
	Caption   string  `json:"caption"`
	Author    string  `json:"author"`
	AuthorURL string  `json:"authorURL"`
	PostURL   string  `json:"postURL"`
	Medias    []Media `json:"medias"`
}

func getTweet(cookiesFilePath string, tweetId string) {
	errorOutput := log.New(os.Stderr, "", 0)

	if len(cookiesFilePath) == 0 || len(tweetId) == 0 {
		errorOutput.Println("Empty file paths for cookie file or tweetId")
		return
	}

	scraper := twitterscraper.New()

	cookiesFile, err := os.Open(cookiesFilePath)
	if err != nil {
		errorOutput.Println("Cookies file " + cookiesFilePath + " cannot be read")
		return
	}

	var cookies []*http.Cookie
	json.NewDecoder(cookiesFile).Decode(&cookies)

	scraper.SetCookies(cookies)

	if !scraper.IsLoggedIn() {
		errorOutput.Println("Cannot login with cookies from file " + cookiesFilePath)
		return
	}

	tweet, err := scraper.GetTweet(tweetId)
	if err != nil {
		errorOutput.Println("Error with getting tweet:")
		errorOutput.Println(err)
		return
	}

	cookiesJSON, err := json.Marshal(scraper.GetCookies())
	if err != nil {
		errorOutput.Println("Cannot serialize cookies to JSON")
	} else {
		cookiesFile, err = os.Create(cookiesFilePath)
		if err != nil {
			errorOutput.Println("Cannot create file with cookies")
		} else {
			_, err = cookiesFile.Write(cookiesJSON)
			if err != nil {
				errorOutput.Println("Cannot write file with cookies")
			}
		}
	}

	var socialPost SocialPost

	socialPost.Caption = tweet.Text
	socialPost.Author = tweet.Username
	socialPost.AuthorURL = "https://twitter.com/" + tweet.Username
	socialPost.PostURL = tweet.PermanentURL
	socialPost.Medias = make([]Media, 0)

	for i := 0; i < len(tweet.Photos); i++ {
		socialPost.Medias = append(socialPost.Medias, Media{"photo", tweet.Photos[i].URL, tweet.Photos[i].URL + ":orig"})
	}

	for i := 0; i < len(tweet.Videos); i++ {
		socialPost.Medias = append(socialPost.Medias, Media{"video", tweet.Videos[i].URL, tweet.Videos[i].URL})
	}

	for i := 0; i < len(tweet.GIFs); i++ {
		socialPost.Medias = append(socialPost.Medias, Media{"gif", tweet.GIFs[i].URL, tweet.GIFs[i].URL})
	}

	data, err := json.Marshal(socialPost)
	if err != nil {
		errorOutput.Println("Error with JSON encoding of socialPost:")
		errorOutput.Println(err)
		return
	}

	fmt.Println(string(data))
}
