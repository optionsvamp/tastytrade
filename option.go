package tastytrade

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type OptionDataDetailed struct {
	HaltedAt                       string `json:"halted-at"`
	InstrumentType                 string `json:"instrument-type"`
	RootSymbol                     string `json:"root-symbol"`
	Active                         bool   `json:"active"`
	IsClosingOnly                  bool   `json:"is-closing-only"`
	UnderlyingSymbol               string `json:"underlying-symbol"`
	DaysToExpiration               int    `json:"days-to-expiration"`
	ExpirationDate                 string `json:"expiration-date"`
	ExpiresAt                      string `json:"expires-at"`
	ListedMarket                   string `json:"listed-market"`
	StrikePrice                    string `json:"strike-price"`
	OldSecurityNumber              string `json:"old-security-number"`
	OptionType                     string `json:"option-type"`
	MarketTimeInstrumentCollection string `json:"market-time-instrument-collection"`
	Symbol                         string `json:"symbol"`
	StreamerSymbol                 string `json:"streamer-symbol"`
	ExpirationType                 string `json:"expiration-type"`
	SharesPerContract              int    `json:"shares-per-contract"`
	StopsTradingAt                 string `json:"stops-trading-at"`
	ExerciseStyle                  string `json:"exercise-style"`
	SettlementType                 string `json:"settlement-type"`
	OptionChainType                string `json:"option-chain-type"`
}

type OptionChainsDetailedResponse struct {
	Context string `json:"context"`
	Data    struct {
		Items []OptionDataDetailed `json:"items"`
	} `json:"data"`
}

type StrikeNested struct {
	StrikePrice        string `json:"strike-price"`
	Call               string `json:"call"`
	CallStreamerSymbol string `json:"call-streamer-symbol"`
	Put                string `json:"put"`
	PutStreamerSymbol  string `json:"put-streamer-symbol"`
}

type ExpirationNested struct {
	ExpirationType   string         `json:"expiration-type"`
	ExpirationDate   string         `json:"expiration-date"`
	DaysToExpiration int            `json:"days-to-expiration"`
	SettlementType   string         `json:"settlement-type"`
	Strikes          []StrikeNested `json:"strikes"`
}

type OptionChainItemNested struct {
	UnderlyingSymbol  string             `json:"underlying-symbol"`
	RootSymbol        string             `json:"root-symbol"`
	OptionChainType   string             `json:"option-chain-type"`
	SharesPerContract int                `json:"shares-per-contract"`
	Expirations       []ExpirationNested `json:"expirations"`
}

type OptionChainsNestedResponse struct {
	Data struct {
		Items []OptionChainItemNested `json:"items"`
	} `json:"data"`
	Context string `json:"context"`
}

type DeliverableCompact struct {
	ID              int    `json:"id"`
	RootSymbol      string `json:"root-symbol"`
	DeliverableType string `json:"deliverable-type"`
	Description     string `json:"description"`
	Amount          string `json:"amount"`
	Symbol          string `json:"symbol"`
	InstrumentType  string `json:"instrument-type"`
	Percent         string `json:"percent"`
}

type OptionChainItemCompact struct {
	UnderlyingSymbol  string               `json:"underlying-symbol"`
	RootSymbol        string               `json:"root-symbol"`
	OptionChainType   string               `json:"option-chain-type"`
	SettlementType    string               `json:"settlement-type"`
	SharesPerContract int                  `json:"shares-per-contract"`
	ExpirationType    string               `json:"expiration-type"`
	Deliverables      []DeliverableCompact `json:"deliverables"`
	Symbols           []string             `json:"symbols"`
}

type OptionChainsCompactResponse struct {
	Data struct {
		Items []OptionChainItemCompact `json:"items"`
	} `json:"data"`
	Context string `json:"context"`
}

