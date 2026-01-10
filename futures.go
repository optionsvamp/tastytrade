package tastytrade

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// FutureETFEquivalent represents an ETF equivalent for a future contract.
type FutureETFEquivalent struct {
	Symbol        string `json:"symbol"`         // ETF symbol
	ShareQuantity int    `json:"share-quantity"` // Number of shares equivalent
}

// FutureProduct represents a future product with configuration and roll information.
type FutureProduct struct {
	RootSymbol                   string   `json:"root-symbol"`
	Code                         string   `json:"code"`
	Description                  string   `json:"description"`
	ClearingCode                 string   `json:"clearing-code"`
	ClearingExchangeCode         string   `json:"clearing-exchange-code"`
	ClearportCode                string   `json:"clearport-code"`
	LegacyCode                   string   `json:"legacy-code"`
	Exchange                     string   `json:"exchange"`
	LegacyExchangeCode           string   `json:"legacy-exchange-code"`
	ProductType                  string   `json:"product-type"`
	ListedMonths                 []string `json:"listed-months"`
	ActiveMonths                 []string `json:"active-months"`
	NotionalMultiplier           string   `json:"notional-multiplier"`
	TickSize                     string   `json:"tick-size"`
	DisplayFactor                string   `json:"display-factor"`
	StreamerExchangeCode         string   `json:"streamer-exchange-code"`
	SmallNotional                bool     `json:"small-notional"`
	BackMonthFirstCalendarSymbol bool     `json:"back-month-first-calendar-symbol"`
	FirstNotice                  bool     `json:"first-notice"`
	CashSettled                  bool     `json:"cash-settled"`
	SecurityGroup                string   `json:"security-group"`
	MarketSector                 string   `json:"market-sector"`
	Roll                         struct {
		Name               string `json:"name"`
		ActiveCount        int    `json:"active-count"`
		CashSettled        bool   `json:"cash-settled"`
		BusinessDaysOffset int    `json:"business-days-offset"`
		FirstNotice        bool   `json:"first-notice"`
	} `json:"roll"`
}

// TickSize represents a tick size with optional threshold and symbol.
type TickSize struct {
	Value     string `json:"value"`               // Tick size value
	Threshold string `json:"threshold,omitempty"` // Optional threshold for the tick size
	Symbol    string `json:"symbol,omitempty"`    // Optional symbol identifier
}

// Future represents a future contract with comprehensive details.
// It contains information about the contract, product, expiration, and trading characteristics.
type Future struct {
	Symbol                       string              `json:"symbol"`                           // Future contract symbol
	ProductCode                  string              `json:"product-code"`                     // Product code
	ContractSize                 string              `json:"contract-size"`                    // Contract size
	TickSize                     string              `json:"tick-size"`                        // Tick size
	NotionalMultiplier           string              `json:"notional-multiplier"`              // Notional multiplier
	MainFraction                 string              `json:"main-fraction"`                    // Main fraction
	SubFraction                  string              `json:"sub-fraction"`                     // Sub fraction
	DisplayFactor                string              `json:"display-factor"`                   // Display factor
	LastTradeDate                string              `json:"last-trade-date"`                  // Last trade date
	ExpirationDate               string              `json:"expiration-date"`                  // Expiration date
	ClosingOnlyDate              string              `json:"closing-only-date"`                // Closing only date
	Active                       bool                `json:"active"`                           // Whether the future is active
	ActiveMonth                  bool                `json:"active-month"`                     // Whether this is the active month
	NextActiveMonth              bool                `json:"next-active-month"`                // Whether this is the next active month
	IsClosingOnly                bool                `json:"is-closing-only"`                  // Whether only closing positions are allowed
	StopsTradingAt               string              `json:"stops-trading-at"`                 // Time when trading stops
	ExpiresAt                    string              `json:"expires-at"`                       // Expiration timestamp
	ProductGroup                 string              `json:"product-group"`                    // Product group
	Exchange                     string              `json:"exchange"`                         // Exchange where the future trades
	RollTargetSymbol             string              `json:"roll-target-symbol"`               // Target symbol for rolling
	StreamerExchangeCode         string              `json:"streamer-exchange-code"`           // Streamer exchange code
	StreamerSymbol               string              `json:"streamer-symbol"`                  // Symbol used for streaming quotes
	BackMonthFirstCalendarSymbol bool                `json:"back-month-first-calendar-symbol"` // Whether back month uses first calendar symbol
	IsTradeable                  bool                `json:"is-tradeable"`                     // Whether the future is tradeable
	FutureETFEquivalent          FutureETFEquivalent `json:"future-etf-equivalent"`            // ETF equivalent information
	FutureProduct                FutureProduct       `json:"future-product"`                   // Future product information
	TickSizes                    []TickSize          `json:"tick-sizes"`                       // Array of tick sizes
	OptionTickSizes              []TickSize          `json:"option-tick-sizes"`                // Array of option tick sizes
	SpreadTickSizes              []TickSize          `json:"spread-tick-sizes"`                // Array of spread tick sizes
}

