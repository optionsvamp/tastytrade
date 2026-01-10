package tastytrade

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// QuoteData represents market data quote information for any security type.
// Different instrument types will have different fields populated.
type QuoteData struct {
	Symbol             string `json:"symbol"`                          // Security symbol
	InstrumentType     string `json:"instrument-type"`                 // Type: "Equity", "Equity Option", "Cryptocurrency", "Index", "Future", "Future Option"
	UpdatedAt          string `json:"updated-at"`                      // Timestamp of last update
	Bid                string `json:"bid,omitempty"`                   // Bid price
	BidSize            string `json:"bid-size,omitempty"`              // Bid size
	Ask                string `json:"ask,omitempty"`                   // Ask price
	AskSize            string `json:"ask-size,omitempty"`              // Ask size
	Mid                string `json:"mid,omitempty"`                   // Mid price (average of bid and ask)
	Mark               string `json:"mark,omitempty"`                  // Mark price
	Last               string `json:"last,omitempty"`                  // Last trade price
	LastMkt            string `json:"last-mkt,omitempty"`              // Last market price
	Beta               string `json:"beta,omitempty"`                  // Beta (for equities)
	DividendAmount     string `json:"dividend-amount,omitempty"`       // Dividend amount (for equities)
	DividendFrequency  string `json:"dividend-frequency,omitempty"`    // Dividend frequency (for equities)
	Open               string `json:"open,omitempty"`                  // Opening price
	DayHighPrice       string `json:"day-high-price,omitempty"`        // Day high price
	DayLowPrice        string `json:"day-low-price,omitempty"`         // Day low price
	Close              string `json:"close,omitempty"`                 // Closing price
	ClosePriceType     string `json:"close-price-type,omitempty"`      // Close price type (e.g., "Final", "Regular")
	PrevClose          string `json:"prev-close,omitempty"`            // Previous close price
	PrevClosePriceType string `json:"prev-close-price-type,omitempty"` // Previous close price type
	SummaryDate        string `json:"summary-date,omitempty"`          // Summary date
	PrevCloseDate      string `json:"prev-close-date,omitempty"`       // Previous close date
	LowLimitPrice      string `json:"low-limit-price,omitempty"`       // Low limit price (for equities)
	HighLimitPrice     string `json:"high-limit-price,omitempty"`      // High limit price (for equities)
	IsTradingHalted    bool   `json:"is-trading-halted,omitempty"`     // Whether trading is halted
	HaltStartTime      int64  `json:"halt-start-time,omitempty"`       // Halt start time (-1 if not halted)
	HaltEndTime        int64  `json:"halt-end-time,omitempty"`         // Halt end time (-1 if not halted)
	YearLowPrice       string `json:"year-low-price,omitempty"`        // Year low price
	YearHighPrice      string `json:"year-high-price,omitempty"`       // Year high price
	Volume             string `json:"volume,omitempty"`                // Trading volume
}

// QuotesResponse represents the response structure returned by GetQuotesByType.
// It contains a list of quote data items and optional pagination information.
type QuotesResponse struct {
	Data struct {
		Items []QuoteData `json:"items"` // Array of quote data
	} `json:"data"`
	Pagination interface{} `json:"pagination"` // Pagination info (null in this endpoint)
}

// QuoteQueryParams represents query parameters for fetching quotes by type.
// Each field accepts a comma-separated list of symbols for that security type.
type QuoteQueryParams struct {
	Cryptocurrency []string `json:"cryptocurrency"` // Comma-separated list of cryptocurrency symbols (e.g., "BTC/USD")
	Equity         []string `json:"equity"`         // Comma-separated list of equity symbols (e.g., "AAPL")
	EquityOption   []string `json:"equity-option"`  // Comma-separated list of equity option symbols (e.g., "SPY 250428P00355000")
	Index          []string `json:"index"`          // Comma-separated list of index symbols (e.g., "SPX")
	Future         []string `json:"future"`         // Comma-separated list of future symbols (e.g., "/CLM5")
	FutureOption   []string `json:"future-option"`  // Comma-separated list of future option symbols (e.g., "/MESU5EX3M5 250620C6450")
}

