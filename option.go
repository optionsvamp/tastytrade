package tastytrade

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// OptionDataDetailed represents detailed option chain data returned by ListOptionsChainsDetailed.
// It contains comprehensive information about each option contract in the chain.
// Example values:
//
//	Symbol: "AAPL  260116C00005000" (OCC format with spaces)
//	StreamerSymbol: ".AAPL260116C5" (streamer format with dot prefix, no spaces)
//	OptionType: "C" (single letter: "C" for Call, "P" for Put)
//	StrikePrice: "5.0"
//	ExpirationDate: "2026-01-16"
type OptionDataDetailed struct {
	HaltedAt                       string `json:"halted-at"`                         // Time when trading was halted
	InstrumentType                 string `json:"instrument-type"`                   // Type of instrument (e.g., "Equity Option")
	RootSymbol                     string `json:"root-symbol"`                       // Root symbol for the option
	Active                         bool   `json:"active"`                            // Whether the option is currently active
	IsClosingOnly                  bool   `json:"is-closing-only"`                   // Whether only closing positions are allowed
	UnderlyingSymbol               string `json:"underlying-symbol"`                 // Underlying equity symbol
	DaysToExpiration               int    `json:"days-to-expiration"`                // Number of days until expiration
	ExpirationDate                 string `json:"expiration-date"`                   // Expiration date (YYYY-MM-DD)
	ExpiresAt                      string `json:"expires-at"`                        // Expiration timestamp
	ListedMarket                   string `json:"listed-market"`                     // Market where the option is listed
	StrikePrice                    string `json:"strike-price"`                      // Strike price of the option (e.g., "5.0")
	OldSecurityNumber              string `json:"old-security-number"`               // Old security number
	OptionType                     string `json:"option-type"`                       // Type: "C" for Call or "P" for Put
	MarketTimeInstrumentCollection string `json:"market-time-instrument-collection"` // Market time instrument collection identifier
	Symbol                         string `json:"symbol"`                            // Option symbol in OCC format (e.g., "AAPL  260116C00005000")
	StreamerSymbol                 string `json:"streamer-symbol"`                   // Symbol used for streaming quotes with dot prefix (e.g., ".AAPL260116C5")
	ExpirationType                 string `json:"expiration-type"`                   // Expiration type (e.g., "Standard", "Weekly")
	SharesPerContract              int    `json:"shares-per-contract"`               // Number of shares per contract (typically 100)
	StopsTradingAt                 string `json:"stops-trading-at"`                  // Time when trading stops
	ExerciseStyle                  string `json:"exercise-style"`                    // Exercise style: "American" or "European"
	SettlementType                 string `json:"settlement-type"`                   // Settlement type (e.g., "Physical", "Cash")
	OptionChainType                string `json:"option-chain-type"`                 // Type of option chain
}

// OptionChainsDetailedResponse represents the response structure returned by ListOptionsChainsDetailed.
// It contains a detailed list of all option contracts in the chain.
type OptionChainsDetailedResponse struct {
	Context string `json:"context"` // API context identifier
	Data    struct {
		Items []OptionDataDetailed `json:"items"` // Array of detailed option contracts
	} `json:"data"`
}

// StrikeNested represents strike price information in nested option chain responses.
// Example values:
//
//	StrikePrice: "2800.0"
//	Call: "SPXW  260112C02800000" (OCC format with spaces)
//	CallStreamerSymbol: ".SPXW260112C2800" (streamer format with dot prefix, no spaces)
//	Put: "SPXW  260112P02800000" (OCC format with spaces)
//	PutStreamerSymbol: ".SPXW260112P2800" (streamer format with dot prefix, no spaces)
type StrikeNested struct {
	StrikePrice        string `json:"strike-price"`         // Strike price (e.g., "2800.0")
	Call               string `json:"call"`                 // Call option symbol in OCC format (e.g., "SPXW  260112C02800000")
	CallStreamerSymbol string `json:"call-streamer-symbol"` // Call option streamer symbol with dot prefix (e.g., ".SPXW260112C2800")
	Put                string `json:"put"`                  // Put option symbol in OCC format (e.g., "SPXW  260112P02800000")
	PutStreamerSymbol  string `json:"put-streamer-symbol"`  // Put option streamer symbol with dot prefix (e.g., ".SPXW260112P2800")
}

