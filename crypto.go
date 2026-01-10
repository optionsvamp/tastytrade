package tastytrade

import (
	"encoding/json"
	"fmt"
)

// DestinationVenueSymbols represents destination venue symbol information for cryptocurrency trading.
type DestinationVenueSymbols struct {
	Id                   int    `json:"id"`                     // Destination venue ID
	Symbol               string `json:"symbol"`                 // Symbol at the destination venue
	DestinationVenue     string `json:"destination-venue"`      // Name of the destination venue
	MaxQuantityPrecision int    `json:"max-quantity-precision"` // Maximum quantity precision
	MaxPricePrecision    int    `json:"max-price-precision"`    // Maximum price precision
	Routable             bool   `json:"routable"`               // Whether the symbol is routable
}

// Cryptocurrency represents cryptocurrency instrument information.
// It contains details about the cryptocurrency including trading characteristics and venue information.
type Cryptocurrency struct {
	Id                      int                       `json:"id"`                        // Instrument ID
	Symbol                  string                    `json:"symbol"`                    // Cryptocurrency symbol
	InstrumentType          string                    `json:"instrument-type"`           // Type of instrument (e.g., "Cryptocurrency")
	ShortDescription        string                    `json:"short-description"`         // Short description
	Description             string                    `json:"description"`               // Full description
	IsClosingOnly           bool                      `json:"is-closing-only"`           // Whether only closing positions are allowed
	Active                  bool                      `json:"active"`                    // Whether the cryptocurrency is currently active
	TickSize                string                    `json:"tick-size"`                 // Tick size for trading
	StreamerSymbol          string                    `json:"streamer-symbol"`           // Symbol used for streaming quotes
	DestinationVenueSymbols []DestinationVenueSymbols `json:"destination-venue-symbols"` // Array of destination venue symbols
}

// ListCryptocurrenciesResult represents the response structure returned by ListCryptocurrencies.
// It contains a list of cryptocurrency instruments and context information.
type ListCryptocurrenciesResult struct {
	Data struct {
		Items []Cryptocurrency `json:"items"` // Array of cryptocurrency instruments
	} `json:"data"`
	Context string `json:"context"` // API context identifier
}

// GetCryptocurrencyResult represents the response structure returned by GetCryptocurrency.
// It contains detailed information about a specific cryptocurrency instrument.
type GetCryptocurrencyResult struct {
	Data    Cryptocurrency `json:"data"`    // Cryptocurrency data
	Context string         `json:"context"` // API context identifier
}

// ListCryptocurrencies retrieves a list of all available cryptocurrencies.
// Returns a ListCryptocurrenciesResult containing all cryptocurrency instruments
// with their trading characteristics and venue information.
func (api *TastytradeAPI) ListCryptocurrencies() (ListCryptocurrenciesResult, error) {
	url := fmt.Sprintf("%s/instruments/cryptocurrencies", api.host)
	data, err := api.fetchData(url)
	if err != nil {
		return ListCryptocurrenciesResult{}, err
	}

	var response ListCryptocurrenciesResult
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ListCryptocurrenciesResult{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return ListCryptocurrenciesResult{}, err
	}

	return response, nil
}

// GetCryptocurrency retrieves data for a specific cryptocurrency symbol.
// Returns a GetCryptocurrencyResult containing detailed information about the cryptocurrency
// including tick size, venue symbols, and trading status.
func (api *TastytradeAPI) GetCryptocurrency(symbol string) (GetCryptocurrencyResult, error) {
	url := fmt.Sprintf("%s/instruments/cryptocurrencies/%s", api.host, symbol)
	data, err := api.fetchData(url)
	if err != nil {
		return GetCryptocurrencyResult{}, err
	}

	var response GetCryptocurrencyResult
	jsonData, err := json.Marshal(data)
	if err != nil {
		return GetCryptocurrencyResult{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return GetCryptocurrencyResult{}, err
	}

	return response, nil
}
