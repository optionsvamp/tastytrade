package tastytrade

import (
	"encoding/json"
	"fmt"
)

// Position represents a trading position in an account.
// It contains information about the position including quantity, prices, P&L, and status.
type Position struct {
	AccountNumber                 string `json:"account-number"`                    // Account number holding the position
	Symbol                        string `json:"symbol"`                            // Symbol of the position
	InstrumentType                string `json:"instrument-type"`                   // Type of instrument (Equity, Option, Future, etc.)
	UnderlyingSymbol              string `json:"underlying-symbol"`                 // Underlying symbol for derivatives
	Quantity                      string `json:"quantity"`                          // Position quantity
	QuantityDirection             string `json:"quantity-direction"`                // Direction: "Long" or "Short"
	ClosePrice                    string `json:"close-price"`                       // Closing price
	AverageOpenPrice              string `json:"average-open-price"`                // Average price at which position was opened
	AverageYearlyMarketClosePrice string `json:"average-yearly-market-close-price"` // Average yearly market close price
	AverageDailyMarketClosePrice  string `json:"average-daily-market-close-price"`  // Average daily market close price
	Multiplier                    int    `json:"multiplier"`                        // Contract multiplier
	CostEffect                    string `json:"cost-effect"`                       // Cost effect: "Debit" or "Credit"
	IsSuppressed                  bool   `json:"is-suppressed"`                     // Whether position is suppressed
	IsFrozen                      bool   `json:"is-frozen"`                         // Whether position is frozen
	RestrictedQuantity            string `json:"restricted-quantity"`               // Quantity that is restricted
	RealizedDayGain               string `json:"realized-day-gain"`                 // Realized gain for the day
	RealizedDayGainEffect         string `json:"realized-day-gain-effect"`          // Effect of day gain: "Debit" or "Credit"
	RealizedDayGainDate           string `json:"realized-day-gain-date"`            // Date of realized day gain
	RealizedToday                 string `json:"realized-today"`                    // Realized P&L for today
	RealizedTodayEffect           string `json:"realized-today-effect"`             // Effect of today's realized P&L: "Debit" or "Credit"
	RealizedTodayDate             string `json:"realized-today-date"`               // Date of today's realized P&L
	CreatedAt                     string `json:"created-at"`                        // Position creation timestamp
	UpdatedAt                     string `json:"updated-at"`                        // Position last update timestamp
}

// PositionsResponse represents the response structure returned by GetPositions.
// It contains a list of positions for the account and context information.
type PositionsResponse struct {
	Context string `json:"context"` // API context identifier
	Data    struct {
		Items []Position `json:"items"` // Array of positions
	} `json:"data"`
}

// GetPositions retrieves all positions for a specific account.
// Returns a PositionsResponse containing a list of all open positions including
// equities, options, futures, and other instruments.
func (api *TastytradeAPI) GetPositions(accountNumber string) (PositionsResponse, error) {
	url := fmt.Sprintf("%s/accounts/%s/positions", api.host, accountNumber)
	data, err := api.fetchData(url)
	if err != nil {
		return PositionsResponse{}, err
	}

	var response PositionsResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return PositionsResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return PositionsResponse{}, err
	}

	return response, nil
}