// FuturesQueryResponse represents the response structure returned by QueryFutures.
// It contains a list of future contracts matching the query parameters.
type FuturesQueryResponse struct {
	Data struct {
		Items []Future `json:"items"` // Array of future contracts
	} `json:"data"`
}

// FuturesQueryParams represents query parameters for filtering futures.
type FuturesQueryParams struct {
	Symbol      []string `json:"symbol"`       // Array of future symbols to filter by
	ProductCode []string `json:"product-code"` // Array of product codes to filter by
}

// FutureResponse represents the response structure returned by GetFuture.
// It contains detailed information about a specific future contract.
type FutureResponse struct {
	Data    Future `json:"data"`    // Future contract data
	Context string `json:"context"` // API context identifier
}

// FutureOptionProduct represents a future option product configuration.
type FutureOptionProduct struct {
	RootSymbol              string `json:"root-symbol"`               // Root symbol
	CashSettled             bool   `json:"cash-settled"`              // Whether cash settled
	Code                    string `json:"code"`                      // Product code
	LegacyCode              string `json:"legacy-code"`               // Legacy code
	ClearportCode           string `json:"clearport-code"`            // Clearport code
	ClearingCode            string `json:"clearing-code"`             // Clearing code
	ClearingExchangeCode    string `json:"clearing-exchange-code"`    // Clearing exchange code
	ClearingPriceMultiplier string `json:"clearing-price-multiplier"` // Clearing price multiplier
	DisplayFactor           string `json:"display-factor"`            // Display factor
	Exchange                string `json:"exchange"`                  // Exchange
	ProductType             string `json:"product-type"`              // Product type
	ExpirationType          string `json:"expiration-type"`           // Expiration type
	SettlementDelayDays     int    `json:"settlement-delay-days"`     // Settlement delay in days
	IsRollover              bool   `json:"is-rollover"`               // Whether rollover is enabled
	MarketSector            string `json:"market-sector"`             // Market sector
}

// FutureProductsResponse represents the response structure returned by ListFutureProducts.
// It contains a list of future products.
type FutureProductsResponse struct {
	Data struct {
		Items []FutureProduct `json:"items"` // Array of future products
	} `json:"data"`
	Context string `json:"context"` // API context identifier
}

// FutureProductResponse represents the response structure returned by GetFutureProduct.
// It contains detailed information about a specific future product.
type FutureProductResponse struct {
	Data    FutureProduct `json:"data"`    // Future product data
	Context string        `json:"context"` // API context identifier
}

// FuturesOptionExpirationNested represents expiration information in nested future option chain responses.
type FuturesOptionExpirationNested struct {
	UnderlyingSymbol     string         `json:"underlying-symbol"`      // Underlying future symbol
	RootSymbol           string         `json:"root-symbol"`            // Root symbol
	OptionRootSymbol     string         `json:"option-root-symbol"`     // Option root symbol
	OptionContractSymbol string         `json:"option-contract-symbol"` // Option contract symbol
	Asset                string         `json:"asset"`                  // Asset type
	ExpirationDate       string         `json:"expiration-date"`        // Expiration date (YYYY-MM-DD)
	DaysToExpiration     int            `json:"days-to-expiration"`     // Number of days until expiration
	ExpirationType       string         `json:"expiration-type"`        // Expiration type
	SettlementType       string         `json:"settlement-type"`        // Settlement type
	NotionalValue        string         `json:"notional-value"`         // Notional value
	DisplayFactor        string         `json:"display-factor"`         // Display factor
	StrikeFactor         string         `json:"strike-factor"`          // Strike factor
	StopsTradingAt       string         `json:"stops-trading-at"`       // Time when trading stops
	ExpiresAt            string         `json:"expires-at"`             // Expiration timestamp
	TickSizes            []TickSize     `json:"tick-sizes"`             // Array of tick sizes
	Strikes              []StrikeNested `json:"strikes"`                // Array of strike prices (StrikeNested is defined in option.go). For futures options, Call/Put may use different format (e.g., "./ESH6 EW2G6 260213C5200")
}

