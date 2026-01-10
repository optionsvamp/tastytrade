package tastytrade

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Warrant represents warrant instrument information.
// It contains details about the warrant including trading status and market information.
type Warrant struct {
	Symbol         string `json:"symbol"`          // Warrant symbol
	InstrumentType string `json:"instrument-type"` // Type of instrument (e.g., "Warrant")
	ListedMarket   string `json:"listed-market"`   // Market where the warrant is listed
	Description    string `json:"description"`     // Description of the warrant
	IsClosingOnly  bool   `json:"is-closing-only"` // Whether only closing positions are allowed
	Active         bool   `json:"active"`          // Whether the warrant is currently active
}

// ListWarrantsResult represents the response structure returned by ListWarrants.
// It contains a list of warrant instruments and context information.
type ListWarrantsResult struct {
	Data struct {
		Items []Warrant `json:"items"` // Array of warrant instruments
	} `json:"data"`
	Context string `json:"context"` // API context identifier
}

// GetWarrantResult represents the response structure returned by GetWarrant.
// It contains detailed information about a specific warrant instrument.
type GetWarrantResult struct {
	Data    Warrant `json:"data"`    // Warrant data
	Context string  `json:"context"` // API context identifier
}

// ListWarrants retrieves a list of warrants, optionally filtered by symbols.
// If symbols are provided, only those warrants are returned. If no symbols are provided,
// all available warrants are returned.
// Returns a ListWarrantsResult containing matching warrant instruments.
func (api *TastytradeAPI) ListWarrants(symbols ...string) (ListWarrantsResult, error) {
	url := fmt.Sprintf("%s/instruments/warrants", api.host)
	if len(symbols) > 0 {
		url = fmt.Sprintf("%s?symbols=%s", url, strings.Join(symbols, ","))
	}
	data, err := api.fetchData(url)
	if err != nil {
		return ListWarrantsResult{}, err
	}

	var response ListWarrantsResult
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ListWarrantsResult{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return ListWarrantsResult{}, err
	}

	return response, nil
}

// GetWarrant retrieves data for a specific warrant symbol.
// Returns a GetWarrantResult containing detailed information about the warrant
// including market, description, and trading status.
func (api *TastytradeAPI) GetWarrant(symbol string) (GetWarrantResult, error) {
	url := fmt.Sprintf("%s/instruments/warrants/%s", api.host, symbol)
	data, err := api.fetchData(url)
	if err != nil {
		return GetWarrantResult{}, err
	}

	var response GetWarrantResult
	jsonData, err := json.Marshal(data)
	if err != nil {
		return GetWarrantResult{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return GetWarrantResult{}, err
	}

	return response, nil
}