// ExpirationNested represents expiration information in nested option chain responses.
type ExpirationNested struct {
	ExpirationType   string         `json:"expiration-type"`    // Expiration type (e.g., "Standard", "Weekly")
	ExpirationDate   string         `json:"expiration-date"`    // Expiration date (YYYY-MM-DD)
	DaysToExpiration int            `json:"days-to-expiration"` // Number of days until expiration
	SettlementType   string         `json:"settlement-type"`    // Settlement type
	Strikes          []StrikeNested `json:"strikes"`            // Array of strike prices with call/put symbols
}

// OptionChainItemNested represents a nested option chain item with expirations and strikes.
type OptionChainItemNested struct {
	UnderlyingSymbol  string             `json:"underlying-symbol"`   // Underlying equity symbol
	RootSymbol        string             `json:"root-symbol"`         // Root symbol for the option
	OptionChainType   string             `json:"option-chain-type"`   // Type of option chain
	SharesPerContract int                `json:"shares-per-contract"` // Number of shares per contract
	Expirations       []ExpirationNested `json:"expirations"`         // Array of expiration dates with strikes
}

// OptionChainsNestedResponse represents the response structure returned by ListOptionChainsNested.
// It contains a nested structure organizing options by expiration and strike.
type OptionChainsNestedResponse struct {
	Data struct {
		Items []OptionChainItemNested `json:"items"` // Array of nested option chain items
	} `json:"data"`
	Context string `json:"context"` // API context identifier
}

// DeliverableCompact represents deliverable information in compact option chain responses.
type DeliverableCompact struct {
	ID              int    `json:"id"`               // Deliverable ID
	RootSymbol      string `json:"root-symbol"`      // Root symbol
	DeliverableType string `json:"deliverable-type"` // Type of deliverable
	Description     string `json:"description"`      // Description of the deliverable
	Amount          string `json:"amount"`           // Amount of the deliverable
	Symbol          string `json:"symbol"`           // Symbol of the deliverable
	InstrumentType  string `json:"instrument-type"`  // Type of instrument
	Percent         string `json:"percent"`          // Percentage of the deliverable
}

// OptionChainItemCompact represents a compact option chain item with deliverables and symbols.
type OptionChainItemCompact struct {
	UnderlyingSymbol  string               `json:"underlying-symbol"`   // Underlying equity symbol
	RootSymbol        string               `json:"root-symbol"`         // Root symbol for the option
	OptionChainType   string               `json:"option-chain-type"`   // Type of option chain
	SettlementType    string               `json:"settlement-type"`     // Settlement type
	SharesPerContract int                  `json:"shares-per-contract"` // Number of shares per contract
	ExpirationType    string               `json:"expiration-type"`     // Expiration type
	Deliverables      []DeliverableCompact `json:"deliverables"`        // Array of deliverables
	Symbols           []string             `json:"symbols"`             // Array of option symbols
}

// OptionChainsCompactResponse represents the response structure returned by GetOptionChainsCompact.
// It contains a compact representation of the option chain with deliverables and symbols.
type OptionChainsCompactResponse struct {
	Data struct {
		Items []OptionChainItemCompact `json:"items"` // Array of compact option chain items
	} `json:"data"`
	Context string `json:"context"` // API context identifier
}