// OptionChain represents a future option chain with expirations.
type OptionChain struct {
	UnderlyingSymbol string                          `json:"underlying-symbol"` // Underlying future symbol
	RootSymbol       string                          `json:"root-symbol"`       // Root symbol
	ExerciseStyle    string                          `json:"exercise-style"`    // Exercise style: "American" or "European"
	Expirations      []FuturesOptionExpirationNested `json:"expirations"`       // Array of expiration dates with strikes
}

// FutureOptionChainsNestedData contains futures and their associated option chains.
type FutureOptionChainsNestedData struct {
	Futures      []Future      `json:"futures"`       // Array of future contracts
	OptionChains []OptionChain `json:"option-chains"` // Array of option chains for the futures
}

// FutureOptionChainsNestedResponse represents the response structure returned by ListFutureOptionChainsNested.
// It contains a nested structure with futures and their option chains organized by expiration.
type FutureOptionChainsNestedResponse struct {
	Data    FutureOptionChainsNestedData `json:"data"`    // Nested data containing futures and option chains
	Context string                       `json:"context"` // API context identifier
}

// FutureOption represents a future option contract with comprehensive details.
type FutureOption struct {
	Symbol               string              `json:"symbol"`
	UnderlyingSymbol     string              `json:"underlying-symbol"`
	ProductCode          string              `json:"product-code"`
	ExpirationDate       string              `json:"expiration-date"`
	RootSymbol           string              `json:"root-symbol"`
	OptionRootSymbol     string              `json:"option-root-symbol"`
	StrikePrice          string              `json:"strike-price"`
	Exchange             string              `json:"exchange"`
	ExchangeSymbol       string              `json:"exchange-symbol"`
	StreamerSymbol       string              `json:"streamer-symbol"`
	OptionType           string              `json:"option-type"`
	ExerciseStyle        string              `json:"exercise-style"`
	IsVanilla            bool                `json:"is-vanilla"`
	IsPrimaryDeliverable bool                `json:"is-primary-deliverable"`
	FuturePriceRatio     string              `json:"future-price-ratio"`
	Multiplier           string              `json:"multiplier"`
	UnderlyingCount      string              `json:"underlying-count"`
	IsConfirmed          bool                `json:"is-confirmed"`
	NotionalValue        string              `json:"notional-value"`
	DisplayFactor        string              `json:"display-factor"`
	SecurityExchange     string              `json:"security-exchange"`
	SxID                 string              `json:"sx-id"`
	SettlementType       string              `json:"settlement-type"`
	StrikeFactor         string              `json:"strike-factor"`
	MaturityDate         string              `json:"maturity-date"`
	IsExercisableWeekly  bool                `json:"is-exercisable-weekly"`
	LastTradeTime        string              `json:"last-trade-time"`
	DaysToExpiration     int                 `json:"days-to-expiration"`
	IsClosingOnly        bool                `json:"is-closing-only"`
	Active               bool                `json:"active"`
	StopsTradingAt       string              `json:"stops-trading-at"`
	ExpiresAt            string              `json:"expires-at"`
	FutureOptionProduct  FutureOptionProduct `json:"future-option-product"`
}

// FutureOptionChainsDetailedResponse represents the response structure returned by ListFutureOptionChainsDetailed.
// It contains a detailed list of all future option contracts in the chain.
type FutureOptionChainsDetailedResponse struct {
	Data struct {
		Items []FutureOption `json:"items"` // Array of detailed future option contracts
	} `json:"data"`
	Context string `json:"context"` // API context identifier
}

// FutureOptionsDetailedResponse represents the response structure returned by ListFutureOptions.
// It contains a list of future options matching the query parameters.
type FutureOptionsDetailedResponse struct {
	Data struct {
		Items []FutureOption `json:"items"` // Array of future option contracts
	} `json:"data"`
	Context string `json:"context"` // API context identifier
}

