package tastytrade

import (
	"encoding/json"
	"fmt"
)

// Address represents a physical address structure.
type Address struct {
	StreetOne   string `json:"street-one"`   // First line of street address
	City        string `json:"city"`         // City name
	StateRegion string `json:"state-region"` // State or region
	PostalCode  string `json:"postal-code"`  // Postal/ZIP code
	Country     string `json:"country"`      // Country name
	IsForeign   bool   `json:"is-foreign"`   // Whether the address is foreign
	IsDomestic  bool   `json:"is-domestic"`  // Whether the address is domestic
}

// CustomerSuitability represents customer suitability information for trading.
type CustomerSuitability struct {
	ID                                int    `json:"id"`                                   // Suitability record ID
	MaritalStatus                     string `json:"marital-status"`                       // Marital status
	NumberOfDependents                int    `json:"number-of-dependents"`                 // Number of dependents
	EmploymentStatus                  string `json:"employment-status"`                    // Employment status
	Occupation                        string `json:"occupation"`                           // Occupation
	EmployerName                      string `json:"employer-name"`                        // Employer name
	JobTitle                          string `json:"job-title"`                            // Job title
	AnnualNetIncome                   int    `json:"annual-net-income"`                    // Annual net income
	NetWorth                          int    `json:"net-worth"`                            // Total net worth
	LiquidNetWorth                    int    `json:"liquid-net-worth"`                     // Liquid net worth
	StockTradingExperience            string `json:"stock-trading-experience"`             // Stock trading experience level
	CoveredOptionsTradingExperience   string `json:"covered-options-trading-experience"`   // Covered options trading experience level
	UncoveredOptionsTradingExperience string `json:"uncovered-options-trading-experience"` // Uncovered options trading experience level
	FuturesTradingExperience          string `json:"futures-trading-experience"`           // Futures trading experience level
}

// Person represents personal information about a customer.
type Person struct {
	ExternalID         string `json:"external-id"`          // External identifier
	FirstName          string `json:"first-name"`           // First name
	LastName           string `json:"last-name"`            // Last name
	BirthDate          string `json:"birth-date"`           // Date of birth
	CitizenshipCountry string `json:"citizenship-country"`  // Country of citizenship
	USACitizenshipType string `json:"usa-citizenship-type"` // USA citizenship type
	MaritalStatus      string `json:"marital-status"`       // Marital status
	NumberOfDependents int    `json:"number-of-dependents"` // Number of dependents
	EmploymentStatus   string `json:"employment-status"`    // Employment status
	Occupation         string `json:"occupation"`           // Occupation
	EmployerName       string `json:"employer-name"`        // Employer name
	JobTitle           string `json:"job-title"`            // Job title
}

// CustomerData represents comprehensive customer information returned by GetCustomerInfo.
type CustomerData struct {
	ID                              string              `json:"id"`                                  // Customer ID
	FirstName                       string              `json:"first-name"`                          // First name
	LastName                        string              `json:"last-name"`                           // Last name
	Address                         Address             `json:"address"`                             // Physical address
	MailingAddress                  Address             `json:"mailing-address"`                     // Mailing address
	CustomerSuitability             CustomerSuitability `json:"customer-suitability"`                // Suitability information
	USACitizenshipType              string              `json:"usa-citizenship-type"`                // USA citizenship type
	IsForeign                       bool                `json:"is-foreign"`                          // Whether customer is foreign
	MobilePhoneNumber               string              `json:"mobile-phone-number"`                 // Mobile phone number
	Email                           string              `json:"email"`                               // Email address
	TaxNumberType                   string              `json:"tax-number-type"`                     // Type of tax number (SSN, EIN, etc.)
	TaxNumber                       string              `json:"tax-number"`                          // Tax identification number
	BirthDate                       string              `json:"birth-date"`                          // Date of birth
	ExternalID                      string              `json:"external-id"`                         // External identifier
	CitizenshipCountry              string              `json:"citizenship-country"`                 // Country of citizenship
	SubjectToTaxWithholding         bool                `json:"subject-to-tax-withholding"`          // Whether subject to tax withholding
	AgreedToMargining               bool                `json:"agreed-to-margining"`                 // Whether agreed to margining
	AgreedToTerms                   bool                `json:"agreed-to-terms"`                     // Whether agreed to terms
	HasIndustryAffiliation          bool                `json:"has-industry-affiliation"`            // Whether has industry affiliation
	HasPoliticalAffiliation         bool                `json:"has-political-affiliation"`           // Whether has political affiliation
	HasListedAffiliation            bool                `json:"has-listed-affiliation"`              // Whether has listed affiliation
	IsProfessional                  bool                `json:"is-professional"`                     // Whether is a professional trader
	HasDelayedQuotes                bool                `json:"has-delayed-quotes"`                  // Whether has delayed quotes
	HasPendingOrApprovedApplication bool                `json:"has-pending-or-approved-application"` // Whether has pending/approved application
	IdentifiableType                string              `json:"identifiable-type"`                   // Type of identifiable entity
	Person                          Person              `json:"person"`                              // Person information
}

