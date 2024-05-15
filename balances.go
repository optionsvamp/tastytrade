package tastytrade

import (
	"fmt"
)

type BalanceData struct {
	AccountNumber                      string  `json:"account-number"`
	CashBalance                        float64 `json:"cash-balance,string"`
	LongEquityValue                    float64 `json:"long-equity-value,string"`
	ShortEquityValue                   float64 `json:"short-equity-value,string"`
	LongDerivativeValue                float64 `json:"long-derivative-value,string"`
	ShortDerivativeValue               float64 `json:"short-derivative-value,string"`
	LongFuturesValue                   float64 `json:"long-futures-value,string"`
	ShortFuturesValue                  float64 `json:"short-futures-value,string"`
	LongFuturesDerivativeValue         float64 `json:"long-futures-derivative-value,string"`
	ShortFuturesDerivativeValue        float64 `json:"short-futures-derivative-value,string"`
	LongMargineableValue               float64 `json:"long-margineable-value,string"`
	ShortMargineableValue              float64 `json:"short-margineable-value,string"`
	MarginEquity                       float64 `json:"margin-equity,string"`
	EquityBuyingPower                  float64 `json:"equity-buying-power,string"`
	DerivativeBuyingPower              float64 `json:"derivative-buying-power,string"`
	DayTradingBuyingPower              float64 `json:"day-trading-buying-power,string"`
	FuturesMarginRequirement           float64 `json:"futures-margin-requirement,string"`
	AvailableTradingFunds              float64 `json:"available-trading-funds,string"`
	MaintenanceRequirement             float64 `json:"maintenance-requirement,string"`
	MaintenanceCallValue               float64 `json:"maintenance-call-value,string"`
	RegTCallValue                      float64 `json:"reg-t-call-value,string"`
	DayTradingCallValue                float64 `json:"day-trading-call-value,string"`
	DayEquityCallValue                 float64 `json:"day-equity-call-value,string"`
	NetLiquidatingValue                float64 `json:"net-liquidating-value,string"`
	CashAvailableToWithdraw            float64 `json:"cash-available-to-withdraw,string"`
	DayTradeExcess                     float64 `json:"day-trade-excess,string"`
	PendingCash                        float64 `json:"pending-cash,string"`
	PendingCashEffect                  string  `json:"pending-cash-effect"`
	LongCryptocurrencyValue            float64 `json:"long-cryptocurrency-value,string"`
	ShortCryptocurrencyValue           float64 `json:"short-cryptocurrency-value,string"`
	CryptocurrencyMarginRequirement    float64 `json:"cryptocurrency-margin-requirement,string"`
	UnsettledCryptocurrencyFiatAmount  float64 `json:"unsettled-cryptocurrency-fiat-amount,string"`
	UnsettledCryptocurrencyFiatEffect  string  `json:"unsettled-cryptocurrency-fiat-effect"`
	ClosedLoopAvailableBalance         float64 `json:"closed-loop-available-balance,string"`
	EquityOfferingMarginRequirement    float64 `json:"equity-offering-margin-requirement,string"`
	LongBondValue                      float64 `json:"long-bond-value,string"`
	BondMarginRequirement              float64 `json:"bond-margin-requirement,string"`
	UsedDerivativeBuyingPower          float64 `json:"used-derivative-buying-power,string"`
	SnapshotDate                       string  `json:"snapshot-date"`
	RegTMarginRequirement              float64 `json:"reg-t-margin-requirement,string"`
	FuturesOvernightMarginRequirement  float64 `json:"futures-overnight-margin-requirement,string"`
	FuturesIntradayMarginRequirement   float64 `json:"futures-intraday-margin-requirement,string"`
	MaintenanceExcess                  float64 `json:"maintenance-excess,string"`
	PendingMarginInterest              float64 `json:"pending-margin-interest,string"`
	EffectiveCryptocurrencyBuyingPower float64 `json:"effective-cryptocurrency-buying-power,string"`
	UpdatedAt                          string  `json:"updated-at"`
}

type BalanceResponse struct {
	Data    BalanceData `json:"data"`
	Context string      `json:"context"`
}

type AccountBalanceSnapshot struct {
	AccountNumber            string  `json:"account-number"`
	CashBalance              float64 `json:"cash-balance,string"`
	LongEquityValue          float64 `json:"long-equity-value,string"`
	ShortEquityValue         float64 `json:"short-equity-value,string"`
	LongDerivativeValue      float64 `json:"long-derivative-value,string"`
	ShortDerivativeValue     float64 `json:"short-derivative-value,string"`
	LongFuturesValue         float64 `json:"long-futures-value,string"`
	ShortFuturesValue        float64 `json:"short-futures-value,string"`
	LongMargineableValue     float64 `json:"long-margineable-value,string"`
	ShortMargineableValue    float64 `json:"short-margineable-value,string"`
	MarginEquity             float64 `json:"margin-equity,string"`
	EquityBuyingPower        float64 `json:"equity-buying-power,string"`
	DerivativeBuyingPower    float64 `json:"derivative-buying-power,string"`
	DayTradingBuyingPower    float64 `json:"day-trading-buying-power,string"`
	FuturesMarginRequirement float64 `json:"futures-margin-requirement,string"`
	AvailableTradingFunds    float64 `json:"available-trading-funds,string"`
	MaintenanceRequirement   float64 `json:"maintenance-requirement,string"`
	MaintenanceCallValue     float64 `json:"maintenance-call-value,string"`
	RegTCallValue            float64 `json:"reg-t-call-value,string"`
	DayTradingCallValue      float64 `json:"day-trading-call-value,string"`
	DayEquityCallValue       float64 `json:"day-equity-call-value,string"`
	NetLiquidatingValue      float64 `json:"net-liquidating-value,string"`
	DayTradeExcess           float64 `json:"day-trade-excess,string"`
	PendingCash              float64 `json:"pending-cash,string"`
	PendingCashEffect        string  `json:"pending-cash-effect"`
	SnapshotDate             string  `json:"snapshot-date"`
	TimeOfDay                string  `json:"time-of-day"`
}

type AccountBalanceSnapshotResponse struct {
	Data struct {
		Items []AccountBalanceSnapshot `json:"items"`
	} `json:"data"`
	Context string `json:"context"`
}

// GetAccountBalances retrieves balances for a specific account
func (api *TastytradeAPI) GetAccountBalances(accountNumber string) (BalanceResponse, error) {
	url := fmt.Sprintf("%s/accounts/%s/balances", api.host, accountNumber)
	var response BalanceResponse
	err := api.fetchDataAndUnmarshal(url, &response)
	if err != nil {
		return BalanceResponse{}, err
	}
	return response, nil
}

// GetAccountBalanceSnapshots retrieves balance snapshots for a specific account
func (api *TastytradeAPI) GetAccountBalanceSnapshots(accountNumber string, snapshotDate string, timeOfDay string) (AccountBalanceSnapshotResponse, error) {
	url := fmt.Sprintf("%s/accounts/%s/balance-snapshots?snapshot-date=%s&time-of-day=%s", api.host, accountNumber, snapshotDate, timeOfDay)
	var response AccountBalanceSnapshotResponse
	err := api.fetchDataAndUnmarshal(url, &response)
	if err != nil {
		return AccountBalanceSnapshotResponse{}, err
	}
	return response, nil
}