// FutureOptionProductsResponse represents the response structure returned by ListFutureOptionProducts.
// It contains a list of future option products.
type FutureOptionProductsResponse struct {
	Data struct {
		Items []FutureOptionProduct `json:"items"` // Array of future option products
	} `json:"data"`
	Context string `json:"context"` // API context identifier
}

// FutureOptionProductDetailedResponse represents the response structure returned by GetFutureOptionProduct.
// It contains detailed information about a specific future option product.
type FutureOptionProductDetailedResponse struct {
	Data    FutureOptionProduct `json:"data"`    // Future option product data
	Context string              `json:"context"` // API context identifier
}

// FutureOptionsQueryParams represents query parameters for filtering future options.
type FutureOptionsQueryParams struct {
	Symbol           []string `json:"symbol"`             // Array of future option symbols to filter by
	OptionRootSymbol string   `json:"option-root-symbol"` // Option root symbol to filter by
	ExpirationDate   string   `json:"expiration-date"`    // Expiration date to filter by
	OptionType       string   `json:"option-type"`        // Option type to filter by (Call/Put)
	StrikePrice      float64  `json:"strike-price"`       // Strike price to filter by
}

// FutureOptionDetailedResponse represents the response structure returned by GetFutureOption.
// It contains detailed information about a specific future option contract.
type FutureOptionDetailedResponse struct {
	Data    FutureOption `json:"data"`    // Future option data
	Context string       `json:"context"` // API context identifier
}

// QueryFutures retrieves a list of futures based on optional query parameters.
// params can be nil to retrieve all futures, or can filter by symbol or product code.
// Returns a FuturesQueryResponse containing a list of matching future contracts.
func (api *TastytradeAPI) QueryFutures(params *FuturesQueryParams) (FuturesQueryResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/futures", api.host)

	if params != nil {
		queryParams := url.Values{}
		for _, symbol := range params.Symbol {
			queryParams.Add("symbol[]", symbol)
		}
		for _, productCode := range params.ProductCode {
			queryParams.Add("product-code[]", productCode)
		}
		urlVal = fmt.Sprintf("%s?%s", urlVal, queryParams.Encode())
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

// GetFuture retrieves data for a specific future symbol.
// Returns a FutureResponse containing detailed information about the future contract
// including expiration, trading characteristics, and product details.
func (api *TastytradeAPI) GetFuture(symbol string) (FutureResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/futures/%s", api.host, url.PathEscape(symbol))
	data, err := api.fetchInstrumentData(urlVal)
	if err != nil {
		return FutureResponse{}, err
	}

	var response FutureResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return FutureResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return FutureResponse{}, err
	}

	return response, nil
}

// ListFutureProducts retrieves a list of all future products.
// Returns a FutureProductsResponse containing all available future products with
// their configuration and roll information.
func (api *TastytradeAPI) ListFutureProducts() (FutureProductsResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/future-products", api.host)
	data, err := api.fetchInstrumentData(urlVal)
	if err != nil {
		return FutureProductsResponse{}, err
	}

	var response FutureProductsResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return FutureProductsResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return FutureProductsResponse{}, err
	}

	return response, nil
}

// GetFutureProduct retrieves data for a specific future product by exchange and symbol.
// Returns a FutureProductResponse containing detailed product configuration including
// roll settings, clearing codes, and trading parameters.
func (api *TastytradeAPI) GetFutureProduct(exchange string, symbol string) (FutureProductResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/future-products/%s/%s", api.host, exchange, url.PathEscape(symbol))
	data, err := api.fetchInstrumentData(urlVal)
	if err != nil {
		return FutureProductResponse{}, err
	}

	var response FutureProductResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return FutureProductResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return FutureProductResponse{}, err
	}

	return response, nil
}

