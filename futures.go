package tastytrade

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type FutureETFEquivalent struct {
	Symbol        string `json:"symbol"`
	ShareQuantity int    `json:"share-quantity"`
}

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

type TickSize struct {
	Value     string `json:"value"`
	Threshold string `json:"threshold,omitempty"`
	Symbol    string `json:"symbol,omitempty"`
}

type Future struct {
	Symbol                       string              `json:"symbol"`
	ProductCode                  string              `json:"product-code"`
	ContractSize                 string              `json:"contract-size"`
	TickSize                     string              `json:"tick-size"`
	NotionalMultiplier           string              `json:"notional-multiplier"`
	MainFraction                 string              `json:"main-fraction"`
	SubFraction                  string              `json:"sub-fraction"`
	DisplayFactor                string              `json:"display-factor"`
	LastTradeDate                string              `json:"last-trade-date"`
	ExpirationDate               string              `json:"expiration-date"`
	ClosingOnlyDate              string              `json:"closing-only-date"`
	Active                       bool                `json:"active"`
	ActiveMonth                  bool                `json:"active-month"`
	NextActiveMonth              bool                `json:"next-active-month"`
	IsClosingOnly                bool                `json:"is-closing-only"`
	StopsTradingAt               string              `json:"stops-trading-at"`
	ExpiresAt                    string              `json:"expires-at"`
	ProductGroup                 string              `json:"product-group"`
	Exchange                     string              `json:"exchange"`
	RollTargetSymbol             string              `json:"roll-target-symbol"`
	StreamerExchangeCode         string              `json:"streamer-exchange-code"`
	StreamerSymbol               string              `json:"streamer-symbol"`
	BackMonthFirstCalendarSymbol bool                `json:"back-month-first-calendar-symbol"`
	IsTradeable                  bool                `json:"is-tradeable"`
	FutureETFEquivalent          FutureETFEquivalent `json:"future-etf-equivalent"`
	FutureProduct                FutureProduct       `json:"future-product"`
	TickSizes                    []TickSize          `json:"tick-sizes"`
	OptionTickSizes              []TickSize          `json:"option-tick-sizes"`
	SpreadTickSizes              []TickSize          `json:"spread-tick-sizes"`
}

type FuturesQueryResponse struct {
	Data struct {
		Items []Future `json:"items"`
	} `json:"data"`
}

type FuturesQueryParams struct {
	Symbol      []string `json:"symbol"`
	ProductCode []string `json:"product-code"`
}

type FutureResponse struct {
	Data    Future `json:"data"`
	Context string `json:"context"`
}

type FutureOptionProduct struct {
	RootSymbol              string `json:"root-symbol"`
	CashSettled             bool   `json:"cash-settled"`
	Code                    string `json:"code"`
	LegacyCode              string `json:"legacy-code"`
	ClearportCode           string `json:"clearport-code"`
	ClearingCode            string `json:"clearing-code"`
	ClearingExchangeCode    string `json:"clearing-exchange-code"`
	ClearingPriceMultiplier string `json:"clearing-price-multiplier"`
	DisplayFactor           string `json:"display-factor"`
	Exchange                string `json:"exchange"`
	ProductType             string `json:"product-type"`
	ExpirationType          string `json:"expiration-type"`
	SettlementDelayDays     int    `json:"settlement-delay-days"`
	IsRollover              bool   `json:"is-rollover"`
	MarketSector            string `json:"market-sector"`
}

type FutureProductsResponse struct {
	Data struct {
		Items []FutureProduct `json:"items"`
	} `json:"data"`
	Context string `json:"context"`
}

type FutureProductResponse struct {
	Data    FutureProduct `json:"data"`
	Context string        `json:"context"`
}

type FuturesOptionExpirationNested struct {
	UnderlyingSymbol     string         `json:"underlying-symbol"`
	RootSymbol           string         `json:"root-symbol"`
	OptionRootSymbol     string         `json:"option-root-symbol"`
	OptionContractSymbol string         `json:"option-contract-symbol"`
	Asset                string         `json:"asset"`
	ExpirationDate       string         `json:"expiration-date"`
	DaysToExpiration     int            `json:"days-to-expiration"`
	ExpirationType       string         `json:"expiration-type"`
	SettlementType       string         `json:"settlement-type"`
	NotionalValue        string         `json:"notional-value"`
	DisplayFactor        string         `json:"display-factor"`
	StrikeFactor         string         `json:"strike-factor"`
	StopsTradingAt       string         `json:"stops-trading-at"`
	ExpiresAt            string         `json:"expires-at"`
	TickSizes            []TickSize     `json:"tick-sizes"`
	Strikes              []StrikeNested `json:"strikes"`
}