type EquityOptionData struct {
	Symbol                         string `json:"symbol"`
	InstrumentType                 string `json:"instrument-type"`
	Active                         bool   `json:"active"`
	StrikePrice                    string `json:"strike-price"`
	RootSymbol                     string `json:"root-symbol"`
	UnderlyingSymbol               string `json:"underlying-symbol"`
	ExpirationDate                 string `json:"expiration-date"`
	ExerciseStyle                  string `json:"exercise-style"`
	SharesPerContract              int    `json:"shares-per-contract"`
	OptionType                     string `json:"option-type"`
	OptionChainType                string `json:"option-chain-type"`
	ExpirationType                 string `json:"expiration-type"`
	SettlementType                 string `json:"settlement-type"`
	StopsTradingAt                 string `json:"stops-trading-at"`
	MarketTimeInstrumentCollection string `json:"market-time-instrument-collection"`
	DaysToExpiration               int    `json:"days-to-expiration"`
	ExpiresAt                      string `json:"expires-at"`
	IsClosingOnly                  bool   `json:"is-closing-only"`
	StreamerSymbol                 string `json:"streamer-symbol"`
}

type EquityOptionsListResponse struct {
	Context string `json:"context"`
	Data    struct {
		Items []EquityOptionData `json:"items"`
	} `json:"data"`
}

type EquityOptionsQueryParams struct {
	Symbol      []string `json:"symbol"`
	Active      *bool    `json:"active"`
	WithExpired *bool    `json:"with-expired"`
}

type EquityOptionResponse struct {
	Data    EquityOptionData `json:"data"`
	Context string           `json:"context"`
}

// ListOptionsChainsDetailed retrieves option chain data for a specific symbol
func (api *TastytradeAPI) ListOptionsChainsDetailed(symbol string) (OptionChainsDetailedResponse, error) {
	urlVal := fmt.Sprintf("%s/option-chains/%s", api.host, symbol)
	data, err := api.fetchData(urlVal)
	if err != nil {
		return OptionChainsDetailedResponse{}, err
	}

	var response OptionChainsDetailedResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return OptionChainsDetailedResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return OptionChainsDetailedResponse{}, err
	}

	return response, nil
}

// ListOptionChainsNested retrieves nested option chain data for a specific symbol
func (api *TastytradeAPI) ListOptionChainsNested(symbol string) (OptionChainsNestedResponse, error) {
	urlVal := fmt.Sprintf("%s/option-chains/%s/nested", api.host, symbol)
	data, err := api.fetchData(urlVal)
	if err != nil {
		return OptionChainsNestedResponse{}, err
	}

	var response OptionChainsNestedResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return OptionChainsNestedResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return OptionChainsNestedResponse{}, err
	}

	return response, nil
}

// GetOptionChainsCompact retrieves compact option chain data for a specific symbol
func (api *TastytradeAPI) GetOptionChainsCompact(symbol string) (OptionChainsCompactResponse, error) {
	urlVal := fmt.Sprintf("%s/option-chains/%s/compact", api.host, symbol)
	data, err := api.fetchData(urlVal)
	if err != nil {
		return OptionChainsCompactResponse{}, err
	}

	var response OptionChainsCompactResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return OptionChainsCompactResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return OptionChainsCompactResponse{}, err
	}

	return response, nil
}

// GetEquityOptions retrieves a list of equity options
func (api *TastytradeAPI) GetEquityOptions(params *EquityOptionsQueryParams) (EquityOptionsListResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/equity-options", api.host)

	if params != nil {
		queryParams := url.Values{}
		for _, symbol := range params.Symbol {
			queryParams.Add("symbol[]", symbol)
		}
		if params.Active != nil {
			queryParams.Add("active", fmt.Sprintf("%t", *params.Active))
		}
		if params.WithExpired != nil {
			queryParams.Add("with-expired", fmt.Sprintf("%t", *params.WithExpired))
		}
		urlVal = fmt.Sprintf("%s?%s", urlVal, queryParams.Encode())
	}

	data, err := api.fetchData(urlVal)
	if err != nil {
		return EquityOptionsListResponse{}, err
	}

	var response EquityOptionsListResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return EquityOptionsListResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return EquityOptionsListResponse{}, err
	}

	return response, nil
}

// GetEquityOption retrieves data for a specific equity option symbol
func (api *TastytradeAPI) GetEquityOption(symbol string) (EquityOptionResponse, error) {
	url := fmt.Sprintf("%s/instruments/equity-options/%s", api.host, url.PathEscape(symbol))
	data, err := api.fetchData(url)
	if err != nil {
		return EquityOptionResponse{}, err
	}

	var response EquityOptionResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return EquityOptionResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return EquityOptionResponse{}, err
	}

	return response, nil
}
