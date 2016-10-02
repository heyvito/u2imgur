package remote

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/victorgama/u2imgur/config"
)

var client = &http.Client{}

func GetTokenFromPin(pin string) (*config.Session, error) {
	resp, err := http.PostForm("https://api.imgur.com/oauth2/token", url.Values{
		"client_id":     {config.EnvironData.ID},
		"client_secret": {config.EnvironData.Secret},
		"grant_type":    {"pin"},
		"pin":           {pin},
	})

	if err != nil {
		return nil, err
	}

	return getSessionFromResponse(resp)
}

func GetTokenFromRefreshToken(token string) (*config.Session, error) {
	resp, err := http.PostForm("https://api.imgur.com/oauth2/token", url.Values{
		"client_id":     {config.EnvironData.ID},
		"client_secret": {config.EnvironData.Secret},
		"grant_type":    {"refresh_token"},
		"refresh_token": {token},
	})

	if err != nil {
		return nil, err
	}

	return getSessionFromResponse(resp)
}

func getSessionFromResponse(resp *http.Response) (*config.Session, error) {
	defer resp.Body.Close()
	result := make(map[string]interface{})
	err := json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return nil, err
	}

	t := time.Now().Add(time.Duration(result["expires_in"].(float64)) * time.Second)

	var sess = config.Session{
		AccessToken:  result["access_token"].(string),
		RefreshToken: result["refresh_token"].(string),
		ExpiresAt:    t,
	}

	return &sess, nil
}
