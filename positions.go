package tastytrade

import (
	"encoding/json"
	"fmt"
)

type Position struct {
	AccountNumber                 string `json:"account-number"`
	Symbol                        string `json:"symbol"`
	InstrumentType                string `json:"instrument-type"`
	UnderlyingSymbol              string `json:"underlying-symbol"`
	Quantity                      string `json:"quantity"`
	QuantityDirection             string `json:"quantity-direction"`
	ClosePrice                    string `json:"close-price"`
	AverageOpenPrice              string `json:"average-open-price"`
	AverageYearlyMarketClosePrice string `json:"average-yearly-market-close-price"`
	AverageDailyMarketClosePrice  string `json:"average-daily-market-close-price"`
	Multiplier                    int    `json:"multiplier"`
	CostEffect                    string `json:"cost-effect"`
	IsSuppressed                  bool   `json:"is-suppressed"`
	IsFrozen                      bool   `json:"is-frozen"`
	RestrictedQuantity            string `json:"restricted-quantity"`
	RealizedDayGain               string `json:"realized-day-gain"`
	RealizedDayGainEffect         string `json:"realized-day-gain-effect"`
	RealizedDayGainDate           string `json:"realized-day-gain-date"`
	RealizedToday                 string `json:"realized-today"`
	RealizedTodayEffect           string `json:"realized-today-effect"`
	RealizedTodayDate             string `json:"realized-today-date"`
	CreatedAt                     string `json:"created-at"`
	UpdatedAt                     string `json:"updated-at"`
}

type PositionsResponse struct {
	Context string `json:"context"`
	Data    struct {
		Items []Position `json:"items"`
	} `json:"data"`
}

// GetPositions retrieves positions for a specific account
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
