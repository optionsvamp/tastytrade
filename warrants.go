package tastytrade

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Warrant struct {
	Symbol         string `json:"symbol"`
	InstrumentType string `json:"instrument-type"`
	ListedMarket   string `json:"listed-market"`
	Description    string `json:"description"`
	IsClosingOnly  bool   `json:"is-closing-only"`
	Active         bool   `json:"active"`
}

type ListWarrantsResult struct {
	Data struct {
		Items []Warrant `json:"items"`
	} `json:"data"`
	Context string `json:"context"`
}

type GetWarrantResult struct {
	Data    Warrant `json:"data"`
	Context string  `json:"context"`
}

// ListWarrants retrieves list of warrants
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

// GetWarrant retrieves a specific warrant
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
