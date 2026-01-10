package tastytrade

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// User represents user information returned after authentication.
type User struct {
	Email       string `json:"email"`        // User's email address
	Username    string `json:"username"`     // Username
	ExternalID  string `json:"external-id"`  // External identifier
	IsConfirmed bool   `json:"is-confirmed"` // Whether the account is confirmed
}

// AuthData represents authentication data returned by the Authenticate endpoint.
// It contains user information and a session token for subsequent API requests.
type AuthData struct {
	User         User   `json:"user"`          // User information
	SessionToken string `json:"session-token"` // Session token for API authentication
}

// AuthResponse represents the response structure returned by Authenticate.
// It contains authentication data and context information.
type AuthResponse struct {
	Data    AuthData `json:"data"`    // Authentication data
	Context string   `json:"context"` // API context identifier
}

// Authenticate authenticates the client with the Tastytrade API using username and password.
// On success, the session token is stored in the API client for use in subsequent requests.
// Returns an error if authentication fails.
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
