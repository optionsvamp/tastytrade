package tastytrade

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// BacktestLeg represents a leg in a backtest strategy.
type BacktestLeg struct {
	InstrumentType string  `json:"instrument-type"` // Type: "Equity Option", "Equity", "Future Option", "Future"
	Symbol         string  `json:"symbol"`          // Symbol for the leg
	Action         string  `json:"action"`          // Action: "Buy to Open", "Sell to Open", "Buy to Close", "Sell to Close"
	Quantity       float64 `json:"quantity"`        // Quantity for the leg
}

// BacktestEntryConditions represents entry conditions for a backtest.
type BacktestEntryConditions struct {
	UnderlyingPrice *struct {
		Min *float64 `json:"min,omitempty"` // Minimum underlying price
		Max *float64 `json:"max,omitempty"` // Maximum underlying price
	} `json:"underlying-price,omitempty"`
	DaysToExpiration *struct {
		Min *int `json:"min,omitempty"` // Minimum days to expiration
		Max *int `json:"max,omitempty"` // Maximum days to expiration
	} `json:"days-to-expiration,omitempty"`
	ImpliedVolatility *struct {
		Min *float64 `json:"min,omitempty"` // Minimum implied volatility
		Max *float64 `json:"max,omitempty"` // Maximum implied volatility
	} `json:"implied-volatility,omitempty"`
}

// BacktestExitConditions represents exit conditions for a backtest.
type BacktestExitConditions struct {
	DaysToExpiration *int     `json:"days-to-expiration,omitempty"` // Exit at days to expiration
	ProfitTarget     *float64 `json:"profit-target,omitempty"`      // Profit target percentage
	StopLoss         *float64 `json:"stop-loss,omitempty"`          // Stop loss percentage
	TimeBased        *bool    `json:"time-based,omitempty"`         // Whether exit is time-based
}

// BacktestRequest represents a backtest submission request.
type BacktestRequest struct {
	Symbol          string                   `json:"symbol"`                     // Underlying symbol (e.g., "SPY")
	StartDate       string                   `json:"start-date"`                 // Start date (YYYY-MM-DD format)
	EndDate         string                   `json:"end-date"`                   // End date (YYYY-MM-DD format)
	Legs            []BacktestLeg            `json:"legs"`                       // Array of strategy legs
	EntryConditions *BacktestEntryConditions `json:"entry-conditions,omitempty"` // Entry conditions
	ExitConditions  *BacktestExitConditions  `json:"exit-conditions,omitempty"`  // Exit conditions
}

// BacktestResult represents the result of a backtest.
type BacktestResult struct {
	ID                string  `json:"id"`                // Backtest ID
	Symbol            string  `json:"symbol"`            // Underlying symbol
	StartDate         string  `json:"start-date"`        // Start date
	EndDate           string  `json:"end-date"`          // End date
	TotalTrades       int     `json:"total-trades"`      // Total number of trades
	WinningTrades     int     `json:"winning-trades"`    // Number of winning trades
	LosingTrades      int     `json:"losing-trades"`     // Number of losing trades
	TotalProfitLoss   float64 `json:"total-profit-loss"` // Total profit/loss
	AverageProfitLoss float64 `json:"avg-profit-loss"`   // Average profit/loss per trade
	MaxProfit         float64 `json:"max-profit"`        // Maximum profit
	MaxLoss           float64 `json:"max-loss"`          // Maximum loss
	WinRate           float64 `json:"win-rate"`          // Win rate percentage
	CreatedAt         string  `json:"created-at"`        // Creation timestamp
	Status            string  `json:"status"`            // Status: "completed", "running", "failed"
}

// BacktestResponse represents the response structure returned by backtest endpoints.
type BacktestResponse struct {
	Data    BacktestResult `json:"data"`    // Backtest result data
	Context string         `json:"context"` // API context identifier
}

// BacktestsResponse represents a list of backtest results.
type BacktestsResponse struct {
	Data struct {
		Items []BacktestResult `json:"items"` // Array of backtest results
	} `json:"data"`
	Context string `json:"context"` // API context identifier
}

// postData sends a POST request with JSON body to the specified URL with authorization
func (api *TastytradeAPI) postData(urlVal string, payload interface{}) (map[string]interface{}, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", urlVal, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", api.authToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return nil, fmt.Errorf("client error occurred: status code %d", resp.StatusCode)
	} else if resp.StatusCode >= 500 {
		return nil, fmt.Errorf("server error occurred: status code %d", resp.StatusCode)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

// SubmitBacktest submits a backtest request to the API.
// Returns a BacktestResponse containing the backtest result or status.
func (api *TastytradeAPI) SubmitBacktest(request BacktestRequest) (BacktestResponse, error) {
	urlVal := fmt.Sprintf("%s/backtesting", api.host)

	data, err := api.postData(urlVal, request)
	if err != nil {
		return BacktestResponse{}, err
	}

	var response BacktestResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return BacktestResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return BacktestResponse{}, err
	}

	return response, nil
}

// GetBacktest retrieves a specific backtest by ID.
// Returns a BacktestResponse containing the backtest result.
func (api *TastytradeAPI) GetBacktest(backtestID string) (BacktestResponse, error) {
	urlVal := fmt.Sprintf("%s/backtesting/%s", api.host, url.PathEscape(backtestID))

	data, err := api.fetchData(urlVal)
	if err != nil {
		return BacktestResponse{}, err
	}

	var response BacktestResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return BacktestResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return BacktestResponse{}, err
	}

	return response, nil
}

// ListBacktests retrieves a list of all backtests for the authenticated account.
// Returns a BacktestsResponse containing an array of backtest results.
func (api *TastytradeAPI) ListBacktests() (BacktestsResponse, error) {
	urlVal := fmt.Sprintf("%s/backtesting", api.host)

	data, err := api.fetchData(urlVal)
	if err != nil {
		return BacktestsResponse{}, err
	}

	var response BacktestsResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return BacktestsResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return BacktestsResponse{}, err
	}

	return response, nil
}
