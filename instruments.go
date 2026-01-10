package tastytrade

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// QuantityDecimalPrecision represents quantity decimal precision information for an instrument.
type QuantityDecimalPrecision struct {
	Symbol                    string `json:"symbol"`                      // Instrument symbol
	InstrumentType            string `json:"instrument-type"`             // Type of instrument
	Value                     int    `json:"value"`                       // Decimal precision value
	MinimumIncrementPrecision int    `json:"minimum-increment-precision"` // Minimum increment precision
}

// QuantityDecimalPrecisionsResponse represents the response structure returned by GetQuantityDecimalPrecisions.
// It contains a list of quantity decimal precision configurations.
type QuantityDecimalPrecisionsResponse struct {
	Data    []QuantityDecimalPrecision `json:"data"`    // Array of quantity decimal precisions
	Context string                     `json:"context"` // API context identifier
}

// CryptocurrenciesQueryParams represents query parameters for listing cryptocurrencies.
type CryptocurrenciesQueryParams struct {
	Symbol []string `json:"symbol"` // Array of cryptocurrency symbols to filter by
}

// CryptocurrenciesListResponse represents the response structure for listing cryptocurrencies.
// The API returns an array directly, not wrapped in a data object.
type CryptocurrenciesListResponse []Cryptocurrency

// FutureOptionProductsQueryParams represents query parameters for listing future option products with pagination.
type FutureOptionProductsQueryParams struct {
	PageOffset int `json:"page-offset"` // Page offset for pagination (default: 0)
	PerPage    int `json:"per-page"`    // Number of items per page (default: 1000, max: 1000)
}

// FutureProductsQueryParams represents query parameters for listing future products.
type FutureProductsQueryParams struct {
	PageOffset int `json:"page-offset"` // Page offset for pagination (default: 0)
	PerPage    int `json:"per-page"`    // Number of items per page (default: 1000, max: 1000)
}

// FuturesQueryParamsV2 represents query parameters for querying futures (updated API version).
type FuturesQueryParamsV2 struct {
	Symbol            []string `json:"symbol"`              // Array of future symbols (e.g., "ESZ9")
	ProductCode       []string `json:"product-code"`        // Array of product codes (e.g., "ES", "6A")
	SecurityID        []string `json:"security-id"`         // Array of exchange-specific security IDs
	Exchange          string   `json:"exchange"`            // Exchange name (only used to avoid security id collisions)
	OnlyActiveFutures *bool    `json:"only-active-futures"` // If true (defaults to true), only active futures are returned
	PageOffset        int      `json:"page-offset"`         // Page offset for pagination
	PerPage           int      `json:"per-page"`            // Number of items per page (default: 1000, max: 1000)
}

// GetQuantityDecimalPrecisions retrieves all quantity decimal precisions for instruments.
// Returns a QuantityDecimalPrecisionsResponse containing precision configurations for all instrument types.
func (api *TastytradeAPI) GetQuantityDecimalPrecisions() (QuantityDecimalPrecisionsResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/quantity-decimal-precisions", api.host)

	data, err := api.fetchInstrumentData(urlVal)
	if err != nil {
		return QuantityDecimalPrecisionsResponse{}, err
	}

	var response QuantityDecimalPrecisionsResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return QuantityDecimalPrecisionsResponse{}, err
	}

	// Check if data is wrapped or direct array
	if dataValue, ok := data["data"]; ok {
		if dataItems, ok := dataValue.([]interface{}); ok {
			// Wrapped array in "data" key
			itemsData, _ := json.Marshal(dataItems)
			err = json.Unmarshal(itemsData, &response.Data)
		} else {
			// Single object in "data" key - wrap it in an array
			itemsData, _ := json.Marshal([]interface{}{dataValue})
			err = json.Unmarshal(itemsData, &response.Data)
		}
		if context, ok := data["context"].(string); ok {
			response.Context = context
		}
	} else if items, ok := data["items"].([]interface{}); ok {
		// Wrapped in "items" key
		itemsData, _ := json.Marshal(items)
		err = json.Unmarshal(itemsData, &response.Data)
		if context, ok := data["context"].(string); ok {
			response.Context = context
		}
	} else {
		// Try direct array format
		var rawArray []interface{}
		if err2 := json.Unmarshal(jsonData, &rawArray); err2 == nil {
			itemsData, _ := json.Marshal(rawArray)
			err = json.Unmarshal(itemsData, &response.Data)
		} else {
			// Try wrapped format
			var wrapped struct {
				Data    []QuantityDecimalPrecision `json:"data"`
				Context string                     `json:"context"`
			}
			if err2 := json.Unmarshal(jsonData, &wrapped); err2 == nil {
				response.Data = wrapped.Data
				response.Context = wrapped.Context
				err = nil
			} else {
				// Last resort: try unmarshaling the whole thing as array
				err = json.Unmarshal(jsonData, &response.Data)
			}
		}
	}

	if err != nil {
		return QuantityDecimalPrecisionsResponse{}, err
	}

	if context, ok := data["context"].(string); ok {
		response.Context = context
	}

	return response, nil
}

