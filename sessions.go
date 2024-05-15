package tastytrade

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Email       string `json:"email"`
	Username    string `json:"username"`
	ExternalID  string `json:"external-id"`
	IsConfirmed bool   `json:"is-confirmed"`
}

type AuthData struct {
	User         User   `json:"user"`
	SessionToken string `json:"session-token"`
}

type AuthResponse struct {
	Data    AuthData `json:"data"`
	Context string   `json:"context"`
}

// Authenticate authenticates the client with the Tastytrade API
func (api *TastytradeAPI) Authenticate(username, password string) error {
	authURL := fmt.Sprintf("%s/sessions", api.host)
	authData := map[string]string{
		"login":    username,
		"password": password,
	}
	authBody, err := json.Marshal(authData)
	if err != nil {
		return err
	}

	resp, err := api.httpClient.Post(authURL, "application/json", bytes.NewReader(authBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("authentication failed: status code %d", resp.StatusCode)
	}

	authResponse := AuthResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return err
	}

	api.authToken = authResponse.Data.SessionToken
	return nil
}
