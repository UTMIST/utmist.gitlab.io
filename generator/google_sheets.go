package generator

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

type credentials struct {
	Installed installed `json:"installed"`
}

type installed struct {
	ClientID                string   `json:"client_id"`
	ProjectID               string   `json:"project_id"`
	AuthURI                 string   `json:"auth_uri"`
	TokenURI                string   `json:"token_uri"`
	AuthProviderX509CertURL string   `json:"auth_provider_x509_cert_url"`
	ClientSecret            string   `json:"client_secret"`
	RedirectURIs            []string `json:"redirect_uris"`
}

func getCredentials() ([]byte, error) {
	var exists bool
	clientID, exists := os.LookupEnv("CLIENT_ID")
	if !exists {
		return nil, nil
	}
	projectID, exists := os.LookupEnv("PROJECT_ID")
	if !exists {
		return nil, nil
	}
	authURI, exists := os.LookupEnv("AUTH_URI")
	if !exists {
		return nil, nil
	}
	tokenURI, exists := os.LookupEnv("TOKEN_URI")
	if !exists {
		return nil, nil
	}
	authProviderX509CertURL, exists := os.LookupEnv("AUTH_PROVIDER_X509_CERT_URL")
	if !exists {
		return nil, nil
	}
	clientSecret, exists := os.LookupEnv("CLIENT_SECRET")
	if !exists {
		return nil, nil
	}
	redirectURIs := []string{}
	if redirectURI, exists := os.LookupEnv("REDIRECT_URI1"); exists {
		redirectURIs = append(redirectURIs, redirectURI)
	}
	if redirectURI, exists := os.LookupEnv("REDIRECT_URI2"); exists {
		redirectURIs = append(redirectURIs, redirectURI)
	}

	creds := credentials{
		Installed: installed{
			ClientID:                clientID,
			ProjectID:               projectID,
			AuthURI:                 authURI,
			TokenURI:                tokenURI,
			AuthProviderX509CertURL: authProviderX509CertURL,
			ClientSecret:            clientSecret,
			RedirectURIs:            redirectURIs,
		},
	}

	return json.MarshalIndent(creds, "", "  ")
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok := tokenFromFile(tokFile)
	if tok == nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) *oauth2.Token {

	accessToken, exists := os.LookupEnv("ACCESS_TOKEN")
	if !exists {
		return nil
	}
	tokenType, exists := os.LookupEnv("TOKEN_TYPE")
	if !exists {
		return nil
	}
	refreshToken, exists := os.LookupEnv("REFRESH_TOKEN")
	if !exists {
		return nil
	}
	expiryStr, exists := os.LookupEnv("EXPIRY")
	if !exists {
		return nil
	}
	expiry, err := time.Parse(time.RFC3339, expiryStr)
	if err != nil {
		return nil
	}

	tok := &oauth2.Token{
		AccessToken:  accessToken,
		TokenType:    tokenType,
		RefreshToken: refreshToken,
		Expiry:       expiry,
	}
	return tok
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