// CustomerResponse represents the response structure returned by GetCustomerInfo.
// It contains customer data and context information.
type CustomerResponse struct {
	Context string       `json:"context"` // API context identifier
	Data    CustomerData `json:"data"`    // Customer information
}

// Account represents account information for a customer.
type Account struct {
	AccountNumber         string `json:"account-number"`          // Account number
	ExternalID            string `json:"external-id"`             // External identifier
	OpenedAt              string `json:"opened-at"`               // Account opening date
	Nickname              string `json:"nickname"`                // Account nickname
	AccountTypeName       string `json:"account-type-name"`       // Type of account
	DayTraderStatus       bool   `json:"day-trader-status"`       // Whether account has day trader status
	IsClosed              bool   `json:"is-closed"`               // Whether account is closed
	IsFirmError           bool   `json:"is-firm-error"`           // Whether account has firm error
	IsFirmProprietary     bool   `json:"is-firm-proprietary"`     // Whether account is firm proprietary
	IsFuturesApproved     bool   `json:"is-futures-approved"`     // Whether account is approved for futures
	IsTestDrive           bool   `json:"is-test-drive"`           // Whether account is a test drive account
	MarginOrCash          string `json:"margin-or-cash"`          // Account type: margin or cash
	IsForeign             bool   `json:"is-foreign"`              // Whether account is foreign
	FundingDate           string `json:"funding-date"`            // Account funding date
	InvestmentObjective   string `json:"investment-objective"`    // Investment objective
	FuturesAccountPurpose string `json:"futures-account-purpose"` // Futures account purpose
	SuitableOptionsLevel  string `json:"suitable-options-level"`  // Suitable options trading level
	CreatedAt             string `json:"created-at"`              // Account creation timestamp
}

// AccountContainer wraps an Account with authority level information.
type AccountContainer struct {
	Account        Account `json:"account"`         // Account information
	AuthorityLevel string  `json:"authority-level"` // Authority level for the account
}

// AccountData contains a list of account containers.
type AccountData struct {
	Items []AccountContainer `json:"items"` // Array of account containers
}

// AccountResponse represents the response structure returned by GetAccount.
// It contains a single account and context information.
type AccountResponse struct {
	Context string  `json:"context"` // API context identifier
	Data    Account `json:"data"`    // Account information
}

// AccountsResponse represents the response structure returned by ListCustomerAccounts.
// It contains a list of accounts and context information.
type AccountsResponse struct {
	Context string      `json:"context"` // API context identifier
	Data    AccountData `json:"data"`    // Account data containing list of accounts
}

// GetCustomerInfo retrieves customer information for the authenticated user.
// Returns a CustomerResponse containing personal information, addresses, suitability data,
// and account preferences.
func (api *TastytradeAPI) GetCustomerInfo() (CustomerResponse, error) {
	url := fmt.Sprintf("%s/customers/me", api.host)
	data, err := api.fetchData(url)
	if err != nil {
		return CustomerResponse{}, err
	}

	var response CustomerResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return CustomerResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return CustomerResponse{}, err
	}

	return response, nil
}

// ListCustomerAccounts retrieves a list of all accounts for the authenticated customer.
// Returns an AccountsResponse containing an array of account containers with account
// information and authority levels.
func (api *TastytradeAPI) ListCustomerAccounts() (AccountsResponse, error) {
	url := fmt.Sprintf("%s/customers/me/accounts", api.host)
	data, err := api.fetchData(url)
	if err != nil {
		return AccountsResponse{}, err
	}

	var response AccountsResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return AccountsResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return AccountsResponse{}, err
	}

	return response, nil
}

// GetAccount retrieves account information for a specific account.
// Returns an AccountResponse containing detailed account information including
// account type, status, trading permissions, and configuration.
func (api *TastytradeAPI) GetAccount(accountNumber string) (AccountResponse, error) {
	url := fmt.Sprintf("%s/customers/me/accounts/%s", api.host, accountNumber)
	data, err := api.fetchData(url)
	if err != nil {
		return AccountResponse{}, err
	}

	var response AccountResponse
	jsonData, err := json.Marshal(data)
	if err != nil {
		return AccountResponse{}, err
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return AccountResponse{}, err
	}

	return response, nil
}
