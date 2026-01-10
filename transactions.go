package tastytrade

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Transaction represents a single transaction in an account.
// It contains details about trades, deposits, withdrawals, fees, and other account activities.
type Transaction struct {
	ID                 int    `json:"id"`                   // Transaction ID
	AccountNumber      string `json:"account-number"`       // Account number
	Symbol             string `json:"symbol"`               // Symbol of the instrument
	InstrumentType     string `json:"instrument-type"`      // Type of instrument (Equity, Option, Future, etc.)
	UnderlyingSymbol   string `json:"underlying-symbol"`    // Underlying symbol for derivatives
	TransactionType    string `json:"transaction-type"`     // Type of transaction (e.g., "Money Movement", "Trade", "Fee", "Deposit")
	TransactionSubType string `json:"transaction-sub-type"` // Subtype of transaction (e.g., "Withdrawal", "Deposit", "Fee")
	Description        string `json:"description"`          // Transaction description
	Action             string `json:"action"`               // Action taken (e.g., "Buy", "Sell", or empty string for non-trade transactions)
	Quantity           string `json:"quantity"`             // Transaction quantity
	Price              string `json:"price"`                // Transaction price
	ExecutedAt         string `json:"executed-at"`          // Execution timestamp
	TransactionDate    string `json:"transaction-date"`     // Transaction date
	Value              string `json:"value"`                // Transaction value
	ValueEffect        string `json:"value-effect"`         // Value effect: "Debit" or "Credit" (example: "Debit")
	NetValue           string `json:"net-value"`            // Net transaction value
	NetValueEffect     string `json:"net-value-effect"`     // Net value effect: "Debit" or "Credit" (example: "Debit")
	IsEstimatedFee     bool   `json:"is-estimated-fee"`     // Whether fee is estimated
}

// TransactionResponse represents the response structure returned by GetTransaction.
// It contains a single transaction and context information.
type TransactionResponse struct {
	Data    Transaction `json:"data"`    // Transaction data
	Context string      `json:"context"` // API context identifier
}

// Pagination contains pagination information for paginated responses.
type Pagination struct {
	PerPage            int     `json:"per-page"`             // Number of items per page
	PageOffset         int     `json:"page-offset"`          // Current page offset
	ItemOffset         int     `json:"item-offset"`          // Current item offset
	TotalItems         int     `json:"total-items"`          // Total number of items
	TotalPages         int     `json:"total-pages"`          // Total number of pages
	CurrentItemCount   int     `json:"current-item-count"`   // Number of items in current page
	PreviousLink       *string `json:"previous-link"`        // Link to previous page (if available)
	NextLink           *string `json:"next-link"`            // Link to next page (if available)
	PagingLinkTemplate *string `json:"paging-link-template"` // Template for paging links
}

// TransactionsResponse represents the response structure returned by GetTransactions.
// It contains a paginated list of transactions, API version, context, and pagination information.
type TransactionsResponse struct {
	Data struct {
		Items []Transaction `json:"items"` // Array of transactions
	} `json:"data"`
	APIVersion string     `json:"api-version"` // API version
	Context    string     `json:"context"`     // API context identifier
	Pagination Pagination `json:"pagination"`  // Pagination information
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

// GetTransactions retrieves transactions for a specific account with optional filtering.
// params can be nil to retrieve all transactions, or can filter by type, date range,
// symbol, instrument type, action, and other criteria.
// Returns a TransactionsResponse containing a paginated list of matching transactions.
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

// GetTransaction retrieves a specific transaction by ID for a specific account.
// Returns a TransactionResponse containing detailed information about the transaction.
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