// GetQuotesByType fetches market data quotes for multiple securities by type.
// This endpoint allows you to fetch quotes for several securities at once by passing
// the security type as the query parameter key and a comma-delimited list of symbols as the value.
// Only available to funded account holders. Does not provide delayed quotes.
//
// Example usage:
//
//	params := &QuoteQueryParams{
//	    Equity: []string{"AAPL", "TSLA"},
//	    Cryptocurrency: []string{"BTC/USD"},
//	}
//	quotes, err := api.GetQuotesByType(params)
//
// Returns a QuotesResponse containing quote data for all requested securities.
func (api *TastytradeAPI) GetQuotesByType(params *QuoteQueryParams) (QuotesResponse, error) {
	urlVal := fmt.Sprintf("%s/market-data/by-type", api.host)

	if params != nil {
		queryParams := url.Values{}
		if len(params.Cryptocurrency) > 0 {
			queryParams.Add("cryptocurrency", strings.Join(params.Cryptocurrency, ","))
		}
		if len(params.Equity) > 0 {
			queryParams.Add("equity", strings.Join(params.Equity, ","))
		}
		if len(params.EquityOption) > 0 {
			queryParams.Add("equity-option", strings.Join(params.EquityOption, ","))
		}
		if len(params.Index) > 0 {
			queryParams.Add("index", strings.Join(params.Index, ","))
		}
		if len(params.Future) > 0 {
			queryParams.Add("future", strings.Join(params.Future, ","))
		}
		if len(params.FutureOption) > 0 {
			queryParams.Add("future-option", strings.Join(params.FutureOption, ","))
		}
		if len(queryParams) > 0 {
			urlVal = fmt.Sprintf("%s?%s", urlVal, queryParams.Encode())
		}
	}

	data, err := api.fetchData(urlVal)
	if err != nil {
		return QuotesResponse{}, err
	}

	var response QuotesResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return QuotesResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return QuotesResponse{}, err
	}

	return response, nil
}

// OptionExpirationImpliedVolatility represents implied volatility data for a specific option expiration.
type OptionExpirationImpliedVolatility struct {
	ExpirationDate    string  `json:"expiration-date"`    // Option expiration date (date-time format)
	SettlementType    string  `json:"settlement-type"`    // Settlement type: "AM" or "PM"
	OptionChainType   string  `json:"option-chain-type"`  // Option chain type: "Standard" or "Non-standard"
	ImpliedVolatility float64 `json:"implied-volatility"` // Implied volatility of option expiration
}

// MarketMetricInfo represents volatility and liquidity data for a symbol.
// Includes underlying volatility and liquidity data as well as option volatility data.
type MarketMetricInfo struct {
	Symbol                              string                              `json:"symbol"`                                 // Symbol
	ImpliedVolatilityIndex              float64                             `json:"implied-volatility-index"`               // IV Index of underlying
	ImpliedVolatilityIndex5DayChange    float64                             `json:"implied-volatility-index-5-day-change"`  // 5 day change of IV index of underlying
	ImpliedVolatilityRank               float64                             `json:"implied-volatility-rank"`                // IV Rank of underlying
	ImpliedVolatilityPercentile         float64                             `json:"implied-volatility-percentile"`          // IV percentile of underlying
	Liquidity                           float64                             `json:"liquidity"`                              // Liquidity of underlying
	LiquidityRank                       float64                             `json:"liquidity-rank"`                         // Liquidity rank of underlying
	LiquidityRating                     int32                               `json:"liquidity-rating"`                       // Liquidity rating of underlying
	OptionExpirationImpliedVolatilities []OptionExpirationImpliedVolatility `json:"option-expiration-implied-volatilities"` // List of option volatility data
}

// DividendInfo represents historical dividend data for a symbol.
type DividendInfo struct {
	OccurredDate string  `json:"occurred-date"` // Date of dividend (YYYY-MM-DD format)
	Amount       float64 `json:"amount"`        // Per share amount
}

// EarningsInfo represents historical earnings data for a symbol.
type EarningsInfo struct {
	OccurredDate string  `json:"occurred-date"` // Date of earnings announcement (YYYY-MM-DD format)
	EPS          float64 `json:"eps"`           // Earnings per share amount
}

