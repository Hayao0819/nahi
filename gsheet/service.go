package gsheet

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"os"

	"github.com/cockroachdb/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) (*http.Client, error) {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := getTokenFromFile(tokFile)
	if err != nil {
		tok, err = getTokenFromWeb(config)
		if err != nil {
			// log.Fatalf("Unable to retrieve token from web: %v", err)
			return nil, errors.Wrapf(err, "Unable to retrieve token from web: %v")
		}
	}
	saveToken(tokFile, tok)

	return config.Client(context.Background(), tok), nil
}

func getAuthCodeFromRedirectedURL(u string) (string, error) {
	parsed, err := url.Parse(u)
	if err != nil {
		return "", errors.Wrapf(err, "Unable to parse URL: %v")
	}
	return parsed.Query().Get("code"), nil
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Println("URL: ", authURL)
	fmt.Print("Go to this link in your browser then type the reirected URL: ")

	var stdinUrl string
	if _, err := fmt.Scan(&stdinUrl); err != nil {
		return nil, errors.Wrapf(err, "Unable to read authorization code: %v")
	}

	authCode, err := getAuthCodeFromRedirectedURL(stdinUrl)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to get authorization code: %v")
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to retrieve token from web: %v")
	}
	return tok, nil
}

// Retrieves a token from a local file.
func getTokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	slog.Debug("Saving credential", "file", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

var service *sheets.Service = nil

func GetService() (*sheets.Service, error) {
	if service != nil {
		return service, nil
	}
	return NewService("./assets/credentials.json")
}

func NewService(credentialFile string) (*sheets.Service, error) {
	ctx := context.Background()
	b, err := os.ReadFile(credentialFile)
	if err != nil {
		// log.Fatalf("Unable to read client secret file: %v", err)
		return nil, errors.Wrapf(err, "Unable to read client secret file: %v")
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		// log.Fatalf("Unable to parse client secret file to config: %v", err)
		return nil, errors.Wrapf(err, "Unable to parse client secret file to config: %v")
	}
	client, err := getClient(config)
	if err != nil {
		// log.Fatalf("Unable to get client: %v", err)
		return nil, errors.Wrapf(err, "Unable to get client: %v")

	}

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		// log.Fatalf("Unable to retrieve Sheets client: %v", err)
		return nil, errors.Wrapf(err, "Unable to retrieve Sheets client: %v")
	}
	return srv, nil
}
