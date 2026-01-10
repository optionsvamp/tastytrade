package tastytrade

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	baseURL = "https://api.tastytrade.com"
)

// TastytradeAPI represents the Tastytrade API client
type TastytradeAPI struct {
	httpClient *http.Client
	authToken  string
	host       string
	apiVersion string // API version for Accept-Version header (e.g., "20250715")
}

// NewTastytradeAPI creates a new instance of TastytradeAPI
func NewTastytradeAPI(hosts ...string) *TastytradeAPI {
	host := baseURL
	if len(hosts) > 0 {
		host = hosts[0]
	}
	return &TastytradeAPI{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		host:       host,
	}
}

// SetAPIVersion sets the API version for subsequent requests.
// If version is empty, the header will not be sent (defaults to 20250714).
// Example: api.SetAPIVersion("20250715")
func (api *TastytradeAPI) SetAPIVersion(version string) {
	api.apiVersion = version
}

// fetchData sends a GET request to the specified URL with authorization
func (api *TastytradeAPI) fetchData(url string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", api.authToken)

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return nil, fmt.Errorf("client error occurred: status code %d", resp.StatusCode)
	} else if resp.StatusCode >= 500 {
		return nil, fmt.Errorf("server error occurred: status code %d", resp.StatusCode)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

// Helper function to fetch and unmarshal data
func (api *TastytradeAPI) fetchDataAndUnmarshal(urlVal string, v interface{}) error {
	req, err := http.NewRequest("GET", urlVal, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", api.authToken)

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return fmt.Errorf("client error occurred: status code %d", resp.StatusCode)
	} else if resp.StatusCode >= 500 {
		return fmt.Errorf("server error occurred: status code %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}

// fetchInstrumentData sends a GET request to an instrument endpoint with authorization and Accept-Version header.
// This is used specifically for instrument endpoints that support versioning.
func (api *TastytradeAPI) fetchInstrumentData(url string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", api.authToken)
	if api.apiVersion != "" {
		req.Header.Set("Accept-Version", api.apiVersion)
	}

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return nil, fmt.Errorf("client error occurred: status code %d", resp.StatusCode)
	} else if resp.StatusCode >= 500 {
		return nil, fmt.Errorf("server error occurred: status code %d", resp.StatusCode)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
