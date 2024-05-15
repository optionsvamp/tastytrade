package tastytrade

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type EquityResponse struct {
	Context string     `json:"context"`
	Data    EquityData `json:"data"`
}

type Tick struct {
	Threshold string `json:"threshold,omitempty"`
	Value     string `json:"value"`
}

type EquityData struct {
	Active                         bool   `json:"active"`
	BorrowRate                     string `json:"borrow-rate"`
	Cusip                          string `json:"cusip"`
	Description                    string `json:"description"`
	ID                             int    `json:"id"`
	InstrumentType                 string `json:"instrument-type"`
	IsClosingOnly                  bool   `json:"is-closing-only"`
	IsETF                          bool   `json:"is-etf"`
	IsFractionalQuantityEligible   bool   `json:"is-fractional-quantity-eligible"`
	IsIlliquid                     bool   `json:"is-illiquid"`
	IsIndex                        bool   `json:"is-index"`
	IsOptionsClosingOnly           bool   `json:"is-options-closing-only"`
	Lendability                    string `json:"lendability"`
	ListedMarket                   string `json:"listed-market"`
	MarketTimeInstrumentCollection string `json:"market-time-instrument-collection"`
	OptionTickSizes                []Tick `json:"option-tick-sizes"`
	ShortDescription               string `json:"short-description"`
	StreamerSymbol                 string `json:"streamer-symbol"`
	Symbol                         string `json:"symbol"`
	TickSizes                      []Tick `json:"tick-sizes"`
}

type EquityListResponse struct {
	Context string `json:"context"`
	Data    struct {
		Items []EquityData `json:"items"`
	} `json:"data"`
}

type EquityQueryParams struct {
	Symbol      []string `json:"symbol"`
	Lendability string   `json:"lendability"`
	IsIndex     *bool    `json:"is-index"`
	IsETF       *bool    `json:"is-etf"`
}

type ActiveEquityQueryParams struct {
	Lendability string `json:"lendability"`
	PerPage     int    `json:"per-page"`
	PageOffset  int    `json:"page-offset"`
}

// GetEquityData retrieves data for a specific equity symbol
func (api *TastytradeAPI) GetEquityData(symbol string) (EquityResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/equities/%s", api.host, symbol)
	data, err := api.fetchData(urlVal)
	if err != nil {
		return EquityResponse{}, err
	}

	var response EquityResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return EquityResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return EquityResponse{}, err
	}

	return response, nil
}

// ListEquities retrieves a list of all equities
func (api *TastytradeAPI) ListEquities(params *EquityQueryParams) (EquityListResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/equities", api.host)

	if params != nil {
		queryParams := url.Values{}
		for _, symbol := range params.Symbol {
			queryParams.Add("symbol[]", symbol)
		}
		if params.Lendability != "" {
			queryParams.Add("lendability", params.Lendability)
		}
		if params.IsIndex != nil {
			queryParams.Add("is-index", fmt.Sprintf("%t", *params.IsIndex))
		}
		if params.IsETF != nil {
			queryParams.Add("is-etf", fmt.Sprintf("%t", *params.IsETF))
		}
		urlVal = fmt.Sprintf("%s?%s", urlVal, queryParams.Encode())
	}

	data, err := api.fetchData(urlVal)
	if err != nil {
		return EquityListResponse{}, err
	}

	var response EquityListResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return EquityListResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return EquityListResponse{}, err
	}

	return response, nil
}

// ListActiveEquities retrieves a list of all active equities
func (api *TastytradeAPI) ListActiveEquities(params *ActiveEquityQueryParams) (EquityListResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/equities/active", api.host)

	if params != nil {
		queryParams := url.Values{}
		if params.Lendability != "" {
			queryParams.Add("lendability", params.Lendability)
		}
		if params.PerPage != 0 {
			queryParams.Add("per-page", fmt.Sprintf("%d", params.PerPage))
		}
		if params.PageOffset != 0 {
			queryParams.Add("page-offset", fmt.Sprintf("%d", params.PageOffset))
		}
		urlVal = fmt.Sprintf("%s?%s", urlVal, queryParams.Encode())
	}

	data, err := api.fetchData(urlVal)
	if err != nil {
		return EquityListResponse{}, err
	}

	var response EquityListResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return EquityListResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return EquityListResponse{}, err
	}

	return response, nil
}