// ListCryptocurrenciesWithParams retrieves a list of cryptocurrencies with optional symbol filtering.
// params can be nil to retrieve all cryptocurrencies, or can filter by symbol array.
// Returns a CryptocurrenciesListResponse containing matching cryptocurrency instruments.
func (api *TastytradeAPI) ListCryptocurrenciesWithParams(params *CryptocurrenciesQueryParams) (CryptocurrenciesListResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/cryptocurrencies", api.host)

	if params != nil && len(params.Symbol) > 0 {
		queryParams := url.Values{}
		for _, symbol := range params.Symbol {
			queryParams.Add("symbol[]", symbol)
		}
		urlVal = fmt.Sprintf("%s?%s", urlVal, queryParams.Encode())
	}

	data, err := api.fetchInstrumentData(urlVal)
	if err != nil {
		return nil, err
	}

	var response CryptocurrenciesListResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Check if data is wrapped or direct array
	if items, ok := data["items"].([]interface{}); ok {
		itemsData, _ := json.Marshal(items)
		err = json.Unmarshal(itemsData, &response)
	} else {
		// Direct array format
		err = json.Unmarshal(jsonData, &response)
		if err != nil {
			// Try wrapped format
			var wrapped struct {
				Data struct {
					Items []Cryptocurrency `json:"items"`
				} `json:"data"`
			}
			if err2 := json.Unmarshal(jsonData, &wrapped); err2 == nil {
				response = wrapped.Data.Items
				err = nil
			}
		}
	}

	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetEquityOptionWithActive retrieves data for a specific equity option by symbol with optional active filter.
// symbol should be in OCC format (e.g., "FB    180629C00200000").
// active is optional and filters for options available for trading (defaults to filtering non-standard/flex options).
// Returns an EquityOptionResponse containing detailed information about the equity option.
func (api *TastytradeAPI) GetEquityOptionWithActive(symbol string, active *bool) (EquityOptionResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/equity-options/%s", api.host, url.PathEscape(symbol))

	if active != nil {
		queryParams := url.Values{}
		queryParams.Add("active", strconv.FormatBool(*active))
		urlVal = fmt.Sprintf("%s?%s", urlVal, queryParams.Encode())
	}

	data, err := api.fetchInstrumentData(urlVal)
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

// ListFutureOptionProductsWithPagination retrieves metadata for all supported future option products with pagination.
// params can be nil to use defaults, or can specify page-offset and per-page.
// Returns a FutureOptionProductsResponse containing a list of future option products.
func (api *TastytradeAPI) ListFutureOptionProductsWithPagination(params *FutureOptionProductsQueryParams) (FutureOptionProductsResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/future-option-products", api.host)

	if params != nil {
		queryParams := url.Values{}
		if params.PageOffset > 0 {
			queryParams.Add("page-offset", strconv.Itoa(params.PageOffset))
		}
		if params.PerPage > 0 {
			queryParams.Add("per-page", strconv.Itoa(params.PerPage))
		}
		if len(queryParams) > 0 {
			urlVal = fmt.Sprintf("%s?%s", urlVal, queryParams.Encode())
		}
	}

	data, err := api.fetchInstrumentData(urlVal)
	if err != nil {
		return FutureOptionProductsResponse{}, err
	}

	var response FutureOptionProductsResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return FutureOptionProductsResponse{}, err
	}

	// Check if data is wrapped or direct array
	if items, ok := data["items"].([]interface{}); ok {
		itemsData, _ := json.Marshal(items)
		err = json.Unmarshal(itemsData, &response.Data.Items)
	} else {
		// Direct array format
		err = json.Unmarshal(jsonData, &response.Data.Items)
		if err != nil {
			// Try wrapped format
			var wrapped struct {
				Data struct {
					Items []FutureOptionProduct `json:"items"`
				} `json:"data"`
				Context string `json:"context"`
			}
			if err2 := json.Unmarshal(jsonData, &wrapped); err2 == nil {
				response.Data.Items = wrapped.Data.Items
				response.Context = wrapped.Context
				err = nil
			}
		}
	}

	if err != nil {
		return FutureOptionProductsResponse{}, err
	}

	if context, ok := data["context"].(string); ok {
		response.Context = context
	}

	return response, nil
}

// ListFutureProductsWithPagination retrieves metadata for all supported future products with pagination.
// params can be nil to use defaults, or can specify page-offset and per-page.
// Returns a FutureProductsResponse containing a list of future products.
func (api *TastytradeAPI) ListFutureProductsWithPagination(params *FutureProductsQueryParams) (FutureProductsResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/future-products", api.host)

	if params != nil {
		queryParams := url.Values{}
		if params.PageOffset > 0 {
			queryParams.Add("page-offset", strconv.Itoa(params.PageOffset))
		}
		if params.PerPage > 0 {
			queryParams.Add("per-page", strconv.Itoa(params.PerPage))
		}
		if len(queryParams) > 0 {
			urlVal = fmt.Sprintf("%s?%s", urlVal, queryParams.Encode())
		}
	}

	data, err := api.fetchInstrumentData(urlVal)
	if err != nil {
		return FutureProductsResponse{}, err
	}

	var response FutureProductsResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return FutureProductsResponse{}, err
	}

	// Check if data is wrapped or direct array
	if items, ok := data["items"].([]interface{}); ok {
		itemsData, _ := json.Marshal(items)
		err = json.Unmarshal(itemsData, &response.Data.Items)
	} else {
		// Direct array format
		err = json.Unmarshal(jsonData, &response.Data.Items)
		if err != nil {
			// Try wrapped format
			var wrapped struct {
				Data struct {
					Items []FutureProduct `json:"items"`
				} `json:"data"`
				Context string `json:"context"`
			}
			if err2 := json.Unmarshal(jsonData, &wrapped); err2 == nil {
				response.Data.Items = wrapped.Data.Items
				response.Context = wrapped.Context
				err = nil
			}
		}
	}

	if err != nil {
		return FutureProductsResponse{}, err
	}

	if context, ok := data["context"].(string); ok {
		response.Context = context
	}

	return response, nil
}

