package tastytrade

import (
	"encoding/json"
	"fmt"
)

type TradingStatusData struct {
	AccountNumber                            string `json:"account-number"`
	DayTradeCount                            int    `json:"day-trade-count"`
	EquitiesMarginCalculationType            string `json:"equities-margin-calculation-type"`
	FeeScheduleName                          string `json:"fee-schedule-name"`
	FuturesMarginRateMultiplier              string `json:"futures-margin-rate-multiplier"`
	HasIntradayEquitiesMargin                bool   `json:"has-intraday-equities-margin"`
	ID                                       int    `json:"id"`
	IsAggregatedAtClearing                   bool   `json:"is-aggregated-at-clearing"`
	IsClosed                                 bool   `json:"is-closed"`
	IsClosingOnly                            bool   `json:"is-closing-only"`
	IsCryptocurrencyClosingOnly              bool   `json:"is-cryptocurrency-closing-only"`
	IsCryptocurrencyEnabled                  bool   `json:"is-cryptocurrency-enabled"`
	IsFrozen                                 bool   `json:"is-frozen"`
	IsFullEquityMarginRequired               bool   `json:"is-full-equity-margin-required"`
	IsFuturesClosingOnly                     bool   `json:"is-futures-closing-only"`
	IsFuturesIntraDayEnabled                 bool   `json:"is-futures-intra-day-enabled"`
	IsFuturesEnabled                         bool   `json:"is-futures-enabled"`
	IsInDayTradeEquityMaintenanceCall        bool   `json:"is-in-day-trade-equity-maintenance-call"`
	IsInMarginCall                           bool   `json:"is-in-margin-call"`
	IsPatternDayTrader                       bool   `json:"is-pattern-day-trader"`
	IsRiskReducingOnly                       bool   `json:"is-risk-reducing-only"`
	IsSmallNotionalFuturesIntraDayEnabled    bool   `json:"is-small-notional-futures-intra-day-enabled"`
	IsRollTheDayForwardEnabled               bool   `json:"is-roll-the-day-forward-enabled"`
	AreFarOtmNetOptionsRestricted            bool   `json:"are-far-otm-net-options-restricted"`
	OptionsLevel                             string `json:"options-level"`
	ShortCallsEnabled                        bool   `json:"short-calls-enabled"`
	SmallNotionalFuturesMarginRateMultiplier string `json:"small-notional-futures-margin-rate-multiplier"`
	IsEquityOfferingEnabled                  bool   `json:"is-equity-offering-enabled"`
	IsEquityOfferingClosingOnly              bool   `json:"is-equity-offering-closing-only"`
	EnhancedFraudSafeguardsEnabledAt         string `json:"enhanced-fraud-safeguards-enabled-at"`
	UpdatedAt                                string `json:"updated-at"`
}

type TradingStatusResponse struct {
	Context string            `json:"context"`
	Data    TradingStatusData `json:"data"`
}

// GetAccountTradingStatus retrieves trading status for a specific account
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
