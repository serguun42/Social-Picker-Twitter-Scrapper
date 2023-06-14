# Social Picker Twitter Scrapper

[twitter-scraper](https://github.com/n0madic/twitter-scraper) wrapped for [Social-Picker-API](https://github.com/serguun42/Social-Picker-API)

## Usage

1. [Install Go](https://go.dev/)
2. Create `credentials.json` somewhere – a JSON file, containing username and password. See an [`example/credentials.json`](./example/credentials.json)
3. Install dependencies – `go get .`
4. Build package for you platform – `go build -o Social-Picker-Twitter-Scrapper`
    - run `chmod +x Social-Picker-Twitter-Scrapper` if created binary is not executable
5. Login to retrieve this-app-specific cookies – `./Social-Picker-Twitter-Scrapper login cookies.json credentials.json` – where
    - `login` is a method for this app to get authentication
    - `cookies.json` is a path to persistent cookies file managed by this app
    - `credentials.json` is a path to credentials file you created at step #2
6. Check `cookies.json` (or your filepath) to exist and to be a JSON (_array or `{ "Name": "…", "Value": "…" }`_)
7. Run `./Social-Picker-Twitter-Scrapper getTweet cookies.json 123456789` – where
    - `getTweet` is a method for this app to get authentication
    - `cookies.json` is a path to persistent cookies file managed by this app
    - `123456789` is ID of a tweet you want to get
8. Outputs errors to `stderr`, fine tweet in SocialPost format to `stdout`

---

### [BSL-1.0 License](./LICENSE)
