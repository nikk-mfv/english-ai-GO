package config

import (
	"os"

	"encoding/json"
	"englishAI/entities"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewGoogleOauthConfig() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}
}

// User represents a user authenticated via Google OAuth2.
func FetchGoogleUserInfo(token *oauth2.Token) (entities.User, error) {
	req, _ := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return entities.User{}, err
	}
	defer resp.Body.Close()

	var data struct {
		ID       string `json:"id"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		ImageUrl string `json:"picture"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return entities.User{}, err
	}
	return entities.User{
		Provider:      "google",
		ProviderID:    data.ID,
		Email:         data.Email,
		Username:      data.Name,
		EmailImageUrl: data.ImageUrl,
	}, nil
}
