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

// positionRaw is used for flexible unmarshaling of Position fields that may be numbers or strings
type positionRaw struct {
	AccountNumber                 interface{} `json:"account-number"`
	Symbol                        interface{} `json:"symbol"`
	InstrumentType                interface{} `json:"instrument-type"`
	UnderlyingSymbol              interface{} `json:"underlying-symbol"`
	Quantity                      interface{} `json:"quantity"`
	QuantityDirection             interface{} `json:"quantity-direction"`
	ClosePrice                    interface{} `json:"close-price"`
	AverageOpenPrice              interface{} `json:"average-open-price"`
	AverageYearlyMarketClosePrice interface{} `json:"average-yearly-market-close-price"`
	AverageDailyMarketClosePrice  interface{} `json:"average-daily-market-close-price"`
	Multiplier                    interface{} `json:"multiplier"`
	CostEffect                    interface{} `json:"cost-effect"`
	IsSuppressed                  interface{} `json:"is-suppressed"`
	IsFrozen                      interface{} `json:"is-frozen"`
	RestrictedQuantity            interface{} `json:"restricted-quantity"`
	RealizedDayGain               interface{} `json:"realized-day-gain"`
	RealizedDayGainEffect         interface{} `json:"realized-day-gain-effect"`
	RealizedDayGainDate           interface{} `json:"realized-day-gain-date"`
	RealizedToday                 interface{} `json:"realized-today"`
	RealizedTodayEffect           interface{} `json:"realized-today-effect"`
	RealizedTodayDate             interface{} `json:"realized-today-date"`
	CreatedAt                     interface{} `json:"created-at"`
	UpdatedAt                     interface{} `json:"updated-at"`
}

// convertToString converts an interface{} to string, handling both numbers and strings
func convertToString(v interface{}) string {
	if v == nil {
		return ""
	}
	switch val := v.(type) {
	case string:
		return val
	case float64:
		// Remove decimal if it's a whole number
		if val == float64(int64(val)) {
			return fmt.Sprintf("%.0f", val)
		}
		return fmt.Sprintf("%g", val)
	case int:
		return fmt.Sprintf("%d", val)
	case int64:
		return fmt.Sprintf("%d", val)
	default:
		return fmt.Sprintf("%v", val)
	}
}

// convertToInt converts an interface{} to int, handling both numbers and strings
func convertToInt(v interface{}) int {
	if v == nil {
		return 0
	}
	switch val := v.(type) {
	case int:
		return val
	case int64:
		return int(val)
	case float64:
		return int(val)
	case string:
		var result int
		fmt.Sscanf(val, "%d", &result)
		return result
	default:
		return 0
	}
}

// convertToBool converts an interface{} to bool
func convertToBool(v interface{}) bool {
	if v == nil {
		return false
	}
	switch val := v.(type) {
	case bool:
		return val
	case string:
		return val == "true" || val == "1"
	case int:
		return val != 0
	case float64:
		return val != 0
	default:
		return false
	}
}

// convertPositionRaw converts a positionRaw to Position
func convertPositionRaw(raw positionRaw) Position {
	return Position{
		AccountNumber:                 convertToString(raw.AccountNumber),
		Symbol:                        convertToString(raw.Symbol),
		InstrumentType:                convertToString(raw.InstrumentType),
		UnderlyingSymbol:              convertToString(raw.UnderlyingSymbol),
		Quantity:                      convertToString(raw.Quantity),
		QuantityDirection:             convertToString(raw.QuantityDirection),
		ClosePrice:                    convertToString(raw.ClosePrice),
		AverageOpenPrice:              convertToString(raw.AverageOpenPrice),
		AverageYearlyMarketClosePrice: convertToString(raw.AverageYearlyMarketClosePrice),
		AverageDailyMarketClosePrice:  convertToString(raw.AverageDailyMarketClosePrice),
		Multiplier:                    convertToInt(raw.Multiplier),
		CostEffect:                    convertToString(raw.CostEffect),
		IsSuppressed:                  convertToBool(raw.IsSuppressed),
		IsFrozen:                      convertToBool(raw.IsFrozen),
		RestrictedQuantity:            convertToString(raw.RestrictedQuantity),
		RealizedDayGain:               convertToString(raw.RealizedDayGain),
		RealizedDayGainEffect:         convertToString(raw.RealizedDayGainEffect),
		RealizedDayGainDate:           convertToString(raw.RealizedDayGainDate),
		RealizedToday:                 convertToString(raw.RealizedToday),
		RealizedTodayEffect:           convertToString(raw.RealizedTodayEffect),
		RealizedTodayDate:             convertToString(raw.RealizedTodayDate),
		CreatedAt:                     convertToString(raw.CreatedAt),
		UpdatedAt:                     convertToString(raw.UpdatedAt),
	}
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

	// Extract context if available
	var context string
	if ctx, ok := data["context"].(string); ok {
		context = ctx
	}

	// Extract items array
	var items []positionRaw
	if dataValue, ok := data["data"]; ok {
		if dataMap, ok := dataValue.(map[string]interface{}); ok {
			if itemsArray, ok := dataMap["items"].([]interface{}); ok {
				itemsData, _ := json.Marshal(itemsArray)
				json.Unmarshal(itemsData, &items)
			}
		}
	}

	// Convert raw positions to Position structs
	positions := make([]Position, len(items))
	for i, raw := range items {
		positions[i] = convertPositionRaw(raw)
	}

	response := PositionsResponse{
		Context: context,
	}
	response.Data.Items = positions

	return response, nil
}