type OptionChain struct {
	UnderlyingSymbol string                          `json:"underlying-symbol"`
	RootSymbol       string                          `json:"root-symbol"`
	ExerciseStyle    string                          `json:"exercise-style"`
	Expirations      []FuturesOptionExpirationNested `json:"expirations"`
}

type FutureOptionChainsNestedData struct {
	Futures      []Future      `json:"futures"`
	OptionChains []OptionChain `json:"option-chains"`
}

type FutureOptionChainsNestedResponse struct {
	Data    FutureOptionChainsNestedData `json:"data"`
	Context string                       `json:"context"`
}

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

type FutureOptionChainsDetailedResponse struct {
	Data struct {
		Items []FutureOption `json:"items"`
	} `json:"data"`
	Context string `json:"context"`
}

type FutureOptionsDetailedResponse struct {
	Data struct {
		Items []FutureOption `json:"items"`
	} `json:"data"`
	Context string `json:"context"`
}

type FutureOptionProductsResponse struct {
	Data struct {
		Items []FutureOptionProduct `json:"items"`
	} `json:"data"`
	Context string `json:"context"`
}

type FutureOptionProductDetailedResponse struct {
	Data    FutureOptionProduct `json:"data"`
	Context string              `json:"context"`
}

type FutureOptionsQueryParams struct {
	Symbol           []string `json:"symbol"`
	OptionRootSymbol string   `json:"option-root-symbol"`
	ExpirationDate   string   `json:"expiration-date"`
	OptionType       string   `json:"option-type"`
	StrikePrice      float64  `json:"strike-price"`
}

type FutureOptionDetailedResponse struct {
	Data    FutureOption `json:"data"`
	Context string       `json:"context"`
}

// QueryFutures retrieves a list of futures
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

	data, err := api.fetchData(urlVal)
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

// GetFuture retrieves data for a specific future symbol
func (api *TastytradeAPI) GetFuture(symbol string) (FutureResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/futures/%s", api.host, url.PathEscape(symbol))
	data, err := api.fetchData(urlVal)
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

// ListFutureProducts retrieves a list of future products
func (api *TastytradeAPI) ListFutureProducts() (FutureProductsResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/future-products", api.host)
	data, err := api.fetchData(urlVal)
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

// GetFutureProduct retrieves data for a specific future product
func (api *TastytradeAPI) GetFutureProduct(exchange string, symbol string) (FutureProductResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/future-products/%s/%s", api.host, exchange, url.PathEscape(symbol))
	data, err := api.fetchData(urlVal)
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

// ListFutureOptionChainsNested retrieves nested future option chain data for a specific symbol
func (api *TastytradeAPI) ListFutureOptionChainsNested(symbol string) (FutureOptionChainsNestedResponse, error) {
	urlVal := fmt.Sprintf("%s/futures-option-chains/%s/nested", api.host, symbol)
	data, err := api.fetchData(urlVal)
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

// ListFutureOptionChainsDetailed retrieves detailed future option chain data for a specific symbol
func (api *TastytradeAPI) ListFutureOptionChainsDetailed(symbol string) (FutureOptionChainsDetailedResponse, error) {
	urlVal := fmt.Sprintf("%s/futures-option-chains/%s", api.host, symbol)
	data, err := api.fetchData(urlVal)
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

// ListFutureOptions retrieves future option data for a specific symbol with query parameters
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

	data, err := api.fetchData(urlVal)
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

// GetFutureOption retrieves data for a specific future option symbol
func (api *TastytradeAPI) GetFutureOption(symbol string) (FutureOptionDetailedResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/future-options/%s", api.host, url.PathEscape(symbol))
	data, err := api.fetchData(urlVal)
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

// ListFutureOptionProducts retrieves a list of future option products
func (api *TastytradeAPI) ListFutureOptionProducts() (FutureOptionProductsResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/future-option-products", api.host)
	data, err := api.fetchData(urlVal)
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

// GetFutureOptionProduct retrieves data for a specific future option product
func (api *TastytradeAPI) GetFutureOptionProduct(exchange string, rootSymbol string) (FutureOptionProductDetailedResponse, error) {
	urlVal := fmt.Sprintf("%s/instruments/future-option-products/%s/%s", api.host, exchange, rootSymbol)
	data, err := api.fetchData(urlVal)
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
