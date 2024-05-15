package tastytrade

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Transaction struct {
	ID                 int    `json:"id"`
	AccountNumber      string `json:"account-number"`
	Symbol             string `json:"symbol"`
	InstrumentType     string `json:"instrument-type"`
	UnderlyingSymbol   string `json:"underlying-symbol"`
	TransactionType    string `json:"transaction-type"`
	TransactionSubType string `json:"transaction-sub-type"`
	Description        string `json:"description"`
	Action             string `json:"action"`
	Quantity           string `json:"quantity"`
	Price              string `json:"price"`
	ExecutedAt         string `json:"executed-at"`
	TransactionDate    string `json:"transaction-date"`
	Value              string `json:"value"`
	ValueEffect        string `json:"value-effect"`
	NetValue           string `json:"net-value"`
	NetValueEffect     string `json:"net-value-effect"`
	IsEstimatedFee     bool   `json:"is-estimated-fee"`
}

type TransactionResponse struct {
	Data    Transaction `json:"data"`
	Context string      `json:"context"`
}

type Pagination struct {
	PerPage            int     `json:"per-page"`
	PageOffset         int     `json:"page-offset"`
	ItemOffset         int     `json:"item-offset"`
	TotalItems         int     `json:"total-items"`
	TotalPages         int     `json:"total-pages"`
	CurrentItemCount   int     `json:"current-item-count"`
	PreviousLink       *string `json:"previous-link"`
	NextLink           *string `json:"next-link"`
	PagingLinkTemplate *string `json:"paging-link-template"`
}

type TransactionsResponse struct {
	Data struct {
		Items []Transaction `json:"items"`
	} `json:"data"`
	APIVersion string     `json:"api-version"`
	Context    string     `json:"context"`
	Pagination Pagination `json:"pagination"`
}

type TransactionQueryParams struct {
	Sort             string   `json:"sort"`
	Type             string   `json:"type"`
	SubType          []string `json:"sub-type"`
	Types            []string `json:"types"`
	StartDate        string   `json:"start-date"`
	EndDate          string   `json:"end-date"`
	InstrumentType   string   `json:"instrument-type"`
	Symbol           string   `json:"symbol"`
	UnderlyingSymbol string   `json:"underlying-symbol"`
	Action           string   `json:"action"`
	PartitionKey     string   `json:"partition-key"`
	FuturesSymbol    string   `json:"futures-symbol"`
	StartAt          string   `json:"start-at"`
	EndAt            string   `json:"end-at"`
}

// GetTransactions retrieves transactions for a specific account
func (api *TastytradeAPI) GetTransactions(accountNumber string, params *TransactionQueryParams) (TransactionsResponse, error) {
	urlVal := fmt.Sprintf("%s/accounts/%s/transactions", api.host, accountNumber)

	if params != nil {
		queryParams := url.Values{}
		if params.Sort != "" {
			queryParams.Add("sort", params.Sort)
		}
		if params.Type != "" {
			queryParams.Add("type", params.Type)
		}
		for _, subType := range params.SubType {
			queryParams.Add("sub-type", subType)
		}
		for _, types := range params.Types {
			queryParams.Add("types", types)
		}
		if params.StartDate != "" {
			queryParams.Add("start-date", params.StartDate)
		}
		if params.EndDate != "" {
			queryParams.Add("end-date", params.EndDate)
		}
		if params.InstrumentType != "" {
			queryParams.Add("instrument-type", params.InstrumentType)
		}
		if params.Symbol != "" {
			queryParams.Add("symbol", params.Symbol)
		}
		if params.UnderlyingSymbol != "" {
			queryParams.Add("underlying-symbol", params.UnderlyingSymbol)
		}
		if params.Action != "" {
			queryParams.Add("action", params.Action)
		}
		if params.PartitionKey != "" {
			queryParams.Add("partition-key", params.PartitionKey)
		}
		if params.FuturesSymbol != "" {
			queryParams.Add("futures-symbol", params.FuturesSymbol)
		}
		if params.StartAt != "" {
			queryParams.Add("start-at", params.StartAt)
		}
		if params.EndAt != "" {
			queryParams.Add("end-at", params.EndAt)
		}
		urlVal = fmt.Sprintf("%s?%s", urlVal, queryParams.Encode())
	}

	data, err := api.fetchData(urlVal)
	if err != nil {
		return TransactionsResponse{}, err
	}

	var response TransactionsResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return TransactionsResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return TransactionsResponse{}, err
	}

	return response, nil
}

// GetTransaction retrieves a specific transaction for a specific account
func (api *TastytradeAPI) GetTransaction(accountNumber string, transactionID string) (TransactionResponse, error) {
	urlVal := fmt.Sprintf("%s/accounts/%s/transactions/%s", api.host, accountNumber, transactionID)
	data, err := api.fetchData(urlVal)
	if err != nil {
		return TransactionResponse{}, err
	}

	var response TransactionResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return TransactionResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return TransactionResponse{}, err
	}

	return response, nil
}
