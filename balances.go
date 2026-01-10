package tastytrade

import (
	"fmt"
)

// BalanceData represents the account balance information returned by GetAccountBalances.
// It contains comprehensive balance details including cash, equity, derivatives, futures,
// cryptocurrency positions, margin requirements, and buying power calculations.
type BalanceData struct {
	AccountNumber                      string  `json:"account-number"`                               // Account identifier
	CashBalance                        float64 `json:"cash-balance,string"`                          // Current cash balance
	LongEquityValue                    float64 `json:"long-equity-value,string"`                     // Total value of long equity positions
	ShortEquityValue                   float64 `json:"short-equity-value,string"`                    // Total value of short equity positions
	LongDerivativeValue                float64 `json:"long-derivative-value,string"`                 // Total value of long derivative positions
	ShortDerivativeValue               float64 `json:"short-derivative-value,string"`                // Total value of short derivative positions
	LongFuturesValue                   float64 `json:"long-futures-value,string"`                    // Total value of long futures positions
	ShortFuturesValue                  float64 `json:"short-futures-value,string"`                   // Total value of short futures positions
	LongFuturesDerivativeValue         float64 `json:"long-futures-derivative-value,string"`         // Total value of long futures derivative positions
	ShortFuturesDerivativeValue        float64 `json:"short-futures-derivative-value,string"`        // Total value of short futures derivative positions
	LongMargineableValue               float64 `json:"long-margineable-value,string"`                // Total value of long margineable positions
	ShortMargineableValue              float64 `json:"short-margineable-value,string"`               // Total value of short margineable positions
	MarginEquity                       float64 `json:"margin-equity,string"`                         // Margin equity value
	EquityBuyingPower                  float64 `json:"equity-buying-power,string"`                   // Available equity buying power
	DerivativeBuyingPower              float64 `json:"derivative-buying-power,string"`               // Available derivative buying power
	DayTradingBuyingPower              float64 `json:"day-trading-buying-power,string"`              // Available day trading buying power
	FuturesMarginRequirement           float64 `json:"futures-margin-requirement,string"`            // Required margin for futures positions
	AvailableTradingFunds              float64 `json:"available-trading-funds,string"`               // Funds available for trading
	MaintenanceRequirement             float64 `json:"maintenance-requirement,string"`               // Maintenance margin requirement
	MaintenanceCallValue               float64 `json:"maintenance-call-value,string"`                // Maintenance call amount if applicable
	RegTCallValue                      float64 `json:"reg-t-call-value,string"`                      // Reg T call amount if applicable
	DayTradingCallValue                float64 `json:"day-trading-call-value,string"`                // Day trading call amount if applicable
	DayEquityCallValue                 float64 `json:"day-equity-call-value,string"`                 // Day equity call amount if applicable
	NetLiquidatingValue                float64 `json:"net-liquidating-value,string"`                 // Net liquidating value of the account
	CashAvailableToWithdraw            float64 `json:"cash-available-to-withdraw,string"`            // Cash available for withdrawal
	DayTradeExcess                     float64 `json:"day-trade-excess,string"`                      // Day trade excess amount
	PendingCash                        float64 `json:"pending-cash,string"`                          // Pending cash transactions
	PendingCashEffect                  string  `json:"pending-cash-effect"`                          // Effect of pending cash: "Credit", "Debit", or "None" (example: "None")
	LongCryptocurrencyValue            float64 `json:"long-cryptocurrency-value,string"`             // Total value of long cryptocurrency positions
	ShortCryptocurrencyValue           float64 `json:"short-cryptocurrency-value,string"`            // Total value of short cryptocurrency positions
	CryptocurrencyMarginRequirement    float64 `json:"cryptocurrency-margin-requirement,string"`     // Required margin for cryptocurrency positions
	UnsettledCryptocurrencyFiatAmount  float64 `json:"unsettled-cryptocurrency-fiat-amount,string"`  // Unsettled cryptocurrency fiat amount
	UnsettledCryptocurrencyFiatEffect  string  `json:"unsettled-cryptocurrency-fiat-effect"`         // Effect of unsettled cryptocurrency fiat: "Credit", "Debit", or "None" (example: "None")
	ClosedLoopAvailableBalance         float64 `json:"closed-loop-available-balance,string"`         // Closed loop available balance
	EquityOfferingMarginRequirement    float64 `json:"equity-offering-margin-requirement,string"`    // Margin requirement for equity offerings
	LongBondValue                      float64 `json:"long-bond-value,string"`                       // Total value of long bond positions
	BondMarginRequirement              float64 `json:"bond-margin-requirement,string"`               // Required margin for bond positions
	UsedDerivativeBuyingPower          float64 `json:"used-derivative-buying-power,string"`          // Used derivative buying power
	SnapshotDate                       string  `json:"snapshot-date"`                                // Date of the balance snapshot
	RegTMarginRequirement              float64 `json:"reg-t-margin-requirement,string"`              // Reg T margin requirement
	FuturesOvernightMarginRequirement  float64 `json:"futures-overnight-margin-requirement,string"`  // Overnight margin requirement for futures
	FuturesIntradayMarginRequirement   float64 `json:"futures-intraday-margin-requirement,string"`   // Intraday margin requirement for futures
	MaintenanceExcess                  float64 `json:"maintenance-excess,string"`                    // Maintenance excess amount
	PendingMarginInterest              float64 `json:"pending-margin-interest,string"`               // Pending margin interest
	EffectiveCryptocurrencyBuyingPower float64 `json:"effective-cryptocurrency-buying-power,string"` // Effective cryptocurrency buying power
	UpdatedAt                          string  `json:"updated-at"`                                   // Timestamp of last update
}

