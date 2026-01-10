package tastytrade

import (
	"encoding/json"
	"fmt"
)

// TradingStatusData represents account trading status information.
// It contains comprehensive details about account permissions, restrictions, and trading capabilities.
type TradingStatusData struct {
	AccountNumber                            string `json:"account-number"`                                // Account number
	DayTradeCount                            int    `json:"day-trade-count"`                               // Number of day trades
	EquitiesMarginCalculationType            string `json:"equities-margin-calculation-type"`              // Type of margin calculation for equities
	FeeScheduleName                          string `json:"fee-schedule-name"`                             // Name of the fee schedule
	FuturesMarginRateMultiplier              string `json:"futures-margin-rate-multiplier"`                // Futures margin rate multiplier
	HasIntradayEquitiesMargin                bool   `json:"has-intraday-equities-margin"`                  // Whether intraday equities margin is enabled
	ID                                       int    `json:"id"`                                            // Trading status record ID
	IsAggregatedAtClearing                   bool   `json:"is-aggregated-at-clearing"`                     // Whether account is aggregated at clearing
	IsClosed                                 bool   `json:"is-closed"`                                     // Whether account is closed
	IsClosingOnly                            bool   `json:"is-closing-only"`                               // Whether account is closing-only
	IsCryptocurrencyClosingOnly              bool   `json:"is-cryptocurrency-closing-only"`                // Whether cryptocurrency is closing-only
	IsCryptocurrencyEnabled                  bool   `json:"is-cryptocurrency-enabled"`                     // Whether cryptocurrency trading is enabled
	IsFrozen                                 bool   `json:"is-frozen"`                                     // Whether account is frozen
	IsFullEquityMarginRequired               bool   `json:"is-full-equity-margin-required"`                // Whether full equity margin is required
	IsFuturesClosingOnly                     bool   `json:"is-futures-closing-only"`                       // Whether futures are closing-only
	IsFuturesIntraDayEnabled                 bool   `json:"is-futures-intra-day-enabled"`                  // Whether futures intraday trading is enabled
	IsFuturesEnabled                         bool   `json:"is-futures-enabled"`                            // Whether futures trading is enabled
	IsInDayTradeEquityMaintenanceCall        bool   `json:"is-in-day-trade-equity-maintenance-call"`       // Whether in day trade equity maintenance call
	IsInMarginCall                           bool   `json:"is-in-margin-call"`                             // Whether account is in margin call
	IsPatternDayTrader                       bool   `json:"is-pattern-day-trader"`                         // Whether account has pattern day trader status
	IsRiskReducingOnly                       bool   `json:"is-risk-reducing-only"`                         // Whether only risk-reducing trades are allowed
	IsSmallNotionalFuturesIntraDayEnabled    bool   `json:"is-small-notional-futures-intra-day-enabled"`   // Whether small notional futures intraday is enabled
	IsRollTheDayForwardEnabled               bool   `json:"is-roll-the-day-forward-enabled"`               // Whether roll the day forward is enabled
	AreFarOtmNetOptionsRestricted            bool   `json:"are-far-otm-net-options-restricted"`            // Whether far OTM net options are restricted
	OptionsLevel                             string `json:"options-level"`                                 // Options trading level
	ShortCallsEnabled                        bool   `json:"short-calls-enabled"`                           // Whether short calls are enabled
	SmallNotionalFuturesMarginRateMultiplier string `json:"small-notional-futures-margin-rate-multiplier"` // Small notional futures margin rate multiplier
	IsEquityOfferingEnabled                  bool   `json:"is-equity-offering-enabled"`                    // Whether equity offerings are enabled
	IsEquityOfferingClosingOnly              bool   `json:"is-equity-offering-closing-only"`               // Whether equity offerings are closing-only
	EnhancedFraudSafeguardsEnabledAt         string `json:"enhanced-fraud-safeguards-enabled-at"`          // Timestamp when enhanced fraud safeguards were enabled
	UpdatedAt                                string `json:"updated-at"`                                    // Timestamp of last update
}

// TradingStatusResponse represents the response structure returned by GetAccountTradingStatus.
// It contains comprehensive trading status information for the account.
type TradingStatusResponse struct {
	Context string            `json:"context"` // API context identifier
	Data    TradingStatusData `json:"data"`    // Trading status data
}

// GetAccountTradingStatus retrieves trading status for a specific account.
// Returns a TradingStatusResponse containing detailed information about account permissions,
// restrictions, trading capabilities, and status flags.
func (api *TastytradeAPI) GetAccountTradingStatus(accountNumber string) (TradingStatusResponse, error) {
	url := fmt.Sprintf("%s/accounts/%s/trading-status", api.host, accountNumber)
	data, err := api.fetchData(url)
	if err != nil {
		return TradingStatusResponse{}, err
	}

	var response TradingStatusResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return TradingStatusResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return TradingStatusResponse{}, err
	}

	return response, nil
}