// GetMarketMetrics returns an array of volatility data for given symbols.
// symbols is a comma-separated list of symbols (e.g., "AAPL,FB,BRK/B").
// Returns an array of MarketMetricInfo containing volatility and liquidity data
// for each symbol, including option expiration implied volatilities.
func (api *TastytradeAPI) GetMarketMetrics(symbols []string) ([]MarketMetricInfo, error) {
	urlVal := fmt.Sprintf("%s/market-metrics", api.host)

	if len(symbols) > 0 {
		queryParams := url.Values{}
		queryParams.Add("symbols", strings.Join(symbols, ","))
		urlVal = fmt.Sprintf("%s?%s", urlVal, queryParams.Encode())
	}

	data, err := api.fetchData(urlVal)
	if err != nil {
		return nil, err
	}

	// The API may return an array directly or wrapped in a data object
	var response []MarketMetricInfo
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Check if data is wrapped in a "data" key
	if dataValue, ok := data["data"]; ok {
		if dataItems, ok := dataValue.([]interface{}); ok {
			// Wrapped array in "data" key
			itemsData, _ := json.Marshal(dataItems)
			err = json.Unmarshal(itemsData, &response)
		} else {
			// Single object in "data" key - wrap it in an array
			itemsData, _ := json.Marshal([]interface{}{dataValue})
			err = json.Unmarshal(itemsData, &response)
		}
	} else if items, ok := data["items"].([]interface{}); ok {
		// Wrapped in "items" key
		itemsData, _ := json.Marshal(items)
		err = json.Unmarshal(itemsData, &response)
	} else {
		// Try direct array format
		if _, isArray := data[""].(interface{}); !isArray {
			// If the root is an array (not a map), try unmarshaling directly
			var rawArray []interface{}
			if err2 := json.Unmarshal(jsonData, &rawArray); err2 == nil {
				itemsData, _ := json.Marshal(rawArray)
				err = json.Unmarshal(itemsData, &response)
			} else {
				// Last resort: try unmarshaling the whole thing as array
				err = json.Unmarshal(jsonData, &response)
			}
		} else {
			err = json.Unmarshal(jsonData, &response)
		}
	}

	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetHistoricalDividends retrieves historical dividend data for a symbol.
// Returns an array of DividendInfo containing dividend dates and amounts.
func (api *TastytradeAPI) GetHistoricalDividends(symbol string) ([]DividendInfo, error) {
	urlVal := fmt.Sprintf("%s/market-metrics/historic-corporate-events/dividends/%s", api.host, url.PathEscape(symbol))

	data, err := api.fetchData(urlVal)
	if err != nil {
		return nil, err
	}

	var response []DividendInfo
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Check if data is wrapped or direct array
	if dataValue, ok := data["data"]; ok {
		if dataItems, ok := dataValue.([]interface{}); ok {
			// Wrapped array in "data" key
			itemsData, _ := json.Marshal(dataItems)
			err = json.Unmarshal(itemsData, &response)
		} else {
			// Single object in "data" key - wrap it in an array
			itemsData, _ := json.Marshal([]interface{}{dataValue})
			err = json.Unmarshal(itemsData, &response)
		}
	} else if items, ok := data["items"].([]interface{}); ok {
		// Wrapped in "items" key
		itemsData, _ := json.Marshal(items)
		err = json.Unmarshal(itemsData, &response)
	} else {
		// Try direct array format
		var rawArray []interface{}
		if err2 := json.Unmarshal(jsonData, &rawArray); err2 == nil {
			itemsData, _ := json.Marshal(rawArray)
			err = json.Unmarshal(itemsData, &response)
		} else {
			// Last resort: try unmarshaling the whole thing as array
			err = json.Unmarshal(jsonData, &response)
		}
	}

	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetHistoricalEarnings retrieves historical earnings data for a symbol.
// startDate is required and limits earnings data from start-date until now.
// endDate is optional and limits earnings data from start-date until end-date.
// Date format should be YYYY-MM-DD.
// Returns an array of EarningsInfo containing earnings announcement dates and EPS amounts.
func (api *TastytradeAPI) GetHistoricalEarnings(symbol string, startDate string, endDate *string) ([]EarningsInfo, error) {
	urlVal := fmt.Sprintf("%s/market-metrics/historic-corporate-events/earnings-reports/%s", api.host, url.PathEscape(symbol))

	queryParams := url.Values{}
	queryParams.Add("start-date", startDate)
	if endDate != nil && *endDate != "" {
		queryParams.Add("end-date", *endDate)
	}
	urlVal = fmt.Sprintf("%s?%s", urlVal, queryParams.Encode())

	data, err := api.fetchData(urlVal)
	if err != nil {
		return nil, err
	}

	var response []EarningsInfo
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Check if data is wrapped or direct array
	if dataValue, ok := data["data"]; ok {
		if dataItems, ok := dataValue.([]interface{}); ok {
			// Wrapped array in "data" key
			itemsData, _ := json.Marshal(dataItems)
			err = json.Unmarshal(itemsData, &response)
		} else {
			// Single object in "data" key - wrap it in an array
			itemsData, _ := json.Marshal([]interface{}{dataValue})
			err = json.Unmarshal(itemsData, &response)
		}
	} else if items, ok := data["items"].([]interface{}); ok {
		// Wrapped in "items" key
		itemsData, _ := json.Marshal(items)
		err = json.Unmarshal(itemsData, &response)
	} else {
		// Try direct array format
		var rawArray []interface{}
		if err2 := json.Unmarshal(jsonData, &rawArray); err2 == nil {
			itemsData, _ := json.Marshal(rawArray)
			err = json.Unmarshal(itemsData, &response)
		} else {
			// Last resort: try unmarshaling the whole thing as array
			err = json.Unmarshal(jsonData, &response)
		}
	}

	if err != nil {
		return nil, err
	}

	return response, nil
}