// EquityOptionData represents equity option instrument information.
// It contains details about a specific equity option contract.
// Example values:
//
//	Symbol: "AAPL  260116C00005000" (OCC format with spaces)
//	StreamerSymbol: ".AAPL260116C5" (streamer format with dot prefix)
//	OptionType: "C" (single letter: "C" for Call, "P" for Put)
type EquityOptionData struct {
	Symbol                         string `json:"symbol"`                            // Option symbol in OCC format (e.g., "AAPL  260116C00005000")
	InstrumentType                 string `json:"instrument-type"`                   // Type of instrument (e.g., "Equity Option")
	Active                         bool   `json:"active"`                            // Whether the option is currently active
	StrikePrice                    string `json:"strike-price"`                      // Strike price (e.g., "5.0")
	RootSymbol                     string `json:"root-symbol"`                       // Root symbol for the option
	UnderlyingSymbol               string `json:"underlying-symbol"`                 // Underlying equity symbol
	ExpirationDate                 string `json:"expiration-date"`                   // Expiration date (YYYY-MM-DD)
	ExerciseStyle                  string `json:"exercise-style"`                    // Exercise style: "American" or "European"
	SharesPerContract              int    `json:"shares-per-contract"`               // Number of shares per contract (typically 100)
	OptionType                     string `json:"option-type"`                       // Type: "C" for Call or "P" for Put
	OptionChainType                string `json:"option-chain-type"`                 // Type of option chain
	ExpirationType                 string `json:"expiration-type"`                   // Expiration type (e.g., "Standard", "Weekly")
	SettlementType                 string `json:"settlement-type"`                   // Settlement type (e.g., "Physical", "Cash")
	StopsTradingAt                 string `json:"stops-trading-at"`                  // Time when trading stops
	MarketTimeInstrumentCollection string `json:"market-time-instrument-collection"` // Market time instrument collection identifier
	DaysToExpiration               int    `json:"days-to-expiration"`                // Number of days until expiration
	ExpiresAt                      string `json:"expires-at"`                        // Expiration timestamp
	IsClosingOnly                  bool   `json:"is-closing-only"`                   // Whether only closing positions are allowed
	StreamerSymbol                 string `json:"streamer-symbol"`                   // Symbol used for streaming quotes with dot prefix (e.g., ".AAPL260116C5")
}

// EquityOptionsListResponse represents the response structure returned by GetEquityOptions.
// It contains a list of equity option instruments matching the query parameters.
type EquityOptionsListResponse struct {
	Context string `json:"context"` // API context identifier
	Data    struct {
		Items []EquityOptionData `json:"items"` // Array of equity option instruments
	} `json:"data"`
}

// EquityOptionsQueryParams represents query parameters for filtering equity options.
type EquityOptionsQueryParams struct {
	Symbol      []string `json:"symbol"`       // Array of option symbols to filter by
	Active      *bool    `json:"active"`       // Filter by active status (nil = no filter)
	WithExpired *bool    `json:"with-expired"` // Include expired options (nil = no filter)
}

// EquityOptionResponse represents the response structure returned by GetEquityOption.
// It contains detailed information about a specific equity option contract.
type EquityOptionResponse struct {
	Data    EquityOptionData `json:"data"`    // Equity option data
	Context string           `json:"context"` // API context identifier
}

// ListOptionsChainsDetailed retrieves detailed option chain data for a specific equity symbol.
// Returns an OptionChainsDetailedResponse containing a flat list of all option contracts
// in the chain with comprehensive details for each contract.
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

// ListOptionChainsNested retrieves nested option chain data for a specific equity symbol.
// Returns an OptionChainsNestedResponse containing a hierarchical structure organized
// by expiration dates and strike prices, making it easier to navigate the chain.
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

// GetOptionChainsCompact retrieves compact option chain data for a specific equity symbol.
// Returns an OptionChainsCompactResponse containing a minimal representation with
// deliverables and option symbols, useful for quick lookups.
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

// GetEquityOptions retrieves a list of equity options based on query parameters.
// params can be nil to retrieve all equity options, or can filter by symbol,
// active status, and whether to include expired options.
// Returns an EquityOptionsListResponse containing matching equity option instruments.
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

// GetEquityOption retrieves data for a specific equity option symbol.
// Returns an EquityOptionResponse containing detailed information about the option
// contract including strike, expiration, exercise style, and settlement details.
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
