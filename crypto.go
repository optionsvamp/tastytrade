package tastytrade

import (
	"encoding/json"
	"fmt"
)

type DestinationVenueSymbols struct {
	Id                   int    `json:"id"`
	Symbol               string `json:"symbol"`
	DestinationVenue     string `json:"destination-venue"`
	MaxQuantityPrecision int    `json:"max-quantity-precision"`
	MaxPricePrecision    int    `json:"max-price-precision"`
	Routable             bool   `json:"routable"`
}

type Cryptocurrency struct {
	Id                      int                       `json:"id"`
	Symbol                  string                    `json:"symbol"`
	InstrumentType          string                    `json:"instrument-type"`
	ShortDescription        string                    `json:"short-description"`
	Description             string                    `json:"description"`
	IsClosingOnly           bool                      `json:"is-closing-only"`
	Active                  bool                      `json:"active"`
	TickSize                string                    `json:"tick-size"`
	StreamerSymbol          string                    `json:"streamer-symbol"`
	DestinationVenueSymbols []DestinationVenueSymbols `json:"destination-venue-symbols"`
}

type ListCryptocurrenciesResult struct {
	Data struct {
		Items []Cryptocurrency `json:"items"`
	} `json:"data"`
	Context string `json:"context"`
}

type GetCryptocurrencyResult struct {
	Data    Cryptocurrency `json:"data"`
	Context string         `json:"context"`
}

// ListCryptocurrencies retrieves list of cryptocurrencies
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

// GetCryptocurrency retrieves a specific cryptocurrency
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