// ListFutureOptionChainsNested retrieves nested future option chain data for a specific future symbol.
// Returns a FutureOptionChainsNestedResponse containing a hierarchical structure organized
// by expiration dates and strikes, making it easier to navigate the option chain.
func (api *TastytradeAPI) ListFutureOptionChainsNested(symbol string) (FutureOptionChainsNestedResponse, error) {
	urlVal := fmt.Sprintf("%s/futures-option-chains/%s/nested", api.host, symbol)
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

// ListFutureOptionChainsDetailed retrieves detailed future option chain data for a specific future symbol.
// Returns a FutureOptionChainsDetailedResponse containing a flat list of all future option
// contracts in the chain with comprehensive details for each contract.
func (api *TastytradeAPI) ListFutureOptionChainsDetailed(symbol string) (FutureOptionChainsDetailedResponse, error) {
	urlVal := fmt.Sprintf("%s/futures-option-chains/%s", api.host, symbol)
	data, err := api.fetchInstrumentData(urlVal)
	if err != nil {
		return FutureOptionChainsDetailedResponse{}, err
	}

	var response FutureOptionChainsDetailedResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return FutureOptionChainsDetailedResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return FutureOptionChainsDetailedResponse{}, err
	}

	return response, nil
}

// ListFutureOptions retrieves future option data based on query parameters.
// params can be nil to retrieve all future options, or can filter by symbol, option root symbol,
// expiration date, option type, and strike price.
// Returns a FutureOptionsDetailedResponse containing matching future option contracts.
func (api *TastytradeAPI) ListFutureOptions(params *FutureOptionsQueryParams) (FutureOptionsDetailedResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/future-options", api.host)

	if params != nil {
		queryParams := url.Values{}
		for _, symbol := range params.Symbol {
			queryParams.Add("symbol[]", symbol)
		}
		if params.OptionRootSymbol != "" {
			queryParams.Add("option-root-symbol", params.OptionRootSymbol)
		}
		if params.ExpirationDate != "" {
			queryParams.Add("expiration-date", params.ExpirationDate)
		}
		if params.OptionType != "" {
			queryParams.Add("option-type", params.OptionType)
		}
		if params.StrikePrice != 0 {
			queryParams.Add("strike-price", fmt.Sprintf("%.2f", params.StrikePrice))
		}
		urlVal = fmt.Sprintf("%s?%s", urlVal, queryParams.Encode())
	}

	data, err := api.fetchInstrumentData(urlVal)
	if err != nil {
		return FutureOptionsDetailedResponse{}, err
	}

	var response FutureOptionsDetailedResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return FutureOptionsDetailedResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return FutureOptionsDetailedResponse{}, err
	}

	return response, nil
}

// GetFutureOption retrieves data for a specific future option symbol.
// Returns a FutureOptionDetailedResponse containing detailed information about the future option
// contract including strike, expiration, exercise style, and settlement details.
func (api *TastytradeAPI) GetFutureOption(symbol string) (FutureOptionDetailedResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/future-options/%s", api.host, url.PathEscape(symbol))
	data, err := api.fetchInstrumentData(urlVal)
	if err != nil {
		return FutureOptionDetailedResponse{}, err
	}

	var response FutureOptionDetailedResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return FutureOptionDetailedResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return FutureOptionDetailedResponse{}, err
	}

	return response, nil
}

// ListFutureOptionProducts retrieves a list of all future option products.
// Returns a FutureOptionProductsResponse containing all available future option products
// with their configuration and trading parameters.
func (api *TastytradeAPI) ListFutureOptionProducts() (FutureOptionProductsResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/future-option-products", api.host)
	data, err := api.fetchInstrumentData(urlVal)
	if err != nil {
		return FutureOptionProductsResponse{}, err
	}

	var response FutureOptionProductsResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return FutureOptionProductsResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return FutureOptionProductsResponse{}, err
	}

	return response, nil
}

// GetFutureOptionProduct retrieves data for a specific future option product by exchange and root symbol.
// Returns a FutureOptionProductDetailedResponse containing detailed product configuration
// including clearing codes, settlement parameters, and expiration settings.
func (api *TastytradeAPI) GetFutureOptionProduct(exchange string, rootSymbol string) (FutureOptionProductDetailedResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/future-option-products/%s/%s", api.host, exchange, rootSymbol)
	data, err := api.fetchInstrumentData(urlVal)
	if err != nil {
		return FutureOptionProductDetailedResponse{}, err
	}

	var response FutureOptionProductDetailedResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return FutureOptionProductDetailedResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return FutureOptionProductDetailedResponse{}, err
	}

	return response, nil
}