// QueryFuturesV2 retrieves a list of futures with enhanced query parameters (updated API version).
// params can be nil to retrieve all futures, or can filter by symbol, product-code, security-id,
// exchange, and only-active-futures flag.
// Returns a FuturesQueryResponse containing a list of matching future contracts.
func (api *TastytradeAPI) QueryFuturesV2(params *FuturesQueryParamsV2) (FuturesQueryResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/futures", api.host)

	if params != nil {
		queryParams := url.Values{}
		for _, symbol := range params.Symbol {
			queryParams.Add("symbol[]", symbol)
		}
		for _, productCode := range params.ProductCode {
			queryParams.Add("product-code[]", productCode)
		}
		for _, securityID := range params.SecurityID {
			queryParams.Add("security-id[]", securityID)
		}
		if params.Exchange != "" {
			queryParams.Add("exchange", params.Exchange)
		}
		if params.OnlyActiveFutures != nil {
			queryParams.Add("only-active-futures", strconv.FormatBool(*params.OnlyActiveFutures))
		}
		if params.PageOffset > 0 {
			queryParams.Add("page-offset", strconv.Itoa(params.PageOffset))
		}
		if params.PerPage > 0 {
			queryParams.Add("per-page", strconv.Itoa(params.PerPage))
		}
		if len(queryParams) > 0 {
			urlVal = fmt.Sprintf("%s?%s", urlVal, queryParams.Encode())
		}
	}

	data, err := api.fetchInstrumentData(urlVal)
	if err != nil {
		return FuturesQueryResponse{}, err
	}

	var response FuturesQueryResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return FuturesQueryResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return FuturesQueryResponse{}, err
	}

	return response, nil
}

// GetOptionChainBySymbolID retrieves an option chain given an underlying symbol ID (integer).
// symbolID is the integer ID of the underlying symbol (not the symbol string).
// Returns an OptionChainsDetailedResponse containing all option contracts in the chain.
func (api *TastytradeAPI) GetOptionChainBySymbolID(symbolID int) (OptionChainsDetailedResponse, error) {
	urlVal := fmt.Sprintf("%s/option-chains/%d", api.host, symbolID)

	data, err := api.fetchInstrumentData(urlVal)
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

// GetFuturesOptionChainBySymbolID retrieves a futures option chain given a futures product code ID (integer).
// symbolID is the integer ID of the futures product code (not the symbol string).
// Returns a FutureOptionChainsNestedResponse containing the futures option chain.
func (api *TastytradeAPI) GetFuturesOptionChainBySymbolID(symbolID int) (FutureOptionChainsNestedResponse, error) {
	urlVal := fmt.Sprintf("%s/futures-option-chains/%d", api.host, symbolID)

	data, err := api.fetchInstrumentData(urlVal)
	if err != nil {
		return FutureOptionChainsNestedResponse{}, err
	}

	var response FutureOptionChainsNestedResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return FutureOptionChainsNestedResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return FutureOptionChainsNestedResponse{}, err
	}

	return response, nil
}
