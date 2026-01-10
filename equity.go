package tastytrade

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// EquityResponse represents the response structure returned by GetEquityData.
// It contains equity instrument data and context information.
type EquityResponse struct {
	Context string     `json:"context"` // API context identifier
	Data    EquityData `json:"data"`    // Equity instrument data
}

// Tick represents a tick size with optional threshold.
type Tick struct {
	Threshold string `json:"threshold,omitempty"` // Optional threshold for the tick size
	Value     string `json:"value"`               // Tick size value
}

// EquityData represents equity instrument information returned by equity endpoints.
// It contains details about the equity symbol including trading characteristics,
// market information, and tick sizes.
type EquityData struct {
	Active                         bool   `json:"active"`                            // Whether the equity is currently active
	BorrowRate                     string `json:"borrow-rate"`                       // Stock borrow rate
	Cusip                          string `json:"cusip"`                             // CUSIP identifier
	Description                    string `json:"description"`                       // Full description of the equity
	ID                             int    `json:"id"`                                // Instrument ID
	InstrumentType                 string `json:"instrument-type"`                   // Type of instrument (e.g., "Equity")
	IsClosingOnly                  bool   `json:"is-closing-only"`                   // Whether only closing positions are allowed
	IsETF                          bool   `json:"is-etf"`                            // Whether this is an ETF
	IsFractionalQuantityEligible   bool   `json:"is-fractional-quantity-eligible"`   // Whether fractional shares are allowed
	IsIlliquid                     bool   `json:"is-illiquid"`                       // Whether the equity is considered illiquid
	IsIndex                        bool   `json:"is-index"`                          // Whether this is an index
	IsOptionsClosingOnly           bool   `json:"is-options-closing-only"`           // Whether options are closing-only
	Lendability                    string `json:"lendability"`                       // Stock lendability status
	ListedMarket                   string `json:"listed-market"`                     // Market where the equity is listed
	MarketTimeInstrumentCollection string `json:"market-time-instrument-collection"` // Market time instrument collection identifier
	OptionTickSizes                []Tick `json:"option-tick-sizes"`                 // Tick sizes for options on this equity
	ShortDescription               string `json:"short-description"`                 // Short description of the equity
	StreamerSymbol                 string `json:"streamer-symbol"`                   // Symbol used for streaming quotes
	Symbol                         string `json:"symbol"`                            // Equity symbol
	TickSizes                      []Tick `json:"tick-sizes"`                        // Tick sizes for trading this equity
}

// EquityListResponse represents the response structure returned by ListEquities and ListActiveEquities.
// It contains a list of equity instruments and context information.
type EquityListResponse struct {
	Context string `json:"context"` // API context identifier
	Data    struct {
		Items []EquityData `json:"items"` // Array of equity instruments
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

// GetEquityData retrieves data for a specific equity symbol.
// Returns an EquityResponse containing detailed information about the equity including
// trading characteristics, market information, and tick sizes.
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

// ListEquities retrieves a list of equities based on optional query parameters.
// params can be nil to retrieve all equities, or can filter by symbol, lendability,
// is-index, or is-etf flags.
// Returns an EquityListResponse containing a list of matching equity instruments.
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

// ListActiveEquities retrieves a paginated list of all active equities.
// params can be nil or can specify lendability filter and pagination parameters
// (per-page and page-offset).
// Returns an EquityListResponse containing a list of active equity instruments.
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