// BalanceResponse represents the response structure returned by GetAccountBalances.
// It wraps the balance data with a context field.
type BalanceResponse struct {
	Data    BalanceData `json:"data"`    // Balance data for the account
	Context string      `json:"context"` // API context identifier
}

// AccountBalanceSnapshot represents a historical balance snapshot for an account.
// It contains balance information at a specific point in time.
type AccountBalanceSnapshot struct {
	AccountNumber            string  `json:"account-number"`                    // Account identifier
	CashBalance              float64 `json:"cash-balance,string"`               // Cash balance at snapshot time
	LongEquityValue          float64 `json:"long-equity-value,string"`          // Long equity value at snapshot time
	ShortEquityValue         float64 `json:"short-equity-value,string"`         // Short equity value at snapshot time
	LongDerivativeValue      float64 `json:"long-derivative-value,string"`      // Long derivative value at snapshot time
	ShortDerivativeValue     float64 `json:"short-derivative-value,string"`     // Short derivative value at snapshot time
	LongFuturesValue         float64 `json:"long-futures-value,string"`         // Long futures value at snapshot time
	ShortFuturesValue        float64 `json:"short-futures-value,string"`        // Short futures value at snapshot time
	LongMargineableValue     float64 `json:"long-margineable-value,string"`     // Long margineable value at snapshot time
	ShortMargineableValue    float64 `json:"short-margineable-value,string"`    // Short margineable value at snapshot time
	MarginEquity             float64 `json:"margin-equity,string"`              // Margin equity at snapshot time
	EquityBuyingPower        float64 `json:"equity-buying-power,string"`        // Equity buying power at snapshot time
	DerivativeBuyingPower    float64 `json:"derivative-buying-power,string"`    // Derivative buying power at snapshot time
	DayTradingBuyingPower    float64 `json:"day-trading-buying-power,string"`   // Day trading buying power at snapshot time
	FuturesMarginRequirement float64 `json:"futures-margin-requirement,string"` // Futures margin requirement at snapshot time
	AvailableTradingFunds    float64 `json:"available-trading-funds,string"`    // Available trading funds at snapshot time
	MaintenanceRequirement   float64 `json:"maintenance-requirement,string"`    // Maintenance requirement at snapshot time
	MaintenanceCallValue     float64 `json:"maintenance-call-value,string"`     // Maintenance call value at snapshot time
	RegTCallValue            float64 `json:"reg-t-call-value,string"`           // Reg T call value at snapshot time
	DayTradingCallValue      float64 `json:"day-trading-call-value,string"`     // Day trading call value at snapshot time
	DayEquityCallValue       float64 `json:"day-equity-call-value,string"`      // Day equity call value at snapshot time
	NetLiquidatingValue      float64 `json:"net-liquidating-value,string"`      // Net liquidating value at snapshot time
	DayTradeExcess           float64 `json:"day-trade-excess,string"`           // Day trade excess at snapshot time
	PendingCash              float64 `json:"pending-cash,string"`               // Pending cash at snapshot time
	PendingCashEffect        string  `json:"pending-cash-effect"`               // Effect of pending cash (Credit/Debit)
	SnapshotDate             string  `json:"snapshot-date"`                     // Date of the snapshot
	TimeOfDay                string  `json:"time-of-day"`                       // Time of day for the snapshot (e.g., "BOD", "EOD")
}

// AccountBalanceSnapshotResponse represents the response structure returned by GetAccountBalanceSnapshots.
// It contains a list of balance snapshots for the specified account and date.
type AccountBalanceSnapshotResponse struct {
	Data struct {
		Items []AccountBalanceSnapshot `json:"items"` // Array of balance snapshots
	} `json:"data"`
	Context string `json:"context"` // API context identifier
}

// GetAccountBalances retrieves balances for a specific account.
// Returns a BalanceResponse containing current account balance information including
// cash, equity, derivatives, futures, margin requirements, and buying power.
func (api *TastytradeAPI) GetAccountBalances(accountNumber string) (BalanceResponse, error) {
	url := fmt.Sprintf("%s/accounts/%s/balances", api.host, accountNumber)
	var response BalanceResponse
	err := api.fetchDataAndUnmarshal(url, &response)
	if err != nil {
		return BalanceResponse{}, err
	}
	return response, nil
}

// GetAccountBalanceSnapshots retrieves balance snapshots for a specific account.
// snapshotDate should be in YYYY-MM-DD format.
// timeOfDay should be "BOD" (beginning of day) or "EOD" (end of day).
// Returns an AccountBalanceSnapshotResponse containing historical balance snapshots.
func (api *TastytradeAPI) GetAccountBalanceSnapshots(accountNumber string, snapshotDate string, timeOfDay string) (AccountBalanceSnapshotResponse, error) {
	url := fmt.Sprintf("%s/accounts/%s/balance-snapshots?snapshot-date=%s&time-of-day=%s", api.host, accountNumber, snapshotDate, timeOfDay)
	var response AccountBalanceSnapshotResponse
	err := api.fetchDataAndUnmarshal(url, &response)
	if err != nil {
		return AccountBalanceSnapshotResponse{}, err
	}
	return response, nil
}
